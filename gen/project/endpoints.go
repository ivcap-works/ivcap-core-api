// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package project

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "project" service endpoints.
type Endpoints struct {
	List               goa.Endpoint
	CreateProject      goa.Endpoint
	Delete             goa.Endpoint
	Read               goa.Endpoint
	ListProjectMembers goa.Endpoint
	UpdateMembership   goa.Endpoint
	RemoveMembership   goa.Endpoint
	DefaultProject     goa.Endpoint
	SetDefaultProject  goa.Endpoint
	ProjectAccount     goa.Endpoint
	SetProjectAccount  goa.Endpoint
}

// NewEndpoints wraps the methods of the "project" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:               NewListEndpoint(s, a.JWTAuth),
		CreateProject:      NewCreateProjectEndpoint(s, a.JWTAuth),
		Delete:             NewDeleteEndpoint(s, a.JWTAuth),
		Read:               NewReadEndpoint(s, a.JWTAuth),
		ListProjectMembers: NewListProjectMembersEndpoint(s, a.JWTAuth),
		UpdateMembership:   NewUpdateMembershipEndpoint(s, a.JWTAuth),
		RemoveMembership:   NewRemoveMembershipEndpoint(s, a.JWTAuth),
		DefaultProject:     NewDefaultProjectEndpoint(s, a.JWTAuth),
		SetDefaultProject:  NewSetDefaultProjectEndpoint(s, a.JWTAuth),
		ProjectAccount:     NewProjectAccountEndpoint(s, a.JWTAuth),
		SetProjectAccount:  NewSetProjectAccountEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "project" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.CreateProject = m(e.CreateProject)
	e.Delete = m(e.Delete)
	e.Read = m(e.Read)
	e.ListProjectMembers = m(e.ListProjectMembers)
	e.UpdateMembership = m(e.UpdateMembership)
	e.RemoveMembership = m(e.RemoveMembership)
	e.DefaultProject = m(e.DefaultProject)
	e.SetDefaultProject = m(e.SetDefaultProject)
	e.ProjectAccount = m(e.ProjectAccount)
	e.SetProjectAccount = m(e.SetProjectAccount)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "project".
func NewListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedProjectListRT(res, view)
		return vres, nil
	}
}

// NewCreateProjectEndpoint returns an endpoint function that calls the method
// "CreateProject" of service "project".
func NewCreateProjectEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateProjectPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:write"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.CreateProject(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedProjectStatusRT(res, view)
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "project".
func NewDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:write"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Delete(ctx, p)
	}
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "project".
func NewReadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReadPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Read(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedProjectStatusRT(res, view)
		return vres, nil
	}
}

// NewListProjectMembersEndpoint returns an endpoint function that calls the
// method "ListProjectMembers" of service "project".
func NewListProjectMembersEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListProjectMembersPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return s.ListProjectMembers(ctx, p)
	}
}

// NewUpdateMembershipEndpoint returns an endpoint function that calls the
// method "UpdateMembership" of service "project".
func NewUpdateMembershipEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateMembershipPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.UpdateMembership(ctx, p)
	}
}

// NewRemoveMembershipEndpoint returns an endpoint function that calls the
// method "RemoveMembership" of service "project".
func NewRemoveMembershipEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RemoveMembershipPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.RemoveMembership(ctx, p)
	}
}

// NewDefaultProjectEndpoint returns an endpoint function that calls the method
// "DefaultProject" of service "project".
func NewDefaultProjectEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DefaultProjectPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.DefaultProject(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedProjectStatusRT(res, view)
		return vres, nil
	}
}

// NewSetDefaultProjectEndpoint returns an endpoint function that calls the
// method "SetDefaultProject" of service "project".
func NewSetDefaultProjectEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SetDefaultProjectPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:write"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.SetDefaultProject(ctx, p)
	}
}

// NewProjectAccountEndpoint returns an endpoint function that calls the method
// "ProjectAccount" of service "project".
func NewProjectAccountEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ProjectAccountPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:read"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return s.ProjectAccount(ctx, p)
	}
}

// NewSetProjectAccountEndpoint returns an endpoint function that calls the
// method "SetProjectAccount" of service "project".
func NewSetProjectAccountEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SetProjectAccountPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:write"},
		}
		ctx, err = authJWTFn(ctx, p.JWT, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.SetProjectAccount(ctx, p)
	}
}
