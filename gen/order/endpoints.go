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

// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package order

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "order" service endpoints.
type Endpoints struct {
	Read   goa.Endpoint
	List   goa.Endpoint
	Create goa.Endpoint
}

// NewEndpoints wraps the methods of the "order" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Read:   NewReadEndpoint(s, a.JWTAuth),
		List:   NewListEndpoint(s, a.JWTAuth),
		Create: NewCreateEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "order" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Read = m(e.Read)
	e.List = m(e.List)
	e.Create = m(e.Create)
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "order".
func NewReadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
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
		vres := NewViewedOrderStatusRT(res, view)
		return vres, nil
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "order".
func NewListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
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
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedOrderListRT(res, "default")
		return vres, nil
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "order".
func NewCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
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
		res, view, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedOrderStatusRT(res, view)
		return vres, nil
	}
}
