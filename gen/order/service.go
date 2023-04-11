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

package order

import (
	orderviews "github.com/reinventingscience/ivcap-core-api/gen/order/views"
	"context"

	"goa.design/goa/v3/security"
)

// Manage the life cycle of an order for CRE services.
type Service interface {
	// orders
	List(context.Context, *ListPayload) (res *OrderListRT, err error)
	// Create a new orders and return its status.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Create(context.Context, *CreatePayload) (res *OrderStatusRT, view string, err error)
	// Show orders by ID
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Read(context.Context, *ReadPayload) (res *OrderStatusRT, view string, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "order"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list", "create", "read"}

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

// ListPayload is the payload type of the order service list method.
type ListPayload struct {
	// The $filter system query option allows clients to filter a collection of
	// resources that are addressed by a request URL. The expression specified with
	// $filter
	// is evaluated for each resource in the collection, and only items where the
	// expression
	// evaluates to true are included in the response.
	Filter string
	// The $orderby query option allows clients to request resources in either
	// ascending order using asc or descending order using desc. If asc or desc not
	// specified,
	// then the resources will be ordered in ascending order. The request below
	// orders Trips on
	// property EndsAt in descending order.
	Orderby string
	// The $top system query option requests the number of items in the queried
	// collection to be included in the result.
	Top int
	// The $skip query option requests the number of items in the queried collection
	// that are to be skipped and not included in the result.
	Skip int
	// The $select system query option allows the clients to requests a limited set
	// of properties for each entity or complex type. The example returns Name and
	// IcaoCode
	// of all Airports.
	Select string
	// DEPRECATED: List offset. Use '$skip' instead
	Offset *int
	// DEPRECATED: Max. number of records to return. Use '$top' instead
	Limit *int
	// DEPRECATED: Page token
	PageToken string
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

type OrderListItem struct {
	// Order ID
	ID *string
	// Optional customer provided name
	Name *string
	// Order status
	Status *string
	// DateTime order was placed
	OrderedAt *string
	// DateTime processing of order started
	StartedAt *string
	// DateTime order was finished
	FinishedAt *string
	// ID of ordered service
	ServiceID *string
	// ID of ordered service
	AccountID *string
	Links     *SelfT
}

// OrderListRT is the result type of the order service list method.
type OrderListRT struct {
	// Orders
	Orders []*OrderListItem
	// Navigation links
	Links *NavT
}

type OrderRequestT struct {
	// Reference to service requested
	ServiceID string
	// Reference to billable account
	AccountID string
	// Optional customer provided name
	Name *string
	// Service parameters
	Parameters []*ParameterT
}

// OrderStatusRT is the result type of the order service create method.
type OrderStatusRT struct {
	// Order ID
	ID string
	// Order status
	Status *string
	// DateTime order was placed
	OrderedAt *string
	// DateTime order processing started
	StartedAt *string
	// DateTime order processing finished
	FinishedAt *string
	// Products delivered for this order
	Products []*ProductT
	// Reference to service requested
	Service *RefT
	// Reference to billable account
	Account *RefT
	Links   *SelfT
	// Optional customer provided name
	Name *string
	// Service parameters
	Parameters []*ParameterT
}

type ParameterT struct {
	Name  *string
	Value *string
}

type ProductT struct {
	ID       *string
	Name     *string
	Status   *string
	MimeType *string
	Size     *int64
	Links    *SelfWithDataT
}

// ReadPayload is the payload type of the order service read method.
type ReadPayload struct {
	// ID of orders to show
	ID string
	// JWT used for authentication
	JWT string
}

type RefT struct {
	ID    *string
	Links *SelfT
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

type SelfWithDataT struct {
	Self        *string
	DescribedBy *DescribedByT
	Data        *string
}

// ServiceNotAvailable is the type returned when the service necessary to
// fulfill the request is currently not available.
type ServiceNotAvailableT struct {
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
	return "ServiceNotAvailable is the type returned when the service necessary to fulfill the request is currently not available."
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

// NewOrderStatusRT initializes result type OrderStatusRT from viewed result
// type OrderStatusRT.
func NewOrderStatusRT(vres *orderviews.OrderStatusRT) *OrderStatusRT {
	var res *OrderStatusRT
	switch vres.View {
	case "default", "":
		res = newOrderStatusRT(vres.Projected)
	case "tiny":
		res = newOrderStatusRTTiny(vres.Projected)
	}
	return res
}

// NewViewedOrderStatusRT initializes viewed result type OrderStatusRT from
// result type OrderStatusRT using the given view.
func NewViewedOrderStatusRT(res *OrderStatusRT, view string) *orderviews.OrderStatusRT {
	var vres *orderviews.OrderStatusRT
	switch view {
	case "default", "":
		p := newOrderStatusRTView(res)
		vres = &orderviews.OrderStatusRT{Projected: p, View: "default"}
	case "tiny":
		p := newOrderStatusRTViewTiny(res)
		vres = &orderviews.OrderStatusRT{Projected: p, View: "tiny"}
	}
	return vres
}

// newOrderListRT converts projected type OrderListRT to service type
// OrderListRT.
func newOrderListRT(vres *orderviews.OrderListRTView) *OrderListRT {
	res := &OrderListRT{}
	if vres.Orders != nil {
		res.Orders = make([]*OrderListItem, len(vres.Orders))
		for i, val := range vres.Orders {
			res.Orders[i] = transformOrderviewsOrderListItemViewToOrderListItem(val)
		}
	}
	if vres.Links != nil {
		res.Links = transformOrderviewsNavTViewToNavT(vres.Links)
	}
	return res
}

// newOrderListRTView projects result type OrderListRT to projected type
// OrderListRTView using the "default" view.
func newOrderListRTView(res *OrderListRT) *orderviews.OrderListRTView {
	vres := &orderviews.OrderListRTView{}
	if res.Orders != nil {
		vres.Orders = make([]*orderviews.OrderListItemView, len(res.Orders))
		for i, val := range res.Orders {
			vres.Orders[i] = transformOrderListItemToOrderviewsOrderListItemView(val)
		}
	}
	if res.Links != nil {
		vres.Links = transformNavTToOrderviewsNavTView(res.Links)
	}
	return vres
}

// newOrderStatusRT converts projected type OrderStatusRT to service type
// OrderStatusRT.
func newOrderStatusRT(vres *orderviews.OrderStatusRTView) *OrderStatusRT {
	res := &OrderStatusRT{
		Name:       vres.Name,
		Status:     vres.Status,
		OrderedAt:  vres.OrderedAt,
		StartedAt:  vres.StartedAt,
		FinishedAt: vres.FinishedAt,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Parameters != nil {
		res.Parameters = make([]*ParameterT, len(vres.Parameters))
		for i, val := range vres.Parameters {
			res.Parameters[i] = transformOrderviewsParameterTViewToParameterT(val)
		}
	}
	if vres.Products != nil {
		res.Products = make([]*ProductT, len(vres.Products))
		for i, val := range vres.Products {
			res.Products[i] = transformOrderviewsProductTViewToProductT(val)
		}
	}
	if vres.Service != nil {
		res.Service = transformOrderviewsRefTViewToRefT(vres.Service)
	}
	if vres.Account != nil {
		res.Account = transformOrderviewsRefTViewToRefT(vres.Account)
	}
	if vres.Links != nil {
		res.Links = transformOrderviewsSelfTViewToSelfT(vres.Links)
	}
	return res
}

// newOrderStatusRTTiny converts projected type OrderStatusRT to service type
// OrderStatusRT.
func newOrderStatusRTTiny(vres *orderviews.OrderStatusRTView) *OrderStatusRT {
	res := &OrderStatusRT{
		Name:   vres.Name,
		Status: vres.Status,
	}
	if vres.Links != nil {
		res.Links = transformOrderviewsSelfTViewToSelfT(vres.Links)
	}
	return res
}

// newOrderStatusRTView projects result type OrderStatusRT to projected type
// OrderStatusRTView using the "default" view.
func newOrderStatusRTView(res *OrderStatusRT) *orderviews.OrderStatusRTView {
	vres := &orderviews.OrderStatusRTView{
		ID:         &res.ID,
		Status:     res.Status,
		OrderedAt:  res.OrderedAt,
		StartedAt:  res.StartedAt,
		FinishedAt: res.FinishedAt,
		Name:       res.Name,
	}
	if res.Products != nil {
		vres.Products = make([]*orderviews.ProductTView, len(res.Products))
		for i, val := range res.Products {
			vres.Products[i] = transformProductTToOrderviewsProductTView(val)
		}
	}
	if res.Service != nil {
		vres.Service = transformRefTToOrderviewsRefTView(res.Service)
	}
	if res.Account != nil {
		vres.Account = transformRefTToOrderviewsRefTView(res.Account)
	}
	if res.Links != nil {
		vres.Links = transformSelfTToOrderviewsSelfTView(res.Links)
	}
	if res.Parameters != nil {
		vres.Parameters = make([]*orderviews.ParameterTView, len(res.Parameters))
		for i, val := range res.Parameters {
			vres.Parameters[i] = transformParameterTToOrderviewsParameterTView(val)
		}
	}
	return vres
}

// newOrderStatusRTViewTiny projects result type OrderStatusRT to projected
// type OrderStatusRTView using the "tiny" view.
func newOrderStatusRTViewTiny(res *OrderStatusRT) *orderviews.OrderStatusRTView {
	vres := &orderviews.OrderStatusRTView{
		Status: res.Status,
		Name:   res.Name,
	}
	if res.Links != nil {
		vres.Links = transformSelfTToOrderviewsSelfTView(res.Links)
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
		ID:         v.ID,
		Name:       v.Name,
		Status:     v.Status,
		OrderedAt:  v.OrderedAt,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		ServiceID:  v.ServiceID,
		AccountID:  v.AccountID,
	}
	if v.Links != nil {
		res.Links = transformOrderviewsSelfTViewToSelfT(v.Links)
	}

	return res
}

// transformOrderviewsSelfTViewToSelfT builds a value of type *SelfT from a
// value of type *orderviews.SelfTView.
func transformOrderviewsSelfTViewToSelfT(v *orderviews.SelfTView) *SelfT {
	res := &SelfT{
		Self: v.Self,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = transformOrderviewsDescribedByTViewToDescribedByT(v.DescribedBy)
	}

	return res
}

// transformOrderviewsDescribedByTViewToDescribedByT builds a value of type
// *DescribedByT from a value of type *orderviews.DescribedByTView.
func transformOrderviewsDescribedByTViewToDescribedByT(v *orderviews.DescribedByTView) *DescribedByT {
	if v == nil {
		return nil
	}
	res := &DescribedByT{
		Href: v.Href,
		Type: v.Type,
	}

	return res
}

// transformOrderviewsNavTViewToNavT builds a value of type *NavT from a value
// of type *orderviews.NavTView.
func transformOrderviewsNavTViewToNavT(v *orderviews.NavTView) *NavT {
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

// transformOrderListItemToOrderviewsOrderListItemView builds a value of type
// *orderviews.OrderListItemView from a value of type *OrderListItem.
func transformOrderListItemToOrderviewsOrderListItemView(v *OrderListItem) *orderviews.OrderListItemView {
	res := &orderviews.OrderListItemView{
		ID:         v.ID,
		Name:       v.Name,
		Status:     v.Status,
		OrderedAt:  v.OrderedAt,
		StartedAt:  v.StartedAt,
		FinishedAt: v.FinishedAt,
		ServiceID:  v.ServiceID,
		AccountID:  v.AccountID,
	}
	if v.Links != nil {
		res.Links = transformSelfTToOrderviewsSelfTView(v.Links)
	}

	return res
}

// transformSelfTToOrderviewsSelfTView builds a value of type
// *orderviews.SelfTView from a value of type *SelfT.
func transformSelfTToOrderviewsSelfTView(v *SelfT) *orderviews.SelfTView {
	res := &orderviews.SelfTView{
		Self: v.Self,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = transformDescribedByTToOrderviewsDescribedByTView(v.DescribedBy)
	}

	return res
}

// transformDescribedByTToOrderviewsDescribedByTView builds a value of type
// *orderviews.DescribedByTView from a value of type *DescribedByT.
func transformDescribedByTToOrderviewsDescribedByTView(v *DescribedByT) *orderviews.DescribedByTView {
	if v == nil {
		return nil
	}
	res := &orderviews.DescribedByTView{
		Href: v.Href,
		Type: v.Type,
	}

	return res
}

// transformNavTToOrderviewsNavTView builds a value of type
// *orderviews.NavTView from a value of type *NavT.
func transformNavTToOrderviewsNavTView(v *NavT) *orderviews.NavTView {
	res := &orderviews.NavTView{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}

// transformOrderviewsParameterTViewToParameterT builds a value of type
// *ParameterT from a value of type *orderviews.ParameterTView.
func transformOrderviewsParameterTViewToParameterT(v *orderviews.ParameterTView) *ParameterT {
	if v == nil {
		return nil
	}
	res := &ParameterT{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}

// transformOrderviewsProductTViewToProductT builds a value of type *ProductT
// from a value of type *orderviews.ProductTView.
func transformOrderviewsProductTViewToProductT(v *orderviews.ProductTView) *ProductT {
	if v == nil {
		return nil
	}
	res := &ProductT{
		ID:       v.ID,
		Name:     v.Name,
		Status:   v.Status,
		MimeType: v.MimeType,
		Size:     v.Size,
	}
	if v.Links != nil {
		res.Links = transformOrderviewsSelfWithDataTViewToSelfWithDataT(v.Links)
	}

	return res
}

// transformOrderviewsSelfWithDataTViewToSelfWithDataT builds a value of type
// *SelfWithDataT from a value of type *orderviews.SelfWithDataTView.
func transformOrderviewsSelfWithDataTViewToSelfWithDataT(v *orderviews.SelfWithDataTView) *SelfWithDataT {
	if v == nil {
		return nil
	}
	res := &SelfWithDataT{
		Self: v.Self,
		Data: v.Data,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = transformOrderviewsDescribedByTViewToDescribedByT(v.DescribedBy)
	}

	return res
}

// transformOrderviewsRefTViewToRefT builds a value of type *RefT from a value
// of type *orderviews.RefTView.
func transformOrderviewsRefTViewToRefT(v *orderviews.RefTView) *RefT {
	if v == nil {
		return nil
	}
	res := &RefT{
		ID: v.ID,
	}
	if v.Links != nil {
		res.Links = transformOrderviewsSelfTViewToSelfT(v.Links)
	}

	return res
}

// transformProductTToOrderviewsProductTView builds a value of type
// *orderviews.ProductTView from a value of type *ProductT.
func transformProductTToOrderviewsProductTView(v *ProductT) *orderviews.ProductTView {
	if v == nil {
		return nil
	}
	res := &orderviews.ProductTView{
		ID:       v.ID,
		Name:     v.Name,
		Status:   v.Status,
		MimeType: v.MimeType,
		Size:     v.Size,
	}
	if v.Links != nil {
		res.Links = transformSelfWithDataTToOrderviewsSelfWithDataTView(v.Links)
	}

	return res
}

// transformSelfWithDataTToOrderviewsSelfWithDataTView builds a value of type
// *orderviews.SelfWithDataTView from a value of type *SelfWithDataT.
func transformSelfWithDataTToOrderviewsSelfWithDataTView(v *SelfWithDataT) *orderviews.SelfWithDataTView {
	if v == nil {
		return nil
	}
	res := &orderviews.SelfWithDataTView{
		Self: v.Self,
		Data: v.Data,
	}
	if v.DescribedBy != nil {
		res.DescribedBy = transformDescribedByTToOrderviewsDescribedByTView(v.DescribedBy)
	}

	return res
}

// transformRefTToOrderviewsRefTView builds a value of type
// *orderviews.RefTView from a value of type *RefT.
func transformRefTToOrderviewsRefTView(v *RefT) *orderviews.RefTView {
	if v == nil {
		return nil
	}
	res := &orderviews.RefTView{
		ID: v.ID,
	}
	if v.Links != nil {
		res.Links = transformSelfTToOrderviewsSelfTView(v.Links)
	}

	return res
}

// transformParameterTToOrderviewsParameterTView builds a value of type
// *orderviews.ParameterTView from a value of type *ParameterT.
func transformParameterTToOrderviewsParameterTView(v *ParameterT) *orderviews.ParameterTView {
	res := &orderviews.ParameterTView{
		Name:  v.Name,
		Value: v.Value,
	}

	return res
}
