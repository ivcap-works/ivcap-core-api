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

	order "github.com/ivcap-works/ivcap-core-api/gen/order"
	orderviews "github.com/ivcap-works/ivcap-core-api/gen/order/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "order" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListOrderPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the order list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "list", "*order.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the order
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *order.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("order", "list", err)
			}
			p := NewListOrderListRTOK(&body)
			view := "default"
			vres := &orderviews.OrderListRT{Projected: p, View: view}
			if err = orderviews.ValidateOrderListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("order", "list", err)
			}
			res := order.NewOrderListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "order" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*order.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("order", "read", "*order.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadOrderPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the order read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "read", "*order.ReadPayload", v)
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

// DecodeReadResponse returns a decoder for responses returned by the order
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *order.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("order", "read", err)
			}
			err = ValidateReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "read", err)
			}
			res := NewReadOrderStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "read", err)
			}
			err = ValidateReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "read", err)
			}
			return nil, NewReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildProductsRequest instantiates a HTTP request object with method and path
// set to call the "order" service "products" endpoint
func (c *Client) BuildProductsRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		orderID string
	)
	{
		p, ok := v.(*order.ProductsPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("order", "products", "*order.ProductsPayload", v)
		}
		orderID = p.OrderID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ProductsOrderPath(orderID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "products", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeProductsRequest returns an encoder for requests sent to the order
// products server.
func EncodeProductsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.ProductsPayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "products", "*order.ProductsPayload", v)
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
		if p.OrderBy != nil {
			values.Add("order-by", *p.OrderBy)
		}
		values.Add("order-desc", fmt.Sprintf("%v", p.OrderDesc))
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeProductsResponse returns a decoder for responses returned by the order
// products endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeProductsResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *order.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *order.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeProductsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ProductsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "products", err)
			}
			err = ValidateProductsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "products", err)
			}
			res := NewProductsPartialProductListTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ProductsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "products", err)
			}
			err = ValidateProductsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "products", err)
			}
			return nil, NewProductsBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ProductsInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "products", err)
			}
			err = ValidateProductsInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "products", err)
			}
			return nil, NewProductsInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ProductsInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "products", err)
			}
			err = ValidateProductsInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "products", err)
			}
			return nil, NewProductsInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ProductsNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "products", err)
			}
			err = ValidateProductsNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "products", err)
			}
			return nil, NewProductsNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ProductsNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "products", err)
			}
			err = ValidateProductsNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "products", err)
			}
			return nil, NewProductsNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewProductsNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewProductsNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "products", resp.StatusCode, string(body))
		}
	}
}

// BuildMetadataRequest instantiates a HTTP request object with method and path
// set to call the "order" service "metadata" endpoint
func (c *Client) BuildMetadataRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		orderID string
	)
	{
		p, ok := v.(*order.MetadataPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("order", "metadata", "*order.MetadataPayload", v)
		}
		orderID = p.OrderID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MetadataOrderPath(orderID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "metadata", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMetadataRequest returns an encoder for requests sent to the order
// metadata server.
func EncodeMetadataRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.MetadataPayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "metadata", "*order.MetadataPayload", v)
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
		if p.OrderBy != nil {
			values.Add("order-by", *p.OrderBy)
		}
		values.Add("order-desc", fmt.Sprintf("%v", p.OrderDesc))
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeMetadataResponse returns a decoder for responses returned by the order
// metadata endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeMetadataResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *order.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *order.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeMetadataResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body MetadataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "metadata", err)
			}
			err = ValidateMetadataResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "metadata", err)
			}
			res := NewMetadataPartialMetaListTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body MetadataBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "metadata", err)
			}
			err = ValidateMetadataBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "metadata", err)
			}
			return nil, NewMetadataBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body MetadataInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "metadata", err)
			}
			err = ValidateMetadataInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "metadata", err)
			}
			return nil, NewMetadataInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body MetadataInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "metadata", err)
			}
			err = ValidateMetadataInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "metadata", err)
			}
			return nil, NewMetadataInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body MetadataNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "metadata", err)
			}
			err = ValidateMetadataNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "metadata", err)
			}
			return nil, NewMetadataNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body MetadataNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "metadata", err)
			}
			err = ValidateMetadataNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "metadata", err)
			}
			return nil, NewMetadataNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewMetadataNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewMetadataNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "metadata", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "order" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateOrderPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the order create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "create", "*order.CreatePayload", v)
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
			return goahttp.ErrEncodingError("order", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the order
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *order.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *order.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create", err)
			}
			res := NewCreateOrderStatusRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create", err)
			}
			return nil, NewCreateBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body CreateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create", err)
			}
			err = ValidateCreateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create", err)
			}
			return nil, NewCreateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create", err)
			}
			err = ValidateCreateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create", err)
			}
			return nil, NewCreateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create", err)
			}
			err = ValidateCreateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create", err)
			}
			return nil, NewCreateNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body CreateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "create", err)
			}
			err = ValidateCreateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "create", err)
			}
			return nil, NewCreateNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewCreateNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewCreateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildLogsRequest instantiates a HTTP request object with method and path set
