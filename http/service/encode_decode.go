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
	"os"
	"strconv"
	"strings"

	service "github.com/ivcap-works/ivcap-core-api/gen/service"
	serviceviews "github.com/ivcap-works/ivcap-core-api/gen/service/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildServiceListRequest instantiates a HTTP request object with method and
// path set to call the "service" service "service-list" endpoint
func (c *Client) BuildServiceListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ServiceListServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "service-list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeServiceListRequest returns an encoder for requests sent to the service
// service-list server.
func EncodeServiceListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ServiceListPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "service-list", "*service.ServiceListPayload", v)
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

// DecodeServiceListResponse returns a decoder for responses returned by the
// service service-list endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeServiceListResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeServiceListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ServiceListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-list", err)
			}
			p := NewServiceListRTViewOK(&body)
			view := "default"
			vres := &serviceviews.ServiceListRT{Projected: p, View: view}
			if err = serviceviews.ValidateServiceListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "service-list", err)
			}
			res := service.NewServiceListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ServiceListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-list", err)
			}
			err = ValidateServiceListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-list", err)
			}
			return nil, NewServiceListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ServiceListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-list", err)
			}
			err = ValidateServiceListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-list", err)
			}
			return nil, NewServiceListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ServiceListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-list", err)
			}
			err = ValidateServiceListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-list", err)
			}
			return nil, NewServiceListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ServiceListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-list", err)
			}
			err = ValidateServiceListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-list", err)
			}
			return nil, NewServiceListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewServiceListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewServiceListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "service-list", resp.StatusCode, string(body))
		}
	}
}

// BuildServiceCreateRequest instantiates a HTTP request object with method and
// path set to call the "service" service "service-create" endpoint
func (c *Client) BuildServiceCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ServiceCreateServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "service-create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeServiceCreateRequest returns an encoder for requests sent to the
// service service-create server.
func EncodeServiceCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ServiceCreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "service-create", "*service.ServiceCreatePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewServiceCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("service", "service-create", err)
		}
		return nil
	}
}

// DecodeServiceCreateResponse returns a decoder for responses returned by the
// service service-create endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeServiceCreateResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "already-created" (type *service.ResourceAlreadyCreatedT): http.StatusConflict
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeServiceCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ServiceCreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			res := NewServiceCreateServiceStatusRTCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ServiceCreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			return nil, NewServiceCreateBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ServiceCreateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			return nil, NewServiceCreateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ServiceCreateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			return nil, NewServiceCreateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ServiceCreateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			return nil, NewServiceCreateNotImplemented(&body)
		case http.StatusConflict:
			var (
				body ServiceCreateAlreadyCreatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateAlreadyCreatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			return nil, NewServiceCreateAlreadyCreated(&body)
		case http.StatusNotFound:
			var (
				body ServiceCreateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-create", err)
			}
			err = ValidateServiceCreateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-create", err)
			}
			return nil, NewServiceCreateNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewServiceCreateNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewServiceCreateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "service-create", resp.StatusCode, string(body))
		}
	}
}

// BuildServiceReadRequest instantiates a HTTP request object with method and
// path set to call the "service" service "service-read" endpoint
func (c *Client) BuildServiceReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.ServiceReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "service-read", "*service.ServiceReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ServiceReadServicePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "service-read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeServiceReadRequest returns an encoder for requests sent to the service
// service-read server.
func EncodeServiceReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ServiceReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "service-read", "*service.ServiceReadPayload", v)
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

// DecodeServiceReadResponse returns a decoder for responses returned by the
// service service-read endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeServiceReadResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeServiceReadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ServiceReadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-read", err)
			}
			err = ValidateServiceReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-read", err)
			}
			res := NewServiceReadServiceStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ServiceReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-read", err)
			}
			err = ValidateServiceReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-read", err)
			}
			return nil, NewServiceReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ServiceReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-read", err)
			}
			err = ValidateServiceReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-read", err)
			}
			return nil, NewServiceReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ServiceReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-read", err)
			}
			err = ValidateServiceReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-read", err)
			}
			return nil, NewServiceReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ServiceReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-read", err)
			}
			err = ValidateServiceReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-read", err)
			}
			return nil, NewServiceReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewServiceReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewServiceReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "service-read", resp.StatusCode, string(body))
		}
	}
}

