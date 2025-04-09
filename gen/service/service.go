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

package service

import (
	"context"
	"io"

	serviceviews "github.com/ivcap-works/ivcap-core-api/gen/service/views"
	"goa.design/goa/v3/security"
)

// Manage the life cycle of a service offered on IVCAP.
type Service interface {
	// list services
	ServiceList(context.Context, *ServiceListPayload) (res *ServiceListRT, err error)
	// Create a new service and return its status.
	ServiceCreate(context.Context, *ServiceCreatePayload) (res *ServiceStatusRT, err error)
	// Show service by ID
	ServiceRead(context.Context, *ServiceReadPayload) (res *ServiceStatusRT, err error)
	// Update an existing service and return its status.
	ServiceUpdate(context.Context, *ServiceUpdatePayload) (res *ServiceStatusRT, err error)
	// Delete an existing service.
	ServiceDelete(context.Context, *ServiceDeletePayload) (err error)
	// list jobs for a specific service
	JobList(context.Context, *JobListPayload) (res *JobListRT, err error)
	// Create a job in the context of a specific service.
	JobCreate(context.Context, *JobCreatePayload, io.ReadCloser) (res *JobCreateResult, body io.ReadCloser, err error)
	// show the status of a job within the context of a service
	JobRead(context.Context, *JobReadPayload) (res *JobStatusRT, err error)
	// Return the result of a job.
	JobOutput(context.Context, *JobOutputPayload) (res *JobOutputResult, body io.ReadCloser, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ivcap"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.44"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [9]string{"service-list", "service-create", "service-read", "service-update", "service-delete", "job-list", "job-create", "job-read", "job-output"}

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

// JobCreatePayload is the payload type of the service service job-create
// method.
type JobCreatePayload struct {
	// ID of service for which to show the list of jobs
	ServiceID     string
	InContentType string
	InOrderID     *string
	ForwardHost   *string
	ForwardProto  *string
	Timeout       *int
	// JWT used for authentication
	JWT string
}

// JobCreateResult is the result type of the service service job-create method.
type JobCreateResult struct {
	OutContentType string
	OutOrderID     string
	JobID          string
	JobURL         *string
}

// Job failed due to an innternal error
type JobInternalErrorT struct {
	// more infomration about the error
	Message string
}

type JobListItem struct {
	// ID
	ID string
	// Optional customer provided name
	Name *string
	// Job status
	Status string
	// DateTime job processing started
	StartedAt *string
	// DateTime job processing finished
	FinishedAt *string
	// Reference to service requested
	Service string
	// Reference to order
	Order *string
	Href  string `json:"href,omitempty"`
}

// JobListPayload is the payload type of the service service job-list method.
type JobListPayload struct {
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
	// ID of service for which to show the list of jobs
	ServiceID string
}

// JobListRT is the result type of the service service job-list method.
type JobListRT struct {
	// Jobs
	Items []*JobListItem
	// Time at which this list was valid
	AtTime string
	Links  []*LinkT
}

// Job did not produce any result
type JobNoResultT struct {
}

// JobOutputPayload is the payload type of the service service job-output
// method.
type JobOutputPayload struct {
	// ID of service for which to show the list of jobs
	ServiceID string
	// ID of job for this service
	JobID string
	// JWT used for authentication
	JWT string
}

// JobOutputResult is the result type of the service service job-output method.
type JobOutputResult struct {
	ContentType string
	OrderID     string
	JobID       string
	JobURL      string
}

// JobReadPayload is the payload type of the service service job-read method.
type JobReadPayload struct {
	// ID of job to show
	ID string
	// JWT used for authentication
	JWT string
	// ID of service for which to show the list of jobs
	ServiceID string
	// include request content if possible
	WithRequestContent *bool
	// include result content if possible
	WithResultContent *bool
}

// Job failed because the requested parameters were incorrect
type JobRequestErrorT struct {
	// more infomration about the error
	Message string
}

// The information returned if the job hasn't finished yet
type JobRetryLaterT struct {
	// the ID of the job
	JobID *string
	// the URL for the job
	Location string
	// The time in seconds after which an update may be available
	RetryLater int
}

// JobStatusRT is the result type of the service service job-read method.
type JobStatusRT struct {
	// ID
	ID string
	// Job status
	Status string
	// Optional customer provided name
	Name *string
	// Optional customer provided tags
	Tags []string
	// Reference to order
	Order string
	// Reference to service requested
	Service string
	// Mime type of request
	RequestContentType string
	// Request content
	RequestContent any
	// Mime type of result
	ResultContentType *string
	// Result content
	ResultContent any
	Products      *PartialProductList2T
	// Additional error message id status is 'Error' or 'Failed'
	ErrorMessage *string
	// Reference to billable account
	Account string
	// Reference to policy used
	Policy string
	// DateTime job was submitted
	RequestedAt string
	// DateTime job processing started
	StartedAt *string
	// DateTime job processing finished
	FinishedAt *string
	Links      []*LinkT
}

type LinkT struct {
	// relation type
	Rel string
	// mime type
	Type string
	// web link
	Href string
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

type ParameterDefT struct {
	Name        string
	Label       *string
	Type        string
	Description string
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

type PartialProductList2T struct {
	// (Partial) list of products delivered by this order
	Items []*ProductListItem2T
	// Links to more products, if there are any
	Links []*LinkT
}

type ProductListItem2T struct {
	ID       string
	Name     *string
	Status   string
	MimeType *string `json:"mime-type,omitempty"`
	Size     *int64
	Href     string  `json:"href,omitempty"`
	DataHref *string `json:"dataRef,omitempty"`
}

// Will be returned when receiving a request to create and already existing
// resource.
type ResourceAlreadyCreatedT struct {
	// ID of already existing resource
	ID string
	// Message of error
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

// ServiceCreatePayload is the payload type of the service service
// service-create method.
type ServiceCreatePayload struct {
	// New service description
	Service *ServiceDefinitionT
	// JWT used for authentication
	JWT string
}

type ServiceDefinitionT struct {
	// type of controller used for this service
	ControllerSchema string
	// controller definition
	Controller any
	// Reference to policy used
	Policy string
	// ID
	ID string
	// Optional provider provided name
	Name *string
	// More detailed description of the service
	Description string
	// Optional tags defined for service to help in categorising them
	Tags []string
	// Service parameter definitions
	Parameters []*ParameterDefT
}

// ServiceDeletePayload is the payload type of the service service
// service-delete method.
type ServiceDeletePayload struct {
	// ID of service to update
	ID string
	// JWT used for authentication
	JWT string
}

type ServiceListItemT struct {
	// ID
	ID string
	// Optional customer provided name
	Name *string
	// Optional description of the service
	Description *string
	// Optional tags defined for service to help in categorising them
	Tags []string
	// type of controller used for this service
	ControllerSchema string
	// time this service has been available from
	ValidFrom *string
	// time this service has been available to
	ValidTo *string
	Href    string `json:"href,omitempty"`
}

// ServiceListPayload is the payload type of the service service service-list
// method.
type ServiceListPayload struct {
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

// ServiceListRT is the result type of the service service service-list method.
type ServiceListRT struct {
	// Services
	Items []*ServiceListItemT
	// Time at which this list was valid
	AtTime string
	Links  []*LinkT
}

// Service necessary to fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// ServiceReadPayload is the payload type of the service service service-read
// method.
type ServiceReadPayload struct {
	// ID of service to show
	ID string
	// JWT used for authentication
	JWT string
}

// ServiceStatusRT is the result type of the service service service-create
// method.
type ServiceStatusRT struct {
	// Service status
	Status string
	// type of controller used for this service
	ControllerSchema string
	// controller definition
	Controller any
	// Reference to billable account
	Account string
	// Reference to policy used
	Policy string
	// time this service has been available from
	ValidFrom *string
	// time this service has been available to
	ValidTo *string
	Links   []*LinkT
	// ID
	ID string
	// Optional provider provided name
	Name *string
	// More detailed description of the service
	Description string
	// Optional tags defined for service to help in categorising them
	Tags []string
	// Service parameter definitions
	Parameters []*ParameterDefT
}

// ServiceUpdatePayload is the payload type of the service service
// service-update method.
type ServiceUpdatePayload struct {
	// ID of service to update
	ID *string
	// Create if not already exist
	ForceCreate *bool
	// Updated service description
	Service *ServiceDefinitionT
	// JWT used for authentication
	JWT string
}

// Temporarily redirecting to a different URL
type TemporaryRedirectT struct {
	// the URL for the job
	Location string
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
func (e *JobInternalErrorT) Error() string {
	return "Job failed due to an innternal error"
}

// ErrorName returns "JobInternalErrorT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *JobInternalErrorT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "JobInternalErrorT".
func (e *JobInternalErrorT) GoaErrorName() string {
	return "job-internal-error"
}

// Error returns an error description.
func (e *JobNoResultT) Error() string {
	return "Job did not produce any result"
}

// ErrorName returns "JobNoResultT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *JobNoResultT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "JobNoResultT".
func (e *JobNoResultT) GoaErrorName() string {
	return "job-no-result"
}

// Error returns an error description.
func (e *JobRequestErrorT) Error() string {
	return "Job failed because the requested parameters were incorrect"
}

// ErrorName returns "JobRequestErrorT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *JobRequestErrorT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "JobRequestErrorT".
func (e *JobRequestErrorT) GoaErrorName() string {
	return "job-request-error"
}

// Error returns an error description.
func (e *JobRetryLaterT) Error() string {
	return "The information returned if the job hasn't finished yet"
}

// ErrorName returns "JobRetryLaterT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *JobRetryLaterT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "JobRetryLaterT".
func (e *JobRetryLaterT) GoaErrorName() string {
	return "not-ready-yet"
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
func (e *TemporaryRedirectT) Error() string {
	return "Temporarily redirecting to a different URL"
}

// ErrorName returns "TemporaryRedirectT".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *TemporaryRedirectT) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "TemporaryRedirectT".
func (e *TemporaryRedirectT) GoaErrorName() string {
	return "temporary-redirect"
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

// NewJobListRT initializes result type JobListRT from viewed result type
// JobListRT.
func NewJobListRT(vres *serviceviews.JobListRT) *JobListRT {
	return newJobListRT(vres.Projected)
}

// NewViewedJobListRT initializes viewed result type JobListRT from result type
// JobListRT using the given view.
func NewViewedJobListRT(res *JobListRT, view string) *serviceviews.JobListRT {
	p := newJobListRTView(res)
	return &serviceviews.JobListRT{Projected: p, View: "default"}
}

// newServiceListRT converts projected type ServiceListRT to service type
// ServiceListRT.
func newServiceListRT(vres *serviceviews.ServiceListRTView) *ServiceListRT {
	res := &ServiceListRT{}
	if vres.AtTime != nil {
		res.AtTime = *vres.AtTime
	}
	if vres.Items != nil {
		res.Items = make([]*ServiceListItemT, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformServiceviewsServiceListItemTViewToServiceListItemT(val)
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
		vres.Items = make([]*serviceviews.ServiceListItemTView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformServiceListItemTToServiceviewsServiceListItemTView(val)
		}
	} else {
		vres.Items = []*serviceviews.ServiceListItemTView{}
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

// newJobListRT converts projected type JobListRT to service type JobListRT.
func newJobListRT(vres *serviceviews.JobListRTView) *JobListRT {
	res := &JobListRT{}
	if vres.AtTime != nil {
		res.AtTime = *vres.AtTime
	}
	if vres.Items != nil {
		res.Items = make([]*JobListItem, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformServiceviewsJobListItemViewToJobListItem(val)
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

// newJobListRTView projects result type JobListRT to projected type
// JobListRTView using the "default" view.
func newJobListRTView(res *JobListRT) *serviceviews.JobListRTView {
	vres := &serviceviews.JobListRTView{
		AtTime: &res.AtTime,
	}
	if res.Items != nil {
		vres.Items = make([]*serviceviews.JobListItemView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformJobListItemToServiceviewsJobListItemView(val)
		}
	} else {
		vres.Items = []*serviceviews.JobListItemView{}
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

// transformServiceviewsServiceListItemTViewToServiceListItemT builds a value
// of type *ServiceListItemT from a value of type
// *serviceviews.ServiceListItemTView.
func transformServiceviewsServiceListItemTViewToServiceListItemT(v *serviceviews.ServiceListItemTView) *ServiceListItemT {
	if v == nil {
		return nil
	}
	res := &ServiceListItemT{
		ID:               *v.ID,
		Name:             v.Name,
		Description:      v.Description,
		ControllerSchema: *v.ControllerSchema,
		ValidFrom:        v.ValidFrom,
		ValidTo:          v.ValidTo,
		Href:             *v.Href,
	}
	if v.Tags != nil {
		res.Tags = make([]string, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = val
		}
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

// transformServiceListItemTToServiceviewsServiceListItemTView builds a value
// of type *serviceviews.ServiceListItemTView from a value of type
// *ServiceListItemT.
func transformServiceListItemTToServiceviewsServiceListItemTView(v *ServiceListItemT) *serviceviews.ServiceListItemTView {
	res := &serviceviews.ServiceListItemTView{
		ID:               &v.ID,
		Name:             v.Name,
		Description:      v.Description,
		ControllerSchema: &v.ControllerSchema,
		ValidFrom:        v.ValidFrom,
		ValidTo:          v.ValidTo,
		Href:             &v.Href,
	}
	if v.Tags != nil {
		res.Tags = make([]string, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = val
		}
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

// transformServiceviewsJobListItemViewToJobListItem builds a value of type
// *JobListItem from a value of type *serviceviews.JobListItemView.
func transformServiceviewsJobListItemViewToJobListItem(v *serviceviews.JobListItemView) *JobListItem {
	if v == nil {
		return nil
	}
	res := &JobListItem{
		ID:         *v.ID,
		Name:       v.Name,
		Status:     *v.Status,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		Service:    *v.Service,
		Order:      v.Order,
		Href:       *v.Href,
	}

	return res
}

// transformJobListItemToServiceviewsJobListItemView builds a value of type
// *serviceviews.JobListItemView from a value of type *JobListItem.
func transformJobListItemToServiceviewsJobListItemView(v *JobListItem) *serviceviews.JobListItemView {
	res := &serviceviews.JobListItemView{
		ID:         &v.ID,
		Name:       v.Name,
		Status:     &v.Status,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		Service:    &v.Service,
		Order:      v.Order,
		Href:       &v.Href,
	}

	return res
}
