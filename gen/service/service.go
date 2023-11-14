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

package service

import (
	serviceviews "github.com/reinventingscience/ivcap-core-api/gen/service/views"
	"context"

	"goa.design/goa/v3/security"
)

// Manage the life cycle of a service offered on the CRE marketplace.
type Service interface {
	// list services
	List(context.Context, *ListPayload) (res *ServiceListRT, err error)
	// Create a new services and return its status.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	CreateService(context.Context, *CreateServicePayload) (res *ServiceStatusRT, view string, err error)
	// Show services by ID
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Read(context.Context, *ReadPayload) (res *ServiceStatusRT, view string, err error)
	// Update an existing services and return its status.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Update(context.Context, *UpdatePayload) (res *ServiceStatusRT, view string, err error)
	// Delete an existing services.
	Delete(context.Context, *DeletePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

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
}

// CreateServicePayload is the payload type of the service service
// create_service method.
type CreateServicePayload struct {
	// New services description
	Services *ServiceDescriptionT
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

type DescribedByT struct {
	Href *string
	Type *string
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

// ListPayload is the payload type of the service service list method.
type ListPayload struct {
	// The $limit system query option requests the number of items in the queried
	// collection to be included in the result.
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
	// When set order result in descending order. Ascending order is the default.
	OrderDesc bool
	// Return the state of the respective resources at that time [now]
	AtTime *string
	// The content of 'page' is returned in the 'links' part of a previous query and
	// will when set, ALL other parameters, except for 'limit' are ignored.
	Page *string
	// JWT used for authentication
	JWT string
}

type NavT struct {
	Self  *string
	First *string
	Next  *string
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

type ParameterT struct {
	Name  *string
	Value *string
}

// ReadPayload is the payload type of the service service read method.
type ReadPayload struct {
	// ID of services to show
	ID string
	// JWT used for authentication
	JWT string
}

type RefT struct {
	ID    *string
	Links *SelfT
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

type SelfT struct {
	Self        *string
	DescribedBy *DescribedByT
}

type ServiceDescriptionT struct {
	// Provider provided reference. Should to be a single string with punctuations
	// allowed. Might be changed, so please check result
	ProviderRef *string `json:"provider-ref,omitempty"`
	// Reference to service provider
	ProviderID string `json:"provider-id,omitempty"`
	// More detailed description of the service
	Description string
	// Optional provider provided meta tags
	Metadata []*ParameterT
	// Reference to account revenues for this service should be credited to
	References []*ReferenceT
	// Link to banner image oprionally used for this service
	Banner *string
	// Definition of the workflow to use for executing this service
	Workflow *WorkflowT
	// Reference to policy controlling access
	PolicyID *string `json:"policy-id,omitempty"`
	// Optional provider provided name
	Name *string
	// Optional provider provided tags
	Tags []string
	// Service parameter definitions
	Parameters []*ParameterDefT
}

type ServiceListItem struct {
	// Service ID
	ID *string
	// Optional customer provided name
	Name *string
	// Optional description of the service
	Description *string
	// Optional provider link
	Provider *RefT
	Links    *SelfT
}

// ServiceListRT is the result type of the service service list method.
type ServiceListRT struct {
	// Services
	Services []*ServiceListItem
	// Time at which this list was valid
	AtTime string
	// Navigation links
	Links *NavT
}

// ServiceStatusRT is the result type of the service service create_service
// method.
type ServiceStatusRT struct {
	// Service ID
	ID string
	// Provider provided ID. Needs to be a single string with punctuations allowed.
	// Might have been changed
	ProviderRef *string
	// More detailed description of the service
	Description *string
	// Service status
	Status *string
	// Optional provider provided meta tags
	Metadata []*ParameterT
	// Reference to service provider
	Provider *RefT
	// Reference to billable account
	Account *RefT
	Links   *SelfT
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
	Services *ServiceDescriptionT
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
	Type *string
	// Type of workflow
	Basic *BasicWorkflowOptsT
	// Defines the workflow using argo's WF schema
	Argo interface{}
	// Type specific options - left for backward compatibility, if possible use
	// type specific elements
	Opts interface{}
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

// NewServiceStatusRT initializes result type ServiceStatusRT from viewed
// result type ServiceStatusRT.
func NewServiceStatusRT(vres *serviceviews.ServiceStatusRT) *ServiceStatusRT {
	var res *ServiceStatusRT
	switch vres.View {
	case "default", "":
		res = newServiceStatusRT(vres.Projected)
	case "tiny":
		res = newServiceStatusRTTiny(vres.Projected)
	}
	return res
}

// NewViewedServiceStatusRT initializes viewed result type ServiceStatusRT from
// result type ServiceStatusRT using the given view.
func NewViewedServiceStatusRT(res *ServiceStatusRT, view string) *serviceviews.ServiceStatusRT {
	var vres *serviceviews.ServiceStatusRT
	switch view {
	case "default", "":
		p := newServiceStatusRTView(res)
		vres = &serviceviews.ServiceStatusRT{Projected: p, View: "default"}
	case "tiny":
		p := newServiceStatusRTViewTiny(res)
		vres = &serviceviews.ServiceStatusRT{Projected: p, View: "tiny"}
	}
	return vres
}

// newServiceListRT converts projected type ServiceListRT to service type
// ServiceListRT.
func newServiceListRT(vres *serviceviews.ServiceListRTView) *ServiceListRT {
	res := &ServiceListRT{}
	if vres.AtTime != nil {
		res.AtTime = *vres.AtTime
	}
	if vres.Services != nil {
		res.Services = make([]*ServiceListItem, len(vres.Services))
		for i, val := range vres.Services {
			res.Services[i] = transformServiceviewsServiceListItemViewToServiceListItem(val)
		}
	}
	if vres.Links != nil {
		res.Links = transformServiceviewsNavTViewToNavT(vres.Links)
	}
	return res
}

// newServiceListRTView projects result type ServiceListRT to projected type
// ServiceListRTView using the "default" view.
func newServiceListRTView(res *ServiceListRT) *serviceviews.ServiceListRTView {
	vres := &serviceviews.ServiceListRTView{
		AtTime: &res.AtTime,
	}
	if res.Services != nil {
		vres.Services = make([]*serviceviews.ServiceListItemView, len(res.Services))
		for i, val := range res.Services {
			vres.Services[i] = transformServiceListItemToServiceviewsServiceListItemView(val)
		}
	}
	if res.Links != nil {
		vres.Links = transformNavTToServiceviewsNavTView(res.Links)
	}
	return vres
}

// newServiceStatusRT converts projected type ServiceStatusRT to service type
// ServiceStatusRT.
func newServiceStatusRT(vres *serviceviews.ServiceStatusRTView) *ServiceStatusRT {
	res := &ServiceStatusRT{
		Name:        vres.Name,
		Description: vres.Description,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Tags != nil {
		res.Tags = make([]string, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = val
		}
	}
	if vres.Metadata != nil {
		res.Metadata = make([]*ParameterT, len(vres.Metadata))
		for i, val := range vres.Metadata {
			res.Metadata[i] = transformServiceviewsParameterTViewToParameterT(val)
		}
	}
	if vres.Parameters != nil {
		res.Parameters = make([]*ParameterDefT, len(vres.Parameters))
		for i, val := range vres.Parameters {
			res.Parameters[i] = transformServiceviewsParameterDefTViewToParameterDefT(val)
		}
	}
	if vres.Provider != nil {
		res.Provider = transformServiceviewsRefTViewToRefT(vres.Provider)
	}
	if vres.Account != nil {
		res.Account = transformServiceviewsRefTViewToRefT(vres.Account)
	}
	if vres.Links != nil {
		res.Links = transformServiceviewsSelfTViewToSelfT(vres.Links)
	}
	return res
}

// newServiceStatusRTTiny converts projected type ServiceStatusRT to service
// type ServiceStatusRT.
func newServiceStatusRTTiny(vres *serviceviews.ServiceStatusRTView) *ServiceStatusRT {
	res := &ServiceStatusRT{
		Name: vres.Name,
	}
	if vres.Links != nil {
		res.Links = transformServiceviewsSelfTViewToSelfT(vres.Links)
	}
	return res
}

// newServiceStatusRTView projects result type ServiceStatusRT to projected
// type ServiceStatusRTView using the "default" view.
func newServiceStatusRTView(res *ServiceStatusRT) *serviceviews.ServiceStatusRTView {
	vres := &serviceviews.ServiceStatusRTView{
		ID:          &res.ID,
		Description: res.Description,
		Name:        res.Name,
	}
	if res.Metadata != nil {
		vres.Metadata = make([]*serviceviews.ParameterTView, len(res.Metadata))
		for i, val := range res.Metadata {
			vres.Metadata[i] = transformParameterTToServiceviewsParameterTView(val)
		}
	}
	if res.Provider != nil {
		vres.Provider = transformRefTToServiceviewsRefTView(res.Provider)
	}
	if res.Account != nil {
		vres.Account = transformRefTToServiceviewsRefTView(res.Account)
	}
	if res.Links != nil {
		vres.Links = transformSelfTToServiceviewsSelfTView(res.Links)
	}
	if res.Tags != nil {
		vres.Tags = make([]string, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = val
		}
	}
	if res.Parameters != nil {
		vres.Parameters = make([]*serviceviews.ParameterDefTView, len(res.Parameters))
		for i, val := range res.Parameters {
			vres.Parameters[i] = transformParameterDefTToServiceviewsParameterDefTView(val)
		}
	}
	return vres
}

// newServiceStatusRTViewTiny projects result type ServiceStatusRT to projected
// type ServiceStatusRTView using the "tiny" view.
func newServiceStatusRTViewTiny(res *ServiceStatusRT) *serviceviews.ServiceStatusRTView {
	vres := &serviceviews.ServiceStatusRTView{
		Name: res.Name,
	}
	if res.Links != nil {
		vres.Links = transformSelfTToServiceviewsSelfTView(res.Links)
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
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
	}
	if v.Provider != nil {
		res.Provider = transformServiceviewsRefTViewToRefT(v.Provider)
	}
	if v.Links != nil {
		res.Links = transformServiceviewsSelfTViewToSelfT(v.Links)
	}

	return res
}

// transformServiceviewsRefTViewToRefT builds a value of type *RefT from a
// value of type *serviceviews.RefTView.
func transformServiceviewsRefTViewToRefT(v *serviceviews.RefTView) *RefT {
	if v == nil {
		return nil
	}
	res := &RefT{
		ID: v.ID,
	}
	if v.Links != nil {
		res.Links = transformServiceviewsSelfTViewToSelfT(v.Links)
	}

	return res
}

// transformServiceviewsSelfTViewToSelfT builds a value of type *SelfT from a
// value of type *serviceviews.SelfTView.
func transformServiceviewsSelfTViewToSelfT(v *serviceviews.SelfTView) *SelfT {
	if v == nil {
		return nil
	}
	res := &SelfT{
		Self: v.Self,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = transformServiceviewsDescribedByTViewToDescribedByT(v.DescribedBy)
	}

	return res
}

// transformServiceviewsDescribedByTViewToDescribedByT builds a value of type
// *DescribedByT from a value of type *serviceviews.DescribedByTView.
func transformServiceviewsDescribedByTViewToDescribedByT(v *serviceviews.DescribedByTView) *DescribedByT {
	if v == nil {
		return nil
	}
	res := &DescribedByT{
		Href: v.Href,
		Type: v.Type,
	}

	return res
}

// transformServiceviewsNavTViewToNavT builds a value of type *NavT from a
// value of type *serviceviews.NavTView.
func transformServiceviewsNavTViewToNavT(v *serviceviews.NavTView) *NavT {
	if v == nil {
		return nil
	}
	res := &NavT{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}

// transformServiceListItemToServiceviewsServiceListItemView builds a value of
// type *serviceviews.ServiceListItemView from a value of type *ServiceListItem.
func transformServiceListItemToServiceviewsServiceListItemView(v *ServiceListItem) *serviceviews.ServiceListItemView {
	res := &serviceviews.ServiceListItemView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
	}
	if v.Provider != nil {
		res.Provider = transformRefTToServiceviewsRefTView(v.Provider)
	}
	if v.Links != nil {
		res.Links = transformSelfTToServiceviewsSelfTView(v.Links)
	}

	return res
}

// transformRefTToServiceviewsRefTView builds a value of type
// *serviceviews.RefTView from a value of type *RefT.
func transformRefTToServiceviewsRefTView(v *RefT) *serviceviews.RefTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.RefTView{
		ID: v.ID,
	}
	if v.Links != nil {
		res.Links = transformSelfTToServiceviewsSelfTView(v.Links)
	}

	return res
}

// transformSelfTToServiceviewsSelfTView builds a value of type
// *serviceviews.SelfTView from a value of type *SelfT.
func transformSelfTToServiceviewsSelfTView(v *SelfT) *serviceviews.SelfTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.SelfTView{
		Self: v.Self,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = transformDescribedByTToServiceviewsDescribedByTView(v.DescribedBy)
	}

	return res
}

// transformDescribedByTToServiceviewsDescribedByTView builds a value of type
// *serviceviews.DescribedByTView from a value of type *DescribedByT.
func transformDescribedByTToServiceviewsDescribedByTView(v *DescribedByT) *serviceviews.DescribedByTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.DescribedByTView{
		Href: v.Href,
		Type: v.Type,
	}

	return res
}

// transformNavTToServiceviewsNavTView builds a value of type
// *serviceviews.NavTView from a value of type *NavT.
func transformNavTToServiceviewsNavTView(v *NavT) *serviceviews.NavTView {
	res := &serviceviews.NavTView{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}

// transformServiceviewsParameterTViewToParameterT builds a value of type
// *ParameterT from a value of type *serviceviews.ParameterTView.
func transformServiceviewsParameterTViewToParameterT(v *serviceviews.ParameterTView) *ParameterT {
	if v == nil {
		return nil
	}
	res := &ParameterT{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// transformServiceviewsParameterDefTViewToParameterDefT builds a value of type
// *ParameterDefT from a value of type *serviceviews.ParameterDefTView.
func transformServiceviewsParameterDefTViewToParameterDefT(v *serviceviews.ParameterDefTView) *ParameterDefT {
	if v == nil {
		return nil
	}
	res := &ParameterDefT{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
		Unary:       v.Unary,
	}
	if v.Options != nil {
		res.Options = make([]*ParameterOptT, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = transformServiceviewsParameterOptTViewToParameterOptT(val)
		}
	}

	return res
}

// transformServiceviewsParameterOptTViewToParameterOptT builds a value of type
// *ParameterOptT from a value of type *serviceviews.ParameterOptTView.
func transformServiceviewsParameterOptTViewToParameterOptT(v *serviceviews.ParameterOptTView) *ParameterOptT {
	if v == nil {
		return nil
	}
	res := &ParameterOptT{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}

// transformParameterTToServiceviewsParameterTView builds a value of type
// *serviceviews.ParameterTView from a value of type *ParameterT.
func transformParameterTToServiceviewsParameterTView(v *ParameterT) *serviceviews.ParameterTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.ParameterTView{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// transformParameterDefTToServiceviewsParameterDefTView builds a value of type
// *serviceviews.ParameterDefTView from a value of type *ParameterDefT.
func transformParameterDefTToServiceviewsParameterDefTView(v *ParameterDefT) *serviceviews.ParameterDefTView {
	res := &serviceviews.ParameterDefTView{
		Name:        v.Name,
		Label:       v.Label,
		Type:        v.Type,
		Description: v.Description,
		Unit:        v.Unit,
		Constant:    v.Constant,
		Optional:    v.Optional,
		Default:     v.Default,
		Unary:       v.Unary,
	}
	if v.Options != nil {
		res.Options = make([]*serviceviews.ParameterOptTView, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = transformParameterOptTToServiceviewsParameterOptTView(val)
		}
	}

	return res
}

// transformParameterOptTToServiceviewsParameterOptTView builds a value of type
// *serviceviews.ParameterOptTView from a value of type *ParameterOptT.
func transformParameterOptTToServiceviewsParameterOptTView(v *ParameterOptT) *serviceviews.ParameterOptTView {
	if v == nil {
		return nil
	}
	res := &serviceviews.ParameterOptTView{
		Value:       v.Value,
		Description: v.Description,
	}

	return res
}
