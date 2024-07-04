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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	service "github.com/ivcap-works/ivcap-core-api/gen/service"
	serviceviews "github.com/ivcap-works/ivcap-core-api/gen/service/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "service" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
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
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "list", "*service.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the service
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
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
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
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
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateServiceRequest instantiates a HTTP request object with method and
// path set to call the "service" service "create_service" endpoint
func (c *Client) BuildCreateServiceRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateServiceServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "create_service", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateServiceRequest returns an encoder for requests sent to the
// service create_service server.
func EncodeCreateServiceRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.CreateServicePayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "create_service", "*service.CreateServicePayload", v)
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
			return goahttp.ErrEncodingError("service", "create_service", err)
		}
		return nil
	}
}

// DecodeCreateServiceResponse returns a decoder for responses returned by the
// service create_service endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateServiceResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "already-created" (type *service.ResourceAlreadyCreatedT): http.StatusConflict
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			res := NewCreateServiceServiceStatusRTCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateServiceBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			return nil, NewCreateServiceBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body CreateServiceInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			return nil, NewCreateServiceInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateServiceInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			return nil, NewCreateServiceInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateServiceNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			return nil, NewCreateServiceNotImplemented(&body)
		case http.StatusConflict:
			var (
				body CreateServiceAlreadyCreatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceAlreadyCreatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			return nil, NewCreateServiceAlreadyCreated(&body)
		case http.StatusNotFound:
			var (
				body CreateServiceNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "create_service", err)
			}
			err = ValidateCreateServiceNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "create_service", err)
			}
			return nil, NewCreateServiceNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewCreateServiceNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewCreateServiceNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "create_service", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "service" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
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

// EncodeReadRequest returns an encoder for requests sent to the service read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "read", "*service.ReadPayload", v)
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

// DecodeReadResponse returns a decoder for responses returned by the service
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *service.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("service", "read", err)
			}
			err = ValidateReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "read", err)
			}
			res := NewReadServiceStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
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
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("service", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
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
		case http.StatusServiceUnavailable:
			return nil, NewReadNotAvailable()
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
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
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
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
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
//   - "invalid-parameter" (type *service.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *service.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("service", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "update", err)
			}
			res := NewUpdateServiceStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
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
		case http.StatusServiceUnavailable:
			return nil, NewUpdateNotAvailable()
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
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
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
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
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
//   - "invalid-scopes" (type *service.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *service.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *service.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *service.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("service", "delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
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
		case http.StatusServiceUnavailable:
			return nil, NewDeleteNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewDeleteNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "delete", resp.StatusCode, string(body))
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
		Banner:      v.Banner,
		PublishedAt: v.PublishedAt,
		Policy:      v.Policy,
		Account:     v.Account,
		Href:        v.Href,
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
		res.Memory = marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody(v.Memory)
	}
	if v.CPU != nil {
		res.CPU = marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody(v.CPU)
	}
	if v.EphemeralStorage != nil {
		res.EphemeralStorage = marshalServiceResourceMemoryTToResourceMemoryTRequestBodyRequestBody(v.EphemeralStorage)
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
		res.Memory = marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT(v.Memory)
	}
	if v.CPU != nil {
		res.CPU = marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT(v.CPU)
	}
	if v.EphemeralStorage != nil {
		res.EphemeralStorage = marshalResourceMemoryTRequestBodyRequestBodyToServiceResourceMemoryT(v.EphemeralStorage)
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
