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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	secret "github.com/ivcap-works/ivcap-core-api/gen/secret"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "secret" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListSecretPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secret", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the secret list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*secret.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("secret", "list", "*secret.ListPayload", v)
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
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		if p.Filter != nil {
			values.Add("filter", *p.Filter)
		}
		if p.Offset != nil {
			values.Add("offset", *p.Offset)
		}
		if p.Limit != nil {
			values.Add("limit", fmt.Sprintf("%v", *p.Limit))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the secret
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *secret.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *secret.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *secret.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *secret.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *secret.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *secret.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("secret", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "list", err)
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
				return nil, goahttp.ErrDecodingError("secret", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secret", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "secret" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetSecretPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secret", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the secret get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*secret.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("secret", "get", "*secret.GetPayload", v)
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
		values.Add("secret-name", p.SecretName)
		if p.SecretType != nil {
			values.Add("secret-type", *p.SecretType)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the secret get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeGetResponse may return the following errors:
//   - "bad-request" (type *secret.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *secret.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *secret.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *secret.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *secret.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *secret.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *secret.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "get", err)
			}
			res := NewGetSecretResultTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "get", err)
			}
			return nil, NewGetBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body GetInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "get", err)
			}
			err = ValidateGetInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "get", err)
			}
			return nil, NewGetInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body GetInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "get", err)
			}
			err = ValidateGetInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "get", err)
			}
			return nil, NewGetInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body GetNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "get", err)
			}
			err = ValidateGetNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "get", err)
			}
			return nil, NewGetNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "get", err)
			}
			err = ValidateGetNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "get", err)
			}
			return nil, NewGetNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewGetNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewGetNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secret", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildSetRequest instantiates a HTTP request object with method and path set
// to call the "secret" service "set" endpoint
func (c *Client) BuildSetRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetSecretPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secret", "set", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetRequest returns an encoder for requests sent to the secret set
// server.
func EncodeSetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*secret.SetPayload)
		if !ok {
			return goahttp.ErrInvalidType("secret", "set", "*secret.SetPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSetRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("secret", "set", err)
		}
		return nil
	}
}

// DecodeSetResponse returns a decoder for responses returned by the secret set
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeSetResponse may return the following errors:
//   - "bad-request" (type *secret.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *secret.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *secret.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *secret.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *secret.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *secret.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *secret.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeSetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body SetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "set", err)
			}
			err = ValidateSetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "set", err)
			}
			return nil, NewSetBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body SetInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "set", err)
			}
			err = ValidateSetInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "set", err)
			}
			return nil, NewSetInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body SetInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "set", err)
			}
			err = ValidateSetInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "set", err)
			}
			return nil, NewSetInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body SetNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "set", err)
			}
			err = ValidateSetNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "set", err)
			}
			return nil, NewSetNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body SetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secret", "set", err)
			}
			err = ValidateSetNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secret", "set", err)
			}
			return nil, NewSetNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewSetNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewSetNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secret", "set", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSecretListItemResponseBodyToSecretSecretListItem builds a value of
// type *secret.SecretListItem from a value of type *SecretListItemResponseBody.
func unmarshalSecretListItemResponseBodyToSecretSecretListItem(v *SecretListItemResponseBody) *secret.SecretListItem {
	res := &secret.SecretListItem{
		SecretName: *v.SecretName,
		ExpiryTime: *v.ExpiryTime,
	}

	return res
}

// unmarshalLinkTResponseBodyToSecretLinkT builds a value of type *secret.LinkT
// from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToSecretLinkT(v *LinkTResponseBody) *secret.LinkT {
	res := &secret.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}
