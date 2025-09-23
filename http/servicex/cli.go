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

	servicex "github.com/ivcap-works/ivcap-core-api/gen/servicex"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the servicex list endpoint from CLI
// flags.
func BuildListPayload(servicexListLimit string, servicexListPage string, servicexListFilter string, servicexListOrderBy string, servicexListOrderDesc string, servicexListAtTime string, servicexListJWT string) (*servicex.ListPayload, error) {
	var err error
	var limit int
	{
		if servicexListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(servicexListLimit, 10, strconv.IntSize)
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
		if servicexListPage != "" {
			page = &servicexListPage
		}
	}
	var filter *string
	{
		if servicexListFilter != "" {
			filter = &servicexListFilter
		}
	}
	var orderBy *string
	{
		if servicexListOrderBy != "" {
			orderBy = &servicexListOrderBy
		}
	}
	var orderDesc bool
	{
		if servicexListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(servicexListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if servicexListAtTime != "" {
			atTime = &servicexListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = servicexListJWT
	}
	v := &servicex.ListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildCreateServicePayload builds the payload for the servicex create_service
// endpoint from CLI flags.
func BuildCreateServicePayload(servicexCreateServiceBody string, servicexCreateServiceJWT string) (*servicex.CreateServicePayload, error) {
	var err error
	var body CreateServiceRequestBody
	{
		err = json.Unmarshal([]byte(servicexCreateServiceBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"banner\": \"http://keelingbatz.biz/ally\",\n      \"description\": \"This service ...\",\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"policy\": \"urn:ivcap:policy:123e4567-e89b-12d3-a456-426614174000\",\n      \"references\": [\n         {\n            \"title\": \"Fuga veritatis qui asperiores quaerat molestiae.\",\n            \"uri\": \"http://jacobsonolson.name/rosalind_schmidt\"\n         },\n         {\n            \"title\": \"Fuga veritatis qui asperiores quaerat molestiae.\",\n            \"uri\": \"http://jacobsonolson.name/rosalind_schmidt\"\n         }\n      ],\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ],\n      \"workflow\": {\n         \"argo\": \"Repellat nemo recusandae asperiores dolorum voluptas.\",\n         \"basic\": {\n            \"command\": [\n               \"/bin/sh\",\n               \"-c\",\n               \"echo $PATH\"\n            ],\n            \"cpu\": {\n               \"limit\": \"100m\",\n               \"request\": \"10m\"\n            },\n            \"ephemeral-storage\": {\n               \"limit\": \"4Gi\",\n               \"request\": \"2Gi\"\n            },\n            \"gpu-number\": 2,\n            \"gpu-type\": \"nvidia-tesla-t4\",\n            \"image\": \"alpine\",\n            \"image-pull-policy\": \"Praesentium saepe minima non qui quos voluptate.\",\n            \"memory\": {\n               \"limit\": \"100Mi\",\n               \"request\": \"10Mi\"\n            },\n            \"shared-memory\": \"1Gi\"\n         },\n         \"type\": \"basic\"\n      }\n   }'")
		}
		if body.Workflow == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("workflow", "body"))
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		for _, e := range body.References {
			if e != nil {
				if err2 := ValidateXReferenceTRequestBodyRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if body.Banner != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.banner", *body.Banner, goa.FormatURI))
		}
		if body.Workflow != nil {
			if err2 := ValidateXWorkflowTRequestBodyRequestBody(body.Workflow); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
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
		jwt = servicexCreateServiceJWT
	}
	v := &servicex.XServiceDefinitionT{
		Description: body.Description,
		Banner:      body.Banner,
		Policy:      body.Policy,
		Name:        body.Name,
	}
	if body.References != nil {
		v.References = make([]*servicex.XReferenceT, len(body.References))
		for i, val := range body.References {
			v.References[i] = marshalXReferenceTRequestBodyRequestBodyToServicexXReferenceT(val)
		}
	}
	if body.Workflow != nil {
		v.Workflow = marshalXWorkflowTRequestBodyRequestBodyToServicexXWorkflowT(body.Workflow)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.Parameters != nil {
		v.Parameters = make([]*servicex.ParameterDefT, len(body.Parameters))
		for i, val := range body.Parameters {
			v.Parameters[i] = marshalParameterDefTToServicexParameterDefT(val)
		}
	} else {
		v.Parameters = []*servicex.ParameterDefT{}
	}
	res := &servicex.CreateServicePayload{
		Services: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildReadPayload builds the payload for the servicex read endpoint from CLI
// flags.
func BuildReadPayload(servicexReadID string, servicexReadJWT string) (*servicex.ReadPayload, error) {
	var id string
	{
		id = servicexReadID
	}
	var jwt string
	{
		jwt = servicexReadJWT
	}
	v := &servicex.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdatePayload builds the payload for the servicex update endpoint from
// CLI flags.
func BuildUpdatePayload(servicexUpdateBody string, servicexUpdateID string, servicexUpdateForceCreate string, servicexUpdateJWT string) (*servicex.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(servicexUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"banner\": \"http://quigley.net/hollie\",\n      \"description\": \"This service ...\",\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"policy\": \"urn:ivcap:policy:123e4567-e89b-12d3-a456-426614174000\",\n      \"references\": [\n         {\n            \"title\": \"Fuga veritatis qui asperiores quaerat molestiae.\",\n            \"uri\": \"http://jacobsonolson.name/rosalind_schmidt\"\n         },\n         {\n            \"title\": \"Fuga veritatis qui asperiores quaerat molestiae.\",\n            \"uri\": \"http://jacobsonolson.name/rosalind_schmidt\"\n         },\n         {\n            \"title\": \"Fuga veritatis qui asperiores quaerat molestiae.\",\n            \"uri\": \"http://jacobsonolson.name/rosalind_schmidt\"\n         },\n         {\n            \"title\": \"Fuga veritatis qui asperiores quaerat molestiae.\",\n            \"uri\": \"http://jacobsonolson.name/rosalind_schmidt\"\n         }\n      ],\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ],\n      \"workflow\": {\n         \"argo\": \"Repellat nemo recusandae asperiores dolorum voluptas.\",\n         \"basic\": {\n            \"command\": [\n               \"/bin/sh\",\n               \"-c\",\n               \"echo $PATH\"\n            ],\n            \"cpu\": {\n               \"limit\": \"100m\",\n               \"request\": \"10m\"\n            },\n            \"ephemeral-storage\": {\n               \"limit\": \"4Gi\",\n               \"request\": \"2Gi\"\n            },\n            \"gpu-number\": 2,\n            \"gpu-type\": \"nvidia-tesla-t4\",\n            \"image\": \"alpine\",\n            \"image-pull-policy\": \"Praesentium saepe minima non qui quos voluptate.\",\n            \"memory\": {\n               \"limit\": \"100Mi\",\n               \"request\": \"10Mi\"\n            },\n            \"shared-memory\": \"1Gi\"\n         },\n         \"type\": \"basic\"\n      }\n   }'")
		}
		if body.Workflow == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("workflow", "body"))
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		for _, e := range body.References {
			if e != nil {
				if err2 := ValidateXReferenceTRequestBodyRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if body.Banner != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.banner", *body.Banner, goa.FormatURI))
		}
		if body.Workflow != nil {
			if err2 := ValidateXWorkflowTRequestBodyRequestBody(body.Workflow); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Policy != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", *body.Policy, goa.FormatURI))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = servicexUpdateID
	}
	var forceCreate *bool
	{
		if servicexUpdateForceCreate != "" {
			var val bool
			val, err = strconv.ParseBool(servicexUpdateForceCreate)
			forceCreate = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for forceCreate, must be BOOL")
			}
		}
	}
	var jwt string
	{
		jwt = servicexUpdateJWT
	}
	v := &servicex.XServiceDefinitionT{
		Description: body.Description,
		Banner:      body.Banner,
		Policy:      body.Policy,
		Name:        body.Name,
	}
	if body.References != nil {
		v.References = make([]*servicex.XReferenceT, len(body.References))
		for i, val := range body.References {
			v.References[i] = marshalXReferenceTRequestBodyRequestBodyToServicexXReferenceT(val)
		}
	}
	if body.Workflow != nil {
		v.Workflow = marshalXWorkflowTRequestBodyRequestBodyToServicexXWorkflowT(body.Workflow)
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	if body.Parameters != nil {
		v.Parameters = make([]*servicex.ParameterDefT, len(body.Parameters))
		for i, val := range body.Parameters {
			v.Parameters[i] = marshalParameterDefTToServicexParameterDefT(val)
		}
	} else {
		v.Parameters = []*servicex.ParameterDefT{}
	}
	res := &servicex.UpdatePayload{
		Services: v,
	}
	res.ID = &id
	res.ForceCreate = forceCreate
	res.JWT = jwt

	return res, nil
}

// BuildDeletePayload builds the payload for the servicex delete endpoint from
// CLI flags.
func BuildDeletePayload(servicexDeleteID string, servicexDeleteJWT string) (*servicex.DeletePayload, error) {
	var id string
	{
		id = servicexDeleteID
	}
	var jwt string
	{
		jwt = servicexDeleteJWT
	}
	v := &servicex.DeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}
