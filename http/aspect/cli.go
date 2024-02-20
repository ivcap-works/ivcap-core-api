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
	"encoding/json"
	"fmt"
	"strconv"

	aspect "github.com/ivcap-works/ivcap-core-api/gen/aspect"
	goa "goa.design/goa/v3/pkg"
)

// BuildReadPayload builds the payload for the aspect read endpoint from CLI
// flags.
func BuildReadPayload(aspectReadID string, aspectReadJWT string) (*aspect.ReadPayload, error) {
	var id string
	{
		id = aspectReadID
	}
	var jwt string
	{
		jwt = aspectReadJWT
	}
	v := &aspect.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildListPayload builds the payload for the aspect list endpoint from CLI
// flags.
func BuildListPayload(aspectListEntity string, aspectListSchema string, aspectListContentPath string, aspectListAtTime string, aspectListLimit string, aspectListFilter string, aspectListOrderBy string, aspectListOrderDirection string, aspectListIncludeContent string, aspectListPage string, aspectListJWT string) (*aspect.ListPayload, error) {
	var err error
	var entity *string
	{
		if aspectListEntity != "" {
			entity = &aspectListEntity
			err = goa.MergeErrors(err, goa.ValidateFormat("entity", *entity, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var schema *string
	{
		if aspectListSchema != "" {
			schema = &aspectListSchema
		}
	}
	var contentPath *string
	{
		if aspectListContentPath != "" {
			contentPath = &aspectListContentPath
		}
	}
	var atTime *string
	{
		if aspectListAtTime != "" {
			atTime = &aspectListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if aspectListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(aspectListLimit, 10, strconv.IntSize)
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
	var filter string
	{
		if aspectListFilter != "" {
			filter = aspectListFilter
		}
	}
	var orderBy string
	{
		if aspectListOrderBy != "" {
			orderBy = aspectListOrderBy
		}
	}
	var orderDirection string
	{
		if aspectListOrderDirection != "" {
			orderDirection = aspectListOrderDirection
		}
	}
	var includeContent *bool
	{
		if aspectListIncludeContent != "" {
			var val bool
			val, err = strconv.ParseBool(aspectListIncludeContent)
			includeContent = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for includeContent, must be BOOL")
			}
		}
	}
	var page *string
	{
		if aspectListPage != "" {
			page = &aspectListPage
		}
	}
	var jwt string
	{
		jwt = aspectListJWT
	}
	v := &aspect.ListPayload{}
	v.Entity = entity
	v.Schema = schema
	v.ContentPath = contentPath
	v.AtTime = atTime
	v.Limit = limit
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDirection = orderDirection
	v.IncludeContent = includeContent
	v.Page = page
	v.JWT = jwt

	return v, nil
}

// BuildCreatePayload builds the payload for the aspect create endpoint from
// CLI flags.
func BuildCreatePayload(aspectCreateBody string, aspectCreateEntity string, aspectCreateSchema string, aspectCreatePolicy string, aspectCreateJWT string, aspectCreateContentType string) (*aspect.CreatePayload, error) {
	var err error
	var body any
	{
		err = json.Unmarshal([]byte(aspectCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "\"{\\\"$schema\\\": ...}\"")
		}
	}
	var entity string
	{
		entity = aspectCreateEntity
	}
	var schema string
	{
		schema = aspectCreateSchema
	}
	var policy *string
	{
		if aspectCreatePolicy != "" {
			policy = &aspectCreatePolicy
		}
	}
	var jwt string
	{
		jwt = aspectCreateJWT
	}
	var contentType string
	{
		contentType = aspectCreateContentType
	}
	v := body
	res := &aspect.CreatePayload{
		Content: &v,
	}
	res.Entity = entity
	res.Schema = schema
	res.Policy = policy
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}

// BuildUpdatePayload builds the payload for the aspect update endpoint from
// CLI flags.
func BuildUpdatePayload(aspectUpdateBody string, aspectUpdateID string, aspectUpdateEntity string, aspectUpdateSchema string, aspectUpdateJWT string, aspectUpdateContentType string) (*aspect.UpdatePayload, error) {
	var err error
	var body any
	{
		err = json.Unmarshal([]byte(aspectUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "\"{\\\"$schema\\\": ...}\"")
		}
	}
	var id string
	{
		id = aspectUpdateID
	}
	var entity string
	{
		entity = aspectUpdateEntity
	}
	var schema string
	{
		schema = aspectUpdateSchema
		err = goa.MergeErrors(err, goa.ValidateFormat("schema", schema, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = aspectUpdateJWT
	}
	var contentType string
	{
		contentType = aspectUpdateContentType
	}
	v := body
	res := &aspect.UpdatePayload{
		Content: v,
	}
	res.ID = id
	res.Entity = entity
	res.Schema = schema
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}

// BuildRetractPayload builds the payload for the aspect retract endpoint from
// CLI flags.
func BuildRetractPayload(aspectRetractID string, aspectRetractJWT string) (*aspect.RetractPayload, error) {
	var err error
	var id string
	{
		id = aspectRetractID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = aspectRetractJWT
	}
	v := &aspect.RetractPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}
