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
	artifact "github.com/ivcap-works/ivcap-core-api/gen/artifact"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "artifact" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Artifacts
	Items []*ArtifactListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// ReadResponseBody is the type of the "artifact" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// Artifact ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Artifact status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Mime-type of data
	MimeType *string `form:"mime-type,omitempty" json:"mime-type,omitempty" xml:"mime-type,omitempty"`
	// Size of data
	Size *int64 `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// URL of object this artifact is caching
	CacheOf *string `form:"cache-of,omitempty" json:"cache-of,omitempty" xml:"cache-of,omitempty"`
	// ETAG of artifact
	Etag *string `form:"etag,omitempty" json:"etag,omitempty" xml:"etag,omitempty"`
	// DateTime artifact was created
	CreatedAt *string `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	// DateTime artifact was last modified
	LastModifiedAt *string `form:"last-modified-at,omitempty" json:"last-modified-at,omitempty" xml:"last-modified-at,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// Reference to billable account
	Account  *string              `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	DataHref *string              `form:"data-href,omitempty" json:"data-href,omitempty" xml:"data-href,omitempty"`
	Links    []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// UploadResponseBody is the type of the "artifact" service "upload" endpoint
// HTTP response body.
type UploadResponseBody struct {
	// Artifact ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Artifact status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Mime-type of data
	MimeType *string `form:"mime-type,omitempty" json:"mime-type,omitempty" xml:"mime-type,omitempty"`
	// Size of data
	Size *int64 `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// URL of object this artifact is caching
	CacheOf *string `form:"cache-of,omitempty" json:"cache-of,omitempty" xml:"cache-of,omitempty"`
	// ETAG of artifact
	Etag *string `form:"etag,omitempty" json:"etag,omitempty" xml:"etag,omitempty"`
	// DateTime artifact was created
	CreatedAt *string `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	// DateTime artifact was last modified
	LastModifiedAt *string `form:"last-modified-at,omitempty" json:"last-modified-at,omitempty" xml:"last-modified-at,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// Reference to billable account
	Account  *string              `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	DataHref *string              `form:"data-href,omitempty" json:"data-href,omitempty" xml:"data-href,omitempty"`
	Links    []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// ListBadRequestResponseBody is the type of the "artifact" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "artifact" service
// "list" endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "artifact" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "artifact" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "artifact" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "artifact" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "artifact" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "artifact" service "read"
// endpoint HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UploadBadRequestResponseBody is the type of the "artifact" service "upload"
// endpoint HTTP response body for the "bad-request" error.
type UploadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UploadInvalidScopesResponseBody is the type of the "artifact" service
// "upload" endpoint HTTP response body for the "invalid-scopes" error.
type UploadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UploadNotImplementedResponseBody is the type of the "artifact" service
// "upload" endpoint HTTP response body for the "not-implemented" error.
type UploadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ArtifactListItemResponseBody is used to define fields on response body types.
type ArtifactListItemResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Artifact status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Size of artifact in bytes
	Size *int64 `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Mime (content) type of artifact
	MimeType *string `form:"mime-type,omitempty" json:"mime-type,omitempty" xml:"mime-type,omitempty"`
	// time this artifact was created
	CreatedAt *string `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
	Href      *string `json:"href,omitempty"`
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

// NewListArtifactListRTOK builds a "artifact" service "list" endpoint result
// from a HTTP "OK" response.
func NewListArtifactListRTOK(body *ListResponseBody) *artifact.ArtifactListRT {
	v := &artifact.ArtifactListRT{
		AtTime: body.AtTime,
	}
	v.Items = make([]*artifact.ArtifactListItem, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalArtifactListItemResponseBodyToArtifactArtifactListItem(val)
	}
	v.Links = make([]*artifact.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToArtifactLinkT(val)
	}

	return v
}

// NewListBadRequest builds a artifact service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *artifact.BadRequestT {
	v := &artifact.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidParameter builds a artifact service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *artifact.InvalidParameterT {
	v := &artifact.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a artifact service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *artifact.InvalidScopesT {
	v := &artifact.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a artifact service list endpoint
// not-implemented error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *artifact.NotImplementedT {
	v := &artifact.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAvailable builds a artifact service list endpoint not-available
// error.
func NewListNotAvailable() *artifact.ServiceNotAvailableT {
	v := &artifact.ServiceNotAvailableT{}

	return v
}

// NewListNotAuthorized builds a artifact service list endpoint not-authorized
// error.
func NewListNotAuthorized() *artifact.UnauthorizedT {
	v := &artifact.UnauthorizedT{}

	return v
}

// NewReadArtifactStatusRTCreated builds a "artifact" service "read" endpoint
// result from a HTTP "Created" response.
func NewReadArtifactStatusRTCreated(body *ReadResponseBody) *artifact.ArtifactStatusRT {
	v := &artifact.ArtifactStatusRT{
		ID:             *body.ID,
		Name:           body.Name,
		Status:         *body.Status,
		MimeType:       body.MimeType,
		Size:           body.Size,
		CacheOf:        body.CacheOf,
		Etag:           body.Etag,
		CreatedAt:      body.CreatedAt,
		LastModifiedAt: body.LastModifiedAt,
		Policy:         body.Policy,
		Account:        body.Account,
		DataHref:       body.DataHref,
	}
	v.Links = make([]*artifact.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToArtifactLinkT(val)
	}

	return v
}

// NewReadBadRequest builds a artifact service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *artifact.BadRequestT {
	v := &artifact.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidScopes builds a artifact service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *artifact.InvalidScopesT {
	v := &artifact.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a artifact service read endpoint
// not-implemented error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *artifact.NotImplementedT {
	v := &artifact.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a artifact service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *artifact.ResourceNotFoundT {
	v := &artifact.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAvailable builds a artifact service read endpoint not-available
// error.
func NewReadNotAvailable() *artifact.ServiceNotAvailableT {
	v := &artifact.ServiceNotAvailableT{}

	return v
}

// NewReadNotAuthorized builds a artifact service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *artifact.UnauthorizedT {
	v := &artifact.UnauthorizedT{}

	return v
}

// NewUploadArtifactUploadRTCreated builds a "artifact" service "upload"
// endpoint result from a HTTP "Created" response.
func NewUploadArtifactUploadRTCreated(body *UploadResponseBody, location string, tusResumable *string, tusOffset *int64) *artifact.ArtifactUploadRT {
	v := &artifact.ArtifactUploadRT{
		ID:             *body.ID,
		Name:           body.Name,
		Status:         *body.Status,
		MimeType:       body.MimeType,
		Size:           body.Size,
		CacheOf:        body.CacheOf,
		Etag:           body.Etag,
		CreatedAt:      body.CreatedAt,
		LastModifiedAt: body.LastModifiedAt,
		Policy:         body.Policy,
		Account:        body.Account,
		DataHref:       body.DataHref,
	}
	v.Links = make([]*artifact.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToArtifactLinkT(val)
	}
	v.Location = location
	v.TusResumable = tusResumable
	v.TusOffset = tusOffset

	return v
}

// NewUploadBadRequest builds a artifact service upload endpoint bad-request
// error.
func NewUploadBadRequest(body *UploadBadRequestResponseBody) *artifact.BadRequestT {
	v := &artifact.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewUploadInvalidScopes builds a artifact service upload endpoint
// invalid-scopes error.
func NewUploadInvalidScopes(body *UploadInvalidScopesResponseBody) *artifact.InvalidScopesT {
	v := &artifact.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUploadNotImplemented builds a artifact service upload endpoint
// not-implemented error.
func NewUploadNotImplemented(body *UploadNotImplementedResponseBody) *artifact.NotImplementedT {
	v := &artifact.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewUploadNotAvailable builds a artifact service upload endpoint
// not-available error.
func NewUploadNotAvailable() *artifact.ServiceNotAvailableT {
	v := &artifact.ServiceNotAvailableT{}

	return v
}

// NewUploadNotAuthorized builds a artifact service upload endpoint
// not-authorized error.
func NewUploadNotAuthorized() *artifact.UnauthorizedT {
	v := &artifact.UnauthorizedT{}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	for _, e := range body.Items {
		if e != nil {
			if err2 := ValidateArtifactListItemResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
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

// ValidateReadResponseBody runs the validations defined on ReadResponseBody
func ValidateReadResponseBody(body *ReadResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Status != nil {
		if !(*body.Status == "pending" || *body.Status == "partial" || *body.Status == "ready" || *body.Status == "error" || *body.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"pending", "partial", "ready", "error", "unknown"}))
		}
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created-at", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.LastModifiedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.last-modified-at", *body.LastModifiedAt, goa.FormatDateTime))
	}
	if body.Policy != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
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

// ValidateUploadResponseBody runs the validations defined on UploadResponseBody
func ValidateUploadResponseBody(body *UploadResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Status != nil {
		if !(*body.Status == "pending" || *body.Status == "partial" || *body.Status == "ready" || *body.Status == "error" || *body.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"pending", "partial", "ready", "error", "unknown"}))
		}
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created-at", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.LastModifiedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.last-modified-at", *body.LastModifiedAt, goa.FormatDateTime))
	}
	if body.Policy != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
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

// ValidateReadBadRequestResponseBody runs the validations defined on
// read_bad-request_response_body
func ValidateReadBadRequestResponseBody(body *ReadBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateReadInvalidScopesResponseBody runs the validations defined on
// read_invalid-scopes_response_body
func ValidateReadInvalidScopesResponseBody(body *ReadInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateReadNotImplementedResponseBody runs the validations defined on
// read_not-implemented_response_body
func ValidateReadNotImplementedResponseBody(body *ReadNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateReadNotFoundResponseBody runs the validations defined on
// read_not-found_response_body
func ValidateReadNotFoundResponseBody(body *ReadNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	return
}

// ValidateUploadBadRequestResponseBody runs the validations defined on
// upload_bad-request_response_body
func ValidateUploadBadRequestResponseBody(body *UploadBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUploadInvalidScopesResponseBody runs the validations defined on
// upload_invalid-scopes_response_body
func ValidateUploadInvalidScopesResponseBody(body *UploadInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUploadNotImplementedResponseBody runs the validations defined on
// upload_not-implemented_response_body
func ValidateUploadNotImplementedResponseBody(body *UploadNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateArtifactListItemResponseBody runs the validations defined on
// ArtifactListItemResponseBody
func ValidateArtifactListItemResponseBody(body *ArtifactListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created-at", "body"))
	}
	if body.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Status != nil {
		if !(*body.Status == "pending" || *body.Status == "partial" || *body.Status == "ready" || *body.Status == "error" || *body.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"pending", "partial", "ready", "error", "unknown"}))
		}
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created-at", *body.CreatedAt, goa.FormatDateTime))
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
