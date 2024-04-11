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

package service

import (
	"context"

	serviceviews "github.com/ivcap-works/ivcap-core-api/gen/service/views"
	"goa.design/goa/v3/security"
)

// Manage the life cycle of a service offered on the CRE marketplace.
type Service interface {
	// list services
	List(context.Context, *ListPayload) (res *ServiceListRT, err error)
	// Create a new services and return its status.
	CreateService(context.Context, *CreateServicePayload) (res *ServiceStatusRT, err error)
	// Show services by ID
	Read(context.Context, *ReadPayload) (res *ServiceStatusRT, err error)
	// Update an existing services and return its status.
	Update(context.Context, *UpdatePayload) (res *ServiceStatusRT, err error)
	// Delete an existing services.
	Delete(context.Context, *DeletePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ivcap"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.35"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "create_service", "read", "update", "delete"}

// Bad arguments supplied.
type BadRequestT struct {
	// Information message
	Message string
}

type BasicWorkflowOptsT struct {
	// container image name
	Image string
	// Command to start the container - needed for some container runtimes
	Command []string
	// Defines memory resource requests and limits
	Memory *ResourceMemoryT
	// Defines cpu resource requests and limits
	// (see
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu)
	CPU *ResourceMemoryT
	// Defines ephemeral storage resource requests and limits
	// (see
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#local-ephemeral-storage)
	EphemeralStorage *ResourceMemoryT `json:"ephemeral-storage,omitempty"`
	// Defines required gpu type
	GpuType *string `json:"gpu-type,omitempty"`
	// Defines number of required gpu
	GpuNumber *int `json:"gpu-number,omitempty"`
}

// CreateServicePayload is the payload type of the service service
// create_service method.
type CreateServicePayload struct {
	// New services description
	Services *ServiceDefinitionT
	// JWT used for authentication
	JWT string
}

// DeletePayload is the payload type of the service service delete method.
type DeletePayload struct {
	// ID of services to update
	ID string
	// JWT used for authentication
	JWT string
}

// Provided credential is not valid.
type InvalidCredentialsT struct {
}

