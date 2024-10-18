// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

package secret

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "secret" service endpoints.
type Endpoints struct {
	List goa.Endpoint
	Get  goa.Endpoint
	Set  goa.Endpoint
}

// NewEndpoints wraps the methods of the "secret" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List: NewListEndpoint(s, a.JWTAuth),
		Get:  NewGetEndpoint(s, a.JWTAuth),
		Set:  NewSetEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "secret" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Get = m(e.Get)
	e.Set = m(e.Set)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "secret".
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

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "secret".
func NewGetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetPayload)
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
		return s.Get(ctx, p)
	}
}

// NewSetEndpoint returns an endpoint function that calls the method "set" of
// service "secret".
func NewSetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SetPayload)
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
		return nil, s.Set(ctx, p)
	}
}
