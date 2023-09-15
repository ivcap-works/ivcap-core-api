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

package metadata

import (
	"context"

	metadataviews "github.com/reinventingscience/ivcap-core-api/gen/metadata/views"

	"goa.design/goa/v3/security"
)

// Manages the life cycle of metadata attached to an entity.
type Service interface {
	// Show metadata by ID
	Read(context.Context, *ReadPayload) (res *MetadataRecordRT, err error)
	// Return a list of metadata records.
	List(context.Context, *ListPayload) (res *ListMetaRT, err error)
	// Attach new metadata to an entity.
	Add(context.Context, *AddPayload) (res *AddMetaRT, err error)
	// Revoke a record for the same entity and same schema and create new one
	// with the provided properties. __NOTE__, this method will fail if there is
	// more than one active record for the entity/schema pair.
	UpdateOne(context.Context, *UpdateOnePayload) (res *AddMetaRT, err error)
	// Revoke this record and create a new one with the information provided.
	// For any field not provided, the value from the current record is used.
	UpdateRecord(context.Context, *UpdateRecordPayload) (res *AddMetaRT, err error)
	// Retract a previously created statement.
	Revoke(context.Context, *RevokePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "metadata"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"read", "list", "add", "update_one", "update_record", "revoke"}

// AddMetaRT is the result type of the metadata service add method.
type AddMetaRT struct {
	// Reference to record created
	RecordID string
}

// AddPayload is the payload type of the metadata service add method.
type AddPayload struct {
	// Entity to which attach metadata
	EntityID string
	// Schema of metadata
	Schema string
	// Aspect content
	Aspect interface{}
	// Content-Type header, MUST be of application/json.
	ContentType string
	// Policy guiding visibility and actions performed
	PolicyID *string
	// JWT used for authentication
	JWT string
}

// Bad arguments supplied.
type BadRequestT struct {
	// Information message
	Message string
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

// ListMetaRT is the result type of the metadata service list method.
type ListMetaRT struct {
	// List of metadata records
	Records []*MetadataListItemRT
	// Entity for which to request metadata
	EntityID *string
	// Optional schema to filter on
	Schema *string
	// Optional json path to further filter on returned list
	AspectPath *string
	// Time at which this list was valid
	AtTime *string
	// Navigation links
	Links *NavT
}

// ListPayload is the payload type of the metadata service list method.
type ListPayload struct {
	// Entity for which to request metadata
	EntityID *string
	// Schema prefix using '%' as wildcard indicator
	Schema *string
	// To learn more about the supported format, see
	// https://www.postgresql.org/docs/current/datatype-json.html#DATATYPE-JSONPATH
	AspectPath *string
	// Return metadata which where valid at that time [now]
	AtTime *string
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
	OrderBy string
	// When set order result in descending order. Ascending order is the default.
	OrderDesc *bool
	// The content of '$page' is returned in the 'links' part of a previous query
	// and
	// will when set, ALL other parameters, except for 'limit' are ignored.
	Page *string
	// JWT used for authentication
	JWT string
}

type MetadataListItemRT struct {
	// Record ID
	RecordID *string
	// Entity ID
	Entity *string
	// Schema ID
	Schema *string
	// Attached metadata aspect
	Aspect interface{}
	// If aspectPath was defined, this is what matched the query
	AspectContext *string
}

// MetadataRecordRT is the result type of the metadata service read method.
type MetadataRecordRT struct {
	// Record ID
	RecordID *string
	// Entity ID
	Entity *string
	// Schema ID
	Schema *string
	// Attached metadata aspect
	Aspect interface{}
	// Time this record was asserted
	ValidFrom *string
	// Time this record was revoked
	ValidTo *string
	// Entity asserting this metadata record at 'valid-from'
	Asserter *string
	// Entity revoking this record at 'valid-to'
	Revoker *string
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

// ReadPayload is the payload type of the metadata service read method.
type ReadPayload struct {
	// ID of metadata to show
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

// RevokePayload is the payload type of the metadata service revoke method.
type RevokePayload struct {
	// Record ID to restract
	ID *string
	// JWT used for authentication
	JWT string
}

// ServiceNotAvailable is the type returned when the service necessary to
// fulfill the request is currently not available.
type ServiceNotAvailableT struct {
}

// Unauthorized access to resource
type UnauthorizedT struct {
}

// UpdateOnePayload is the payload type of the metadata service update_one
// method.
type UpdateOnePayload struct {
	// Record ID to update
	ID *string
	// Entity to which attach metadata
	EntityID string
	// Schema of metadata
	Schema string
	// Aspect content
	Aspect interface{}
	// Content-Type header, MUST be of application/json.
	ContentType *string
	// Policy guiding visibility and actions performed
	PolicyID *string
	// JWT used for authentication
	JWT string
}

// UpdateRecordPayload is the payload type of the metadata service
// update_record method.
type UpdateRecordPayload struct {
	// Record ID to update
	ID string
	// Entity to which attach metadata
	EntityID *string
	// Schema of metadata
	Schema *string
	// Aspect content
	Aspect interface{}
	// Content-Type header, MUST be of application/json.
	ContentType *string
	// Policy guiding visibility and actions performed
	PolicyID *string
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

// NewMetadataRecordRT initializes result type MetadataRecordRT from viewed
// result type MetadataRecordRT.
func NewMetadataRecordRT(vres *metadataviews.MetadataRecordRT) *MetadataRecordRT {
	return newMetadataRecordRT(vres.Projected)
}

// NewViewedMetadataRecordRT initializes viewed result type MetadataRecordRT
// from result type MetadataRecordRT using the given view.
func NewViewedMetadataRecordRT(res *MetadataRecordRT, view string) *metadataviews.MetadataRecordRT {
	p := newMetadataRecordRTView(res)
	return &metadataviews.MetadataRecordRT{Projected: p, View: "default"}
}

// NewListMetaRT initializes result type ListMetaRT from viewed result type
// ListMetaRT.
func NewListMetaRT(vres *metadataviews.ListMetaRT) *ListMetaRT {
	return newListMetaRT(vres.Projected)
}

// NewViewedListMetaRT initializes viewed result type ListMetaRT from result
// type ListMetaRT using the given view.
func NewViewedListMetaRT(res *ListMetaRT, view string) *metadataviews.ListMetaRT {
	p := newListMetaRTView(res)
	return &metadataviews.ListMetaRT{Projected: p, View: "default"}
}

// NewAddMetaRT initializes result type AddMetaRT from viewed result type
// AddMetaRT.
func NewAddMetaRT(vres *metadataviews.AddMetaRT) *AddMetaRT {
	return newAddMetaRT(vres.Projected)
}

// NewViewedAddMetaRT initializes viewed result type AddMetaRT from result type
// AddMetaRT using the given view.
func NewViewedAddMetaRT(res *AddMetaRT, view string) *metadataviews.AddMetaRT {
	p := newAddMetaRTView(res)
	return &metadataviews.AddMetaRT{Projected: p, View: "default"}
}

// newMetadataRecordRT converts projected type MetadataRecordRT to service type
// MetadataRecordRT.
func newMetadataRecordRT(vres *metadataviews.MetadataRecordRTView) *MetadataRecordRT {
	res := &MetadataRecordRT{
		RecordID:  vres.RecordID,
		Entity:    vres.Entity,
		Schema:    vres.Schema,
		Aspect:    vres.Aspect,
		ValidFrom: vres.ValidFrom,
		ValidTo:   vres.ValidTo,
		Asserter:  vres.Asserter,
		Revoker:   vres.Revoker,
	}
	return res
}

// newMetadataRecordRTView projects result type MetadataRecordRT to projected
// type MetadataRecordRTView using the "default" view.
func newMetadataRecordRTView(res *MetadataRecordRT) *metadataviews.MetadataRecordRTView {
	vres := &metadataviews.MetadataRecordRTView{
		RecordID:  res.RecordID,
		Entity:    res.Entity,
		Schema:    res.Schema,
		Aspect:    res.Aspect,
		ValidFrom: res.ValidFrom,
		ValidTo:   res.ValidTo,
		Asserter:  res.Asserter,
		Revoker:   res.Revoker,
	}
	return vres
}

// newListMetaRT converts projected type ListMetaRT to service type ListMetaRT.
func newListMetaRT(vres *metadataviews.ListMetaRTView) *ListMetaRT {
	res := &ListMetaRT{
		EntityID:   vres.EntityID,
		Schema:     vres.Schema,
		AspectPath: vres.AspectPath,
		AtTime:     vres.AtTime,
	}
	if vres.Records != nil {
		res.Records = make([]*MetadataListItemRT, len(vres.Records))
		for i, val := range vres.Records {
			res.Records[i] = transformMetadataviewsMetadataListItemRTViewToMetadataListItemRT(val)
		}
	}
	if vres.Links != nil {
		res.Links = transformMetadataviewsNavTViewToNavT(vres.Links)
	}
	return res
}

// newListMetaRTView projects result type ListMetaRT to projected type
// ListMetaRTView using the "default" view.
func newListMetaRTView(res *ListMetaRT) *metadataviews.ListMetaRTView {
	vres := &metadataviews.ListMetaRTView{
		EntityID:   res.EntityID,
		Schema:     res.Schema,
		AspectPath: res.AspectPath,
		AtTime:     res.AtTime,
	}
	if res.Records != nil {
		vres.Records = make([]*metadataviews.MetadataListItemRTView, len(res.Records))
		for i, val := range res.Records {
			vres.Records[i] = transformMetadataListItemRTToMetadataviewsMetadataListItemRTView(val)
		}
	}
	if res.Links != nil {
		vres.Links = transformNavTToMetadataviewsNavTView(res.Links)
	}
	return vres
}

// newMetadataListItemRT converts projected type MetadataListItemRT to service
// type MetadataListItemRT.
func newMetadataListItemRT(vres *metadataviews.MetadataListItemRTView) *MetadataListItemRT {
	res := &MetadataListItemRT{
		RecordID:      vres.RecordID,
		Entity:        vres.Entity,
		Schema:        vres.Schema,
		Aspect:        vres.Aspect,
		AspectContext: vres.AspectContext,
	}
	return res
}

// newMetadataListItemRTView projects result type MetadataListItemRT to
// projected type MetadataListItemRTView using the "default" view.
func newMetadataListItemRTView(res *MetadataListItemRT) *metadataviews.MetadataListItemRTView {
	vres := &metadataviews.MetadataListItemRTView{
		RecordID:      res.RecordID,
		Entity:        res.Entity,
		Schema:        res.Schema,
		Aspect:        res.Aspect,
		AspectContext: res.AspectContext,
	}
	return vres
}

// newAddMetaRT converts projected type AddMetaRT to service type AddMetaRT.
func newAddMetaRT(vres *metadataviews.AddMetaRTView) *AddMetaRT {
	res := &AddMetaRT{}
	if vres.RecordID != nil {
		res.RecordID = *vres.RecordID
	}
	return res
}

// newAddMetaRTView projects result type AddMetaRT to projected type
// AddMetaRTView using the "default" view.
func newAddMetaRTView(res *AddMetaRT) *metadataviews.AddMetaRTView {
	vres := &metadataviews.AddMetaRTView{
		RecordID: &res.RecordID,
	}
	return vres
}

// transformMetadataviewsMetadataListItemRTViewToMetadataListItemRT builds a
// value of type *MetadataListItemRT from a value of type
// *metadataviews.MetadataListItemRTView.
func transformMetadataviewsMetadataListItemRTViewToMetadataListItemRT(v *metadataviews.MetadataListItemRTView) *MetadataListItemRT {
	if v == nil {
		return nil
	}
	res := &MetadataListItemRT{
		RecordID:      v.RecordID,
		Entity:        v.Entity,
		Schema:        v.Schema,
		Aspect:        v.Aspect,
		AspectContext: v.AspectContext,
	}

	return res
}

// transformMetadataviewsNavTViewToNavT builds a value of type *NavT from a
// value of type *metadataviews.NavTView.
func transformMetadataviewsNavTViewToNavT(v *metadataviews.NavTView) *NavT {
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

// transformMetadataListItemRTToMetadataviewsMetadataListItemRTView builds a
// value of type *metadataviews.MetadataListItemRTView from a value of type
// *MetadataListItemRT.
func transformMetadataListItemRTToMetadataviewsMetadataListItemRTView(v *MetadataListItemRT) *metadataviews.MetadataListItemRTView {
	res := &metadataviews.MetadataListItemRTView{
		RecordID:      v.RecordID,
		Entity:        v.Entity,
		Schema:        v.Schema,
		Aspect:        v.Aspect,
		AspectContext: v.AspectContext,
	}

	return res
}

// transformNavTToMetadataviewsNavTView builds a value of type
// *metadataviews.NavTView from a value of type *NavT.
func transformNavTToMetadataviewsNavTView(v *NavT) *metadataviews.NavTView {
	res := &metadataviews.NavTView{
		Self:  v.Self,
		First: v.First,
		Next:  v.Next,
	}

	return res
}
