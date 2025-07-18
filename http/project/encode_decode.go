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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	project "github.com/ivcap-works/ivcap-core-api/gen/project"
	projectviews "github.com/ivcap-works/ivcap-core-api/gen/project/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "project" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListProjectPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the project list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "list", "*project.ListPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		if p.Filter != nil {
			values.Add("filter", *p.Filter)
		}
		if p.OrderBy != nil {
			values.Add("order-by", *p.OrderBy)
		}
		values.Add("order-desc", fmt.Sprintf("%v", p.OrderDesc))
		if p.AtTime != nil {
			values.Add("at-time", *p.AtTime)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the project
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "list", err)
			}
			p := NewListProjectListRTOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &projectviews.ProjectListRT{Projected: p, View: view}
			if err = projectviews.ValidateProjectListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("project", "list", err)
			}
			res := project.NewProjectListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateProjectRequest instantiates a HTTP request object with method and
// path set to call the "project" service "CreateProject" endpoint
func (c *Client) BuildCreateProjectRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateProjectProjectPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "CreateProject", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateProjectRequest returns an encoder for requests sent to the
// project CreateProject server.
func EncodeCreateProjectRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.CreateProjectPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "CreateProject", "*project.CreateProjectPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateProjectRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("project", "CreateProject", err)
		}
		return nil
	}
}

// DecodeCreateProjectResponse returns a decoder for responses returned by the
// project CreateProject endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateProjectResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateProjectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateProjectResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "CreateProject", err)
			}
			p := NewCreateProjectProjectStatusRTOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &projectviews.ProjectStatusRT{Projected: p, View: view}
			if err = projectviews.ValidateProjectStatusRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("project", "CreateProject", err)
			}
			res := project.NewProjectStatusRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateProjectBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "CreateProject", err)
			}
			err = ValidateCreateProjectBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "CreateProject", err)
			}
			return nil, NewCreateProjectBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body CreateProjectInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "CreateProject", err)
			}
			err = ValidateCreateProjectInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "CreateProject", err)
			}
			return nil, NewCreateProjectInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateProjectInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "CreateProject", err)
			}
			err = ValidateCreateProjectInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "CreateProject", err)
			}
			return nil, NewCreateProjectInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateProjectNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "CreateProject", err)
			}
			err = ValidateCreateProjectNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "CreateProject", err)
			}
			return nil, NewCreateProjectNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewCreateProjectNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewCreateProjectNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "CreateProject", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "project" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*project.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "delete", "*project.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteProjectPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the project
// delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "delete", "*project.DeletePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the project
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		case http.StatusForbidden:
			var (
				body DeleteInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "delete", err)
			}
			err = ValidateDeleteInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "delete", err)
			}
			return nil, NewDeleteInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body DeleteNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "delete", err)
			}
			err = ValidateDeleteNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "delete", err)
			}
			return nil, NewDeleteNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewDeleteNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewDeleteNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "delete", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "project" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*project.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "read", "*project.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadProjectPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the project read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "read", "*project.ReadPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeReadResponse returns a decoder for responses returned by the project
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeReadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ReadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "read", err)
			}
			p := NewReadProjectStatusRTOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &projectviews.ProjectStatusRT{Projected: p, View: view}
			if err = projectviews.ValidateProjectStatusRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("project", "read", err)
			}
			res := project.NewProjectStatusRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "read", err)
			}
			err = ValidateReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "read", err)
			}
			return nil, NewReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildListProjectMembersRequest instantiates a HTTP request object with
