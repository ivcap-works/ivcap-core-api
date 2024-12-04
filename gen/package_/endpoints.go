// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

package package_

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "package" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Pull   goa.Endpoint
	Push   goa.Endpoint
	Status goa.Endpoint
	Remove goa.Endpoint
}

// PullResponseData holds both the result and the HTTP response body reader of
// the "pull" method.
type PullResponseData struct {
	// Result is the method result.
	Result *PullResultT
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// PushRequestData holds both the payload and the HTTP request body reader of
// the "push" method.
type PushRequestData struct {
	// Payload is the method payload.
	Payload *PushPayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "package" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:   NewListEndpoint(s, a.JWTAuth),
		Pull:   NewPullEndpoint(s, a.JWTAuth),
		Push:   NewPushEndpoint(s, a.JWTAuth),
		Status: NewStatusEndpoint(s, a.JWTAuth),
		Remove: NewRemoveEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "package" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Pull = m(e.Pull)
	e.Push = m(e.Push)
	e.Status = m(e.Status)
	e.Remove = m(e.Remove)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "package".
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
		return s.List(ctx, p)
	}
}

// NewPullEndpoint returns an endpoint function that calls the method "pull" of
// service "package".
func NewPullEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*PullPayload)
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
		res, body, err := s.Pull(ctx, p)
		if err != nil {
			return nil, err
		}
		return &PullResponseData{Result: res, Body: body}, nil
	}
}

// NewPushEndpoint returns an endpoint function that calls the method "push" of
// service "package".
func NewPushEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		ep := req.(*PushRequestData)
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
		return s.Push(ctx, ep.Payload, ep.Body)
	}
}

// NewStatusEndpoint returns an endpoint function that calls the method
// "status" of service "package".
func NewStatusEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*StatusPayload)
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
		return s.Status(ctx, p)
	}
}

// NewRemoveEndpoint returns an endpoint function that calls the method
// "remove" of service "package".
func NewRemoveEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RemovePayload)
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
		return nil, s.Remove(ctx, p)
	}
}
