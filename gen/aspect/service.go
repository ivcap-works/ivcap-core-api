// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package aspect

import (
	aspectviews "github.com/reinventingscience/ivcap-core-api/gen/aspect/views"
	"context"

	"goa.design/goa/v3/security"
)

// Manages the life cycle of aspect(s) attached to an entity.
type Service interface {
	// Show aspect by ID
	Read(context.Context, *ReadPayload) (res *AspectRT, err error)
	// Return a list of aspect aspects.
	List(context.Context, *ListPayload) (res *AspectListRT, err error)
	// Attach new aspect to an entity.
	Create(context.Context, *CreatePayload) (res *AspectIDRT, err error)
	// Retract this aspect and create a new one with the information provided.
	// For any field not provided, the value from the current aspect is used.
	Update(context.Context, *UpdatePayload) (res *AspectIDRT, err error)
	// Retract a previously created statement.
	Retract(context.Context, *RetractPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "aspect"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"read", "list", "create", "update", "retract"}

// AspectIDRT is the result type of the aspect service create method.
type AspectIDRT struct {
	// ID to specific aspect
	ID string
}

type AspectListItemRT struct {
	// Record URN
	ID string
	// Entity URN
	Entity string
	// Schema URN
	Schema string
	// Attached aspect aspect
	Content any
	// Content-Type header, MUST be of application/json.
	ContentType string `json:"content-type,omitempty"`
}

// AspectListRT is the result type of the aspect service list method.
type AspectListRT struct {
	// List of aspect descriptions
	Items []*AspectListItemRT
	// Entity for which to request aspect
	Entity *string
	// Optional schema to filter on
	Schema *string
	// Optional json path to further filter on returned list
	AspectPath *string
	// Time at which this list was valid
	AtTime string
	// Navigation links
	Links *NavT
}

// AspectRT is the result type of the aspect service read method.
type AspectRT struct {
	// Record URN
	ID *string
	// Entity URN
	Entity *string
	// Schema URN
	Schema *string
	// Description of aspect encoded as 'content-type'
	Content any
	// Content-Type header, MUST be of application/json.
	ContentType *string `json:"content-type,omitempty"`
	// Time this aspect was asserted
	ValidFrom *string
	// Time this aspect was retractd
	ValidTo *string
	// Entity asserting this aspect aspect at 'valid-from'
	Asserter *string
	// Entity retracting this aspect at 'valid-to'
	Retracter *string
}

// Bad arguments supplied.
type BadRequestT struct {
	// Information message
	Message string
}

// CreatePayload is the payload type of the aspect service create method.
type CreatePayload struct {
	// Entity to which attach aspect
	Entity string `json:"entity,omitempty"`
	// Schema of the aspect in payload
	Schema string
	// Aspect content
	Content any
	// Content-Type header, MUST be of application/json.
	ContentType string `json:"content-type,omitempty"`
	// Policy guiding visibility and actions performed
	Policy *string `json:"policy,omitempty"`
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

// ListPayload is the payload type of the aspect service list method.
type ListPayload struct {
	// Optioanl entity for which to request aspects
	Entity *string `json:"entity,omitempty"`
	// Schema prefix using '%' as wildcard indicator
	Schema *string
	// To learn more about the supported format, see
	// https://www.postgresql.org/docs/current/datatype-json.html#DATATYPE-JSONPATH
	ContentPath *string `json:"aspect-path,omitempty"`
	// Return aspect which where valid at that time [now]
	AtTime *string `json:"at-time,omitempty"`
	// The 'limit' system query option requests the number of items in the queried
	// collection to be included in the result.
	Limit int
	// The 'filter' system query option allows clients to filter a collection of
	// resources that are addressed by a request URL. The expression specified with
	// 'filter'
	// is evaluated for each resource in the collection, and only items where the
	// expression
	// evaluates to true are included in the response.
	Filter string
	// The 'orderby' query option allows clients to request resources in either
	// ascending order using asc or descending order using desc. If asc or desc not
	// specified,
	// then the resources will be ordered in ascending order. The request below
	// orders Trips on
	// property EndsAt in descending order.
	OrderBy string `json:"order-by,omitempty"`
	// When set order result in descending order. Ascending order is the default.
	OrderDesc *bool `json:"order-desc,omitempty"`
	// The content of '$page' is returned in the 'links' part of a previous query
	// and
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

// ReadPayload is the payload type of the aspect service read method.
type ReadPayload struct {
	// ID of aspect to show
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

// RetractPayload is the payload type of the aspect service retract method.
type RetractPayload struct {
	// Aspect ID to restract
	ID string
	// JWT used for authentication
	JWT string
}

// ServiceNotAvailable is the type returned when the service necessary to
// fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// Unauthorized access to resource
type UnauthorizedT struct {
}

// UnsupportedContentType is the error returned when the provided content type
// is not supported.
type UnsupportedContentType struct {
	// message describing expected type or pattern.
	Message string
}

// UpdatePayload is the payload type of the aspect service update method.
type UpdatePayload struct {
	// Aspect to update
	ID string
	// Entity to which attach aspect
	Entity string `json:"entity,omitempty"`
	// Schema of aspect
	Schema string
	// Aspect content
	Content any
	// Content-Type header, MUST be of application/json.
	ContentType string `json:"content-type,omitempty"`
	// JWT used for authentication
	JWT string
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

// Error returns an error description.
func (e *UnsupportedContentType) Error() string {
	return "UnsupportedContentType is the error returned when the provided content type is not supported."
}

// ErrorName returns "UnsupportedContentType".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *UnsupportedContentType) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "UnsupportedContentType".
func (e *UnsupportedContentType) GoaErrorName() string {
	return "unsupported-content-type"
}

// NewAspectRT initializes result type AspectRT from viewed result type
// AspectRT.
func NewAspectRT(vres *aspectviews.AspectRT) *AspectRT {
	return newAspectRT(vres.Projected)
}

// NewViewedAspectRT initializes viewed result type AspectRT from result type
// AspectRT using the given view.
func NewViewedAspectRT(res *AspectRT, view string) *aspectviews.AspectRT {
	p := newAspectRTView(res)
	return &aspectviews.AspectRT{Projected: p, View: "default"}
}

// NewAspectListRT initializes result type AspectListRT from viewed result type
// AspectListRT.
func NewAspectListRT(vres *aspectviews.AspectListRT) *AspectListRT {
	return newAspectListRT(vres.Projected)
}

// NewViewedAspectListRT initializes viewed result type AspectListRT from
// result type AspectListRT using the given view.
func NewViewedAspectListRT(res *AspectListRT, view string) *aspectviews.AspectListRT {
	p := newAspectListRTView(res)
	return &aspectviews.AspectListRT{Projected: p, View: "default"}
}

// NewAspectIDRT initializes result type AspectIDRT from viewed result type
// AspectIDRT.
func NewAspectIDRT(vres *aspectviews.AspectIDRT) *AspectIDRT {
	return newAspectIDRT(vres.Projected)
}

// NewViewedAspectIDRT initializes viewed result type AspectIDRT from result
// type AspectIDRT using the given view.
func NewViewedAspectIDRT(res *AspectIDRT, view string) *aspectviews.AspectIDRT {
	p := newAspectIDRTView(res)
	return &aspectviews.AspectIDRT{Projected: p, View: "default"}
}

// newAspectRT converts projected type AspectRT to service type AspectRT.
func newAspectRT(vres *aspectviews.AspectRTView) *AspectRT {
	res := &AspectRT{
		ID:          vres.ID,
		Entity:      vres.Entity,
		Schema:      vres.Schema,
		Content:     vres.Content,
		ContentType: vres.ContentType,
		ValidFrom:   vres.ValidFrom,
		ValidTo:     vres.ValidTo,
		Asserter:    vres.Asserter,
		Retracter:   vres.Retracter,
	}
	return res
}

// newAspectRTView projects result type AspectRT to projected type AspectRTView
// using the "default" view.
func newAspectRTView(res *AspectRT) *aspectviews.AspectRTView {
	vres := &aspectviews.AspectRTView{
		ID:          res.ID,
		Entity:      res.Entity,
		Schema:      res.Schema,
		Content:     res.Content,
		ContentType: res.ContentType,
		ValidFrom:   res.ValidFrom,
		ValidTo:     res.ValidTo,
		Asserter:    res.Asserter,
		Retracter:   res.Retracter,
	}
	return vres
}

// newAspectListRT converts projected type AspectListRT to service type
// AspectListRT.
func newAspectListRT(vres *aspectviews.AspectListRTView) *AspectListRT {
	res := &AspectListRT{
		Entity:     vres.Entity,
		Schema:     vres.Schema,
		AspectPath: vres.AspectPath,
	}
	if vres.AtTime != nil {
		res.AtTime = *vres.AtTime
	}
	if vres.Items != nil {
		res.Items = make([]*AspectListItemRT, len(vres.Items))
		for i, val := range vres.Items {
			res.Items[i] = transformAspectviewsAspectListItemRTViewToAspectListItemRT(val)
		}
	}
	if vres.Links != nil {
		res.Links = transformAspectviewsNavTViewToNavT(vres.Links)
	}
	return res
}

// newAspectListRTView projects result type AspectListRT to projected type
// AspectListRTView using the "default" view.
func newAspectListRTView(res *AspectListRT) *aspectviews.AspectListRTView {
	vres := &aspectviews.AspectListRTView{
		Entity:     res.Entity,
		Schema:     res.Schema,
		AspectPath: res.AspectPath,
		AtTime:     &res.AtTime,
	}
	if res.Items != nil {
		vres.Items = make([]*aspectviews.AspectListItemRTView, len(res.Items))
		for i, val := range res.Items {
			vres.Items[i] = transformAspectListItemRTToAspectviewsAspectListItemRTView(val)
		}
	} else {
		vres.Items = []*aspectviews.AspectListItemRTView{}
	}
	if res.Links != nil {
		vres.Links = transformNavTToAspectviewsNavTView(res.Links)
	}
	return vres
}

// newAspectListItemRT converts projected type AspectListItemRT to service type
// AspectListItemRT.
func newAspectListItemRT(vres *aspectviews.AspectListItemRTView) *AspectListItemRT {
	res := &AspectListItemRT{
		Content: vres.Content,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Entity != nil {
		res.Entity = *vres.Entity
	}
	if vres.Schema != nil {
		res.Schema = *vres.Schema
	}
	if vres.ContentType != nil {
		res.ContentType = *vres.ContentType
	}
	return res
}

// newAspectListItemRTView projects result type AspectListItemRT to projected
// type AspectListItemRTView using the "default" view.
func newAspectListItemRTView(res *AspectListItemRT) *aspectviews.AspectListItemRTView {
	vres := &aspectviews.AspectListItemRTView{
		ID:          &res.ID,
		Entity:      &res.Entity,
		Schema:      &res.Schema,
		Content:     res.Content,
		ContentType: &res.ContentType,
	}
	return vres
}

// newAspectIDRT converts projected type AspectIDRT to service type AspectIDRT.
func newAspectIDRT(vres *aspectviews.AspectIDRTView) *AspectIDRT {
	res := &AspectIDRT{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	return res
}

// newAspectIDRTView projects result type AspectIDRT to projected type
// AspectIDRTView using the "default" view.
func newAspectIDRTView(res *AspectIDRT) *aspectviews.AspectIDRTView {
	vres := &aspectviews.AspectIDRTView{
		ID: &res.ID,
	}
	return vres
}

// transformAspectviewsAspectListItemRTViewToAspectListItemRT builds a value of
// type *AspectListItemRT from a value of type
// *aspectviews.AspectListItemRTView.
func transformAspectviewsAspectListItemRTViewToAspectListItemRT(v *aspectviews.AspectListItemRTView) *AspectListItemRT {
	if v == nil {
		return nil
	}
	res := &AspectListItemRT{
		ID:          *v.ID,
		Entity:      *v.Entity,
		Schema:      *v.Schema,
		Content:     v.Content,
		ContentType: *v.ContentType,
	}

	return res
}

// transformAspectviewsNavTViewToNavT builds a value of type *NavT from a value
// of type *aspectviews.NavTView.
func transformAspectviewsNavTViewToNavT(v *aspectviews.NavTView) *NavT {
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

// transformAspectListItemRTToAspectviewsAspectListItemRTView builds a value of
// type *aspectviews.AspectListItemRTView from a value of type
// *AspectListItemRT.
func transformAspectListItemRTToAspectviewsAspectListItemRTView(v *AspectListItemRT) *aspectviews.AspectListItemRTView {
	res := &aspectviews.AspectListItemRTView{
		ID:          &v.ID,
		Entity:      &v.Entity,
		Schema:      &v.Schema,
		Content:     v.Content,
		ContentType: &v.ContentType,
	}

	return res
}

// transformNavTToAspectviewsNavTView builds a value of type
// *aspectviews.NavTView from a value of type *NavT.
func transformNavTToAspectviewsNavTView(v *NavT) *aspectviews.NavTView {
	res := &aspectviews.NavTView{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}
