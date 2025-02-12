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

package queue

import (
	"context"

	queueviews "github.com/ivcap-works/ivcap-core-api/gen/queue/views"
	"goa.design/goa/v3/security"
)

// Manage the life cycle of a queue.
type Service interface {
	// Create a new queues and return its status.
	Create(context.Context, *CreatePayload) (res *Createqueueresponse, err error)
	// Show queues by ID
	Read(context.Context, *ReadPayload) (res *Readqueueresponse, err error)
	// Delete an existing queues.
	Delete(context.Context, *DeletePayload) (err error)
	// list queues
	List(context.Context, *ListPayload) (res *QueueListResult, err error)
	// Send a message to a specific queues.
	Enqueue(context.Context, *EnqueuePayload) (res *Messagestatus, err error)
	// Read a message from a specific queues.
	Dequeue(context.Context, *DequeuePayload) (res *MessageList, err error)
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
const ServiceName = "queue"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"create", "read", "delete", "list", "enqueue", "dequeue"}

// Something wasn't right with this request
type BadRequestT struct {
	// Information message
	Message string
}

// CreatePayload is the payload type of the queue service create method.
type CreatePayload struct {
	// New queues description
	Queues *PayloadForCreateEndpoint
	// JWT used for authentication
	JWT string
}

// Createqueueresponse is the result type of the queue service create method.
type Createqueueresponse struct {
	// queue
	ID string
	// Name of the created queue.
	Name string
	// Description of the created queue.
	Description *string
	// Reference to billable account
	Account *string
}

// DeletePayload is the payload type of the queue service delete method.
type DeletePayload struct {
	// ID of queues to update
	ID string
	// JWT used for authentication
	JWT string
}

// DequeuePayload is the payload type of the queue service dequeue method.
type DequeuePayload struct {
	// queue
	ID string
	// Maximum number of messages to dequeue
	Limit *int
	// JWT used for authentication
	JWT string
}

