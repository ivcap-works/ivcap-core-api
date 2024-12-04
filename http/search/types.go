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
	search "github.com/ivcap-works/ivcap-core-api/gen/search"
	goa "goa.design/goa/v3/pkg"
)

// SearchResponseBody is the type of the "search" service "search" endpoint
// HTTP response body.
type SearchResponseBody struct {
	// List of search result
	Items []any `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// SearchBadRequestResponseBody is the type of the "search" service "search"
// endpoint HTTP response body for the "bad-request" error.
type SearchBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SearchInvalidParameterResponseBody is the type of the "search" service
// "search" endpoint HTTP response body for the "invalid-parameter" error.
type SearchInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// SearchInvalidScopesResponseBody is the type of the "search" service "search"
// endpoint HTTP response body for the "invalid-scopes" error.
type SearchInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SearchNotImplementedResponseBody is the type of the "search" service
// "search" endpoint HTTP response body for the "not-implemented" error.
type SearchNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SearchUnsupportedContentTypeResponseBody is the type of the "search" service
// "search" endpoint HTTP response body for the "unsupported-content-type"
// error.
type SearchUnsupportedContentTypeResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LinkTResponseBody is used to define fields on response body types.
type LinkTResponseBody struct {
	// relation type
	Rel *string `form:"rel,omitempty" json:"rel,omitempty" xml:"rel,omitempty"`
	// mime type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// web link
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
}

// NewSearchListRTOK builds a "search" service "search" endpoint result from a
// HTTP "OK" response.
func NewSearchListRTOK(body *SearchResponseBody) *search.SearchListRT {
	v := &search.SearchListRT{
		AtTime: *body.AtTime,
	}
	v.Items = make([]any, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = val
	}
	v.Links = make([]*search.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToSearchLinkT(val)
	}

	return v
}

// NewSearchBadRequest builds a search service search endpoint bad-request
// error.
func NewSearchBadRequest(body *SearchBadRequestResponseBody) *search.BadRequestT {
	v := &search.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewSearchInvalidParameter builds a search service search endpoint
// invalid-parameter error.
func NewSearchInvalidParameter(body *SearchInvalidParameterResponseBody) *search.InvalidParameterT {
	v := &search.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewSearchInvalidScopes builds a search service search endpoint
// invalid-scopes error.
func NewSearchInvalidScopes(body *SearchInvalidScopesResponseBody) *search.InvalidScopesT {
	v := &search.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSearchNotImplemented builds a search service search endpoint
// not-implemented error.
func NewSearchNotImplemented(body *SearchNotImplementedResponseBody) *search.NotImplementedT {
	v := &search.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewSearchNotAvailable builds a search service search endpoint not-available
// error.
func NewSearchNotAvailable() *search.ServiceNotAvailableT {
	v := &search.ServiceNotAvailableT{}

	return v
}

// NewSearchNotAuthorized builds a search service search endpoint
// not-authorized error.
func NewSearchNotAuthorized() *search.UnauthorizedT {
	v := &search.UnauthorizedT{}

	return v
}

// NewSearchUnsupportedContentType builds a search service search endpoint
// unsupported-content-type error.
func NewSearchUnsupportedContentType(body *SearchUnsupportedContentTypeResponseBody) *search.UnsupportedContentTypeT {
	v := &search.UnsupportedContentTypeT{
		Message: *body.Message,
	}

	return v
}

// ValidateSearchResponseBody runs the validations defined on SearchResponseBody
func ValidateSearchResponseBody(body *SearchResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.AtTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("at-time", "body"))
	}
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.at-time", *body.AtTime, goa.FormatDateTime))
	}
	for _, e := range body.Links {
		if e != nil {
			if err2 := ValidateLinkTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateSearchBadRequestResponseBody runs the validations defined on
// search_bad-request_response_body
func ValidateSearchBadRequestResponseBody(body *SearchBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSearchInvalidParameterResponseBody runs the validations defined on
// search_invalid-parameter_response_body
func ValidateSearchInvalidParameterResponseBody(body *SearchInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSearchInvalidScopesResponseBody runs the validations defined on
// search_invalid-scopes_response_body
func ValidateSearchInvalidScopesResponseBody(body *SearchInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateSearchNotImplementedResponseBody runs the validations defined on
// search_not-implemented_response_body
func ValidateSearchNotImplementedResponseBody(body *SearchNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSearchUnsupportedContentTypeResponseBody runs the validations
// defined on search_unsupported-content-type_response_body
func ValidateSearchUnsupportedContentTypeResponseBody(body *SearchUnsupportedContentTypeResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLinkTResponseBody runs the validations defined on LinkTResponseBody
func ValidateLinkTResponseBody(body *LinkTResponseBody) (err error) {
	if body.Rel == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rel", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "body"))
	}
	return
}
