// Copyright 2023 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	order "github.com/reinventingscience/ivcap-core-api/gen/order"
	orderviews "github.com/reinventingscience/ivcap-core-api/gen/order/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "order" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Reference to service requested
	ServiceID string `form:"service-id" json:"service-id" xml:"service-id"`
	// Reference to billable account
	AccountID *string `form:"account-id,omitempty" json:"account-id,omitempty" xml:"account-id,omitempty"`
	// Policy to control access to record an all generated artifacts
	PolicyID *string `form:"policy-id,omitempty" json:"policy-id,omitempty" xml:"policy-id,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional customer provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameters
	Parameters []*ParameterT `form:"parameters" json:"parameters" xml:"parameters"`
}

// LogsRequestBody is the type of the "order" service "logs" endpoint HTTP
// request body.
type LogsRequestBody struct {
	// From unix time, seconds since 1970-01-01
	From *int64 `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	// To unix time, seconds since 1970-01-01
	To *int64 `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`
	// Reference to namespace name
	NamespaceName *string `form:"namespace-name,omitempty" json:"namespace-name,omitempty" xml:"namespace-name,omitempty"`
	// Reference to container name
	ContainerName *string `form:"container-name,omitempty" json:"container-name,omitempty" xml:"container-name,omitempty"`
	// Reference to order requested
	OrderID string `form:"order-id" json:"order-id" xml:"order-id"`
	// Policy to control access to record an all generated artifacts
	PolicyID *string `form:"policy-id,omitempty" json:"policy-id,omitempty" xml:"policy-id,omitempty"`
}

// ReadResponseBody is the type of the "order" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// Order ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Order status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime order was placed
	OrderedAt *string `form:"ordered-at,omitempty" json:"ordered-at,omitempty" xml:"ordered-at,omitempty"`
	// DateTime order processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime order processing finished
	FinishedAt *string `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	// Products delivered for this order
	Products []*ProductTResponseBody `form:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
	// Reference to service requested
	Service *RefTResponseBody `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Reference to billable account
	Account *RefTResponseBody  `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	Links   *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional customer provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameters
	Parameters []*ParameterTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// ListResponseBody is the type of the "order" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Orders
	Orders []*OrderListItemResponseBody `form:"orders,omitempty" json:"orders,omitempty" xml:"orders,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	// Navigation links
	Links *NavTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// CreateResponseBody is the type of the "order" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// Order ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Order status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime order was placed
	OrderedAt *string `form:"ordered-at,omitempty" json:"ordered-at,omitempty" xml:"ordered-at,omitempty"`
	// DateTime order processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime order processing finished
	FinishedAt *string `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	// Products delivered for this order
	Products []*ProductTResponseBody `form:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
	// Reference to service requested
	Service *RefTResponseBody `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Reference to billable account
	Account *RefTResponseBody  `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	Links   *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional customer provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameters
	Parameters []*ParameterTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "order" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "order" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "order" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "order" service "read" endpoint
// HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListBadRequestResponseBody is the type of the "order" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "order" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "order" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "order" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "order" service "create"
// endpoint HTTP response body for the "bad-request" error.
type CreateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateInvalidParameterResponseBody is the type of the "order" service
// "create" endpoint HTTP response body for the "invalid-parameter" error.
type CreateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// CreateInvalidScopesResponseBody is the type of the "order" service "create"
// endpoint HTTP response body for the "invalid-scopes" error.
type CreateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateNotImplementedResponseBody is the type of the "order" service "create"
// endpoint HTTP response body for the "not-implemented" error.
type CreateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateNotFoundResponseBody is the type of the "order" service "create"
// endpoint HTTP response body for the "not-found" error.
type CreateNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LogsBadRequestResponseBody is the type of the "order" service "logs"
// endpoint HTTP response body for the "bad-request" error.
type LogsBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LogsInvalidParameterResponseBody is the type of the "order" service "logs"
// endpoint HTTP response body for the "invalid-parameter" error.
type LogsInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// LogsInvalidScopesResponseBody is the type of the "order" service "logs"
// endpoint HTTP response body for the "invalid-scopes" error.
type LogsInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LogsNotImplementedResponseBody is the type of the "order" service "logs"
// endpoint HTTP response body for the "not-implemented" error.
type LogsNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LogsNotFoundResponseBody is the type of the "order" service "logs" endpoint
// HTTP response body for the "not-found" error.
type LogsNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ProductTResponseBody is used to define fields on response body types.
type ProductTResponseBody struct {
	ID       *string                    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name     *string                    `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Status   *string                    `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	MimeType *string                    `form:"mime-type,omitempty" json:"mime-type,omitempty" xml:"mime-type,omitempty"`
	Size     *int64                     `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	Links    *SelfWithDataTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// SelfWithDataTResponseBody is used to define fields on response body types.