// BuildServiceUpdateRequest instantiates a HTTP request object with method and
// path set to call the "service" service "service-update" endpoint
func (c *Client) BuildServiceUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.ServiceUpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "service-update", "*service.ServiceUpdatePayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ServiceUpdateServicePath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "service-update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeServiceUpdateRequest returns an encoder for requests sent to the
// service service-update server.
func EncodeServiceUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ServiceUpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "service-update", "*service.ServiceUpdatePayload", v)
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
		body := NewServiceUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("service", "service-update", err)
		}
		return nil
	}
}

// DecodeServiceUpdateResponse returns a decoder for responses returned by the
// service service-update endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeServiceUpdateResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeServiceUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ServiceUpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-update", err)
			}
			err = ValidateServiceUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-update", err)
			}
			res := NewServiceUpdateServiceStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ServiceUpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-update", err)
			}
			err = ValidateServiceUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-update", err)
			}
			return nil, NewServiceUpdateBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ServiceUpdateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-update", err)
			}
			err = ValidateServiceUpdateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-update", err)
			}
			return nil, NewServiceUpdateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ServiceUpdateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-update", err)
			}
			err = ValidateServiceUpdateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-update", err)
			}
			return nil, NewServiceUpdateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ServiceUpdateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-update", err)
			}
			err = ValidateServiceUpdateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-update", err)
			}
			return nil, NewServiceUpdateNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ServiceUpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-update", err)
			}
			err = ValidateServiceUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-update", err)
			}
			return nil, NewServiceUpdateNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewServiceUpdateNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewServiceUpdateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "service-update", resp.StatusCode, string(body))
		}
	}
}

// BuildServiceDeleteRequest instantiates a HTTP request object with method and
// path set to call the "service" service "service-delete" endpoint
func (c *Client) BuildServiceDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*service.ServiceDeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "service-delete", "*service.ServiceDeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ServiceDeleteServicePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "service-delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeServiceDeleteRequest returns an encoder for requests sent to the
// service service-delete server.
func EncodeServiceDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ServiceDeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "service-delete", "*service.ServiceDeletePayload", v)
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

// DecodeServiceDeleteResponse returns a decoder for responses returned by the
// service service-delete endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeServiceDeleteResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeServiceDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ServiceDeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-delete", err)
			}
			err = ValidateServiceDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-delete", err)
			}
			return nil, NewServiceDeleteBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ServiceDeleteInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-delete", err)
			}
			err = ValidateServiceDeleteInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-delete", err)
			}
			return nil, NewServiceDeleteInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ServiceDeleteNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "service-delete", err)
			}
			err = ValidateServiceDeleteNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "service-delete", err)
			}
			return nil, NewServiceDeleteNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewServiceDeleteNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewServiceDeleteNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "service-delete", resp.StatusCode, string(body))
		}
	}
}

// BuildJobListRequest instantiates a HTTP request object with method and path
// set to call the "service" service "job-list" endpoint
func (c *Client) BuildJobListRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		serviceID string
	)
	{
		p, ok := v.(*service.JobListPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "job-list", "*service.JobListPayload", v)
		}
		serviceID = p.ServiceID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: JobListServicePath(serviceID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "job-list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeJobListRequest returns an encoder for requests sent to the service
// job-list server.
func EncodeJobListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.JobListPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "job-list", "*service.JobListPayload", v)
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

// DecodeJobListResponse returns a decoder for responses returned by the
// service job-list endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeJobListResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeJobListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body JobListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-list", err)
			}
			p := NewJobListRTViewOK(&body)
			view := "default"
			vres := &serviceviews.JobListRT{Projected: p, View: view}
			if err = serviceviews.ValidateJobListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("service", "job-list", err)
			}
			res := service.NewJobListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body JobListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-list", err)
			}
			err = ValidateJobListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-list", err)
			}
			return nil, NewJobListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body JobListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-list", err)
			}
			err = ValidateJobListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-list", err)
			}
			return nil, NewJobListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body JobListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-list", err)
			}
			err = ValidateJobListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-list", err)
			}
			return nil, NewJobListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body JobListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-list", err)
			}
			err = ValidateJobListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-list", err)
			}
			return nil, NewJobListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewJobListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewJobListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "job-list", resp.StatusCode, string(body))
		}
	}
}

// BuildJobCreateRequest instantiates a HTTP request object with method and
// path set to call the "service" service "job-create" endpoint
func (c *Client) BuildJobCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		serviceID string
		body      io.Reader
	)
	{
		rd, ok := v.(*service.JobCreateRequestData)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "job-create", "service.JobCreateRequestData", v)
		}
		p := rd.Payload
		body = rd.Body
		serviceID = p.ServiceID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: JobCreateServicePath(serviceID)}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "job-create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeJobCreateRequest returns an encoder for requests sent to the service
