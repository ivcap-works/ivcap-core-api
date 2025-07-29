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

// ReadAspectPath returns the URL path to the aspect service read HTTP endpoint.
func ReadAspectPath(id string) string {
	return fmt.Sprintf("/1/aspects/%v", id)
}

// ListAspectPath returns the URL path to the aspect service list HTTP endpoint.
func ListAspectPath() string {
	return "/1/aspects"
}

// CreateAspectPath returns the URL path to the aspect service create HTTP endpoint.
func CreateAspectPath() string {
	return "/1/aspects"
}

// UpdateAspectPath returns the URL path to the aspect service update HTTP endpoint.
func UpdateAspectPath() string {
	return "/1/aspects"
}

// RetractAspectPath returns the URL path to the aspect service retract HTTP endpoint.
func RetractAspectPath(id string) string {
	return fmt.Sprintf("/1/aspects/%v", id)
}