// method and path set to call the "project" service "ListProjectMembers"
// endpoint
func (c *Client) BuildListProjectMembersRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		urn string
	)
	{
		p, ok := v.(*project.ListProjectMembersPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "ListProjectMembers", "*project.ListProjectMembersPayload", v)
		}
		urn = p.Urn
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListProjectMembersProjectPath(urn)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "ListProjectMembers", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListProjectMembersRequest returns an encoder for requests sent to the
// project ListProjectMembers server.
func EncodeListProjectMembersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.ListProjectMembersPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "ListProjectMembers", "*project.ListProjectMembersPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		if p.Role != nil {
			values.Add("role", *p.Role)
		}
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListProjectMembersResponse returns a decoder for responses returned by
// the project ListProjectMembers endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeListProjectMembersResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeListProjectMembersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListProjectMembersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ListProjectMembers", err)
			}
			err = ValidateListProjectMembersResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ListProjectMembers", err)
			}
			res := NewListProjectMembersMembersListOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListProjectMembersBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ListProjectMembers", err)
			}
			err = ValidateListProjectMembersBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ListProjectMembers", err)
			}
			return nil, NewListProjectMembersBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListProjectMembersInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ListProjectMembers", err)
			}
			err = ValidateListProjectMembersInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ListProjectMembers", err)
			}
			return nil, NewListProjectMembersInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListProjectMembersInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ListProjectMembers", err)
			}
			err = ValidateListProjectMembersInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ListProjectMembers", err)
			}
			return nil, NewListProjectMembersInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListProjectMembersNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ListProjectMembers", err)
			}
			err = ValidateListProjectMembersNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ListProjectMembers", err)
			}
			return nil, NewListProjectMembersNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ListProjectMembersNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ListProjectMembers", err)
			}
			err = ValidateListProjectMembersNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ListProjectMembers", err)
			}
			return nil, NewListProjectMembersNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListProjectMembersNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListProjectMembersNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "ListProjectMembers", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateMembershipRequest instantiates a HTTP request object with method
// and path set to call the "project" service "UpdateMembership" endpoint
func (c *Client) BuildUpdateMembershipRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		projectUrn string
		userUrn    string
	)
	{
		p, ok := v.(*project.UpdateMembershipPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "UpdateMembership", "*project.UpdateMembershipPayload", v)
		}
		projectUrn = p.ProjectUrn
		userUrn = p.UserUrn
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateMembershipProjectPath(projectUrn, userUrn)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "UpdateMembership", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateMembershipRequest returns an encoder for requests sent to the
// project UpdateMembership server.
func EncodeUpdateMembershipRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.UpdateMembershipPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "UpdateMembership", "*project.UpdateMembershipPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateMembershipRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("project", "UpdateMembership", err)
		}
		return nil
	}
}

// DecodeUpdateMembershipResponse returns a decoder for responses returned by
// the project UpdateMembership endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeUpdateMembershipResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateMembershipResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body UpdateMembershipBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "UpdateMembership", err)
			}
			err = ValidateUpdateMembershipBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "UpdateMembership", err)
			}
			return nil, NewUpdateMembershipBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body UpdateMembershipInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "UpdateMembership", err)
			}
			err = ValidateUpdateMembershipInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "UpdateMembership", err)
			}
			return nil, NewUpdateMembershipInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body UpdateMembershipInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "UpdateMembership", err)
			}
			err = ValidateUpdateMembershipInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "UpdateMembership", err)
			}
			return nil, NewUpdateMembershipInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UpdateMembershipNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "UpdateMembership", err)
			}
			err = ValidateUpdateMembershipNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "UpdateMembership", err)
			}
			return nil, NewUpdateMembershipNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body UpdateMembershipNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "UpdateMembership", err)
			}
			err = ValidateUpdateMembershipNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "UpdateMembership", err)
			}
			return nil, NewUpdateMembershipNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewUpdateMembershipNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewUpdateMembershipNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "UpdateMembership", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveMembershipRequest instantiates a HTTP request object with method
