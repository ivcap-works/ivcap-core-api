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

package client

import (
	queue "github.com/ivcap-works/ivcap-core-api/gen/queue"
	queueviews "github.com/ivcap-works/ivcap-core-api/gen/queue/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "queue" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Optional Name for the queue. Cannot contain whitespace, ., *, >, path
	// separators (forward or backwards slash), and non-printable characters.
	Name string `form:"name" json:"name" xml:"name"`
	// More detailed description of the queue.
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Reference to policy used
	Policy *string `form:"policy,omitempty" json:"policy,omitempty" xml:"policy,omitempty"`
}

// CreateResponseBody is the type of the "queue" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// queue
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the created queue.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the created queue.
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Reference to billable account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
}

// ReadResponseBody is the type of the "queue" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the queue.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the queue.
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Number of messages sent to the queue
	TotalMessages *uint64 `form:"total-messages,omitempty" json:"total-messages,omitempty" xml:"total-messages,omitempty"`
	// Number of bytes in the queue
	Bytes *uint64 `form:"bytes,omitempty" json:"bytes,omitempty" xml:"bytes,omitempty"`
	// Timestamp of the first message in the queue
	FirstTime *string `form:"first-time,omitempty" json:"first-time,omitempty" xml:"first-time,omitempty"`
	// Timestamp of the last message in the queue
	LastTime *string `form:"last-time,omitempty" json:"last-time,omitempty" xml:"last-time,omitempty"`
	// Number of consumers
	ConsumerCount *int `form:"consumer-count,omitempty" json:"consumer-count,omitempty" xml:"consumer-count,omitempty"`
	// Timestamp when the queue was created
	CreatedAt *string `form:"created-at,omitempty" json:"created-at,omitempty" xml:"created-at,omitempty"`
}

// ListResponseBody is the type of the "queue" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Queues
	Items []*QueueListItemResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Time at which this list was valid
	AtTime *string              `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	Links  []*LinkTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// EnqueueResponseBody is the type of the "queue" service "enqueue" endpoint
// HTTP response body.
type EnqueueResponseBody struct {
	// queue
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// DequeueResponseBody is the type of the "queue" service "dequeue" endpoint
// HTTP response body.
type DequeueResponseBody struct {
	// Messages in the queue
	Messages []*PublishedmessageResponseBody `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "queue" service "create"
