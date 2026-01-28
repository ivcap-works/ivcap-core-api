// Copyright 2026 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	servicex "github.com/ivcap-works/ivcap-core-api/gen/servicex"
	servicexviews "github.com/ivcap-works/ivcap-core-api/gen/servicex/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "servicex" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListServicexPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("servicex", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the servicex list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*servicex.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("servicex", "list", "*servicex.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the servicex
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *servicex.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *servicex.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *servicex.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *servicex.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *servicex.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *servicex.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("servicex", "list", err)
			}
			p := NewListXServiceListRTOK(&body)
			view := "default"
			vres := &servicexviews.XServiceListRT{Projected: p, View: view}
			if err = servicexviews.ValidateXServiceListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("servicex", "list", err)
			}
			res := servicex.NewXServiceListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("servicex", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateServiceRequest instantiates a HTTP request object with method and
// path set to call the "servicex" service "create_service" endpoint
func (c *Client) BuildCreateServiceRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateServiceServicexPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("servicex", "create_service", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateServiceRequest returns an encoder for requests sent to the
// servicex create_service server.
func EncodeCreateServiceRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*servicex.CreateServicePayload)
		if !ok {
			return goahttp.ErrInvalidType("servicex", "create_service", "*servicex.CreateServicePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateServiceRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("servicex", "create_service", err)
		}
		return nil
	}
}

// DecodeCreateServiceResponse returns a decoder for responses returned by the
// servicex create_service endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateServiceResponse may return the following errors:
//   - "bad-request" (type *servicex.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *servicex.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *servicex.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *servicex.NotImplementedT): http.StatusNotImplemented
//   - "already-created" (type *servicex.ResourceAlreadyCreatedT): http.StatusConflict
//   - "not-found" (type *servicex.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *servicex.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *servicex.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateServiceResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusCreated:
			var (
				body CreateServiceResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			res := NewCreateServiceXServiceStatusRTCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateServiceBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			return nil, NewCreateServiceBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body CreateServiceInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			return nil, NewCreateServiceInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateServiceInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			return nil, NewCreateServiceInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateServiceNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			return nil, NewCreateServiceNotImplemented(&body)
		case http.StatusConflict:
			var (
				body CreateServiceAlreadyCreatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceAlreadyCreatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			return nil, NewCreateServiceAlreadyCreated(&body)
		case http.StatusNotFound:
			var (
				body CreateServiceNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "create_service", err)
			}
			err = ValidateCreateServiceNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "create_service", err)
			}
			return nil, NewCreateServiceNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewCreateServiceNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewCreateServiceNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("servicex", "create_service", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "servicex" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*servicex.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("servicex", "read", "*servicex.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadServicexPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("servicex", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the servicex read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*servicex.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("servicex", "read", "*servicex.ReadPayload", v)
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

// DecodeReadResponse returns a decoder for responses returned by the servicex
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *servicex.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *servicex.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *servicex.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *servicex.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *servicex.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *servicex.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("servicex", "read", err)
			}
			err = ValidateReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "read", err)
			}
			res := NewReadXServiceStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "read", err)
			}
			err = ValidateReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "read", err)
			}
			return nil, NewReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("servicex", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "servicex" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*servicex.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("servicex", "update", "*servicex.UpdatePayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateServicexPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("servicex", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the servicex
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*servicex.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("servicex", "update", "*servicex.UpdatePayload", v)
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
			return goahttp.ErrEncodingError("servicex", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// servicex update endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "bad-request" (type *servicex.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *servicex.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *servicex.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *servicex.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *servicex.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *servicex.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *servicex.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "update", err)
			}
			res := NewUpdateXServiceStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body UpdateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "update", err)
			}
			err = ValidateUpdateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "update", err)
			}
			return nil, NewUpdateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body UpdateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "update", err)
			}
			err = ValidateUpdateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "update", err)
			}
			return nil, NewUpdateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UpdateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "update", err)
			}
			err = ValidateUpdateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "update", err)
			}
			return nil, NewUpdateNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewUpdateNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewUpdateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("servicex", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "servicex" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*servicex.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("servicex", "delete", "*servicex.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteServicexPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("servicex", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the servicex
// delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*servicex.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("servicex", "delete", "*servicex.DeletePayload", v)
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

// DecodeDeleteResponse returns a decoder for responses returned by the
// servicex delete endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "bad-request" (type *servicex.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *servicex.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *servicex.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *servicex.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *servicex.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("servicex", "delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		case http.StatusForbidden:
			var (
				body DeleteInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "delete", err)
			}
			err = ValidateDeleteInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "delete", err)
			}
			return nil, NewDeleteInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body DeleteNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("servicex", "delete", err)
			}
			err = ValidateDeleteNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("servicex", "delete", err)
			}
			return nil, NewDeleteNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewDeleteNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewDeleteNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("servicex", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalXServiceListItemResponseBodyToServicexviewsXServiceListItemView
// builds a value of type *servicexviews.XServiceListItemView from a value of
// type *XServiceListItemResponseBody.
func unmarshalXServiceListItemResponseBodyToServicexviewsXServiceListItemView(v *XServiceListItemResponseBody) *servicexviews.XServiceListItemView {
	res := &servicexviews.XServiceListItemView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Banner:      v.Banner,
		PublishedAt: v.PublishedAt,
		Policy:      v.Policy,
		Account:     v.Account,
		Href:        v.Href,
	}

	return res
}

