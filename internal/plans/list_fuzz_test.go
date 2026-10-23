package plans

import (
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
)

type fuzzPlanService struct {
	fakes.PlanService
	plans []cherrygo.Plan
}

func (s *fuzzPlanService) List(_ context.Context, _ int, _ *cherrygo.GetOptions) ([]cherrygo.Plan, *cherrygo.Response, error) {
	return s.plans, nil, nil
}

type discardPlanOutput struct{}

func (discardPlanOutput) Output(any, []string, *[][]string) error {
	return nil
}

func (discardPlanOutput) SetFormat(outputs.Format) {}

type fuzzPlanDeps struct {
	svc *fuzzPlanService
}

func (d fuzzPlanDeps) Client() cherrygo.PlansService { return d.svc }
func (d fuzzPlanDeps) GetOpts() *cherrygo.GetOptions { return nil }
func (d fuzzPlanDeps) Outputer() outputs.Outputer    { return discardPlanOutput{} }

// FuzzPlanListResponse checks that decoded responses do not cause a panic
func FuzzPlanListResponse(f *testing.F) {
	for _, seed := range []struct {
		response string
		region   string
	}{
		{`null`, ""},
		{`[]`, ""},
		{`[{}]`, ""},
		{`[{"slug":"test","available_regions":[{"id":1,"slug":"lt","stock_qty":2,"spot_qty":1}],"pricing":[{"unit":"Hourly","price":1},{"unit":"Spot hourly","price":0.5}]}]`, ""},
		{`[{"slug":"test","available_regions":[{"id":1,"slug":"lt"},{"id":2,"slug":"nl"}]}]`, "lt"},
		{`[{"slug":"test","available_regions":[{"id":1,"slug":"lt"}]}]`, "1"},
		{`[{"slug":"test","available_regions":[{"id":1,"slug":"lt"}]}]`, "missing"},
		{`[{"available_regions":[{}]}]`, ""},
		{`[{"available_regions":[{"stock_qty":1}]}]`, "lt"},
		{`[{"available_regions":[null]}]`, ""},
		{`[{"slug":"test","available_regions":[{"id":1,"slug":"lt"},{}]}]`, ""},
	} {
		f.Add([]byte(seed.response), seed.region)
	}

	f.Fuzz(func(t *testing.T, response []byte, region string) {
		var plans []cherrygo.Plan
		if err := json.Unmarshal(response, &plans); err != nil {
			return // Invalid JSON or a type mismatch is an expected rejection.
		}

		missingRegion := false
		for _, plan := range plans {
			for _, available := range plan.AvailableRegions {
				if available.Region == nil {
					missingRegion = true
				}
			}
		}

		svc := &fuzzPlanService{plans: plans}
		cmd := NewCommand(fuzzPlanDeps{svc: svc}).list()
		cmd.SetArgs([]string{"--team-id=1", "--region=" + region})
		cmd.SetOut(io.Discard)
		cmd.SetErr(io.Discard)
		cmd.SilenceUsage, cmd.SilenceErrors = true, true
		if err := cmd.Execute(); err != nil && !missingRegion {
			t.Fatalf("complete response was rejected: %v", err)
		}
	})
}
