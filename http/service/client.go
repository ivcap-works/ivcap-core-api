// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the service service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the service service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListDoer:            doer,
		CreateDoer:          doer,
		ReadDoer:            doer,
		UpdateDoer:          doer,
		DeleteDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the service service
// list server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Create returns an endpoint that makes HTTP requests to the service service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "create", err)
		}
		return decodeResponse(resp)
	}
}

// Read returns an endpoint that makes HTTP requests to the service service
// read server.
func (c *Client) Read() goa.Endpoint {
	var (
		encodeRequest  = EncodeReadRequest(c.encoder)
		decodeResponse = DecodeReadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "read", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the service service
// update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the service service
// delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteRequest(c.encoder)
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "delete", err)
		}
		return decodeResponse(resp)
	}
}
