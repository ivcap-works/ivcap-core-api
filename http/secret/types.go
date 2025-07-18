// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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
	secret "github.com/ivcap-works/ivcap-core-api/gen/secret"
	goa "goa.design/goa/v3/pkg"
)

// SetRequestBody is the type of the "secret" service "set" endpoint HTTP
// request body.
type SetRequestBody struct {
	// Secret name
	SecretName string `form:"secret-name" json:"secret-name" xml:"secret-name"`
	// Secret type
	SecretType *string `form:"secret-type,omitempty" json:"secret-type,omitempty" xml:"secret-type,omitempty"`
	// Secret value
	SecretValue string `form:"secret-value" json:"secret-value" xml:"secret-value"`
	// Expiry time
	ExpiryTime int64 `form:"expiry-time" json:"expiry-time" xml:"expiry-time"`
}

// ListResponseBody is the type of the "secret" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// secrets
	Items []*SecretListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	Links []*LinkTResponseBody          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// GetResponseBody is the type of the "secret" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Secret name
	SecretName *string `form:"secret-name,omitempty" json:"secret-name,omitempty" xml:"secret-name,omitempty"`
	// Secret value
	SecretValue *string `form:"secret-value,omitempty" json:"secret-value,omitempty" xml:"secret-value,omitempty"`
	// Expiry time
	ExpiryTime *int64 `form:"expiry-time,omitempty" json:"expiry-time,omitempty" xml:"expiry-time,omitempty"`
}

