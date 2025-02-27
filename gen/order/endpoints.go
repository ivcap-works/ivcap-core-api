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

package order

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "order" service endpoints.
type Endpoints struct {
	List     goa.Endpoint
	Read     goa.Endpoint
	Products goa.Endpoint
	Metadata goa.Endpoint
	Create   goa.Endpoint
	Logs     goa.Endpoint
	Top      goa.Endpoint
}

// LogsResponseData holds both the result and the HTTP response body reader of
// the "logs" method.
type LogsResponseData struct {
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "order" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:     NewListEndpoint(s, a.JWTAuth),
		Read:     NewReadEndpoint(s, a.JWTAuth),
		Products: NewProductsEndpoint(s, a.JWTAuth),
		Metadata: NewMetadataEndpoint(s, a.JWTAuth),
		Create:   NewCreateEndpoint(s, a.JWTAuth),
		Logs:     NewLogsEndpoint(s, a.JWTAuth),
		Top:      NewTopEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "order" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Read = m(e.Read)
	e.Products = m(e.Products)
	e.Metadata = m(e.Metadata)
	e.Create = m(e.Create)
	e.Logs = m(e.Logs)
	e.Top = m(e.Top)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "order".
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
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedOrderListRT(res, "default")
		return vres, nil
	}
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "order".
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

// NewProductsEndpoint returns an endpoint function that calls the method
// "products" of service "order".
func NewProductsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ProductsPayload)
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
		return s.Products(ctx, p)
	}
}

// NewMetadataEndpoint returns an endpoint function that calls the method
// "metadata" of service "order".
func NewMetadataEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MetadataPayload)
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
		return s.Metadata(ctx, p)
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "order".
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

// NewLogsEndpoint returns an endpoint function that calls the method "logs" of
// service "order".
func NewLogsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*LogsPayload)
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
		body, err := s.Logs(ctx, p)
		if err != nil {
			return nil, err
		}
		return &LogsResponseData{Body: body}, nil
	}
}

// NewTopEndpoint returns an endpoint function that calls the method "top" of
// service "order".
func NewTopEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*TopPayload)
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
		res, err := s.Top(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedOrderTopResultItemCollection(res, "default")
		return vres, nil
	}
}
