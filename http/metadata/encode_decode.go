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
	metadata "github.com/reinventingscience/ivcap-core-api/gen/metadata"
	metadataviews "github.com/reinventingscience/ivcap-core-api/gen/metadata/views"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "metadata" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListMetadataPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("metadata", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the metadata list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*metadata.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("metadata", "list", "*metadata.ListPayload", v)
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
		if p.EntityID != nil {
			values.Add("entity-id", *p.EntityID)
		}
		if p.Schema != nil {
			values.Add("schema", *p.Schema)
		}
		if p.AspectPath != nil {
			values.Add("aspect-path", *p.AspectPath)
		}
		if p.AtTime != nil {
			values.Add("at-time", *p.AtTime)
		}
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("filter", p.Filter)
		values.Add("order-by", p.OrderBy)
		if p.OrderDesc != nil {
			values.Add("order-desc", fmt.Sprintf("%v", *p.OrderDesc))
		}
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the metadata
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *metadata.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *metadata.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *metadata.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *metadata.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *metadata.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *metadata.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("metadata", "list", err)
			}
			p := NewListMetaRTViewOK(&body)
			view := "default"
			vres := &metadataviews.ListMetaRT{Projected: p, View: view}
			if err = metadataviews.ValidateListMetaRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("metadata", "list", err)
			}
			res := metadata.NewListMetaRT(vres)
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
					return nil, goahttp.ErrDecodingError("metadata", "list", err)
				}
				err = ValidateListBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("metadata", "list", err)
				}
				return nil, NewListBadRequest(&body)
			case "invalid-credential":
				return nil, NewListInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("metadata", "list", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("metadata", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "metadata" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*metadata.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("metadata", "read", "*metadata.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadMetadataPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("metadata", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the metadata read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*metadata.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("metadata", "read", "*metadata.ReadPayload", v)
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

// DecodeReadResponse returns a decoder for responses returned by the metadata
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *metadata.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *metadata.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-scopes" (type *metadata.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *metadata.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *metadata.ResourceNotFoundT): http.StatusNotFound
//   - "not-authorized" (type *metadata.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("metadata", "read", err)
			}
			p := NewReadMetadataRecordRTOK(&body)
			view := "default"
			vres := &metadataviews.MetadataRecordRT{Projected: p, View: view}
			if err = metadataviews.ValidateMetadataRecordRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("metadata", "read", err)
			}
			res := metadata.NewMetadataRecordRT(vres)
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
					return nil, goahttp.ErrDecodingError("metadata", "read", err)
				}
				err = ValidateReadBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("metadata", "read", err)
				}
				return nil, NewReadBadRequest(&body)
			case "invalid-credential":
				return nil, NewReadInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("metadata", "read", resp.StatusCode, string(body))
			}
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("metadata", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "metadata" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddMetadataPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("metadata", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the metadata add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*metadata.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("metadata", "add", "*metadata.AddPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		{
			head := p.ContentType
			req.Header.Set("Content-Type", head)
		}
		values := req.URL.Query()
		values.Add("entity-id", p.EntityID)
		values.Add("schema", p.Schema)
		if p.PolicyID != nil {
			values.Add("policy-id", *p.PolicyID)
		}
		req.URL.RawQuery = values.Encode()
		body := p.Aspect
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("metadata", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the metadata
// add endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeAddResponse may return the following errors:
//   - "bad-request" (type *metadata.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *metadata.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *metadata.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *metadata.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *metadata.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *metadata.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body AddResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "add", err)
			}
			p := NewAddMetaRTViewOK(&body)
			view := "default"
			vres := &metadataviews.AddMetaRT{Projected: p, View: view}
			if err = metadataviews.ValidateAddMetaRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("metadata", "add", err)
			}
			res := metadata.NewAddMetaRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body AddBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("metadata", "add", err)
				}
				err = ValidateAddBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("metadata", "add", err)
				}
				return nil, NewAddBadRequest(&body)
			case "invalid-credential":
				return nil, NewAddInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("metadata", "add", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body AddInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "add", err)
			}
			err = ValidateAddInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "add", err)
			}
			return nil, NewAddInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body AddInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "add", err)
			}
			err = ValidateAddInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "add", err)
			}
			return nil, NewAddInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body AddNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "add", err)
			}
			err = ValidateAddNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "add", err)
			}
			return nil, NewAddNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewAddNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("metadata", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateOneRequest instantiates a HTTP request object with method and
// path set to call the "metadata" service "update_one" endpoint
func (c *Client) BuildUpdateOneRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateOneMetadataPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("metadata", "update_one", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateOneRequest returns an encoder for requests sent to the metadata
// update_one server.
func EncodeUpdateOneRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*metadata.UpdateOnePayload)
		if !ok {
			return goahttp.ErrInvalidType("metadata", "update_one", "*metadata.UpdateOnePayload", v)
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
		values.Add("entity-id", p.EntityID)
		values.Add("schema", p.Schema)
		if p.PolicyID != nil {
			values.Add("policy-id", *p.PolicyID)
		}
		req.URL.RawQuery = values.Encode()
		body := p.Aspect
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("metadata", "update_one", err)
		}
		return nil
	}
}

