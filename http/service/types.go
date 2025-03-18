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
	service "github.com/ivcap-works/ivcap-core-api/gen/service"
	serviceviews "github.com/ivcap-works/ivcap-core-api/gen/service/views"
	goa "goa.design/goa/v3/pkg"
)

// ServiceCreateRequestBody is the type of the "service" service
// "service-create" endpoint HTTP request body.
type ServiceCreateRequestBody struct {
	// type of controller used for this service
	ControllerSchema string `form:"controller-schema" json:"controller-schema" xml:"controller-schema"`
	// controller definition
	Controller any `form:"controller" json:"controller" xml:"controller"`
	// Reference to policy used
	Policy string `form:"policy" json:"policy" xml:"policy"`
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// More detailed description of the service
	Description string `form:"description" json:"description" xml:"description"`
	// Optional tags defined for service to help in categorising them
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefT `form:"parameters" json:"parameters" xml:"parameters"`
}

// ServiceUpdateRequestBody is the type of the "service" service
// "service-update" endpoint HTTP request body.
type ServiceUpdateRequestBody struct {
	// type of controller used for this service
	ControllerSchema string `form:"controller-schema" json:"controller-schema" xml:"controller-schema"`
	// controller definition
	Controller any `form:"controller" json:"controller" xml:"controller"`
	// Reference to policy used
	Policy string `form:"policy" json:"policy" xml:"policy"`
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// More detailed description of the service
	Description string `form:"description" json:"description" xml:"description"`
	// Optional tags defined for service to help in categorising them
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefT `form:"parameters" json:"parameters" xml:"parameters"`
}

