// Copyright 2026 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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
	"strconv"

	package_ "github.com/ivcap-works/ivcap-core-api/gen/package_"
)

// BuildListPayload builds the payload for the package list endpoint from CLI
// flags.
func BuildListPayload(package_ListTag string, package_ListPage string, package_ListLimit string, package_ListJWT string) (*package_.ListPayload, error) {
	var err error
	var tag *string
	{
		if package_ListTag != "" {
			tag = &package_ListTag
		}
	}
	var page *string
	{
		if package_ListPage != "" {
			page = &package_ListPage
		}
	}
	var limit *int
	{
		if package_ListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(package_ListLimit, 10, strconv.IntSize)
			val := int(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
		}
	}
	var jwt string
	{
		jwt = package_ListJWT
	}
	v := &package_.ListPayload{}
	v.Tag = tag
	v.Page = page
	v.Limit = limit
	v.JWT = jwt

	return v, nil
}

// BuildRemovePayload builds the payload for the package remove endpoint from
// CLI flags.
func BuildRemovePayload(package_RemoveTag string, package_RemoveJWT string) (*package_.RemovePayload, error) {
	var tag string
	{
		tag = package_RemoveTag
	}
	var jwt string
	{
		jwt = package_RemoveJWT
	}
	v := &package_.RemovePayload{}
	v.Tag = tag
	v.JWT = jwt

	return v, nil
}