// EnqueuePayload is the payload type of the queue service enqueue method.
type EnqueuePayload struct {
	// queue
	ID string
	// Content-Type header, MUST be of application/json.
	ContentType *string `json:"content-type,omitempty"`
	// Schema used for message
	Schema *string
	// Message content
	Content any `json:"content,omitempty"`
	// JWT used for authentication
	JWT string
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

type LinkT struct {
	// relation type
	Rel string
	// mime type
	Type string
	// web link
	Href string
}

// ListPayload is the payload type of the queue service list method.
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

// MessageList is the result type of the queue service dequeue method.
type MessageList struct {
	// Messages in the queue
	Messages []*Publishedmessage
	// Time at which this list was valid
	AtTime *string
}

// Messagestatus is the result type of the queue service enqueue method.
type Messagestatus struct {
	// queue
	ID *string
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

type PayloadForCreateEndpoint struct {
	// Optional Name for the queue. Cannot contain whitespace, ., *, >, path
	// separators (forward or backwards slash), and non-printable characters.
	Name string
	// More detailed description of the queue.
	Description *string
	// Reference to policy used
	Policy *string
}

type Publishedmessage struct {
	// Message identifier
	ID *string
	// Message content in JSON format.
	Content any
	// Schema used for message
	Schema *string
	// Encoding type of message content (defaults to 'application/json')
	ContentType *string
}

type QueueListItem struct {
	// queue
	ID string
	// Name of the created queue.
	Name *string
	// Description of the created queue.
	Description *string
	// Reference to billable account
	Account string
	Href    string
}

// QueueListResult is the result type of the queue service list method.
type QueueListResult struct {
	// Queues
	Items []*QueueListItem
	// Time at which this list was valid
	AtTime string
	Links  []*LinkT
}

// ReadPayload is the payload type of the queue service read method.
type ReadPayload struct {
	// ID of queues to show
	ID string
	// JWT used for authentication
	JWT string
}

// Readqueueresponse is the result type of the queue service read method.
type Readqueueresponse struct {
	// ID
	ID string
	// Name of the queue.
	Name string
	// Description of the queue.
	Description *string
	// Number of messages sent to the queue
	TotalMessages *uint64
	// Number of bytes in the queue
	Bytes *uint64
	// Timestamp of the first message in the queue
	FirstTime *string
	// Timestamp of the last message in the queue
	LastTime *string
	// Number of consumers
	ConsumerCount *int
	// Timestamp when the queue was created
	CreatedAt string
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

// NewCreatequeueresponse initializes result type Createqueueresponse from
// viewed result type Createqueueresponse.
func NewCreatequeueresponse(vres *queueviews.Createqueueresponse) *Createqueueresponse {
	return newCreatequeueresponse(vres.Projected)
}

// NewViewedCreatequeueresponse initializes viewed result type
// Createqueueresponse from result type Createqueueresponse using the given
// view.
func NewViewedCreatequeueresponse(res *Createqueueresponse, view string) *queueviews.Createqueueresponse {
	p := newCreatequeueresponseView(res)
	return &queueviews.Createqueueresponse{Projected: p, View: "default"}
}

// NewReadqueueresponse initializes result type Readqueueresponse from viewed
// result type Readqueueresponse.
func NewReadqueueresponse(vres *queueviews.Readqueueresponse) *Readqueueresponse {
	return newReadqueueresponse(vres.Projected)
}

// NewViewedReadqueueresponse initializes viewed result type Readqueueresponse
// from result type Readqueueresponse using the given view.
func NewViewedReadqueueresponse(res *Readqueueresponse, view string) *queueviews.Readqueueresponse {
	p := newReadqueueresponseView(res)
	return &queueviews.Readqueueresponse{Projected: p, View: "default"}
}

// NewMessagestatus initializes result type Messagestatus from viewed result
// type Messagestatus.
func NewMessagestatus(vres *queueviews.Messagestatus) *Messagestatus {
	return newMessagestatus(vres.Projected)
}

// NewViewedMessagestatus initializes viewed result type Messagestatus from
// result type Messagestatus using the given view.
func NewViewedMessagestatus(res *Messagestatus, view string) *queueviews.Messagestatus {
	p := newMessagestatusView(res)
	return &queueviews.Messagestatus{Projected: p, View: "default"}
}

// newCreatequeueresponse converts projected type Createqueueresponse to
// service type Createqueueresponse.
func newCreatequeueresponse(vres *queueviews.CreatequeueresponseView) *Createqueueresponse {
	res := &Createqueueresponse{
		Description: vres.Description,
		Account:     vres.Account,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newCreatequeueresponseView projects result type Createqueueresponse to
// projected type CreatequeueresponseView using the "default" view.
func newCreatequeueresponseView(res *Createqueueresponse) *queueviews.CreatequeueresponseView {
	vres := &queueviews.CreatequeueresponseView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: res.Description,
		Account:     res.Account,
	}
	return vres
}

// newReadqueueresponse converts projected type Readqueueresponse to service
// type Readqueueresponse.
func newReadqueueresponse(vres *queueviews.ReadqueueresponseView) *Readqueueresponse {
	res := &Readqueueresponse{
		Description:   vres.Description,
		TotalMessages: vres.TotalMessages,
		Bytes:         vres.Bytes,
		FirstTime:     vres.FirstTime,
		LastTime:      vres.LastTime,
		ConsumerCount: vres.ConsumerCount,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	return res
}

// newReadqueueresponseView projects result type Readqueueresponse to projected
// type ReadqueueresponseView using the "default" view.
func newReadqueueresponseView(res *Readqueueresponse) *queueviews.ReadqueueresponseView {
	vres := &queueviews.ReadqueueresponseView{
		ID:            &res.ID,
		Name:          &res.Name,
		Description:   res.Description,
		TotalMessages: res.TotalMessages,
		Bytes:         res.Bytes,
		FirstTime:     res.FirstTime,
		LastTime:      res.LastTime,
		ConsumerCount: res.ConsumerCount,
		CreatedAt:     &res.CreatedAt,
	}
	return vres
}

// newMessagestatus converts projected type Messagestatus to service type
// Messagestatus.
func newMessagestatus(vres *queueviews.MessagestatusView) *Messagestatus {
	res := &Messagestatus{
		ID: vres.ID,
	}
	return res
}

// newMessagestatusView projects result type Messagestatus to projected type
// MessagestatusView using the "default" view.
func newMessagestatusView(res *Messagestatus) *queueviews.MessagestatusView {
	vres := &queueviews.MessagestatusView{
		ID: res.ID,
	}
	return vres
}
