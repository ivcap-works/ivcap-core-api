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
	service "github.com/ivcap-works/ivcap-core-api/gen/service"
	serviceviews "github.com/ivcap-works/ivcap-core-api/gen/service/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateServiceRequestBody is the type of the "service" service
// "create_service" endpoint HTTP request body.
type CreateServiceRequestBody struct {
	// More detailed description of the service
	Description string `form:"description" json:"description" xml:"description"`
	// Reference to account revenues for this service should be credited to
	References []*ReferenceTRequestBodyRequestBody `form:"references,omitempty" json:"references,omitempty" xml:"references,omitempty"`
	// Link to banner image optionally used for this service
	Banner *string `form:"banner,omitempty" json:"banner,omitempty" xml:"banner,omitempty"`
	// Definition of the workflow to use for executing this service
	Workflow *WorkflowTRequestBodyRequestBody `form:"workflow" json:"workflow" xml:"workflow"`
	// Reference to policy used
	Policy *string `json:"policy"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional provider provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefT `form:"parameters" json:"parameters" xml:"parameters"`
}

// UpdateRequestBody is the type of the "service" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// More detailed description of the service
	Description string `form:"description" json:"description" xml:"description"`
	// Reference to account revenues for this service should be credited to
	References []*ReferenceTRequestBodyRequestBody `form:"references,omitempty" json:"references,omitempty" xml:"references,omitempty"`
	// Link to banner image optionally used for this service
	Banner *string `form:"banner,omitempty" json:"banner,omitempty" xml:"banner,omitempty"`
	// Definition of the workflow to use for executing this service
	Workflow *WorkflowTRequestBodyRequestBody `form:"workflow" json:"workflow" xml:"workflow"`
	// Reference to policy used
	Policy *string `json:"policy"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional provider provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefT `form:"parameters" json:"parameters" xml:"parameters"`
}

// ListResponseBody is the type of the "service" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Services
	Items []*ServiceListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// CreateServiceResponseBody is the type of the "service" service