// ListBadRequestResponseBody is the type of the "secret" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "secret" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "secret" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "secret" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetBadRequestResponseBody is the type of the "secret" service "get" endpoint
// HTTP response body for the "bad-request" error.
type GetBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetInvalidParameterResponseBody is the type of the "secret" service "get"
// endpoint HTTP response body for the "invalid-parameter" error.
type GetInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// GetInvalidScopesResponseBody is the type of the "secret" service "get"
// endpoint HTTP response body for the "invalid-scopes" error.
type GetInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetNotImplementedResponseBody is the type of the "secret" service "get"
// endpoint HTTP response body for the "not-implemented" error.
type GetNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetNotFoundResponseBody is the type of the "secret" service "get" endpoint
// HTTP response body for the "not-found" error.
type GetNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetBadRequestResponseBody is the type of the "secret" service "set" endpoint
// HTTP response body for the "bad-request" error.
type SetBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetInvalidParameterResponseBody is the type of the "secret" service "set"
// endpoint HTTP response body for the "invalid-parameter" error.
type SetInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// SetInvalidScopesResponseBody is the type of the "secret" service "set"
// endpoint HTTP response body for the "invalid-scopes" error.
type SetInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetNotImplementedResponseBody is the type of the "secret" service "set"
// endpoint HTTP response body for the "not-implemented" error.
type SetNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetNotFoundResponseBody is the type of the "secret" service "set" endpoint
// HTTP response body for the "not-found" error.
type SetNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SecretListItemResponseBody is used to define fields on response body types.
type SecretListItemResponseBody struct {
	// Secret name
	SecretName *string `form:"secret-name,omitempty" json:"secret-name,omitempty" xml:"secret-name,omitempty"`
	// Expiry time
	ExpiryTime *int64 `form:"expiry-time,omitempty" json:"expiry-time,omitempty" xml:"expiry-time,omitempty"`
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

// NewSetRequestBody builds the HTTP request body from the payload of the "set"
// endpoint of the "secret" service.
func NewSetRequestBody(p *secret.SetPayload) *SetRequestBody {
	body := &SetRequestBody{
		SecretName:  p.Secrets.SecretName,
		SecretType:  p.Secrets.SecretType,
		SecretValue: p.Secrets.SecretValue,
		ExpiryTime:  p.Secrets.ExpiryTime,
	}
	return body
}

// NewListResultOK builds a "secret" service "list" endpoint result from a HTTP
// "OK" response.
func NewListResultOK(body *ListResponseBody) *secret.ListResult {
	v := &secret.ListResult{}
	v.Items = make([]*secret.SecretListItem, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalSecretListItemResponseBodyToSecretSecretListItem(val)
	}
	v.Links = make([]*secret.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToSecretLinkT(val)
	}

	return v
}

// NewListBadRequest builds a secret service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *secret.BadRequestT {
	v := &secret.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidParameter builds a secret service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *secret.InvalidParameterT {
	v := &secret.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a secret service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *secret.InvalidScopesT {
	v := &secret.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a secret service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *secret.NotImplementedT {
	v := &secret.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAvailable builds a secret service list endpoint not-available
// error.
func NewListNotAvailable() *secret.ServiceNotAvailableT {
	v := &secret.ServiceNotAvailableT{}

	return v
}

// NewListNotAuthorized builds a secret service list endpoint not-authorized
// error.
func NewListNotAuthorized() *secret.UnauthorizedT {
	v := &secret.UnauthorizedT{}

	return v
}

// NewGetSecretResultTOK builds a "secret" service "get" endpoint result from a
// HTTP "OK" response.
func NewGetSecretResultTOK(body *GetResponseBody) *secret.SecretResultT {
	v := &secret.SecretResultT{
		SecretName:  *body.SecretName,
		SecretValue: *body.SecretValue,
		ExpiryTime:  *body.ExpiryTime,
	}

	return v
}

// NewGetBadRequest builds a secret service get endpoint bad-request error.
func NewGetBadRequest(body *GetBadRequestResponseBody) *secret.BadRequestT {
	v := &secret.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewGetInvalidParameter builds a secret service get endpoint
// invalid-parameter error.
func NewGetInvalidParameter(body *GetInvalidParameterResponseBody) *secret.InvalidParameterT {
	v := &secret.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewGetInvalidScopes builds a secret service get endpoint invalid-scopes
// error.
func NewGetInvalidScopes(body *GetInvalidScopesResponseBody) *secret.InvalidScopesT {
	v := &secret.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewGetNotImplemented builds a secret service get endpoint not-implemented
// error.
func NewGetNotImplemented(body *GetNotImplementedResponseBody) *secret.NotImplementedT {
	v := &secret.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewGetNotFound builds a secret service get endpoint not-found error.
func NewGetNotFound(body *GetNotFoundResponseBody) *secret.ResourceNotFoundT {
	v := &secret.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewGetNotAvailable builds a secret service get endpoint not-available error.
func NewGetNotAvailable() *secret.ServiceNotAvailableT {
	v := &secret.ServiceNotAvailableT{}

	return v
}

// NewGetNotAuthorized builds a secret service get endpoint not-authorized
// error.
func NewGetNotAuthorized() *secret.UnauthorizedT {
	v := &secret.UnauthorizedT{}

	return v
}

// NewSetBadRequest builds a secret service set endpoint bad-request error.
func NewSetBadRequest(body *SetBadRequestResponseBody) *secret.BadRequestT {
	v := &secret.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewSetInvalidParameter builds a secret service set endpoint
// invalid-parameter error.
func NewSetInvalidParameter(body *SetInvalidParameterResponseBody) *secret.InvalidParameterT {
	v := &secret.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewSetInvalidScopes builds a secret service set endpoint invalid-scopes
// error.
func NewSetInvalidScopes(body *SetInvalidScopesResponseBody) *secret.InvalidScopesT {
	v := &secret.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSetNotImplemented builds a secret service set endpoint not-implemented
// error.
func NewSetNotImplemented(body *SetNotImplementedResponseBody) *secret.NotImplementedT {
	v := &secret.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewSetNotFound builds a secret service set endpoint not-found error.
func NewSetNotFound(body *SetNotFoundResponseBody) *secret.ResourceNotFoundT {
	v := &secret.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSetNotAvailable builds a secret service set endpoint not-available error.
func NewSetNotAvailable() *secret.ServiceNotAvailableT {
	v := &secret.ServiceNotAvailableT{}

	return v
}

// NewSetNotAuthorized builds a secret service set endpoint not-authorized
// error.
func NewSetNotAuthorized() *secret.UnauthorizedT {
	v := &secret.UnauthorizedT{}

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
			if err2 := ValidateSecretListItemResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
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

// ValidateGetResponseBody runs the validations defined on GetResponseBody
func ValidateGetResponseBody(body *GetResponseBody) (err error) {
	if body.SecretName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("secret-name", "body"))
	}
	if body.SecretValue == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("secret-value", "body"))
	}
	if body.ExpiryTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expiry-time", "body"))
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

// ValidateGetBadRequestResponseBody runs the validations defined on
// get_bad-request_response_body
func ValidateGetBadRequestResponseBody(body *GetBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetInvalidParameterResponseBody runs the validations defined on
// get_invalid-parameter_response_body
func ValidateGetInvalidParameterResponseBody(body *GetInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetInvalidScopesResponseBody runs the validations defined on
// get_invalid-scopes_response_body
func ValidateGetInvalidScopesResponseBody(body *GetInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateGetNotImplementedResponseBody runs the validations defined on
// get_not-implemented_response_body
func ValidateGetNotImplementedResponseBody(body *GetNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetNotFoundResponseBody runs the validations defined on
// get_not-found_response_body
func ValidateGetNotFoundResponseBody(body *GetNotFoundResponseBody) (err error) {
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

// ValidateSetBadRequestResponseBody runs the validations defined on
// set_bad-request_response_body
func ValidateSetBadRequestResponseBody(body *SetBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetInvalidParameterResponseBody runs the validations defined on
// set_invalid-parameter_response_body
func ValidateSetInvalidParameterResponseBody(body *SetInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetInvalidScopesResponseBody runs the validations defined on
// set_invalid-scopes_response_body
func ValidateSetInvalidScopesResponseBody(body *SetInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateSetNotImplementedResponseBody runs the validations defined on
// set_not-implemented_response_body
func ValidateSetNotImplementedResponseBody(body *SetNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetNotFoundResponseBody runs the validations defined on
// set_not-found_response_body
func ValidateSetNotFoundResponseBody(body *SetNotFoundResponseBody) (err error) {
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

// ValidateSecretListItemResponseBody runs the validations defined on
// SecretListItemResponseBody
func ValidateSecretListItemResponseBody(body *SecretListItemResponseBody) (err error) {
	if body.SecretName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("secret-name", "body"))
	}
	if body.ExpiryTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expiry-time", "body"))
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
