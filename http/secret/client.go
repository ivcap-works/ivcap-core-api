// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the secret service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Get Doer is the HTTP client used to make requests to the get endpoint.
	GetDoer goahttp.Doer

	// Set Doer is the HTTP client used to make requests to the set endpoint.
	SetDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the secret service servers.
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
		GetDoer:             doer,
		SetDoer:             doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the secret service list
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
			return nil, goahttp.ErrRequestError("secret", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Get returns an endpoint that makes HTTP requests to the secret service get
// server.
func (c *Client) Get() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetRequest(c.encoder)
		decodeResponse = DecodeGetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("secret", "get", err)
		}
		return decodeResponse(resp)
	}
}

// Set returns an endpoint that makes HTTP requests to the secret service set
// server.
func (c *Client) Set() goa.Endpoint {
	var (
		encodeRequest  = EncodeSetRequest(c.encoder)
		decodeResponse = DecodeSetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("secret", "set", err)
		}
		return decodeResponse(resp)
	}
}
