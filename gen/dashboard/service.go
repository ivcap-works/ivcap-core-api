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

package dashboard

import (
	"context"

	dashboardviews "github.com/ivcap-works/ivcap-core-api/gen/dashboard/views"
	"goa.design/goa/v3/security"
)

// list dashboards
type Service interface {
	// list dashboards
	List(context.Context, *ListPayload) (res *DashboardListRT, err error)
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
const ServiceName = "dashboard"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"list"}

// Something wasn't right with this request
type BadRequestT struct {
	// Information message
	Message string
}

type DashboardListItem struct {
	// dashboard id
	ID int
	// dashboard uid
	UID string
	// Dashboard title
	Title string
	// Dashboard url
	URL string
}

// DashboardListRT is the result type of the dashboard service list method.
type DashboardListRT struct {
	// Dashboards
	Items []*DashboardListItem
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

// ListPayload is the payload type of the dashboard service list method.
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

// Service necessary to fulfil the request is currently not available.
type ServiceNotAvailableT struct {
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

// NewDashboardListRT initializes result type DashboardListRT from viewed
// result type DashboardListRT.
func NewDashboardListRT(vres *dashboardviews.DashboardListRT) *DashboardListRT {
	return newDashboardListRT(vres.Projected)
}

// NewViewedDashboardListRT initializes viewed result type DashboardListRT from
// result type DashboardListRT using the given view.
func NewViewedDashboardListRT(res *DashboardListRT, view string) *dashboardviews.DashboardListRT {
	p := newDashboardListRTView(res)
	return &dashboardviews.DashboardListRT{Projected: p, View: "default"}
}

// newDashboardListRT converts projected type DashboardListRT to service type
// DashboardListRT.
func newDashboardListRT(vres *dashboardviews.DashboardListRTView) *DashboardListRT {
	res := &DashboardListRT{}
	if vres.Items != nil {
		res.Items = make([]*DashboardListItem, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformDashboardviewsDashboardListItemViewToDashboardListItem(val)
		}
	}
	return res
}

// newDashboardListRTView projects result type DashboardListRT to projected
// type DashboardListRTView using the "default" view.
func newDashboardListRTView(res *DashboardListRT) *dashboardviews.DashboardListRTView {
	vres := &dashboardviews.DashboardListRTView{}
	if res.Items != nil {
		vres.Items = make([]*dashboardviews.DashboardListItemView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformDashboardListItemToDashboardviewsDashboardListItemView(val)
		}
	} else {
		vres.Items = []*dashboardviews.DashboardListItemView{}
	}
	return vres
}

// transformDashboardviewsDashboardListItemViewToDashboardListItem builds a
// value of type *DashboardListItem from a value of type
// *dashboardviews.DashboardListItemView.
func transformDashboardviewsDashboardListItemViewToDashboardListItem(v *dashboardviews.DashboardListItemView) *DashboardListItem {
	if v == nil {
		return nil
	}
	res := &DashboardListItem{
		ID:    *v.ID,
		UID:   *v.UID,
		Title: *v.Title,
		URL:   *v.URL,
	}

	return res
}

// transformDashboardListItemToDashboardviewsDashboardListItemView builds a
// value of type *dashboardviews.DashboardListItemView from a value of type
// *DashboardListItem.
func transformDashboardListItemToDashboardviewsDashboardListItemView(v *DashboardListItem) *dashboardviews.DashboardListItemView {
	res := &dashboardviews.DashboardListItemView{
		ID:    &v.ID,
		UID:   &v.UID,
		Title: &v.Title,
		URL:   &v.URL,
	}

	return res
}