// "create_service" endpoint HTTP response body.
type CreateServiceResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// More detailed description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Service status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Reference to billable account
	Account *string              `json:"account"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional provider provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// ReadResponseBody is the type of the "service" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// More detailed description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Service status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Reference to billable account
	Account *string              `json:"account"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional provider provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// UpdateResponseBody is the type of the "service" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// More detailed description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Service status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Reference to billable account
	Account *string              `json:"account"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Optional provider provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional provider provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameter definitions
	Parameters []*ParameterDefTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// ListBadRequestResponseBody is the type of the "service" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "service" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "service" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "service" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateServiceBadRequestResponseBody is the type of the "service" service
// "create_service" endpoint HTTP response body for the "bad-request" error.
type CreateServiceBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateServiceInvalidParameterResponseBody is the type of the "service"
// service "create_service" endpoint HTTP response body for the
// "invalid-parameter" error.
type CreateServiceInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// CreateServiceInvalidScopesResponseBody is the type of the "service" service
// "create_service" endpoint HTTP response body for the "invalid-scopes" error.
type CreateServiceInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateServiceNotImplementedResponseBody is the type of the "service" service
// "create_service" endpoint HTTP response body for the "not-implemented" error.
type CreateServiceNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateServiceAlreadyCreatedResponseBody is the type of the "service" service
// "create_service" endpoint HTTP response body for the "already-created" error.
type CreateServiceAlreadyCreatedResponseBody struct {
	// ID of already existing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateServiceNotFoundResponseBody is the type of the "service" service
// "create_service" endpoint HTTP response body for the "not-found" error.
type CreateServiceNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "service" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "service" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "service" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "service" service "read"
// endpoint HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateBadRequestResponseBody is the type of the "service" service "update"
// endpoint HTTP response body for the "bad-request" error.
type UpdateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateInvalidParameterResponseBody is the type of the "service" service
// "update" endpoint HTTP response body for the "invalid-parameter" error.
type UpdateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// UpdateInvalidScopesResponseBody is the type of the "service" service
// "update" endpoint HTTP response body for the "invalid-scopes" error.
type UpdateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateNotImplementedResponseBody is the type of the "service" service
// "update" endpoint HTTP response body for the "not-implemented" error.
type UpdateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateNotFoundResponseBody is the type of the "service" service "update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteBadRequestResponseBody is the type of the "service" service "delete"
// endpoint HTTP response body for the "bad-request" error.
type DeleteBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteInvalidScopesResponseBody is the type of the "service" service
// "delete" endpoint HTTP response body for the "invalid-scopes" error.
type DeleteInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteNotImplementedResponseBody is the type of the "service" service
// "delete" endpoint HTTP response body for the "not-implemented" error.
type DeleteNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ServiceListItemResponseBody is used to define fields on response body types.
type ServiceListItemResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional description of the service
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Reference to billable account
	Account *string `json:"account"`
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

// ReferenceTRequestBodyRequestBody is used to define fields on request body
// types.
type ReferenceTRequestBodyRequestBody struct {
	// Title of reference document
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Link to document
	URI *string `form:"uri,omitempty" json:"uri,omitempty" xml:"uri,omitempty"`
}

// WorkflowTRequestBodyRequestBody is used to define fields on request body
// types.
type WorkflowTRequestBodyRequestBody struct {
	// Type of workflow
	Type string `form:"type" json:"type" xml:"type"`
	// Type of workflow
	Basic *BasicWorkflowOptsTRequestBodyRequestBody `form:"basic,omitempty" json:"basic,omitempty" xml:"basic,omitempty"`
	// Defines the workflow using argo's WF schema
	Argo any `form:"argo,omitempty" json:"argo,omitempty" xml:"argo,omitempty"`
	// Type specific options - left for backward compatibility, if possible use
	// type specific elements
	Opts any `form:"opts,omitempty" json:"opts,omitempty" xml:"opts,omitempty"`
}

// BasicWorkflowOptsTRequestBodyRequestBody is used to define fields on request
// body types.
type BasicWorkflowOptsTRequestBodyRequestBody struct {
	// container image name
	Image string `form:"image" json:"image" xml:"image"`
	// Command to start the container - needed for some container runtimes
	Command []string `form:"command" json:"command" xml:"command"`
	// Defines memory resource requests and limits
	Memory *ResourceMemoryTRequestBodyRequestBody `form:"memory,omitempty" json:"memory,omitempty" xml:"memory,omitempty"`
	// Defines cpu resource requests and limits
	// (see
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu)
	CPU *ResourceMemoryTRequestBodyRequestBody `form:"cpu,omitempty" json:"cpu,omitempty" xml:"cpu,omitempty"`
	// Defines ephemeral storage resource requests and limits
	// (see
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#local-ephemeral-storage)
	EphemeralStorage *ResourceMemoryTRequestBodyRequestBody `json:"ephemeral-storage,omitempty"`
	// Defines required gpu type
	GpuType *string `json:"gpu-type,omitempty"`
	// Defines number of required gpu
	GpuNumber *int `json:"gpu-number,omitempty"`
}

// ResourceMemoryTRequestBodyRequestBody is used to define fields on request
// body types.
type ResourceMemoryTRequestBodyRequestBody struct {
	// minimal requirements [0]
	Request *string `form:"request,omitempty" json:"request,omitempty" xml:"request,omitempty"`
	// minimal requirements [system limit]
	Limit *string `form:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
}

