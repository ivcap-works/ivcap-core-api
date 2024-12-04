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

package package_

import (
	"context"
	"io"

	"goa.design/goa/v3/security"
)

// Manage the life cycle of a service package.
type Service interface {
	// list ivcap service's docker images under account
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// pull ivcap service's docker image
	Pull(context.Context, *PullPayload) (res *PullResultT, body io.ReadCloser, err error)
	// upload service's docker image to container registry
	Push(context.Context, *PushPayload, io.ReadCloser) (res *PushResult, err error)
	// check push status of a layer
	Status(context.Context, *StatusPayload) (res *PushStatusT, err error)
	// remove ivcap service's docker image
	Remove(context.Context, *RemovePayload) (err error)
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
const ServiceName = "package"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "pull", "push", "status", "remove"}

// Something wasn't right with this request
type BadRequestT struct {
	// Information message
	Message string
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

// ListPayload is the payload type of the package service list method.
type ListPayload struct {
	// docker image tag
	Tag *string
	// maximum number of repository items, which can have multiple tags
	Limit *int
	// page url to list
	Page *string
	// JWT used for authentication
	JWT string
}

// ListResult is the result type of the package service list method.
type ListResult struct {
	// docker image tags
	Items []string
	Links []*LinkT
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

// PullPayload is the payload type of the package service pull method.
type PullPayload struct {
	// docker image tag or layer digest
	Ref string
	// pull type, either be manifest, config or layer
	Type string
	// offset of the layer chunk
	Offset *int
	// JWT used for authentication
	JWT string
}

// PullResultT is the result type of the package service pull method.
type PullResultT struct {
	// total size in bytes of layer
	Total int
	// available size in bytes of layer to read
	Available int
}

// PushPayload is the payload type of the package service push method.
type PushPayload struct {
	// docker image tag
	Tag string
	// force to override
	Force *bool
	// push type, either be manifest, config or layer
	Type string
	// digest of the push
	Digest string
	// start of the layer chunk
	Start *int
	// end of the layer chunk
	End *int
	// total size of the layer
	Total *int
	// JWT used for authentication
	JWT string
}

// PushResult is the result type of the package service push method.
type PushResult struct {
	// uploaded image digest or tag
	Digest string
	// layer exists or not
	Exists bool
}

// PushStatusT is the result type of the package service status method.
type PushStatusT struct {
	// Push status
	Status string
	// Message
	Message string
}

// RemovePayload is the payload type of the package service remove method.
type RemovePayload struct {
	// docker image tag
	Tag string
	// JWT used for authentication
	JWT string
}

// Will be returned when receiving a request to create and already existing
// resource.
type ResourceAlreadyCreatedT struct {
	// ID of already existing resource
	ID string
	// Message of error
	Message string
}

// Service necessary to fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// StatusPayload is the payload type of the package service status method.
type StatusPayload struct {
	// docker image tag
	Tag string
	// docker image layer digest
	Digest string
	// JWT used for authentication
	JWT string
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
func (e *ResourceAlreadyCreatedT) Error() string {
	return "Will be returned when receiving a request to create and already existing resource."
}

// ErrorName returns "ResourceAlreadyCreatedT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ResourceAlreadyCreatedT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ResourceAlreadyCreatedT".
func (e *ResourceAlreadyCreatedT) GoaErrorName() string {
	return "already-created"
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
