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
	"bytes"
	service "github.com/reinventingscience/ivcap-core-api/gen/service"
	serviceviews "github.com/reinventingscience/ivcap-core-api/gen/service/views"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "service" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the service list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*service.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "list", "*service.ListPayload", v)
		}
		values := req.URL.Query()
		values.Add("$filter", p.Filter)
		values.Add("$orderby", p.Orderby)
		values.Add("$top", fmt.Sprintf("%v", p.Top))
		values.Add("$skip", fmt.Sprintf("%v", p.Skip))
		values.Add("$select", p.Select)
		if p.Offset != nil {
			values.Add("offset", fmt.Sprintf("%v", *p.Offset))
		}
		if p.Limit != nil {
			values.Add("limit", fmt.Sprintf("%v", *p.Limit))
		}
		values.Add("pageToken", p.PageToken)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the service
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *service.InvalidCredentialsT): http.StatusBadRequest
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("service", "list", err)
			}
			p := NewListServiceListRTOK(&body)
			view := "default"
			vres := &serviceviews.ServiceListRT{Projected: p, View: view}
			if err = serviceviews.ValidateServiceListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "list", err)
			}
			res := service.NewServiceListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body ListBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "list", err)
				}
				err = ValidateListBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "list", err)
				}
				return nil, NewListBadRequest(&body)
			case "invalid-credential":
				return nil, NewListInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "list", resp.StatusCode, string(body))
			}
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "service" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the service
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*service.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "create", "*service.CreatePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("service", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the service
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *service.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "already-created" (type *service.ResourceAlreadyCreatedT): http.StatusConflict
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
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
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create", err)
			}
			p := NewCreateServiceStatusRTCreated(&body)
			view := resp.Header.Get("goa-view")
			vres := &serviceviews.ServiceStatusRT{Projected: p, View: view}
			if err = serviceviews.ValidateServiceStatusRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "create", err)
			}
			res := service.NewServiceStatusRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body CreateBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "create", err)
				}
				err = ValidateCreateBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "create", err)
				}
				return nil, NewCreateBadRequest(&body)
			case "invalid-credential":
				return nil, NewCreateInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "create", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body CreateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create", err)
			}
			err = ValidateCreateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create", err)
			}
			return nil, NewCreateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create", err)
			}
			err = ValidateCreateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create", err)
			}
			return nil, NewCreateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create", err)
			}
			err = ValidateCreateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create", err)
			}
			return nil, NewCreateNotImplemented(&body)
		case http.StatusConflict:
			var (
				body CreateAlreadyCreatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create", err)
			}
			err = ValidateCreateAlreadyCreatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create", err)
			}
			return nil, NewCreateAlreadyCreated(&body)
		case http.StatusNotFound:
			var (
				body CreateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create", err)
			}
			err = ValidateCreateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create", err)
			}
			return nil, NewCreateNotFound(&body)
		case http.StatusUnauthorized:
			return nil, NewCreateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "service" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "read", "*service.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadServicePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeReadResponse returns a decoder for responses returned by the service
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *service.InvalidCredentialsT): http.StatusBadRequest
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeReadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("service", "read", err)
			}
			p := NewReadServiceStatusRTOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &serviceviews.ServiceStatusRT{Projected: p, View: view}
			if err = serviceviews.ValidateServiceStatusRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "read", err)
			}
			res := service.NewServiceStatusRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body ReadBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "read", err)
				}
				err = ValidateReadBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "read", err)
				}
				return nil, NewReadBadRequest(&body)
			case "invalid-credential":
				return nil, NewReadInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "read", resp.StatusCode, string(body))
			}
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "service" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "update", "*service.UpdatePayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateServicePath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the service
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*service.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "update", "*service.UpdatePayload", v)
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
		if p.ForceCreate != nil {
			values.Add("force-create", fmt.Sprintf("%v", *p.ForceCreate))
		}
		req.URL.RawQuery = values.Encode()
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("service", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the service
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *service.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
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
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "update", err)
			}
			p := NewUpdateServiceStatusRTOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &serviceviews.ServiceStatusRT{Projected: p, View: view}
			if err = serviceviews.ValidateServiceStatusRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "update", err)
			}
			res := service.NewServiceStatusRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body UpdateBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "update", err)
				}
				err = ValidateUpdateBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "update", err)
				}
				return nil, NewUpdateBadRequest(&body)
			case "invalid-credential":
				return nil, NewUpdateInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "update", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body UpdateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "update", err)
			}
			err = ValidateUpdateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "update", err)
			}
			return nil, NewUpdateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body UpdateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "update", err)
			}
			err = ValidateUpdateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "update", err)
			}
			return nil, NewUpdateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UpdateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "update", err)
			}
			err = ValidateUpdateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "update", err)
			}
			return nil, NewUpdateNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		case http.StatusUnauthorized:
			return nil, NewUpdateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "service" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "delete", "*service.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteServicePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the service
// delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*service.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "delete", "*service.DeletePayload", v)
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

// DecodeDeleteResponse returns a decoder for responses returned by the service
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *service.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
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
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body DeleteBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "delete", err)
				}
				err = ValidateDeleteBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "delete", err)
				}
				return nil, NewDeleteBadRequest(&body)
			case "invalid-credential":
				return nil, NewDeleteInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "delete", resp.StatusCode, string(body))
			}
		case http.StatusForbidden:
			var (
				body DeleteInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "delete", err)
			}
			err = ValidateDeleteInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "delete", err)
			}
			return nil, NewDeleteInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body DeleteNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "delete", err)
			}
			err = ValidateDeleteNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "delete", err)
			}
			return nil, NewDeleteNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewDeleteNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "delete", resp.StatusCode, string(body))
		}
	}
}

// BuildListOrdersRequest instantiates a HTTP request object with method and
// path set to call the "service" service "listOrders" endpoint
func (c *Client) BuildListOrdersRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.ListOrdersPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "listOrders", "*service.ListOrdersPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListOrdersServicePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "listOrders", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListOrdersRequest returns an encoder for requests sent to the service
// listOrders server.
func EncodeListOrdersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*service.ListOrdersPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "listOrders", "*service.ListOrdersPayload", v)
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

// DecodeListOrdersResponse returns a decoder for responses returned by the
// service listOrders endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListOrdersResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *service.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeListOrdersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
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
				body ListOrdersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "listOrders", err)
			}
			p := NewListOrdersOrderListRTOK(&body)
			view := "default"
			vres := &serviceviews.OrderListRT{Projected: p, View: view}
			if err = serviceviews.ValidateOrderListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "listOrders", err)
			}
			res := service.NewOrderListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body ListOrdersBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "listOrders", err)
				}
				err = ValidateListOrdersBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "listOrders", err)
				}
				return nil, NewListOrdersBadRequest(&body)
			case "invalid-credential":
				return nil, NewListOrdersInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "listOrders", resp.StatusCode, string(body))
			}
		case http.StatusForbidden:
			var (
				body ListOrdersInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "listOrders", err)
			}
			err = ValidateListOrdersInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "listOrders", err)
			}
			return nil, NewListOrdersInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListOrdersNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "listOrders", err)
			}
			err = ValidateListOrdersNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "listOrders", err)
			}
			return nil, NewListOrdersNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ListOrdersNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "listOrders", err)
			}
			err = ValidateListOrdersNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "listOrders", err)
			}
			return nil, NewListOrdersNotFound(&body)
		case http.StatusUnauthorized:
			return nil, NewListOrdersNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "listOrders", resp.StatusCode, string(body))
		}
	}
}

// unmarshalServiceListItemResponseBodyToServiceviewsServiceListItemView builds
// a value of type *serviceviews.ServiceListItemView from a value of type
// *ServiceListItemResponseBody.
func unmarshalServiceListItemResponseBodyToServiceviewsServiceListItemView(v *ServiceListItemResponseBody) *serviceviews.ServiceListItemView {
	res := &serviceviews.ServiceListItemView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
	}
	if v.Provider != nil {
		res.Provider = unmarshalRefTResponseBodyToServiceviewsRefTView(v.Provider)
	}
	res.Links = unmarshalSelfTResponseBodyToServiceviewsSelfTView(v.Links)

	return res
}

// unmarshalRefTResponseBodyToServiceviewsRefTView builds a value of type
// *serviceviews.RefTView from a value of type *RefTResponseBody.
func unmarshalRefTResponseBodyToServiceviewsRefTView(v *RefTResponseBody) *serviceviews.RefTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.RefTView{
		ID: v.ID,
	}
	if v.Links != nil {
		res.Links = unmarshalSelfTResponseBodyToServiceviewsSelfTView(v.Links)
	}

	return res
}

// unmarshalSelfTResponseBodyToServiceviewsSelfTView builds a value of type
// *serviceviews.SelfTView from a value of type *SelfTResponseBody.
func unmarshalSelfTResponseBodyToServiceviewsSelfTView(v *SelfTResponseBody) *serviceviews.SelfTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.SelfTView{
		Self: v.Self,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = unmarshalDescribedByTResponseBodyToServiceviewsDescribedByTView(v.DescribedBy)
	}

	return res
}

// unmarshalDescribedByTResponseBodyToServiceviewsDescribedByTView builds a
// value of type *serviceviews.DescribedByTView from a value of type
// *DescribedByTResponseBody.
func unmarshalDescribedByTResponseBodyToServiceviewsDescribedByTView(v *DescribedByTResponseBody) *serviceviews.DescribedByTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.DescribedByTView{
		Href: v.Href,
		Type: v.Type,
	}

	return res
}

