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

	service "github.com/ivcap-works/ivcap-core-api/gen/service"
	goa "goa.design/goa/v3/pkg"
)

// BuildServiceListPayload builds the payload for the service service-list
// endpoint from CLI flags.
func BuildServiceListPayload(serviceServiceListLimit string, serviceServiceListPage string, serviceServiceListFilter string, serviceServiceListOrderBy string, serviceServiceListOrderDesc string, serviceServiceListAtTime string, serviceServiceListJWT string) (*service.ServiceListPayload, error) {
	var err error
	var limit int
	{
		if serviceServiceListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(serviceServiceListLimit, 10, strconv.IntSize)
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
		if serviceServiceListPage != "" {
			page = &serviceServiceListPage
		}
	}
	var filter *string
	{
		if serviceServiceListFilter != "" {
			filter = &serviceServiceListFilter
		}
	}
	var orderBy *string
	{
		if serviceServiceListOrderBy != "" {
			orderBy = &serviceServiceListOrderBy
		}
	}
	var orderDesc bool
	{
		if serviceServiceListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(serviceServiceListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if serviceServiceListAtTime != "" {
			atTime = &serviceServiceListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = serviceServiceListJWT
	}
	v := &service.ServiceListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildServiceCreatePayload builds the payload for the service service-create
// endpoint from CLI flags.
func BuildServiceCreatePayload(serviceServiceCreateBody string, serviceServiceCreateJWT string) (*service.ServiceCreatePayload, error) {
	var err error
	var body ServiceCreateRequestBody
	{
		err = json.Unmarshal([]byte(serviceServiceCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"controller\": [\n         {\n            \"$schema\": \"urn:ivcap:schema.service.rest.1\",\n            \"command\": [\n               \"python\",\n               \"/app/tool-service.py\"\n            ],\n            \"image\": \"your-docker-image:latest\",\n            \"port\": 8090,\n            \"ready-url\": \"/_healtz\",\n            \"resources\": {\n               \"limits\": {\n                  \"cpu\": \"500m\",\n                  \"ephemeral-storage\": \"1Gi\",\n                  \"memory\": \"1Gi\"\n               },\n               \"requests\": {\n                  \"cpu\": \"500m\",\n                  \"ephemeral-storage\": \"1Gi\",\n                  \"memory\": \"1Gi\"\n               }\n            }\n         }\n      ],\n      \"controller-schema\": \"urn:ivcap:schema.service.argo.1\",\n      \"description\": \"This service ...\",\n      \"id\": \"urn:ivcap:service:123e4567-e89b-12d3-a456-426614174000\",\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"policy\": \"urn:ivcap:policy:123e4567-e89b-12d3-a456-426614174000\",\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ]\n   }'")
		}
		if body.Controller == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("controller", "body"))
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.controller-schema", body.ControllerSchema, goa.FormatURI))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", body.Policy, goa.FormatURI))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", body.ID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = serviceServiceCreateJWT
	}
	v := &service.ServiceDefinitionT{
		ControllerSchema: body.ControllerSchema,
		Controller:       body.Controller,
		Policy:           body.Policy,
		ID:               body.ID,
		Name:             body.Name,
		Description:      body.Description,
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
	} else {
		v.Parameters = []*service.ParameterDefT{}
	}
	res := &service.ServiceCreatePayload{
		Service: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildServiceReadPayload builds the payload for the service service-read
// endpoint from CLI flags.
func BuildServiceReadPayload(serviceServiceReadID string, serviceServiceReadJWT string) (*service.ServiceReadPayload, error) {
	var id string
	{
		id = serviceServiceReadID
	}
	var jwt string
	{
		jwt = serviceServiceReadJWT
	}
	v := &service.ServiceReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildServiceUpdatePayload builds the payload for the service service-update
// endpoint from CLI flags.
func BuildServiceUpdatePayload(serviceServiceUpdateBody string, serviceServiceUpdateID string, serviceServiceUpdateForceCreate string, serviceServiceUpdateJWT string) (*service.ServiceUpdatePayload, error) {
	var err error
	var body ServiceUpdateRequestBody
	{
		err = json.Unmarshal([]byte(serviceServiceUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"controller\": [\n         {\n            \"$schema\": \"urn:ivcap:schema.service.rest.1\",\n            \"command\": [\n               \"python\",\n               \"/app/tool-service.py\"\n            ],\n            \"image\": \"your-docker-image:latest\",\n            \"port\": 8090,\n            \"ready-url\": \"/_healtz\",\n            \"resources\": {\n               \"limits\": {\n                  \"cpu\": \"500m\",\n                  \"ephemeral-storage\": \"1Gi\",\n                  \"memory\": \"1Gi\"\n               },\n               \"requests\": {\n                  \"cpu\": \"500m\",\n                  \"ephemeral-storage\": \"1Gi\",\n                  \"memory\": \"1Gi\"\n               }\n            }\n         }\n      ],\n      \"controller-schema\": \"urn:ivcap:schema.service.argo.1\",\n      \"description\": \"This service ...\",\n      \"id\": \"urn:ivcap:service:123e4567-e89b-12d3-a456-426614174000\",\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"policy\": \"urn:ivcap:policy:123e4567-e89b-12d3-a456-426614174000\",\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ]\n   }'")
		}
		if body.Controller == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("controller", "body"))
		}
		if body.Parameters == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "body"))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.controller-schema", body.ControllerSchema, goa.FormatURI))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.policy", body.Policy, goa.FormatURI))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", body.ID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = serviceServiceUpdateID
	}
	var forceCreate *bool
	{
		if serviceServiceUpdateForceCreate != "" {
			var val bool
			val, err = strconv.ParseBool(serviceServiceUpdateForceCreate)
			forceCreate = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for forceCreate, must be BOOL")
			}
		}
	}
	var jwt string
	{
		jwt = serviceServiceUpdateJWT
	}
	v := &service.ServiceDefinitionT{
		ControllerSchema: body.ControllerSchema,
		Controller:       body.Controller,
		Policy:           body.Policy,
		ID:               body.ID,
		Name:             body.Name,
		Description:      body.Description,
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
	} else {
		v.Parameters = []*service.ParameterDefT{}
	}
	res := &service.ServiceUpdatePayload{
		Service: v,
	}
	res.ID = &id
	res.ForceCreate = forceCreate
	res.JWT = jwt

	return res, nil
}

