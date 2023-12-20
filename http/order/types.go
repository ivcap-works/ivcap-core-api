// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	order "github.com/ivcap-works/ivcap-core-api/gen/order"
	orderviews "github.com/ivcap-works/ivcap-core-api/gen/order/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "order" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Reference to service requested
	Service string `json:"service"`
	// Reference to policy used
	Policy *string `json:"policy"`
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
	NamespaceName *string `json:"namespace-name,omitempty"`
	// Reference to container name
	ContainerName *string `json:"container-name,omitempty"`
	// Reference to order
	Order string `json:"order"`
}

// TopRequestBody is the type of the "order" service "top" endpoint HTTP
// request body.
type TopRequestBody struct {
	// Reference to order
	Order string `json:"order"`
	// Reference to namespace name
	NamespaceName *string `json:"namespace-name,omitempty"`
}

// ReadResponseBody is the type of the "order" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Order status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime order was placed
	OrderedAt *string `form:"ordered-at,omitempty" json:"ordered-at,omitempty" xml:"ordered-at,omitempty"`
	// DateTime order processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime order processing finished
	FinishedAt *string                          `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	Products   *PartialProductListTResponseBody `form:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
	// Reference to service requested
	Service *string `json:"service"`
	// Reference to billable account
	Account *string              `json:"account"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
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
	Items []*OrderListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// CreateResponseBody is the type of the "order" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Order status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime order was placed
	OrderedAt *string `form:"ordered-at,omitempty" json:"ordered-at,omitempty" xml:"ordered-at,omitempty"`
	// DateTime order processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime order processing finished
	FinishedAt *string                          `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	Products   *PartialProductListTResponseBody `form:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
	// Reference to service requested
	Service *string `json:"service"`
	// Reference to billable account
	Account *string              `json:"account"`
	Links   []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Optional customer provided tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Service parameters
	Parameters []*ParameterTResponseBody `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// TopResponseBody is the type of the "order" service "top" endpoint HTTP
// response body.
type TopResponseBody []*OrderTopResultItemResponse

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