// DecodeUpdateOneResponse returns a decoder for responses returned by the
// metadata update_one endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateOneResponse may return the following errors:
//   - "bad-request" (type *metadata.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *metadata.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *metadata.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *metadata.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *metadata.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *metadata.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateOneResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body UpdateOneResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_one", err)
			}
			p := NewUpdateOneAddMetaRTOK(&body)
			view := "default"
			vres := &metadataviews.AddMetaRT{Projected: p, View: view}
			if err = metadataviews.ValidateAddMetaRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_one", err)
			}
			res := metadata.NewAddMetaRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body UpdateOneBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("metadata", "update_one", err)
				}
				err = ValidateUpdateOneBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("metadata", "update_one", err)
				}
				return nil, NewUpdateOneBadRequest(&body)
			case "invalid-credential":
				return nil, NewUpdateOneInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("metadata", "update_one", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body UpdateOneInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_one", err)
			}
			err = ValidateUpdateOneInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_one", err)
			}
			return nil, NewUpdateOneInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body UpdateOneInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_one", err)
			}
			err = ValidateUpdateOneInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_one", err)
			}
			return nil, NewUpdateOneInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UpdateOneNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_one", err)
			}
			err = ValidateUpdateOneNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_one", err)
			}
			return nil, NewUpdateOneNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewUpdateOneNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("metadata", "update_one", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRecordRequest instantiates a HTTP request object with method and
// path set to call the "metadata" service "update_record" endpoint
func (c *Client) BuildUpdateRecordRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*metadata.UpdateRecordPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("metadata", "update_record", "*metadata.UpdateRecordPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateRecordMetadataPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("metadata", "update_record", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRecordRequest returns an encoder for requests sent to the
// metadata update_record server.
func EncodeUpdateRecordRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*metadata.UpdateRecordPayload)
		if !ok {
			return goahttp.ErrInvalidType("metadata", "update_record", "*metadata.UpdateRecordPayload", v)
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
		if p.EntityID != nil {
			values.Add("entity-id", *p.EntityID)
		}
		if p.Schema != nil {
			values.Add("schema", *p.Schema)
		}
		if p.PolicyID != nil {
			values.Add("policy-id", *p.PolicyID)
		}
		req.URL.RawQuery = values.Encode()
		body := p.Aspect
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("metadata", "update_record", err)
		}
		return nil
	}
}

