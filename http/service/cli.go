// Copyright 2023 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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
func BuildListPayload(serviceListLimit string, serviceListPage string, serviceListFilter string, serviceListOrderBy string, serviceListOrderDesc string, serviceListAtTime string, serviceListJWT string) (*service.ListPayload, error) {
	var err error
	var limit int
	{
		if serviceListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(serviceListLimit, 10, strconv.IntSize)
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
		if serviceListPage != "" {
			page = &serviceListPage
		}
	}
	var filter *string
	{
		if serviceListFilter != "" {
			filter = &serviceListFilter
		}
	}
	var orderBy *string
	{
		if serviceListOrderBy != "" {
			orderBy = &serviceListOrderBy
		}
	}
	var orderDesc bool
	{
		if serviceListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(serviceListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if serviceListAtTime != "" {
			atTime = &serviceListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("atTime", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = serviceListJWT
	}
	v := &service.ListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

	return v, nil
}

// BuildCreateServicePayload builds the payload for the service create_service
// endpoint from CLI flags.
func BuildCreateServicePayload(serviceCreateServiceBody string, serviceCreateServiceJWT string) (*service.CreateServicePayload, error) {
	var err error
	var body CreateServiceRequestBody
	{
		err = json.Unmarshal([]byte(serviceCreateServiceBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"account-id\": \"urn:ivcap:account:0f0e3f57-80f7-4899-9b69-459af2efd789\",\n      \"banner\": \"http://boyer.org/otha\",\n      \"description\": \"This service ...\",\n      \"metadata\": [\n         {\n            \"name\": \"Porro minima qui ut cumque quia.\",\n            \"value\": \"Exercitationem ratione.\"\n         },\n         {\n            \"name\": \"Porro minima qui ut cumque quia.\",\n            \"value\": \"Exercitationem ratione.\"\n         },\n         {\n            \"name\": \"Porro minima qui ut cumque quia.\",\n            \"value\": \"Exercitationem ratione.\"\n         }\n      ],\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"policy-id\": \"Amet accusamus assumenda nobis ipsum.\",\n      \"provider-id\": \"urn:ivcap:provider:0f0e3f57-80f7-4899-9b69-459af2efd789\",\n      \"provider-ref\": \"service_foo_patch_1\",\n      \"references\": [\n         {\n            \"title\": \"Inventore voluptatem incidunt vel et voluptas.\",\n            \"uri\": \"http://rodriguez.name/sterling.schmitt\"\n         },\n         {\n            \"title\": \"Inventore voluptatem incidunt vel et voluptas.\",\n            \"uri\": \"http://rodriguez.name/sterling.schmitt\"\n         },\n         {\n            \"title\": \"Inventore voluptatem incidunt vel et voluptas.\",\n            \"uri\": \"http://rodriguez.name/sterling.schmitt\"\n         }\n      ],\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ],\n      \"workflow\": {\n         \"argo\": \"Exercitationem nobis perspiciatis voluptate quia.\",\n         \"basic\": {\n            \"command\": [\n               \"/bin/sh\",\n               \"-c\",\n               \"echo $PATH\"\n            ],\n            \"cpu\": {\n               \"limit\": \"100m\",\n               \"request\": \"10m\"\n            },\n            \"ephemeral-storage\": {\n               \"limit\": \"4Gi\",\n               \"request\": \"2Gi\"\n            },\n            \"image\": \"alpine\",\n            \"memory\": {\n               \"limit\": \"100Mi\",\n               \"request\": \"10Mi\"\n            }\n         },\n         \"opts\": \"Facilis alias.\",\n         \"type\": \"basic\"\n      }\n   }'")
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
		if body.Workflow != nil {
			if err2 := ValidateWorkflowTRequestBodyRequestBody(body.Workflow); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = serviceCreateServiceJWT
	}
	v := &service.ServiceDescriptionT{
		ProviderRef: body.ProviderRef,
		ProviderID:  body.ProviderID,
		Description: body.Description,
		AccountID:   body.AccountID,
		Banner:      body.Banner,
		PolicyID:    body.PolicyID,
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
	res := &service.CreateServicePayload{
		Services: v,
	}
	res.JWT = jwt

	return res, nil
}

// BuildReadPayload builds the payload for the service read endpoint from CLI
// flags.
func BuildReadPayload(serviceReadID string, serviceReadJWT string) (*service.ReadPayload, error) {
	var id string
	{
		id = serviceReadID
	}
	var jwt string
	{
		jwt = serviceReadJWT
	}
	v := &service.ReadPayload{}
	v.ID = id
	v.JWT = jwt

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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"account-id\": \"urn:ivcap:account:0f0e3f57-80f7-4899-9b69-459af2efd789\",\n      \"banner\": \"http://colegottlieb.biz/domenick_bins\",\n      \"description\": \"This service ...\",\n      \"metadata\": [\n         {\n            \"name\": \"Porro minima qui ut cumque quia.\",\n            \"value\": \"Exercitationem ratione.\"\n         },\n         {\n            \"name\": \"Porro minima qui ut cumque quia.\",\n            \"value\": \"Exercitationem ratione.\"\n         }\n      ],\n      \"name\": \"Fire risk for Lot2\",\n      \"parameters\": [\n         {\n            \"description\": \"The name of the region as according to ...\",\n            \"label\": \"Region Name\",\n            \"name\": \"region\",\n            \"type\": \"string\"\n         },\n         {\n            \"label\": \"Rainfall/month threshold\",\n            \"name\": \"threshold\",\n            \"type\": \"float\",\n            \"unit\": \"m\"\n         }\n      ],\n      \"policy-id\": \"Inventore inventore totam.\",\n      \"provider-id\": \"urn:ivcap:provider:0f0e3f57-80f7-4899-9b69-459af2efd789\",\n      \"provider-ref\": \"service_foo_patch_1\",\n      \"references\": [\n         {\n            \"title\": \"Inventore voluptatem incidunt vel et voluptas.\",\n            \"uri\": \"http://rodriguez.name/sterling.schmitt\"\n         },\n         {\n            \"title\": \"Inventore voluptatem incidunt vel et voluptas.\",\n            \"uri\": \"http://rodriguez.name/sterling.schmitt\"\n         },\n         {\n            \"title\": \"Inventore voluptatem incidunt vel et voluptas.\",\n            \"uri\": \"http://rodriguez.name/sterling.schmitt\"\n         }\n      ],\n      \"tags\": [\n         \"tag1\",\n         \"tag2\"\n      ],\n      \"workflow\": {\n         \"argo\": \"Exercitationem nobis perspiciatis voluptate quia.\",\n         \"basic\": {\n            \"command\": [\n               \"/bin/sh\",\n               \"-c\",\n               \"echo $PATH\"\n            ],\n            \"cpu\": {\n               \"limit\": \"100m\",\n               \"request\": \"10m\"\n            },\n            \"ephemeral-storage\": {\n               \"limit\": \"4Gi\",\n               \"request\": \"2Gi\"\n            },\n            \"image\": \"alpine\",\n            \"memory\": {\n               \"limit\": \"100Mi\",\n               \"request\": \"10Mi\"\n            }\n         },\n         \"opts\": \"Facilis alias.\",\n         \"type\": \"basic\"\n      }\n   }'")
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
		if body.Workflow != nil {
			if err2 := ValidateWorkflowTRequestBodyRequestBody(body.Workflow); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
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
		PolicyID:    body.PolicyID,
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
