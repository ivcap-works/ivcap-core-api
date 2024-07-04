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
	"os"
	"strconv"
	"strings"

	package_ "github.com/ivcap-works/ivcap-core-api/gen/package_"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
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

// BuildPullRequest instantiates a HTTP request object with method and path set
// to call the "package" service "pull" endpoint
func (c *Client) BuildPullRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PullPackagePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("package", "pull", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePullRequest returns an encoder for requests sent to the package pull
// server.
func EncodePullRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*package_.PullPayload)
		if !ok {
			return goahttp.ErrInvalidType("package", "pull", "*package_.PullPayload", v)
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
		values.Add("ref", p.Ref)
		values.Add("type", p.Type)
		if p.Offset != nil {
			values.Add("offset", fmt.Sprintf("%v", *p.Offset))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodePullResponse returns a decoder for responses returned by the package
// pull endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodePullResponse may return the following errors:
//   - "bad-request" (type *package_.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *package_.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *package_.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *package_.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *package_.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *package_.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodePullResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				total     int
				available int
				err       error
			)
			{
				totalRaw := resp.Header.Get("Total")
				if totalRaw == "" {
					return nil, goahttp.ErrValidationError("package", "pull", goa.MissingFieldError("total", "header"))
				}
				v, err2 := strconv.ParseInt(totalRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("total", totalRaw, "integer"))
				}
				total = int(v)
			}
			{
				availableRaw := resp.Header.Get("Available")
				if availableRaw == "" {
					return nil, goahttp.ErrValidationError("package", "pull", goa.MissingFieldError("available", "header"))
				}
				v, err2 := strconv.ParseInt(availableRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("available", availableRaw, "integer"))
				}
				available = int(v)
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "pull", err)
			}
			res := NewPullResultTOK(total, available)
			return res, nil
		case http.StatusBadRequest:
			var (
				body PullBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "pull", err)
			}
			err = ValidatePullBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "pull", err)
			}
			return nil, NewPullBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body PullInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "pull", err)
			}
			err = ValidatePullInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "pull", err)
			}
			return nil, NewPullInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body PullInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "pull", err)
			}
			err = ValidatePullInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "pull", err)
			}
			return nil, NewPullInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body PullNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "pull", err)
			}
			err = ValidatePullNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "pull", err)
			}
			return nil, NewPullNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewPullNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewPullNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("package", "pull", resp.StatusCode, string(body))
		}
	}
}

// BuildPushRequest instantiates a HTTP request object with method and path set
// to call the "package" service "push" endpoint
func (c *Client) BuildPushRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		body io.Reader
	)
	rd, ok := v.(*package_.PushRequestData)
	if !ok {
		return nil, goahttp.ErrInvalidType("package", "push", "package_.PushRequestData", v)
	}
	body = rd.Body
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PushPackagePath()}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("package", "push", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePushRequest returns an encoder for requests sent to the package push
// server.
func EncodePushRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		data, ok := v.(*package_.PushRequestData)
		if !ok {
			return goahttp.ErrInvalidType("package", "push", "*package_.PushRequestData", v)
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
		values := req.URL.Query()
		values.Add("tag", p.Tag)
		if p.Force != nil {
			values.Add("force", fmt.Sprintf("%v", *p.Force))
		}
		values.Add("type", p.Type)
		values.Add("digest", p.Digest)
		if p.Total != nil {
			values.Add("total", fmt.Sprintf("%v", *p.Total))
		}
		if p.Start != nil {
			values.Add("start", fmt.Sprintf("%v", *p.Start))
		}
		if p.End != nil {
			values.Add("end", fmt.Sprintf("%v", *p.End))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodePushResponse returns a decoder for responses returned by the package
// push endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodePushResponse may return the following errors:
//   - "bad-request" (type *package_.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *package_.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *package_.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *package_.NotImplementedT): http.StatusNotImplemented
//   - "already-created" (type *package_.ResourceAlreadyCreatedT): http.StatusConflict
//   - "not-available" (type *package_.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *package_.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodePushResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body PushResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "push", err)
			}
			res := NewPushResultCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body PushBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "push", err)
			}
			err = ValidatePushBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "push", err)
			}
			return nil, NewPushBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body PushInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "push", err)
			}
			err = ValidatePushInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "push", err)
			}
			return nil, NewPushInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body PushInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "push", err)
			}
			err = ValidatePushInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "push", err)
			}
			return nil, NewPushInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body PushNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "push", err)
			}
			err = ValidatePushNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "push", err)
			}
			return nil, NewPushNotImplemented(&body)
		case http.StatusConflict:
			var (
				body PushAlreadyCreatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "push", err)
			}
			err = ValidatePushAlreadyCreatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "push", err)
			}
			return nil, NewPushAlreadyCreated(&body)
		case http.StatusServiceUnavailable:
			return nil, NewPushNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewPushNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("package", "push", resp.StatusCode, string(body))
		}
	}
}

// // BuildPushStreamPayload creates a streaming endpoint request payload from the
// method payload and the path to the file to be streamed
func BuildPushStreamPayload(payload any, fpath string) (*package_.PushRequestData, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return &package_.PushRequestData{
		Payload: payload.(*package_.PushPayload),
		Body:    f,
	}, nil
}

// BuildStatusRequest instantiates a HTTP request object with method and path
// set to call the "package" service "status" endpoint
func (c *Client) BuildStatusRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StatusPackagePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("package", "status", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStatusRequest returns an encoder for requests sent to the package
// status server.
func EncodeStatusRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*package_.StatusPayload)
		if !ok {
			return goahttp.ErrInvalidType("package", "status", "*package_.StatusPayload", v)
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
		values.Add("digest", p.Digest)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeStatusResponse returns a decoder for responses returned by the package
// status endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeStatusResponse may return the following errors:
//   - "bad-request" (type *package_.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *package_.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *package_.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *package_.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *package_.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *package_.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeStatusResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body StatusResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "status", err)
			}
			err = ValidateStatusResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "status", err)
			}
			res := NewStatusPushStatusTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body StatusBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "status", err)
			}
			err = ValidateStatusBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "status", err)
			}
			return nil, NewStatusBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body StatusInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "status", err)
			}
			err = ValidateStatusInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "status", err)
			}
			return nil, NewStatusInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body StatusInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "status", err)
			}
			err = ValidateStatusInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "status", err)
			}
			return nil, NewStatusInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body StatusNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("package", "status", err)
			}
			err = ValidateStatusNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("package", "status", err)
			}
			return nil, NewStatusNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewStatusNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewStatusNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("package", "status", resp.StatusCode, string(body))
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