// TopBadRequestResponseBody is the type of the "order" service "top" endpoint
// HTTP response body for the "bad-request" error.
type TopBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// TopInvalidParameterResponseBody is the type of the "order" service "top"
// endpoint HTTP response body for the "invalid-parameter" error.
type TopInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// TopInvalidScopesResponseBody is the type of the "order" service "top"
// endpoint HTTP response body for the "invalid-scopes" error.
type TopInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// TopNotImplementedResponseBody is the type of the "order" service "top"
// endpoint HTTP response body for the "not-implemented" error.
type TopNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// TopNotFoundResponseBody is the type of the "order" service "top" endpoint
// HTTP response body for the "not-found" error.
type TopNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PartialProductListTResponseBody is used to define fields on response body
// types.
type PartialProductListTResponseBody struct {
	// (Partial) list of products delivered by this order
	Items []*ProductListItemTResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Links to more products, if there are any
	Links []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// ProductListItemTResponseBody is used to define fields on response body types.
type ProductListItemTResponseBody struct {
	ID       *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Status   *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	MimeType *string `json:"mime-type,omitempty"`
	Size     *int64  `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	Href     *string `json:"href,omitempty"`
	DataHref *string `json:"dataRef,omitempty"`
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

// ParameterTResponseBody is used to define fields on response body types.
type ParameterTResponseBody struct {
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// OrderListItemResponseBody is used to define fields on response body types.
type OrderListItemResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Optional customer provided name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Order status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime order was placed
	OrderedAt *string `form:"ordered-at,omitempty" json:"ordered-at,omitempty" xml:"ordered-at,omitempty"`
	// DateTime order processing started
	StartedAt *string `form:"started-at,omitempty" json:"started-at,omitempty" xml:"started-at,omitempty"`
	// DateTime order processing finished
	FinishedAt *string `form:"finished-at,omitempty" json:"finished-at,omitempty" xml:"finished-at,omitempty"`
	// Reference to service requested
	Service *string `json:"service"`
	// Reference to billable account
	Account *string `json:"account"`
	Href    *string `json:"href,omitempty"`
}

// ParameterT is used to define fields on request body types.
type ParameterT struct {
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// OrderTopResultItemResponse is used to define fields on response body types.
type OrderTopResultItemResponse struct {
	// container
	Container *string `form:"container,omitempty" json:"container,omitempty" xml:"container,omitempty"`
	// cpu
	CPU *string `form:"cpu,omitempty" json:"cpu,omitempty" xml:"cpu,omitempty"`
	// memory
	Memory *string `form:"memory,omitempty" json:"memory,omitempty" xml:"memory,omitempty"`
	// storage
	Storage *string `form:"storage,omitempty" json:"storage,omitempty" xml:"storage,omitempty"`
	// ephemeral-storage
	EphemeralStorage *string `form:"ephemeral-storage,omitempty" json:"ephemeral-storage,omitempty" xml:"ephemeral-storage,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "order" service.
func NewCreateRequestBody(p *order.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Service: p.Orders.Service,
		Policy:  p.Orders.Policy,
		Name:    p.Orders.Name,
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
	} else {
		body.Parameters = []*ParameterT{}
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
		Order:         p.DownloadLogRequest.Order,
	}
	return body
}

// NewTopRequestBody builds the HTTP request body from the payload of the "top"
// endpoint of the "order" service.
func NewTopRequestBody(p *order.TopPayload) *TopRequestBody {
	body := &TopRequestBody{
		Order:         p.OrderTopRequest.Order,
		NamespaceName: p.OrderTopRequest.NamespaceName,
	}
	return body
}

// NewReadOrderStatusRTOK builds a "order" service "read" endpoint result from
// a HTTP "OK" response.
func NewReadOrderStatusRTOK(body *ReadResponseBody) *order.OrderStatusRT {
	v := &order.OrderStatusRT{
		ID:         *body.ID,
		Status:     *body.Status,
		OrderedAt:  body.OrderedAt,
		StartedAt:  body.StartedAt,
		FinishedAt: body.FinishedAt,
		Service:    *body.Service,
		Account:    *body.Account,
		Name:       body.Name,
	}
	v.Products = unmarshalPartialProductListTResponseBodyToOrderPartialProductListT(body.Products)
	v.Links = make([]*order.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToOrderLinkT(val)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*order.ParameterT, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterTResponseBodyToOrderParameterT(val)
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
	v.Items = make([]*orderviews.OrderListItemView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalOrderListItemResponseBodyToOrderviewsOrderListItemView(val)
	}
	v.Links = make([]*orderviews.LinkTView, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToOrderviewsLinkTView(val)
	}

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
func NewCreateOrderStatusRTOK(body *CreateResponseBody) *order.OrderStatusRT {
	v := &order.OrderStatusRT{
		ID:         *body.ID,
		Status:     *body.Status,
		OrderedAt:  body.OrderedAt,
		StartedAt:  body.StartedAt,
		FinishedAt: body.FinishedAt,
		Service:    *body.Service,
		Account:    *body.Account,
		Name:       body.Name,
	}
	v.Products = unmarshalPartialProductListTResponseBodyToOrderPartialProductListT(body.Products)
	v.Links = make([]*order.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToOrderLinkT(val)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Parameters = make([]*order.ParameterT, len(body.Parameters))
	for i, val := range body.Parameters {
		v.Parameters[i] = unmarshalParameterTResponseBodyToOrderParameterT(val)
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

// NewTopOrderTopResultItemCollectionOK builds a "order" service "top" endpoint
// result from a HTTP "OK" response.
func NewTopOrderTopResultItemCollectionOK(body TopResponseBody) orderviews.OrderTopResultItemCollectionView {
	v := make([]*orderviews.OrderTopResultItemView, len(body))
	for i, val := range body {
		v[i] = unmarshalOrderTopResultItemResponseToOrderviewsOrderTopResultItemView(val)
	}

	return v
}

// NewTopBadRequest builds a order service top endpoint bad-request error.
func NewTopBadRequest(body *TopBadRequestResponseBody) *order.BadRequestT {
	v := &order.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewTopInvalidCredential builds a order service top endpoint
// invalid-credential error.
func NewTopInvalidCredential() *order.InvalidCredentialsT {
	v := &order.InvalidCredentialsT{}

	return v
}

// NewTopInvalidParameter builds a order service top endpoint invalid-parameter
// error.
func NewTopInvalidParameter(body *TopInvalidParameterResponseBody) *order.InvalidParameterValue {
	v := &order.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewTopInvalidScopes builds a order service top endpoint invalid-scopes error.
func NewTopInvalidScopes(body *TopInvalidScopesResponseBody) *order.InvalidScopesT {
	v := &order.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewTopNotImplemented builds a order service top endpoint not-implemented
// error.
func NewTopNotImplemented(body *TopNotImplementedResponseBody) *order.NotImplementedT {
	v := &order.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewTopNotFound builds a order service top endpoint not-found error.
func NewTopNotFound(body *TopNotFoundResponseBody) *order.ResourceNotFoundT {
	v := &order.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewTopNotAuthorized builds a order service top endpoint not-authorized error.
func NewTopNotAuthorized() *order.UnauthorizedT {
	v := &order.UnauthorizedT{}

	return v
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
	if body.Products == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("products", "body"))
	}
	if body.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "body"))
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
		if !(*body.Status == "unknown" || *body.Status == "pending" || *body.Status == "scheduled" || *body.Status == "executing" || *body.Status == "succeeded" || *body.Status == "failed" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if body.OrderedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.ordered-at", *body.OrderedAt, goa.FormatDateTime))
	}
	if body.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.started-at", *body.StartedAt, goa.FormatDateTime))
	}
	if body.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.finished-at", *body.FinishedAt, goa.FormatDateTime))
	}
	if body.Products != nil {
		if err2 := ValidatePartialProductListTResponseBody(body.Products); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Service != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.service", *body.Service, goa.FormatURI))
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

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Products == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("products", "body"))
	}
	if body.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "body"))
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
		if !(*body.Status == "unknown" || *body.Status == "pending" || *body.Status == "scheduled" || *body.Status == "executing" || *body.Status == "succeeded" || *body.Status == "failed" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if body.OrderedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.ordered-at", *body.OrderedAt, goa.FormatDateTime))
	}
	if body.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.started-at", *body.StartedAt, goa.FormatDateTime))
	}
	if body.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.finished-at", *body.FinishedAt, goa.FormatDateTime))
	}
	if body.Products != nil {
		if err2 := ValidatePartialProductListTResponseBody(body.Products); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Service != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.service", *body.Service, goa.FormatURI))
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

// ValidateTopBadRequestResponseBody runs the validations defined on
// top_bad-request_response_body
func ValidateTopBadRequestResponseBody(body *TopBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateTopInvalidParameterResponseBody runs the validations defined on
// top_invalid-parameter_response_body
func ValidateTopInvalidParameterResponseBody(body *TopInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateTopInvalidScopesResponseBody runs the validations defined on
// top_invalid-scopes_response_body
func ValidateTopInvalidScopesResponseBody(body *TopInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateTopNotImplementedResponseBody runs the validations defined on
// top_not-implemented_response_body
func ValidateTopNotImplementedResponseBody(body *TopNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateTopNotFoundResponseBody runs the validations defined on
// top_not-found_response_body
func ValidateTopNotFoundResponseBody(body *TopNotFoundResponseBody) (err error) {
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

// ValidatePartialProductListTResponseBody runs the validations defined on
// PartialProductListTResponseBody
func ValidatePartialProductListTResponseBody(body *PartialProductListTResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	for _, e := range body.Items {
		if e != nil {
			if err2 := ValidateProductListItemTResponseBody(e); err2 != nil {
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

// ValidateProductListItemTResponseBody runs the validations defined on
// ProductListItemTResponseBody
func ValidateProductListItemTResponseBody(body *ProductListItemTResponseBody) (err error) {
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

// ValidateOrderListItemResponseBody runs the validations defined on
// OrderListItemResponseBody
func ValidateOrderListItemResponseBody(body *OrderListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "body"))
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
	if body.Status != nil {
		if !(*body.Status == "unknown" || *body.Status == "pending" || *body.Status == "scheduled" || *body.Status == "executing" || *body.Status == "succeeded" || *body.Status == "failed" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if body.OrderedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.ordered-at", *body.OrderedAt, goa.FormatDateTime))
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
	if body.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account", *body.Account, goa.FormatURI))
	}
	return
}

// ValidateOrderTopResultItemResponse runs the validations defined on
// OrderTopResultItemResponse
func ValidateOrderTopResultItemResponse(body *OrderTopResultItemResponse) (err error) {
	if body.Container == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("container", "body"))
	}
	if body.CPU == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cpu", "body"))
	}
	if body.Memory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("memory", "body"))
	}
	if body.Storage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("storage", "body"))
	}
	if body.EphemeralStorage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ephemeral-storage", "body"))
	}
	return
}
