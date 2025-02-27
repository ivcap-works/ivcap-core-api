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

package project

import (
	"context"

	projectviews "github.com/ivcap-works/ivcap-core-api/gen/project/views"
	"goa.design/goa/v3/security"
)

// Create, Join, Manage Projects
type Service interface {
	// list project
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	List(context.Context, *ListPayload) (res *ProjectListRT, view string, err error)
	// Create a new project and return its status.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	CreateProject(context.Context, *CreateProjectPayload) (res *ProjectStatusRT, view string, err error)
	// Delete an existing project.
	Delete(context.Context, *DeletePayload) (err error)
	// Show project by ID
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Read(context.Context, *ReadPayload) (res *ProjectStatusRT, view string, err error)
	// Lists the current members of a project.
	ListProjectMembers(context.Context, *ListProjectMembersPayload) (res *MembersList, err error)
	// Adds or Updates the roles of a user in a project.
	UpdateMembership(context.Context, *UpdateMembershipPayload) (err error)
	// Remove a user from a project.
	RemoveMembership(context.Context, *RemoveMembershipPayload) (err error)
	// Retrieves the user's current default project
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	DefaultProject(context.Context, *DefaultProjectPayload) (res *ProjectStatusRT, view string, err error)
	// Sets the default project of a user.
	SetDefaultProject(context.Context, *SetDefaultProjectPayload) (err error)
	// Retrieves the project's billing account urn
	ProjectAccount(context.Context, *ProjectAccountPayload) (res *AccountResult, err error)
	// Sets the billing account of a project
	SetProjectAccount(context.Context, *SetProjectAccountPayload) (err error)
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
const ServiceName = "project"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [11]string{"list", "CreateProject", "delete", "read", "ListProjectMembers", "UpdateMembership", "RemoveMembership", "DefaultProject", "SetDefaultProject", "ProjectAccount", "SetProjectAccount"}

// AccountResult is the result type of the project service ProjectAccount
// method.
type AccountResult struct {
	// Account URN
	AccountUrn string
}

// Something wasn't right with this request
type BadRequestT struct {
	// Information message
	Message string
}

// CreateProjectPayload is the payload type of the project service
// CreateProject method.
type CreateProjectPayload struct {
	// New project description
	Project *ProjectCreateRequest
	// JWT used for authentication
	JWT string
}

// DefaultProjectPayload is the payload type of the project service
// DefaultProject method.
type DefaultProjectPayload struct {
	// JWT used for authentication
	JWT string
}

// DeletePayload is the payload type of the project service delete method.
type DeletePayload struct {
	// ID of project to update
	ID string
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

// ListPayload is the payload type of the project service list method.
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

// ListProjectMembersPayload is the payload type of the project service
// ListProjectMembers method.
type ListProjectMembersPayload struct {
	// Project URN
	Urn string
	// Role
	Role *string
	// The 'limit' query option sets the maximum number of items
	// to be included in the result.
	Limit int
	// A pagination token to retrieve the next set of results. Empty if there are
	// no more results
	Page *string
	// JWT used for authentication
	JWT string
}

// MembersList is the result type of the project service ListProjectMembers
// method.
type MembersList struct {
	// Members
	Members []*UserListItem
	// A pagination token to retrieve the next set of results. Empty if there are
	// no more results
	Page *string
	// Time at which this list was valid
	AtTime *string
}

// Method is not yet implemented.
type NotImplementedT struct {
	// Information message
	Message string
}

// ProjectAccountPayload is the payload type of the project service
// ProjectAccount method.
type ProjectAccountPayload struct {
	// Project URN
	ProjectUrn string
	// JWT used for authentication
	JWT string
}

type ProjectCreateRequest struct {
	// Project name
	Name string
	// URN of the billing account
	AccountUrn *string
	// URN of the parent project
	ParentProjectUrn *string
	// Additional Metadata
	Properties *ProjectProperties
}

type ProjectListItem struct {
	// Project Name
	Name *string
	// User Role
	Role *string
	// Project URN
	Urn *string
	// DateTime project was created
	CreatedAt *string
	// DateTime project last modified
	ModifiedAt *string
	// Time at which this list was valid
	AtTime *string
}

type ProjectListItemCollection []*ProjectListItem

// ProjectListRT is the result type of the project service list method.
type ProjectListRT struct {
	// Projects
	Projects ProjectListItemCollection
	// Time at which this list was valid
	AtTime *string
	// A pagination token to retrieve the next set of results. Empty if there are
	// no more results
	Page *string
}

type ProjectProperties struct {
	// String metadata for detailing the use of this project
	Details *string
}

// ProjectStatusRT is the result type of the project service CreateProject
// method.
type ProjectStatusRT struct {
	// Project status
	Status *string
	// DateTime project was created
	CreatedAt *string
	// DateTime project last modified
	ModifiedAt *string
	// Account URN
	Account *string
	// Parent Project URN
	Parent *string
	// Additional Metadata
	Properties *ProjectProperties
	// Project URN
	Urn string
	// Project name
	Name *string
}

// ReadPayload is the payload type of the project service read method.
type ReadPayload struct {
	// ID of project to show
	ID string
	// JWT used for authentication
	JWT string
}

// RemoveMembershipPayload is the payload type of the project service
// RemoveMembership method.
type RemoveMembershipPayload struct {
	// Project URN
	ProjectUrn string
	// User URN
	UserUrn string
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

// Service necessary to fulfil the request is currently not available.
type ServiceNotAvailableT struct {
}

// SetDefaultProjectPayload is the payload type of the project service
// SetDefaultProject method.
type SetDefaultProjectPayload struct {
	// Project URN
	ProjectUrn string
	// User URN
	UserUrn *string
	// JWT used for authentication
	JWT string
}

// SetProjectAccountPayload is the payload type of the project service
// SetProjectAccount method.
type SetProjectAccountPayload struct {
	// Project URN
	ProjectUrn string
	// Account URN
	AccountUrn string
	// JWT used for authentication
	JWT string
}

// Unauthorized access to resource
type UnauthorizedT struct {
}

// UpdateMembershipPayload is the payload type of the project service
// UpdateMembership method.
type UpdateMembershipPayload struct {
	// Project URN
	ProjectUrn string
	// User URN
	UserUrn string
	// Role
	Role string
	// JWT used for authentication
	JWT string
}

type UserListItem struct {
	// User URN
	Urn *string
	// Email
	Email *string
	// Role
	Role *string
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

// NewProjectListRT initializes result type ProjectListRT from viewed result
// type ProjectListRT.
func NewProjectListRT(vres *projectviews.ProjectListRT) *ProjectListRT {
	var res *ProjectListRT
	switch vres.View {
	case "default", "":
		res = newProjectListRT(vres.Projected)
	case "tiny":
		res = newProjectListRTTiny(vres.Projected)
	}
	return res
}

// NewViewedProjectListRT initializes viewed result type ProjectListRT from
// result type ProjectListRT using the given view.
func NewViewedProjectListRT(res *ProjectListRT, view string) *projectviews.ProjectListRT {
	var vres *projectviews.ProjectListRT
	switch view {
	case "default", "":
		p := newProjectListRTView(res)
		vres = &projectviews.ProjectListRT{Projected: p, View: "default"}
	case "tiny":
		p := newProjectListRTViewTiny(res)
		vres = &projectviews.ProjectListRT{Projected: p, View: "tiny"}
	}
	return vres
}

// NewProjectStatusRT initializes result type ProjectStatusRT from viewed
// result type ProjectStatusRT.
func NewProjectStatusRT(vres *projectviews.ProjectStatusRT) *ProjectStatusRT {
	var res *ProjectStatusRT
	switch vres.View {
	case "default", "":
		res = newProjectStatusRT(vres.Projected)
	case "tiny":
		res = newProjectStatusRTTiny(vres.Projected)
	}
	return res
}

// NewViewedProjectStatusRT initializes viewed result type ProjectStatusRT from
// result type ProjectStatusRT using the given view.
func NewViewedProjectStatusRT(res *ProjectStatusRT, view string) *projectviews.ProjectStatusRT {
	var vres *projectviews.ProjectStatusRT
	switch view {
	case "default", "":
		p := newProjectStatusRTView(res)
		vres = &projectviews.ProjectStatusRT{Projected: p, View: "default"}
	case "tiny":
		p := newProjectStatusRTViewTiny(res)
		vres = &projectviews.ProjectStatusRT{Projected: p, View: "tiny"}
	}
	return vres
}

// newProjectListRT converts projected type ProjectListRT to service type
// ProjectListRT.
func newProjectListRT(vres *projectviews.ProjectListRTView) *ProjectListRT {
	res := &ProjectListRT{
		AtTime: vres.AtTime,
		Page:   vres.Page,
	}
	if vres.Projects != nil {
		res.Projects = newProjectListItemCollection(vres.Projects)
	}
	return res
}

// newProjectListRTTiny converts projected type ProjectListRT to service type
// ProjectListRT.
func newProjectListRTTiny(vres *projectviews.ProjectListRTView) *ProjectListRT {
	res := &ProjectListRT{
		AtTime: vres.AtTime,
		Page:   vres.Page,
	}
	if vres.Projects != nil {
		res.Projects = newProjectListItemCollectionTiny(vres.Projects)
	}
	return res
}

// newProjectListRTView projects result type ProjectListRT to projected type
// ProjectListRTView using the "default" view.
func newProjectListRTView(res *ProjectListRT) *projectviews.ProjectListRTView {
	vres := &projectviews.ProjectListRTView{
		AtTime: res.AtTime,
		Page:   res.Page,
	}
	if res.Projects != nil {
		vres.Projects = newProjectListItemCollectionView(res.Projects)
	}
	return vres
}

// newProjectListRTViewTiny projects result type ProjectListRT to projected
// type ProjectListRTView using the "tiny" view.
func newProjectListRTViewTiny(res *ProjectListRT) *projectviews.ProjectListRTView {
	vres := &projectviews.ProjectListRTView{
		AtTime: res.AtTime,
		Page:   res.Page,
	}
	if res.Projects != nil {
		vres.Projects = newProjectListItemCollectionViewTiny(res.Projects)
	}
	return vres
}

// newProjectListItemCollection converts projected type
// ProjectListItemCollection to service type ProjectListItemCollection.
func newProjectListItemCollection(vres projectviews.ProjectListItemCollectionView) ProjectListItemCollection {
	res := make(ProjectListItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newProjectListItem(n)
	}
	return res
}

// newProjectListItemCollectionTiny converts projected type
// ProjectListItemCollection to service type ProjectListItemCollection.
func newProjectListItemCollectionTiny(vres projectviews.ProjectListItemCollectionView) ProjectListItemCollection {
	res := make(ProjectListItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newProjectListItemTiny(n)
	}
	return res
}

// newProjectListItemCollectionView projects result type
// ProjectListItemCollection to projected type ProjectListItemCollectionView
// using the "default" view.
func newProjectListItemCollectionView(res ProjectListItemCollection) projectviews.ProjectListItemCollectionView {
	vres := make(projectviews.ProjectListItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newProjectListItemView(n)
	}
	return vres
}

// newProjectListItemCollectionViewTiny projects result type
// ProjectListItemCollection to projected type ProjectListItemCollectionView
// using the "tiny" view.
func newProjectListItemCollectionViewTiny(res ProjectListItemCollection) projectviews.ProjectListItemCollectionView {
	vres := make(projectviews.ProjectListItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newProjectListItemViewTiny(n)
	}
	return vres
}

// newProjectListItem converts projected type ProjectListItem to service type
// ProjectListItem.
func newProjectListItem(vres *projectviews.ProjectListItemView) *ProjectListItem {
	res := &ProjectListItem{
		Name:       vres.Name,
		Role:       vres.Role,
		Urn:        vres.Urn,
		CreatedAt:  vres.CreatedAt,
		ModifiedAt: vres.ModifiedAt,
		AtTime:     vres.AtTime,
	}
	return res
}

// newProjectListItemTiny converts projected type ProjectListItem to service
// type ProjectListItem.
func newProjectListItemTiny(vres *projectviews.ProjectListItemView) *ProjectListItem {
	res := &ProjectListItem{
		Urn: vres.Urn,
	}
	return res
}

// newProjectListItemView projects result type ProjectListItem to projected
// type ProjectListItemView using the "default" view.
func newProjectListItemView(res *ProjectListItem) *projectviews.ProjectListItemView {
	vres := &projectviews.ProjectListItemView{
		Name:       res.Name,
		Role:       res.Role,
		Urn:        res.Urn,
		CreatedAt:  res.CreatedAt,
		ModifiedAt: res.ModifiedAt,
		AtTime:     res.AtTime,
	}
	return vres
}

// newProjectListItemViewTiny projects result type ProjectListItem to projected
// type ProjectListItemView using the "tiny" view.
func newProjectListItemViewTiny(res *ProjectListItem) *projectviews.ProjectListItemView {
	vres := &projectviews.ProjectListItemView{
		Urn: res.Urn,
	}
	return vres
}

// newProjectStatusRT converts projected type ProjectStatusRT to service type
// ProjectStatusRT.
func newProjectStatusRT(vres *projectviews.ProjectStatusRTView) *ProjectStatusRT {
	res := &ProjectStatusRT{
		Name:       vres.Name,
		Account:    vres.Account,
		Parent:     vres.Parent,
		Status:     vres.Status,
		CreatedAt:  vres.CreatedAt,
		ModifiedAt: vres.ModifiedAt,
	}
	if vres.Urn != nil {
		res.Urn = *vres.Urn
	}
	if vres.Properties != nil {
		res.Properties = transformProjectviewsProjectPropertiesViewToProjectProperties(vres.Properties)
	}
	return res
}

// newProjectStatusRTTiny converts projected type ProjectStatusRT to service
// type ProjectStatusRT.
func newProjectStatusRTTiny(vres *projectviews.ProjectStatusRTView) *ProjectStatusRT {
	res := &ProjectStatusRT{
		Name:   vres.Name,
		Status: vres.Status,
	}
	return res
}

// newProjectStatusRTView projects result type ProjectStatusRT to projected
// type ProjectStatusRTView using the "default" view.
func newProjectStatusRTView(res *ProjectStatusRT) *projectviews.ProjectStatusRTView {
	vres := &projectviews.ProjectStatusRTView{
		Status:     res.Status,
		CreatedAt:  res.CreatedAt,
		ModifiedAt: res.ModifiedAt,
		Account:    res.Account,
		Parent:     res.Parent,
		Urn:        &res.Urn,
		Name:       res.Name,
	}
	if res.Properties != nil {
		vres.Properties = transformProjectPropertiesToProjectviewsProjectPropertiesView(res.Properties)
	}
	return vres
}

// newProjectStatusRTViewTiny projects result type ProjectStatusRT to projected
// type ProjectStatusRTView using the "tiny" view.
func newProjectStatusRTViewTiny(res *ProjectStatusRT) *projectviews.ProjectStatusRTView {
	vres := &projectviews.ProjectStatusRTView{
		Status: res.Status,
		Name:   res.Name,
	}
	return vres
}

// transformProjectviewsProjectPropertiesViewToProjectProperties builds a value
// of type *ProjectProperties from a value of type
// *projectviews.ProjectPropertiesView.
func transformProjectviewsProjectPropertiesViewToProjectProperties(v *projectviews.ProjectPropertiesView) *ProjectProperties {
	if v == nil {
		return nil
	}
	res := &ProjectProperties{
		Details: v.Details,
	}

	return res
}

// transformProjectPropertiesToProjectviewsProjectPropertiesView builds a value
// of type *projectviews.ProjectPropertiesView from a value of type
// *ProjectProperties.
func transformProjectPropertiesToProjectviewsProjectPropertiesView(v *ProjectProperties) *projectviews.ProjectPropertiesView {
	if v == nil {
		return nil
	}
	res := &projectviews.ProjectPropertiesView{
		Details: v.Details,
	}

	return res
}