// BuildServiceDeletePayload builds the payload for the service service-delete
// endpoint from CLI flags.
func BuildServiceDeletePayload(serviceServiceDeleteID string, serviceServiceDeleteJWT string) (*service.ServiceDeletePayload, error) {
	var id string
	{
		id = serviceServiceDeleteID
	}
	var jwt string
	{
		jwt = serviceServiceDeleteJWT
	}
	v := &service.ServiceDeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildJobListPayload builds the payload for the service job-list endpoint
// from CLI flags.
func BuildJobListPayload(serviceJobListServiceID string, serviceJobListLimit string, serviceJobListPage string, serviceJobListFilter string, serviceJobListOrderBy string, serviceJobListOrderDesc string, serviceJobListAtTime string, serviceJobListJWT string) (*service.JobListPayload, error) {
	var err error
	var serviceID string
	{
		serviceID = serviceJobListServiceID
	}
	var limit int
	{
		if serviceJobListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(serviceJobListLimit, 10, strconv.IntSize)
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
		if serviceJobListPage != "" {
			page = &serviceJobListPage
		}
	}
	var filter *string
	{
		if serviceJobListFilter != "" {
			filter = &serviceJobListFilter
		}
	}
	var orderBy *string
	{
		if serviceJobListOrderBy != "" {
			orderBy = &serviceJobListOrderBy
		}
	}
	var orderDesc bool
	{
		if serviceJobListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(serviceJobListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if serviceJobListAtTime != "" {
			atTime = &serviceJobListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = serviceJobListJWT
	}
	v := &service.JobListPayload{}
	v.ServiceID = serviceID
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildJobCreatePayload builds the payload for the service job-create endpoint
// from CLI flags.
func BuildJobCreatePayload(serviceJobCreateServiceID string, serviceJobCreateJWT string, serviceJobCreateInContentType string, serviceJobCreateInOrderID string, serviceJobCreateForwardHost string, serviceJobCreateForwardProto string, serviceJobCreateTimeout string) (*service.JobCreatePayload, error) {
	var err error
	var serviceID string
	{
		serviceID = serviceJobCreateServiceID
	}
	var jwt string
	{
		jwt = serviceJobCreateJWT
	}
	var inContentType string
	{
		inContentType = serviceJobCreateInContentType
	}
	var inOrderID *string
	{
		if serviceJobCreateInOrderID != "" {
			inOrderID = &serviceJobCreateInOrderID
		}
	}
	var forwardHost *string
	{
		if serviceJobCreateForwardHost != "" {
			forwardHost = &serviceJobCreateForwardHost
		}
	}
	var forwardProto *string
	{
		if serviceJobCreateForwardProto != "" {
			forwardProto = &serviceJobCreateForwardProto
		}
	}
	var timeout *int
	{
		if serviceJobCreateTimeout != "" {
			var v int64
			v, err = strconv.ParseInt(serviceJobCreateTimeout, 10, strconv.IntSize)
			val := int(v)
			timeout = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for timeout, must be INT")
			}
		}
	}
	v := &service.JobCreatePayload{}
	v.ServiceID = serviceID
	v.JWT = jwt
	v.InContentType = inContentType
	v.InOrderID = inOrderID
	v.ForwardHost = forwardHost
	v.ForwardProto = forwardProto
	v.Timeout = timeout

	return v, nil
}

// BuildJobReadPayload builds the payload for the service job-read endpoint
// from CLI flags.
func BuildJobReadPayload(serviceJobReadServiceID string, serviceJobReadID string, serviceJobReadWithRequestContent string, serviceJobReadWithResultContent string, serviceJobReadJWT string) (*service.JobReadPayload, error) {
	var err error
	var serviceID string
	{
		serviceID = serviceJobReadServiceID
	}
	var id string
	{
		id = serviceJobReadID
	}
	var withRequestContent *bool
	{
		if serviceJobReadWithRequestContent != "" {
			var val bool
			val, err = strconv.ParseBool(serviceJobReadWithRequestContent)
			withRequestContent = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for withRequestContent, must be BOOL")
			}
		}
	}
	var withResultContent *bool
	{
		if serviceJobReadWithResultContent != "" {
			var val bool
			val, err = strconv.ParseBool(serviceJobReadWithResultContent)
			withResultContent = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for withResultContent, must be BOOL")
			}
		}
	}
	var jwt string
	{
		jwt = serviceJobReadJWT
	}
	v := &service.JobReadPayload{}
	v.ServiceID = serviceID
	v.ID = id
	v.WithRequestContent = withRequestContent
	v.WithResultContent = withResultContent
	v.JWT = jwt

	return v, nil
}
