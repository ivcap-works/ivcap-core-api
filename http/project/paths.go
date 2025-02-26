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

package client

import (
	"fmt"
)

// ListProjectPath returns the URL path to the project service list HTTP endpoint.
func ListProjectPath() string {
	return "/1/project"
}

// CreateProjectProjectPath returns the URL path to the project service CreateProject HTTP endpoint.
func CreateProjectProjectPath() string {
	return "/1/project"
}

// DeleteProjectPath returns the URL path to the project service delete HTTP endpoint.
func DeleteProjectPath(id string) string {
	return fmt.Sprintf("/1/project/%v", id)
}

// ReadProjectPath returns the URL path to the project service read HTTP endpoint.
func ReadProjectPath(id string) string {
	return fmt.Sprintf("/1/project/%v", id)
}

// ListProjectMembersProjectPath returns the URL path to the project service ListProjectMembers HTTP endpoint.
func ListProjectMembersProjectPath(urn string) string {
	return fmt.Sprintf("/1/project/%v/members", urn)
}

// UpdateMembershipProjectPath returns the URL path to the project service UpdateMembership HTTP endpoint.
func UpdateMembershipProjectPath(projectUrn string, userUrn string) string {
	return fmt.Sprintf("/1/project/%v/memberships/%v", projectUrn, userUrn)
}

// RemoveMembershipProjectPath returns the URL path to the project service RemoveMembership HTTP endpoint.
func RemoveMembershipProjectPath(projectUrn string, userUrn string) string {
	return fmt.Sprintf("/1/project/%v/memberships/%v", projectUrn, userUrn)
}

// DefaultProjectProjectPath returns the URL path to the project service DefaultProject HTTP endpoint.
func DefaultProjectProjectPath() string {
	return "/1/project/default"
}

// SetDefaultProjectProjectPath returns the URL path to the project service SetDefaultProject HTTP endpoint.
func SetDefaultProjectProjectPath() string {
	return "/1/project/default"
}

// ProjectAccountProjectPath returns the URL path to the project service ProjectAccount HTTP endpoint.
func ProjectAccountProjectPath(projectUrn string) string {
	return fmt.Sprintf("/1/project/%v/account", projectUrn)
}

// SetProjectAccountProjectPath returns the URL path to the project service SetProjectAccount HTTP endpoint.
func SetProjectAccountProjectPath(projectUrn string) string {
	return fmt.Sprintf("/1/project/%v/account", projectUrn)
}
