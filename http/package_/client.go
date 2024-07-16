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

	package_ "github.com/ivcap-works/ivcap-core-api/gen/package_"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the package service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Pull Doer is the HTTP client used to make requests to the pull endpoint.
	PullDoer goahttp.Doer

	// Push Doer is the HTTP client used to make requests to the push endpoint.
	PushDoer goahttp.Doer

	// Patch Doer is the HTTP client used to make requests to the patch endpoint.
	PatchDoer goahttp.Doer

	// Put Doer is the HTTP client used to make requests to the put endpoint.
	PutDoer goahttp.Doer

	// Remove Doer is the HTTP client used to make requests to the remove endpoint.
	RemoveDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the package service servers.
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
		PullDoer:            doer,
		PushDoer:            doer,
		PatchDoer:           doer,
		PutDoer:             doer,
		RemoveDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the package service
// list server.
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
			return nil, goahttp.ErrRequestError("package", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Pull returns an endpoint that makes HTTP requests to the package service
// pull server.
func (c *Client) Pull() goa.Endpoint {
	var (
		encodeRequest  = EncodePullRequest(c.encoder)
		decodeResponse = DecodePullResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPullRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PullDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "pull", err)
		}
		res, err := decodeResponse(resp)
		if err != nil {
			resp.Body.Close()
			return nil, err
		}
		return &package_.PullResponseData{Result: res.(*package_.PullResultT), Body: resp.Body}, nil
	}
}

// Push returns an endpoint that makes HTTP requests to the package service
// push server.
func (c *Client) Push() goa.Endpoint {
	var (
		encodeRequest  = EncodePushRequest(c.encoder)
		decodeResponse = DecodePushResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPushRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PushDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "push", err)
		}
		return decodeResponse(resp)
	}
}

// Patch returns an endpoint that makes HTTP requests to the package service
// patch server.
func (c *Client) Patch() goa.Endpoint {
	var (
		encodeRequest  = EncodePatchRequest(c.encoder)
		decodeResponse = DecodePatchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPatchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PatchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "patch", err)
		}
		return decodeResponse(resp)
	}
}

// Put returns an endpoint that makes HTTP requests to the package service put
// server.
func (c *Client) Put() goa.Endpoint {
	var (
		encodeRequest  = EncodePutRequest(c.encoder)
		decodeResponse = DecodePutResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPutRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PutDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "put", err)
		}
		return decodeResponse(resp)
	}
}

// Remove returns an endpoint that makes HTTP requests to the package service
// remove server.
func (c *Client) Remove() goa.Endpoint {
	var (
		encodeRequest  = EncodeRemoveRequest(c.encoder)
		decodeResponse = DecodeRemoveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRemoveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("package", "remove", err)
		}
		return decodeResponse(resp)
	}
}
