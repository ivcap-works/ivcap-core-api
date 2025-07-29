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
	"fmt"
)

// ServiceListServicePath returns the URL path to the service service service-list HTTP endpoint.
func ServiceListServicePath() string {
	return "/1/services2"
}

// ServiceCreateServicePath returns the URL path to the service service service-create HTTP endpoint.
func ServiceCreateServicePath() string {
	return "/1/services2"
}

// ServiceReadServicePath returns the URL path to the service service service-read HTTP endpoint.
func ServiceReadServicePath(id string) string {
	return fmt.Sprintf("/1/services2/%v", id)
}

// ServiceUpdateServicePath returns the URL path to the service service service-update HTTP endpoint.
func ServiceUpdateServicePath(id string) string {
	return fmt.Sprintf("/1/services2/%v", id)
}

// ServiceDeleteServicePath returns the URL path to the service service service-delete HTTP endpoint.
func ServiceDeleteServicePath(id string) string {
	return fmt.Sprintf("/1/services2/%v", id)
}

// JobListServicePath returns the URL path to the service service job-list HTTP endpoint.
func JobListServicePath(serviceID string) string {
	return fmt.Sprintf("/1/services2/%v/jobs", serviceID)
}

// JobCreateServicePath returns the URL path to the service service job-create HTTP endpoint.
func JobCreateServicePath(serviceID string) string {
	return fmt.Sprintf("/1/services2/%v/jobs", serviceID)
}

// JobReadServicePath returns the URL path to the service service job-read HTTP endpoint.
func JobReadServicePath(serviceID string, id string) string {
	return fmt.Sprintf("/1/services2/%v/jobs/%v", serviceID, id)
}

// JobOutputServicePath returns the URL path to the service service job-output HTTP endpoint.
func JobOutputServicePath(serviceID string, jobID string) string {
	return fmt.Sprintf("/1/services2/%v/jobs/%v/output", serviceID, jobID)
}
