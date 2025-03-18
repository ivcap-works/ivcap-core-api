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

// ListServicexPath returns the URL path to the servicex service list HTTP endpoint.
func ListServicexPath() string {
	return "/1/services"
}

// CreateServiceServicexPath returns the URL path to the servicex service create_service HTTP endpoint.
func CreateServiceServicexPath() string {
	return "/1/services"
}

// ReadServicexPath returns the URL path to the servicex service read HTTP endpoint.
func ReadServicexPath(id string) string {
	return fmt.Sprintf("/1/services/%v", id)
}

// UpdateServicexPath returns the URL path to the servicex service update HTTP endpoint.
func UpdateServicexPath(id string) string {
	return fmt.Sprintf("/1/services/%v", id)
}

// DeleteServicexPath returns the URL path to the servicex service delete HTTP endpoint.
func DeleteServicexPath(id string) string {
	return fmt.Sprintf("/1/services/%v", id)
}
