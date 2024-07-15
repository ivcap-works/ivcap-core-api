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

package search

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "search" service endpoints.
type Endpoints struct {
	Search goa.Endpoint
}

// NewEndpoints wraps the methods of the "search" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Search: NewSearchEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "search" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Search = m(e.Search)
}

// NewSearchEndpoint returns an endpoint function that calls the method
// "search" of service "search".
func NewSearchEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SearchPayload)
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
		return s.Search(ctx, p)
	}
}
