// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

	artifact "github.com/ivcap-works/ivcap-core-api/gen/artifact"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "artifact" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListArtifactPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("artifact", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the artifact list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*artifact.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("artifact", "list", "*artifact.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the artifact
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *artifact.BadRequestT): http.StatusFailedDependency
//   - "invalid-parameter" (type *artifact.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *artifact.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *artifact.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *artifact.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *artifact.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("artifact", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "list", err)
			}
			res := NewListArtifactListRTOK(&body)
			return res, nil
		case http.StatusFailedDependency:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("artifact", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "artifact" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*artifact.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("artifact", "read", "*artifact.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadArtifactPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("artifact", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the artifact read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*artifact.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("artifact", "read", "*artifact.ReadPayload", v)
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

// DecodeReadResponse returns a decoder for responses returned by the artifact
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *artifact.BadRequestT): http.StatusFailedDependency
//   - "invalid-scopes" (type *artifact.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *artifact.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *artifact.ResourceNotFoundT): http.StatusNotFound
//   - "not-available" (type *artifact.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *artifact.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("artifact", "read", err)
			}
			err = ValidateReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "read", err)
			}
			res := NewReadArtifactStatusRTOK(&body)
			return res, nil
		case http.StatusFailedDependency:
			var (
				body ReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "read", err)
			}
			err = ValidateReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "read", err)
			}
			return nil, NewReadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusServiceUnavailable:
			return nil, NewReadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("artifact", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildUploadRequest instantiates a HTTP request object with method and path
// set to call the "artifact" service "upload" endpoint
func (c *Client) BuildUploadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		body io.Reader
	)
	rd, ok := v.(*artifact.UploadRequestData)
	if !ok {
		return nil, goahttp.ErrInvalidType("artifact", "upload", "artifact.UploadRequestData", v)
	}
	body = rd.Body
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UploadArtifactPath()}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("artifact", "upload", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUploadRequest returns an encoder for requests sent to the artifact
// upload server.
func EncodeUploadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		data, ok := v.(*artifact.UploadRequestData)
		if !ok {
			return goahttp.ErrInvalidType("artifact", "upload", "*artifact.UploadRequestData", v)
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
		if p.ContentType != nil {
			head := *p.ContentType
			req.Header.Set("Content-Type", head)
		}
		if p.ContentEncoding != nil {
			head := *p.ContentEncoding
			req.Header.Set("Content-Encoding", head)
		}
		if p.ContentLength != nil {
			head := *p.ContentLength
			headStr := strconv.Itoa(head)
			req.Header.Set("Content-Length", headStr)
		}
		if p.Name != nil {
			head := *p.Name
			req.Header.Set("X-Name", head)
		}
		if p.Collection != nil {
			head := *p.Collection
			req.Header.Set("X-Collection", head)
		}
		if p.Policy != nil {
			head := *p.Policy
			req.Header.Set("X-Policy", head)
		}
		if p.XContentType != nil {
			head := *p.XContentType
			req.Header.Set("X-Content-Type", head)
		}
		if p.XContentLength != nil {
			head := *p.XContentLength
			headStr := strconv.Itoa(head)
			req.Header.Set("X-Content-Length", headStr)
		}
		if p.UploadLength != nil {
			head := *p.UploadLength
			headStr := strconv.Itoa(head)
			req.Header.Set("Upload-Length", headStr)
		}
		if p.TusResumable != nil {
			head := *p.TusResumable
			req.Header.Set("Tus-Resumable", head)
		}
		return nil
	}
}

// DecodeUploadResponse returns a decoder for responses returned by the
// artifact upload endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUploadResponse may return the following errors:
//   - "bad-request" (type *artifact.BadRequestT): http.StatusFailedDependency
//   - "invalid-scopes" (type *artifact.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *artifact.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *artifact.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *artifact.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUploadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body UploadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "upload", err)
			}
			err = ValidateUploadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "upload", err)
			}
			var (
				location     string
				tusResumable *string
				tusOffset    *int64
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("location", "header"))
			}
			location = locationRaw
			tusResumableRaw := resp.Header.Get("Tus-Resumable")
			if tusResumableRaw != "" {
				tusResumable = &tusResumableRaw
			}
			{
				tusOffsetRaw := resp.Header.Get("Upload-Offset")
				if tusOffsetRaw != "" {
					v, err2 := strconv.ParseInt(tusOffsetRaw, 10, 64)
					if err2 != nil {
						err = goa.MergeErrors(err, goa.InvalidFieldTypeError("tus-offset", tusOffsetRaw, "integer"))
					}
					tusOffset = &v
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "upload", err)
			}
			res := NewUploadArtifactUploadRTCreated(&body, location, tusResumable, tusOffset)
			return res, nil
		case http.StatusFailedDependency:
			var (
				body UploadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "upload", err)
			}
			err = ValidateUploadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "upload", err)
			}
			return nil, NewUploadBadRequest(&body)
		case http.StatusForbidden:
			var (
				body UploadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "upload", err)
			}
			err = ValidateUploadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "upload", err)
			}
			return nil, NewUploadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UploadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("artifact", "upload", err)
			}
			err = ValidateUploadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("artifact", "upload", err)
			}
			return nil, NewUploadNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewUploadNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewUploadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("artifact", "upload", resp.StatusCode, string(body))
		}
	}
}

// // BuildUploadStreamPayload creates a streaming endpoint request payload from
// the method payload and the path to the file to be streamed
func BuildUploadStreamPayload(payload any, fpath string) (*artifact.UploadRequestData, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return &artifact.UploadRequestData{
		Payload: payload.(*artifact.UploadPayload),
		Body:    f,
	}, nil
}

// unmarshalArtifactListItemResponseBodyToArtifactArtifactListItem builds a
// value of type *artifact.ArtifactListItem from a value of type
// *ArtifactListItemResponseBody.
func unmarshalArtifactListItemResponseBodyToArtifactArtifactListItem(v *ArtifactListItemResponseBody) *artifact.ArtifactListItem {
	res := &artifact.ArtifactListItem{
		ID:       *v.ID,
		Name:     v.Name,
		Status:   *v.Status,
		Size:     v.Size,
		MimeType: v.MimeType,
		Href:     *v.Href,
	}

	return res
}

// unmarshalLinkTResponseBodyToArtifactLinkT builds a value of type
// *artifact.LinkT from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToArtifactLinkT(v *LinkTResponseBody) *artifact.LinkT {
	res := &artifact.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}