// DecodeUpdateRecordResponse returns a decoder for responses returned by the
// metadata update_record endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeUpdateRecordResponse may return the following errors:
//   - "bad-request" (type *metadata.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *metadata.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *metadata.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *metadata.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *metadata.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *metadata.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateRecordResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body UpdateRecordResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_record", err)
			}
			p := NewUpdateRecordAddMetaRTOK(&body)
			view := "default"
			vres := &metadataviews.AddMetaRT{Projected: p, View: view}
			if err = metadataviews.ValidateAddMetaRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_record", err)
			}
			res := metadata.NewAddMetaRT(vres)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body UpdateRecordBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("metadata", "update_record", err)
				}
				err = ValidateUpdateRecordBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("metadata", "update_record", err)
				}
				return nil, NewUpdateRecordBadRequest(&body)
			case "invalid-credential":
				return nil, NewUpdateRecordInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("metadata", "update_record", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body UpdateRecordInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_record", err)
			}
			err = ValidateUpdateRecordInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_record", err)
			}
			return nil, NewUpdateRecordInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body UpdateRecordInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_record", err)
			}
			err = ValidateUpdateRecordInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_record", err)
			}
			return nil, NewUpdateRecordInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UpdateRecordNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "update_record", err)
			}
			err = ValidateUpdateRecordNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "update_record", err)
			}
			return nil, NewUpdateRecordNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewUpdateRecordNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("metadata", "update_record", resp.StatusCode, string(body))
		}
	}
}

// BuildRevokeRequest instantiates a HTTP request object with method and path
// set to call the "metadata" service "revoke" endpoint
func (c *Client) BuildRevokeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*metadata.RevokePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("metadata", "revoke", "*metadata.RevokePayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RevokeMetadataPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("metadata", "revoke", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRevokeRequest returns an encoder for requests sent to the metadata
// revoke server.
func EncodeRevokeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*metadata.RevokePayload)
		if !ok {
			return goahttp.ErrInvalidType("metadata", "revoke", "*metadata.RevokePayload", v)
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

// DecodeRevokeResponse returns a decoder for responses returned by the
// metadata revoke endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRevokeResponse may return the following errors:
//   - "bad-request" (type *metadata.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *metadata.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *metadata.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *metadata.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *metadata.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *metadata.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeRevokeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
					body RevokeBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("metadata", "revoke", err)
				}
				err = ValidateRevokeBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("metadata", "revoke", err)
				}
				return nil, NewRevokeBadRequest(&body)
			case "invalid-credential":
				return nil, NewRevokeInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("metadata", "revoke", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body RevokeInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "revoke", err)
			}
			err = ValidateRevokeInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "revoke", err)
			}
			return nil, NewRevokeInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body RevokeInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "revoke", err)
			}
			err = ValidateRevokeInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "revoke", err)
			}
			return nil, NewRevokeInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body RevokeNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("metadata", "revoke", err)
			}
			err = ValidateRevokeNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("metadata", "revoke", err)
			}
			return nil, NewRevokeNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewRevokeNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("metadata", "revoke", resp.StatusCode, string(body))
		}
	}
}

// unmarshalMetadataListItemRTResponseBodyToMetadataviewsMetadataListItemRTView
// builds a value of type *metadataviews.MetadataListItemRTView from a value of
// type *MetadataListItemRTResponseBody.
func unmarshalMetadataListItemRTResponseBodyToMetadataviewsMetadataListItemRTView(v *MetadataListItemRTResponseBody) *metadataviews.MetadataListItemRTView {
	res := &metadataviews.MetadataListItemRTView{
		RecordID:      v.RecordID,
		Entity:        v.Entity,
		Schema:        v.Schema,
		Aspect:        v.Aspect,
		AspectContext: v.AspectContext,
	}

	return res
}

// unmarshalNavTResponseBodyToMetadataviewsNavTView builds a value of type
// *metadataviews.NavTView from a value of type *NavTResponseBody.
func unmarshalNavTResponseBodyToMetadataviewsNavTView(v *NavTResponseBody) *metadataviews.NavTView {
	res := &metadataviews.NavTView{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}