// job-create server.
func EncodeJobCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		data, ok := v.(*service.JobCreateRequestData)
		if !ok {
			return goahttp.ErrInvalidType("service", "job-create", "*service.JobCreateRequestData", v)
		}
		p := data.Payload
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		{
			head := p.InContentType
			req.Header.Set("Content-Type", head)
		}
		if p.InOrderID != nil {
			head := *p.InOrderID
			req.Header.Set("IVCAP-Order-Id", head)
		}
		if p.ForwardHost != nil {
			head := *p.ForwardHost
			req.Header.Set("X-Forwarded-Host", head)
		}
		if p.ForwardProto != nil {
			head := *p.ForwardProto
			req.Header.Set("X-Forwarded-Proto", head)
		}
		if p.Timeout != nil {
			head := *p.Timeout
			headStr := strconv.Itoa(head)
			req.Header.Set("Timeout", headStr)
		}
		return nil
	}
}

// DecodeJobCreateResponse returns a decoder for responses returned by the
// service job-create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeJobCreateResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-ready-yet" (type *service.JobRetryLaterT): http.StatusAccepted
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "temporary-redirect" (type *service.TemporaryRedirectT): http.StatusTemporaryRedirect
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeJobCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				outContentType string
				outOrderID     string
				jobID          string
				jobURL         *string
				err            error
			)
			outContentTypeRaw := resp.Header.Get("Content-Type")
			if outContentTypeRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("outContentType", "header"))
			}
			outContentType = outContentTypeRaw
			outOrderIDRaw := resp.Header.Get("Ivcap-Order-Id")
			if outOrderIDRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("outOrderID", "header"))
			}
			outOrderID = outOrderIDRaw
			jobIDRaw := resp.Header.Get("Ivcap-Job-Id")
			if jobIDRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("jobID", "header"))
			}
			jobID = jobIDRaw
			jobURLRaw := resp.Header.Get("Ivcap-Job-Url")
			if jobURLRaw != "" {
				jobURL = &jobURLRaw
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			res := NewJobCreateResultOK(outContentType, outOrderID, jobID, jobURL)
			return res, nil
		case http.StatusBadRequest:
			var (
				body JobCreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-create", err)
			}
			err = ValidateJobCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body JobCreateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-create", err)
			}
			err = ValidateJobCreateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body JobCreateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-create", err)
			}
			err = ValidateJobCreateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateInvalidScopes(&body)
		case http.StatusAccepted:
			var (
				body JobCreateNotReadyYetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-create", err)
			}
			err = ValidateJobCreateNotReadyYetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateNotReadyYet(&body)
		case http.StatusNotImplemented:
			var (
				body JobCreateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-create", err)
			}
			err = ValidateJobCreateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body JobCreateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-create", err)
			}
			err = ValidateJobCreateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewJobCreateNotAvailable()
		case http.StatusTemporaryRedirect:
			var (
				location string
				err      error
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("location", "header"))
			}
			location = locationRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-create", err)
			}
			return nil, NewJobCreateTemporaryRedirect(location)
		case http.StatusUnauthorized:
			return nil, NewJobCreateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "job-create", resp.StatusCode, string(body))
		}
	}
}

// // BuildJobCreateStreamPayload creates a streaming endpoint request payload
// from the method payload and the path to the file to be streamed
func BuildJobCreateStreamPayload(payload any, fpath string) (*service.JobCreateRequestData, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return &service.JobCreateRequestData{
		Payload: payload.(*service.JobCreatePayload),
		Body:    f,
	}, nil
}

// BuildJobReadRequest instantiates a HTTP request object with method and path
// set to call the "service" service "job-read" endpoint
func (c *Client) BuildJobReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		serviceID string
		id        string
	)
	{
		p, ok := v.(*service.JobReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "job-read", "*service.JobReadPayload", v)
		}
		serviceID = p.ServiceID
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: JobReadServicePath(serviceID, id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "job-read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeJobReadRequest returns an encoder for requests sent to the service
// job-read server.
func EncodeJobReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.JobReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "job-read", "*service.JobReadPayload", v)
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
		if p.WithRequestContent != nil {
			values.Add("with-request-content", fmt.Sprintf("%v", *p.WithRequestContent))
		}
		values.Add("with-result-content", fmt.Sprintf("%v", p.WithResultContent))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeJobReadResponse returns a decoder for responses returned by the
