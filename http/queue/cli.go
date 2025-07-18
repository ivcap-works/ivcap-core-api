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

	queue "github.com/ivcap-works/ivcap-core-api/gen/queue"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the queue create endpoint from CLI
// flags.
func BuildCreatePayload(queueCreateBody string, queueCreateJWT string) (*queue.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(queueCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Events for the event service\",\n      \"name\": \"events\",\n      \"policy\": \"urn:ivcap:policy:123e4567-e89b-12d3-a456-426614174000\"\n   }'")
		}
		if body.Policy != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = queueCreateJWT
	}
	v := &queue.PayloadForCreateEndpoint{
		Name:        body.Name,
		Description: body.Description,
		Policy:      body.Policy,
	}
	res := &queue.CreatePayload{
		Queues: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildReadPayload builds the payload for the queue read endpoint from CLI
// flags.
func BuildReadPayload(queueReadID string, queueReadJWT string) (*queue.ReadPayload, error) {
	var id string
	{
		id = queueReadID
	}
	var jwt string
	{
		jwt = queueReadJWT
	}
	v := &queue.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeletePayload builds the payload for the queue delete endpoint from CLI
// flags.
func BuildDeletePayload(queueDeleteID string, queueDeleteJWT string) (*queue.DeletePayload, error) {
	var id string
	{
		id = queueDeleteID
	}
	var jwt string
	{
		jwt = queueDeleteJWT
	}
	v := &queue.DeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildListPayload builds the payload for the queue list endpoint from CLI
// flags.
func BuildListPayload(queueListLimit string, queueListPage string, queueListFilter string, queueListOrderBy string, queueListOrderDesc string, queueListAtTime string, queueListJWT string) (*queue.ListPayload, error) {
	var err error
	var limit int
	{
		if queueListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(queueListLimit, 10, strconv.IntSize)
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
		if queueListPage != "" {
			page = &queueListPage
		}
	}
	var filter *string
	{
		if queueListFilter != "" {
			filter = &queueListFilter
		}
	}
	var orderBy *string
	{
		if queueListOrderBy != "" {
			orderBy = &queueListOrderBy
		}
	}
	var orderDesc bool
	{
		if queueListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(queueListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if queueListAtTime != "" {
			atTime = &queueListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = queueListJWT
	}
	v := &queue.ListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildEnqueuePayload builds the payload for the queue enqueue endpoint from
// CLI flags.
func BuildEnqueuePayload(queueEnqueueBody string, queueEnqueueID string, queueEnqueueSchema string, queueEnqueueJWT string, queueEnqueueContentType string) (*queue.EnqueuePayload, error) {
	var err error
	var body any
	{
		err = json.Unmarshal([]byte(queueEnqueueBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "\"{\\\"temperature\\\": \\\"21\\\", \\\"location\\\": \\\"Buoy101\\\", \\\"timestamp\\\": \\\"2024-05-20T14:30:00Z\\\"}\"")
		}
	}
	var id string
	{
		id = queueEnqueueID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var schema *string
	{
		if queueEnqueueSchema != "" {
			schema = &queueEnqueueSchema
			err = goa.MergeErrors(err, goa.ValidateFormat("schema", *schema, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = queueEnqueueJWT
	}
	var contentType *string
	{
		if queueEnqueueContentType != "" {
			contentType = &queueEnqueueContentType
		}
	}
	v := body
	res := &queue.EnqueuePayload{
		Content: &v,
	}
	res.ID = id
	res.Schema = schema
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}

// BuildDequeuePayload builds the payload for the queue dequeue endpoint from
// CLI flags.
func BuildDequeuePayload(queueDequeueID string, queueDequeueLimit string, queueDequeueJWT string) (*queue.DequeuePayload, error) {
	var err error
	var id string
	{
		id = queueDequeueID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var limit *int
	{
		if queueDequeueLimit != "" {
			var v int64
			v, err = strconv.ParseInt(queueDequeueLimit, 10, strconv.IntSize)
			val := int(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
		}
	}
	var jwt string
	{
		jwt = queueDequeueJWT
	}
	v := &queue.DequeuePayload{}
	v.ID = id
	v.Limit = limit
	v.JWT = jwt

	return v, nil
}