// and path set to call the "project" service "RemoveMembership" endpoint
func (c *Client) BuildRemoveMembershipRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		projectUrn string
		userUrn    string
	)
	{
		p, ok := v.(*project.RemoveMembershipPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "RemoveMembership", "*project.RemoveMembershipPayload", v)
		}
		projectUrn = p.ProjectUrn
		userUrn = p.UserUrn
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveMembershipProjectPath(projectUrn, userUrn)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "RemoveMembership", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRemoveMembershipRequest returns an encoder for requests sent to the
// project RemoveMembership server.
func EncodeRemoveMembershipRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.RemoveMembershipPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "RemoveMembership", "*project.RemoveMembershipPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeRemoveMembershipResponse returns a decoder for responses returned by
// the project RemoveMembership endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeRemoveMembershipResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeRemoveMembershipResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body RemoveMembershipBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "RemoveMembership", err)
			}
			err = ValidateRemoveMembershipBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "RemoveMembership", err)
			}
			return nil, NewRemoveMembershipBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body RemoveMembershipInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "RemoveMembership", err)
			}
			err = ValidateRemoveMembershipInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "RemoveMembership", err)
			}
			return nil, NewRemoveMembershipInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body RemoveMembershipInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "RemoveMembership", err)
			}
			err = ValidateRemoveMembershipInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "RemoveMembership", err)
			}
			return nil, NewRemoveMembershipInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body RemoveMembershipNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "RemoveMembership", err)
			}
			err = ValidateRemoveMembershipNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "RemoveMembership", err)
			}
			return nil, NewRemoveMembershipNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body RemoveMembershipNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "RemoveMembership", err)
			}
			err = ValidateRemoveMembershipNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "RemoveMembership", err)
			}
			return nil, NewRemoveMembershipNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewRemoveMembershipNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewRemoveMembershipNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "RemoveMembership", resp.StatusCode, string(body))
		}
	}
}

// BuildDefaultProjectRequest instantiates a HTTP request object with method
// and path set to call the "project" service "DefaultProject" endpoint
func (c *Client) BuildDefaultProjectRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DefaultProjectProjectPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "DefaultProject", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDefaultProjectRequest returns an encoder for requests sent to the
// project DefaultProject server.
func EncodeDefaultProjectRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.DefaultProjectPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "DefaultProject", "*project.DefaultProjectPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDefaultProjectResponse returns a decoder for responses returned by the
// project DefaultProject endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeDefaultProjectResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeDefaultProjectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DefaultProjectResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "DefaultProject", err)
			}
			p := NewDefaultProjectProjectStatusRTOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &projectviews.ProjectStatusRT{Projected: p, View: view}
			if err = projectviews.ValidateProjectStatusRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("project", "DefaultProject", err)
			}
			res := project.NewProjectStatusRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body DefaultProjectBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "DefaultProject", err)
			}
			err = ValidateDefaultProjectBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "DefaultProject", err)
			}
			return nil, NewDefaultProjectBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body DefaultProjectInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "DefaultProject", err)
			}
			err = ValidateDefaultProjectInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "DefaultProject", err)
			}
			return nil, NewDefaultProjectInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body DefaultProjectInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "DefaultProject", err)
			}
			err = ValidateDefaultProjectInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "DefaultProject", err)
			}
			return nil, NewDefaultProjectInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body DefaultProjectNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "DefaultProject", err)
			}
			err = ValidateDefaultProjectNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "DefaultProject", err)
			}
			return nil, NewDefaultProjectNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body DefaultProjectNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "DefaultProject", err)
			}
			err = ValidateDefaultProjectNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "DefaultProject", err)
			}
			return nil, NewDefaultProjectNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewDefaultProjectNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewDefaultProjectNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "DefaultProject", resp.StatusCode, string(body))
		}
	}
}

// BuildSetDefaultProjectRequest instantiates a HTTP request object with method
// and path set to call the "project" service "SetDefaultProject" endpoint
func (c *Client) BuildSetDefaultProjectRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetDefaultProjectProjectPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "SetDefaultProject", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetDefaultProjectRequest returns an encoder for requests sent to the
// project SetDefaultProject server.
func EncodeSetDefaultProjectRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.SetDefaultProjectPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "SetDefaultProject", "*project.SetDefaultProjectPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSetDefaultProjectRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("project", "SetDefaultProject", err)
		}
		return nil
	}
}