// ServiceListResponseBody is the type of the "service" service "service-list"
// endpoint HTTP response body.
type ServiceListResponseBody struct {
	// Services
	Items []*ServiceListItemTResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// ServiceCreateResponseBody is the type of the "service" service
// "service-create" endpoint HTTP response body.
type ServiceCreateResponseBody struct {
	// Service status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// type of controller used for this service
	ControllerSchema *string `form:"controller-schema,omitempty" json:"controller-schema,omitempty" xml:"controller-schema,omitempty"`
	// controller definition
	Controller any `form:"controller,omitempty" json:"controller,omitempty" xml:"controller,omitempty"`
	// Reference to billable account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// time this service has been available from
	ValidFrom *string `form:"valid-from,omitempty" json:"valid-from,omitempty" xml:"valid-from,omitempty"`
	// time this service has been available to
	ValidTo *string              `form:"valid-to,omitempty" json:"valid-to,omitempty" xml:"valid-to,omitempty"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// More detailed description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Optional tags defined for service to help in categorising them
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// ServiceReadResponseBody is the type of the "service" service "service-read"
// endpoint HTTP response body.
type ServiceReadResponseBody struct {
	// Service status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// type of controller used for this service
	ControllerSchema *string `form:"controller-schema,omitempty" json:"controller-schema,omitempty" xml:"controller-schema,omitempty"`
	// controller definition
	Controller any `form:"controller,omitempty" json:"controller,omitempty" xml:"controller,omitempty"`
	// Reference to billable account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// time this service has been available from
	ValidFrom *string `form:"valid-from,omitempty" json:"valid-from,omitempty" xml:"valid-from,omitempty"`
	// time this service has been available to
	ValidTo *string              `form:"valid-to,omitempty" json:"valid-to,omitempty" xml:"valid-to,omitempty"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// More detailed description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Optional tags defined for service to help in categorising them
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// ServiceUpdateResponseBody is the type of the "service" service
// "service-update" endpoint HTTP response body.
type ServiceUpdateResponseBody struct {
	// Service status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// type of controller used for this service
	ControllerSchema *string `form:"controller-schema,omitempty" json:"controller-schema,omitempty" xml:"controller-schema,omitempty"`
	// controller definition
	Controller any `form:"controller,omitempty" json:"controller,omitempty" xml:"controller,omitempty"`
	// Reference to billable account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// time this service has been available from
	ValidFrom *string `form:"valid-from,omitempty" json:"valid-from,omitempty" xml:"valid-from,omitempty"`
	// time this service has been available to
	ValidTo *string              `form:"valid-to,omitempty" json:"valid-to,omitempty" xml:"valid-to,omitempty"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// More detailed description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Optional tags defined for service to help in categorising them
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// JobListResponseBody is the type of the "service" service "job-list" endpoint
// HTTP response body.
type JobListResponseBody struct {
	// Jobs
	Items []*JobListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// JobReadResponseBody is the type of the "service" service "job-read" endpoint
// HTTP response body.
type JobReadResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Job status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional customer provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Reference to order
	Order *string `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	// Reference to service requested
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Mime type of request
	RequestContentType *string `form:"request-content-type,omitempty" json:"request-content-type,omitempty" xml:"request-content-type,omitempty"`
	// Request content
	RequestContent any `form:"request-content,omitempty" json:"request-content,omitempty" xml:"request-content,omitempty"`
	// Mime type of result
	ResultContentType *string `form:"result-content-type,omitempty" json:"result-content-type,omitempty" xml:"result-content-type,omitempty"`
	// Result content
	ResultContent any                               `form:"result-content,omitempty" json:"result-content,omitempty" xml:"result-content,omitempty"`
	Products      *PartialProductList2TResponseBody `form:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
	// Additional error message id status is 'Error' or 'Failed'
	ErrorMessage *string `form:"error-message,omitempty" json:"error-message,omitempty" xml:"error-message,omitempty"`
	// Reference to billable account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
	// DateTime job was submitted
	RequestedAt *string `form:"requested-at,omitempty" json:"requested-at,omitempty" xml:"requested-at,omitempty"`
	// DateTime job processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime job processing finished
	FinishedAt *string              `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	Links      []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// ServiceListBadRequestResponseBody is the type of the "service" service
// "service-list" endpoint HTTP response body for the "bad-request" error.
type ServiceListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceListInvalidParameterResponseBody is the type of the "service" service
// "service-list" endpoint HTTP response body for the "invalid-parameter" error.
type ServiceListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ServiceListInvalidScopesResponseBody is the type of the "service" service
// "service-list" endpoint HTTP response body for the "invalid-scopes" error.
type ServiceListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceListNotImplementedResponseBody is the type of the "service" service
// "service-list" endpoint HTTP response body for the "not-implemented" error.
type ServiceListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceCreateBadRequestResponseBody is the type of the "service" service
// "service-create" endpoint HTTP response body for the "bad-request" error.
type ServiceCreateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceCreateInvalidParameterResponseBody is the type of the "service"
// service "service-create" endpoint HTTP response body for the
// "invalid-parameter" error.
type ServiceCreateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ServiceCreateInvalidScopesResponseBody is the type of the "service" service
// "service-create" endpoint HTTP response body for the "invalid-scopes" error.
type ServiceCreateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceCreateNotImplementedResponseBody is the type of the "service" service
// "service-create" endpoint HTTP response body for the "not-implemented" error.
type ServiceCreateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceCreateAlreadyCreatedResponseBody is the type of the "service" service
// "service-create" endpoint HTTP response body for the "already-created" error.
type ServiceCreateAlreadyCreatedResponseBody struct {
	// ID of already existing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceCreateNotFoundResponseBody is the type of the "service" service
// "service-create" endpoint HTTP response body for the "not-found" error.
type ServiceCreateNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceReadBadRequestResponseBody is the type of the "service" service
// "service-read" endpoint HTTP response body for the "bad-request" error.
type ServiceReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceReadInvalidScopesResponseBody is the type of the "service" service
// "service-read" endpoint HTTP response body for the "invalid-scopes" error.
type ServiceReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceReadNotImplementedResponseBody is the type of the "service" service
// "service-read" endpoint HTTP response body for the "not-implemented" error.
type ServiceReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceReadNotFoundResponseBody is the type of the "service" service
// "service-read" endpoint HTTP response body for the "not-found" error.
type ServiceReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceUpdateBadRequestResponseBody is the type of the "service" service
// "service-update" endpoint HTTP response body for the "bad-request" error.
type ServiceUpdateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceUpdateInvalidParameterResponseBody is the type of the "service"
// service "service-update" endpoint HTTP response body for the
// "invalid-parameter" error.
type ServiceUpdateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ServiceUpdateInvalidScopesResponseBody is the type of the "service" service
// "service-update" endpoint HTTP response body for the "invalid-scopes" error.
type ServiceUpdateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceUpdateNotImplementedResponseBody is the type of the "service" service
// "service-update" endpoint HTTP response body for the "not-implemented" error.
type ServiceUpdateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceUpdateNotFoundResponseBody is the type of the "service" service
// "service-update" endpoint HTTP response body for the "not-found" error.
type ServiceUpdateNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceDeleteBadRequestResponseBody is the type of the "service" service
// "service-delete" endpoint HTTP response body for the "bad-request" error.
type ServiceDeleteBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceDeleteInvalidScopesResponseBody is the type of the "service" service
// "service-delete" endpoint HTTP response body for the "invalid-scopes" error.
type ServiceDeleteInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceDeleteNotImplementedResponseBody is the type of the "service" service
// "service-delete" endpoint HTTP response body for the "not-implemented" error.
type ServiceDeleteNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobListBadRequestResponseBody is the type of the "service" service
// "job-list" endpoint HTTP response body for the "bad-request" error.
type JobListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobListInvalidParameterResponseBody is the type of the "service" service
// "job-list" endpoint HTTP response body for the "invalid-parameter" error.
type JobListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// JobListInvalidScopesResponseBody is the type of the "service" service
// "job-list" endpoint HTTP response body for the "invalid-scopes" error.
type JobListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobListNotImplementedResponseBody is the type of the "service" service
// "job-list" endpoint HTTP response body for the "not-implemented" error.
type JobListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobCreateBadRequestResponseBody is the type of the "service" service
// "job-create" endpoint HTTP response body for the "bad-request" error.
type JobCreateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobCreateInvalidParameterResponseBody is the type of the "service" service
// "job-create" endpoint HTTP response body for the "invalid-parameter" error.
type JobCreateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// JobCreateInvalidScopesResponseBody is the type of the "service" service
// "job-create" endpoint HTTP response body for the "invalid-scopes" error.
type JobCreateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobCreateNotReadyYetResponseBody is the type of the "service" service
// "job-create" endpoint HTTP response body for the "not-ready-yet" error.
type JobCreateNotReadyYetResponseBody struct {
	// the ID of the job
	JobID *string `form:"job-id,omitempty" json:"job-id,omitempty" xml:"job-id,omitempty"`
	// the URL for the job
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The time in seconds after which an update may be available
	RetryLater *int `form:"retry-later,omitempty" json:"retry-later,omitempty" xml:"retry-later,omitempty"`
}

// JobCreateNotImplementedResponseBody is the type of the "service" service
// "job-create" endpoint HTTP response body for the "not-implemented" error.
type JobCreateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobCreateNotFoundResponseBody is the type of the "service" service
// "job-create" endpoint HTTP response body for the "not-found" error.
type JobCreateNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobReadBadRequestResponseBody is the type of the "service" service
// "job-read" endpoint HTTP response body for the "bad-request" error.
type JobReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobReadInvalidScopesResponseBody is the type of the "service" service
// "job-read" endpoint HTTP response body for the "invalid-scopes" error.
type JobReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobReadNotImplementedResponseBody is the type of the "service" service
// "job-read" endpoint HTTP response body for the "not-implemented" error.
type JobReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// JobReadNotFoundResponseBody is the type of the "service" service "job-read"
// endpoint HTTP response body for the "not-found" error.
type JobReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceListItemTResponseBody is used to define fields on response body types.
type ServiceListItemTResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Optional tags defined for service to help in categorising them
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// type of controller used for this service
	ControllerSchema *string `form:"controller-schema,omitempty" json:"controller-schema,omitempty" xml:"controller-schema,omitempty"`
	// time this service has been available from
	ValidFrom *string `form:"valid-from,omitempty" json:"valid-from,omitempty" xml:"valid-from,omitempty"`
	// time this service has been available to
	ValidTo *string `form:"valid-to,omitempty" json:"valid-to,omitempty" xml:"valid-to,omitempty"`
	Href    *string `json:"href,omitempty"`
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

// ParameterDefT is used to define fields on request body types.
type ParameterDefT struct {
	Name        string           `form:"name" json:"name" xml:"name"`
	Label       *string          `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	Type        string           `form:"type" json:"type" xml:"type"`
	Description string           `form:"description" json:"description" xml:"description"`
	Unit        *string          `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	Constant    *bool            `form:"constant,omitempty" json:"constant,omitempty" xml:"constant,omitempty"`
	Optional    *bool            `form:"optional,omitempty" json:"optional,omitempty" xml:"optional,omitempty"`
	Default     *string          `form:"default,omitempty" json:"default,omitempty" xml:"default,omitempty"`
	Options     []*ParameterOptT `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	Unary       *bool            `form:"unary,omitempty" json:"unary,omitempty" xml:"unary,omitempty"`
}

// ParameterOptT is used to define fields on request body types.
type ParameterOptT struct {
	Value       *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// ParameterDefTResponseBody is used to define fields on response body types.
type ParameterDefTResponseBody struct {
	Name        *string                      `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Label       *string                      `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	Type        *string                      `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Description *string                      `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Unit        *string                      `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	Constant    *bool                        `form:"constant,omitempty" json:"constant,omitempty" xml:"constant,omitempty"`
	Optional    *bool                        `form:"optional,omitempty" json:"optional,omitempty" xml:"optional,omitempty"`
	Default     *string                      `form:"default,omitempty" json:"default,omitempty" xml:"default,omitempty"`
	Options     []*ParameterOptTResponseBody `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	Unary       *bool                        `form:"unary,omitempty" json:"unary,omitempty" xml:"unary,omitempty"`
}

// ParameterOptTResponseBody is used to define fields on response body types.
type ParameterOptTResponseBody struct {
	Value       *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// JobListItemResponseBody is used to define fields on response body types.
type JobListItemResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Job status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime job processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime job processing finished
	FinishedAt *string `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	// Reference to service requested
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Reference to order
	Order *string `form:"order,omitempty" json:"order,omitempty" xml:"order,omitempty"`
	Href  *string `json:"href,omitempty"`
}

// PartialProductList2TResponseBody is used to define fields on response body
// types.
type PartialProductList2TResponseBody struct {
	// (Partial) list of products delivered by this order
	Items []*ProductListItem2TResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Links to more products, if there are any
	Links []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// ProductListItem2TResponseBody is used to define fields on response body
// types.
type ProductListItem2TResponseBody struct {
	ID       *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Status   *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	MimeType *string `json:"mime-type,omitempty"`
	Size     *int64  `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	Href     *string `json:"href,omitempty"`
	DataHref *string `json:"dataRef,omitempty"`
}

// NewServiceCreateRequestBody builds the HTTP request body from the payload of
// the "service-create" endpoint of the "service" service.
func NewServiceCreateRequestBody(p *service.ServiceCreatePayload) *ServiceCreateRequestBody {
	body := &ServiceCreateRequestBody{
		ControllerSchema: p.Service.ControllerSchema,
		Controller:       p.Service.Controller,
		Policy:           p.Service.Policy,
		ID:               p.Service.ID,
		Name:             p.Service.Name,
		Description:      p.Service.Description,
	}
	if p.Service.Tags != nil {
		body.Tags = make([]string, len(p.Service.Tags))
		for i, val := range p.Service.Tags {
			body.Tags[i] = val
		}
	}
	if p.Service.Parameters != nil {
		body.Parameters = make([]*ParameterDefT, len(p.Service.Parameters))
		for i, val := range p.Service.Parameters {
			body.Parameters[i] = marshalServiceParameterDefTToParameterDefT(val)
		}
	} else {
		body.Parameters = []*ParameterDefT{}
	}
	return body
}

// NewServiceUpdateRequestBody builds the HTTP request body from the payload of
// the "service-update" endpoint of the "service" service.
func NewServiceUpdateRequestBody(p *service.ServiceUpdatePayload) *ServiceUpdateRequestBody {
	body := &ServiceUpdateRequestBody{
		ControllerSchema: p.Service.ControllerSchema,
		Controller:       p.Service.Controller,
		Policy:           p.Service.Policy,
		ID:               p.Service.ID,
		Name:             p.Service.Name,
		Description:      p.Service.Description,
	}
	if p.Service.Tags != nil {
		body.Tags = make([]string, len(p.Service.Tags))
		for i, val := range p.Service.Tags {
			body.Tags[i] = val
		}
	}
	if p.Service.Parameters != nil {
		body.Parameters = make([]*ParameterDefT, len(p.Service.Parameters))
		for i, val := range p.Service.Parameters {
			body.Parameters[i] = marshalServiceParameterDefTToParameterDefT(val)
		}
	} else {
		body.Parameters = []*ParameterDefT{}
	}
	return body
}

// NewServiceListRTViewOK builds a "service" service "service-list" endpoint
// result from a HTTP "OK" response.
func NewServiceListRTViewOK(body *ServiceListResponseBody) *serviceviews.ServiceListRTView {
	v := &serviceviews.ServiceListRTView{
		AtTime: body.AtTime,
	}
	v.Items = make([]*serviceviews.ServiceListItemTView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalServiceListItemTResponseBodyToServiceviewsServiceListItemTView(val)
	}
	v.Links = make([]*serviceviews.LinkTView, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceviewsLinkTView(val)
	}

	return v
}

// NewServiceListBadRequest builds a service service service-list endpoint
// bad-request error.
func NewServiceListBadRequest(body *ServiceListBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewServiceListInvalidParameter builds a service service service-list
// endpoint invalid-parameter error.
func NewServiceListInvalidParameter(body *ServiceListInvalidParameterResponseBody) *service.InvalidParameterT {
	v := &service.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewServiceListInvalidScopes builds a service service service-list endpoint
// invalid-scopes error.
func NewServiceListInvalidScopes(body *ServiceListInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceListNotImplemented builds a service service service-list endpoint
// not-implemented error.
func NewServiceListNotImplemented(body *ServiceListNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewServiceListNotAvailable builds a service service service-list endpoint
// not-available error.
func NewServiceListNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewServiceListNotAuthorized builds a service service service-list endpoint
// not-authorized error.
func NewServiceListNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewServiceCreateServiceStatusRTCreated builds a "service" service
// "service-create" endpoint result from a HTTP "Created" response.
func NewServiceCreateServiceStatusRTCreated(body *ServiceCreateResponseBody) *service.ServiceStatusRT {
	v := &service.ServiceStatusRT{
		Status:           *body.Status,
		ControllerSchema: *body.ControllerSchema,
		Controller:       body.Controller,
		Account:          *body.Account,
		Policy:           *body.Policy,
		ValidFrom:        body.ValidFrom,
		ValidTo:          body.ValidTo,
		ID:               *body.ID,
		Name:             body.Name,
		Description:      *body.Description,
	}
	v.Links = make([]*service.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceLinkT(val)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*service.ParameterDefT, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterDefTResponseBodyToServiceParameterDefT(val)
	}

	return v
}

// NewServiceCreateBadRequest builds a service service service-create endpoint
// bad-request error.
func NewServiceCreateBadRequest(body *ServiceCreateBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewServiceCreateInvalidParameter builds a service service service-create
// endpoint invalid-parameter error.
func NewServiceCreateInvalidParameter(body *ServiceCreateInvalidParameterResponseBody) *service.InvalidParameterT {
	v := &service.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewServiceCreateInvalidScopes builds a service service service-create
// endpoint invalid-scopes error.
func NewServiceCreateInvalidScopes(body *ServiceCreateInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceCreateNotImplemented builds a service service service-create
// endpoint not-implemented error.
func NewServiceCreateNotImplemented(body *ServiceCreateNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewServiceCreateAlreadyCreated builds a service service service-create
// endpoint already-created error.
func NewServiceCreateAlreadyCreated(body *ServiceCreateAlreadyCreatedResponseBody) *service.ResourceAlreadyCreatedT {
	v := &service.ResourceAlreadyCreatedT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceCreateNotFound builds a service service service-create endpoint
// not-found error.
func NewServiceCreateNotFound(body *ServiceCreateNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceCreateNotAvailable builds a service service service-create
// endpoint not-available error.
func NewServiceCreateNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewServiceCreateNotAuthorized builds a service service service-create
// endpoint not-authorized error.
func NewServiceCreateNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewServiceReadServiceStatusRTOK builds a "service" service "service-read"
// endpoint result from a HTTP "OK" response.
func NewServiceReadServiceStatusRTOK(body *ServiceReadResponseBody) *service.ServiceStatusRT {
	v := &service.ServiceStatusRT{
		Status:           *body.Status,
		ControllerSchema: *body.ControllerSchema,
		Controller:       body.Controller,
		Account:          *body.Account,
		Policy:           *body.Policy,
		ValidFrom:        body.ValidFrom,
		ValidTo:          body.ValidTo,
		ID:               *body.ID,
		Name:             body.Name,
		Description:      *body.Description,
	}
	v.Links = make([]*service.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceLinkT(val)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*service.ParameterDefT, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterDefTResponseBodyToServiceParameterDefT(val)
	}

	return v
}

// NewServiceReadBadRequest builds a service service service-read endpoint
// bad-request error.
func NewServiceReadBadRequest(body *ServiceReadBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewServiceReadInvalidScopes builds a service service service-read endpoint
// invalid-scopes error.
func NewServiceReadInvalidScopes(body *ServiceReadInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceReadNotImplemented builds a service service service-read endpoint
// not-implemented error.
func NewServiceReadNotImplemented(body *ServiceReadNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewServiceReadNotFound builds a service service service-read endpoint
// not-found error.
func NewServiceReadNotFound(body *ServiceReadNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceReadNotAvailable builds a service service service-read endpoint
// not-available error.
func NewServiceReadNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewServiceReadNotAuthorized builds a service service service-read endpoint
// not-authorized error.
func NewServiceReadNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewServiceUpdateServiceStatusRTOK builds a "service" service
// "service-update" endpoint result from a HTTP "OK" response.
func NewServiceUpdateServiceStatusRTOK(body *ServiceUpdateResponseBody) *service.ServiceStatusRT {
	v := &service.ServiceStatusRT{
		Status:           *body.Status,
		ControllerSchema: *body.ControllerSchema,
		Controller:       body.Controller,
		Account:          *body.Account,
		Policy:           *body.Policy,
		ValidFrom:        body.ValidFrom,
		ValidTo:          body.ValidTo,
		ID:               *body.ID,
		Name:             body.Name,
		Description:      *body.Description,
	}
	v.Links = make([]*service.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceLinkT(val)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*service.ParameterDefT, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterDefTResponseBodyToServiceParameterDefT(val)
	}

	return v
}

// NewServiceUpdateBadRequest builds a service service service-update endpoint
// bad-request error.
func NewServiceUpdateBadRequest(body *ServiceUpdateBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewServiceUpdateInvalidParameter builds a service service service-update
// endpoint invalid-parameter error.
func NewServiceUpdateInvalidParameter(body *ServiceUpdateInvalidParameterResponseBody) *service.InvalidParameterT {
	v := &service.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewServiceUpdateInvalidScopes builds a service service service-update
// endpoint invalid-scopes error.
func NewServiceUpdateInvalidScopes(body *ServiceUpdateInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceUpdateNotImplemented builds a service service service-update
// endpoint not-implemented error.
func NewServiceUpdateNotImplemented(body *ServiceUpdateNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewServiceUpdateNotFound builds a service service service-update endpoint
// not-found error.
func NewServiceUpdateNotFound(body *ServiceUpdateNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceUpdateNotAvailable builds a service service service-update
// endpoint not-available error.
func NewServiceUpdateNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewServiceUpdateNotAuthorized builds a service service service-update
// endpoint not-authorized error.
func NewServiceUpdateNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewServiceDeleteBadRequest builds a service service service-delete endpoint
// bad-request error.
func NewServiceDeleteBadRequest(body *ServiceDeleteBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewServiceDeleteInvalidScopes builds a service service service-delete
// endpoint invalid-scopes error.
func NewServiceDeleteInvalidScopes(body *ServiceDeleteInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewServiceDeleteNotImplemented builds a service service service-delete
// endpoint not-implemented error.
func NewServiceDeleteNotImplemented(body *ServiceDeleteNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewServiceDeleteNotAvailable builds a service service service-delete
// endpoint not-available error.
func NewServiceDeleteNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewServiceDeleteNotAuthorized builds a service service service-delete
// endpoint not-authorized error.
func NewServiceDeleteNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewJobListRTViewOK builds a "service" service "job-list" endpoint result
// from a HTTP "OK" response.
func NewJobListRTViewOK(body *JobListResponseBody) *serviceviews.JobListRTView {
	v := &serviceviews.JobListRTView{
		AtTime: body.AtTime,
	}
	v.Items = make([]*serviceviews.JobListItemView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalJobListItemResponseBodyToServiceviewsJobListItemView(val)
	}
	v.Links = make([]*serviceviews.LinkTView, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceviewsLinkTView(val)
	}

	return v
}

// NewJobListBadRequest builds a service service job-list endpoint bad-request
// error.
func NewJobListBadRequest(body *JobListBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewJobListInvalidParameter builds a service service job-list endpoint
// invalid-parameter error.
func NewJobListInvalidParameter(body *JobListInvalidParameterResponseBody) *service.InvalidParameterT {
	v := &service.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewJobListInvalidScopes builds a service service job-list endpoint
// invalid-scopes error.
func NewJobListInvalidScopes(body *JobListInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewJobListNotImplemented builds a service service job-list endpoint
// not-implemented error.
func NewJobListNotImplemented(body *JobListNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewJobListNotAvailable builds a service service job-list endpoint
// not-available error.
func NewJobListNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewJobListNotAuthorized builds a service service job-list endpoint
// not-authorized error.
func NewJobListNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewJobCreateResultOK builds a "service" service "job-create" endpoint result
// from a HTTP "OK" response.
func NewJobCreateResultOK(outContentType string, outOrderID string, jobID string, jobURL *string) *service.JobCreateResult {
	v := &service.JobCreateResult{}
	v.OutContentType = outContentType
	v.OutOrderID = outOrderID
	v.JobID = jobID
	v.JobURL = jobURL

	return v
}

// NewJobCreateBadRequest builds a service service job-create endpoint
// bad-request error.
func NewJobCreateBadRequest(body *JobCreateBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewJobCreateInvalidParameter builds a service service job-create endpoint
// invalid-parameter error.
func NewJobCreateInvalidParameter(body *JobCreateInvalidParameterResponseBody) *service.InvalidParameterT {
	v := &service.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewJobCreateInvalidScopes builds a service service job-create endpoint
// invalid-scopes error.
func NewJobCreateInvalidScopes(body *JobCreateInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewJobCreateNotReadyYet builds a service service job-create endpoint
// not-ready-yet error.
func NewJobCreateNotReadyYet(body *JobCreateNotReadyYetResponseBody) *service.JobRetryLaterT {
	v := &service.JobRetryLaterT{
		JobID:      body.JobID,
		Location:   *body.Location,
		RetryLater: *body.RetryLater,
	}

	return v
}

// NewJobCreateNotImplemented builds a service service job-create endpoint
// not-implemented error.
func NewJobCreateNotImplemented(body *JobCreateNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewJobCreateNotFound builds a service service job-create endpoint not-found
// error.
func NewJobCreateNotFound(body *JobCreateNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewJobCreateNotAvailable builds a service service job-create endpoint
// not-available error.
func NewJobCreateNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewJobCreateTemporaryRedirect builds a service service job-create endpoint
// temporary-redirect error.
func NewJobCreateTemporaryRedirect(location string) *service.TemporaryRedirectT {
	v := &service.TemporaryRedirectT{}
	v.Location = location

	return v
}

// NewJobCreateNotAuthorized builds a service service job-create endpoint
// not-authorized error.
func NewJobCreateNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewJobReadJobStatusRTOK builds a "service" service "job-read" endpoint
// result from a HTTP "OK" response.
func NewJobReadJobStatusRTOK(body *JobReadResponseBody) *service.JobStatusRT {
	v := &service.JobStatusRT{
		ID:                 *body.ID,
		Status:             *body.Status,
		Name:               body.Name,
		Order:              *body.Order,
		Service:            *body.Service,
		RequestContentType: *body.RequestContentType,
		RequestContent:     body.RequestContent,
		ResultContentType:  body.ResultContentType,
		ResultContent:      body.ResultContent,
		ErrorMessage:       body.ErrorMessage,
		Account:            *body.Account,
		Policy:             *body.Policy,
		RequestedAt:        *body.RequestedAt,
		StartedAt:          body.StartedAt,
		FinishedAt:         body.FinishedAt,
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.Products != nil {
		v.Products = unmarshalPartialProductList2TResponseBodyToServicePartialProductList2T(body.Products)
	}
	v.Links = make([]*service.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceLinkT(val)
	}

	return v
}

// NewJobReadBadRequest builds a service service job-read endpoint bad-request
// error.
func NewJobReadBadRequest(body *JobReadBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewJobReadInvalidScopes builds a service service job-read endpoint
// invalid-scopes error.
func NewJobReadInvalidScopes(body *JobReadInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewJobReadNotImplemented builds a service service job-read endpoint
// not-implemented error.
func NewJobReadNotImplemented(body *JobReadNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewJobReadNotFound builds a service service job-read endpoint not-found
// error.
func NewJobReadNotFound(body *JobReadNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewJobReadNotAvailable builds a service service job-read endpoint
// not-available error.
func NewJobReadNotAvailable() *service.ServiceNotAvailableT {
	v := &service.ServiceNotAvailableT{}

	return v
}

// NewJobReadNotAuthorized builds a service service job-read endpoint
// not-authorized error.
func NewJobReadNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// ValidateServiceCreateResponseBody runs the validations defined on
// Service-CreateResponseBody
func ValidateServiceCreateResponseBody(body *ServiceCreateResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.ControllerSchema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller-schema", "body"))
	}
	if body.Controller == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller", "body"))
	}
	if body.Policy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("policy", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "inactive" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "inactive", "error"}))
		}
	}
	if body.ControllerSchema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.controller-schema", *body.ControllerSchema, goa.FormatURI))
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
	}
	if body.Policy != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
	}
	if body.ValidFrom != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-from", *body.ValidFrom, goa.FormatDateTime))
	}
	if body.ValidTo != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-to", *body.ValidTo, goa.FormatDateTime))
	}
	for _, e := range body.Links {
		if e != nil {
			if err2 := ValidateLinkTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	for _, e := range body.Parameters {
		if e != nil {
			if err2 := ValidateParameterDefTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateServiceReadResponseBody runs the validations defined on
// Service-ReadResponseBody
func ValidateServiceReadResponseBody(body *ServiceReadResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.ControllerSchema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller-schema", "body"))
	}
	if body.Controller == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller", "body"))
	}
	if body.Policy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("policy", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "inactive" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "inactive", "error"}))
		}
	}
	if body.ControllerSchema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.controller-schema", *body.ControllerSchema, goa.FormatURI))
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
	}
	if body.Policy != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
	}
	if body.ValidFrom != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-from", *body.ValidFrom, goa.FormatDateTime))
	}
	if body.ValidTo != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-to", *body.ValidTo, goa.FormatDateTime))
	}
	for _, e := range body.Links {
		if e != nil {
			if err2 := ValidateLinkTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	for _, e := range body.Parameters {
		if e != nil {
			if err2 := ValidateParameterDefTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateServiceUpdateResponseBody runs the validations defined on
// Service-UpdateResponseBody
func ValidateServiceUpdateResponseBody(body *ServiceUpdateResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.ControllerSchema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller-schema", "body"))
	}
	if body.Controller == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller", "body"))
	}
	if body.Policy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("policy", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "inactive" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "inactive", "error"}))
		}
	}
	if body.ControllerSchema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.controller-schema", *body.ControllerSchema, goa.FormatURI))
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
	}
	if body.Policy != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
	}
	if body.ValidFrom != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-from", *body.ValidFrom, goa.FormatDateTime))
	}
	if body.ValidTo != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-to", *body.ValidTo, goa.FormatDateTime))
	}
	for _, e := range body.Links {
		if e != nil {
			if err2 := ValidateLinkTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	for _, e := range body.Parameters {
		if e != nil {
			if err2 := ValidateParameterDefTResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateJobReadResponseBody runs the validations defined on
// Job-ReadResponseBody
func ValidateJobReadResponseBody(body *JobReadResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Order == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("order", "body"))
	}
	if body.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "body"))
	}
	if body.RequestContentType == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("request-content-type", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.Policy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("policy", "body"))
	}
	if body.RequestedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("requested-at", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Status != nil {
		if !(*body.Status == "unknown" || *body.Status == "pending" || *body.Status == "scheduled" || *body.Status == "executing" || *body.Status == "succeeded" || *body.Status == "failed" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if body.Order != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.order", *body.Order, goa.FormatURI))
	}
	if body.Service != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.service", *body.Service, goa.FormatURI))
	}
	if body.Products != nil {
		if err2 := ValidatePartialProductList2TResponseBody(body.Products); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
	}
	if body.Policy != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
	}
	if body.RequestedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.requested-at", *body.RequestedAt, goa.FormatDateTime))
	}
	if body.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.started-at", *body.StartedAt, goa.FormatDateTime))
	}
	if body.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.finished-at", *body.FinishedAt, goa.FormatDateTime))
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

// ValidateServiceListBadRequestResponseBody runs the validations defined on
// service-list_bad-request_response_body
func ValidateServiceListBadRequestResponseBody(body *ServiceListBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceListInvalidParameterResponseBody runs the validations defined
// on service-list_invalid-parameter_response_body
func ValidateServiceListInvalidParameterResponseBody(body *ServiceListInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceListInvalidScopesResponseBody runs the validations defined on
// service-list_invalid-scopes_response_body
func ValidateServiceListInvalidScopesResponseBody(body *ServiceListInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateServiceListNotImplementedResponseBody runs the validations defined
// on service-list_not-implemented_response_body
func ValidateServiceListNotImplementedResponseBody(body *ServiceListNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceCreateBadRequestResponseBody runs the validations defined on
// service-create_bad-request_response_body
func ValidateServiceCreateBadRequestResponseBody(body *ServiceCreateBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceCreateInvalidParameterResponseBody runs the validations
// defined on service-create_invalid-parameter_response_body
func ValidateServiceCreateInvalidParameterResponseBody(body *ServiceCreateInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceCreateInvalidScopesResponseBody runs the validations defined
// on service-create_invalid-scopes_response_body
func ValidateServiceCreateInvalidScopesResponseBody(body *ServiceCreateInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateServiceCreateNotImplementedResponseBody runs the validations defined
// on service-create_not-implemented_response_body
func ValidateServiceCreateNotImplementedResponseBody(body *ServiceCreateNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceCreateAlreadyCreatedResponseBody runs the validations defined
// on service-create_already-created_response_body
func ValidateServiceCreateAlreadyCreatedResponseBody(body *ServiceCreateAlreadyCreatedResponseBody) (err error) {
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

// ValidateServiceCreateNotFoundResponseBody runs the validations defined on
// service-create_not-found_response_body
func ValidateServiceCreateNotFoundResponseBody(body *ServiceCreateNotFoundResponseBody) (err error) {
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

// ValidateServiceReadBadRequestResponseBody runs the validations defined on
// service-read_bad-request_response_body
func ValidateServiceReadBadRequestResponseBody(body *ServiceReadBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceReadInvalidScopesResponseBody runs the validations defined on
// service-read_invalid-scopes_response_body
func ValidateServiceReadInvalidScopesResponseBody(body *ServiceReadInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateServiceReadNotImplementedResponseBody runs the validations defined
// on service-read_not-implemented_response_body
func ValidateServiceReadNotImplementedResponseBody(body *ServiceReadNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceReadNotFoundResponseBody runs the validations defined on
// service-read_not-found_response_body
func ValidateServiceReadNotFoundResponseBody(body *ServiceReadNotFoundResponseBody) (err error) {
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

// ValidateServiceUpdateBadRequestResponseBody runs the validations defined on
// service-update_bad-request_response_body
func ValidateServiceUpdateBadRequestResponseBody(body *ServiceUpdateBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceUpdateInvalidParameterResponseBody runs the validations
// defined on service-update_invalid-parameter_response_body
func ValidateServiceUpdateInvalidParameterResponseBody(body *ServiceUpdateInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceUpdateInvalidScopesResponseBody runs the validations defined
// on service-update_invalid-scopes_response_body
func ValidateServiceUpdateInvalidScopesResponseBody(body *ServiceUpdateInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateServiceUpdateNotImplementedResponseBody runs the validations defined
// on service-update_not-implemented_response_body
func ValidateServiceUpdateNotImplementedResponseBody(body *ServiceUpdateNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceUpdateNotFoundResponseBody runs the validations defined on
// service-update_not-found_response_body
func ValidateServiceUpdateNotFoundResponseBody(body *ServiceUpdateNotFoundResponseBody) (err error) {
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

// ValidateServiceDeleteBadRequestResponseBody runs the validations defined on
// service-delete_bad-request_response_body
func ValidateServiceDeleteBadRequestResponseBody(body *ServiceDeleteBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceDeleteInvalidScopesResponseBody runs the validations defined
// on service-delete_invalid-scopes_response_body
func ValidateServiceDeleteInvalidScopesResponseBody(body *ServiceDeleteInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateServiceDeleteNotImplementedResponseBody runs the validations defined
// on service-delete_not-implemented_response_body
func ValidateServiceDeleteNotImplementedResponseBody(body *ServiceDeleteNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobListBadRequestResponseBody runs the validations defined on
// job-list_bad-request_response_body
func ValidateJobListBadRequestResponseBody(body *JobListBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobListInvalidParameterResponseBody runs the validations defined on
// job-list_invalid-parameter_response_body
func ValidateJobListInvalidParameterResponseBody(body *JobListInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobListInvalidScopesResponseBody runs the validations defined on
// job-list_invalid-scopes_response_body
func ValidateJobListInvalidScopesResponseBody(body *JobListInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateJobListNotImplementedResponseBody runs the validations defined on
// job-list_not-implemented_response_body
func ValidateJobListNotImplementedResponseBody(body *JobListNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobCreateBadRequestResponseBody runs the validations defined on
// job-create_bad-request_response_body
func ValidateJobCreateBadRequestResponseBody(body *JobCreateBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobCreateInvalidParameterResponseBody runs the validations defined
// on job-create_invalid-parameter_response_body
func ValidateJobCreateInvalidParameterResponseBody(body *JobCreateInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobCreateInvalidScopesResponseBody runs the validations defined on
// job-create_invalid-scopes_response_body
func ValidateJobCreateInvalidScopesResponseBody(body *JobCreateInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateJobCreateNotReadyYetResponseBody runs the validations defined on
// job-create_not-ready-yet_response_body
func ValidateJobCreateNotReadyYetResponseBody(body *JobCreateNotReadyYetResponseBody) (err error) {
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	if body.RetryLater == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("retry-later", "body"))
	}
	return
}

// ValidateJobCreateNotImplementedResponseBody runs the validations defined on
// job-create_not-implemented_response_body
func ValidateJobCreateNotImplementedResponseBody(body *JobCreateNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobCreateNotFoundResponseBody runs the validations defined on
// job-create_not-found_response_body
func ValidateJobCreateNotFoundResponseBody(body *JobCreateNotFoundResponseBody) (err error) {
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

// ValidateJobReadBadRequestResponseBody runs the validations defined on
// job-read_bad-request_response_body
func ValidateJobReadBadRequestResponseBody(body *JobReadBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobReadInvalidScopesResponseBody runs the validations defined on
// job-read_invalid-scopes_response_body
func ValidateJobReadInvalidScopesResponseBody(body *JobReadInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateJobReadNotImplementedResponseBody runs the validations defined on
// job-read_not-implemented_response_body
func ValidateJobReadNotImplementedResponseBody(body *JobReadNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateJobReadNotFoundResponseBody runs the validations defined on
// job-read_not-found_response_body
func ValidateJobReadNotFoundResponseBody(body *JobReadNotFoundResponseBody) (err error) {
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

// ValidateServiceListItemTResponseBody runs the validations defined on
// ServiceListItemTResponseBody
func ValidateServiceListItemTResponseBody(body *ServiceListItemTResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.ControllerSchema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller-schema", "body"))
	}
	if body.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.ValidFrom != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-from", *body.ValidFrom, goa.FormatDateTime))
	}
	if body.ValidTo != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.valid-to", *body.ValidTo, goa.FormatDateTime))
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

// ValidateParameterDefTResponseBody runs the validations defined on
// ParameterDefTResponseBody
func ValidateParameterDefTResponseBody(body *ParameterDefTResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}

// ValidateJobListItemResponseBody runs the validations defined on
// JobListItemResponseBody
func ValidateJobListItemResponseBody(body *JobListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "body"))
	}
	if body.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Status != nil {
		if !(*body.Status == "unknown" || *body.Status == "pending" || *body.Status == "scheduled" || *body.Status == "executing" || *body.Status == "succeeded" || *body.Status == "failed" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if body.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.started-at", *body.StartedAt, goa.FormatDateTime))
	}
	if body.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.finished-at", *body.FinishedAt, goa.FormatDateTime))
	}
	if body.Service != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.service", *body.Service, goa.FormatURI))
	}
	if body.Order != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.order", *body.Order, goa.FormatURI))
	}
	return
}

// ValidatePartialProductList2TResponseBody runs the validations defined on
// PartialProductList2TResponseBody
func ValidatePartialProductList2TResponseBody(body *PartialProductList2TResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	for _, e := range body.Items {
		if e != nil {
			if err2 := ValidateProductListItem2TResponseBody(e); err2 != nil {
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

// ValidateProductListItem2TResponseBody runs the validations defined on
// ProductListItem2TResponseBody
func ValidateProductListItem2TResponseBody(body *ProductListItem2TResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "body"))
	}
	return
}
