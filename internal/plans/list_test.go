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
	return nil
}

func (fd fakeDeps) Client() cherrygo.PlansService {
	return fd.svc
}

func (fd fakeDeps) Outputer() outputs.Outputer {
	return fd.out
}

func TestList(t *testing.T) {
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
	cmd.SetArgs([]string{"--team-id", "1"})
	cmd.SilenceUsage = true

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(fakeSvc.Calls) != 1 {
		t.Fatalf("want 1 api call, got %d", len(fakeSvc.Calls))
	}
	fakeSvc.Calls[0].AssertMethod(t, "List")
	fakeSvc.Calls[0].AssertParams(t, 1, (*cherrygo.GetOptions)(nil))

	if len(fakeOut.Calls) != 1{
		t.Fatalf("want 1 output call, got %d", len(fakeOut.Calls))
	}

	wantPlan := fakeSvc.Plan()
	wantTh := []string{"Plan Slug", "Region Slug", "Stock Hourly", "Hourly Price", "Stock Spot", "Spot Price"}
	wantTd := [][]string{{"test-slug", "test-slug", "1", fmt.Sprintf("%f", 1.0), "2", fmt.Sprintf("%f", 0.5)}}
	fakeOut.Calls[0].Assert(t, []cherrygo.Plan{wantPlan}, wantTh, wantTd)
}
