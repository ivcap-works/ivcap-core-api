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

	service "github.com/ivcap-works/ivcap-core-api/gen/service"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the service service endpoint HTTP clients.
type Client struct {
	// ServiceList Doer is the HTTP client used to make requests to the
	// service-list endpoint.
	ServiceListDoer goahttp.Doer

	// ServiceCreate Doer is the HTTP client used to make requests to the
	// service-create endpoint.
	ServiceCreateDoer goahttp.Doer

	// ServiceRead Doer is the HTTP client used to make requests to the
	// service-read endpoint.
	ServiceReadDoer goahttp.Doer

	// ServiceUpdate Doer is the HTTP client used to make requests to the
	// service-update endpoint.
	ServiceUpdateDoer goahttp.Doer

	// ServiceDelete Doer is the HTTP client used to make requests to the
	// service-delete endpoint.
	ServiceDeleteDoer goahttp.Doer

	// JobList Doer is the HTTP client used to make requests to the job-list
	// endpoint.
	JobListDoer goahttp.Doer

	// JobCreate Doer is the HTTP client used to make requests to the job-create
	// endpoint.
	JobCreateDoer goahttp.Doer

	// JobRead Doer is the HTTP client used to make requests to the job-read
	// endpoint.
	JobReadDoer goahttp.Doer

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
		ServiceListDoer:     doer,
		ServiceCreateDoer:   doer,
		ServiceReadDoer:     doer,
		ServiceUpdateDoer:   doer,
		ServiceDeleteDoer:   doer,
		JobListDoer:         doer,
		JobCreateDoer:       doer,
		JobReadDoer:         doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// ServiceList returns an endpoint that makes HTTP requests to the service
// service service-list server.
func (c *Client) ServiceList() goa.Endpoint {
	var (
		encodeRequest  = EncodeServiceListRequest(c.encoder)
		decodeResponse = DecodeServiceListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildServiceListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ServiceListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "service-list", err)
		}
		return decodeResponse(resp)
	}
}

// ServiceCreate returns an endpoint that makes HTTP requests to the service
// service service-create server.
func (c *Client) ServiceCreate() goa.Endpoint {
	var (
		encodeRequest  = EncodeServiceCreateRequest(c.encoder)
		decodeResponse = DecodeServiceCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildServiceCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ServiceCreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "service-create", err)
		}
		return decodeResponse(resp)
	}
}

// ServiceRead returns an endpoint that makes HTTP requests to the service
// service service-read server.
func (c *Client) ServiceRead() goa.Endpoint {
	var (
		encodeRequest  = EncodeServiceReadRequest(c.encoder)
		decodeResponse = DecodeServiceReadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildServiceReadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ServiceReadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "service-read", err)
		}
		return decodeResponse(resp)
	}
}

// ServiceUpdate returns an endpoint that makes HTTP requests to the service
// service service-update server.
func (c *Client) ServiceUpdate() goa.Endpoint {
	var (
		encodeRequest  = EncodeServiceUpdateRequest(c.encoder)
		decodeResponse = DecodeServiceUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildServiceUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ServiceUpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "service-update", err)
		}
		return decodeResponse(resp)
	}
}

// ServiceDelete returns an endpoint that makes HTTP requests to the service
// service service-delete server.
func (c *Client) ServiceDelete() goa.Endpoint {
	var (
		encodeRequest  = EncodeServiceDeleteRequest(c.encoder)
		decodeResponse = DecodeServiceDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildServiceDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ServiceDeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "service-delete", err)
		}
		return decodeResponse(resp)
	}
}

// JobList returns an endpoint that makes HTTP requests to the service service
// job-list server.
func (c *Client) JobList() goa.Endpoint {
	var (
		encodeRequest  = EncodeJobListRequest(c.encoder)
		decodeResponse = DecodeJobListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildJobListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.JobListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "job-list", err)
		}
		return decodeResponse(resp)
	}
}

// JobCreate returns an endpoint that makes HTTP requests to the service
// service job-create server.
func (c *Client) JobCreate() goa.Endpoint {
	var (
		encodeRequest  = EncodeJobCreateRequest(c.encoder)
		decodeResponse = DecodeJobCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildJobCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.JobCreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "job-create", err)
		}
		res, err := decodeResponse(resp)
		if err != nil {
			resp.Body.Close()
			return nil, err
		}
		return &service.JobCreateResponseData{Result: res.(*service.JobCreateResult), Body: resp.Body}, nil
	}
}

// JobRead returns an endpoint that makes HTTP requests to the service service
// job-read server.
func (c *Client) JobRead() goa.Endpoint {
	var (
		encodeRequest  = EncodeJobReadRequest(c.encoder)
		decodeResponse = DecodeJobReadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildJobReadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.JobReadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("service", "job-read", err)
		}
		return decodeResponse(resp)
	}
}