// unmarshalNavTResponseBodyToServiceviewsNavTView builds a value of type
// *serviceviews.NavTView from a value of type *NavTResponseBody.
func unmarshalNavTResponseBodyToServiceviewsNavTView(v *NavTResponseBody) *serviceviews.NavTView {
	res := &serviceviews.NavTView{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}

// marshalServiceParameterTToParameterTRequestBodyRequestBody builds a value of
// type *ParameterTRequestBodyRequestBody from a value of type
// *service.ParameterT.
func marshalServiceParameterTToParameterTRequestBodyRequestBody(v *service.ParameterT) *ParameterTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &ParameterTRequestBodyRequestBody{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// marshalServiceReferenceTToReferenceTRequestBodyRequestBody builds a value of
// type *ReferenceTRequestBodyRequestBody from a value of type
// *service.ReferenceT.
func marshalServiceReferenceTToReferenceTRequestBodyRequestBody(v *service.ReferenceT) *ReferenceTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &ReferenceTRequestBodyRequestBody{
		Title: v.Title,
		URI:   v.URI,
	}

	return res
}

// marshalServiceWorkflowTToWorkflowTRequestBodyRequestBody builds a value of
// type *WorkflowTRequestBodyRequestBody from a value of type
// *service.WorkflowT.
func marshalServiceWorkflowTToWorkflowTRequestBodyRequestBody(v *service.WorkflowT) *WorkflowTRequestBodyRequestBody {
	res := &WorkflowTRequestBodyRequestBody{
		Type: v.Type,
		Argo: v.Argo,
		Opts: v.Opts,
	}
	if v.Basic != nil {
		res.Basic = marshalServiceBasicWorkflowOptsTToBasicWorkflowOptsTRequestBodyRequestBody(v.Basic)
	}

	return res
}

// marshalServiceBasicWorkflowOptsTToBasicWorkflowOptsTRequestBodyRequestBody
// builds a value of type *BasicWorkflowOptsTRequestBodyRequestBody from a
// value of type *service.BasicWorkflowOptsT.
func marshalServiceBasicWorkflowOptsTToBasicWorkflowOptsTRequestBodyRequestBody(v *service.BasicWorkflowOptsT) *BasicWorkflowOptsTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &BasicWorkflowOptsTRequestBodyRequestBody{
		Image: v.Image,
	}
	if v.Command != nil {
		res.Command = make([]string, len(v.Command))
		for i, val := range v.Command {
			res.Command[i] = val
		}
	}
	if v.Memory != nil {
		res.Memory = marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody(v.Memory)
	}
	if v.CPU != nil {
		res.CPU = marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody(v.CPU)
	}

	return res
}

// marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody builds
// a value of type *ResourceMemoryTRequestBodyRequestBody from a value of type
// *service.ResourceMemoryT.
func marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody(v *service.ResourceMemoryT) *ResourceMemoryTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &ResourceMemoryTRequestBodyRequestBody{
		Request: v.Request,
		Limit:   v.Limit,
	}

	return res
}

// marshalServiceParameterDefTToParameterDefT builds a value of type
// *ParameterDefT from a value of type *service.ParameterDefT.
func marshalServiceParameterDefTToParameterDefT(v *service.ParameterDefT) *ParameterDefT {
	res := &ParameterDefT{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
	}
	if v.Options != nil {
		res.Options = make([]*ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = marshalServiceParameterOptTToParameterOptT(val)
		}
	}

	return res
}

