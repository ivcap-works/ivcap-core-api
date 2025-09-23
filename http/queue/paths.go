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

// CreateQueuePath returns the URL path to the queue service create HTTP endpoint.
func CreateQueuePath() string {
	return "/1/queues"
}

// ReadQueuePath returns the URL path to the queue service read HTTP endpoint.
func ReadQueuePath(id string) string {
	return fmt.Sprintf("/1/queues/%v", id)
}

// DeleteQueuePath returns the URL path to the queue service delete HTTP endpoint.
func DeleteQueuePath(id string) string {
	return fmt.Sprintf("/1/queues/%v", id)
}

// ListQueuePath returns the URL path to the queue service list HTTP endpoint.
func ListQueuePath() string {
	return "/1/queues"
}

// EnqueueQueuePath returns the URL path to the queue service enqueue HTTP endpoint.
func EnqueueQueuePath(id string) string {
	return fmt.Sprintf("/1/queues/%v/messages", id)
}

// DequeueQueuePath returns the URL path to the queue service dequeue HTTP endpoint.
func DequeueQueuePath(id string) string {
	return fmt.Sprintf("/1/queues/%v/messages", id)
}