// DecodeSetDefaultProjectResponse returns a decoder for responses returned by
// the project SetDefaultProject endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeSetDefaultProjectResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeSetDefaultProjectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body SetDefaultProjectBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetDefaultProject", err)
			}
			err = ValidateSetDefaultProjectBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetDefaultProject", err)
			}
			return nil, NewSetDefaultProjectBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body SetDefaultProjectInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetDefaultProject", err)
			}
			err = ValidateSetDefaultProjectInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetDefaultProject", err)
			}
			return nil, NewSetDefaultProjectInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body SetDefaultProjectInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetDefaultProject", err)
			}
			err = ValidateSetDefaultProjectInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetDefaultProject", err)
			}
			return nil, NewSetDefaultProjectInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body SetDefaultProjectNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetDefaultProject", err)
			}
			err = ValidateSetDefaultProjectNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetDefaultProject", err)
			}
			return nil, NewSetDefaultProjectNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body SetDefaultProjectNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetDefaultProject", err)
			}
			err = ValidateSetDefaultProjectNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetDefaultProject", err)
			}
			return nil, NewSetDefaultProjectNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewSetDefaultProjectNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewSetDefaultProjectNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "SetDefaultProject", resp.StatusCode, string(body))
		}
	}
}

// BuildProjectAccountRequest instantiates a HTTP request object with method
// and path set to call the "project" service "ProjectAccount" endpoint
func (c *Client) BuildProjectAccountRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		projectUrn string
	)
	{
		p, ok := v.(*project.ProjectAccountPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "ProjectAccount", "*project.ProjectAccountPayload", v)
		}
		projectUrn = p.ProjectUrn
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ProjectAccountProjectPath(projectUrn)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "ProjectAccount", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeProjectAccountRequest returns an encoder for requests sent to the
// project ProjectAccount server.
func EncodeProjectAccountRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.ProjectAccountPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "ProjectAccount", "*project.ProjectAccountPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeProjectAccountResponse returns a decoder for responses returned by the
// project ProjectAccount endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeProjectAccountResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeProjectAccountResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ProjectAccountResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ProjectAccount", err)
			}
			err = ValidateProjectAccountResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ProjectAccount", err)
			}
			res := NewProjectAccountAccountResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ProjectAccountBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ProjectAccount", err)
			}
			err = ValidateProjectAccountBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ProjectAccount", err)
			}
			return nil, NewProjectAccountBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ProjectAccountInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ProjectAccount", err)
			}
			err = ValidateProjectAccountInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ProjectAccount", err)
			}
			return nil, NewProjectAccountInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ProjectAccountInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ProjectAccount", err)
			}
			err = ValidateProjectAccountInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ProjectAccount", err)
			}
			return nil, NewProjectAccountInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ProjectAccountNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ProjectAccount", err)
			}
			err = ValidateProjectAccountNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ProjectAccount", err)
			}
			return nil, NewProjectAccountNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ProjectAccountNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "ProjectAccount", err)
			}
			err = ValidateProjectAccountNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "ProjectAccount", err)
			}
			return nil, NewProjectAccountNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewProjectAccountNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewProjectAccountNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "ProjectAccount", resp.StatusCode, string(body))
		}
	}
}

// BuildSetProjectAccountRequest instantiates a HTTP request object with method
// and path set to call the "project" service "SetProjectAccount" endpoint
func (c *Client) BuildSetProjectAccountRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		projectUrn string
	)
	{
		p, ok := v.(*project.SetProjectAccountPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("project", "SetProjectAccount", "*project.SetProjectAccountPayload", v)
		}
		projectUrn = p.ProjectUrn
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetProjectAccountProjectPath(projectUrn)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("project", "SetProjectAccount", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetProjectAccountRequest returns an encoder for requests sent to the
// project SetProjectAccount server.
func EncodeSetProjectAccountRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*project.SetProjectAccountPayload)
		if !ok {
			return goahttp.ErrInvalidType("project", "SetProjectAccount", "*project.SetProjectAccountPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSetProjectAccountRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("project", "SetProjectAccount", err)
		}
		return nil
	}
}

