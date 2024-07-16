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
	package_ "github.com/ivcap-works/ivcap-core-api/gen/package_"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "package" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// docker image tags
	Items []string             `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	Links []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// PushResponseBody is the type of the "package" service "push" endpoint HTTP
// response body.
type PushResponseBody struct {
	// digest or tag of image or layer
	Digest *string `form:"digest,omitempty" json:"digest,omitempty" xml:"digest,omitempty"`
	// location url for patch
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// layer mounted or not
	Mounted *bool `form:"mounted,omitempty" json:"mounted,omitempty" xml:"mounted,omitempty"`
}

// PatchResponseBody is the type of the "package" service "patch" endpoint HTTP
// response body.
type PatchResponseBody struct {
	// location url for patch
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
}

// PutResponseBody is the type of the "package" service "put" endpoint HTTP
// response body.
type PutResponseBody struct {
	// uploaded image digest or tag
	Digest *string `form:"digest,omitempty" json:"digest,omitempty" xml:"digest,omitempty"`
}

// ListBadRequestResponseBody is the type of the "package" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "package" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "package" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "package" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PullBadRequestResponseBody is the type of the "package" service "pull"
// endpoint HTTP response body for the "bad-request" error.
type PullBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PullInvalidParameterResponseBody is the type of the "package" service "pull"
// endpoint HTTP response body for the "invalid-parameter" error.
type PullInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// PullInvalidScopesResponseBody is the type of the "package" service "pull"
// endpoint HTTP response body for the "invalid-scopes" error.
type PullInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PullNotImplementedResponseBody is the type of the "package" service "pull"
// endpoint HTTP response body for the "not-implemented" error.
type PullNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PushBadRequestResponseBody is the type of the "package" service "push"
// endpoint HTTP response body for the "bad-request" error.
type PushBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PushInvalidParameterResponseBody is the type of the "package" service "push"
// endpoint HTTP response body for the "invalid-parameter" error.
type PushInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// PushInvalidScopesResponseBody is the type of the "package" service "push"
// endpoint HTTP response body for the "invalid-scopes" error.
type PushInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PushNotImplementedResponseBody is the type of the "package" service "push"
// endpoint HTTP response body for the "not-implemented" error.
type PushNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PushAlreadyCreatedResponseBody is the type of the "package" service "push"
// endpoint HTTP response body for the "already-created" error.
type PushAlreadyCreatedResponseBody struct {
	// ID of already existing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PatchBadRequestResponseBody is the type of the "package" service "patch"
// endpoint HTTP response body for the "bad-request" error.
type PatchBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PatchInvalidParameterResponseBody is the type of the "package" service
// "patch" endpoint HTTP response body for the "invalid-parameter" error.
type PatchInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// PatchInvalidScopesResponseBody is the type of the "package" service "patch"
// endpoint HTTP response body for the "invalid-scopes" error.
type PatchInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PatchNotImplementedResponseBody is the type of the "package" service "patch"
// endpoint HTTP response body for the "not-implemented" error.
type PatchNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PatchAlreadyCreatedResponseBody is the type of the "package" service "patch"
// endpoint HTTP response body for the "already-created" error.
type PatchAlreadyCreatedResponseBody struct {
	// ID of already existing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PutBadRequestResponseBody is the type of the "package" service "put"
// endpoint HTTP response body for the "bad-request" error.
type PutBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PutInvalidParameterResponseBody is the type of the "package" service "put"
// endpoint HTTP response body for the "invalid-parameter" error.
type PutInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// PutInvalidScopesResponseBody is the type of the "package" service "put"
// endpoint HTTP response body for the "invalid-scopes" error.
type PutInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PutNotImplementedResponseBody is the type of the "package" service "put"
// endpoint HTTP response body for the "not-implemented" error.
type PutNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PutAlreadyCreatedResponseBody is the type of the "package" service "put"
// endpoint HTTP response body for the "already-created" error.
type PutAlreadyCreatedResponseBody struct {
	// ID of already existing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveBadRequestResponseBody is the type of the "package" service "remove"
// endpoint HTTP response body for the "bad-request" error.
type RemoveBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveInvalidParameterResponseBody is the type of the "package" service
// "remove" endpoint HTTP response body for the "invalid-parameter" error.
type RemoveInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// RemoveInvalidScopesResponseBody is the type of the "package" service
// "remove" endpoint HTTP response body for the "invalid-scopes" error.
type RemoveInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveNotImplementedResponseBody is the type of the "package" service
// "remove" endpoint HTTP response body for the "not-implemented" error.
type RemoveNotImplementedResponseBody struct {
	// Information message
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

// NewListResultOK builds a "package" service "list" endpoint result from a
// HTTP "OK" response.
func NewListResultOK(body *ListResponseBody) *package_.ListResult {
	v := &package_.ListResult{}
	v.Items = make([]string, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = val
	}
	v.Links = make([]*package_.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToPackageLinkT(val)
	}

	return v
}

// NewListBadRequest builds a package service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *package_.BadRequestT {
	v := &package_.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidParameter builds a package service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *package_.InvalidParameterT {
	v := &package_.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a package service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *package_.InvalidScopesT {
	v := &package_.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a package service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *package_.NotImplementedT {
	v := &package_.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAvailable builds a package service list endpoint not-available
// error.
func NewListNotAvailable() *package_.ServiceNotAvailableT {
	v := &package_.ServiceNotAvailableT{}

	return v
}

// NewListNotAuthorized builds a package service list endpoint not-authorized
// error.
func NewListNotAuthorized() *package_.UnauthorizedT {
	v := &package_.UnauthorizedT{}

	return v
}

// NewPullResultTOK builds a "package" service "pull" endpoint result from a
// HTTP "OK" response.
func NewPullResultTOK(total int, available int) *package_.PullResultT {
	v := &package_.PullResultT{}
	v.Total = total
	v.Available = available

	return v
}

// NewPullBadRequest builds a package service pull endpoint bad-request error.
func NewPullBadRequest(body *PullBadRequestResponseBody) *package_.BadRequestT {
	v := &package_.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewPullInvalidParameter builds a package service pull endpoint
// invalid-parameter error.
func NewPullInvalidParameter(body *PullInvalidParameterResponseBody) *package_.InvalidParameterT {
	v := &package_.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewPullInvalidScopes builds a package service pull endpoint invalid-scopes
// error.
func NewPullInvalidScopes(body *PullInvalidScopesResponseBody) *package_.InvalidScopesT {
	v := &package_.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPullNotImplemented builds a package service pull endpoint not-implemented
// error.
func NewPullNotImplemented(body *PullNotImplementedResponseBody) *package_.NotImplementedT {
	v := &package_.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewPullNotAvailable builds a package service pull endpoint not-available
// error.
func NewPullNotAvailable() *package_.ServiceNotAvailableT {
	v := &package_.ServiceNotAvailableT{}

	return v
}

// NewPullNotAuthorized builds a package service pull endpoint not-authorized
// error.
func NewPullNotAuthorized() *package_.UnauthorizedT {
	v := &package_.UnauthorizedT{}

	return v
}

// NewPushResultCreated builds a "package" service "push" endpoint result from
// a HTTP "Created" response.
func NewPushResultCreated(body *PushResponseBody) *package_.PushResult {
	v := &package_.PushResult{
		Digest:   *body.Digest,
		Location: *body.Location,
		Mounted:  *body.Mounted,
	}

	return v
}

// NewPushBadRequest builds a package service push endpoint bad-request error.
func NewPushBadRequest(body *PushBadRequestResponseBody) *package_.BadRequestT {
	v := &package_.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewPushInvalidParameter builds a package service push endpoint
// invalid-parameter error.
func NewPushInvalidParameter(body *PushInvalidParameterResponseBody) *package_.InvalidParameterT {
	v := &package_.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewPushInvalidScopes builds a package service push endpoint invalid-scopes
// error.
func NewPushInvalidScopes(body *PushInvalidScopesResponseBody) *package_.InvalidScopesT {
	v := &package_.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPushNotImplemented builds a package service push endpoint not-implemented
// error.
func NewPushNotImplemented(body *PushNotImplementedResponseBody) *package_.NotImplementedT {
	v := &package_.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewPushAlreadyCreated builds a package service push endpoint already-created
// error.
func NewPushAlreadyCreated(body *PushAlreadyCreatedResponseBody) *package_.ResourceAlreadyCreatedT {
	v := &package_.ResourceAlreadyCreatedT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPushNotAvailable builds a package service push endpoint not-available
// error.
func NewPushNotAvailable() *package_.ServiceNotAvailableT {
	v := &package_.ServiceNotAvailableT{}

	return v
}

// NewPushNotAuthorized builds a package service push endpoint not-authorized
// error.
func NewPushNotAuthorized() *package_.UnauthorizedT {
	v := &package_.UnauthorizedT{}

	return v
}

// NewPatchResultCreated builds a "package" service "patch" endpoint result
// from a HTTP "Created" response.
func NewPatchResultCreated(body *PatchResponseBody) *package_.PatchResult {
	v := &package_.PatchResult{
		Location: *body.Location,
	}

	return v
}

// NewPatchBadRequest builds a package service patch endpoint bad-request error.
func NewPatchBadRequest(body *PatchBadRequestResponseBody) *package_.BadRequestT {
	v := &package_.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewPatchInvalidParameter builds a package service patch endpoint
// invalid-parameter error.
func NewPatchInvalidParameter(body *PatchInvalidParameterResponseBody) *package_.InvalidParameterT {
	v := &package_.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewPatchInvalidScopes builds a package service patch endpoint invalid-scopes
// error.
func NewPatchInvalidScopes(body *PatchInvalidScopesResponseBody) *package_.InvalidScopesT {
	v := &package_.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPatchNotImplemented builds a package service patch endpoint
// not-implemented error.
func NewPatchNotImplemented(body *PatchNotImplementedResponseBody) *package_.NotImplementedT {
	v := &package_.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewPatchAlreadyCreated builds a package service patch endpoint
// already-created error.
func NewPatchAlreadyCreated(body *PatchAlreadyCreatedResponseBody) *package_.ResourceAlreadyCreatedT {
	v := &package_.ResourceAlreadyCreatedT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPatchNotAvailable builds a package service patch endpoint not-available
// error.
func NewPatchNotAvailable() *package_.ServiceNotAvailableT {
	v := &package_.ServiceNotAvailableT{}

	return v
}

// NewPatchNotAuthorized builds a package service patch endpoint not-authorized
// error.
func NewPatchNotAuthorized() *package_.UnauthorizedT {
	v := &package_.UnauthorizedT{}

	return v
}

// NewPutResultCreated builds a "package" service "put" endpoint result from a
// HTTP "Created" response.
func NewPutResultCreated(body *PutResponseBody) *package_.PutResult {
	v := &package_.PutResult{
		Digest: *body.Digest,
	}

	return v
}

// NewPutBadRequest builds a package service put endpoint bad-request error.
func NewPutBadRequest(body *PutBadRequestResponseBody) *package_.BadRequestT {
	v := &package_.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewPutInvalidParameter builds a package service put endpoint
// invalid-parameter error.
func NewPutInvalidParameter(body *PutInvalidParameterResponseBody) *package_.InvalidParameterT {
	v := &package_.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewPutInvalidScopes builds a package service put endpoint invalid-scopes
// error.
func NewPutInvalidScopes(body *PutInvalidScopesResponseBody) *package_.InvalidScopesT {
	v := &package_.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPutNotImplemented builds a package service put endpoint not-implemented
// error.
func NewPutNotImplemented(body *PutNotImplementedResponseBody) *package_.NotImplementedT {
	v := &package_.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewPutAlreadyCreated builds a package service put endpoint already-created
// error.
func NewPutAlreadyCreated(body *PutAlreadyCreatedResponseBody) *package_.ResourceAlreadyCreatedT {
	v := &package_.ResourceAlreadyCreatedT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewPutNotAvailable builds a package service put endpoint not-available error.
func NewPutNotAvailable() *package_.ServiceNotAvailableT {
	v := &package_.ServiceNotAvailableT{}

	return v
}

// NewPutNotAuthorized builds a package service put endpoint not-authorized
// error.
func NewPutNotAuthorized() *package_.UnauthorizedT {
	v := &package_.UnauthorizedT{}

	return v
}

// NewRemoveBadRequest builds a package service remove endpoint bad-request
// error.
func NewRemoveBadRequest(body *RemoveBadRequestResponseBody) *package_.BadRequestT {
	v := &package_.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewRemoveInvalidParameter builds a package service remove endpoint
// invalid-parameter error.
func NewRemoveInvalidParameter(body *RemoveInvalidParameterResponseBody) *package_.InvalidParameterT {
	v := &package_.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewRemoveInvalidScopes builds a package service remove endpoint
// invalid-scopes error.
func NewRemoveInvalidScopes(body *RemoveInvalidScopesResponseBody) *package_.InvalidScopesT {
	v := &package_.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewRemoveNotImplemented builds a package service remove endpoint
// not-implemented error.
func NewRemoveNotImplemented(body *RemoveNotImplementedResponseBody) *package_.NotImplementedT {
	v := &package_.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewRemoveNotAvailable builds a package service remove endpoint not-available
// error.
func NewRemoveNotAvailable() *package_.ServiceNotAvailableT {
	v := &package_.ServiceNotAvailableT{}

	return v
}

// NewRemoveNotAuthorized builds a package service remove endpoint
// not-authorized error.
func NewRemoveNotAuthorized() *package_.UnauthorizedT {
	v := &package_.UnauthorizedT{}

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
	for _, e := range body.Links {
		if e != nil {
			if err2 := ValidateLinkTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidatePushResponseBody runs the validations defined on PushResponseBody
func ValidatePushResponseBody(body *PushResponseBody) (err error) {
	if body.Digest == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("digest", "body"))
	}
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	if body.Mounted == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mounted", "body"))
	}
	return
}

// ValidatePatchResponseBody runs the validations defined on PatchResponseBody
func ValidatePatchResponseBody(body *PatchResponseBody) (err error) {
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	return
}

// ValidatePutResponseBody runs the validations defined on PutResponseBody
func ValidatePutResponseBody(body *PutResponseBody) (err error) {
	if body.Digest == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("digest", "body"))
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

// ValidatePullBadRequestResponseBody runs the validations defined on
// pull_bad-request_response_body
func ValidatePullBadRequestResponseBody(body *PullBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePullInvalidParameterResponseBody runs the validations defined on
// pull_invalid-parameter_response_body
func ValidatePullInvalidParameterResponseBody(body *PullInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePullInvalidScopesResponseBody runs the validations defined on
// pull_invalid-scopes_response_body
func ValidatePullInvalidScopesResponseBody(body *PullInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidatePullNotImplementedResponseBody runs the validations defined on
// pull_not-implemented_response_body
func ValidatePullNotImplementedResponseBody(body *PullNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePushBadRequestResponseBody runs the validations defined on
// push_bad-request_response_body
func ValidatePushBadRequestResponseBody(body *PushBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePushInvalidParameterResponseBody runs the validations defined on
// push_invalid-parameter_response_body
func ValidatePushInvalidParameterResponseBody(body *PushInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePushInvalidScopesResponseBody runs the validations defined on
// push_invalid-scopes_response_body
func ValidatePushInvalidScopesResponseBody(body *PushInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidatePushNotImplementedResponseBody runs the validations defined on
// push_not-implemented_response_body
func ValidatePushNotImplementedResponseBody(body *PushNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePushAlreadyCreatedResponseBody runs the validations defined on
// push_already-created_response_body
func ValidatePushAlreadyCreatedResponseBody(body *PushAlreadyCreatedResponseBody) (err error) {
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

// ValidatePatchBadRequestResponseBody runs the validations defined on
// patch_bad-request_response_body
func ValidatePatchBadRequestResponseBody(body *PatchBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePatchInvalidParameterResponseBody runs the validations defined on
// patch_invalid-parameter_response_body
func ValidatePatchInvalidParameterResponseBody(body *PatchInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePatchInvalidScopesResponseBody runs the validations defined on
// patch_invalid-scopes_response_body
func ValidatePatchInvalidScopesResponseBody(body *PatchInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidatePatchNotImplementedResponseBody runs the validations defined on
// patch_not-implemented_response_body
func ValidatePatchNotImplementedResponseBody(body *PatchNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePatchAlreadyCreatedResponseBody runs the validations defined on
// patch_already-created_response_body
func ValidatePatchAlreadyCreatedResponseBody(body *PatchAlreadyCreatedResponseBody) (err error) {
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

// ValidatePutBadRequestResponseBody runs the validations defined on
// put_bad-request_response_body
func ValidatePutBadRequestResponseBody(body *PutBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePutInvalidParameterResponseBody runs the validations defined on
// put_invalid-parameter_response_body
func ValidatePutInvalidParameterResponseBody(body *PutInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePutInvalidScopesResponseBody runs the validations defined on
// put_invalid-scopes_response_body
func ValidatePutInvalidScopesResponseBody(body *PutInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidatePutNotImplementedResponseBody runs the validations defined on
// put_not-implemented_response_body
func ValidatePutNotImplementedResponseBody(body *PutNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidatePutAlreadyCreatedResponseBody runs the validations defined on
// put_already-created_response_body
func ValidatePutAlreadyCreatedResponseBody(body *PutAlreadyCreatedResponseBody) (err error) {
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

// ValidateRemoveBadRequestResponseBody runs the validations defined on
// remove_bad-request_response_body
func ValidateRemoveBadRequestResponseBody(body *RemoveBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRemoveInvalidParameterResponseBody runs the validations defined on
// remove_invalid-parameter_response_body
func ValidateRemoveInvalidParameterResponseBody(body *RemoveInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRemoveInvalidScopesResponseBody runs the validations defined on
// remove_invalid-scopes_response_body
func ValidateRemoveInvalidScopesResponseBody(body *RemoveInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateRemoveNotImplementedResponseBody runs the validations defined on
// remove_not-implemented_response_body
func ValidateRemoveNotImplementedResponseBody(body *RemoveNotImplementedResponseBody) (err error) {
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