// to call the "order" service "logs" endpoint
func (c *Client) BuildLogsRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		orderID string
	)
	{
		p, ok := v.(*order.LogsPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("order", "logs", "*order.LogsPayload", v)
		}
		orderID = p.OrderID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LogsOrderPath(orderID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "logs", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLogsRequest returns an encoder for requests sent to the order logs
// server.
func EncodeLogsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.LogsPayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "logs", "*order.LogsPayload", v)
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
		if p.From != nil {
			values.Add("from", fmt.Sprintf("%v", *p.From))
		}
		if p.To != nil {
			values.Add("to", fmt.Sprintf("%v", *p.To))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeLogsResponse returns a decoder for responses returned by the order
// logs endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeLogsResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *order.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *order.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeLogsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
			return nil, nil
		case http.StatusBadRequest:
			var (
				body LogsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "logs", err)
			}
			err = ValidateLogsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "logs", err)
			}
			return nil, NewLogsBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body LogsInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "logs", err)
			}
			err = ValidateLogsInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "logs", err)
			}
			return nil, NewLogsInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body LogsInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "logs", err)
			}
			err = ValidateLogsInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "logs", err)
			}
			return nil, NewLogsInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body LogsNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "logs", err)
			}
			err = ValidateLogsNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "logs", err)
			}
			return nil, NewLogsNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body LogsNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "logs", err)
			}
			err = ValidateLogsNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "logs", err)
			}
			return nil, NewLogsNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewLogsNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewLogsNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "logs", resp.StatusCode, string(body))
		}
	}
}

// BuildTopRequest instantiates a HTTP request object with method and path set
// to call the "order" service "top" endpoint
func (c *Client) BuildTopRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		orderID string
	)
	{
		p, ok := v.(*order.TopPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("order", "top", "*order.TopPayload", v)
		}
		orderID = p.OrderID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: TopOrderPath(orderID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("order", "top", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeTopRequest returns an encoder for requests sent to the order top
// server.
func EncodeTopRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*order.TopPayload)
		if !ok {
			return goahttp.ErrInvalidType("order", "top", "*order.TopPayload", v)
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

// DecodeTopResponse returns a decoder for responses returned by the order top
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeTopResponse may return the following errors:
//   - "bad-request" (type *order.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *order.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *order.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *order.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *order.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *order.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *order.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeTopResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body TopResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "top", err)
			}
			p := NewTopOrderTopResultItemCollectionOK(body)
			view := "default"
			vres := orderviews.OrderTopResultItemCollection{Projected: p, View: view}
			if err = orderviews.ValidateOrderTopResultItemCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("order", "top", err)
			}
			res := order.NewOrderTopResultItemCollection(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body TopBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "top", err)
			}
			err = ValidateTopBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "top", err)
			}
			return nil, NewTopBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body TopInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "top", err)
			}
			err = ValidateTopInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "top", err)
			}
			return nil, NewTopInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body TopInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "top", err)
			}
			err = ValidateTopInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "top", err)
			}
			return nil, NewTopInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body TopNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "top", err)
			}
			err = ValidateTopNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "top", err)
			}
			return nil, NewTopNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body TopNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("order", "top", err)
			}
			err = ValidateTopNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("order", "top", err)
			}
			return nil, NewTopNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewTopNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewTopNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("order", "top", resp.StatusCode, string(body))
		}
	}
}

