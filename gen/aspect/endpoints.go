// Copyright 2023 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

package aspect

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "aspect" service endpoints.
type Endpoints struct {
	Read    goa.Endpoint
	List    goa.Endpoint
	Create  goa.Endpoint
	Update  goa.Endpoint
	Retract goa.Endpoint
}

// NewEndpoints wraps the methods of the "aspect" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Read:    NewReadEndpoint(s, a.JWTAuth),
		List:    NewListEndpoint(s, a.JWTAuth),
		Create:  NewCreateEndpoint(s, a.JWTAuth),
		Update:  NewUpdateEndpoint(s, a.JWTAuth),
		Retract: NewRetractEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "aspect" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Read = m(e.Read)
	e.List = m(e.List)
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Retract = m(e.Retract)
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "aspect".
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
		return s.Read(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "aspect".
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

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "aspect".
func NewCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
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
		return s.Create(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "aspect".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
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
		return s.Update(ctx, p)
	}
}

// NewRetractEndpoint returns an endpoint function that calls the method
// "retract" of service "aspect".
func NewRetractEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RetractPayload)
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
		return nil, s.Retract(ctx, p)
	}
}
