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
	metadata "github.com/reinventingscience/ivcap-core-api/gen/metadata"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the metadata list endpoint from CLI
// flags.
func BuildListPayload(metadataListEntityID string, metadataListSchema string, metadataListAspectPath string, metadataListAtTime string, metadataListLimit string, metadataListFilter string, metadataListOrderBy string, metadataListOrderDesc string, metadataListPage string, metadataListJWT string) (*metadata.ListPayload, error) {
	var err error
	var entityID *string
	{
		if metadataListEntityID != "" {
			entityID = &metadataListEntityID
			err = goa.MergeErrors(err, goa.ValidateFormat("entityID", *entityID, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var schema *string
	{
		if metadataListSchema != "" {
			schema = &metadataListSchema
		}
	}
	var aspectPath *string
	{
		if metadataListAspectPath != "" {
			aspectPath = &metadataListAspectPath
		}
	}
	var atTime *string
	{
		if metadataListAtTime != "" {
			atTime = &metadataListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("atTime", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if metadataListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(metadataListLimit, 10, strconv.IntSize)
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
		if metadataListFilter != "" {
			filter = metadataListFilter
		}
	}
	var orderBy string
	{
		if metadataListOrderBy != "" {
			orderBy = metadataListOrderBy
		}
	}
	var orderDesc *bool
	{
		if metadataListOrderDesc != "" {
			var val bool
			val, err = strconv.ParseBool(metadataListOrderDesc)
			orderDesc = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var page *string
	{
		if metadataListPage != "" {
			page = &metadataListPage
		}
	}
	var jwt string
	{
		jwt = metadataListJWT
	}
	v := &metadata.ListPayload{}
	v.EntityID = entityID
	v.Schema = schema
	v.AspectPath = aspectPath
	v.AtTime = atTime
	v.Limit = limit
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.Page = page
	v.JWT = jwt

	return v, nil
}

// BuildReadPayload builds the payload for the metadata read endpoint from CLI
// flags.
func BuildReadPayload(metadataReadID string, metadataReadJWT string) (*metadata.ReadPayload, error) {
	var id string
	{
		id = metadataReadID
	}
	var jwt string
	{
		jwt = metadataReadJWT
	}
	v := &metadata.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildAddPayload builds the payload for the metadata add endpoint from CLI
// flags.
func BuildAddPayload(metadataAddBody string, metadataAddEntityID string, metadataAddSchema string, metadataAddPolicyID string, metadataAddJWT string, metadataAddContentType string) (*metadata.AddPayload, error) {
	var err error
	var body interface{}
	{
		err = json.Unmarshal([]byte(metadataAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "\"{\\\"$schema\\\": ...}\"")
		}
	}
	var entityID string
	{
		entityID = metadataAddEntityID
		err = goa.MergeErrors(err, goa.ValidateFormat("entityID", entityID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var schema string
	{
		schema = metadataAddSchema
		err = goa.MergeErrors(err, goa.ValidateFormat("schema", schema, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var policyID *string
	{
		if metadataAddPolicyID != "" {
			policyID = &metadataAddPolicyID
			err = goa.MergeErrors(err, goa.ValidateFormat("policyID", *policyID, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = metadataAddJWT
	}
	var contentType string
	{
		contentType = metadataAddContentType
	}
	v := body
	res := &metadata.AddPayload{
		Aspect: &v,
	}
	res.EntityID = entityID
	res.Schema = schema
	res.PolicyID = policyID
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}

// BuildUpdateOnePayload builds the payload for the metadata update_one
// endpoint from CLI flags.
func BuildUpdateOnePayload(metadataUpdateOneBody string, metadataUpdateOneEntityID string, metadataUpdateOneSchema string, metadataUpdateOnePolicyID string, metadataUpdateOneJWT string, metadataUpdateOneContentType string) (*metadata.UpdateOnePayload, error) {
	var err error
	var body string
	{
		body = metadataUpdateOneBody
	}
	var entityID string
	{
		entityID = metadataUpdateOneEntityID
		err = goa.MergeErrors(err, goa.ValidateFormat("entityID", entityID, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var schema string
	{
		schema = metadataUpdateOneSchema
		err = goa.MergeErrors(err, goa.ValidateFormat("schema", schema, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var policyID *string
	{
		if metadataUpdateOnePolicyID != "" {
			policyID = &metadataUpdateOnePolicyID
			err = goa.MergeErrors(err, goa.ValidateFormat("policyID", *policyID, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = metadataUpdateOneJWT
	}
	var contentType *string
	{
		if metadataUpdateOneContentType != "" {
			contentType = &metadataUpdateOneContentType
		}
	}
	v := body
	res := &metadata.UpdateOnePayload{
		Aspect: &v,
	}
	res.EntityID = entityID
	res.Schema = schema
	res.PolicyID = policyID
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}

// BuildUpdateRecordPayload builds the payload for the metadata update_record
// endpoint from CLI flags.
func BuildUpdateRecordPayload(metadataUpdateRecordBody string, metadataUpdateRecordID string, metadataUpdateRecordEntityID string, metadataUpdateRecordSchema string, metadataUpdateRecordPolicyID string, metadataUpdateRecordJWT string, metadataUpdateRecordContentType string) (*metadata.UpdateRecordPayload, error) {
	var err error
	var body string
	{
		body = metadataUpdateRecordBody
	}
	var id string
	{
		id = metadataUpdateRecordID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var entityID *string
	{
		if metadataUpdateRecordEntityID != "" {
			entityID = &metadataUpdateRecordEntityID
			err = goa.MergeErrors(err, goa.ValidateFormat("entityID", *entityID, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var schema *string
	{
		if metadataUpdateRecordSchema != "" {
			schema = &metadataUpdateRecordSchema
			err = goa.MergeErrors(err, goa.ValidateFormat("schema", *schema, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var policyID *string
	{
		if metadataUpdateRecordPolicyID != "" {
			policyID = &metadataUpdateRecordPolicyID
			err = goa.MergeErrors(err, goa.ValidateFormat("policyID", *policyID, goa.FormatURI))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = metadataUpdateRecordJWT
	}
	var contentType *string
	{
		if metadataUpdateRecordContentType != "" {
			contentType = &metadataUpdateRecordContentType
		}
	}
	v := body
	res := &metadata.UpdateRecordPayload{
		Aspect: v,
	}
	res.ID = &id
	res.EntityID = entityID
	res.Schema = schema
	res.PolicyID = policyID
	res.JWT = jwt
	res.ContentType = contentType

	return res, nil
}

// BuildRevokePayload builds the payload for the metadata revoke endpoint from
// CLI flags.
func BuildRevokePayload(metadataRevokeID string, metadataRevokeJWT string) (*metadata.RevokePayload, error) {
	var err error
	var id string
	{
		id = metadataRevokeID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatURI))
		if err != nil {
			return nil, err
		}
	}
	var jwt string
	{
		jwt = metadataRevokeJWT
	}
	v := &metadata.RevokePayload{}
	v.ID = &id
	v.JWT = jwt

	return v, nil
}
