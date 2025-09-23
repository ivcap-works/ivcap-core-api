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
	"encoding/json"
	"fmt"
	"strconv"

	search "github.com/ivcap-works/ivcap-core-api/gen/search"
	goa "goa.design/goa/v3/pkg"
)

// BuildSearchPayload builds the payload for the search search endpoint from
// CLI flags.
func BuildSearchPayload(searchSearchBody string, searchSearchAtTime string, searchSearchLimit string, searchSearchPage string, searchSearchJWT string, searchSearchContentType string) (*search.SearchPayload, error) {
	var err error
	var body []byte
	{
		body = []byte(searchSearchBody)
	}
	var atTime *string
	{
		if searchSearchAtTime != "" {
			atTime = &searchSearchAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if searchSearchLimit != "" {
			var v int64
			v, err = strconv.ParseInt(searchSearchLimit, 10, strconv.IntSize)
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
	var page any
	{
		if searchSearchPage != "" {
			err = json.Unmarshal([]byte(searchSearchPage), &page)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for page, \nerror: %s, \nexample of valid JSON:\n%s", err, "\"gdsgQwhdgd\"")
			}
		}
	}
	var jwt string
	{
		jwt = searchSearchJWT
	}
	var contentType string
	{
		contentType = searchSearchContentType
	}
	v := body
	res := &search.SearchPayload{
		Query: v,
	}
	res.AtTime = atTime
	res.Limit = limit
	res.Page = page
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}
