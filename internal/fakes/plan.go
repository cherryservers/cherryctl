package fakes

import (
	"context"
	"errors"

	"github.com/cherryservers/cherrygo/v4"
)

type PlanService struct {
	Calls []CallRecord
}

func (s *PlanService) List(ctx context.Context, teamID int, opts *cherrygo.GetOptions) ([]cherrygo.Plan, *cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "List", params: []any{teamID, opts}})
	return []cherrygo.Plan{s.Plan()}, nil, nil
}

func (s *PlanService) GetBySlug(ctx context.Context, slug string, opts *cherrygo.GetOptions) (cherrygo.Plan, *cherrygo.Response, error) {
	s.Calls = append(s.Calls, CallRecord{method: "GetBySlug", params: []any{slug, opts}})
	return s.Plan(), nil, nil
}

func (s *PlanService) GetByID(_ context.Context, _ int, _ *cherrygo.GetOptions) (cherrygo.Plan, *cherrygo.Response, error) {
	return cherrygo.Plan{}, nil, errors.New("not implemented")
}

func (s *PlanService) ListPrebuiltPlans(_ context.Context, _, _ string, _ *cherrygo.GetOptions) ([]cherrygo.PrebuiltPlan, *cherrygo.Response, error) {
	return []cherrygo.PrebuiltPlan{}, nil, errors.New("not implemented")
}

func (s *PlanService) ListPrebuiltTeamPlans(_ context.Context, _, _ string, _ int, _ *cherrygo.GetOptions) ([]cherrygo.PrebuiltPlan, *cherrygo.Response, error) {
	return []cherrygo.PrebuiltPlan{}, nil, errors.New("not implemented")
}

// Plan is the plan that will be returned by the fake methods.
func (s *PlanService) Plan() cherrygo.Plan {
	pricing := []cherrygo.Pricing{
		{
			Unit:  "Hourly",
			Price: 1.0,
		},
		{
			Unit:  "Spot hourly",
			Price: 0.5,
		},
	}
	regions := []cherrygo.AvailableRegions{
		{
			Region:   &cherrygo.Region{ID: 1, Slug: "test-region"},
			StockQty: 1,
			SpotQty:  2,
		},
	}
	return cherrygo.Plan{ID: 1, Slug: "test-plan", Pricing: pricing, AvailableRegions: regions}
}
