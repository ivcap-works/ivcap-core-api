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

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the artifact service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Upload Doer is the HTTP client used to make requests to the upload endpoint.
	UploadDoer goahttp.Doer

	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// AddCollection Doer is the HTTP client used to make requests to the
	// addCollection endpoint.
	AddCollectionDoer goahttp.Doer

	// RemoveCollection Doer is the HTTP client used to make requests to the
	// removeCollection endpoint.
	RemoveCollectionDoer goahttp.Doer

	// AddMetadata Doer is the HTTP client used to make requests to the addMetadata
	// endpoint.
	AddMetadataDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the artifact service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListDoer:             doer,
		UploadDoer:           doer,
		ReadDoer:             doer,
		AddCollectionDoer:    doer,
		RemoveCollectionDoer: doer,
		AddMetadataDoer:      doer,
		CORSDoer:             doer,
		RestoreResponseBody:  restoreBody,
		scheme:               scheme,
		host:                 host,
		decoder:              dec,
		encoder:              enc,
	}
}

// List returns an endpoint that makes HTTP requests to the artifact service
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
			return nil, goahttp.ErrRequestError("artifact", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Upload returns an endpoint that makes HTTP requests to the artifact service
// upload server.
func (c *Client) Upload() goa.Endpoint {
	var (
		encodeRequest  = EncodeUploadRequest(c.encoder)
		decodeResponse = DecodeUploadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUploadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UploadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("artifact", "upload", err)
		}
		return decodeResponse(resp)
	}
}

// Read returns an endpoint that makes HTTP requests to the artifact service
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
			return nil, goahttp.ErrRequestError("artifact", "read", err)
		}
		return decodeResponse(resp)
	}
}

// AddCollection returns an endpoint that makes HTTP requests to the artifact
// service addCollection server.
func (c *Client) AddCollection() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddCollectionRequest(c.encoder)
		decodeResponse = DecodeAddCollectionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddCollectionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddCollectionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("artifact", "addCollection", err)
		}
		return decodeResponse(resp)
	}
}

// RemoveCollection returns an endpoint that makes HTTP requests to the
// artifact service removeCollection server.
func (c *Client) RemoveCollection() goa.Endpoint {
	var (
		encodeRequest  = EncodeRemoveCollectionRequest(c.encoder)
		decodeResponse = DecodeRemoveCollectionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRemoveCollectionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveCollectionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("artifact", "removeCollection", err)
		}
		return decodeResponse(resp)
	}
}

// AddMetadata returns an endpoint that makes HTTP requests to the artifact
// service addMetadata server.
func (c *Client) AddMetadata() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddMetadataRequest(c.encoder)
		decodeResponse = DecodeAddMetadataResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddMetadataRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddMetadataDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("artifact", "addMetadata", err)
		}
		return decodeResponse(resp)
	}
}
