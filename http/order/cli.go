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

	order "github.com/ivcap-works/ivcap-core-api/gen/order"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the order list endpoint from CLI
// flags.
func BuildListPayload(orderListLimit string, orderListOffset string, orderListPage string, orderListFilter string, orderListOrderBy string, orderListOrderDesc string, orderListAtTime string, orderListJWT string) (*order.ListPayload, error) {
	var err error
	var limit int
	{
		if orderListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(orderListLimit, 10, strconv.IntSize)
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
	var offset int
	{
		if orderListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(orderListOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
		}
	}
	var page *string
	{
		if orderListPage != "" {
			page = &orderListPage
		}
	}
	var filter *string
	{
		if orderListFilter != "" {
			filter = &orderListFilter
		}
	}
	var orderBy *string
	{
		if orderListOrderBy != "" {
			orderBy = &orderListOrderBy
		}
	}
	var orderDesc bool
	{
		if orderListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(orderListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if orderListAtTime != "" {
			atTime = &orderListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = orderListJWT
	}
	v := &order.ListPayload{}
	v.Limit = limit
	v.Offset = offset
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildReadPayload builds the payload for the order read endpoint from CLI
// flags.
func BuildReadPayload(orderReadID string, orderReadJWT string) (*order.ReadPayload, error) {
	var id string
	{
		id = orderReadID
	}
	var jwt string
	{
		jwt = orderReadJWT
	}
	v := &order.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildProductsPayload builds the payload for the order products endpoint from
// CLI flags.
func BuildProductsPayload(orderProductsOrderID string, orderProductsLimit string, orderProductsPage string, orderProductsJWT string) (*order.ProductsPayload, error) {
	var err error
	var orderID string
	{
		orderID = orderProductsOrderID
		err = goa.MergeErrors(err, goa.ValidateFormat("orderID", orderID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var limit int
	{
		if orderProductsLimit != "" {
			var v int64
			v, err = strconv.ParseInt(orderProductsLimit, 10, strconv.IntSize)
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
		if orderProductsPage != "" {
			page = &orderProductsPage
		}
	}
	var jwt string
	{
		jwt = orderProductsJWT
	}
	v := &order.ProductsPayload{}
	v.OrderID = orderID
	v.Limit = limit
	v.Page = page
	v.JWT = jwt

	return v, nil
}

// BuildCreatePayload builds the payload for the order create endpoint from CLI
// flags.
func BuildCreatePayload(orderCreateBody string, orderCreateJWT string) (*order.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(orderCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"name\": \"region\",\n            \"value\": \"Upper Valley\"\n         },\n         {\n            \"name\": \"threshold\",\n            \"value\": \"10\"\n         }\n      ],\n      \"policy\": \"urn:ivcap:policy:123e4567-e89b-12d3-a456-426614174000\",\n      \"service\": \"urn:ivcap:service:123e4567-e89b-12d3-a456-426614174000\",\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ]\n   }'")
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.service", body.Service, goa.FormatURI))
		if body.Policy != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = orderCreateJWT
	}
	v := &order.OrderRequestT{
		Service: body.Service,
		Policy:  body.Policy,
		Name:    body.Name,
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.Parameters != nil {
		v.Parameters = make([]*order.ParameterT, len(body.Parameters))
		for i, val := range body.Parameters {
			v.Parameters[i] = marshalParameterTToOrderParameterT(val)
		}
	} else {
		v.Parameters = []*order.ParameterT{}
	}
	res := &order.CreatePayload{
		Orders: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildLogsPayload builds the payload for the order logs endpoint from CLI
// flags.
func BuildLogsPayload(orderLogsOrderID string, orderLogsFrom string, orderLogsTo string, orderLogsJWT string) (*order.LogsPayload, error) {
	var err error
	var orderID string
	{
		orderID = orderLogsOrderID
		err = goa.MergeErrors(err, goa.ValidateFormat("orderID", orderID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var from *int64
	{
		if orderLogsFrom != "" {
			val, err := strconv.ParseInt(orderLogsFrom, 10, 64)
			from = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for from, must be INT64")
			}
		}
	}
	var to *int64
	{
		if orderLogsTo != "" {
			val, err := strconv.ParseInt(orderLogsTo, 10, 64)
			to = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for to, must be INT64")
			}
		}
	}
	var jwt string
	{
		jwt = orderLogsJWT
	}
	v := &order.LogsPayload{}
	v.OrderID = orderID
	v.From = from
	v.To = to
	v.JWT = jwt

	return v, nil
}

// BuildTopPayload builds the payload for the order top endpoint from CLI flags.
func BuildTopPayload(orderTopOrderID string, orderTopJWT string) (*order.TopPayload, error) {
	var err error
	var orderID string
	{
		orderID = orderTopOrderID
		err = goa.MergeErrors(err, goa.ValidateFormat("orderID", orderID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = orderTopJWT
	}
	v := &order.TopPayload{}
	v.OrderID = orderID
	v.JWT = jwt

	return v, nil
}
