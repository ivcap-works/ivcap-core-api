// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// Client lists the project service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// CreateProject Doer is the HTTP client used to make requests to the
	// CreateProject endpoint.
	CreateProjectDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// ListProjectMembers Doer is the HTTP client used to make requests to the
	// ListProjectMembers endpoint.
	ListProjectMembersDoer goahttp.Doer

	// UpdateMembership Doer is the HTTP client used to make requests to the
	// UpdateMembership endpoint.
	UpdateMembershipDoer goahttp.Doer

	// RemoveMembership Doer is the HTTP client used to make requests to the
	// RemoveMembership endpoint.
	RemoveMembershipDoer goahttp.Doer

	// DefaultProject Doer is the HTTP client used to make requests to the
	// DefaultProject endpoint.
	DefaultProjectDoer goahttp.Doer

	// SetDefaultProject Doer is the HTTP client used to make requests to the
	// SetDefaultProject endpoint.
	SetDefaultProjectDoer goahttp.Doer

	// ProjectAccount Doer is the HTTP client used to make requests to the
	// ProjectAccount endpoint.
	ProjectAccountDoer goahttp.Doer

	// SetProjectAccount Doer is the HTTP client used to make requests to the
	// SetProjectAccount endpoint.
	SetProjectAccountDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the project service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListDoer:               doer,
		CreateProjectDoer:      doer,
		DeleteDoer:             doer,
		ReadDoer:               doer,
		ListProjectMembersDoer: doer,
		UpdateMembershipDoer:   doer,
		RemoveMembershipDoer:   doer,
		DefaultProjectDoer:     doer,
		SetDefaultProjectDoer:  doer,
		ProjectAccountDoer:     doer,
		SetProjectAccountDoer:  doer,
		CORSDoer:               doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// List returns an endpoint that makes HTTP requests to the project service
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
			return nil, goahttp.ErrRequestError("project", "list", err)
		}
		return decodeResponse(resp)
	}
}

// CreateProject returns an endpoint that makes HTTP requests to the project
// service CreateProject server.
func (c *Client) CreateProject() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateProjectRequest(c.encoder)
		decodeResponse = DecodeCreateProjectResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreateProjectRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateProjectDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "CreateProject", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the project service
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
			return nil, goahttp.ErrRequestError("project", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// Read returns an endpoint that makes HTTP requests to the project service
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
			return nil, goahttp.ErrRequestError("project", "read", err)
		}
		return decodeResponse(resp)
	}
}

// ListProjectMembers returns an endpoint that makes HTTP requests to the
// project service ListProjectMembers server.
func (c *Client) ListProjectMembers() goa.Endpoint {
	var (
		encodeRequest  = EncodeListProjectMembersRequest(c.encoder)
		decodeResponse = DecodeListProjectMembersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListProjectMembersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListProjectMembersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "ListProjectMembers", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateMembership returns an endpoint that makes HTTP requests to the project
// service UpdateMembership server.
func (c *Client) UpdateMembership() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateMembershipRequest(c.encoder)
		decodeResponse = DecodeUpdateMembershipResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateMembershipRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateMembershipDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "UpdateMembership", err)
		}
		return decodeResponse(resp)
	}
}

// RemoveMembership returns an endpoint that makes HTTP requests to the project
// service RemoveMembership server.
func (c *Client) RemoveMembership() goa.Endpoint {
	var (
		encodeRequest  = EncodeRemoveMembershipRequest(c.encoder)
		decodeResponse = DecodeRemoveMembershipResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRemoveMembershipRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveMembershipDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "RemoveMembership", err)
		}
		return decodeResponse(resp)
	}
}

// DefaultProject returns an endpoint that makes HTTP requests to the project
// service DefaultProject server.
func (c *Client) DefaultProject() goa.Endpoint {
	var (
		encodeRequest  = EncodeDefaultProjectRequest(c.encoder)
		decodeResponse = DecodeDefaultProjectResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDefaultProjectRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DefaultProjectDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "DefaultProject", err)
		}
		return decodeResponse(resp)
	}
}

// SetDefaultProject returns an endpoint that makes HTTP requests to the
// project service SetDefaultProject server.
func (c *Client) SetDefaultProject() goa.Endpoint {
	var (
		encodeRequest  = EncodeSetDefaultProjectRequest(c.encoder)
		decodeResponse = DecodeSetDefaultProjectResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSetDefaultProjectRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SetDefaultProjectDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "SetDefaultProject", err)
		}
		return decodeResponse(resp)
	}
}

// ProjectAccount returns an endpoint that makes HTTP requests to the project
// service ProjectAccount server.
func (c *Client) ProjectAccount() goa.Endpoint {
	var (
		encodeRequest  = EncodeProjectAccountRequest(c.encoder)
		decodeResponse = DecodeProjectAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildProjectAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ProjectAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "ProjectAccount", err)
		}
		return decodeResponse(resp)
	}
}

// SetProjectAccount returns an endpoint that makes HTTP requests to the
// project service SetProjectAccount server.
func (c *Client) SetProjectAccount() goa.Endpoint {
	var (
		encodeRequest  = EncodeSetProjectAccountRequest(c.encoder)
		decodeResponse = DecodeSetProjectAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSetProjectAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SetProjectAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("project", "SetProjectAccount", err)
		}
		return decodeResponse(resp)
	}
}
