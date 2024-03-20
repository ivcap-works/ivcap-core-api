// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// ReadOrderPath returns the URL path to the order service read HTTP endpoint.
func ReadOrderPath(id string) string {
	return fmt.Sprintf("/1/orders/%v", id)
}

// ListOrderPath returns the URL path to the order service list HTTP endpoint.
func ListOrderPath() string {
	return "/1/orders"
}

// CreateOrderPath returns the URL path to the order service create HTTP endpoint.
func CreateOrderPath() string {
	return "/1/orders"
}

// LogsOrderPath returns the URL path to the order service logs HTTP endpoint.
func LogsOrderPath(orderID string) string {
	return fmt.Sprintf("/1/orders/%v/logs", orderID)
}

// TopOrderPath returns the URL path to the order service top HTTP endpoint.
func TopOrderPath(orderID string) string {
	return fmt.Sprintf("/1/orders/%v/top", orderID)
}
