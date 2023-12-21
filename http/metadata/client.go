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

// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the metadata service endpoint HTTP clients.
type Client struct {
	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// UpdateRecord Doer is the HTTP client used to make requests to the
	// update_record endpoint.
	UpdateRecordDoer goahttp.Doer

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
		ReadDoer:            doer,
		ListDoer:            doer,
		AddDoer:             doer,
		UpdateRecordDoer:    doer,
		RevokeDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Read returns an endpoint that makes HTTP requests to the metadata service
// read server.
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
			return nil, goahttp.ErrRequestError("metadata", "read", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the metadata service
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
			return nil, goahttp.ErrRequestError("metadata", "list", err)
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
	return func(ctx context.Context, v any) (any, error) {
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

// UpdateRecord returns an endpoint that makes HTTP requests to the metadata
// service update_record server.
func (c *Client) UpdateRecord() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRecordRequest(c.encoder)
		decodeResponse = DecodeUpdateRecordResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateRecordRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateRecordDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("metadata", "update_record", err)
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
	return func(ctx context.Context, v any) (any, error) {
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