// unmarshalOrderListItemResponseBodyToOrderviewsOrderListItemView builds a
// value of type *orderviews.OrderListItemView from a value of type
// *OrderListItemResponseBody.
func unmarshalOrderListItemResponseBodyToOrderviewsOrderListItemView(v *OrderListItemResponseBody) *orderviews.OrderListItemView {
	res := &orderviews.OrderListItemView{
		ID:         v.ID,
		Name:       v.Name,
		Status:     v.Status,
		OrderedAt:  v.OrderedAt,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		Service:    v.Service,
		Account:    v.Account,
		Href:       v.Href,
	}

	return res
}

// unmarshalLinkTResponseBodyToOrderviewsLinkTView builds a value of type
// *orderviews.LinkTView from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToOrderviewsLinkTView(v *LinkTResponseBody) *orderviews.LinkTView {
	res := &orderviews.LinkTView{
		Rel:  v.Rel,
		Type: v.Type,
		Href: v.Href,
	}

	return res
}

// unmarshalPartialProductListTResponseBodyToOrderPartialProductListT builds a
// value of type *order.PartialProductListT from a value of type
// *PartialProductListTResponseBody.
func unmarshalPartialProductListTResponseBodyToOrderPartialProductListT(v *PartialProductListTResponseBody) *order.PartialProductListT {
	res := &order.PartialProductListT{}
	res.Items = make([]*order.ProductListItemT, len(v.Items))
	for i, val := range v.Items {
		res.Items[i] = unmarshalProductListItemTResponseBodyToOrderProductListItemT(val)
	}
	res.Links = make([]*order.LinkT, len(v.Links))
	for i, val := range v.Links {
		res.Links[i] = unmarshalLinkTResponseBodyToOrderLinkT(val)
	}

	return res
}

// unmarshalProductListItemTResponseBodyToOrderProductListItemT builds a value
// of type *order.ProductListItemT from a value of type
// *ProductListItemTResponseBody.
func unmarshalProductListItemTResponseBodyToOrderProductListItemT(v *ProductListItemTResponseBody) *order.ProductListItemT {
	res := &order.ProductListItemT{
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

// unmarshalLinkTResponseBodyToOrderLinkT builds a value of type *order.LinkT
// from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToOrderLinkT(v *LinkTResponseBody) *order.LinkT {
	res := &order.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}

// unmarshalParameterTResponseBodyToOrderParameterT builds a value of type
// *order.ParameterT from a value of type *ParameterTResponseBody.
func unmarshalParameterTResponseBodyToOrderParameterT(v *ParameterTResponseBody) *order.ParameterT {
	res := &order.ParameterT{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// unmarshalOrderMetadataListItemRTResponseBodyToOrderOrderMetadataListItemRT
// builds a value of type *order.OrderMetadataListItemRT from a value of type
// *OrderMetadataListItemRTResponseBody.
func unmarshalOrderMetadataListItemRTResponseBodyToOrderOrderMetadataListItemRT(v *OrderMetadataListItemRTResponseBody) *order.OrderMetadataListItemRT {
	res := &order.OrderMetadataListItemRT{
		ID:          *v.ID,
		Schema:      *v.Schema,
		Href:        *v.Href,
		ContentType: *v.ContentType,
	}

	return res
}

// marshalOrderParameterTToParameterT builds a value of type *ParameterT from a
// value of type *order.ParameterT.
func marshalOrderParameterTToParameterT(v *order.ParameterT) *ParameterT {
	res := &ParameterT{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// marshalParameterTToOrderParameterT builds a value of type *order.ParameterT
// from a value of type *ParameterT.
func marshalParameterTToOrderParameterT(v *ParameterT) *order.ParameterT {
	res := &order.ParameterT{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// unmarshalOrderTopResultItemResponseToOrderviewsOrderTopResultItemView builds
// a value of type *orderviews.OrderTopResultItemView from a value of type
// *OrderTopResultItemResponse.
func unmarshalOrderTopResultItemResponseToOrderviewsOrderTopResultItemView(v *OrderTopResultItemResponse) *orderviews.OrderTopResultItemView {
	res := &orderviews.OrderTopResultItemView{
		Container:        v.Container,
		CPU:              v.CPU,
		Memory:           v.Memory,
		Storage:          v.Storage,
		EphemeralStorage: v.EphemeralStorage,
	}

	return res
}
