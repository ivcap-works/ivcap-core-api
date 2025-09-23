// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the queue service endpoint HTTP clients.
type Client struct {
	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Enqueue Doer is the HTTP client used to make requests to the enqueue
	// endpoint.
	EnqueueDoer goahttp.Doer

	// Dequeue Doer is the HTTP client used to make requests to the dequeue
	// endpoint.
	DequeueDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the queue service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateDoer:          doer,
		ReadDoer:            doer,
		DeleteDoer:          doer,
		ListDoer:            doer,
		EnqueueDoer:         doer,
		DequeueDoer:         doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Create returns an endpoint that makes HTTP requests to the queue service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
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
			return nil, goahttp.ErrRequestError("queue", "create", err)
		}
		return decodeResponse(resp)
	}
}

// Read returns an endpoint that makes HTTP requests to the queue service read
// server.
func (c *Client) Read() goa.Endpoint {
	var (
		encodeRequest  = EncodeReadRequest(c.encoder)
		decodeResponse = DecodeReadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
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
			return nil, goahttp.ErrRequestError("queue", "read", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the queue service
// delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteRequest(c.encoder)
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
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
			return nil, goahttp.ErrRequestError("queue", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the queue service list
// server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
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
			return nil, goahttp.ErrRequestError("queue", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Enqueue returns an endpoint that makes HTTP requests to the queue service
// enqueue server.
func (c *Client) Enqueue() goa.Endpoint {
	var (
		encodeRequest  = EncodeEnqueueRequest(c.encoder)
		decodeResponse = DecodeEnqueueResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildEnqueueRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.EnqueueDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("queue", "enqueue", err)
		}
		return decodeResponse(resp)
	}
}

// Dequeue returns an endpoint that makes HTTP requests to the queue service
// dequeue server.
func (c *Client) Dequeue() goa.Endpoint {
	var (
		encodeRequest  = EncodeDequeueRequest(c.encoder)
		decodeResponse = DecodeDequeueResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDequeueRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DequeueDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("queue", "dequeue", err)
		}
		return decodeResponse(resp)
	}
}
