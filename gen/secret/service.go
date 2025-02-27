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

package secret

import (
	"context"

	"goa.design/goa/v3/security"
)

// List, set or get secrets
type Service interface {
	// list secrets under account
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Get a secrets
	Get(context.Context, *GetPayload) (res *SecretResultT, err error)
	// Set a secrets
	Set(context.Context, *SetPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ivcap"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.43"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "secret"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list", "get", "set"}

// Something wasn't right with this request
type BadRequestT struct {
	// Information message
	Message string
}

// GetPayload is the payload type of the secret service get method.
type GetPayload struct {
	// Secret name
	SecretName string
	// Secret type
	SecretType *string
	// JWT used for authentication
	JWT string
}

// InvalidParameterT is the error returned when a parameter has the wrong value.
type InvalidParameterT struct {
	// message describing expected type or pattern.
	Message string
	// name of parameter.
	Name string
	// provided parameter value.
	Value *string
}

// Caller not authorized to access required scope.
type InvalidScopesT struct {
	// ID of involved resource
	ID *string
	// Message of error
	Message string
}

type LinkT struct {
	// relation type
	Rel string
	// mime type
	Type string
	// web link
	Href string
}

// ListPayload is the payload type of the secret service list method.
type ListPayload struct {
	// page url to list
	Page *string
	// filter of name pattern
	Filter *string
	// maximum number of secrets
	Limit *int
	// offset token of secrets
	Offset *string
	// JWT used for authentication
	JWT string
}

// ListResult is the result type of the secret service list method.
type ListResult struct {
	// secrets
	Items []*SecretListItem
	Links []*LinkT
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

// NotFound is the type returned when attempting to manage a resource that does
// not exist.
type ResourceNotFoundT struct {
	// ID of missing resource
	ID string
	// Message of error
	Message string
}

type SecretListItem struct {
	// Secret name
	SecretName string
	// Expiry time
	ExpiryTime int64
}

// SecretResultT is the result type of the secret service get method.
type SecretResultT struct {
	// Secret name
	SecretName string
	// Secret value
	SecretValue string
	// Expiry time
	ExpiryTime int64
}

// Service necessary to fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// SetPayload is the payload type of the secret service set method.
type SetPayload struct {
	// New secrets description
	Secrets *SetSecretRequestT
	// JWT used for authentication
	JWT string
}

type SetSecretRequestT struct {
	// Secret name
	SecretName string
	// Secret type
	SecretType *string
	// Secret value
	SecretValue string
	// Expiry time
	ExpiryTime int64
}

// Unauthorized access to resource
type UnauthorizedT struct {
}

// Error returns an error description.
func (e *BadRequestT) Error() string {
	return "Something wasn't right with this request"
}

// ErrorName returns "BadRequestT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *BadRequestT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "BadRequestT".
func (e *BadRequestT) GoaErrorName() string {
	return "bad-request"
}

// Error returns an error description.
func (e *InvalidParameterT) Error() string {
	return "InvalidParameterT is the error returned when a parameter has the wrong value."
}

// ErrorName returns "InvalidParameterT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InvalidParameterT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InvalidParameterT".
func (e *InvalidParameterT) GoaErrorName() string {
	return "invalid-parameter"
}

// Error returns an error description.
func (e *InvalidScopesT) Error() string {
	return "Caller not authorized to access required scope."
}

// ErrorName returns "InvalidScopesT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InvalidScopesT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InvalidScopesT".
func (e *InvalidScopesT) GoaErrorName() string {
	return e.Message
}

// Error returns an error description.
func (e *NotImplementedT) Error() string {
	return "Method is not yet implemented."
}

// ErrorName returns "NotImplementedT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotImplementedT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotImplementedT".
func (e *NotImplementedT) GoaErrorName() string {
	return "not-implemented"
}

// Error returns an error description.
func (e *ResourceNotFoundT) Error() string {
	return "NotFound is the type returned when attempting to manage a resource that does not exist."
}

// ErrorName returns "ResourceNotFoundT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ResourceNotFoundT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ResourceNotFoundT".
func (e *ResourceNotFoundT) GoaErrorName() string {
	return "not-found"
}

// Error returns an error description.
func (e *ServiceNotAvailableT) Error() string {
	return "Service necessary to fulfil the request is currently not available."
}

// ErrorName returns "ServiceNotAvailableT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ServiceNotAvailableT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ServiceNotAvailableT".
func (e *ServiceNotAvailableT) GoaErrorName() string {
	return "not-available"
}

// Error returns an error description.
func (e *UnauthorizedT) Error() string {
	return "Unauthorized access to resource"
}

// ErrorName returns "UnauthorizedT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *UnauthorizedT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "UnauthorizedT".
func (e *UnauthorizedT) GoaErrorName() string {
	return "not-authorized"
}
