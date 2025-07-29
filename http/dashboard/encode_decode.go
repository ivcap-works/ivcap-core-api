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

	dashboard "github.com/ivcap-works/ivcap-core-api/gen/dashboard"
	dashboardviews "github.com/ivcap-works/ivcap-core-api/gen/dashboard/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "dashboard" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListDashboardPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("dashboard", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the dashboard list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*dashboard.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("dashboard", "list", "*dashboard.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the dashboard
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *dashboard.BadRequestT): http.StatusBadRequest
//   - "invalid-parameter" (type *dashboard.InvalidParameterT): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *dashboard.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *dashboard.NotImplementedT): http.StatusNotImplemented
//   - "not-available" (type *dashboard.ServiceNotAvailableT): http.StatusServiceUnavailable
//   - "not-authorized" (type *dashboard.UnauthorizedT): http.StatusUnauthorized
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
				return nil, goahttp.ErrDecodingError("dashboard", "list", err)
			}
			p := NewListDashboardListRTOK(&body)
			view := "default"
			vres := &dashboardviews.DashboardListRT{Projected: p, View: view}
			if err = dashboardviews.ValidateDashboardListRT(vres); err != nil {
				return nil, goahttp.ErrValidationError("dashboard", "list", err)
			}
			res := dashboard.NewDashboardListRT(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("dashboard", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("dashboard", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("dashboard", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("dashboard", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("dashboard", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("dashboard", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("dashboard", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("dashboard", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusServiceUnavailable:
			return nil, NewListNotAvailable()
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("dashboard", "list", resp.StatusCode, string(body))
		}
	}
}

// unmarshalDashboardListItemResponseBodyToDashboardviewsDashboardListItemView
// builds a value of type *dashboardviews.DashboardListItemView from a value of
// type *DashboardListItemResponseBody.
func unmarshalDashboardListItemResponseBodyToDashboardviewsDashboardListItemView(v *DashboardListItemResponseBody) *dashboardviews.DashboardListItemView {
	res := &dashboardviews.DashboardListItemView{
		ID:    v.ID,
		UID:   v.UID,
		Title: v.Title,
		URL:   v.URL,
	}

	return res
}