// unmarshalLinkTResponseBodyToServicexviewsLinkTView builds a value of type
// *servicexviews.LinkTView from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToServicexviewsLinkTView(v *LinkTResponseBody) *servicexviews.LinkTView {
	res := &servicexviews.LinkTView{
		Rel:  v.Rel,
		Type: v.Type,
		Href: v.Href,
	}

	return res
}

// marshalServicexXReferenceTToXReferenceTRequestBodyRequestBody builds a value
// of type *XReferenceTRequestBodyRequestBody from a value of type
// *servicex.XReferenceT.
func marshalServicexXReferenceTToXReferenceTRequestBodyRequestBody(v *servicex.XReferenceT) *XReferenceTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &XReferenceTRequestBodyRequestBody{
		Title: v.Title,
		URI:   v.URI,
	}

	return res
}

// marshalServicexXWorkflowTToXWorkflowTRequestBodyRequestBody builds a value
// of type *XWorkflowTRequestBodyRequestBody from a value of type
// *servicex.XWorkflowT.
func marshalServicexXWorkflowTToXWorkflowTRequestBodyRequestBody(v *servicex.XWorkflowT) *XWorkflowTRequestBodyRequestBody {
	res := &XWorkflowTRequestBodyRequestBody{
		Type: v.Type,
		Argo: v.Argo,
	}
	if v.Basic != nil {
		res.Basic = marshalServicexXBasicWorkflowOptsTToXBasicWorkflowOptsTRequestBodyRequestBody(v.Basic)
	}

	return res
}

// marshalServicexXBasicWorkflowOptsTToXBasicWorkflowOptsTRequestBodyRequestBody
// builds a value of type *XBasicWorkflowOptsTRequestBodyRequestBody from a
// value of type *servicex.XBasicWorkflowOptsT.
func marshalServicexXBasicWorkflowOptsTToXBasicWorkflowOptsTRequestBodyRequestBody(v *servicex.XBasicWorkflowOptsT) *XBasicWorkflowOptsTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &XBasicWorkflowOptsTRequestBodyRequestBody{
		Image:           v.Image,
		ImagePullPolicy: v.ImagePullPolicy,
		GpuType:         v.GpuType,
		GpuNumber:       v.GpuNumber,
		SharedMemory:    v.SharedMemory,
	}
	{
		var zero string
		if res.ImagePullPolicy == zero {
			res.ImagePullPolicy = "IfNotPresent"
		}
	}
	if v.Command != nil {
		res.Command = make([]string, len(v.Command))
		for i, val := range v.Command {
			res.Command[i] = val
		}
	} else {
		res.Command = []string{}
	}
	if v.Memory != nil {
		res.Memory = marshalServicexXResourceMemoryTToXResourceMemoryTRequestBodyRequestBody(v.Memory)
	}
	if v.CPU != nil {
		res.CPU = marshalServicexXResourceMemoryTToXResourceMemoryTRequestBodyRequestBody(v.CPU)
	}
	if v.EphemeralStorage != nil {
		res.EphemeralStorage = marshalServicexXResourceMemoryTToXResourceMemoryTRequestBodyRequestBody(v.EphemeralStorage)
	}

	return res
}

