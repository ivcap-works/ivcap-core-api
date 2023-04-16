// Copyright 2023 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	service "github.com/reinventingscience/ivcap-core-api/gen/service"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the service list endpoint from CLI
// flags.
func BuildListPayload(serviceListFilter string, serviceListOrderby string, serviceListTop string, serviceListSkip string, serviceListSelect string, serviceListOffset string, serviceListLimit string, serviceListPageToken string) (*service.ListPayload, error) {
	var err error
	var filter string
	{
		if serviceListFilter != "" {
			filter = serviceListFilter
		}
	}
	var orderby string
	{
		if serviceListOrderby != "" {
			orderby = serviceListOrderby
		}
	}
	var top int
	{
		if serviceListTop != "" {
			var v int64
			v, err = strconv.ParseInt(serviceListTop, 10, strconv.IntSize)
			top = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for top, must be INT")
			}
			if top < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("top", top, 1, true))
			}
			if top > 50 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("top", top, 50, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var skip int
	{
		if serviceListSkip != "" {
			var v int64
			v, err = strconv.ParseInt(serviceListSkip, 10, strconv.IntSize)
			skip = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for skip, must be INT")
			}
			if skip < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("skip", skip, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var select_ string
	{
		if serviceListSelect != "" {
			select_ = serviceListSelect
		}
	}
	var offset *int
	{
		if serviceListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(serviceListOffset, 10, strconv.IntSize)
			val := int(v)
			offset = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if *offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", *offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit *int
	{
		if serviceListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(serviceListLimit, 10, strconv.IntSize)
			val := int(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if *limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", *limit, 1, true))
			}
			if *limit > 50 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", *limit, 50, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var pageToken string
	{
		if serviceListPageToken != "" {
			pageToken = serviceListPageToken
		}
	}
	v := &service.ListPayload{}
	v.Filter = filter
	v.Orderby = orderby
	v.Top = top
	v.Skip = skip
	v.Select = select_
	v.Offset = offset
	v.Limit = limit
	v.PageToken = pageToken

	return v, nil
}

// BuildCreatePayload builds the payload for the service create endpoint from
// CLI flags.
func BuildCreatePayload(serviceCreateBody string, serviceCreateJWT string) (*service.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(serviceCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"accountID\": \"123e4567-e89b-12d3-a456-426614174000\",\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"name\": \"region\",\n            \"value\": \"Upper Valley\"\n         },\n         {\n            \"name\": \"threshold\",\n            \"value\": 10\n         }\n      ],\n      \"serviceID\": \"123e4567-e89b-12d3-a456-426614174000\"\n   }'")
		}
		if body.Workflow == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("workflow", "body"))
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.provider-id", body.ProviderID, goa.FormatURI))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account-id", body.AccountID, goa.FormatURI))
		for _, e := range body.References {
			if e != nil {
				if err2 := ValidateReferenceTRequestBodyRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if body.Banner != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.banner", *body.Banner, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = serviceCreateJWT
	}
	v := &service.ServiceDescriptionT{
		ProviderRef: body.ProviderRef,
		ProviderID:  body.ProviderID,
		Description: body.Description,
		AccountID:   body.AccountID,
		Banner:      body.Banner,
		Name:        body.Name,
	}
	if body.Metadata != nil {
		v.Metadata = make([]*service.ParameterT, len(body.Metadata))
		for i, val := range body.Metadata {
			v.Metadata[i] = marshalParameterTRequestBodyRequestBodyToServiceParameterT(val)
		}
	}
	if body.References != nil {
		v.References = make([]*service.ReferenceT, len(body.References))
		for i, val := range body.References {
			v.References[i] = marshalReferenceTRequestBodyRequestBodyToServiceReferenceT(val)
		}
	}
	if body.Workflow != nil {
		v.Workflow = marshalWorkflowTRequestBodyRequestBodyToServiceWorkflowT(body.Workflow)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.Parameters != nil {
		v.Parameters = make([]*service.ParameterDefT, len(body.Parameters))
		for i, val := range body.Parameters {
			v.Parameters[i] = marshalParameterDefTToServiceParameterDefT(val)
		}
	}
	res := &service.CreatePayload{
		Services: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildReadPayload builds the payload for the service read endpoint from CLI
// flags.
func BuildReadPayload(serviceReadID string) (*service.ReadPayload, error) {
	var id string
	{
		id = serviceReadID
	}
	v := &service.ReadPayload{}
	v.ID = id

	return v, nil
}

// BuildUpdatePayload builds the payload for the service update endpoint from
// CLI flags.
func BuildUpdatePayload(serviceUpdateBody string, serviceUpdateID string, serviceUpdateForceCreate string, serviceUpdateJWT string) (*service.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(serviceUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"account-id\": \"cayp:account:acme\",\n      \"banner\": \"http://harber.org/guiseppe\",\n      \"description\": \"This service ...\",\n      \"metadata\": [\n         {\n            \"name\": \"Quis rerum dignissimos.\",\n            \"value\": \"Expedita quia deserunt veritatis sequi voluptas.\"\n         },\n         {\n            \"name\": \"Quis rerum dignissimos.\",\n            \"value\": \"Expedita quia deserunt veritatis sequi voluptas.\"\n         },\n         {\n            \"name\": \"Quis rerum dignissimos.\",\n            \"value\": \"Expedita quia deserunt veritatis sequi voluptas.\"\n         }\n      ],\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"provider-id\": \"cayp:provider:acme\",\n      \"provider-ref\": \"service_foo_patch_1\",\n      \"references\": [\n         {\n            \"title\": \"Quod nihil aperiam eligendi ut.\",\n            \"uri\": \"http://schowaltercrist.net/reynold\"\n         },\n         {\n            \"title\": \"Quod nihil aperiam eligendi ut.\",\n            \"uri\": \"http://schowaltercrist.net/reynold\"\n         },\n         {\n            \"title\": \"Quod nihil aperiam eligendi ut.\",\n            \"uri\": \"http://schowaltercrist.net/reynold\"\n         }\n      ],\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ],\n      \"workflow\": {\n         \"argo\": \"Reprehenderit molestiae cupiditate voluptas et voluptatibus illum.\",\n         \"basic\": {\n            \"command\": [\n               \"Aut voluptas.\",\n               \"Ut officiis consequatur corporis autem odit.\",\n               \"Unde fuga sed veniam.\"\n            ],\n            \"cpu\": {\n               \"limit\": \"Quidem nulla quae provident dolor amet nulla.\",\n               \"request\": \"Et aut autem deserunt sit architecto.\"\n            },\n            \"image\": \"Voluptatem explicabo aut adipisci.\",\n            \"memory\": {\n               \"limit\": \"Quidem nulla quae provident dolor amet nulla.\",\n               \"request\": \"Et aut autem deserunt sit architecto.\"\n            }\n         },\n         \"opts\": \"Deserunt fugiat hic eos quaerat voluptas distinctio.\",\n         \"type\": \"Pariatur aut.\"\n      }\n   }'")
		}
		if body.Workflow == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("workflow", "body"))
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.provider-id", body.ProviderID, goa.FormatURI))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.account-id", body.AccountID, goa.FormatURI))
		for _, e := range body.References {
			if e != nil {
				if err2 := ValidateReferenceTRequestBodyRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if body.Banner != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.banner", *body.Banner, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = serviceUpdateID
	}
	var forceCreate *bool
	{
		if serviceUpdateForceCreate != "" {
			var val bool
			val, err = strconv.ParseBool(serviceUpdateForceCreate)
			forceCreate = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for forceCreate, must be BOOL")
			}
		}
	}
	var jwt string
	{
		jwt = serviceUpdateJWT
	}
	v := &service.ServiceDescriptionT{
		ProviderRef: body.ProviderRef,
		ProviderID:  body.ProviderID,
		Description: body.Description,
		AccountID:   body.AccountID,
		Banner:      body.Banner,
		Name:        body.Name,
	}
	if body.Metadata != nil {
		v.Metadata = make([]*service.ParameterT, len(body.Metadata))
		for i, val := range body.Metadata {
			v.Metadata[i] = marshalParameterTRequestBodyRequestBodyToServiceParameterT(val)
		}
	}
	if body.References != nil {
		v.References = make([]*service.ReferenceT, len(body.References))
		for i, val := range body.References {
			v.References[i] = marshalReferenceTRequestBodyRequestBodyToServiceReferenceT(val)
		}
	}
	if body.Workflow != nil {
		v.Workflow = marshalWorkflowTRequestBodyRequestBodyToServiceWorkflowT(body.Workflow)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.Parameters != nil {
		v.Parameters = make([]*service.ParameterDefT, len(body.Parameters))
		for i, val := range body.Parameters {
			v.Parameters[i] = marshalParameterDefTToServiceParameterDefT(val)
		}
	}
	res := &service.UpdatePayload{
		Services: v,
	}
	res.ID = &id
	res.ForceCreate = forceCreate
	res.JWT = jwt

	return res, nil
}

// BuildDeletePayload builds the payload for the service delete endpoint from
// CLI flags.
func BuildDeletePayload(serviceDeleteID string, serviceDeleteJWT string) (*service.DeletePayload, error) {
	var id string
	{
		id = serviceDeleteID
	}
	var jwt string
	{
		jwt = serviceDeleteJWT
	}
	v := &service.DeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}
