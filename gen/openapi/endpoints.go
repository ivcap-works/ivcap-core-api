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

package openapi

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "openapi" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "openapi" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "openapi" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
