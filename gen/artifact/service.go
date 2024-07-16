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

package artifact

import (
	"context"
	"io"

	"goa.design/goa/v3/security"
)

// Manage the life cycle of an artifact stored by this deployment.
type Service interface {
	// list artifacts
	List(context.Context, *ListPayload) (res *ArtifactListRT, err error)
	// Show artifacts by ID
	Read(context.Context, *ReadPayload) (res *ArtifactStatusRT, err error)
	// Upload content and create a artifacts.
	Upload(context.Context, *UploadPayload, io.ReadCloser) (res *ArtifactUploadRT, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ivcap"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.41"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "artifact"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list", "read", "upload"}

type ArtifactListItem struct {
	// ID
	ID string
	// Optional name
	Name *string
	// Artifact status
	Status string
	// Size of artifact in bytes
	Size *int64
	// Mime (content) type of artifact
	MimeType *string
	// time this artifact was created
	CreatedAt string
	Href      string `json:"href,omitempty"`
}

// ArtifactListRT is the result type of the artifact service list method.
type ArtifactListRT struct {
	// Artifacts
	Items []*ArtifactListItem
	// Time at which this list was valid
	AtTime *string
	Links  []*LinkT
}

// ArtifactStatusRT is the result type of the artifact service read method.
type ArtifactStatusRT struct {
	// Artifact ID
	ID string
	// Optional name
	Name *string
	// Artifact status
	Status string
	// Mime-type of data
	MimeType *string
	// Size of data
	Size *int64
	// URL of object this artifact is caching
	CacheOf *string
	// ETAG of artifact
	Etag *string
	// DateTime artifact was created
	CreatedAt *string
	// DateTime artifact was last modified
	LastModifiedAt *string
	// Reference to policy used
	Policy *string
	// Reference to billable account
	Account  *string
	DataHref *string
	Links    []*LinkT
}

// ArtifactUploadRT is the result type of the artifact service upload method.
type ArtifactUploadRT struct {
	// link back to record
	Location string
	// indicate version of TUS supported
	TusResumable *string
	// TUS offset for partially uploaded content
	TusOffset *int64
	// Artifact ID
	ID string
	// Optional name
	Name *string
	// Artifact status
	Status string
	// Mime-type of data
	MimeType *string
	// Size of data
	Size *int64
	// URL of object this artifact is caching
	CacheOf *string
	// ETAG of artifact
	Etag *string
	// DateTime artifact was created
	CreatedAt *string
	// DateTime artifact was last modified
	LastModifiedAt *string
	// Reference to policy used
	Policy *string
	// Reference to billable account
	Account  *string
	DataHref *string
	Links    []*LinkT
}

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

// ListPayload is the payload type of the artifact service list method.
type ListPayload struct {
	// The 'limit' query option sets the maximum number of items
	// to be included in the result.
	Limit int
	// The 'filter' system query option allows clients to filter a collection of
	// resources that are addressed by a request URL. The expression specified with
	// 'filter'
	// is evaluated for each resource in the collection, and only items where the
	// expression
	// evaluates to true are included in the response.
	Filter *string
	// The 'orderby' query option allows clients to request resources in either
	// ascending order using asc or descending order using desc. If asc or desc not
	// specified,
	// then the resources will be ordered in ascending order. The request below
	// orders Trips on
	// property EndsAt in descending order.
	OrderBy *string
	// When set order result in descending order. Ascending order is the lt.
	OrderDesc bool
	// Return the state of the respective resources at that time [now]
	AtTime *string
	// The content of 'page' is returned in the 'links' part of a previous query and
	// will when set, ALL other parameters, except for 'limit' are ignored.
	Page *string
	// JWT used for authentication
	JWT string
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

// ReadPayload is the payload type of the artifact service read method.
type ReadPayload struct {
	// ID of artifacts to show
	ID string
	// JWT used for authentication
	JWT string
}

// NotFound is the type returned when attempting to manage a resource that does
// not exist.
type ResourceNotFoundT struct {
	// ID of missing resource
	ID string
	// Message of error
	Message string
}

// Service necessary to fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// Unauthorized access to resource
type UnauthorizedT struct {
}

// UploadPayload is the payload type of the artifact service upload method.
type UploadPayload struct {
	// Content-Type header, MUST define type of uploaded content.
	ContentType *string `json:"content-type,omitempty"`
	// Content-Encoding header, MAY define encoding of content.
	ContentEncoding *string `json:"content-encoding,omitempty"`
	// Content-Length header, MAY define size of expected upload.
	ContentLength *int `json:"content-length,omitempty"`
	// X-Name header, MAY define a more human friendly name. Reusing a name will
	// NOT override an existing artifact with the same name
	Name *string
	// X-Collection header, MAY define an collection name as a simple way of
	// grouping artifacts
	Collection *string
	// X-Policy header, MAY define a specific policy to control access to this
	// artifact
	Policy *string
	// X-Content-Type header, used for initial, empty content creation requests.
	XContentType *string `json:"x-content-type,omitempty"`
	// X-Content-Length header, used for initial, empty content creation requests.
	XContentLength *int `json:"x-content-length,omitempty"`
	// Upload-Length header, sets the expected content size part of the TUS
	// protocol.
	UploadLength *int `json:"upload-length,omitempty"`
	// Tus-Resumable header, specifies TUS protocol version.
	TusResumable *string `json:"tus-resumable,omitempty"`
	// JWT used for authentication
	JWT string
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
