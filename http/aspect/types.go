// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	aspect "github.com/reinventingscience/ivcap-core-api/gen/aspect"
	aspectviews "github.com/reinventingscience/ivcap-core-api/gen/aspect/views"

	goa "goa.design/goa/v3/pkg"
)

// ListRequestBody is the type of the "aspect" service "list" endpoint HTTP
// request body.
type ListRequestBody struct {
	// To learn more about the supported format, see
	// https://www.postgresql.org/docs/current/datatype-json.html#DATATYPE-JSONPATH
	ContentPath *string `json:"aspect-path,omitempty"`
}

// ReadResponseBody is the type of the "aspect" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// Record URN
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Entity URN
	Entity *string `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
	// Schema URN
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Description of aspect encoded as 'content-type'
	Content any `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Content-Type header, MUST be of application/json.
	ContentType *string `json:"content-type,omitempty"`
	// Time this aspect was asserted
	ValidFrom *string `form:"valid-from,omitempty" json:"valid-from,omitempty" xml:"valid-from,omitempty"`
	// Time this aspect was retractd
	ValidTo *string `form:"valid-to,omitempty" json:"valid-to,omitempty" xml:"valid-to,omitempty"`
	// Entity asserting this aspect aspect at 'valid-from'
	Asserter *string `form:"asserter,omitempty" json:"asserter,omitempty" xml:"asserter,omitempty"`
	// Entity retracting this aspect at 'valid-to'
	Retracter *string `form:"retracter,omitempty" json:"retracter,omitempty" xml:"retracter,omitempty"`
}

// ListResponseBody is the type of the "aspect" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of aspect descriptions
	Items []*AspectListItemRTResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Entity for which to request aspect
	Entity *string `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
	// Optional schema to filter on
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Optional json path to further filter on returned list
	AspectPath *string `form:"aspect-path,omitempty" json:"aspect-path,omitempty" xml:"aspect-path,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	// Navigation links
	Links *NavTResponseBody `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// CreateResponseBody is the type of the "aspect" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID to specific aspect
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// UpdateResponseBody is the type of the "aspect" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID to specific aspect
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "aspect" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "aspect" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "aspect" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "aspect" service "read" endpoint
// HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListBadRequestResponseBody is the type of the "aspect" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "aspect" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "aspect" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "aspect" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListUnsupportedContentTypeResponseBody is the type of the "aspect" service
// "list" endpoint HTTP response body for the "unsupported-content-type" error.
type ListUnsupportedContentTypeResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "aspect" service "create"
// endpoint HTTP response body for the "bad-request" error.
type CreateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateInvalidParameterResponseBody is the type of the "aspect" service
// "create" endpoint HTTP response body for the "invalid-parameter" error.
type CreateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// CreateInvalidScopesResponseBody is the type of the "aspect" service "create"
// endpoint HTTP response body for the "invalid-scopes" error.
type CreateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateNotImplementedResponseBody is the type of the "aspect" service
// "create" endpoint HTTP response body for the "not-implemented" error.
type CreateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateBadRequestResponseBody is the type of the "aspect" service "update"
// endpoint HTTP response body for the "bad-request" error.
type UpdateBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateInvalidParameterResponseBody is the type of the "aspect" service
// "update" endpoint HTTP response body for the "invalid-parameter" error.
type UpdateInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// UpdateInvalidScopesResponseBody is the type of the "aspect" service "update"
// endpoint HTTP response body for the "invalid-scopes" error.
type UpdateInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateNotImplementedResponseBody is the type of the "aspect" service
// "update" endpoint HTTP response body for the "not-implemented" error.
type UpdateNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RetractBadRequestResponseBody is the type of the "aspect" service "retract"
// endpoint HTTP response body for the "bad-request" error.
type RetractBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RetractInvalidParameterResponseBody is the type of the "aspect" service
// "retract" endpoint HTTP response body for the "invalid-parameter" error.
type RetractInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// RetractInvalidScopesResponseBody is the type of the "aspect" service
// "retract" endpoint HTTP response body for the "invalid-scopes" error.
type RetractInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RetractNotImplementedResponseBody is the type of the "aspect" service
// "retract" endpoint HTTP response body for the "not-implemented" error.
type RetractNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// AspectListItemRTResponseBody is used to define fields on response body types.
type AspectListItemRTResponseBody struct {
	// Record URN
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Entity URN
	Entity *string `form:"entity,omitempty" json:"entity,omitempty" xml:"entity,omitempty"`
	// Schema URN
	Schema *string `form:"schema,omitempty" json:"schema,omitempty" xml:"schema,omitempty"`
	// Attached aspect aspect
	Content any `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Content-Type header, MUST be of application/json.
	ContentType *string `json:"content-type,omitempty"`
}

