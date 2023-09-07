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
	metadata "github.com/reinventingscience/ivcap-core-api/gen/metadata"
	metadataviews "github.com/reinventingscience/ivcap-core-api/gen/metadata/views"

	goa "goa.design/goa/v3/pkg"
)

// ReadResponseBody is the type of the "metadata" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// Record ID
	RecordID *string `form:"record-id,omitempty" json:"record-id,omitempty" xml:"record-id,omitempty"`
	// Entity ID
	Entity *string `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
	// Schema ID
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Attached metadata aspect
	Aspect interface{} `form:"aspect,omitempty" json:"aspect,omitempty" xml:"aspect,omitempty"`
	// Time this record was asserted
	ValidFrom *string `form:"valid-from,omitempty" json:"valid-from,omitempty" xml:"valid-from,omitempty"`
	// Time this record was revoked
	ValidTo *string `form:"valid-to,omitempty" json:"valid-to,omitempty" xml:"valid-to,omitempty"`
	// Entity asserting this metadata record at 'valid-from'
	Asserter *string `form:"asserter,omitempty" json:"asserter,omitempty" xml:"asserter,omitempty"`
	// Entity revoking this record at 'valid-to'
	Revoker *string `form:"revoker,omitempty" json:"revoker,omitempty" xml:"revoker,omitempty"`
}

// ListResponseBody is the type of the "metadata" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of metadata records
	Records []*MetadataListItemRTResponseBody `form:"records,omitempty" json:"records,omitempty" xml:"records,omitempty"`
	// Entity for which to request metadata
	EntityID *string `form:"entity-id,omitempty" json:"entity-id,omitempty" xml:"entity-id,omitempty"`
	// Optional schema to filter on
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Optional json path to further filter on returned list
	AspectPath *string `form:"aspect-path,omitempty" json:"aspect-path,omitempty" xml:"aspect-path,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	// Navigation links
	Links *NavTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// AddResponseBody is the type of the "metadata" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// Reference to record created
	RecordID *string `form:"record-id,omitempty" json:"record-id,omitempty" xml:"record-id,omitempty"`
}

// UpdateOneResponseBody is the type of the "metadata" service "update_one"
// endpoint HTTP response body.
type UpdateOneResponseBody struct {
	// Reference to record created
	RecordID *string `form:"record-id,omitempty" json:"record-id,omitempty" xml:"record-id,omitempty"`
}

// UpdateRecordResponseBody is the type of the "metadata" service
// "update_record" endpoint HTTP response body.
type UpdateRecordResponseBody struct {
	// Reference to record created
	RecordID *string `form:"record-id,omitempty" json:"record-id,omitempty" xml:"record-id,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "metadata" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "metadata" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "metadata" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "metadata" service "read"
// endpoint HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListBadRequestResponseBody is the type of the "metadata" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "metadata" service
// "list" endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "metadata" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "metadata" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// AddBadRequestResponseBody is the type of the "metadata" service "add"
// endpoint HTTP response body for the "bad-request" error.
type AddBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// AddInvalidParameterResponseBody is the type of the "metadata" service "add"
// endpoint HTTP response body for the "invalid-parameter" error.
type AddInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// AddInvalidScopesResponseBody is the type of the "metadata" service "add"
// endpoint HTTP response body for the "invalid-scopes" error.
type AddInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// AddNotImplementedResponseBody is the type of the "metadata" service "add"
// endpoint HTTP response body for the "not-implemented" error.
type AddNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateOneBadRequestResponseBody is the type of the "metadata" service
// "update_one" endpoint HTTP response body for the "bad-request" error.
type UpdateOneBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateOneInvalidParameterResponseBody is the type of the "metadata" service
// "update_one" endpoint HTTP response body for the "invalid-parameter" error.
type UpdateOneInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// UpdateOneInvalidScopesResponseBody is the type of the "metadata" service
// "update_one" endpoint HTTP response body for the "invalid-scopes" error.
type UpdateOneInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateOneNotImplementedResponseBody is the type of the "metadata" service
// "update_one" endpoint HTTP response body for the "not-implemented" error.
type UpdateOneNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateRecordBadRequestResponseBody is the type of the "metadata" service
// "update_record" endpoint HTTP response body for the "bad-request" error.
type UpdateRecordBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateRecordInvalidParameterResponseBody is the type of the "metadata"
// service "update_record" endpoint HTTP response body for the
// "invalid-parameter" error.
type UpdateRecordInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// UpdateRecordInvalidScopesResponseBody is the type of the "metadata" service
// "update_record" endpoint HTTP response body for the "invalid-scopes" error.
type UpdateRecordInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateRecordNotImplementedResponseBody is the type of the "metadata" service
// "update_record" endpoint HTTP response body for the "not-implemented" error.
type UpdateRecordNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RevokeBadRequestResponseBody is the type of the "metadata" service "revoke"
// endpoint HTTP response body for the "bad-request" error.
type RevokeBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RevokeInvalidParameterResponseBody is the type of the "metadata" service
// "revoke" endpoint HTTP response body for the "invalid-parameter" error.
type RevokeInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// RevokeInvalidScopesResponseBody is the type of the "metadata" service
// "revoke" endpoint HTTP response body for the "invalid-scopes" error.
type RevokeInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RevokeNotImplementedResponseBody is the type of the "metadata" service
// "revoke" endpoint HTTP response body for the "not-implemented" error.
type RevokeNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// MetadataListItemRTResponseBody is used to define fields on response body
// types.
type MetadataListItemRTResponseBody struct {
	// Record ID
	RecordID *string `form:"record-id,omitempty" json:"record-id,omitempty" xml:"record-id,omitempty"`
	// Entity ID
	Entity *string `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
	// Schema ID
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Attached metadata aspect
	Aspect interface{} `form:"aspect,omitempty" json:"aspect,omitempty" xml:"aspect,omitempty"`
	// If aspectPath was defined, this is what matched the query
	AspectContext *string `form:"aspectContext,omitempty" json:"aspectContext,omitempty" xml:"aspectContext,omitempty"`
}

