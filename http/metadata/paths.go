// Copyright 2023 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	"fmt"
)

// ListMetadataPath returns the URL path to the metadata service list HTTP endpoint.
func ListMetadataPath() string {
	return "/1/metadata"
}

// ReadMetadataPath returns the URL path to the metadata service read HTTP endpoint.
func ReadMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}

// AddMetadataPath returns the URL path to the metadata service add HTTP endpoint.
func AddMetadataPath() string {
	return "/1/metadata"
}

// UpdateOneMetadataPath returns the URL path to the metadata service update_one HTTP endpoint.
func UpdateOneMetadataPath() string {
	return "/1/metadata"
}

// UpdateRecordMetadataPath returns the URL path to the metadata service update_record HTTP endpoint.
func UpdateRecordMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}

// RevokeMetadataPath returns the URL path to the metadata service revoke HTTP endpoint.
func RevokeMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}
