// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the metadata service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Revoke Doer is the HTTP client used to make requests to the revoke endpoint.
	RevokeDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the metadata service servers.
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
		ReadDoer:            doer,
		AddDoer:             doer,
		UpdateDoer:          doer,
		RevokeDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the metadata service
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
			return nil, goahttp.ErrRequestError("metadata", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Read returns an endpoint that makes HTTP requests to the metadata service
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
			return nil, goahttp.ErrRequestError("metadata", "read", err)
		}
		return decodeResponse(resp)
	}
}

// Add returns an endpoint that makes HTTP requests to the metadata service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddRequest(c.encoder)
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("metadata", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the metadata service
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
			return nil, goahttp.ErrRequestError("metadata", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Revoke returns an endpoint that makes HTTP requests to the metadata service
// revoke server.
func (c *Client) Revoke() goa.Endpoint {
	var (
		encodeRequest  = EncodeRevokeRequest(c.encoder)
		decodeResponse = DecodeRevokeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRevokeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RevokeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("metadata", "revoke", err)
		}
		return decodeResponse(resp)
	}
}
