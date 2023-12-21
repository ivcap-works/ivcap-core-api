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

package artifact

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "artifact" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Read   goa.Endpoint
	Upload goa.Endpoint
}

// UploadRequestData holds both the payload and the HTTP request body reader of
// the "upload" method.
type UploadRequestData struct {
	// Payload is the method payload.
	Payload *UploadPayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "artifact" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:   NewListEndpoint(s, a.JWTAuth),
		Read:   NewReadEndpoint(s, a.JWTAuth),
		Upload: NewUploadEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "artifact" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Read = m(e.Read)
	e.Upload = m(e.Upload)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "artifact".
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

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "artifact".
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

// NewUploadEndpoint returns an endpoint function that calls the method
// "upload" of service "artifact".
func NewUploadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		ep := req.(*UploadRequestData)
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
		return s.Upload(ctx, ep.Payload, ep.Body)
	}
}
