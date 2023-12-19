// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	artifact "github.com/reinventingscience/ivcap-core-api/gen/artifact"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the artifact list endpoint from CLI
// flags.
func BuildListPayload(artifactListLimit string, artifactListPage string, artifactListFilter string, artifactListOrderBy string, artifactListOrderDesc string, artifactListAtTime string, artifactListJWT string) (*artifact.ListPayload, error) {
	var err error
	var limit int
	{
		if artifactListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(artifactListLimit, 10, strconv.IntSize)
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
		if artifactListPage != "" {
			page = &artifactListPage
		}
	}
	var filter *string
	{
		if artifactListFilter != "" {
			filter = &artifactListFilter
		}
	}
	var orderBy *string
	{
		if artifactListOrderBy != "" {
			orderBy = &artifactListOrderBy
		}
	}
	var orderDesc bool
	{
		if artifactListOrderDesc != "" {
			orderDesc, err = strconv.ParseBool(artifactListOrderDesc)
			if err != nil {
				return nil, fmt.Errorf("invalid value for orderDesc, must be BOOL")
			}
		}
	}
	var atTime *string
	{
		if artifactListAtTime != "" {
			atTime = &artifactListAtTime
			err = goa.MergeErrors(err, goa.ValidateFormat("at-time", *atTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var jwt string
	{
		jwt = artifactListJWT
	}
	v := &artifact.ListPayload{}
	v.Limit = limit
	v.Page = page
	v.Filter = filter
	v.OrderBy = orderBy
	v.OrderDesc = orderDesc
	v.AtTime = atTime
	v.JWT = jwt

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

// BuildUploadPayload builds the payload for the artifact upload endpoint from
// CLI flags.
func BuildUploadPayload(artifactUploadJWT string, artifactUploadContentType string, artifactUploadContentEncoding string, artifactUploadContentLength string, artifactUploadName string, artifactUploadCollection string, artifactUploadPolicy string, artifactUploadXContentType string, artifactUploadXContentLength string, artifactUploadUploadLength string, artifactUploadTusResumable string) (*artifact.UploadPayload, error) {
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
	var policy *string
	{
		if artifactUploadPolicy != "" {
			policy = &artifactUploadPolicy
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
	v.Policy = policy
	v.XContentType = xContentType
	v.XContentLength = xContentLength
	v.UploadLength = uploadLength
	v.TusResumable = tusResumable

	return v, nil
}
