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
	artifact "github.com/reinventingscience/ivcap-core-api/gen/artifact"
	artifactviews "github.com/reinventingscience/ivcap-core-api/gen/artifact/views"

	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "artifact" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Artifacts
	Artifacts []*ArtifactListItemResponseBody `form:"artifacts,omitempty" json:"artifacts,omitempty" xml:"artifacts,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	// Navigation links
	Links *NavTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
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
	// Reference to policy controlling access
	Policy *RefTResponseBody `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// Reference to billable account
	Account *RefTResponseBody `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Link to retrieve the artifact data
	Data  *SelfTResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
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
	// Reference to policy controlling access
	Policy *RefTResponseBody `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// Reference to billable account
	Account *RefTResponseBody `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Link to retrieve the artifact data
	Data  *SelfTResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// link back to record
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// indicate version of TUS supported
	TusResumable *string `form:"tus_resumable,omitempty" json:"tus_resumable,omitempty" xml:"tus_resumable,omitempty"`
	// TUS offset for partially uploaded content
	TusOffset *int64 `form:"tus_offset,omitempty" json:"tus_offset,omitempty" xml:"tus_offset,omitempty"`
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

// ArtifactListItemResponseBody is used to define fields on response body types.
type ArtifactListItemResponseBody struct {
	// Artifact ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Artifact status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Size of aritfact in bytes
	Size *int64 `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Mime (content) type of artifact
	MimeType *string            `form:"mime-type,omitempty" json:"mime-type,omitempty" xml:"mime-type,omitempty"`
	Links    *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// SelfTResponseBody is used to define fields on response body types.
