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

	queue "github.com/ivcap-works/ivcap-core-api/gen/queue"
	queueviews "github.com/ivcap-works/ivcap-core-api/gen/queue/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "queue" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateQueuePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("queue", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the queue create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*queue.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("queue", "create", "*queue.CreatePayload", v)
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
			return goahttp.ErrEncodingError("queue", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the queue
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "bad-request" (type *queue.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *queue.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *queue.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *queue.NotImplementedT): http.StatusNotImplemented
//   - "already-created" (type *queue.ResourceAlreadyCreatedT): http.StatusConflict
//   - "not-found" (type *queue.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *queue.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *queue.UnauthorizedT): http.StatusUnauthorized
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
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			p := NewCreatequeueresponseViewCreated(&body)
			view := "default"
			vres := &queueviews.Createqueueresponse{Projected: p, View: view}
			if err = queueviews.ValidateCreatequeueresponse(vres); err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			res := queue.NewCreatequeueresponse(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			return nil, NewCreateBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body CreateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			err = ValidateCreateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			return nil, NewCreateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			err = ValidateCreateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			return nil, NewCreateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			err = ValidateCreateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			return nil, NewCreateNotImplemented(&body)
		case http.StatusConflict:
			var (
				body CreateAlreadyCreatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			err = ValidateCreateAlreadyCreatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			return nil, NewCreateAlreadyCreated(&body)
		case http.StatusNotFound:
			var (
				body CreateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "create", err)
			}
			err = ValidateCreateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "create", err)
			}
			return nil, NewCreateNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewCreateNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewCreateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("queue", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "queue" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*queue.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("queue", "read", "*queue.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadQueuePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("queue", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the queue read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*queue.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("queue", "read", "*queue.ReadPayload", v)
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

// DecodeReadResponse returns a decoder for responses returned by the queue
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *queue.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *queue.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *queue.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *queue.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *queue.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *queue.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("queue", "read", err)
			}
			p := NewReadqueueresponseViewOK(&body)
			view := "default"
			vres := &queueviews.Readqueueresponse{Projected: p, View: view}
			if err = queueviews.ValidateReadqueueresponse(vres); err != nil {
				return nil, goahttp.ErrValidationError("queue", "read", err)
			}
			res := queue.NewReadqueueresponse(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "read", err)
			}
			err = ValidateReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "read", err)
			}
			return nil, NewReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("queue", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "queue" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*queue.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("queue", "delete", "*queue.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteQueuePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("queue", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the queue delete
// server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*queue.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("queue", "delete", "*queue.DeletePayload", v)
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

// DecodeDeleteResponse returns a decoder for responses returned by the queue
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "bad-request" (type *queue.BadRequestT): http.StatusBadRequest
//   - "invalid-scopes" (type *queue.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *queue.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *queue.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *queue.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("queue", "delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		case http.StatusForbidden:
			var (
				body DeleteInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "delete", err)
			}
			err = ValidateDeleteInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "delete", err)
			}
			return nil, NewDeleteInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body DeleteNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "delete", err)
			}
			err = ValidateDeleteNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "delete", err)
			}
			return nil, NewDeleteNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewDeleteNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewDeleteNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("queue", "delete", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "queue" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListQueuePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("queue", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the queue list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*queue.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("queue", "list", "*queue.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the queue
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *queue.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *queue.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *queue.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *queue.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *queue.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *queue.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("queue", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "list", err)
			}
			res := NewListQueueListResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("queue", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildEnqueueRequest instantiates a HTTP request object with method and path
// set to call the "queue" service "enqueue" endpoint
func (c *Client) BuildEnqueueRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*queue.EnqueuePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("queue", "enqueue", "*queue.EnqueuePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: EnqueueQueuePath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("queue", "enqueue", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeEnqueueRequest returns an encoder for requests sent to the queue
// enqueue server.
func EncodeEnqueueRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*queue.EnqueuePayload)
		if !ok {
			return goahttp.ErrInvalidType("queue", "enqueue", "*queue.EnqueuePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.ContentType != nil {
			head := *p.ContentType
			req.Header.Set("Content-Type", head)
		}
		values := req.URL.Query()
		if p.Schema != nil {
			values.Add("schema", *p.Schema)
		}
		req.URL.RawQuery = values.Encode()
		body := p.Content
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("queue", "enqueue", err)
		}
		return nil
	}
}

// DecodeEnqueueResponse returns a decoder for responses returned by the queue
// enqueue endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeEnqueueResponse may return the following errors:
//   - "bad-request" (type *queue.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *queue.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *queue.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *queue.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *queue.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *queue.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeEnqueueResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body EnqueueResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "enqueue", err)
			}
			p := NewEnqueueMessagestatusOK(&body)
			view := "default"
			vres := &queueviews.Messagestatus{Projected: p, View: view}
			if err = queueviews.ValidateMessagestatus(vres); err != nil {
				return nil, goahttp.ErrValidationError("queue", "enqueue", err)
			}
			res := queue.NewMessagestatus(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body EnqueueBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "enqueue", err)
			}
			err = ValidateEnqueueBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "enqueue", err)
			}
			return nil, NewEnqueueBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body EnqueueInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "enqueue", err)
			}
			err = ValidateEnqueueInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "enqueue", err)
			}
			return nil, NewEnqueueInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body EnqueueInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "enqueue", err)
			}
			err = ValidateEnqueueInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "enqueue", err)
			}
			return nil, NewEnqueueInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body EnqueueNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "enqueue", err)
			}
			err = ValidateEnqueueNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "enqueue", err)
			}
			return nil, NewEnqueueNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewEnqueueNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewEnqueueNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("queue", "enqueue", resp.StatusCode, string(body))
		}
	}
}

