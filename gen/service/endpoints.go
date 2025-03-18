// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package service

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "service" service endpoints.
type Endpoints struct {
	ServiceList   goa.Endpoint
	ServiceCreate goa.Endpoint
	ServiceRead   goa.Endpoint
	ServiceUpdate goa.Endpoint
	ServiceDelete goa.Endpoint
	JobList       goa.Endpoint
	JobCreate     goa.Endpoint
	JobRead       goa.Endpoint
}

// JobCreateRequestData holds both the payload and the HTTP request body reader
// of the "job-create" method.
type JobCreateRequestData struct {
	// Payload is the method payload.
	Payload *JobCreatePayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// JobCreateResponseData holds both the result and the HTTP response body
// reader of the "job-create" method.
type JobCreateResponseData struct {
	// Result is the method result.
	Result *JobCreateResult
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "service" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ServiceList:   NewServiceListEndpoint(s, a.JWTAuth),
		ServiceCreate: NewServiceCreateEndpoint(s, a.JWTAuth),
		ServiceRead:   NewServiceReadEndpoint(s, a.JWTAuth),
		ServiceUpdate: NewServiceUpdateEndpoint(s, a.JWTAuth),
		ServiceDelete: NewServiceDeleteEndpoint(s, a.JWTAuth),
		JobList:       NewJobListEndpoint(s, a.JWTAuth),
		JobCreate:     NewJobCreateEndpoint(s, a.JWTAuth),
		JobRead:       NewJobReadEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "service" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ServiceList = m(e.ServiceList)
	e.ServiceCreate = m(e.ServiceCreate)
	e.ServiceRead = m(e.ServiceRead)
	e.ServiceUpdate = m(e.ServiceUpdate)
	e.ServiceDelete = m(e.ServiceDelete)
	e.JobList = m(e.JobList)
	e.JobCreate = m(e.JobCreate)
	e.JobRead = m(e.JobRead)
}

// NewServiceListEndpoint returns an endpoint function that calls the method
// "service-list" of service "service".
func NewServiceListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ServiceListPayload)
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
		res, err := s.ServiceList(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedServiceListRT(res, "default")
		return vres, nil
	}
}

// NewServiceCreateEndpoint returns an endpoint function that calls the method
// "service-create" of service "service".
func NewServiceCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ServiceCreatePayload)
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
		return s.ServiceCreate(ctx, p)
	}
}

// NewServiceReadEndpoint returns an endpoint function that calls the method
// "service-read" of service "service".
func NewServiceReadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ServiceReadPayload)
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
		return s.ServiceRead(ctx, p)
	}
}

// NewServiceUpdateEndpoint returns an endpoint function that calls the method
// "service-update" of service "service".
func NewServiceUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ServiceUpdatePayload)
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
		return s.ServiceUpdate(ctx, p)
	}
}

// NewServiceDeleteEndpoint returns an endpoint function that calls the method
// "service-delete" of service "service".
func NewServiceDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ServiceDeletePayload)
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
		return nil, s.ServiceDelete(ctx, p)
	}
}

// NewJobListEndpoint returns an endpoint function that calls the method
// "job-list" of service "service".
func NewJobListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*JobListPayload)
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
		res, err := s.JobList(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJobListRT(res, "default")
		return vres, nil
	}
}

// NewJobCreateEndpoint returns an endpoint function that calls the method
// "job-create" of service "service".
func NewJobCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		ep := req.(*JobCreateRequestData)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"consumer:read", "consumer:write"},
			RequiredScopes: []string{"consumer:write"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.JWT, &sc)
		if err != nil {
			return nil, err
		}
		res, body, err := s.JobCreate(ctx, ep.Payload, ep.Body)
		if err != nil {
			return nil, err
		}
		return &JobCreateResponseData{Result: res, Body: body}, nil
	}
}

// NewJobReadEndpoint returns an endpoint function that calls the method
// "job-read" of service "service".
func NewJobReadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*JobReadPayload)
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
		return s.JobRead(ctx, p)
	}
}