// marshalServicexXResourceMemoryTToXResourceMemoryTRequestBodyRequestBody
// builds a value of type *XResourceMemoryTRequestBodyRequestBody from a value
// of type *servicex.XResourceMemoryT.
func marshalServicexXResourceMemoryTToXResourceMemoryTRequestBodyRequestBody(v *servicex.XResourceMemoryT) *XResourceMemoryTRequestBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &XResourceMemoryTRequestBodyRequestBody{
		Request: v.Request,
		Limit:   v.Limit,
	}

	return res
}

// marshalServicexParameterDefTToParameterDefT builds a value of type
// *ParameterDefT from a value of type *servicex.ParameterDefT.
func marshalServicexParameterDefTToParameterDefT(v *servicex.ParameterDefT) *ParameterDefT {
	res := &ParameterDefT{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
		Unary:       v.Unary,
	}
	if v.Options != nil {
		res.Options = make([]*ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			if val == nil {
				res.Options[i] = nil
				continue
			}
			res.Options[i] = marshalServicexParameterOptTToParameterOptT(val)
		}
	}

	return res
}

// marshalServicexParameterOptTToParameterOptT builds a value of type
// *ParameterOptT from a value of type *servicex.ParameterOptT.
func marshalServicexParameterOptTToParameterOptT(v *servicex.ParameterOptT) *ParameterOptT {
	if v == nil {
		return nil
	}
	res := &ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// marshalXReferenceTRequestBodyRequestBodyToServicexXReferenceT builds a value
// of type *servicex.XReferenceT from a value of type
// *XReferenceTRequestBodyRequestBody.
func marshalXReferenceTRequestBodyRequestBodyToServicexXReferenceT(v *XReferenceTRequestBodyRequestBody) *servicex.XReferenceT {
	if v == nil {
		return nil
	}
	res := &servicex.XReferenceT{
		Title: v.Title,
		URI:   v.URI,
	}

	return res
}

// marshalXWorkflowTRequestBodyRequestBodyToServicexXWorkflowT builds a value
// of type *servicex.XWorkflowT from a value of type
// *XWorkflowTRequestBodyRequestBody.
func marshalXWorkflowTRequestBodyRequestBodyToServicexXWorkflowT(v *XWorkflowTRequestBodyRequestBody) *servicex.XWorkflowT {
	res := &servicex.XWorkflowT{
		Type: v.Type,
		Argo: v.Argo,
	}
	if v.Basic != nil {
		res.Basic = marshalXBasicWorkflowOptsTRequestBodyRequestBodyToServicexXBasicWorkflowOptsT(v.Basic)
	}

	return res
}

// marshalXBasicWorkflowOptsTRequestBodyRequestBodyToServicexXBasicWorkflowOptsT
// builds a value of type *servicex.XBasicWorkflowOptsT from a value of type
// *XBasicWorkflowOptsTRequestBodyRequestBody.
func marshalXBasicWorkflowOptsTRequestBodyRequestBodyToServicexXBasicWorkflowOptsT(v *XBasicWorkflowOptsTRequestBodyRequestBody) *servicex.XBasicWorkflowOptsT {
	if v == nil {
		return nil
	}
	res := &servicex.XBasicWorkflowOptsT{
		Image:           v.Image,
		ImagePullPolicy: v.ImagePullPolicy,
		GpuType:         v.GpuType,
		GpuNumber:       v.GpuNumber,
		SharedMemory:    v.SharedMemory,
	}
	{
		var zero string
		if res.ImagePullPolicy == zero {
			res.ImagePullPolicy = "IfNotPresent"
		}
	}
	if v.Command != nil {
		res.Command = make([]string, len(v.Command))
		for i, val := range v.Command {
			res.Command[i] = val
		}
	} else {
		res.Command = []string{}
	}
	if v.Memory != nil {
		res.Memory = marshalXResourceMemoryTRequestBodyRequestBodyToServicexXResourceMemoryT(v.Memory)
	}
	if v.CPU != nil {
		res.CPU = marshalXResourceMemoryTRequestBodyRequestBodyToServicexXResourceMemoryT(v.CPU)
	}
	if v.EphemeralStorage != nil {
		res.EphemeralStorage = marshalXResourceMemoryTRequestBodyRequestBodyToServicexXResourceMemoryT(v.EphemeralStorage)
	}

	return res
}

// marshalXResourceMemoryTRequestBodyRequestBodyToServicexXResourceMemoryT
// builds a value of type *servicex.XResourceMemoryT from a value of type
// *XResourceMemoryTRequestBodyRequestBody.
func marshalXResourceMemoryTRequestBodyRequestBodyToServicexXResourceMemoryT(v *XResourceMemoryTRequestBodyRequestBody) *servicex.XResourceMemoryT {
	if v == nil {
		return nil
	}
	res := &servicex.XResourceMemoryT{
		Request: v.Request,
		Limit:   v.Limit,
	}

	return res
}

// marshalParameterDefTToServicexParameterDefT builds a value of type
// *servicex.ParameterDefT from a value of type *ParameterDefT.
func marshalParameterDefTToServicexParameterDefT(v *ParameterDefT) *servicex.ParameterDefT {
	res := &servicex.ParameterDefT{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
		Unary:       v.Unary,
	}
	if v.Options != nil {
		res.Options = make([]*servicex.ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			if val == nil {
				res.Options[i] = nil
				continue
			}
			res.Options[i] = marshalParameterOptTToServicexParameterOptT(val)
		}
	}

	return res
}

// marshalParameterOptTToServicexParameterOptT builds a value of type
// *servicex.ParameterOptT from a value of type *ParameterOptT.
func marshalParameterOptTToServicexParameterOptT(v *ParameterOptT) *servicex.ParameterOptT {
	if v == nil {
		return nil
	}
	res := &servicex.ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// unmarshalLinkTResponseBodyToServicexLinkT builds a value of type
// *servicex.LinkT from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToServicexLinkT(v *LinkTResponseBody) *servicex.LinkT {
	res := &servicex.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}

// unmarshalParameterDefTResponseBodyToServicexParameterDefT builds a value of
// type *servicex.ParameterDefT from a value of type *ParameterDefTResponseBody.
func unmarshalParameterDefTResponseBodyToServicexParameterDefT(v *ParameterDefTResponseBody) *servicex.ParameterDefT {
	res := &servicex.ParameterDefT{
		Name:        *v.Name,
		Label:       v.Label,
		Type:        *v.Type,
		Description: *v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
		Unary:       v.Unary,
	}
	if v.Options != nil {
		res.Options = make([]*servicex.ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			if val == nil {
				res.Options[i] = nil
				continue
			}
			res.Options[i] = unmarshalParameterOptTResponseBodyToServicexParameterOptT(val)
		}
	}

	return res
}

// unmarshalParameterOptTResponseBodyToServicexParameterOptT builds a value of
// type *servicex.ParameterOptT from a value of type *ParameterOptTResponseBody.
func unmarshalParameterOptTResponseBodyToServicexParameterOptT(v *ParameterOptTResponseBody) *servicex.ParameterOptT {
	if v == nil {
		return nil
	}
	res := &servicex.ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}