// ParameterDefT is used to define fields on request body types.
type ParameterDefT struct {
	Name        *string          `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Label       *string          `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	Type        *string          `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Description *string          `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
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

// NewCreateServiceRequestBody builds the HTTP request body from the payload of
// the "create_service" endpoint of the "service" service.
func NewCreateServiceRequestBody(p *service.CreateServicePayload) *CreateServiceRequestBody {
	body := &CreateServiceRequestBody{
		Description: p.Services.Description,
		Banner:      p.Services.Banner,
		Policy:      p.Services.Policy,
		Name:        p.Services.Name,
	}
	if p.Services.References != nil {
		body.References = make([]*ReferenceTRequestBodyRequestBody, len(p.Services.References))
		for i, val := range p.Services.References {
			body.References[i] = marshalServiceReferenceTToReferenceTRequestBodyRequestBody(val)
		}
	}
	if p.Services.Workflow != nil {
		body.Workflow = marshalServiceWorkflowTToWorkflowTRequestBodyRequestBody(p.Services.Workflow)
	}
	if p.Services.Tags != nil {
		body.Tags = make([]string, len(p.Services.Tags))
		for i, val := range p.Services.Tags {
			body.Tags[i] = val
		}
	}
	if p.Services.Parameters != nil {
		body.Parameters = make([]*ParameterDefT, len(p.Services.Parameters))
		for i, val := range p.Services.Parameters {
			body.Parameters[i] = marshalServiceParameterDefTToParameterDefT(val)
		}
	} else {
		body.Parameters = []*ParameterDefT{}
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "service" service.
func NewUpdateRequestBody(p *service.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Description: p.Services.Description,
		Banner:      p.Services.Banner,
		Policy:      p.Services.Policy,
		Name:        p.Services.Name,
	}
	if p.Services.References != nil {
		body.References = make([]*ReferenceTRequestBodyRequestBody, len(p.Services.References))
		for i, val := range p.Services.References {
			body.References[i] = marshalServiceReferenceTToReferenceTRequestBodyRequestBody(val)
		}
	}
	if p.Services.Workflow != nil {
		body.Workflow = marshalServiceWorkflowTToWorkflowTRequestBodyRequestBody(p.Services.Workflow)
	}
	if p.Services.Tags != nil {
		body.Tags = make([]string, len(p.Services.Tags))
		for i, val := range p.Services.Tags {
			body.Tags[i] = val
		}
	}
	if p.Services.Parameters != nil {
		body.Parameters = make([]*ParameterDefT, len(p.Services.Parameters))
		for i, val := range p.Services.Parameters {
			body.Parameters[i] = marshalServiceParameterDefTToParameterDefT(val)
		}
	} else {
		body.Parameters = []*ParameterDefT{}
	}
	return body
}

// NewListServiceListRTOK builds a "service" service "list" endpoint result
// from a HTTP "OK" response.
func NewListServiceListRTOK(body *ListResponseBody) *serviceviews.ServiceListRTView {
	v := &serviceviews.ServiceListRTView{
		AtTime: body.AtTime,
	}
	v.Items = make([]*serviceviews.ServiceListItemView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalServiceListItemResponseBodyToServiceviewsServiceListItemView(val)
	}
	v.Links = make([]*serviceviews.LinkTView, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToServiceviewsLinkTView(val)
	}

	return v
}

// NewListBadRequest builds a service service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidCredential builds a service service list endpoint
// invalid-credential error.
func NewListInvalidCredential() *service.InvalidCredentialsT {
	v := &service.InvalidCredentialsT{}

	return v
}

// NewListInvalidParameter builds a service service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *service.InvalidParameterValue {
	v := &service.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a service service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a service service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAuthorized builds a service service list endpoint not-authorized
// error.
func NewListNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewCreateServiceServiceStatusRTCreated builds a "service" service
// "create_service" endpoint result from a HTTP "Created" response.
func NewCreateServiceServiceStatusRTCreated(body *CreateServiceResponseBody) *service.ServiceStatusRT {
	v := &service.ServiceStatusRT{
		ID:          *body.ID,
		Description: body.Description,
		Status:      *body.Status,
		Account:     *body.Account,
		Name:        body.Name,
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

// NewCreateServiceBadRequest builds a service service create_service endpoint
// bad-request error.
func NewCreateServiceBadRequest(body *CreateServiceBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewCreateServiceInvalidCredential builds a service service create_service
// endpoint invalid-credential error.
func NewCreateServiceInvalidCredential() *service.InvalidCredentialsT {
	v := &service.InvalidCredentialsT{}

	return v
}

// NewCreateServiceInvalidParameter builds a service service create_service
// endpoint invalid-parameter error.
func NewCreateServiceInvalidParameter(body *CreateServiceInvalidParameterResponseBody) *service.InvalidParameterValue {
	v := &service.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewCreateServiceInvalidScopes builds a service service create_service
// endpoint invalid-scopes error.
func NewCreateServiceInvalidScopes(body *CreateServiceInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateServiceNotImplemented builds a service service create_service
// endpoint not-implemented error.
func NewCreateServiceNotImplemented(body *CreateServiceNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewCreateServiceAlreadyCreated builds a service service create_service
// endpoint already-created error.
func NewCreateServiceAlreadyCreated(body *CreateServiceAlreadyCreatedResponseBody) *service.ResourceAlreadyCreatedT {
	v := &service.ResourceAlreadyCreatedT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateServiceNotFound builds a service service create_service endpoint
// not-found error.
func NewCreateServiceNotFound(body *CreateServiceNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateServiceNotAuthorized builds a service service create_service
// endpoint not-authorized error.
func NewCreateServiceNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewReadServiceStatusRTOK builds a "service" service "read" endpoint result
// from a HTTP "OK" response.
func NewReadServiceStatusRTOK(body *ReadResponseBody) *service.ServiceStatusRT {
	v := &service.ServiceStatusRT{
		ID:          *body.ID,
		Description: body.Description,
		Status:      *body.Status,
		Account:     *body.Account,
		Name:        body.Name,
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

// NewReadBadRequest builds a service service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidCredential builds a service service read endpoint
// invalid-credential error.
func NewReadInvalidCredential() *service.InvalidCredentialsT {
	v := &service.InvalidCredentialsT{}

	return v
}

// NewReadInvalidScopes builds a service service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a service service read endpoint not-implemented
// error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a service service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAuthorized builds a service service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewUpdateServiceStatusRTOK builds a "service" service "update" endpoint
// result from a HTTP "OK" response.
func NewUpdateServiceStatusRTOK(body *UpdateResponseBody) *service.ServiceStatusRT {
	v := &service.ServiceStatusRT{
		ID:          *body.ID,
		Description: body.Description,
		Status:      *body.Status,
		Account:     *body.Account,
		Name:        body.Name,
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

// NewUpdateBadRequest builds a service service update endpoint bad-request
// error.
func NewUpdateBadRequest(body *UpdateBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateInvalidCredential builds a service service update endpoint
// invalid-credential error.
func NewUpdateInvalidCredential() *service.InvalidCredentialsT {
	v := &service.InvalidCredentialsT{}

	return v
}

// NewUpdateInvalidParameter builds a service service update endpoint
// invalid-parameter error.
func NewUpdateInvalidParameter(body *UpdateInvalidParameterResponseBody) *service.InvalidParameterValue {
	v := &service.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewUpdateInvalidScopes builds a service service update endpoint
// invalid-scopes error.
func NewUpdateInvalidScopes(body *UpdateInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateNotImplemented builds a service service update endpoint
// not-implemented error.
func NewUpdateNotImplemented(body *UpdateNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateNotFound builds a service service update endpoint not-found error.
func NewUpdateNotFound(body *UpdateNotFoundResponseBody) *service.ResourceNotFoundT {
	v := &service.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateNotAuthorized builds a service service update endpoint
// not-authorized error.
func NewUpdateNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// NewDeleteBadRequest builds a service service delete endpoint bad-request
// error.
func NewDeleteBadRequest(body *DeleteBadRequestResponseBody) *service.BadRequestT {
	v := &service.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewDeleteInvalidCredential builds a service service delete endpoint
// invalid-credential error.
func NewDeleteInvalidCredential() *service.InvalidCredentialsT {
	v := &service.InvalidCredentialsT{}

	return v
}

// NewDeleteInvalidScopes builds a service service delete endpoint
// invalid-scopes error.
func NewDeleteInvalidScopes(body *DeleteInvalidScopesResponseBody) *service.InvalidScopesT {
	v := &service.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotImplemented builds a service service delete endpoint
// not-implemented error.
func NewDeleteNotImplemented(body *DeleteNotImplementedResponseBody) *service.NotImplementedT {
	v := &service.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotAuthorized builds a service service delete endpoint
// not-authorized error.
func NewDeleteNotAuthorized() *service.UnauthorizedT {
	v := &service.UnauthorizedT{}

	return v
}

// ValidateCreateServiceResponseBody runs the validations defined on
// create_service_response_body
func ValidateCreateServiceResponseBody(body *CreateServiceResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "inactive" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "inactive", "error"}))
		}
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
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "inactive" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "inactive", "error"}))
		}
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

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "inactive" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "inactive", "error"}))
		}
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

// ValidateCreateServiceBadRequestResponseBody runs the validations defined on
// create_service_bad-request_response_body
func ValidateCreateServiceBadRequestResponseBody(body *CreateServiceBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateServiceInvalidParameterResponseBody runs the validations
// defined on create_service_invalid-parameter_response_body
func ValidateCreateServiceInvalidParameterResponseBody(body *CreateServiceInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateServiceInvalidScopesResponseBody runs the validations defined
// on create_service_invalid-scopes_response_body
func ValidateCreateServiceInvalidScopesResponseBody(body *CreateServiceInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateCreateServiceNotImplementedResponseBody runs the validations defined
// on create_service_not-implemented_response_body
func ValidateCreateServiceNotImplementedResponseBody(body *CreateServiceNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateServiceAlreadyCreatedResponseBody runs the validations defined
// on create_service_already-created_response_body
func ValidateCreateServiceAlreadyCreatedResponseBody(body *CreateServiceAlreadyCreatedResponseBody) (err error) {
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

// ValidateCreateServiceNotFoundResponseBody runs the validations defined on
// create_service_not-found_response_body
func ValidateCreateServiceNotFoundResponseBody(body *CreateServiceNotFoundResponseBody) (err error) {
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

// ValidateUpdateBadRequestResponseBody runs the validations defined on
// update_bad-request_response_body
func ValidateUpdateBadRequestResponseBody(body *UpdateBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateInvalidParameterResponseBody runs the validations defined on
// update_invalid-parameter_response_body
func ValidateUpdateInvalidParameterResponseBody(body *UpdateInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateInvalidScopesResponseBody runs the validations defined on
// update_invalid-scopes_response_body
func ValidateUpdateInvalidScopesResponseBody(body *UpdateInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateNotImplementedResponseBody runs the validations defined on
// update_not-implemented_response_body
func ValidateUpdateNotImplementedResponseBody(body *UpdateNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateNotFoundResponseBody runs the validations defined on
// update_not-found_response_body
func ValidateUpdateNotFoundResponseBody(body *UpdateNotFoundResponseBody) (err error) {
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

// ValidateDeleteBadRequestResponseBody runs the validations defined on
// delete_bad-request_response_body
func ValidateDeleteBadRequestResponseBody(body *DeleteBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDeleteInvalidScopesResponseBody runs the validations defined on
// delete_invalid-scopes_response_body
func ValidateDeleteInvalidScopesResponseBody(body *DeleteInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateDeleteNotImplementedResponseBody runs the validations defined on
// delete_not-implemented_response_body
func ValidateDeleteNotImplementedResponseBody(body *DeleteNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateServiceListItemResponseBody runs the validations defined on
// ServiceListItemResponseBody
func ValidateServiceListItemResponseBody(body *ServiceListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "body"))
	}
	if body.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
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

// ValidateReferenceTRequestBodyRequestBody runs the validations defined on
// ReferenceTRequestBodyRequestBody
func ValidateReferenceTRequestBodyRequestBody(body *ReferenceTRequestBodyRequestBody) (err error) {
	if body.URI != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.uri", *body.URI, goa.FormatURI))
	}
	return
}

// ValidateWorkflowTRequestBodyRequestBody runs the validations defined on
// WorkflowTRequestBodyRequestBody
func ValidateWorkflowTRequestBodyRequestBody(body *WorkflowTRequestBodyRequestBody) (err error) {
	if body.Basic != nil {
		if err2 := ValidateBasicWorkflowOptsTRequestBodyRequestBody(body.Basic); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateBasicWorkflowOptsTRequestBodyRequestBody runs the validations
// defined on BasicWorkflowOptsTRequestBodyRequestBody
func ValidateBasicWorkflowOptsTRequestBodyRequestBody(body *BasicWorkflowOptsTRequestBodyRequestBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	return
}
