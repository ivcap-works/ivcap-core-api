// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	order "github.com/ivcap-works/ivcap-core-api/gen/order"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

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

// BuildListPayload builds the payload for the order list endpoint from CLI
// flags.
func BuildListPayload(orderListLimit string, orderListPage string, orderListFilter string, orderListOrderBy string, orderListOrderDesc string, orderListAtTime string, orderListJWT string) (*order.ListPayload, error) {
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
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
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
func BuildLogsPayload(orderLogsBody string, orderLogsJWT string) (*order.LogsPayload, error) {
	var err error
	var body LogsRequestBody
	{
		err = json.Unmarshal([]byte(orderLogsBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"container-name\": \"main\",\n      \"from\": 1257894000,\n      \"namespace-name\": \"ivcap-develop-runner\",\n      \"order\": \"urn:ivcap:order:123e4567-e89b-12d3-a456-426614174000\",\n      \"to\": 1257894000\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.order", body.Order, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = orderLogsJWT
	}
	v := &order.DownloadLogRequestT{
		From:          body.From,
		To:            body.To,
		NamespaceName: body.NamespaceName,
		ContainerName: body.ContainerName,
		Order:         body.Order,
	}
	res := &order.LogsPayload{
		DownloadLogRequest: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildTopPayload builds the payload for the order top endpoint from CLI flags.
func BuildTopPayload(orderTopBody string, orderTopJWT string) (*order.TopPayload, error) {
	var err error
	var body TopRequestBody
	{
		err = json.Unmarshal([]byte(orderTopBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"namespace-name\": \"ivcap-develop-runner\",\n      \"order\": \"urn:ivcap:order:123e4567-e89b-12d3-a456-426614174000\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.order", body.Order, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = orderTopJWT
	}
	v := &order.OrderTopRequestT{
		Order:         body.Order,
		NamespaceName: body.NamespaceName,
	}
	res := &order.TopPayload{
		OrderTopRequest: v,
	}
	res.JWT = jwt

	return res, nil
}
