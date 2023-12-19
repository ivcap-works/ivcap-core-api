// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package service

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "service" service endpoints.
type Endpoints struct {
	List          goa.Endpoint
	CreateService goa.Endpoint
	Read          goa.Endpoint
	Update        goa.Endpoint
	Delete        goa.Endpoint
}

// NewEndpoints wraps the methods of the "service" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:          NewListEndpoint(s, a.JWTAuth),
		CreateService: NewCreateServiceEndpoint(s, a.JWTAuth),
		Read:          NewReadEndpoint(s, a.JWTAuth),
		Update:        NewUpdateEndpoint(s, a.JWTAuth),
		Delete:        NewDeleteEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "service" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.CreateService = m(e.CreateService)
	e.Read = m(e.Read)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "service".
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
		vres := NewViewedServiceListRT(res, "default")
		return vres, nil
	}
}

// NewCreateServiceEndpoint returns an endpoint function that calls the method
// "create_service" of service "service".
func NewCreateServiceEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateServicePayload)
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
		res, view, err := s.CreateService(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedServiceStatusRT(res, view)
		return vres, nil
	}
}

// NewReadEndpoint returns an endpoint function that calls the method "read" of
// service "service".
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
		vres := NewViewedServiceStatusRT(res, view)
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "service".
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
		res, view, err := s.Update(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedServiceStatusRT(res, view)
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "service".
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