// InvalidParameterValue is the error returned when a parameter has the wrong
// value.
type InvalidParameterValue struct {
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

// ListPayload is the payload type of the service service list method.
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

type ParameterDefT struct {
	Name        *string
	Label       *string
	Type        *string
	Description *string
	Unit        *string
	Constant    *bool
	Optional    *bool
	Default     *string
	Options     []*ParameterOptT
	Unary       *bool
}

type ParameterOptT struct {
	Value       *string
	Description *string
}

// ReadPayload is the payload type of the service service read method.
type ReadPayload struct {
	// ID of services to show
	ID string
	// JWT used for authentication
	JWT string
}

type ReferenceT struct {
	// Title of reference document
	Title *string
	// Link to document
	URI *string
}

// Will be returned when receiving a request to create and already existing
// resource.
type ResourceAlreadyCreatedT struct {
	// ID of already existing resource
	ID string
	// Message of error
	Message string
}

// See
// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
// for units
type ResourceMemoryT struct {
	// minimal requirements [0]
	Request *string
	// minimal requirements [system limit]
	Limit *string
}

// NotFound is the type returned when attempting to manage a resource that does
// not exist.
type ResourceNotFoundT struct {
	// ID of missing resource
	ID string
	// Message of error
	Message string
}

type ServiceDefinitionT struct {
	// More detailed description of the service
	Description string
	// Reference to account revenues for this service should be credited to
	References []*ReferenceT
	// Link to banner image optionally used for this service
	Banner *string
	// Definition of the workflow to use for executing this service
	Workflow *WorkflowT
	// Reference to policy used
	Policy *string `json:"policy"`
	// Optional provider provided name
	Name *string
	// Optional provider provided tags
	Tags []string
	// Service parameter definitions
	Parameters []*ParameterDefT
}

type ServiceListItem struct {
	// ID
	ID string
	// Optional customer provided name
	Name *string
	// Optional description of the service
	Description *string
	// Reference to billable account
	Account string `json:"account"`
	Href    string `json:"href,omitempty"`
}

// ServiceListRT is the result type of the service service list method.
type ServiceListRT struct {
	// Services
	Items []*ServiceListItem
	// Time at which this list was valid
	AtTime string
	Links  []*LinkT
}

// ServiceStatusRT is the result type of the service service create_service
// method.
type ServiceStatusRT struct {
	// ID
	ID string
	// More detailed description of the service
	Description *string
	// Service status
	Status string
	// Reference to billable account
	Account string `json:"account"`
	Links   []*LinkT
	// Optional provider provided name
	Name *string
	// Optional provider provided tags
	Tags []string
	// Service parameter definitions
	Parameters []*ParameterDefT
}

// Unauthorized access to resource
type UnauthorizedT struct {
}

// UpdatePayload is the payload type of the service service update method.
type UpdatePayload struct {
	// ID of services to update
	ID *string
	// Create if not already exist
	ForceCreate *bool
	// Updated services description
	Services *ServiceDefinitionT
	// JWT used for authentication
	JWT string
}

// Defines the workflow to use to execute this service. Currently supported
// 'types' are 'basic'
// and 'argo'. In case of 'basic', use the 'basic' element for further
// parameters. In the current implementation
// 'opts' is expected to contain the same schema as 'basic'
type WorkflowT struct {
	// Type of workflow
	Type string
	// Type of workflow
	Basic *BasicWorkflowOptsT
	// Defines the workflow using argo's WF schema
	Argo any
	// Type specific options - left for backward compatibility, if possible use
	// type specific elements
	Opts any
}

// Error returns an error description.
func (e *BadRequestT) Error() string {
	return "Bad arguments supplied."
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
func (e *InvalidCredentialsT) Error() string {
	return "Provided credential is not valid."
}

// ErrorName returns "InvalidCredentialsT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InvalidCredentialsT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InvalidCredentialsT".
func (e *InvalidCredentialsT) GoaErrorName() string {
	return "invalid-credential"
}

// Error returns an error description.
func (e *InvalidParameterValue) Error() string {
	return "InvalidParameterValue is the error returned when a parameter has the wrong value."
}

// ErrorName returns "InvalidParameterValue".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InvalidParameterValue) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InvalidParameterValue".
func (e *InvalidParameterValue) GoaErrorName() string {
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

// NewServiceListRT initializes result type ServiceListRT from viewed result
// type ServiceListRT.
func NewServiceListRT(vres *serviceviews.ServiceListRT) *ServiceListRT {
	return newServiceListRT(vres.Projected)
}

// NewViewedServiceListRT initializes viewed result type ServiceListRT from
// result type ServiceListRT using the given view.
func NewViewedServiceListRT(res *ServiceListRT, view string) *serviceviews.ServiceListRT {
	p := newServiceListRTView(res)
	return &serviceviews.ServiceListRT{Projected: p, View: "default"}
}

// newServiceListRT converts projected type ServiceListRT to service type
// ServiceListRT.
func newServiceListRT(vres *serviceviews.ServiceListRTView) *ServiceListRT {
	res := &ServiceListRT{}
	if vres.AtTime != nil {
		res.AtTime = *vres.AtTime
	}
	if vres.Items != nil {
		res.Items = make([]*ServiceListItem, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformServiceviewsServiceListItemViewToServiceListItem(val)
		}
	}
	if vres.Links != nil {
		res.Links = make([]*LinkT, len(vres.Links))
		for i, val := range vres.Links {
			res.Links[i] = transformServiceviewsLinkTViewToLinkT(val)
		}
	}
	return res
}

// newServiceListRTView projects result type ServiceListRT to projected type
// ServiceListRTView using the "default" view.
func newServiceListRTView(res *ServiceListRT) *serviceviews.ServiceListRTView {
	vres := &serviceviews.ServiceListRTView{
		AtTime: &res.AtTime,
	}
	if res.Items != nil {
		vres.Items = make([]*serviceviews.ServiceListItemView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformServiceListItemToServiceviewsServiceListItemView(val)
		}
	} else {
		vres.Items = []*serviceviews.ServiceListItemView{}
	}
	if res.Links != nil {
		vres.Links = make([]*serviceviews.LinkTView, len(res.Links))
		for i, val := range res.Links {
			vres.Links[i] = transformLinkTToServiceviewsLinkTView(val)
		}
	} else {
		vres.Links = []*serviceviews.LinkTView{}
	}
	return vres
}

// transformServiceviewsServiceListItemViewToServiceListItem builds a value of
// type *ServiceListItem from a value of type *serviceviews.ServiceListItemView.
func transformServiceviewsServiceListItemViewToServiceListItem(v *serviceviews.ServiceListItemView) *ServiceListItem {
	if v == nil {
		return nil
	}
	res := &ServiceListItem{
		ID:          *v.ID,
		Name:        v.Name,
		Description: v.Description,
		Account:     *v.Account,
		Href:        *v.Href,
	}

	return res
}

// transformServiceviewsLinkTViewToLinkT builds a value of type *LinkT from a
// value of type *serviceviews.LinkTView.
func transformServiceviewsLinkTViewToLinkT(v *serviceviews.LinkTView) *LinkT {
	if v == nil {
		return nil
	}
	res := &LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}

// transformServiceListItemToServiceviewsServiceListItemView builds a value of
// type *serviceviews.ServiceListItemView from a value of type *ServiceListItem.
func transformServiceListItemToServiceviewsServiceListItemView(v *ServiceListItem) *serviceviews.ServiceListItemView {
	res := &serviceviews.ServiceListItemView{
		ID:          &v.ID,
		Name:        v.Name,
		Description: v.Description,
		Account:     &v.Account,
		Href:        &v.Href,
	}

	return res
}

// transformLinkTToServiceviewsLinkTView builds a value of type
// *serviceviews.LinkTView from a value of type *LinkT.
func transformLinkTToServiceviewsLinkTView(v *LinkT) *serviceviews.LinkTView {
	res := &serviceviews.LinkTView{
		Rel:  &v.Rel,
		Type: &v.Type,
		Href: &v.Href,
	}

	return res
}
