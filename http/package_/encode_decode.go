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

	package_ "github.com/ivcap-works/ivcap-core-api/gen/package_"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "package" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListPackagePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("package", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the package list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*package_.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("package", "list", "*package_.ListPayload", v)
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
		if p.Tag != nil {
			values.Add("tag", *p.Tag)
		}
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		if p.Limit != nil {
			values.Add("limit", fmt.Sprintf("%v", *p.Limit))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the package
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *package_.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *package_.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *package_.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *package_.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *package_.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *package_.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("package", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "list", err)
			}
			res := NewListResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("package", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "package" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemovePackagePath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("package", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRemoveRequest returns an encoder for requests sent to the package
// remove server.
func EncodeRemoveRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*package_.RemovePayload)
		if !ok {
			return goahttp.ErrInvalidType("package", "remove", "*package_.RemovePayload", v)
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
		values.Add("tag", p.Tag)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeRemoveResponse returns a decoder for responses returned by the package
// remove endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeRemoveResponse may return the following errors:
//   - "bad-request" (type *package_.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *package_.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *package_.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *package_.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *package_.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *package_.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeRemoveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body RemoveBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "remove", err)
			}
			err = ValidateRemoveBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "remove", err)
			}
			return nil, NewRemoveBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body RemoveInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "remove", err)
			}
			err = ValidateRemoveInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "remove", err)
			}
			return nil, NewRemoveInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body RemoveInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "remove", err)
			}
			err = ValidateRemoveInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "remove", err)
			}
			return nil, NewRemoveInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body RemoveNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "remove", err)
			}
			err = ValidateRemoveNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "remove", err)
			}
			return nil, NewRemoveNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewRemoveNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewRemoveNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("package", "remove", resp.StatusCode, string(body))
		}
	}
}

// unmarshalLinkTResponseBodyToPackageLinkT builds a value of type
// *package_.LinkT from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToPackageLinkT(v *LinkTResponseBody) *package_.LinkT {
	res := &package_.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}
