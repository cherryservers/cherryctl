package plans

import (
	"fmt"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
)

type fakeDeps struct {
	svc *fakes.PlanService
	out *fakes.Outputer
}

func (td fakeDeps) GetOpts() *cherrygo.GetOptions {
	return &cherrygo.GetOptions{}
}

func (fd fakeDeps) Client() cherrygo.PlansService {
	return fd.svc
}

func (fd fakeDeps) Outputer() outputs.Outputer {
	return fd.out
}

func TestList(t *testing.T) {
	cases := []struct {
		title            string
		args             []string
		wantClientParams []any
	}{
		{
			title:            "only team-id",
			args:             []string{"--team-id", "1"},
			wantClientParams: []any{1, &cherrygo.GetOptions{}},
		},
		{
			title:            "team-id and region slug",
			args:             []string{"--team-id", "1", "--region", "test-region"},
			wantClientParams: []any{1, &cherrygo.GetOptions{QueryParams: map[string]string{"region": "test-region"}}},
		},
		{
			title:            "team-id and region id",
			args:             []string{"--team-id", "1", "--region", "1"},
			wantClientParams: []any{1, &cherrygo.GetOptions{QueryParams: map[string]string{"region": "1"}}},
		},
		{
			title:            "shorthands",
			args:             []string{"-t", "1", "-r", "test-region"},
			wantClientParams: []any{1, &cherrygo.GetOptions{QueryParams: map[string]string{"region": "test-region"}}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			fakeSvc := fakes.PlanService{}
			fakeOut := fakes.Outputer{}
			dep := fakeDeps{
				svc: &fakeSvc,
				out: &fakeOut,
			}

			c := Command{
				Deps: dep,
			}
			cmd := c.list()
			cmd.SetArgs(tc.args)
			cmd.SilenceUsage = true

			err := cmd.Execute()
			if err != nil {
				t.Fatal(err.Error())
			}

			if len(fakeSvc.Calls) != 1 {
				t.Fatalf("want 1 api call, got %d", len(fakeSvc.Calls))
			}
			fakeSvc.Calls[0].AssertMethod(t, "List")
			fakeSvc.Calls[0].AssertParams(t, tc.wantClientParams...)

			if len(fakeOut.Calls) != 1 {
				t.Fatalf("want 1 output call, got %d", len(fakeOut.Calls))
			}

			wantPlan := fakeSvc.Plan()
			wantTh := []string{"Plan Slug", "Region Slug", "Stock Hourly", "Hourly Price", "Stock Spot", "Spot Price"}
			wantTd := [][]string{{"test-plan", "test-region", "1", fmt.Sprintf("%f", 1.0), "2", fmt.Sprintf("%f", 0.5)}}
			fakeOut.Calls[0].Assert(t, []cherrygo.Plan{wantPlan}, wantTh, wantTd)
		})
	}
}
