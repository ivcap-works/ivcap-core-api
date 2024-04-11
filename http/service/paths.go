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
	"fmt"
)

// ListServicePath returns the URL path to the service service list HTTP endpoint.
func ListServicePath() string {
	return "/1/services"
}

// CreateServiceServicePath returns the URL path to the service service create_service HTTP endpoint.
func CreateServiceServicePath() string {
	return "/1/services"
}

// ReadServicePath returns the URL path to the service service read HTTP endpoint.
func ReadServicePath(id string) string {
	return fmt.Sprintf("/1/services/%v", id)
}

// UpdateServicePath returns the URL path to the service service update HTTP endpoint.
func UpdateServicePath(id string) string {
	return fmt.Sprintf("/1/services/%v", id)
}

// DeleteServicePath returns the URL path to the service service delete HTTP endpoint.
func DeleteServicePath(id string) string {
	return fmt.Sprintf("/1/services/%v", id)
}
