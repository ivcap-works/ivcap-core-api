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