// service job-read endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeJobReadResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeJobReadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body JobReadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-read", err)
			}
			err = ValidateJobReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-read", err)
			}
			res := NewJobReadJobStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body JobReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-read", err)
			}
			err = ValidateJobReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-read", err)
			}
			return nil, NewJobReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body JobReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-read", err)
			}
			err = ValidateJobReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-read", err)
			}
			return nil, NewJobReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body JobReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-read", err)
			}
			err = ValidateJobReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-read", err)
			}
			return nil, NewJobReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body JobReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-read", err)
			}
			err = ValidateJobReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-read", err)
			}
			return nil, NewJobReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewJobReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewJobReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "job-read", resp.StatusCode, string(body))
		}
	}
}

// BuildJobOutputRequest instantiates a HTTP request object with method and
// path set to call the "service" service "job-output" endpoint
func (c *Client) BuildJobOutputRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		serviceID string
		jobID     string
	)
	{
		p, ok := v.(*service.JobOutputPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "job-output", "*service.JobOutputPayload", v)
		}
		serviceID = p.ServiceID
		jobID = p.JobID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: JobOutputServicePath(serviceID, jobID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "job-output", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeJobOutputRequest returns an encoder for requests sent to the service
// job-output server.
func EncodeJobOutputRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.JobOutputPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "job-output", "*service.JobOutputPayload", v)
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

// DecodeJobOutputResponse returns a decoder for responses returned by the
// service job-output endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeJobOutputResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "job-request-error" (type *service.JobRequestErrorT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "job-internal-error" (type *service.JobInternalErrorT): http.StatusInternalServerError
//   - "job-no-result" (type *service.JobNoResultT): http.StatusNoContent
//   - "not-ready-yet" (type *service.JobRetryLaterT): http.StatusNoContent
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeJobOutputResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				contentType string
				orderID     string
				jobID       string
				jobURL      string
				err         error
			)
			contentTypeRaw := resp.Header.Get("Content-Type")
			if contentTypeRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("contentType", "header"))
			}
			contentType = contentTypeRaw
			orderIDRaw := resp.Header.Get("Ivcap-Order-Id")
			if orderIDRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("orderID", "header"))
			}
			orderID = orderIDRaw
			jobIDRaw := resp.Header.Get("Ivcap-Job-Id")
			if jobIDRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("jobID", "header"))
			}
			jobID = jobIDRaw
			jobURLRaw := resp.Header.Get("Ivcap-Job-Url")
			if jobURLRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("jobURL", "header"))
			}
			jobURL = jobURLRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-output", err)
			}
			res := NewJobOutputResultOK(contentType, orderID, jobID, jobURL)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body JobOutputBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "job-output", err)
				}
				err = ValidateJobOutputBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "job-output", err)
				}
				return nil, NewJobOutputBadRequest(&body)
			case "job-request-error":
				var (
					body JobOutputJobRequestErrorResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "job-output", err)
				}
				err = ValidateJobOutputJobRequestErrorResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "job-output", err)
				}
				return nil, NewJobOutputJobRequestError(&body)
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "job-output", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body JobOutputInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-output", err)
			}
			err = ValidateJobOutputInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-output", err)
			}
			return nil, NewJobOutputInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body JobOutputInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-output", err)
			}
			err = ValidateJobOutputInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-output", err)
			}
			return nil, NewJobOutputInvalidScopes(&body)
		case http.StatusInternalServerError:
			var (
				body JobOutputJobInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-output", err)
			}
			err = ValidateJobOutputJobInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-output", err)
			}
			return nil, NewJobOutputJobInternalError(&body)
		case http.StatusNoContent:
			en := resp.Header.Get("goa-error")
			switch en {
			case "job-no-result":
				return nil, NewJobOutputJobNoResult()
			case "not-ready-yet":
				var (
					body JobOutputNotReadyYetResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("service", "job-output", err)
				}
				err = ValidateJobOutputNotReadyYetResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "job-output", err)
				}
				var (
					location string
				)
				locationRaw := resp.Header.Get("Location")
				if locationRaw == "" {
					err = goa.MergeErrors(err, goa.MissingFieldError("location", "header"))
				}
				location = locationRaw
				if err != nil {
					return nil, goahttp.ErrValidationError("service", "job-output", err)
				}
				return nil, NewJobOutputNotReadyYet(&body, location)
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("service", "job-output", resp.StatusCode, string(body))
			}
		case http.StatusNotImplemented:
			var (
				body JobOutputNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-output", err)
			}
			err = ValidateJobOutputNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-output", err)
			}
			return nil, NewJobOutputNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body JobOutputNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "job-output", err)
			}
			err = ValidateJobOutputNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "job-output", err)
			}
			return nil, NewJobOutputNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewJobOutputNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewJobOutputNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "job-output", resp.StatusCode, string(body))
		}
	}
}