// BuildDequeueRequest instantiates a HTTP request object with method and path
// set to call the "queue" service "dequeue" endpoint
func (c *Client) BuildDequeueRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*queue.DequeuePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("queue", "dequeue", "*queue.DequeuePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DequeueQueuePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("queue", "dequeue", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDequeueRequest returns an encoder for requests sent to the queue
// dequeue server.
func EncodeDequeueRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*queue.DequeuePayload)
		if !ok {
			return goahttp.ErrInvalidType("queue", "dequeue", "*queue.DequeuePayload", v)
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
		if p.Limit != nil {
			values.Add("limit", fmt.Sprintf("%v", *p.Limit))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeDequeueResponse returns a decoder for responses returned by the queue
// dequeue endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDequeueResponse may return the following errors:
//   - "bad-request" (type *queue.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *queue.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *queue.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *queue.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *queue.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *queue.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeDequeueResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body DequeueResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "dequeue", err)
			}
			err = ValidateDequeueResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "dequeue", err)
			}
			res := NewDequeueMessageListOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body DequeueBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "dequeue", err)
			}
			err = ValidateDequeueBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "dequeue", err)
			}
			return nil, NewDequeueBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body DequeueInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "dequeue", err)
			}
			err = ValidateDequeueInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "dequeue", err)
			}
			return nil, NewDequeueInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body DequeueInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "dequeue", err)
			}
			err = ValidateDequeueInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "dequeue", err)
			}
			return nil, NewDequeueInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body DequeueNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("queue", "dequeue", err)
			}
			err = ValidateDequeueNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("queue", "dequeue", err)
			}
			return nil, NewDequeueNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewDequeueNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewDequeueNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("queue", "dequeue", resp.StatusCode, string(body))
		}
	}
}

// unmarshalQueueListItemResponseBodyToQueueQueueListItem builds a value of
// type *queue.QueueListItem from a value of type *QueueListItemResponseBody.
func unmarshalQueueListItemResponseBodyToQueueQueueListItem(v *QueueListItemResponseBody) *queue.QueueListItem {
	res := &queue.QueueListItem{
		ID:          *v.ID,
		Name:        v.Name,
		Description: v.Description,
		Account:     *v.Account,
		Href:        *v.Href,
	}

	return res
}

// unmarshalLinkTResponseBodyToQueueLinkT builds a value of type *queue.LinkT
// from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToQueueLinkT(v *LinkTResponseBody) *queue.LinkT {
	res := &queue.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}

// unmarshalPublishedmessageResponseBodyToQueuePublishedmessage builds a value
// of type *queue.Publishedmessage from a value of type
// *PublishedmessageResponseBody.
func unmarshalPublishedmessageResponseBodyToQueuePublishedmessage(v *PublishedmessageResponseBody) *queue.Publishedmessage {
	res := &queue.Publishedmessage{
		ID:          v.ID,
		Content:     v.Content,
		Schema:      v.Schema,
		ContentType: v.ContentType,
	}

	return res
}