// marshalServiceParameterOptTToParameterOptT builds a value of type
// *ParameterOptT from a value of type *service.ParameterOptT.
func marshalServiceParameterOptTToParameterOptT(v *service.ParameterOptT) *ParameterOptT {
	if v == nil {
		return nil
	}
	res := &ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// marshalParameterTRequestBodyRequestBodyToServiceParameterT builds a value of
// type *service.ParameterT from a value of type
// *ParameterTRequestBodyRequestBody.
func marshalParameterTRequestBodyRequestBodyToServiceParameterT(v *ParameterTRequestBodyRequestBody) *service.ParameterT {
	if v == nil {
		return nil
	}
	res := &service.ParameterT{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// marshalReferenceTRequestBodyRequestBodyToServiceReferenceT builds a value of
// type *service.ReferenceT from a value of type
// *ReferenceTRequestBodyRequestBody.
func marshalReferenceTRequestBodyRequestBodyToServiceReferenceT(v *ReferenceTRequestBodyRequestBody) *service.ReferenceT {
	if v == nil {
		return nil
	}
	res := &service.ReferenceT{
		Title: v.Title,
		URI:   v.URI,
	}

	return res
}

// marshalWorkflowTRequestBodyRequestBodyToServiceWorkflowT builds a value of
// type *service.WorkflowT from a value of type
// *WorkflowTRequestBodyRequestBody.
func marshalWorkflowTRequestBodyRequestBodyToServiceWorkflowT(v *WorkflowTRequestBodyRequestBody) *service.WorkflowT {
	res := &service.WorkflowT{
		Type: v.Type,
		Argo: v.Argo,
		Opts: v.Opts,
	}
	if v.Basic != nil {
		res.Basic = marshalBasicWorkflowOptsTRequestBodyRequestBodyToServiceBasicWorkflowOptsT(v.Basic)
	}

	return res
}

// marshalBasicWorkflowOptsTRequestBodyRequestBodyToServiceBasicWorkflowOptsT
// builds a value of type *service.BasicWorkflowOptsT from a value of type
// *BasicWorkflowOptsTRequestBodyRequestBody.
func marshalBasicWorkflowOptsTRequestBodyRequestBodyToServiceBasicWorkflowOptsT(v *BasicWorkflowOptsTRequestBodyRequestBody) *service.BasicWorkflowOptsT {
	if v == nil {
		return nil
	}
	res := &service.BasicWorkflowOptsT{
		Image: v.Image,
	}
	if v.Command != nil {
		res.Command = make([]string, len(v.Command))
		for i, val := range v.Command {
			res.Command[i] = val
		}
	}
	if v.Memory != nil {
		res.Memory = marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT(v.Memory)
	}
	if v.CPU != nil {
		res.CPU = marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT(v.CPU)
	}

	return res
}

// marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT builds
// a value of type *service.ResourceMemoryT from a value of type
// *ResourceMemoryTRequestBodyRequestBody.
func marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT(v *ResourceMemoryTRequestBodyRequestBody) *service.ResourceMemoryT {
	if v == nil {
		return nil
	}
	res := &service.ResourceMemoryT{
		Request: v.Request,
		Limit:   v.Limit,
	}

	return res
}

// marshalParameterDefTToServiceParameterDefT builds a value of type
// *service.ParameterDefT from a value of type *ParameterDefT.
func marshalParameterDefTToServiceParameterDefT(v *ParameterDefT) *service.ParameterDefT {
	res := &service.ParameterDefT{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
	}
	if v.Options != nil {
		res.Options = make([]*service.ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = marshalParameterOptTToServiceParameterOptT(val)
		}
	}

	return res
}

// marshalParameterOptTToServiceParameterOptT builds a value of type
// *service.ParameterOptT from a value of type *ParameterOptT.
func marshalParameterOptTToServiceParameterOptT(v *ParameterOptT) *service.ParameterOptT {
	if v == nil {
		return nil
	}
	res := &service.ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// unmarshalParameterTResponseBodyToServiceviewsParameterTView builds a value
// of type *serviceviews.ParameterTView from a value of type
// *ParameterTResponseBody.
func unmarshalParameterTResponseBodyToServiceviewsParameterTView(v *ParameterTResponseBody) *serviceviews.ParameterTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.ParameterTView{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// unmarshalParameterDefTResponseBodyToServiceviewsParameterDefTView builds a
// value of type *serviceviews.ParameterDefTView from a value of type
// *ParameterDefTResponseBody.
func unmarshalParameterDefTResponseBodyToServiceviewsParameterDefTView(v *ParameterDefTResponseBody) *serviceviews.ParameterDefTView {
	res := &serviceviews.ParameterDefTView{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
	}
	if v.Options != nil {
		res.Options = make([]*serviceviews.ParameterOptTView, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = unmarshalParameterOptTResponseBodyToServiceviewsParameterOptTView(val)
		}
	}

	return res
}

// unmarshalParameterOptTResponseBodyToServiceviewsParameterOptTView builds a
// value of type *serviceviews.ParameterOptTView from a value of type
// *ParameterOptTResponseBody.
func unmarshalParameterOptTResponseBodyToServiceviewsParameterOptTView(v *ParameterOptTResponseBody) *serviceviews.ParameterOptTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.ParameterOptTView{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// unmarshalOrderListItemResponseBodyToServiceviewsOrderListItemView builds a
// value of type *serviceviews.OrderListItemView from a value of type
// *OrderListItemResponseBody.
func unmarshalOrderListItemResponseBodyToServiceviewsOrderListItemView(v *OrderListItemResponseBody) *serviceviews.OrderListItemView {
	res := &serviceviews.OrderListItemView{
		ID:         v.ID,
		Name:       v.Name,
		Status:     v.Status,
		OrderedAt:  v.OrderedAt,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		ServiceID:  v.ServiceID,
		AccountID:  v.AccountID,
	}
	res.Links = unmarshalSelfTResponseBodyToServiceviewsSelfTView(v.Links)

	return res
}
