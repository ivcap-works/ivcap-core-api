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

package client

import (
	project "github.com/ivcap-works/ivcap-core-api/gen/project"
	projectviews "github.com/ivcap-works/ivcap-core-api/gen/project/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateProjectRequestBody is the type of the "project" service
// "CreateProject" endpoint HTTP request body.
type CreateProjectRequestBody struct {
	// Project name
	Name string `form:"name" json:"name" xml:"name"`
	// URN of the billing account
	AccountUrn *string `form:"account_urn,omitempty" json:"account_urn,omitempty" xml:"account_urn,omitempty"`
	// URN of the parent project
	ParentProjectUrn *string `form:"parent_project_urn,omitempty" json:"parent_project_urn,omitempty" xml:"parent_project_urn,omitempty"`
	// Additional Metadata
	Properties *ProjectPropertiesRequestBodyRequestBody `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
}

// UpdateMembershipRequestBody is the type of the "project" service
// "UpdateMembership" endpoint HTTP request body.
type UpdateMembershipRequestBody struct {
	// Role
	Role string `form:"role" json:"role" xml:"role"`
}

// SetDefaultProjectRequestBody is the type of the "project" service
// "SetDefaultProject" endpoint HTTP request body.
type SetDefaultProjectRequestBody struct {
	// Project URN
	ProjectUrn string `form:"project_urn" json:"project_urn" xml:"project_urn"`
	// User URN
	UserUrn *string `form:"user_urn,omitempty" json:"user_urn,omitempty" xml:"user_urn,omitempty"`
}

// SetProjectAccountRequestBody is the type of the "project" service
// "SetProjectAccount" endpoint HTTP request body.
type SetProjectAccountRequestBody struct {
	// Account URN
	AccountUrn string `form:"account_urn" json:"account_urn" xml:"account_urn"`
}

// ListResponseBody is the type of the "project" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Projects
	Projects ProjectListItemCollectionResponseBody `form:"projects,omitempty" json:"projects,omitempty" xml:"projects,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
	// A pagination token to retrieve the next set of results. Empty if there are
	// no more results
	Page *string `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
}

// CreateProjectResponseBody is the type of the "project" service
// "CreateProject" endpoint HTTP response body.
type CreateProjectResponseBody struct {
	// Project status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime project was created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// DateTime project last modified
	ModifiedAt *string `form:"modified_at,omitempty" json:"modified_at,omitempty" xml:"modified_at,omitempty"`
	// Account URN
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Parent Project URN
	Parent *string `form:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// Additional Metadata
	Properties *ProjectPropertiesResponseBody `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// Project URN
	Urn *string `form:"urn,omitempty" json:"urn,omitempty" xml:"urn,omitempty"`
	// Project name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ReadResponseBody is the type of the "project" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// Project status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime project was created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// DateTime project last modified
	ModifiedAt *string `form:"modified_at,omitempty" json:"modified_at,omitempty" xml:"modified_at,omitempty"`
	// Account URN
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Parent Project URN
	Parent *string `form:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// Additional Metadata
	Properties *ProjectPropertiesResponseBody `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// Project URN
	Urn *string `form:"urn,omitempty" json:"urn,omitempty" xml:"urn,omitempty"`
	// Project name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ListProjectMembersResponseBody is the type of the "project" service
// "ListProjectMembers" endpoint HTTP response body.
type ListProjectMembersResponseBody struct {
	// Members
	Members []*UserListItemResponseBody `form:"members,omitempty" json:"members,omitempty" xml:"members,omitempty"`
	// A pagination token to retrieve the next set of results. Empty if there are
	// no more results
	Page *string `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
}

// DefaultProjectResponseBody is the type of the "project" service
// "DefaultProject" endpoint HTTP response body.
type DefaultProjectResponseBody struct {
	// Project status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// DateTime project was created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// DateTime project last modified
	ModifiedAt *string `form:"modified_at,omitempty" json:"modified_at,omitempty" xml:"modified_at,omitempty"`
	// Account URN
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Parent Project URN
	Parent *string `form:"parent,omitempty" json:"parent,omitempty" xml:"parent,omitempty"`
	// Additional Metadata
	Properties *ProjectPropertiesResponseBody `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// Project URN
	Urn *string `form:"urn,omitempty" json:"urn,omitempty" xml:"urn,omitempty"`
	// Project name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ProjectAccountResponseBody is the type of the "project" service
// "ProjectAccount" endpoint HTTP response body.
type ProjectAccountResponseBody struct {
	// Account URN
	AccountUrn *string `form:"account_urn,omitempty" json:"account_urn,omitempty" xml:"account_urn,omitempty"`
}

// ListBadRequestResponseBody is the type of the "project" service "list"
// endpoint HTTP response body for the "bad-request" error.
type ListBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListInvalidParameterResponseBody is the type of the "project" service "list"
// endpoint HTTP response body for the "invalid-parameter" error.
type ListInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListInvalidScopesResponseBody is the type of the "project" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListNotImplementedResponseBody is the type of the "project" service "list"
// endpoint HTTP response body for the "not-implemented" error.
type ListNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateProjectBadRequestResponseBody is the type of the "project" service
// "CreateProject" endpoint HTTP response body for the "bad-request" error.
type CreateProjectBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateProjectInvalidParameterResponseBody is the type of the "project"
// service "CreateProject" endpoint HTTP response body for the
// "invalid-parameter" error.
type CreateProjectInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// CreateProjectInvalidScopesResponseBody is the type of the "project" service
// "CreateProject" endpoint HTTP response body for the "invalid-scopes" error.
type CreateProjectInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateProjectNotImplementedResponseBody is the type of the "project" service
// "CreateProject" endpoint HTTP response body for the "not-implemented" error.
type CreateProjectNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteBadRequestResponseBody is the type of the "project" service "delete"
// endpoint HTTP response body for the "bad-request" error.
type DeleteBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteInvalidScopesResponseBody is the type of the "project" service
// "delete" endpoint HTTP response body for the "invalid-scopes" error.
type DeleteInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteNotImplementedResponseBody is the type of the "project" service
// "delete" endpoint HTTP response body for the "not-implemented" error.
type DeleteNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "project" service "read"
// endpoint HTTP response body for the "bad-request" error.
type ReadBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadInvalidScopesResponseBody is the type of the "project" service "read"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReadInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotImplementedResponseBody is the type of the "project" service "read"
// endpoint HTTP response body for the "not-implemented" error.
type ReadNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ReadNotFoundResponseBody is the type of the "project" service "read"
// endpoint HTTP response body for the "not-found" error.
type ReadNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListProjectMembersBadRequestResponseBody is the type of the "project"
// service "ListProjectMembers" endpoint HTTP response body for the
// "bad-request" error.
type ListProjectMembersBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListProjectMembersInvalidParameterResponseBody is the type of the "project"
// service "ListProjectMembers" endpoint HTTP response body for the
// "invalid-parameter" error.
type ListProjectMembersInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ListProjectMembersInvalidScopesResponseBody is the type of the "project"
// service "ListProjectMembers" endpoint HTTP response body for the
// "invalid-scopes" error.
type ListProjectMembersInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListProjectMembersNotImplementedResponseBody is the type of the "project"
// service "ListProjectMembers" endpoint HTTP response body for the
// "not-implemented" error.
type ListProjectMembersNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListProjectMembersNotFoundResponseBody is the type of the "project" service
// "ListProjectMembers" endpoint HTTP response body for the "not-found" error.
type ListProjectMembersNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateMembershipBadRequestResponseBody is the type of the "project" service
// "UpdateMembership" endpoint HTTP response body for the "bad-request" error.
type UpdateMembershipBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateMembershipInvalidParameterResponseBody is the type of the "project"
// service "UpdateMembership" endpoint HTTP response body for the
// "invalid-parameter" error.
type UpdateMembershipInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// UpdateMembershipInvalidScopesResponseBody is the type of the "project"
// service "UpdateMembership" endpoint HTTP response body for the
// "invalid-scopes" error.
type UpdateMembershipInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateMembershipNotImplementedResponseBody is the type of the "project"
// service "UpdateMembership" endpoint HTTP response body for the
// "not-implemented" error.
type UpdateMembershipNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateMembershipNotFoundResponseBody is the type of the "project" service
// "UpdateMembership" endpoint HTTP response body for the "not-found" error.
type UpdateMembershipNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveMembershipBadRequestResponseBody is the type of the "project" service
// "RemoveMembership" endpoint HTTP response body for the "bad-request" error.
type RemoveMembershipBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveMembershipInvalidParameterResponseBody is the type of the "project"
// service "RemoveMembership" endpoint HTTP response body for the
// "invalid-parameter" error.
type RemoveMembershipInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// RemoveMembershipInvalidScopesResponseBody is the type of the "project"
// service "RemoveMembership" endpoint HTTP response body for the
// "invalid-scopes" error.
type RemoveMembershipInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveMembershipNotImplementedResponseBody is the type of the "project"
// service "RemoveMembership" endpoint HTTP response body for the
// "not-implemented" error.
type RemoveMembershipNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RemoveMembershipNotFoundResponseBody is the type of the "project" service
// "RemoveMembership" endpoint HTTP response body for the "not-found" error.
type RemoveMembershipNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DefaultProjectBadRequestResponseBody is the type of the "project" service
// "DefaultProject" endpoint HTTP response body for the "bad-request" error.
type DefaultProjectBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DefaultProjectInvalidParameterResponseBody is the type of the "project"
// service "DefaultProject" endpoint HTTP response body for the
// "invalid-parameter" error.
type DefaultProjectInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// DefaultProjectInvalidScopesResponseBody is the type of the "project" service
// "DefaultProject" endpoint HTTP response body for the "invalid-scopes" error.
type DefaultProjectInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DefaultProjectNotImplementedResponseBody is the type of the "project"
// service "DefaultProject" endpoint HTTP response body for the
// "not-implemented" error.
type DefaultProjectNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DefaultProjectNotFoundResponseBody is the type of the "project" service
// "DefaultProject" endpoint HTTP response body for the "not-found" error.
type DefaultProjectNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetDefaultProjectBadRequestResponseBody is the type of the "project" service
// "SetDefaultProject" endpoint HTTP response body for the "bad-request" error.
type SetDefaultProjectBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetDefaultProjectInvalidParameterResponseBody is the type of the "project"
// service "SetDefaultProject" endpoint HTTP response body for the
// "invalid-parameter" error.
type SetDefaultProjectInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// SetDefaultProjectInvalidScopesResponseBody is the type of the "project"
// service "SetDefaultProject" endpoint HTTP response body for the
// "invalid-scopes" error.
type SetDefaultProjectInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetDefaultProjectNotImplementedResponseBody is the type of the "project"
// service "SetDefaultProject" endpoint HTTP response body for the
// "not-implemented" error.
type SetDefaultProjectNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetDefaultProjectNotFoundResponseBody is the type of the "project" service
// "SetDefaultProject" endpoint HTTP response body for the "not-found" error.
type SetDefaultProjectNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ProjectAccountBadRequestResponseBody is the type of the "project" service
// "ProjectAccount" endpoint HTTP response body for the "bad-request" error.
type ProjectAccountBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ProjectAccountInvalidParameterResponseBody is the type of the "project"
// service "ProjectAccount" endpoint HTTP response body for the
// "invalid-parameter" error.
type ProjectAccountInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// ProjectAccountInvalidScopesResponseBody is the type of the "project" service
// "ProjectAccount" endpoint HTTP response body for the "invalid-scopes" error.
type ProjectAccountInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ProjectAccountNotImplementedResponseBody is the type of the "project"
// service "ProjectAccount" endpoint HTTP response body for the
// "not-implemented" error.
type ProjectAccountNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ProjectAccountNotFoundResponseBody is the type of the "project" service
// "ProjectAccount" endpoint HTTP response body for the "not-found" error.
type ProjectAccountNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetProjectAccountBadRequestResponseBody is the type of the "project" service
// "SetProjectAccount" endpoint HTTP response body for the "bad-request" error.
type SetProjectAccountBadRequestResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetProjectAccountInvalidParameterResponseBody is the type of the "project"
// service "SetProjectAccount" endpoint HTTP response body for the
// "invalid-parameter" error.
type SetProjectAccountInvalidParameterResponseBody struct {
	// message describing expected type or pattern.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// name of parameter.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// provided parameter value.
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// SetProjectAccountInvalidScopesResponseBody is the type of the "project"
// service "SetProjectAccount" endpoint HTTP response body for the
// "invalid-scopes" error.
type SetProjectAccountInvalidScopesResponseBody struct {
	// ID of involved resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetProjectAccountNotImplementedResponseBody is the type of the "project"
// service "SetProjectAccount" endpoint HTTP response body for the
// "not-implemented" error.
type SetProjectAccountNotImplementedResponseBody struct {
	// Information message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// SetProjectAccountNotFoundResponseBody is the type of the "project" service
// "SetProjectAccount" endpoint HTTP response body for the "not-found" error.
type SetProjectAccountNotFoundResponseBody struct {
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ProjectListItemCollectionResponseBody is used to define fields on response
// body types.
type ProjectListItemCollectionResponseBody []*ProjectListItemResponseBody

// ProjectListItemResponseBody is used to define fields on response body types.
type ProjectListItemResponseBody struct {
	// Project Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// User Role
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Project URN
	Urn *string `form:"urn,omitempty" json:"urn,omitempty" xml:"urn,omitempty"`
	// DateTime project was created
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// DateTime project last modified
	ModifiedAt *string `form:"modified_at,omitempty" json:"modified_at,omitempty" xml:"modified_at,omitempty"`
	// Time at which this list was valid
	AtTime *string `form:"at-time,omitempty" json:"at-time,omitempty" xml:"at-time,omitempty"`
}

// ProjectPropertiesRequestBodyRequestBody is used to define fields on request
// body types.
type ProjectPropertiesRequestBodyRequestBody struct {
	// String metadata for detailing the use of this project
	Details *string `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
}

// ProjectPropertiesResponseBody is used to define fields on response body
// types.
type ProjectPropertiesResponseBody struct {
	// String metadata for detailing the use of this project
	Details *string `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
}

// UserListItemResponseBody is used to define fields on response body types.
type UserListItemResponseBody struct {
	// User URN
	Urn *string `form:"urn,omitempty" json:"urn,omitempty" xml:"urn,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Role
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

// NewCreateProjectRequestBody builds the HTTP request body from the payload of
// the "CreateProject" endpoint of the "project" service.
func NewCreateProjectRequestBody(p *project.CreateProjectPayload) *CreateProjectRequestBody {
	body := &CreateProjectRequestBody{
		Name:             p.Project.Name,
		AccountUrn:       p.Project.AccountUrn,
		ParentProjectUrn: p.Project.ParentProjectUrn,
	}
	if p.Project.Properties != nil {
		body.Properties = marshalProjectProjectPropertiesToProjectPropertiesRequestBodyRequestBody(p.Project.Properties)
	}
	return body
}

// NewUpdateMembershipRequestBody builds the HTTP request body from the payload
// of the "UpdateMembership" endpoint of the "project" service.
func NewUpdateMembershipRequestBody(p *project.UpdateMembershipPayload) *UpdateMembershipRequestBody {
	body := &UpdateMembershipRequestBody{
		Role: p.Role,
	}
	return body
}

// NewSetDefaultProjectRequestBody builds the HTTP request body from the
// payload of the "SetDefaultProject" endpoint of the "project" service.
func NewSetDefaultProjectRequestBody(p *project.SetDefaultProjectPayload) *SetDefaultProjectRequestBody {
	body := &SetDefaultProjectRequestBody{
		ProjectUrn: p.ProjectUrn,
		UserUrn:    p.UserUrn,
	}
	return body
}

// NewSetProjectAccountRequestBody builds the HTTP request body from the
// payload of the "SetProjectAccount" endpoint of the "project" service.
func NewSetProjectAccountRequestBody(p *project.SetProjectAccountPayload) *SetProjectAccountRequestBody {
	body := &SetProjectAccountRequestBody{
		AccountUrn: p.AccountUrn,
	}
	return body
}

// NewListProjectListRTOK builds a "project" service "list" endpoint result
// from a HTTP "OK" response.
func NewListProjectListRTOK(body *ListResponseBody) *projectviews.ProjectListRTView {
	v := &projectviews.ProjectListRTView{
		AtTime: body.AtTime,
		Page:   body.Page,
	}
	v.Projects = make([]*projectviews.ProjectListItemView, len(body.Projects))
	for i, val := range body.Projects {
		v.Projects[i] = unmarshalProjectListItemResponseBodyToProjectviewsProjectListItemView(val)
	}

	return v
}

// NewListBadRequest builds a project service list endpoint bad-request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListInvalidParameter builds a project service list endpoint
// invalid-parameter error.
func NewListInvalidParameter(body *ListInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListInvalidScopes builds a project service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body *ListInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListNotImplemented builds a project service list endpoint not-implemented
// error.
func NewListNotImplemented(body *ListNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListNotAvailable builds a project service list endpoint not-available
// error.
func NewListNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewListNotAuthorized builds a project service list endpoint not-authorized
// error.
func NewListNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewCreateProjectProjectStatusRTOK builds a "project" service "CreateProject"
// endpoint result from a HTTP "OK" response.
func NewCreateProjectProjectStatusRTOK(body *CreateProjectResponseBody) *projectviews.ProjectStatusRTView {
	v := &projectviews.ProjectStatusRTView{
		Status:     body.Status,
		CreatedAt:  body.CreatedAt,
		ModifiedAt: body.ModifiedAt,
		Account:    body.Account,
		Parent:     body.Parent,
		Urn:        body.Urn,
		Name:       body.Name,
	}
	if body.Properties != nil {
		v.Properties = unmarshalProjectPropertiesResponseBodyToProjectviewsProjectPropertiesView(body.Properties)
	}

	return v
}

// NewCreateProjectBadRequest builds a project service CreateProject endpoint
// bad-request error.
func NewCreateProjectBadRequest(body *CreateProjectBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewCreateProjectInvalidParameter builds a project service CreateProject
// endpoint invalid-parameter error.
func NewCreateProjectInvalidParameter(body *CreateProjectInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewCreateProjectInvalidScopes builds a project service CreateProject
// endpoint invalid-scopes error.
func NewCreateProjectInvalidScopes(body *CreateProjectInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewCreateProjectNotImplemented builds a project service CreateProject
// endpoint not-implemented error.
func NewCreateProjectNotImplemented(body *CreateProjectNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewCreateProjectNotAvailable builds a project service CreateProject endpoint
// not-available error.
func NewCreateProjectNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewCreateProjectNotAuthorized builds a project service CreateProject
// endpoint not-authorized error.
func NewCreateProjectNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewDeleteBadRequest builds a project service delete endpoint bad-request
// error.
func NewDeleteBadRequest(body *DeleteBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewDeleteInvalidScopes builds a project service delete endpoint
// invalid-scopes error.
func NewDeleteInvalidScopes(body *DeleteInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotImplemented builds a project service delete endpoint
// not-implemented error.
func NewDeleteNotImplemented(body *DeleteNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotAvailable builds a project service delete endpoint not-available
// error.
func NewDeleteNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewDeleteNotAuthorized builds a project service delete endpoint
// not-authorized error.
func NewDeleteNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewReadProjectStatusRTOK builds a "project" service "read" endpoint result
// from a HTTP "OK" response.
func NewReadProjectStatusRTOK(body *ReadResponseBody) *projectviews.ProjectStatusRTView {
	v := &projectviews.ProjectStatusRTView{
		Status:     body.Status,
		CreatedAt:  body.CreatedAt,
		ModifiedAt: body.ModifiedAt,
		Account:    body.Account,
		Parent:     body.Parent,
		Urn:        body.Urn,
		Name:       body.Name,
	}
	if body.Properties != nil {
		v.Properties = unmarshalProjectPropertiesResponseBodyToProjectviewsProjectPropertiesView(body.Properties)
	}

	return v
}

// NewReadBadRequest builds a project service read endpoint bad-request error.
func NewReadBadRequest(body *ReadBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewReadInvalidScopes builds a project service read endpoint invalid-scopes
// error.
func NewReadInvalidScopes(body *ReadInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotImplemented builds a project service read endpoint not-implemented
// error.
func NewReadNotImplemented(body *ReadNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewReadNotFound builds a project service read endpoint not-found error.
func NewReadNotFound(body *ReadNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewReadNotAvailable builds a project service read endpoint not-available
// error.
func NewReadNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewReadNotAuthorized builds a project service read endpoint not-authorized
// error.
func NewReadNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewListProjectMembersMembersListOK builds a "project" service
// "ListProjectMembers" endpoint result from a HTTP "OK" response.
func NewListProjectMembersMembersListOK(body *ListProjectMembersResponseBody) *project.MembersList {
	v := &project.MembersList{
		Page:   body.Page,
		AtTime: body.AtTime,
	}
	v.Members = make([]*project.UserListItem, len(body.Members))
	for i, val := range body.Members {
		v.Members[i] = unmarshalUserListItemResponseBodyToProjectUserListItem(val)
	}

	return v
}

// NewListProjectMembersBadRequest builds a project service ListProjectMembers
// endpoint bad-request error.
func NewListProjectMembersBadRequest(body *ListProjectMembersBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewListProjectMembersInvalidParameter builds a project service
// ListProjectMembers endpoint invalid-parameter error.
func NewListProjectMembersInvalidParameter(body *ListProjectMembersInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewListProjectMembersInvalidScopes builds a project service
// ListProjectMembers endpoint invalid-scopes error.
func NewListProjectMembersInvalidScopes(body *ListProjectMembersInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListProjectMembersNotImplemented builds a project service
// ListProjectMembers endpoint not-implemented error.
func NewListProjectMembersNotImplemented(body *ListProjectMembersNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewListProjectMembersNotFound builds a project service ListProjectMembers
// endpoint not-found error.
func NewListProjectMembersNotFound(body *ListProjectMembersNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewListProjectMembersNotAvailable builds a project service
// ListProjectMembers endpoint not-available error.
func NewListProjectMembersNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewListProjectMembersNotAuthorized builds a project service
// ListProjectMembers endpoint not-authorized error.
func NewListProjectMembersNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewUpdateMembershipBadRequest builds a project service UpdateMembership
// endpoint bad-request error.
func NewUpdateMembershipBadRequest(body *UpdateMembershipBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateMembershipInvalidParameter builds a project service
// UpdateMembership endpoint invalid-parameter error.
func NewUpdateMembershipInvalidParameter(body *UpdateMembershipInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewUpdateMembershipInvalidScopes builds a project service UpdateMembership
// endpoint invalid-scopes error.
func NewUpdateMembershipInvalidScopes(body *UpdateMembershipInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateMembershipNotImplemented builds a project service UpdateMembership
// endpoint not-implemented error.
func NewUpdateMembershipNotImplemented(body *UpdateMembershipNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewUpdateMembershipNotFound builds a project service UpdateMembership
// endpoint not-found error.
func NewUpdateMembershipNotFound(body *UpdateMembershipNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewUpdateMembershipNotAvailable builds a project service UpdateMembership
// endpoint not-available error.
func NewUpdateMembershipNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewUpdateMembershipNotAuthorized builds a project service UpdateMembership
// endpoint not-authorized error.
func NewUpdateMembershipNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewRemoveMembershipBadRequest builds a project service RemoveMembership
// endpoint bad-request error.
func NewRemoveMembershipBadRequest(body *RemoveMembershipBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewRemoveMembershipInvalidParameter builds a project service
// RemoveMembership endpoint invalid-parameter error.
func NewRemoveMembershipInvalidParameter(body *RemoveMembershipInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewRemoveMembershipInvalidScopes builds a project service RemoveMembership
// endpoint invalid-scopes error.
func NewRemoveMembershipInvalidScopes(body *RemoveMembershipInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewRemoveMembershipNotImplemented builds a project service RemoveMembership
// endpoint not-implemented error.
func NewRemoveMembershipNotImplemented(body *RemoveMembershipNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewRemoveMembershipNotFound builds a project service RemoveMembership
// endpoint not-found error.
func NewRemoveMembershipNotFound(body *RemoveMembershipNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewRemoveMembershipNotAvailable builds a project service RemoveMembership
// endpoint not-available error.
func NewRemoveMembershipNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewRemoveMembershipNotAuthorized builds a project service RemoveMembership
// endpoint not-authorized error.
func NewRemoveMembershipNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewDefaultProjectProjectStatusRTOK builds a "project" service
// "DefaultProject" endpoint result from a HTTP "OK" response.
func NewDefaultProjectProjectStatusRTOK(body *DefaultProjectResponseBody) *projectviews.ProjectStatusRTView {
	v := &projectviews.ProjectStatusRTView{
		Status:     body.Status,
		CreatedAt:  body.CreatedAt,
		ModifiedAt: body.ModifiedAt,
		Account:    body.Account,
		Parent:     body.Parent,
		Urn:        body.Urn,
		Name:       body.Name,
	}
	if body.Properties != nil {
		v.Properties = unmarshalProjectPropertiesResponseBodyToProjectviewsProjectPropertiesView(body.Properties)
	}

	return v
}

// NewDefaultProjectBadRequest builds a project service DefaultProject endpoint
// bad-request error.
func NewDefaultProjectBadRequest(body *DefaultProjectBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewDefaultProjectInvalidParameter builds a project service DefaultProject
// endpoint invalid-parameter error.
func NewDefaultProjectInvalidParameter(body *DefaultProjectInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewDefaultProjectInvalidScopes builds a project service DefaultProject
// endpoint invalid-scopes error.
func NewDefaultProjectInvalidScopes(body *DefaultProjectInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewDefaultProjectNotImplemented builds a project service DefaultProject
// endpoint not-implemented error.
func NewDefaultProjectNotImplemented(body *DefaultProjectNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewDefaultProjectNotFound builds a project service DefaultProject endpoint
// not-found error.
func NewDefaultProjectNotFound(body *DefaultProjectNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewDefaultProjectNotAvailable builds a project service DefaultProject
// endpoint not-available error.
func NewDefaultProjectNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewDefaultProjectNotAuthorized builds a project service DefaultProject
// endpoint not-authorized error.
func NewDefaultProjectNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewSetDefaultProjectBadRequest builds a project service SetDefaultProject
// endpoint bad-request error.
func NewSetDefaultProjectBadRequest(body *SetDefaultProjectBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewSetDefaultProjectInvalidParameter builds a project service
// SetDefaultProject endpoint invalid-parameter error.
func NewSetDefaultProjectInvalidParameter(body *SetDefaultProjectInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewSetDefaultProjectInvalidScopes builds a project service SetDefaultProject
// endpoint invalid-scopes error.
func NewSetDefaultProjectInvalidScopes(body *SetDefaultProjectInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSetDefaultProjectNotImplemented builds a project service
// SetDefaultProject endpoint not-implemented error.
func NewSetDefaultProjectNotImplemented(body *SetDefaultProjectNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewSetDefaultProjectNotFound builds a project service SetDefaultProject
// endpoint not-found error.
func NewSetDefaultProjectNotFound(body *SetDefaultProjectNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSetDefaultProjectNotAvailable builds a project service SetDefaultProject
// endpoint not-available error.
func NewSetDefaultProjectNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewSetDefaultProjectNotAuthorized builds a project service SetDefaultProject
// endpoint not-authorized error.
func NewSetDefaultProjectNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewProjectAccountAccountResultOK builds a "project" service "ProjectAccount"
// endpoint result from a HTTP "OK" response.
func NewProjectAccountAccountResultOK(body *ProjectAccountResponseBody) *project.AccountResult {
	v := &project.AccountResult{
		AccountUrn: *body.AccountUrn,
	}

	return v
}

// NewProjectAccountBadRequest builds a project service ProjectAccount endpoint
// bad-request error.
func NewProjectAccountBadRequest(body *ProjectAccountBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewProjectAccountInvalidParameter builds a project service ProjectAccount
// endpoint invalid-parameter error.
func NewProjectAccountInvalidParameter(body *ProjectAccountInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewProjectAccountInvalidScopes builds a project service ProjectAccount
// endpoint invalid-scopes error.
func NewProjectAccountInvalidScopes(body *ProjectAccountInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewProjectAccountNotImplemented builds a project service ProjectAccount
// endpoint not-implemented error.
func NewProjectAccountNotImplemented(body *ProjectAccountNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewProjectAccountNotFound builds a project service ProjectAccount endpoint
// not-found error.
func NewProjectAccountNotFound(body *ProjectAccountNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewProjectAccountNotAvailable builds a project service ProjectAccount
// endpoint not-available error.
func NewProjectAccountNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewProjectAccountNotAuthorized builds a project service ProjectAccount
// endpoint not-authorized error.
func NewProjectAccountNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// NewSetProjectAccountBadRequest builds a project service SetProjectAccount
// endpoint bad-request error.
func NewSetProjectAccountBadRequest(body *SetProjectAccountBadRequestResponseBody) *project.BadRequestT {
	v := &project.BadRequestT{
		Message: *body.Message,
	}

	return v
}

// NewSetProjectAccountInvalidParameter builds a project service
// SetProjectAccount endpoint invalid-parameter error.
func NewSetProjectAccountInvalidParameter(body *SetProjectAccountInvalidParameterResponseBody) *project.InvalidParameterT {
	v := &project.InvalidParameterT{
		Message: *body.Message,
		Name:    *body.Name,
		Value:   body.Value,
	}

	return v
}

// NewSetProjectAccountInvalidScopes builds a project service SetProjectAccount
// endpoint invalid-scopes error.
func NewSetProjectAccountInvalidScopes(body *SetProjectAccountInvalidScopesResponseBody) *project.InvalidScopesT {
	v := &project.InvalidScopesT{
		ID:      body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSetProjectAccountNotImplemented builds a project service
// SetProjectAccount endpoint not-implemented error.
func NewSetProjectAccountNotImplemented(body *SetProjectAccountNotImplementedResponseBody) *project.NotImplementedT {
	v := &project.NotImplementedT{
		Message: *body.Message,
	}

	return v
}

// NewSetProjectAccountNotFound builds a project service SetProjectAccount
// endpoint not-found error.
func NewSetProjectAccountNotFound(body *SetProjectAccountNotFoundResponseBody) *project.ResourceNotFoundT {
	v := &project.ResourceNotFoundT{
		ID:      *body.ID,
		Message: *body.Message,
	}

	return v
}

// NewSetProjectAccountNotAvailable builds a project service SetProjectAccount
// endpoint not-available error.
func NewSetProjectAccountNotAvailable() *project.ServiceNotAvailableT {
	v := &project.ServiceNotAvailableT{}

	return v
}

// NewSetProjectAccountNotAuthorized builds a project service SetProjectAccount
// endpoint not-authorized error.
func NewSetProjectAccountNotAuthorized() *project.UnauthorizedT {
	v := &project.UnauthorizedT{}

	return v
}

// ValidateListProjectMembersResponseBody runs the validations defined on
// ListProjectMembersResponseBody
func ValidateListProjectMembersResponseBody(body *ListProjectMembersResponseBody) (err error) {
	if body.Members == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("members", "body"))
	}
	if body.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.at-time", *body.AtTime, goa.FormatDateTime))
	}
	return
}

// ValidateProjectAccountResponseBody runs the validations defined on
// ProjectAccountResponseBody
func ValidateProjectAccountResponseBody(body *ProjectAccountResponseBody) (err error) {
	if body.AccountUrn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account_urn", "body"))
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

// ValidateCreateProjectBadRequestResponseBody runs the validations defined on
// CreateProject_bad-request_Response_Body
func ValidateCreateProjectBadRequestResponseBody(body *CreateProjectBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateProjectInvalidParameterResponseBody runs the validations
// defined on CreateProject_invalid-parameter_Response_Body
func ValidateCreateProjectInvalidParameterResponseBody(body *CreateProjectInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateProjectInvalidScopesResponseBody runs the validations defined
// on CreateProject_invalid-scopes_Response_Body
func ValidateCreateProjectInvalidScopesResponseBody(body *CreateProjectInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateCreateProjectNotImplementedResponseBody runs the validations defined
// on CreateProject_not-implemented_Response_Body
func ValidateCreateProjectNotImplementedResponseBody(body *CreateProjectNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
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

// ValidateListProjectMembersBadRequestResponseBody runs the validations
// defined on ListProjectMembers_bad-request_Response_Body
func ValidateListProjectMembersBadRequestResponseBody(body *ListProjectMembersBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProjectMembersInvalidParameterResponseBody runs the validations
// defined on ListProjectMembers_invalid-parameter_Response_Body
func ValidateListProjectMembersInvalidParameterResponseBody(body *ListProjectMembersInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProjectMembersInvalidScopesResponseBody runs the validations
// defined on ListProjectMembers_invalid-scopes_Response_Body
func ValidateListProjectMembersInvalidScopesResponseBody(body *ListProjectMembersInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateListProjectMembersNotImplementedResponseBody runs the validations
// defined on ListProjectMembers_not-implemented_Response_Body
func ValidateListProjectMembersNotImplementedResponseBody(body *ListProjectMembersNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProjectMembersNotFoundResponseBody runs the validations defined
// on ListProjectMembers_not-found_Response_Body
func ValidateListProjectMembersNotFoundResponseBody(body *ListProjectMembersNotFoundResponseBody) (err error) {
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

// ValidateUpdateMembershipBadRequestResponseBody runs the validations defined
// on UpdateMembership_bad-request_Response_Body
func ValidateUpdateMembershipBadRequestResponseBody(body *UpdateMembershipBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateMembershipInvalidParameterResponseBody runs the validations
// defined on UpdateMembership_invalid-parameter_Response_Body
func ValidateUpdateMembershipInvalidParameterResponseBody(body *UpdateMembershipInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateMembershipInvalidScopesResponseBody runs the validations
// defined on UpdateMembership_invalid-scopes_Response_Body
func ValidateUpdateMembershipInvalidScopesResponseBody(body *UpdateMembershipInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateMembershipNotImplementedResponseBody runs the validations
// defined on UpdateMembership_not-implemented_Response_Body
func ValidateUpdateMembershipNotImplementedResponseBody(body *UpdateMembershipNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateMembershipNotFoundResponseBody runs the validations defined on
// UpdateMembership_not-found_Response_Body
func ValidateUpdateMembershipNotFoundResponseBody(body *UpdateMembershipNotFoundResponseBody) (err error) {
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

// ValidateRemoveMembershipBadRequestResponseBody runs the validations defined
// on RemoveMembership_bad-request_Response_Body
func ValidateRemoveMembershipBadRequestResponseBody(body *RemoveMembershipBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRemoveMembershipInvalidParameterResponseBody runs the validations
// defined on RemoveMembership_invalid-parameter_Response_Body
func ValidateRemoveMembershipInvalidParameterResponseBody(body *RemoveMembershipInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRemoveMembershipInvalidScopesResponseBody runs the validations
// defined on RemoveMembership_invalid-scopes_Response_Body
func ValidateRemoveMembershipInvalidScopesResponseBody(body *RemoveMembershipInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateRemoveMembershipNotImplementedResponseBody runs the validations
// defined on RemoveMembership_not-implemented_Response_Body
func ValidateRemoveMembershipNotImplementedResponseBody(body *RemoveMembershipNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRemoveMembershipNotFoundResponseBody runs the validations defined on
// RemoveMembership_not-found_Response_Body
func ValidateRemoveMembershipNotFoundResponseBody(body *RemoveMembershipNotFoundResponseBody) (err error) {
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

// ValidateDefaultProjectBadRequestResponseBody runs the validations defined on
// DefaultProject_bad-request_Response_Body
func ValidateDefaultProjectBadRequestResponseBody(body *DefaultProjectBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDefaultProjectInvalidParameterResponseBody runs the validations
// defined on DefaultProject_invalid-parameter_Response_Body
func ValidateDefaultProjectInvalidParameterResponseBody(body *DefaultProjectInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDefaultProjectInvalidScopesResponseBody runs the validations defined
// on DefaultProject_invalid-scopes_Response_Body
func ValidateDefaultProjectInvalidScopesResponseBody(body *DefaultProjectInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateDefaultProjectNotImplementedResponseBody runs the validations
// defined on DefaultProject_not-implemented_Response_Body
func ValidateDefaultProjectNotImplementedResponseBody(body *DefaultProjectNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDefaultProjectNotFoundResponseBody runs the validations defined on
// DefaultProject_not-found_Response_Body
func ValidateDefaultProjectNotFoundResponseBody(body *DefaultProjectNotFoundResponseBody) (err error) {
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

// ValidateSetDefaultProjectBadRequestResponseBody runs the validations defined
// on SetDefaultProject_bad-request_Response_Body
func ValidateSetDefaultProjectBadRequestResponseBody(body *SetDefaultProjectBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetDefaultProjectInvalidParameterResponseBody runs the validations
// defined on SetDefaultProject_invalid-parameter_Response_Body
func ValidateSetDefaultProjectInvalidParameterResponseBody(body *SetDefaultProjectInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetDefaultProjectInvalidScopesResponseBody runs the validations
// defined on SetDefaultProject_invalid-scopes_Response_Body
func ValidateSetDefaultProjectInvalidScopesResponseBody(body *SetDefaultProjectInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateSetDefaultProjectNotImplementedResponseBody runs the validations
// defined on SetDefaultProject_not-implemented_Response_Body
func ValidateSetDefaultProjectNotImplementedResponseBody(body *SetDefaultProjectNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetDefaultProjectNotFoundResponseBody runs the validations defined
// on SetDefaultProject_not-found_Response_Body
func ValidateSetDefaultProjectNotFoundResponseBody(body *SetDefaultProjectNotFoundResponseBody) (err error) {
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

// ValidateProjectAccountBadRequestResponseBody runs the validations defined on
// ProjectAccount_bad-request_Response_Body
func ValidateProjectAccountBadRequestResponseBody(body *ProjectAccountBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProjectAccountInvalidParameterResponseBody runs the validations
// defined on ProjectAccount_invalid-parameter_Response_Body
func ValidateProjectAccountInvalidParameterResponseBody(body *ProjectAccountInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProjectAccountInvalidScopesResponseBody runs the validations defined
// on ProjectAccount_invalid-scopes_Response_Body
func ValidateProjectAccountInvalidScopesResponseBody(body *ProjectAccountInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateProjectAccountNotImplementedResponseBody runs the validations
// defined on ProjectAccount_not-implemented_Response_Body
func ValidateProjectAccountNotImplementedResponseBody(body *ProjectAccountNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProjectAccountNotFoundResponseBody runs the validations defined on
// ProjectAccount_not-found_Response_Body
func ValidateProjectAccountNotFoundResponseBody(body *ProjectAccountNotFoundResponseBody) (err error) {
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

// ValidateSetProjectAccountBadRequestResponseBody runs the validations defined
// on SetProjectAccount_bad-request_Response_Body
func ValidateSetProjectAccountBadRequestResponseBody(body *SetProjectAccountBadRequestResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetProjectAccountInvalidParameterResponseBody runs the validations
// defined on SetProjectAccount_invalid-parameter_Response_Body
func ValidateSetProjectAccountInvalidParameterResponseBody(body *SetProjectAccountInvalidParameterResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetProjectAccountInvalidScopesResponseBody runs the validations
// defined on SetProjectAccount_invalid-scopes_Response_Body
func ValidateSetProjectAccountInvalidScopesResponseBody(body *SetProjectAccountInvalidScopesResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateSetProjectAccountNotImplementedResponseBody runs the validations
// defined on SetProjectAccount_not-implemented_Response_Body
func ValidateSetProjectAccountNotImplementedResponseBody(body *SetProjectAccountNotImplementedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSetProjectAccountNotFoundResponseBody runs the validations defined
// on SetProjectAccount_not-found_Response_Body
func ValidateSetProjectAccountNotFoundResponseBody(body *SetProjectAccountNotFoundResponseBody) (err error) {
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

// ValidateProjectListItemCollectionResponseBody runs the validations defined
// on ProjectListItemCollectionResponseBody
func ValidateProjectListItemCollectionResponseBody(body ProjectListItemCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateProjectListItemResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateProjectListItemResponseBody runs the validations defined on
// ProjectListItemResponseBody
func ValidateProjectListItemResponseBody(body *ProjectListItemResponseBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created_at", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.ModifiedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.modified_at", *body.ModifiedAt, goa.FormatDateTime))
	}
	if body.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.at-time", *body.AtTime, goa.FormatDateTime))
	}
	return
}
