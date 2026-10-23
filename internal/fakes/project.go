package fakes

import (
	"context"

	"github.com/cherryservers/cherrygo/v4"
)

var _ cherrygo.ProjectsService = (*ProjectsService)(nil)

type ProjectsService struct {
	fakeService
}

// Create implements [cherrygo.ProjectsService].
func (p *ProjectsService) Create(_ context.Context, _ int, _ *cherrygo.CreateProject) (cherrygo.Project, *cherrygo.Response, error) {
	return cherrygo.Project{}, nil, notImplementedError
}

// Delete implements [cherrygo.ProjectsService].
func (p *ProjectsService) Delete(ctx context.Context, projectID int) (*cherrygo.Response, error) {
	return nil, notImplementedError
}

// Get implements [cherrygo.ProjectsService].
func (p *ProjectsService) Get(ctx context.Context, projectID int, opts *cherrygo.GetOptions) (cherrygo.Project, *cherrygo.Response, error) {
	return cherrygo.Project{}, nil, notImplementedError
}

// List implements [cherrygo.ProjectsService].
func (p *ProjectsService) List(ctx context.Context, teamID int, opts *cherrygo.GetOptions) ([]cherrygo.Project, *cherrygo.Response, error) {
	return nil, nil, notImplementedError
}

// ListSSHKeys implements [cherrygo.ProjectsService].
func (p *ProjectsService) ListSSHKeys(ctx context.Context, projectID int, opts *cherrygo.GetOptions) ([]cherrygo.SSHKey, *cherrygo.Response, error) {
	p.Calls = append(p.Calls, CallRecord{
		params: []any{ctx, projectID, opts},
		method: "ListSSHKeys",
	})
	return []cherrygo.SSHKey{SSHKey()}, nil, p.Err
}

// Update implements [cherrygo.ProjectsService].
func (p *ProjectsService) Update(ctx context.Context, projectID int, request *cherrygo.UpdateProject) (cherrygo.Project, *cherrygo.Response, error) {
	return cherrygo.Project{}, nil, notImplementedError
}
