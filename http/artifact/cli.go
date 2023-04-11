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
	artifact "github.com/reinventingscience/ivcap-core-api/gen/artifact"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the artifact list endpoint from CLI
// flags.
func BuildListPayload(artifactListFilter string, artifactListOrderby string, artifactListTop string, artifactListSkip string, artifactListSelect string, artifactListOffset string, artifactListLimit string, artifactListPageToken string, artifactListJWT string) (*artifact.ListPayload, error) {
	var err error
	var filter string
	{
		if artifactListFilter != "" {
			filter = artifactListFilter
		}
	}
	var orderby string
	{
		if artifactListOrderby != "" {
			orderby = artifactListOrderby
		}
	}
	var top int
	{
		if artifactListTop != "" {
			var v int64
			v, err = strconv.ParseInt(artifactListTop, 10, strconv.IntSize)
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
		if artifactListSkip != "" {
			var v int64
			v, err = strconv.ParseInt(artifactListSkip, 10, strconv.IntSize)
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
		if artifactListSelect != "" {
			select_ = artifactListSelect
		}
	}
	var offset *int
	{
		if artifactListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(artifactListOffset, 10, strconv.IntSize)
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
		if artifactListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(artifactListLimit, 10, strconv.IntSize)
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
		if artifactListPageToken != "" {
			pageToken = artifactListPageToken
		}
	}
	var jwt string
	{
		jwt = artifactListJWT
	}
	v := &artifact.ListPayload{}
	v.Filter = filter
	v.Orderby = orderby
	v.Top = top
	v.Skip = skip
	v.Select = select_
	v.Offset = offset
	v.Limit = limit
	v.PageToken = pageToken
	v.JWT = jwt

	return v, nil
}

// BuildUploadPayload builds the payload for the artifact upload endpoint from
// CLI flags.
func BuildUploadPayload(artifactUploadJWT string, artifactUploadContentType string, artifactUploadContentEncoding string, artifactUploadContentLength string, artifactUploadName string, artifactUploadCollection string, artifactUploadXContentType string, artifactUploadXContentLength string, artifactUploadUploadLength string, artifactUploadTusResumable string) (*artifact.UploadPayload, error) {
	var err error
	var jwt string
	{
		jwt = artifactUploadJWT
	}
	var contentType *string
	{
		if artifactUploadContentType != "" {
			contentType = &artifactUploadContentType
		}
	}
	var contentEncoding *string
	{
		if artifactUploadContentEncoding != "" {
			contentEncoding = &artifactUploadContentEncoding
		}
	}
	var contentLength *int
	{
		if artifactUploadContentLength != "" {
			var v int64
			v, err = strconv.ParseInt(artifactUploadContentLength, 10, strconv.IntSize)
			val := int(v)
			contentLength = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for contentLength, must be INT")
			}
		}
	}
	var name *string
	{
		if artifactUploadName != "" {
			name = &artifactUploadName
		}
	}
	var collection *string
	{
		if artifactUploadCollection != "" {
			collection = &artifactUploadCollection
		}
	}
	var xContentType *string
	{
		if artifactUploadXContentType != "" {
			xContentType = &artifactUploadXContentType
		}
	}
	var xContentLength *int
	{
		if artifactUploadXContentLength != "" {
			var v int64
			v, err = strconv.ParseInt(artifactUploadXContentLength, 10, strconv.IntSize)
			val := int(v)
			xContentLength = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for xContentLength, must be INT")
			}
		}
	}
	var uploadLength *int
	{
		if artifactUploadUploadLength != "" {
			var v int64
			v, err = strconv.ParseInt(artifactUploadUploadLength, 10, strconv.IntSize)
			val := int(v)
			uploadLength = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for uploadLength, must be INT")
			}
		}
	}
	var tusResumable *string
	{
		if artifactUploadTusResumable != "" {
			tusResumable = &artifactUploadTusResumable
		}
	}
	v := &artifact.UploadPayload{}
	v.JWT = jwt
	v.ContentType = contentType
	v.ContentEncoding = contentEncoding
	v.ContentLength = contentLength
	v.Name = name
	v.Collection = collection
	v.XContentType = xContentType
	v.XContentLength = xContentLength
	v.UploadLength = uploadLength
	v.TusResumable = tusResumable

	return v, nil
}

// BuildReadPayload builds the payload for the artifact read endpoint from CLI
// flags.
func BuildReadPayload(artifactReadID string, artifactReadJWT string) (*artifact.ReadPayload, error) {
	var id string
	{
		id = artifactReadID
	}
	var jwt string
	{
		jwt = artifactReadJWT
	}
	v := &artifact.ReadPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildAddCollectionPayload builds the payload for the artifact addCollection
// endpoint from CLI flags.
func BuildAddCollectionPayload(artifactAddCollectionID string, artifactAddCollectionName string, artifactAddCollectionJWT string) (*artifact.AddCollectionPayload, error) {
	var id string
	{
		id = artifactAddCollectionID
	}
	var name string
	{
		name = artifactAddCollectionName
	}
	var jwt string
	{
		jwt = artifactAddCollectionJWT
	}
	v := &artifact.AddCollectionPayload{}
	v.ID = id
	v.Name = name
	v.JWT = jwt

	return v, nil
}

// BuildRemoveCollectionPayload builds the payload for the artifact
// removeCollection endpoint from CLI flags.
func BuildRemoveCollectionPayload(artifactRemoveCollectionID string, artifactRemoveCollectionName string, artifactRemoveCollectionJWT string) (*artifact.RemoveCollectionPayload, error) {
	var id string
	{
		id = artifactRemoveCollectionID
	}
	var name string
	{
		name = artifactRemoveCollectionName
	}
	var jwt string
	{
		jwt = artifactRemoveCollectionJWT
	}
	v := &artifact.RemoveCollectionPayload{}
	v.ID = id
	v.Name = name
	v.JWT = jwt

	return v, nil
}

// BuildAddMetadataPayload builds the payload for the artifact addMetadata
// endpoint from CLI flags.
func BuildAddMetadataPayload(artifactAddMetadataBody string, artifactAddMetadataID string, artifactAddMetadataSchema string, artifactAddMetadataJWT string) (*artifact.AddMetadataPayload, error) {
	var err error
	var body interface{}
	{
		err = json.Unmarshal([]byte(artifactAddMetadataBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "\"Rerum aut.\"")
		}
	}
	var id string
	{
		id = artifactAddMetadataID
	}
	var schema string
	{
		schema = artifactAddMetadataSchema
	}
	var jwt string
	{
		jwt = artifactAddMetadataJWT
	}
	v := body
	res := &artifact.AddMetadataPayload{
		Meta: v,
	}
	res.ID = id
	res.Schema = schema
	res.JWT = jwt

	return res, nil
}