type SelfTResponseBody struct {
	Self        *string                   `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	DescribedBy *DescribedByTResponseBody `form:"describedBy,omitempty" json:"describedBy,omitempty" xml:"describedBy,omitempty"`
}

// DescribedByTResponseBody is used to define fields on response body types.
type DescribedByTResponseBody struct {
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// NavTResponseBody is used to define fields on response body types.
type NavTResponseBody struct {
	Self  *string `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	First *string `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Next  *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
}

// RefTResponseBody is used to define fields on response body types.
type RefTResponseBody struct {
	ID    *string            `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// NewListArtifactListRTOK builds a "artifact" service "list" endpoint result
// from a HTTP "OK" response.
func NewListArtifactListRTOK(body *ListResponseBody) *artifactviews.ArtifactListRTView {
	v := &artifactviews.ArtifactListRTView{
		AtTime: body.AtTime,
	}
	v.Artifacts = make([]*artifactviews.ArtifactListItemView, len(body.Artifacts))
	for i, val := range body.Artifacts {
		v.Artifacts[i] = unmarshalArtifactListItemResponseBodyToArtifactviewsArtifactListItemView(val)
	}
	v.Links = unmarshalNavTResponseBodyToArtifactviewsNavTView(body.Links)

	return v
}

// NewListBadRequest builds a artifact service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *artifact.BadRequestT {
	v := &artifact.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidCredential builds a artifact service list endpoint
// invalid-credential error.
func NewListInvalidCredential() *artifact.InvalidCredentialsT {
	v := &artifact.InvalidCredentialsT{}

	return v
}

// NewListInvalidParameter builds a artifact service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *artifact.InvalidParameterValue {
	v := &artifact.InvalidParameterValue{
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

// NewListNotAuthorized builds a artifact service list endpoint not-authorized
// error.
func NewListNotAuthorized() *artifact.UnauthorizedT {
	v := &artifact.UnauthorizedT{}

	return v
}

// NewUploadArtifactStatusRTCreated builds a "artifact" service "upload"
// endpoint result from a HTTP "Created" response.
func NewUploadArtifactStatusRTCreated(body *UploadResponseBody, location *string, tusResumable *string, tusOffset *int64) *artifactviews.ArtifactStatusRTView {
	v := &artifactviews.ArtifactStatusRTView{
		ID:             body.ID,
		Name:           body.Name,
		Status:         body.Status,
		MimeType:       body.MimeType,
		Size:           body.Size,
		CacheOf:        body.CacheOf,
		Etag:           body.Etag,
		CreatedAt:      body.CreatedAt,
		LastModifiedAt: body.LastModifiedAt,
	}
	if body.Policy != nil {
		v.Policy = unmarshalRefTResponseBodyToArtifactviewsRefTView(body.Policy)
	}
	if body.Account != nil {
		v.Account = unmarshalRefTResponseBodyToArtifactviewsRefTView(body.Account)
	}
	if body.Data != nil {
		v.Data = unmarshalSelfTResponseBodyToArtifactviewsSelfTView(body.Data)
	}
	v.Links = unmarshalSelfTResponseBodyToArtifactviewsSelfTView(body.Links)
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

// NewUploadInvalidCredential builds a artifact service upload endpoint
// invalid-credential error.
func NewUploadInvalidCredential() *artifact.InvalidCredentialsT {
	v := &artifact.InvalidCredentialsT{}

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

// NewUploadNotAuthorized builds a artifact service upload endpoint
// not-authorized error.
func NewUploadNotAuthorized() *artifact.UnauthorizedT {
	v := &artifact.UnauthorizedT{}

	return v
}

// NewReadArtifactStatusRTOK builds a "artifact" service "read" endpoint result
// from a HTTP "OK" response.
func NewReadArtifactStatusRTOK(body *ReadResponseBody) *artifactviews.ArtifactStatusRTView {
	v := &artifactviews.ArtifactStatusRTView{
		ID:             body.ID,
		Name:           body.Name,
		Status:         body.Status,
		MimeType:       body.MimeType,
		Size:           body.Size,
		CacheOf:        body.CacheOf,
		Etag:           body.Etag,
		CreatedAt:      body.CreatedAt,
		LastModifiedAt: body.LastModifiedAt,
		Location:       body.Location,
		TusResumable:   body.TusResumable,
		TusOffset:      body.TusOffset,
	}
	if body.Policy != nil {
		v.Policy = unmarshalRefTResponseBodyToArtifactviewsRefTView(body.Policy)
	}
	if body.Account != nil {
		v.Account = unmarshalRefTResponseBodyToArtifactviewsRefTView(body.Account)
	}
	if body.Data != nil {
		v.Data = unmarshalSelfTResponseBodyToArtifactviewsSelfTView(body.Data)
	}
	v.Links = unmarshalSelfTResponseBodyToArtifactviewsSelfTView(body.Links)

	return v
}

// NewReadBadRequest builds a artifact service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *artifact.BadRequestT {
	v := &artifact.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidCredential builds a artifact service read endpoint
// invalid-credential error.
func NewReadInvalidCredential() *artifact.InvalidCredentialsT {
	v := &artifact.InvalidCredentialsT{}

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

// NewReadNotAuthorized builds a artifact service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *artifact.UnauthorizedT {
	v := &artifact.UnauthorizedT{}

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

// ValidateArtifactListItemResponseBody runs the validations defined on
// ArtifactListItemResponseBody
func ValidateArtifactListItemResponseBody(body *ArtifactListItemResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "pending" || *body.Status == "building" || *body.Status == "ready" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"pending", "building", "ready", "error"}))
		}
	}
	if body.Links != nil {
		if err2 := ValidateSelfTResponseBody(body.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSelfTResponseBody runs the validations defined on SelfTResponseBody
func ValidateSelfTResponseBody(body *SelfTResponseBody) (err error) {
	if body.DescribedBy != nil {
		if err2 := ValidateDescribedByTResponseBody(body.DescribedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDescribedByTResponseBody runs the validations defined on
// DescribedByTResponseBody
func ValidateDescribedByTResponseBody(body *DescribedByTResponseBody) (err error) {
	if body.Href != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.href", *body.Href, goa.FormatURI))
	}
	return
}

// ValidateNavTResponseBody runs the validations defined on NavTResponseBody
func ValidateNavTResponseBody(body *NavTResponseBody) (err error) {
	if body.Self != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.self", *body.Self, goa.FormatURI))
	}
	if body.First != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.first", *body.First, goa.FormatURI))
	}
	if body.Next != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.next", *body.Next, goa.FormatURI))
	}
	return
}

// ValidateRefTResponseBody runs the validations defined on RefTResponseBody
func ValidateRefTResponseBody(body *RefTResponseBody) (err error) {
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Links != nil {
		if err2 := ValidateSelfTResponseBody(body.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
