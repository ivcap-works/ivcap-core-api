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
	"strings"

	search "github.com/ivcap-works/ivcap-core-api/gen/search"
	goahttp "goa.design/goa/v3/http"
)

// BuildSearchRequest instantiates a HTTP request object with method and path
// set to call the "search" service "search" endpoint
func (c *Client) BuildSearchRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SearchSearchPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("search", "search", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSearchRequest returns an encoder for requests sent to the search
// search server.
func EncodeSearchRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*search.SearchPayload)
		if !ok {
			return goahttp.ErrInvalidType("search", "search", "*search.SearchPayload", v)
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
		if p.AtTime != nil {
			values.Add("at-time", *p.AtTime)
		}
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		req.URL.RawQuery = values.Encode()
		body := p.Query
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("search", "search", err)
		}
		return nil
	}
}

// DecodeSearchResponse returns a decoder for responses returned by the search
// search endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeSearchResponse may return the following errors:
//   - "bad-request" (type *search.BadRequestT): http.StatusFailedDependency
//   - "invalid-parameter" (type *search.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *search.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *search.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *search.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *search.UnauthorizedT): http.StatusUnauthorized
//   - "unsupported-content-type" (type *search.UnsupportedContentTypeT): http.StatusUnsupportedMediaType
//   - error: internal error
func DecodeSearchResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body SearchResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			err = ValidateSearchResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			res := NewSearchListRTOK(&body)
			return res, nil
		case http.StatusFailedDependency:
			var (
				body SearchBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			err = ValidateSearchBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			return nil, NewSearchBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body SearchInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			err = ValidateSearchInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			return nil, NewSearchInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body SearchInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			err = ValidateSearchInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			return nil, NewSearchInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body SearchNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			err = ValidateSearchNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			return nil, NewSearchNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewSearchNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewSearchNotAuthorized()
		case http.StatusUnsupportedMediaType:
			var (
				body SearchUnsupportedContentTypeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("search", "search", err)
			}
			err = ValidateSearchUnsupportedContentTypeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("search", "search", err)
			}
			return nil, NewSearchUnsupportedContentType(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("search", "search", resp.StatusCode, string(body))
		}
	}
}

// unmarshalLinkTResponseBodyToSearchLinkT builds a value of type *search.LinkT
// from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToSearchLinkT(v *LinkTResponseBody) *search.LinkT {
	res := &search.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}
