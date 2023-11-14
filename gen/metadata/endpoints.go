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

package metadata

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "metadata" service endpoints.
type Endpoints struct {
	Read         goa.Endpoint
	List         goa.Endpoint
	Add          goa.Endpoint
	UpdateOne    goa.Endpoint
	UpdateRecord goa.Endpoint
	Revoke       goa.Endpoint
}

// NewEndpoints wraps the methods of the "metadata" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Read:         NewReadEndpoint(s, a.JWTAuth),
		List:         NewListEndpoint(s, a.JWTAuth),
		Add:          NewAddEndpoint(s, a.JWTAuth),
		UpdateOne:    NewUpdateOneEndpoint(s, a.JWTAuth),
		UpdateRecord: NewUpdateRecordEndpoint(s, a.JWTAuth),
		Revoke:       NewRevokeEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "metadata" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Read = m(e.Read)
	e.List = m(e.List)
	e.Add = m(e.Add)
	e.UpdateOne = m(e.UpdateOne)
	e.UpdateRecord = m(e.UpdateRecord)
	e.Revoke = m(e.Revoke)
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "metadata".
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
		res, err := s.Read(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMetadataRecordRT(res, "default")
		return vres, nil
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "metadata".
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
		vres := NewViewedListMetaRT(res, "default")
		return vres, nil
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "metadata".
func NewAddEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
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
		res, err := s.Add(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAddMetaRT(res, "default")
		return vres, nil
	}
}

// NewUpdateOneEndpoint returns an endpoint function that calls the method
// "update_one" of service "metadata".
func NewUpdateOneEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateOnePayload)
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
		res, err := s.UpdateOne(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAddMetaRT(res, "default")
		return vres, nil
	}
}

// NewUpdateRecordEndpoint returns an endpoint function that calls the method
// "update_record" of service "metadata".
func NewUpdateRecordEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateRecordPayload)
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
		res, err := s.UpdateRecord(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAddMetaRT(res, "default")
		return vres, nil
	}
}

// NewRevokeEndpoint returns an endpoint function that calls the method
// "revoke" of service "metadata".
func NewRevokeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RevokePayload)
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
		return nil, s.Revoke(ctx, p)
	}
}
