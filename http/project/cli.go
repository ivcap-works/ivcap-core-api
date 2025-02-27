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
	"encoding/json"
	"fmt"
	"strconv"

	project "github.com/ivcap-works/ivcap-core-api/gen/project"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the project list endpoint from CLI
// flags.
func BuildListPayload(projectListLimit string, projectListPage string, projectListFilter string, projectListOrderBy string, projectListOrderDesc string, projectListAtTime string, projectListJWT string) (*project.ListPayload, error) {
	var err error
	var limit int
	{
		if projectListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(projectListLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 50 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 50, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var page *string
	{
		if projectListPage != "" {
			page = &projectListPage
		}
	}
	var filter *string
	{
		if projectListFilter != "" {
			filter = &projectListFilter
		}
	}
	var orderBy *string
	{
		if projectListOrderBy != "" {
			orderBy = &projectListOrderBy
		}
	}
	var orderDesc bool
	{
		if projectListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(projectListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if projectListAtTime != "" {
			atTime = &projectListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = projectListJWT
	}
	v := &project.ListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildCreateProjectPayload builds the payload for the project CreateProject
// endpoint from CLI flags.
func BuildCreateProjectPayload(projectCreateProjectBody string, projectCreateProjectJWT string) (*project.CreateProjectPayload, error) {
	var err error
	var body CreateProjectRequestBody
	{
		err = json.Unmarshal([]byte(projectCreateProjectBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"account_urn\": \"urn:ivcap:account:146d4ac9-244a-4aee-aa32-a28f4b91e60d\",\n      \"name\": \"My project name\",\n      \"parent_project_urn\": \"urn:ivcap:project:8a82775b-27d9-4635-b006-7ef5553656d1\",\n      \"properties\": {\n         \"details\": \"Created for to investigate [objective]\"\n      }\n   }'")
		}
		if body.AccountUrn != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.account_urn", *body.AccountUrn, goa.FormatURI))
		}
		if body.ParentProjectUrn != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.parent_project_urn", *body.ParentProjectUrn, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = projectCreateProjectJWT
	}
	v := &project.ProjectCreateRequest{
		Name:             body.Name,
		AccountUrn:       body.AccountUrn,
		ParentProjectUrn: body.ParentProjectUrn,
	}
	if body.Properties != nil {
		v.Properties = marshalProjectPropertiesRequestBodyRequestBodyToProjectProjectProperties(body.Properties)
	}
	res := &project.CreateProjectPayload{
		Project: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildDeletePayload builds the payload for the project delete endpoint from
// CLI flags.
func BuildDeletePayload(projectDeleteID string, projectDeleteJWT string) (*project.DeletePayload, error) {
	var id string
	{
		id = projectDeleteID
	}
	var jwt string
	{
		jwt = projectDeleteJWT
	}
	v := &project.DeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildReadPayload builds the payload for the project read endpoint from CLI
// flags.
func BuildReadPayload(projectReadID string, projectReadJWT string) (*project.ReadPayload, error) {
	var id string
	{
		id = projectReadID
	}
	var jwt string
	{
		jwt = projectReadJWT
	}
	v := &project.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildListProjectMembersPayload builds the payload for the project
// ListProjectMembers endpoint from CLI flags.
func BuildListProjectMembersPayload(projectListProjectMembersUrn string, projectListProjectMembersRole string, projectListProjectMembersLimit string, projectListProjectMembersPage string, projectListProjectMembersJWT string) (*project.ListProjectMembersPayload, error) {
	var err error
	var urn string
	{
		urn = projectListProjectMembersUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("urn", urn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var role *string
	{
		if projectListProjectMembersRole != "" {
			role = &projectListProjectMembersRole
		}
	}
	var limit int
	{
		if projectListProjectMembersLimit != "" {
			var v int64
			v, err = strconv.ParseInt(projectListProjectMembersLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 50 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 50, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var page *string
	{
		if projectListProjectMembersPage != "" {
			page = &projectListProjectMembersPage
		}
	}
	var jwt string
	{
		jwt = projectListProjectMembersJWT
	}
	v := &project.ListProjectMembersPayload{}
	v.Urn = urn
	v.Role = role
	v.Limit = limit
	v.Page = page
	v.JWT = jwt

	return v, nil
}

// BuildUpdateMembershipPayload builds the payload for the project
// UpdateMembership endpoint from CLI flags.
func BuildUpdateMembershipPayload(projectUpdateMembershipBody string, projectUpdateMembershipProjectUrn string, projectUpdateMembershipUserUrn string, projectUpdateMembershipJWT string) (*project.UpdateMembershipPayload, error) {
	var err error
	var body UpdateMembershipRequestBody
	{
		err = json.Unmarshal([]byte(projectUpdateMembershipBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"role\": \"owner\"\n   }'")
		}
	}
	var projectUrn string
	{
		projectUrn = projectUpdateMembershipProjectUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("project_urn", projectUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var userUrn string
	{
		userUrn = projectUpdateMembershipUserUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("user_urn", userUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = projectUpdateMembershipJWT
	}
	v := &project.UpdateMembershipPayload{
		Role: body.Role,
	}
	v.ProjectUrn = projectUrn
	v.UserUrn = userUrn
	v.JWT = jwt

	return v, nil
}

// BuildRemoveMembershipPayload builds the payload for the project
// RemoveMembership endpoint from CLI flags.
func BuildRemoveMembershipPayload(projectRemoveMembershipProjectUrn string, projectRemoveMembershipUserUrn string, projectRemoveMembershipJWT string) (*project.RemoveMembershipPayload, error) {
	var err error
	var projectUrn string
	{
		projectUrn = projectRemoveMembershipProjectUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("project_urn", projectUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var userUrn string
	{
		userUrn = projectRemoveMembershipUserUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("user_urn", userUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = projectRemoveMembershipJWT
	}
	v := &project.RemoveMembershipPayload{}
	v.ProjectUrn = projectUrn
	v.UserUrn = userUrn
	v.JWT = jwt

	return v, nil
}

// BuildDefaultProjectPayload builds the payload for the project DefaultProject
// endpoint from CLI flags.
func BuildDefaultProjectPayload(projectDefaultProjectJWT string) (*project.DefaultProjectPayload, error) {
	var jwt string
	{
		jwt = projectDefaultProjectJWT
	}
	v := &project.DefaultProjectPayload{}
	v.JWT = jwt

	return v, nil
}

// BuildSetDefaultProjectPayload builds the payload for the project
// SetDefaultProject endpoint from CLI flags.
func BuildSetDefaultProjectPayload(projectSetDefaultProjectBody string, projectSetDefaultProjectJWT string) (*project.SetDefaultProjectPayload, error) {
	var err error
	var body SetDefaultProjectRequestBody
	{
		err = json.Unmarshal([]byte(projectSetDefaultProjectBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"project_urn\": \"urn:ivcap:project:59c76bc8-721b-409d-8a32-6d560680e89f\",\n      \"user_urn\": \"urn:ivcap:user:0b755f67-4d03-4d82-b208-4d6a0ae16468\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.project_urn", body.ProjectUrn, goa.FormatURI))
		if body.UserUrn != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.user_urn", *body.UserUrn, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = projectSetDefaultProjectJWT
	}
	v := &project.SetDefaultProjectPayload{
		ProjectUrn: body.ProjectUrn,
		UserUrn:    body.UserUrn,
	}
	v.JWT = jwt

	return v, nil
}

// BuildProjectAccountPayload builds the payload for the project ProjectAccount
// endpoint from CLI flags.
func BuildProjectAccountPayload(projectProjectAccountProjectUrn string, projectProjectAccountJWT string) (*project.ProjectAccountPayload, error) {
	var err error
	var projectUrn string
	{
		projectUrn = projectProjectAccountProjectUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("project_urn", projectUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = projectProjectAccountJWT
	}
	v := &project.ProjectAccountPayload{}
	v.ProjectUrn = projectUrn
	v.JWT = jwt

	return v, nil
}

// BuildSetProjectAccountPayload builds the payload for the project
// SetProjectAccount endpoint from CLI flags.
func BuildSetProjectAccountPayload(projectSetProjectAccountBody string, projectSetProjectAccountProjectUrn string, projectSetProjectAccountJWT string) (*project.SetProjectAccountPayload, error) {
	var err error
	var body SetProjectAccountRequestBody
	{
		err = json.Unmarshal([]byte(projectSetProjectAccountBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"account_urn\": \"urn:ivcap:account:146d4ac9-244a-4aee-aa32-a28f4b91e60d\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account_urn", body.AccountUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var projectUrn string
	{
		projectUrn = projectSetProjectAccountProjectUrn
		err = goa.MergeErrors(err, goa.ValidateFormat("project_urn", projectUrn, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = projectSetProjectAccountJWT
	}
	v := &project.SetProjectAccountPayload{
		AccountUrn: body.AccountUrn,
	}
	v.ProjectUrn = projectUrn
	v.JWT = jwt

	return v, nil
}
