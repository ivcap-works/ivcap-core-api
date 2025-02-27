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
	dashboard "github.com/ivcap-works/ivcap-core-api/gen/dashboard"
	dashboardviews "github.com/ivcap-works/ivcap-core-api/gen/dashboard/views"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "dashboard" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Dashboards
	Items []*DashboardListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// ListBadRequestResponseBody is the type of the "dashboard" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "dashboard" service
// "list" endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "dashboard" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "dashboard" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DashboardListItemResponseBody is used to define fields on response body
// types.
type DashboardListItemResponseBody struct {
	// dashboard id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// dashboard uid
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Dashboard title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Dashboard url
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// NewListDashboardListRTOK builds a "dashboard" service "list" endpoint result
// from a HTTP "OK" response.
func NewListDashboardListRTOK(body *ListResponseBody) *dashboardviews.DashboardListRTView {
	v := &dashboardviews.DashboardListRTView{}
	v.Items = make([]*dashboardviews.DashboardListItemView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalDashboardListItemResponseBodyToDashboardviewsDashboardListItemView(val)
	}

	return v
}

// NewListBadRequest builds a dashboard service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *dashboard.BadRequestT {
	v := &dashboard.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidParameter builds a dashboard service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *dashboard.InvalidParameterT {
	v := &dashboard.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a dashboard service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *dashboard.InvalidScopesT {
	v := &dashboard.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a dashboard service list endpoint
// not-implemented error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *dashboard.NotImplementedT {
	v := &dashboard.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAvailable builds a dashboard service list endpoint not-available
// error.
func NewListNotAvailable() *dashboard.ServiceNotAvailableT {
	v := &dashboard.ServiceNotAvailableT{}

	return v
}

// NewListNotAuthorized builds a dashboard service list endpoint not-authorized
// error.
func NewListNotAuthorized() *dashboard.UnauthorizedT {
	v := &dashboard.UnauthorizedT{}

	return v
}

// ValidateListBadRequestResponseBody runs the validations defined on
// list_bad-request_response_body
func ValidateListBadRequestResponseBody(body *ListBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListInvalidParameterResponseBody runs the validations defined on
// list_invalid-parameter_response_body
func ValidateListInvalidParameterResponseBody(body *ListInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListInvalidScopesResponseBody runs the validations defined on
// list_invalid-scopes_response_body
func ValidateListInvalidScopesResponseBody(body *ListInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateListNotImplementedResponseBody runs the validations defined on
// list_not-implemented_response_body
func ValidateListNotImplementedResponseBody(body *ListNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDashboardListItemResponseBody runs the validations defined on
// DashboardListItemResponseBody
func ValidateDashboardListItemResponseBody(body *DashboardListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.UID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("uid", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	return
}