// endpoint HTTP response body for the "bad-request" error.
type CreateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateInvalidParameterResponseBody is the type of the "queue" service
// "create" endpoint HTTP response body for the "invalid-parameter" error.
type CreateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// CreateInvalidScopesResponseBody is the type of the "queue" service "create"
// endpoint HTTP response body for the "invalid-scopes" error.
type CreateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateNotImplementedResponseBody is the type of the "queue" service "create"
// endpoint HTTP response body for the "not-implemented" error.
type CreateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateAlreadyCreatedResponseBody is the type of the "queue" service "create"
// endpoint HTTP response body for the "already-created" error.
type CreateAlreadyCreatedResponseBody struct {
	// ID of already existing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateNotFoundResponseBody is the type of the "queue" service "create"
// endpoint HTTP response body for the "not-found" error.
type CreateNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "queue" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "queue" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "queue" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "queue" service "read" endpoint
// HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteBadRequestResponseBody is the type of the "queue" service "delete"
// endpoint HTTP response body for the "bad-request" error.
type DeleteBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteInvalidScopesResponseBody is the type of the "queue" service "delete"
// endpoint HTTP response body for the "invalid-scopes" error.
type DeleteInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteNotImplementedResponseBody is the type of the "queue" service "delete"
// endpoint HTTP response body for the "not-implemented" error.
type DeleteNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListBadRequestResponseBody is the type of the "queue" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "queue" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "queue" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "queue" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// EnqueueBadRequestResponseBody is the type of the "queue" service "enqueue"
// endpoint HTTP response body for the "bad-request" error.
type EnqueueBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// EnqueueInvalidParameterResponseBody is the type of the "queue" service
// "enqueue" endpoint HTTP response body for the "invalid-parameter" error.
type EnqueueInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// EnqueueInvalidScopesResponseBody is the type of the "queue" service
// "enqueue" endpoint HTTP response body for the "invalid-scopes" error.
type EnqueueInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// EnqueueNotImplementedResponseBody is the type of the "queue" service
// "enqueue" endpoint HTTP response body for the "not-implemented" error.
type EnqueueNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DequeueBadRequestResponseBody is the type of the "queue" service "dequeue"
// endpoint HTTP response body for the "bad-request" error.
type DequeueBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DequeueInvalidParameterResponseBody is the type of the "queue" service
// "dequeue" endpoint HTTP response body for the "invalid-parameter" error.
type DequeueInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// DequeueInvalidScopesResponseBody is the type of the "queue" service
// "dequeue" endpoint HTTP response body for the "invalid-scopes" error.
type DequeueInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DequeueNotImplementedResponseBody is the type of the "queue" service
// "dequeue" endpoint HTTP response body for the "not-implemented" error.
type DequeueNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// QueueListItemResponseBody is used to define fields on response body types.
type QueueListItemResponseBody struct {
	// queue
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the created queue.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the created queue.
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Reference to billable account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	Href    *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
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

// PublishedmessageResponseBody is used to define fields on response body types.
type PublishedmessageResponseBody struct {
	// Message identifier
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message content in JSON format.
	Content any `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Schema used for message
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Encoding type of message content (defaults to 'application/json')
	ContentType *string `form:"content-type,omitempty" json:"content-type,omitempty" xml:"content-type,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "queue" service.
func NewCreateRequestBody(p *queue.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:        p.Queues.Name,
		Description: p.Queues.Description,
		Policy:      p.Queues.Policy,
	}
	return body
}

// NewCreatequeueresponseViewCreated builds a "queue" service "create" endpoint
// result from a HTTP "Created" response.
func NewCreatequeueresponseViewCreated(body *CreateResponseBody) *queueviews.CreatequeueresponseView {
	v := &queueviews.CreatequeueresponseView{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Account:     body.Account,
	}

	return v
}

// NewCreateBadRequest builds a queue service create endpoint bad-request error.
func NewCreateBadRequest(body *CreateBadRequestResponseBody) *queue.BadRequestT {
	v := &queue.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewCreateInvalidParameter builds a queue service create endpoint
// invalid-parameter error.
func NewCreateInvalidParameter(body *CreateInvalidParameterResponseBody) *queue.InvalidParameterT {
	v := &queue.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewCreateInvalidScopes builds a queue service create endpoint invalid-scopes
// error.
func NewCreateInvalidScopes(body *CreateInvalidScopesResponseBody) *queue.InvalidScopesT {
	v := &queue.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateNotImplemented builds a queue service create endpoint
// not-implemented error.
func NewCreateNotImplemented(body *CreateNotImplementedResponseBody) *queue.NotImplementedT {
	v := &queue.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewCreateAlreadyCreated builds a queue service create endpoint
// already-created error.
func NewCreateAlreadyCreated(body *CreateAlreadyCreatedResponseBody) *queue.ResourceAlreadyCreatedT {
	v := &queue.ResourceAlreadyCreatedT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateNotFound builds a queue service create endpoint not-found error.
func NewCreateNotFound(body *CreateNotFoundResponseBody) *queue.ResourceNotFoundT {
	v := &queue.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateNotAvailable builds a queue service create endpoint not-available
// error.
func NewCreateNotAvailable() *queue.ServiceNotAvailableT {
	v := &queue.ServiceNotAvailableT{}

	return v
}

// NewCreateNotAuthorized builds a queue service create endpoint not-authorized
// error.
func NewCreateNotAuthorized() *queue.UnauthorizedT {
	v := &queue.UnauthorizedT{}

	return v
}

// NewReadqueueresponseViewCreated builds a "queue" service "read" endpoint
// result from a HTTP "Created" response.
func NewReadqueueresponseViewCreated(body *ReadResponseBody) *queueviews.ReadqueueresponseView {
	v := &queueviews.ReadqueueresponseView{
		ID:            body.ID,
		Name:          body.Name,
		Description:   body.Description,
		TotalMessages: body.TotalMessages,
		Bytes:         body.Bytes,
		FirstTime:     body.FirstTime,
		LastTime:      body.LastTime,
		ConsumerCount: body.ConsumerCount,
		CreatedAt:     body.CreatedAt,
	}

	return v
}

// NewReadBadRequest builds a queue service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *queue.BadRequestT {
	v := &queue.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidScopes builds a queue service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *queue.InvalidScopesT {
	v := &queue.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a queue service read endpoint not-implemented
// error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *queue.NotImplementedT {
	v := &queue.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a queue service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *queue.ResourceNotFoundT {
	v := &queue.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAvailable builds a queue service read endpoint not-available error.
func NewReadNotAvailable() *queue.ServiceNotAvailableT {
	v := &queue.ServiceNotAvailableT{}

	return v
}

// NewReadNotAuthorized builds a queue service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *queue.UnauthorizedT {
	v := &queue.UnauthorizedT{}

	return v
}

// NewDeleteBadRequest builds a queue service delete endpoint bad-request error.
func NewDeleteBadRequest(body *DeleteBadRequestResponseBody) *queue.BadRequestT {
	v := &queue.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewDeleteInvalidScopes builds a queue service delete endpoint invalid-scopes
// error.
func NewDeleteInvalidScopes(body *DeleteInvalidScopesResponseBody) *queue.InvalidScopesT {
	v := &queue.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotImplemented builds a queue service delete endpoint
// not-implemented error.
func NewDeleteNotImplemented(body *DeleteNotImplementedResponseBody) *queue.NotImplementedT {
	v := &queue.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotAvailable builds a queue service delete endpoint not-available
// error.
func NewDeleteNotAvailable() *queue.ServiceNotAvailableT {
	v := &queue.ServiceNotAvailableT{}

	return v
}

// NewDeleteNotAuthorized builds a queue service delete endpoint not-authorized
// error.
func NewDeleteNotAuthorized() *queue.UnauthorizedT {
	v := &queue.UnauthorizedT{}

	return v
}

// NewListQueueListResultOK builds a "queue" service "list" endpoint result
// from a HTTP "OK" response.
func NewListQueueListResultOK(body *ListResponseBody) *queue.QueueListResult {
	v := &queue.QueueListResult{
		AtTime: *body.AtTime,
	}
	v.Items = make([]*queue.QueueListItem, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalQueueListItemResponseBodyToQueueQueueListItem(val)
	}
	v.Links = make([]*queue.LinkT, len(body.Links))
	for i, val := range body.Links {
		v.Links[i] = unmarshalLinkTResponseBodyToQueueLinkT(val)
	}

	return v
}

// NewListBadRequest builds a queue service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *queue.BadRequestT {
	v := &queue.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidParameter builds a queue service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *queue.InvalidParameterT {
	v := &queue.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a queue service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *queue.InvalidScopesT {
	v := &queue.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a queue service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *queue.NotImplementedT {
	v := &queue.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAvailable builds a queue service list endpoint not-available error.
func NewListNotAvailable() *queue.ServiceNotAvailableT {
	v := &queue.ServiceNotAvailableT{}

	return v
}

// NewListNotAuthorized builds a queue service list endpoint not-authorized
// error.
func NewListNotAuthorized() *queue.UnauthorizedT {
	v := &queue.UnauthorizedT{}

	return v
}

// NewEnqueueMessagestatusOK builds a "queue" service "enqueue" endpoint result
// from a HTTP "OK" response.
func NewEnqueueMessagestatusOK(body *EnqueueResponseBody) *queueviews.MessagestatusView {
	v := &queueviews.MessagestatusView{
		ID: body.ID,
	}

	return v
}

// NewEnqueueBadRequest builds a queue service enqueue endpoint bad-request
// error.
func NewEnqueueBadRequest(body *EnqueueBadRequestResponseBody) *queue.BadRequestT {
	v := &queue.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewEnqueueInvalidParameter builds a queue service enqueue endpoint
// invalid-parameter error.
func NewEnqueueInvalidParameter(body *EnqueueInvalidParameterResponseBody) *queue.InvalidParameterT {
	v := &queue.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewEnqueueInvalidScopes builds a queue service enqueue endpoint
// invalid-scopes error.
func NewEnqueueInvalidScopes(body *EnqueueInvalidScopesResponseBody) *queue.InvalidScopesT {
	v := &queue.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewEnqueueNotImplemented builds a queue service enqueue endpoint
// not-implemented error.
func NewEnqueueNotImplemented(body *EnqueueNotImplementedResponseBody) *queue.NotImplementedT {
	v := &queue.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewEnqueueNotAvailable builds a queue service enqueue endpoint not-available
// error.
func NewEnqueueNotAvailable() *queue.ServiceNotAvailableT {
	v := &queue.ServiceNotAvailableT{}

	return v
}

// NewEnqueueNotAuthorized builds a queue service enqueue endpoint
// not-authorized error.
func NewEnqueueNotAuthorized() *queue.UnauthorizedT {
	v := &queue.UnauthorizedT{}

	return v
}

// NewDequeueMessageListOK builds a "queue" service "dequeue" endpoint result
// from a HTTP "OK" response.
func NewDequeueMessageListOK(body *DequeueResponseBody) *queue.MessageList {
	v := &queue.MessageList{
		AtTime: body.AtTime,
	}
	v.Messages = make([]*queue.Publishedmessage, len(body.Messages))
	for i, val := range body.Messages {
		v.Messages[i] = unmarshalPublishedmessageResponseBodyToQueuePublishedmessage(val)
	}

	return v
}

// NewDequeueBadRequest builds a queue service dequeue endpoint bad-request
// error.
func NewDequeueBadRequest(body *DequeueBadRequestResponseBody) *queue.BadRequestT {
	v := &queue.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewDequeueInvalidParameter builds a queue service dequeue endpoint
// invalid-parameter error.
func NewDequeueInvalidParameter(body *DequeueInvalidParameterResponseBody) *queue.InvalidParameterT {
	v := &queue.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewDequeueInvalidScopes builds a queue service dequeue endpoint
// invalid-scopes error.
func NewDequeueInvalidScopes(body *DequeueInvalidScopesResponseBody) *queue.InvalidScopesT {
	v := &queue.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewDequeueNotImplemented builds a queue service dequeue endpoint
// not-implemented error.
func NewDequeueNotImplemented(body *DequeueNotImplementedResponseBody) *queue.NotImplementedT {
	v := &queue.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewDequeueNotAvailable builds a queue service dequeue endpoint not-available
// error.
func NewDequeueNotAvailable() *queue.ServiceNotAvailableT {
	v := &queue.ServiceNotAvailableT{}

	return v
}

// NewDequeueNotAuthorized builds a queue service dequeue endpoint
// not-authorized error.
func NewDequeueNotAuthorized() *queue.UnauthorizedT {
	v := &queue.UnauthorizedT{}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.AtTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("at-time", "body"))
	}
	if body.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "body"))
	}
	for _, e := range body.Items {
		if e != nil {
			if err2 := ValidateQueueListItemResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.at-time", *body.AtTime, goa.FormatDateTime))
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

// ValidateDequeueResponseBody runs the validations defined on
// DequeueResponseBody
func ValidateDequeueResponseBody(body *DequeueResponseBody) (err error) {
	if body.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("messages", "body"))
	}
	for _, e := range body.Messages {
		if e != nil {
			if err2 := ValidatePublishedmessageResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.at-time", *body.AtTime, goa.FormatDateTime))
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

// ValidateCreateAlreadyCreatedResponseBody runs the validations defined on
// create_already-created_response_body
func ValidateCreateAlreadyCreatedResponseBody(body *CreateAlreadyCreatedResponseBody) (err error) {
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

// ValidateEnqueueBadRequestResponseBody runs the validations defined on
// enqueue_bad-request_response_body
func ValidateEnqueueBadRequestResponseBody(body *EnqueueBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateEnqueueInvalidParameterResponseBody runs the validations defined on
// enqueue_invalid-parameter_response_body
func ValidateEnqueueInvalidParameterResponseBody(body *EnqueueInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateEnqueueInvalidScopesResponseBody runs the validations defined on
// enqueue_invalid-scopes_response_body
func ValidateEnqueueInvalidScopesResponseBody(body *EnqueueInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateEnqueueNotImplementedResponseBody runs the validations defined on
// enqueue_not-implemented_response_body
func ValidateEnqueueNotImplementedResponseBody(body *EnqueueNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDequeueBadRequestResponseBody runs the validations defined on
// dequeue_bad-request_response_body
func ValidateDequeueBadRequestResponseBody(body *DequeueBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDequeueInvalidParameterResponseBody runs the validations defined on
// dequeue_invalid-parameter_response_body
func ValidateDequeueInvalidParameterResponseBody(body *DequeueInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDequeueInvalidScopesResponseBody runs the validations defined on
// dequeue_invalid-scopes_response_body
func ValidateDequeueInvalidScopesResponseBody(body *DequeueInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateDequeueNotImplementedResponseBody runs the validations defined on
// dequeue_not-implemented_response_body
func ValidateDequeueNotImplementedResponseBody(body *DequeueNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateQueueListItemResponseBody runs the validations defined on
// QueueListItemResponseBody
func ValidateQueueListItemResponseBody(body *QueueListItemResponseBody) (err error) {
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
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
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

// ValidatePublishedmessageResponseBody runs the validations defined on
// PublishedmessageResponseBody
func ValidatePublishedmessageResponseBody(body *PublishedmessageResponseBody) (err error) {
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	return
}