type SelfWithDataTResponseBody struct {
	Self        *string                   `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	DescribedBy *DescribedByTResponseBody `form:"describedBy,omitempty" json:"describedBy,omitempty" xml:"describedBy,omitempty"`
	Data        *string                   `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// DescribedByTResponseBody is used to define fields on response body types.
type DescribedByTResponseBody struct {
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// RefTResponseBody is used to define fields on response body types.
type RefTResponseBody struct {
	ID    *string            `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// SelfTResponseBody is used to define fields on response body types.
type SelfTResponseBody struct {
	Self        *string                   `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	DescribedBy *DescribedByTResponseBody `form:"describedBy,omitempty" json:"describedBy,omitempty" xml:"describedBy,omitempty"`
}

// ParameterTResponseBody is used to define fields on response body types.
type ParameterTResponseBody struct {
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// OrderListItemResponseBody is used to define fields on response body types.
type OrderListItemResponseBody struct {
	// Order ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Order status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime order was placed
	OrderedAt *string `form:"ordered-at,omitempty" json:"ordered-at,omitempty" xml:"ordered-at,omitempty"`
	// DateTime processing of order started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime order was finished
	FinishedAt *string `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	// ID of ordered service
	ServiceID *string `form:"service-id,omitempty" json:"service-id,omitempty" xml:"service-id,omitempty"`
	// ID of ordered service
	AccountID *string            `form:"account-id,omitempty" json:"account-id,omitempty" xml:"account-id,omitempty"`
	Links     *SelfTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// NavTResponseBody is used to define fields on response body types.
type NavTResponseBody struct {
	Self  *string `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	First *string `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Next  *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
}

// ParameterT is used to define fields on request body types.
type ParameterT struct {
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "order" service.
func NewCreateRequestBody(p *order.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		ServiceID: p.Orders.ServiceID,
		AccountID: p.Orders.AccountID,
		PolicyID:  p.Orders.PolicyID,
		Name:      p.Orders.Name,
	}
	if p.Orders.Tags != nil {
		body.Tags = make([]string, len(p.Orders.Tags))
		for i, val := range p.Orders.Tags {
			body.Tags[i] = val
		}
	}
	if p.Orders.Parameters != nil {
		body.Parameters = make([]*ParameterT, len(p.Orders.Parameters))
		for i, val := range p.Orders.Parameters {
			body.Parameters[i] = marshalOrderParameterTToParameterT(val)
		}
	}
	return body
}

// NewLogsRequestBody builds the HTTP request body from the payload of the
// "logs" endpoint of the "order" service.
func NewLogsRequestBody(p *order.LogsPayload) *LogsRequestBody {
	body := &LogsRequestBody{
		From:          p.DownloadLogRequest.From,
		To:            p.DownloadLogRequest.To,
		NamespaceName: p.DownloadLogRequest.NamespaceName,
		ContainerName: p.DownloadLogRequest.ContainerName,
		OrderID:       p.DownloadLogRequest.OrderID,
		PolicyID:      p.DownloadLogRequest.PolicyID,
	}
	return body
}

// NewReadOrderStatusRTOK builds a "order" service "read" endpoint result from
// a HTTP "OK" response.
func NewReadOrderStatusRTOK(body *ReadResponseBody) *orderviews.OrderStatusRTView {
	v := &orderviews.OrderStatusRTView{
		ID:         body.ID,
		Status:     body.Status,
		OrderedAt:  body.OrderedAt,
		StartedAt:  body.StartedAt,
		FinishedAt: body.FinishedAt,
		Name:       body.Name,
	}
	if body.Products != nil {
		v.Products = make([]*orderviews.ProductTView, len(body.Products))
		for i, val := range body.Products {
			v.Products[i] = unmarshalProductTResponseBodyToOrderviewsProductTView(val)
		}
	}
	if body.Service != nil {
		v.Service = unmarshalRefTResponseBodyToOrderviewsRefTView(body.Service)
	}
	if body.Account != nil {
		v.Account = unmarshalRefTResponseBodyToOrderviewsRefTView(body.Account)
	}
	if body.Links != nil {
		v.Links = unmarshalSelfTResponseBodyToOrderviewsSelfTView(body.Links)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*orderviews.ParameterTView, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterTResponseBodyToOrderviewsParameterTView(val)
	}

	return v
}

// NewReadBadRequest builds a order service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *order.BadRequestT {
	v := &order.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidCredential builds a order service read endpoint
// invalid-credential error.
func NewReadInvalidCredential() *order.InvalidCredentialsT {
	v := &order.InvalidCredentialsT{}

	return v
}

// NewReadInvalidScopes builds a order service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *order.InvalidScopesT {
	v := &order.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a order service read endpoint not-implemented
// error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *order.NotImplementedT {
	v := &order.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a order service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *order.ResourceNotFoundT {
	v := &order.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAuthorized builds a order service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *order.UnauthorizedT {
	v := &order.UnauthorizedT{}

	return v
}

// NewListOrderListRTOK builds a "order" service "list" endpoint result from a
// HTTP "OK" response.
func NewListOrderListRTOK(body *ListResponseBody) *orderviews.OrderListRTView {
	v := &orderviews.OrderListRTView{
		AtTime: body.AtTime,
	}
	v.Orders = make([]*orderviews.OrderListItemView, len(body.Orders))
	for i, val := range body.Orders {
		v.Orders[i] = unmarshalOrderListItemResponseBodyToOrderviewsOrderListItemView(val)
	}
	v.Links = unmarshalNavTResponseBodyToOrderviewsNavTView(body.Links)

	return v
}

// NewListBadRequest builds a order service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *order.BadRequestT {
	v := &order.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidCredential builds a order service list endpoint
// invalid-credential error.
func NewListInvalidCredential() *order.InvalidCredentialsT {
	v := &order.InvalidCredentialsT{}

	return v
}

// NewListInvalidParameter builds a order service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *order.InvalidParameterValue {
	v := &order.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a order service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *order.InvalidScopesT {
	v := &order.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a order service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *order.NotImplementedT {
	v := &order.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAuthorized builds a order service list endpoint not-authorized
// error.
func NewListNotAuthorized() *order.UnauthorizedT {
	v := &order.UnauthorizedT{}

	return v
}

// NewCreateOrderStatusRTOK builds a "order" service "create" endpoint result
// from a HTTP "OK" response.
func NewCreateOrderStatusRTOK(body *CreateResponseBody) *orderviews.OrderStatusRTView {
	v := &orderviews.OrderStatusRTView{
		ID:         body.ID,
		Status:     body.Status,
		OrderedAt:  body.OrderedAt,
		StartedAt:  body.StartedAt,
		FinishedAt: body.FinishedAt,
		Name:       body.Name,
	}
	if body.Products != nil {
		v.Products = make([]*orderviews.ProductTView, len(body.Products))
		for i, val := range body.Products {
			v.Products[i] = unmarshalProductTResponseBodyToOrderviewsProductTView(val)
		}
	}
	if body.Service != nil {
		v.Service = unmarshalRefTResponseBodyToOrderviewsRefTView(body.Service)
	}
	if body.Account != nil {
		v.Account = unmarshalRefTResponseBodyToOrderviewsRefTView(body.Account)
	}
	if body.Links != nil {
		v.Links = unmarshalSelfTResponseBodyToOrderviewsSelfTView(body.Links)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*orderviews.ParameterTView, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterTResponseBodyToOrderviewsParameterTView(val)
	}

	return v
}

// NewCreateBadRequest builds a order service create endpoint bad-request error.
func NewCreateBadRequest(body *CreateBadRequestResponseBody) *order.BadRequestT {
	v := &order.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewCreateInvalidCredential builds a order service create endpoint
// invalid-credential error.
func NewCreateInvalidCredential() *order.InvalidCredentialsT {
	v := &order.InvalidCredentialsT{}

	return v
}

// NewCreateInvalidParameter builds a order service create endpoint
// invalid-parameter error.
func NewCreateInvalidParameter(body *CreateInvalidParameterResponseBody) *order.InvalidParameterValue {
	v := &order.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewCreateInvalidScopes builds a order service create endpoint invalid-scopes
// error.
func NewCreateInvalidScopes(body *CreateInvalidScopesResponseBody) *order.InvalidScopesT {
	v := &order.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateNotImplemented builds a order service create endpoint
// not-implemented error.
func NewCreateNotImplemented(body *CreateNotImplementedResponseBody) *order.NotImplementedT {
	v := &order.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewCreateNotFound builds a order service create endpoint not-found error.
func NewCreateNotFound(body *CreateNotFoundResponseBody) *order.ResourceNotFoundT {
	v := &order.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateNotAvailable builds a order service create endpoint not-available
// error.
func NewCreateNotAvailable() *order.ServiceNotAvailableT {
	v := &order.ServiceNotAvailableT{}

	return v
}

// NewCreateNotAuthorized builds a order service create endpoint not-authorized
// error.
func NewCreateNotAuthorized() *order.UnauthorizedT {
	v := &order.UnauthorizedT{}

	return v
}

// NewLogsBadRequest builds a order service logs endpoint bad-request error.
func NewLogsBadRequest(body *LogsBadRequestResponseBody) *order.BadRequestT {
	v := &order.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewLogsInvalidCredential builds a order service logs endpoint
// invalid-credential error.
func NewLogsInvalidCredential() *order.InvalidCredentialsT {
	v := &order.InvalidCredentialsT{}

	return v
}

// NewLogsInvalidParameter builds a order service logs endpoint
// invalid-parameter error.
func NewLogsInvalidParameter(body *LogsInvalidParameterResponseBody) *order.InvalidParameterValue {
	v := &order.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewLogsInvalidScopes builds a order service logs endpoint invalid-scopes
// error.
func NewLogsInvalidScopes(body *LogsInvalidScopesResponseBody) *order.InvalidScopesT {
	v := &order.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewLogsNotImplemented builds a order service logs endpoint not-implemented
// error.
func NewLogsNotImplemented(body *LogsNotImplementedResponseBody) *order.NotImplementedT {
	v := &order.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewLogsNotFound builds a order service logs endpoint not-found error.
func NewLogsNotFound(body *LogsNotFoundResponseBody) *order.ResourceNotFoundT {
	v := &order.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewLogsNotAuthorized builds a order service logs endpoint not-authorized
// error.
func NewLogsNotAuthorized() *order.UnauthorizedT {
	v := &order.UnauthorizedT{}

	return v
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

// ValidateCreateBadRequestResponseBody runs the validations defined on
// create_bad-request_response_body
func ValidateCreateBadRequestResponseBody(body *CreateBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateInvalidParameterResponseBody runs the validations defined on
// create_invalid-parameter_response_body
func ValidateCreateInvalidParameterResponseBody(body *CreateInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateInvalidScopesResponseBody runs the validations defined on
// create_invalid-scopes_response_body
func ValidateCreateInvalidScopesResponseBody(body *CreateInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateCreateNotImplementedResponseBody runs the validations defined on
// create_not-implemented_response_body
func ValidateCreateNotImplementedResponseBody(body *CreateNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateNotFoundResponseBody runs the validations defined on
// create_not-found_response_body
func ValidateCreateNotFoundResponseBody(body *CreateNotFoundResponseBody) (err error) {
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

// ValidateLogsBadRequestResponseBody runs the validations defined on
// logs_bad-request_response_body
func ValidateLogsBadRequestResponseBody(body *LogsBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLogsInvalidParameterResponseBody runs the validations defined on
// logs_invalid-parameter_response_body
func ValidateLogsInvalidParameterResponseBody(body *LogsInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLogsInvalidScopesResponseBody runs the validations defined on
// logs_invalid-scopes_response_body
func ValidateLogsInvalidScopesResponseBody(body *LogsInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateLogsNotImplementedResponseBody runs the validations defined on
// logs_not-implemented_response_body
func ValidateLogsNotImplementedResponseBody(body *LogsNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLogsNotFoundResponseBody runs the validations defined on
// logs_not-found_response_body
func ValidateLogsNotFoundResponseBody(body *LogsNotFoundResponseBody) (err error) {
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

// ValidateProductTResponseBody runs the validations defined on
// ProductTResponseBody
func ValidateProductTResponseBody(body *ProductTResponseBody) (err error) {
	if body.Links != nil {
		if err2 := ValidateSelfWithDataTResponseBody(body.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSelfWithDataTResponseBody runs the validations defined on
// SelfWithDataTResponseBody
func ValidateSelfWithDataTResponseBody(body *SelfWithDataTResponseBody) (err error) {
	if body.DescribedBy != nil {
		if err2 := ValidateDescribedByTResponseBody(body.DescribedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDescribedByTResponseBody runs the validations defined on
// DescribedByTResponseBody
func ValidateDescribedByTResponseBody(body *DescribedByTResponseBody) (err error) {
	if body.Href != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.href", *body.Href, goa.FormatURI))
	}
	return
}

// ValidateRefTResponseBody runs the validations defined on RefTResponseBody
func ValidateRefTResponseBody(body *RefTResponseBody) (err error) {
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Links != nil {
		if err2 := ValidateSelfTResponseBody(body.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSelfTResponseBody runs the validations defined on SelfTResponseBody
func ValidateSelfTResponseBody(body *SelfTResponseBody) (err error) {
	if body.DescribedBy != nil {
		if err2 := ValidateDescribedByTResponseBody(body.DescribedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateOrderListItemResponseBody runs the validations defined on
// OrderListItemResponseBody
func ValidateOrderListItemResponseBody(body *OrderListItemResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "unknown" || *body.Status == "pending" || *body.Status == "scheduled" || *body.Status == "executing" || *body.Status == "succeeded" || *body.Status == "failed" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if body.Links != nil {
		if err2 := ValidateSelfTResponseBody(body.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateNavTResponseBody runs the validations defined on NavTResponseBody
func ValidateNavTResponseBody(body *NavTResponseBody) (err error) {
	if body.Self != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.self", *body.Self, goa.FormatURI))
	}
	if body.First != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.first", *body.First, goa.FormatURI))
	}
	if body.Next != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.next", *body.Next, goa.FormatURI))
	}
	return
}