// NavTResponseBody is used to define fields on response body types.
type NavTResponseBody struct {
	Self  *string `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
	First *string `form:"first,omitempty" json:"first,omitempty" xml:"first,omitempty"`
	Next  *string `form:"next,omitempty" json:"next,omitempty" xml:"next,omitempty"`
}

// NewListRequestBody builds the HTTP request body from the payload of the
// "list" endpoint of the "aspect" service.
func NewListRequestBody(p *aspect.ListPayload) *ListRequestBody {
	body := &ListRequestBody{
		ContentPath: p.ContentPath,
	}
	return body
}

// NewReadAspectRTOK builds a "aspect" service "read" endpoint result from a
// HTTP "OK" response.
func NewReadAspectRTOK(body *ReadResponseBody) *aspectviews.AspectRTView {
	v := &aspectviews.AspectRTView{
		ID:          body.ID,
		Entity:      body.Entity,
		Schema:      body.Schema,
		Content:     body.Content,
		ContentType: body.ContentType,
		ValidFrom:   body.ValidFrom,
		ValidTo:     body.ValidTo,
		Asserter:    body.Asserter,
		Retracter:   body.Retracter,
	}

	return v
}

// NewReadBadRequest builds a aspect service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *aspect.BadRequestT {
	v := &aspect.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidCredential builds a aspect service read endpoint
// invalid-credential error.
func NewReadInvalidCredential() *aspect.InvalidCredentialsT {
	v := &aspect.InvalidCredentialsT{}

	return v
}

// NewReadInvalidScopes builds a aspect service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *aspect.InvalidScopesT {
	v := &aspect.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a aspect service read endpoint not-implemented
// error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *aspect.NotImplementedT {
	v := &aspect.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a aspect service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *aspect.ResourceNotFoundT {
	v := &aspect.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAuthorized builds a aspect service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *aspect.UnauthorizedT {
	v := &aspect.UnauthorizedT{}

	return v
}

// NewListAspectListRTOK builds a "aspect" service "list" endpoint result from
// a HTTP "OK" response.
func NewListAspectListRTOK(body *ListResponseBody) *aspectviews.AspectListRTView {
	v := &aspectviews.AspectListRTView{
		Entity:     body.Entity,
		Schema:     body.Schema,
		AspectPath: body.AspectPath,
		AtTime:     body.AtTime,
	}
	v.Items = make([]*aspectviews.AspectListItemRTView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalAspectListItemRTResponseBodyToAspectviewsAspectListItemRTView(val)
	}
	v.Links = unmarshalNavTResponseBodyToAspectviewsNavTView(body.Links)

	return v
}

// NewListBadRequest builds a aspect service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *aspect.BadRequestT {
	v := &aspect.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidCredential builds a aspect service list endpoint
// invalid-credential error.
func NewListInvalidCredential() *aspect.InvalidCredentialsT {
	v := &aspect.InvalidCredentialsT{}

	return v
}

// NewListInvalidParameter builds a aspect service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *aspect.InvalidParameterValue {
	v := &aspect.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a aspect service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *aspect.InvalidScopesT {
	v := &aspect.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a aspect service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *aspect.NotImplementedT {
	v := &aspect.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAuthorized builds a aspect service list endpoint not-authorized
// error.
func NewListNotAuthorized() *aspect.UnauthorizedT {
	v := &aspect.UnauthorizedT{}

	return v
}

// NewListUnsupportedContentType builds a aspect service list endpoint
// unsupported-content-type error.
func NewListUnsupportedContentType(body *ListUnsupportedContentTypeResponseBody) *aspect.UnsupportedContentType {
	v := &aspect.UnsupportedContentType{
		Message: *body.Message,
	}

	return v
}

// NewCreateAspectIDRTOK builds a "aspect" service "create" endpoint result
// from a HTTP "OK" response.
func NewCreateAspectIDRTOK(body *CreateResponseBody) *aspectviews.AspectIDRTView {
	v := &aspectviews.AspectIDRTView{
		ID: body.ID,
	}

	return v
}

// NewCreateBadRequest builds a aspect service create endpoint bad-request
// error.
func NewCreateBadRequest(body *CreateBadRequestResponseBody) *aspect.BadRequestT {
	v := &aspect.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewCreateInvalidCredential builds a aspect service create endpoint
// invalid-credential error.
func NewCreateInvalidCredential() *aspect.InvalidCredentialsT {
	v := &aspect.InvalidCredentialsT{}

	return v
}

// NewCreateInvalidParameter builds a aspect service create endpoint
// invalid-parameter error.
func NewCreateInvalidParameter(body *CreateInvalidParameterResponseBody) *aspect.InvalidParameterValue {
	v := &aspect.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewCreateInvalidScopes builds a aspect service create endpoint
// invalid-scopes error.
func NewCreateInvalidScopes(body *CreateInvalidScopesResponseBody) *aspect.InvalidScopesT {
	v := &aspect.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateNotImplemented builds a aspect service create endpoint
// not-implemented error.
func NewCreateNotImplemented(body *CreateNotImplementedResponseBody) *aspect.NotImplementedT {
	v := &aspect.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewCreateNotAuthorized builds a aspect service create endpoint
// not-authorized error.
func NewCreateNotAuthorized() *aspect.UnauthorizedT {
	v := &aspect.UnauthorizedT{}

	return v
}

// NewUpdateAspectIDRTOK builds a "aspect" service "update" endpoint result
// from a HTTP "OK" response.
func NewUpdateAspectIDRTOK(body *UpdateResponseBody) *aspectviews.AspectIDRTView {
	v := &aspectviews.AspectIDRTView{
		ID: body.ID,
	}

	return v
}

// NewUpdateBadRequest builds a aspect service update endpoint bad-request
// error.
func NewUpdateBadRequest(body *UpdateBadRequestResponseBody) *aspect.BadRequestT {
	v := &aspect.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateInvalidCredential builds a aspect service update endpoint
// invalid-credential error.
func NewUpdateInvalidCredential() *aspect.InvalidCredentialsT {
	v := &aspect.InvalidCredentialsT{}

	return v
}

// NewUpdateInvalidParameter builds a aspect service update endpoint
// invalid-parameter error.
func NewUpdateInvalidParameter(body *UpdateInvalidParameterResponseBody) *aspect.InvalidParameterValue {
	v := &aspect.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewUpdateInvalidScopes builds a aspect service update endpoint
// invalid-scopes error.
func NewUpdateInvalidScopes(body *UpdateInvalidScopesResponseBody) *aspect.InvalidScopesT {
	v := &aspect.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateNotImplemented builds a aspect service update endpoint
// not-implemented error.
func NewUpdateNotImplemented(body *UpdateNotImplementedResponseBody) *aspect.NotImplementedT {
	v := &aspect.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateNotAuthorized builds a aspect service update endpoint
// not-authorized error.
func NewUpdateNotAuthorized() *aspect.UnauthorizedT {
	v := &aspect.UnauthorizedT{}

	return v
}

// NewRetractBadRequest builds a aspect service retract endpoint bad-request
// error.
func NewRetractBadRequest(body *RetractBadRequestResponseBody) *aspect.BadRequestT {
	v := &aspect.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewRetractInvalidCredential builds a aspect service retract endpoint
// invalid-credential error.
func NewRetractInvalidCredential() *aspect.InvalidCredentialsT {
	v := &aspect.InvalidCredentialsT{}

	return v
}

// NewRetractInvalidParameter builds a aspect service retract endpoint
// invalid-parameter error.
func NewRetractInvalidParameter(body *RetractInvalidParameterResponseBody) *aspect.InvalidParameterValue {
	v := &aspect.InvalidParameterValue{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewRetractInvalidScopes builds a aspect service retract endpoint
// invalid-scopes error.
func NewRetractInvalidScopes(body *RetractInvalidScopesResponseBody) *aspect.InvalidScopesT {
	v := &aspect.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewRetractNotImplemented builds a aspect service retract endpoint
// not-implemented error.
func NewRetractNotImplemented(body *RetractNotImplementedResponseBody) *aspect.NotImplementedT {
	v := &aspect.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewRetractNotAuthorized builds a aspect service retract endpoint
// not-authorized error.
func NewRetractNotAuthorized() *aspect.UnauthorizedT {
	v := &aspect.UnauthorizedT{}

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

// ValidateListUnsupportedContentTypeResponseBody runs the validations defined
// on list_unsupported-content-type_response_body
func ValidateListUnsupportedContentTypeResponseBody(body *ListUnsupportedContentTypeResponseBody) (err error) {
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

// ValidateUpdateBadRequestResponseBody runs the validations defined on
// update_bad-request_response_body
func ValidateUpdateBadRequestResponseBody(body *UpdateBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateInvalidParameterResponseBody runs the validations defined on
// update_invalid-parameter_response_body
func ValidateUpdateInvalidParameterResponseBody(body *UpdateInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateInvalidScopesResponseBody runs the validations defined on
// update_invalid-scopes_response_body
func ValidateUpdateInvalidScopesResponseBody(body *UpdateInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateNotImplementedResponseBody runs the validations defined on
// update_not-implemented_response_body
func ValidateUpdateNotImplementedResponseBody(body *UpdateNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRetractBadRequestResponseBody runs the validations defined on
// retract_bad-request_response_body
func ValidateRetractBadRequestResponseBody(body *RetractBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRetractInvalidParameterResponseBody runs the validations defined on
// retract_invalid-parameter_response_body
func ValidateRetractInvalidParameterResponseBody(body *RetractInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRetractInvalidScopesResponseBody runs the validations defined on
// retract_invalid-scopes_response_body
func ValidateRetractInvalidScopesResponseBody(body *RetractInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateRetractNotImplementedResponseBody runs the validations defined on
// retract_not-implemented_response_body
func ValidateRetractNotImplementedResponseBody(body *RetractNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateAspectListItemRTResponseBody runs the validations defined on
// AspectListItemRTResponseBody
func ValidateAspectListItemRTResponseBody(body *AspectListItemRTResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Entity == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("entity", "body"))
	}
	if body.Schema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("schema", "body"))
	}
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	if body.ContentType == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content-type", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatURI))
	}
	if body.Entity != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.entity", *body.Entity, goa.FormatURI))
	}
	if body.Schema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.schema", *body.Schema, goa.FormatURI))
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