// DecodeSetProjectAccountResponse returns a decoder for responses returned by
// the project SetProjectAccount endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeSetProjectAccountResponse may return the following errors:
//   - "bad-request" (type *project.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *project.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *project.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *project.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *project.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *project.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *project.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeSetProjectAccountResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body SetProjectAccountBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetProjectAccount", err)
			}
			err = ValidateSetProjectAccountBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetProjectAccount", err)
			}
			return nil, NewSetProjectAccountBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body SetProjectAccountInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetProjectAccount", err)
			}
			err = ValidateSetProjectAccountInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetProjectAccount", err)
			}
			return nil, NewSetProjectAccountInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body SetProjectAccountInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetProjectAccount", err)
			}
			err = ValidateSetProjectAccountInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetProjectAccount", err)
			}
			return nil, NewSetProjectAccountInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body SetProjectAccountNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetProjectAccount", err)
			}
			err = ValidateSetProjectAccountNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetProjectAccount", err)
			}
			return nil, NewSetProjectAccountNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body SetProjectAccountNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("project", "SetProjectAccount", err)
			}
			err = ValidateSetProjectAccountNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("project", "SetProjectAccount", err)
			}
			return nil, NewSetProjectAccountNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewSetProjectAccountNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewSetProjectAccountNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("project", "SetProjectAccount", resp.StatusCode, string(body))
		}
	}
}

// unmarshalProjectListItemResponseBodyToProjectviewsProjectListItemView builds
// a value of type *projectviews.ProjectListItemView from a value of type
// *ProjectListItemResponseBody.
func unmarshalProjectListItemResponseBodyToProjectviewsProjectListItemView(v *ProjectListItemResponseBody) *projectviews.ProjectListItemView {
	res := &projectviews.ProjectListItemView{
		Name:       v.Name,
		Role:       v.Role,
		Urn:        v.Urn,
		CreatedAt:  v.CreatedAt,
		ModifiedAt: v.ModifiedAt,
		AtTime:     v.AtTime,
	}

	return res
}

// marshalProjectProjectPropertiesToProjectPropertiesRequestBodyRequestBody
// builds a value of type *ProjectPropertiesRequestBodyRequestBody from a value
// of type *project.ProjectProperties.
func marshalProjectProjectPropertiesToProjectPropertiesRequestBodyRequestBody(v *project.ProjectProperties) *ProjectPropertiesRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &ProjectPropertiesRequestBodyRequestBody{
		Details: v.Details,
	}

	return res
}

// marshalProjectPropertiesRequestBodyRequestBodyToProjectProjectProperties
// builds a value of type *project.ProjectProperties from a value of type
// *ProjectPropertiesRequestBodyRequestBody.
func marshalProjectPropertiesRequestBodyRequestBodyToProjectProjectProperties(v *ProjectPropertiesRequestBodyRequestBody) *project.ProjectProperties {
	if v == nil {
		return nil
	}
	res := &project.ProjectProperties{
		Details: v.Details,
	}

	return res
}

// unmarshalProjectPropertiesResponseBodyToProjectviewsProjectPropertiesView
// builds a value of type *projectviews.ProjectPropertiesView from a value of
// type *ProjectPropertiesResponseBody.
func unmarshalProjectPropertiesResponseBodyToProjectviewsProjectPropertiesView(v *ProjectPropertiesResponseBody) *projectviews.ProjectPropertiesView {
	if v == nil {
		return nil
	}
	res := &projectviews.ProjectPropertiesView{
		Details: v.Details,
	}

	return res
}

// unmarshalUserListItemResponseBodyToProjectUserListItem builds a value of
// type *project.UserListItem from a value of type *UserListItemResponseBody.
func unmarshalUserListItemResponseBodyToProjectUserListItem(v *UserListItemResponseBody) *project.UserListItem {
	res := &project.UserListItem{
		Urn:   v.Urn,
		Email: v.Email,
		Role:  v.Role,
	}

	return res
}
