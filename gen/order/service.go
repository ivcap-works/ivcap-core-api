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

package order

import (
	"context"
	"io"

	orderviews "github.com/ivcap-works/ivcap-core-api/gen/order/views"
	"goa.design/goa/v3/security"
)

// Manage the life cycle of an order for CRE services.
type Service interface {
	// Show orders by ID
	Read(context.Context, *ReadPayload) (res *OrderStatusRT, err error)
	// list orders
	List(context.Context, *ListPayload) (res *OrderListRT, err error)
	// Create a new orders and return its status.
	Create(context.Context, *CreatePayload) (res *OrderStatusRT, err error)
	// download order logs
	Logs(context.Context, *LogsPayload) (body io.ReadCloser, err error)
	// top order resources
	Top(context.Context, *TopPayload) (res OrderTopResultItemCollection, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ivcap"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.34"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "order"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"read", "list", "create", "logs", "top"}

// Bad arguments supplied.
type BadRequestT struct {
	// Information message
	Message string
}

// CreatePayload is the payload type of the order service create method.
type CreatePayload struct {
	// New orders description
	Orders *OrderRequestT
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

// ListPayload is the payload type of the order service list method.
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

// LogsPayload is the payload type of the order service logs method.
type LogsPayload struct {
	// From unix time, seconds since 1970-01-01
	From *int64
	// To unix time, seconds since 1970-01-01
	To *int64
	// Reference to order requested
	OrderID string
	// JWT used for authentication
	JWT string
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

type OrderListItem struct {
	// ID
	ID string
	// Optional customer provided name
	Name *string
	// Order status
	Status string
	// DateTime order was placed
	OrderedAt *string
	// DateTime order processing started
	StartedAt *string
	// DateTime order processing finished
	FinishedAt *string
	// Reference to service requested
	Service string `json:"service"`
	// Reference to billable account
	Account string `json:"account"`
	Href    string `json:"href,omitempty"`
}

// OrderListRT is the result type of the order service list method.
type OrderListRT struct {
	// Orders
	Items []*OrderListItem
	// Time at which this list was valid
	AtTime string
	Links  []*LinkT
}

type OrderRequestT struct {
	// Reference to service requested
	Service string `json:"service"`
	// Reference to policy used
	Policy *string `json:"policy"`
	// Optional customer provided name
	Name *string
	// Optional customer provided tags
	Tags []string
	// Service parameters
	Parameters []*ParameterT
}

// OrderStatusRT is the result type of the order service read method.
type OrderStatusRT struct {
	// ID
	ID string
	// Order status
	Status string
	// DateTime order was placed
	OrderedAt *string
	// DateTime order processing started
	StartedAt *string
	// DateTime order processing finished
	FinishedAt *string
	Products   *PartialProductListT
	// Reference to service requested
	Service string `json:"service"`
	// Reference to billable account
	Account string `json:"account"`
	Links   []*LinkT
	// Optional customer provided name
	Name *string
	// Optional customer provided tags
	Tags []string
	// Service parameters
	Parameters []*ParameterT
}

type OrderTopResultItem struct {
	// container
	Container string
	// cpu
	CPU string
	// memory
	Memory string
	// storage
	Storage string
	// ephemeral-storage
	EphemeralStorage string
}

// OrderTopResultItemCollection is the result type of the order service top
// method.
type OrderTopResultItemCollection []*OrderTopResultItem

type ParameterT struct {
	Name  *string
	Value *string
}

type PartialProductListT struct {
	// (Partial) list of products delivered by this order
	Items []*ProductListItemT
	// Links to more products, if there are any
	Links []*LinkT
}

type ProductListItemT struct {
	ID       string
	Name     *string
	Status   string
	MimeType *string `json:"mime-type,omitempty"`
	Size     *int64
	Href     string  `json:"href,omitempty"`
	DataHref *string `json:"dataRef,omitempty"`
}

// ReadPayload is the payload type of the order service read method.
type ReadPayload struct {
	// ID of orders to show
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

// ServiceNotAvailable is the type returned when the service necessary to
// fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// TopPayload is the payload type of the order service top method.
type TopPayload struct {
	// Reference to order requested
	OrderID string
	// JWT used for authentication
	JWT string
}

// Unauthorized access to resource
type UnauthorizedT struct {
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
	return "ServiceNotAvailable is the type returned when the service necessary to fulfil the request is currently not available."
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

// NewOrderListRT initializes result type OrderListRT from viewed result type
// OrderListRT.
func NewOrderListRT(vres *orderviews.OrderListRT) *OrderListRT {
	return newOrderListRT(vres.Projected)
}

// NewViewedOrderListRT initializes viewed result type OrderListRT from result
// type OrderListRT using the given view.
func NewViewedOrderListRT(res *OrderListRT, view string) *orderviews.OrderListRT {
	p := newOrderListRTView(res)
	return &orderviews.OrderListRT{Projected: p, View: "default"}
}

// NewOrderTopResultItemCollection initializes result type
// OrderTopResultItemCollection from viewed result type
// OrderTopResultItemCollection.
func NewOrderTopResultItemCollection(vres orderviews.OrderTopResultItemCollection) OrderTopResultItemCollection {
	return newOrderTopResultItemCollection(vres.Projected)
}

// NewViewedOrderTopResultItemCollection initializes viewed result type
// OrderTopResultItemCollection from result type OrderTopResultItemCollection
// using the given view.
func NewViewedOrderTopResultItemCollection(res OrderTopResultItemCollection, view string) orderviews.OrderTopResultItemCollection {
	p := newOrderTopResultItemCollectionView(res)
	return orderviews.OrderTopResultItemCollection{Projected: p, View: "default"}
}

// newOrderListRT converts projected type OrderListRT to service type
// OrderListRT.
func newOrderListRT(vres *orderviews.OrderListRTView) *OrderListRT {
	res := &OrderListRT{}
	if vres.AtTime != nil {
		res.AtTime = *vres.AtTime
	}
	if vres.Items != nil {
		res.Items = make([]*OrderListItem, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformOrderviewsOrderListItemViewToOrderListItem(val)
		}
	}
	if vres.Links != nil {
		res.Links = make([]*LinkT, len(vres.Links))
		for i, val := range vres.Links {
			res.Links[i] = transformOrderviewsLinkTViewToLinkT(val)
		}
	}
	return res
}

// newOrderListRTView projects result type OrderListRT to projected type
// OrderListRTView using the "default" view.
func newOrderListRTView(res *OrderListRT) *orderviews.OrderListRTView {
	vres := &orderviews.OrderListRTView{
		AtTime: &res.AtTime,
	}
	if res.Items != nil {
		vres.Items = make([]*orderviews.OrderListItemView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformOrderListItemToOrderviewsOrderListItemView(val)
		}
	} else {
		vres.Items = []*orderviews.OrderListItemView{}
	}
	if res.Links != nil {
		vres.Links = make([]*orderviews.LinkTView, len(res.Links))
		for i, val := range res.Links {
			vres.Links[i] = transformLinkTToOrderviewsLinkTView(val)
		}
	} else {
		vres.Links = []*orderviews.LinkTView{}
	}
	return vres
}

// newOrderTopResultItemCollection converts projected type
// OrderTopResultItemCollection to service type OrderTopResultItemCollection.
func newOrderTopResultItemCollection(vres orderviews.OrderTopResultItemCollectionView) OrderTopResultItemCollection {
	res := make(OrderTopResultItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newOrderTopResultItem(n)
	}
	return res
}

// newOrderTopResultItemCollectionView projects result type
// OrderTopResultItemCollection to projected type
// OrderTopResultItemCollectionView using the "default" view.
func newOrderTopResultItemCollectionView(res OrderTopResultItemCollection) orderviews.OrderTopResultItemCollectionView {
	vres := make(orderviews.OrderTopResultItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newOrderTopResultItemView(n)
	}
	return vres
}

// newOrderTopResultItem converts projected type OrderTopResultItem to service
// type OrderTopResultItem.
func newOrderTopResultItem(vres *orderviews.OrderTopResultItemView) *OrderTopResultItem {
	res := &OrderTopResultItem{}
	if vres.Container != nil {
		res.Container = *vres.Container
	}
	if vres.CPU != nil {
		res.CPU = *vres.CPU
	}
	if vres.Memory != nil {
		res.Memory = *vres.Memory
	}
	if vres.Storage != nil {
		res.Storage = *vres.Storage
	}
	if vres.EphemeralStorage != nil {
		res.EphemeralStorage = *vres.EphemeralStorage
	}
	return res
}

// newOrderTopResultItemView projects result type OrderTopResultItem to
// projected type OrderTopResultItemView using the "default" view.
func newOrderTopResultItemView(res *OrderTopResultItem) *orderviews.OrderTopResultItemView {
	vres := &orderviews.OrderTopResultItemView{
		Container:        &res.Container,
		CPU:              &res.CPU,
		Memory:           &res.Memory,
		Storage:          &res.Storage,
		EphemeralStorage: &res.EphemeralStorage,
	}
	return vres
}

// transformOrderviewsOrderListItemViewToOrderListItem builds a value of type
// *OrderListItem from a value of type *orderviews.OrderListItemView.
func transformOrderviewsOrderListItemViewToOrderListItem(v *orderviews.OrderListItemView) *OrderListItem {
	if v == nil {
		return nil
	}
	res := &OrderListItem{
		ID:         *v.ID,
		Name:       v.Name,
		Status:     *v.Status,
		OrderedAt:  v.OrderedAt,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		Service:    *v.Service,
		Account:    *v.Account,
		Href:       *v.Href,
	}

	return res
}

// transformOrderviewsLinkTViewToLinkT builds a value of type *LinkT from a
// value of type *orderviews.LinkTView.
func transformOrderviewsLinkTViewToLinkT(v *orderviews.LinkTView) *LinkT {
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

// transformOrderListItemToOrderviewsOrderListItemView builds a value of type
// *orderviews.OrderListItemView from a value of type *OrderListItem.
func transformOrderListItemToOrderviewsOrderListItemView(v *OrderListItem) *orderviews.OrderListItemView {
	res := &orderviews.OrderListItemView{
		ID:         &v.ID,
		Name:       v.Name,
		Status:     &v.Status,
		OrderedAt:  v.OrderedAt,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		Service:    &v.Service,
		Account:    &v.Account,
		Href:       &v.Href,
	}

	return res
}

// transformLinkTToOrderviewsLinkTView builds a value of type
// *orderviews.LinkTView from a value of type *LinkT.
func transformLinkTToOrderviewsLinkTView(v *LinkT) *orderviews.LinkTView {
	res := &orderviews.LinkTView{
		Rel:  &v.Rel,
		Type: &v.Type,
		Href: &v.Href,
	}

	return res
}
