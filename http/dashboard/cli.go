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
	"strconv"

	dashboard "github.com/ivcap-works/ivcap-core-api/gen/dashboard"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the dashboard list endpoint from CLI
// flags.
func BuildListPayload(dashboardListLimit string, dashboardListPage string, dashboardListFilter string, dashboardListOrderBy string, dashboardListOrderDesc string, dashboardListAtTime string, dashboardListJWT string) (*dashboard.ListPayload, error) {
	var err error
	var limit int
	{
		if dashboardListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(dashboardListLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 50 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 50, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var page *string
	{
		if dashboardListPage != "" {
			page = &dashboardListPage
		}
	}
	var filter *string
	{
		if dashboardListFilter != "" {
			filter = &dashboardListFilter
		}
	}
	var orderBy *string
	{
		if dashboardListOrderBy != "" {
			orderBy = &dashboardListOrderBy
		}
	}
	var orderDesc bool
	{
		if dashboardListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(dashboardListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if dashboardListAtTime != "" {
			atTime = &dashboardListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = dashboardListJWT
	}
	v := &dashboard.ListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}