// unmarshalServiceListItemTResponseBodyToServiceviewsServiceListItemTView
// builds a value of type *serviceviews.ServiceListItemTView from a value of
// type *ServiceListItemTResponseBody.
func unmarshalServiceListItemTResponseBodyToServiceviewsServiceListItemTView(v *ServiceListItemTResponseBody) *serviceviews.ServiceListItemTView {
	res := &serviceviews.ServiceListItemTView{
		ID:               v.ID,
		Name:             v.Name,
		Description:      v.Description,
		ControllerSchema: v.ControllerSchema,
		ValidFrom:        v.ValidFrom,
		ValidTo:          v.ValidTo,
		Href:             v.Href,
	}
	if v.Tags != nil {
		res.Tags = make([]string, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = val
		}
	}

	return res
}

// unmarshalLinkTResponseBodyToServiceviewsLinkTView builds a value of type
// *serviceviews.LinkTView from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToServiceviewsLinkTView(v *LinkTResponseBody) *serviceviews.LinkTView {
	res := &serviceviews.LinkTView{
		Rel:  v.Rel,
		Type: v.Type,
		Href: v.Href,
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
		Unary:       v.Unary,
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
		Unary:       v.Unary,
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

// unmarshalLinkTResponseBodyToServiceLinkT builds a value of type
// *service.LinkT from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToServiceLinkT(v *LinkTResponseBody) *service.LinkT {
	res := &service.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}

// unmarshalParameterDefTResponseBodyToServiceParameterDefT builds a value of
// type *service.ParameterDefT from a value of type *ParameterDefTResponseBody.
func unmarshalParameterDefTResponseBodyToServiceParameterDefT(v *ParameterDefTResponseBody) *service.ParameterDefT {
	res := &service.ParameterDefT{
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
		res.Options = make([]*service.ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = unmarshalParameterOptTResponseBodyToServiceParameterOptT(val)
		}
	}

	return res
}

// unmarshalParameterOptTResponseBodyToServiceParameterOptT builds a value of
// type *service.ParameterOptT from a value of type *ParameterOptTResponseBody.
func unmarshalParameterOptTResponseBodyToServiceParameterOptT(v *ParameterOptTResponseBody) *service.ParameterOptT {
	if v == nil {
		return nil
	}
	res := &service.ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// unmarshalJobListItemResponseBodyToServiceviewsJobListItemView builds a value
// of type *serviceviews.JobListItemView from a value of type
// *JobListItemResponseBody.
func unmarshalJobListItemResponseBodyToServiceviewsJobListItemView(v *JobListItemResponseBody) *serviceviews.JobListItemView {
	res := &serviceviews.JobListItemView{
		ID:         v.ID,
		Name:       v.Name,
		Status:     v.Status,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		Service:    v.Service,
		Order:      v.Order,
		Href:       v.Href,
	}

	return res
}

// unmarshalPartialProductList2TResponseBodyToServicePartialProductList2T
// builds a value of type *service.PartialProductList2T from a value of type
// *PartialProductList2TResponseBody.
func unmarshalPartialProductList2TResponseBodyToServicePartialProductList2T(v *PartialProductList2TResponseBody) *service.PartialProductList2T {
	if v == nil {
		return nil
	}
	res := &service.PartialProductList2T{}
	res.Items = make([]*service.ProductListItem2T, len(v.Items))
	for i, val := range v.Items {
		res.Items[i] = unmarshalProductListItem2TResponseBodyToServiceProductListItem2T(val)
	}
	res.Links = make([]*service.LinkT, len(v.Links))
	for i, val := range v.Links {
		res.Links[i] = unmarshalLinkTResponseBodyToServiceLinkT(val)
	}

	return res
}

// unmarshalProductListItem2TResponseBodyToServiceProductListItem2T builds a
// value of type *service.ProductListItem2T from a value of type
// *ProductListItem2TResponseBody.
func unmarshalProductListItem2TResponseBodyToServiceProductListItem2T(v *ProductListItem2TResponseBody) *service.ProductListItem2T {
	res := &service.ProductListItem2T{
		ID:       *v.ID,
		Name:     v.Name,
		Status:   *v.Status,
		MimeType: v.MimeType,
		Size:     v.Size,
		Href:     *v.Href,
		DataHref: v.DataHref,
	}

	return res
}