// NavTResponseBody is used to define fields on response body types.
type NavTResponseBody struct {
	Self  *string `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	First *string `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Next  *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
}

// NewReadMetadataRecordRTOK builds a "metadata" service "read" endpoint result
// from a HTTP "OK" response.
func NewReadMetadataRecordRTOK(body *ReadResponseBody) *metadataviews.MetadataRecordRTView {
	v := &metadataviews.MetadataRecordRTView{
		RecordID:  body.RecordID,
		Entity:    body.Entity,
		Schema:    body.Schema,
		Aspect:    body.Aspect,
		ValidFrom: body.ValidFrom,
		ValidTo:   body.ValidTo,
		Asserter:  body.Asserter,
		Revoker:   body.Revoker,
	}

	return v
}

// NewReadBadRequest builds a metadata service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *metadata.BadRequestT {
	v := &metadata.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidCredential builds a metadata service read endpoint
// invalid-credential error.
func NewReadInvalidCredential() *metadata.InvalidCredentialsT {
	v := &metadata.InvalidCredentialsT{}

	return v
}

// NewReadInvalidScopes builds a metadata service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *metadata.InvalidScopesT {
	v := &metadata.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a metadata service read endpoint
// not-implemented error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *metadata.NotImplementedT {
	v := &metadata.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a metadata service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *metadata.ResourceNotFoundT {
	v := &metadata.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAuthorized builds a metadata service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *metadata.UnauthorizedT {
	v := &metadata.UnauthorizedT{}

	return v
}

// NewListMetaRTViewOK builds a "metadata" service "list" endpoint result from
// a HTTP "OK" response.
func NewListMetaRTViewOK(body *ListResponseBody) *metadataviews.ListMetaRTView {
	v := &metadataviews.ListMetaRTView{
		EntityID:   body.EntityID,
		Schema:     body.Schema,
		AspectPath: body.AspectPath,
		AtTime:     body.AtTime,
	}
	v.Records = make([]*metadataviews.MetadataListItemRTView, len(body.Records))
	for i, val := range body.Records {
		v.Records[i] = unmarshalMetadataListItemRTResponseBodyToMetadataviewsMetadataListItemRTView(val)
	}
	v.Links = unmarshalNavTResponseBodyToMetadataviewsNavTView(body.Links)

	return v
}

// NewListBadRequest builds a metadata service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *metadata.BadRequestT {
	v := &metadata.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidCredential builds a metadata service list endpoint
// invalid-credential error.
func NewListInvalidCredential() *metadata.InvalidCredentialsT {
	v := &metadata.InvalidCredentialsT{}

	return v
}

// NewListInvalidParameter builds a metadata service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *metadata.InvalidParameterValue {
	v := &metadata.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a metadata service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *metadata.InvalidScopesT {
	v := &metadata.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a metadata service list endpoint
// not-implemented error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *metadata.NotImplementedT {
	v := &metadata.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAuthorized builds a metadata service list endpoint not-authorized
// error.
func NewListNotAuthorized() *metadata.UnauthorizedT {
	v := &metadata.UnauthorizedT{}

	return v
}

// NewAddMetaRTViewOK builds a "metadata" service "add" endpoint result from a
// HTTP "OK" response.
func NewAddMetaRTViewOK(body *AddResponseBody) *metadataviews.AddMetaRTView {
	v := &metadataviews.AddMetaRTView{
		RecordID: body.RecordID,
	}

	return v
}

// NewAddBadRequest builds a metadata service add endpoint bad-request error.
func NewAddBadRequest(body *AddBadRequestResponseBody) *metadata.BadRequestT {
	v := &metadata.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewAddInvalidCredential builds a metadata service add endpoint
// invalid-credential error.
func NewAddInvalidCredential() *metadata.InvalidCredentialsT {
	v := &metadata.InvalidCredentialsT{}

	return v
}

// NewAddInvalidParameter builds a metadata service add endpoint
// invalid-parameter error.
func NewAddInvalidParameter(body *AddInvalidParameterResponseBody) *metadata.InvalidParameterValue {
	v := &metadata.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewAddInvalidScopes builds a metadata service add endpoint invalid-scopes
// error.
func NewAddInvalidScopes(body *AddInvalidScopesResponseBody) *metadata.InvalidScopesT {
	v := &metadata.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewAddNotImplemented builds a metadata service add endpoint not-implemented
// error.
func NewAddNotImplemented(body *AddNotImplementedResponseBody) *metadata.NotImplementedT {
	v := &metadata.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewAddNotAuthorized builds a metadata service add endpoint not-authorized
// error.
func NewAddNotAuthorized() *metadata.UnauthorizedT {
	v := &metadata.UnauthorizedT{}

	return v
}

// NewUpdateOneAddMetaRTOK builds a "metadata" service "update_one" endpoint
// result from a HTTP "OK" response.
func NewUpdateOneAddMetaRTOK(body *UpdateOneResponseBody) *metadataviews.AddMetaRTView {
	v := &metadataviews.AddMetaRTView{
		RecordID: body.RecordID,
	}

	return v
}

// NewUpdateOneBadRequest builds a metadata service update_one endpoint
// bad-request error.
func NewUpdateOneBadRequest(body *UpdateOneBadRequestResponseBody) *metadata.BadRequestT {
	v := &metadata.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateOneInvalidCredential builds a metadata service update_one endpoint
// invalid-credential error.
func NewUpdateOneInvalidCredential() *metadata.InvalidCredentialsT {
	v := &metadata.InvalidCredentialsT{}

	return v
}

// NewUpdateOneInvalidParameter builds a metadata service update_one endpoint
// invalid-parameter error.
func NewUpdateOneInvalidParameter(body *UpdateOneInvalidParameterResponseBody) *metadata.InvalidParameterValue {
	v := &metadata.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewUpdateOneInvalidScopes builds a metadata service update_one endpoint
// invalid-scopes error.
func NewUpdateOneInvalidScopes(body *UpdateOneInvalidScopesResponseBody) *metadata.InvalidScopesT {
	v := &metadata.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateOneNotImplemented builds a metadata service update_one endpoint
// not-implemented error.
func NewUpdateOneNotImplemented(body *UpdateOneNotImplementedResponseBody) *metadata.NotImplementedT {
	v := &metadata.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateOneNotAuthorized builds a metadata service update_one endpoint
// not-authorized error.
func NewUpdateOneNotAuthorized() *metadata.UnauthorizedT {
	v := &metadata.UnauthorizedT{}

	return v
}

// NewUpdateRecordAddMetaRTOK builds a "metadata" service "update_record"
// endpoint result from a HTTP "OK" response.
func NewUpdateRecordAddMetaRTOK(body *UpdateRecordResponseBody) *metadataviews.AddMetaRTView {
	v := &metadataviews.AddMetaRTView{
		RecordID: body.RecordID,
	}

	return v
}

// NewUpdateRecordBadRequest builds a metadata service update_record endpoint
// bad-request error.
func NewUpdateRecordBadRequest(body *UpdateRecordBadRequestResponseBody) *metadata.BadRequestT {
	v := &metadata.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateRecordInvalidCredential builds a metadata service update_record
// endpoint invalid-credential error.
func NewUpdateRecordInvalidCredential() *metadata.InvalidCredentialsT {
	v := &metadata.InvalidCredentialsT{}

	return v
}

// NewUpdateRecordInvalidParameter builds a metadata service update_record
// endpoint invalid-parameter error.
func NewUpdateRecordInvalidParameter(body *UpdateRecordInvalidParameterResponseBody) *metadata.InvalidParameterValue {
	v := &metadata.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewUpdateRecordInvalidScopes builds a metadata service update_record
// endpoint invalid-scopes error.
func NewUpdateRecordInvalidScopes(body *UpdateRecordInvalidScopesResponseBody) *metadata.InvalidScopesT {
	v := &metadata.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateRecordNotImplemented builds a metadata service update_record
// endpoint not-implemented error.
func NewUpdateRecordNotImplemented(body *UpdateRecordNotImplementedResponseBody) *metadata.NotImplementedT {
	v := &metadata.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateRecordNotAuthorized builds a metadata service update_record
// endpoint not-authorized error.
func NewUpdateRecordNotAuthorized() *metadata.UnauthorizedT {
	v := &metadata.UnauthorizedT{}

	return v
}

// NewRevokeBadRequest builds a metadata service revoke endpoint bad-request
// error.
func NewRevokeBadRequest(body *RevokeBadRequestResponseBody) *metadata.BadRequestT {
	v := &metadata.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewRevokeInvalidCredential builds a metadata service revoke endpoint
// invalid-credential error.
func NewRevokeInvalidCredential() *metadata.InvalidCredentialsT {
	v := &metadata.InvalidCredentialsT{}

	return v
}

// NewRevokeInvalidParameter builds a metadata service revoke endpoint
// invalid-parameter error.
func NewRevokeInvalidParameter(body *RevokeInvalidParameterResponseBody) *metadata.InvalidParameterValue {
	v := &metadata.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewRevokeInvalidScopes builds a metadata service revoke endpoint
// invalid-scopes error.
func NewRevokeInvalidScopes(body *RevokeInvalidScopesResponseBody) *metadata.InvalidScopesT {
	v := &metadata.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewRevokeNotImplemented builds a metadata service revoke endpoint
// not-implemented error.
func NewRevokeNotImplemented(body *RevokeNotImplementedResponseBody) *metadata.NotImplementedT {
	v := &metadata.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewRevokeNotAuthorized builds a metadata service revoke endpoint
// not-authorized error.
func NewRevokeNotAuthorized() *metadata.UnauthorizedT {
	v := &metadata.UnauthorizedT{}

	return v
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

// ValidateAddBadRequestResponseBody runs the validations defined on
// add_bad-request_response_body
func ValidateAddBadRequestResponseBody(body *AddBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateAddInvalidParameterResponseBody runs the validations defined on
// add_invalid-parameter_response_body
func ValidateAddInvalidParameterResponseBody(body *AddInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateAddInvalidScopesResponseBody runs the validations defined on
// add_invalid-scopes_response_body
func ValidateAddInvalidScopesResponseBody(body *AddInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateAddNotImplementedResponseBody runs the validations defined on
// add_not-implemented_response_body
func ValidateAddNotImplementedResponseBody(body *AddNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateOneBadRequestResponseBody runs the validations defined on
// update_one_bad-request_response_body
func ValidateUpdateOneBadRequestResponseBody(body *UpdateOneBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateOneInvalidParameterResponseBody runs the validations defined
// on update_one_invalid-parameter_response_body
func ValidateUpdateOneInvalidParameterResponseBody(body *UpdateOneInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateOneInvalidScopesResponseBody runs the validations defined on
// update_one_invalid-scopes_response_body
func ValidateUpdateOneInvalidScopesResponseBody(body *UpdateOneInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateOneNotImplementedResponseBody runs the validations defined on
// update_one_not-implemented_response_body
func ValidateUpdateOneNotImplementedResponseBody(body *UpdateOneNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateRecordBadRequestResponseBody runs the validations defined on
// update_record_bad-request_response_body
func ValidateUpdateRecordBadRequestResponseBody(body *UpdateRecordBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateRecordInvalidParameterResponseBody runs the validations
// defined on update_record_invalid-parameter_response_body
func ValidateUpdateRecordInvalidParameterResponseBody(body *UpdateRecordInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateRecordInvalidScopesResponseBody runs the validations defined
// on update_record_invalid-scopes_response_body
func ValidateUpdateRecordInvalidScopesResponseBody(body *UpdateRecordInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateRecordNotImplementedResponseBody runs the validations defined
// on update_record_not-implemented_response_body
func ValidateUpdateRecordNotImplementedResponseBody(body *UpdateRecordNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRevokeBadRequestResponseBody runs the validations defined on
// revoke_bad-request_response_body
func ValidateRevokeBadRequestResponseBody(body *RevokeBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRevokeInvalidParameterResponseBody runs the validations defined on
// revoke_invalid-parameter_response_body
func ValidateRevokeInvalidParameterResponseBody(body *RevokeInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRevokeInvalidScopesResponseBody runs the validations defined on
// revoke_invalid-scopes_response_body
func ValidateRevokeInvalidScopesResponseBody(body *RevokeInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateRevokeNotImplementedResponseBody runs the validations defined on
// revoke_not-implemented_response_body
func ValidateRevokeNotImplementedResponseBody(body *RevokeNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateMetadataListItemRTResponseBody runs the validations defined on
// MetadataListItemRTResponseBody
func ValidateMetadataListItemRTResponseBody(body *MetadataListItemRTResponseBody) (err error) {
	if body.RecordID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.record-id", *body.RecordID, goa.FormatURI))
	}
	if body.Entity != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.entity", *body.Entity, goa.FormatURI))
	}
	if body.Schema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.schema", *body.Schema, goa.FormatURI))
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
