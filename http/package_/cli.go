// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	"fmt"
	"strconv"

	package_ "github.com/ivcap-works/ivcap-core-api/gen/package_"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the package list endpoint from CLI
// flags.
func BuildListPayload(package_ListTag string, package_ListPage string, package_ListLimit string, package_ListJWT string) (*package_.ListPayload, error) {
	var err error
	var tag *string
	{
		if package_ListTag != "" {
			tag = &package_ListTag
		}
	}
	var page *string
	{
		if package_ListPage != "" {
			page = &package_ListPage
		}
	}
	var limit *int
	{
		if package_ListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(package_ListLimit, 10, strconv.IntSize)
			val := int(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
		}
	}
	var jwt string
	{
		jwt = package_ListJWT
	}
	v := &package_.ListPayload{}
	v.Tag = tag
	v.Page = page
	v.Limit = limit
	v.JWT = jwt

	return v, nil
}

// BuildPullPayload builds the payload for the package pull endpoint from CLI
// flags.
func BuildPullPayload(package_PullRef string, package_PullType string, package_PullOffset string, package_PullJWT string) (*package_.PullPayload, error) {
	var err error
	var ref string
	{
		ref = package_PullRef
	}
	var type_ string
	{
		type_ = package_PullType
		if !(type_ == "manifest" || type_ == "config" || type_ == "layer") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("type", type_, []any{"manifest", "config", "layer"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var offset *int
	{
		if package_PullOffset != "" {
			var v int64
			v, err = strconv.ParseInt(package_PullOffset, 10, strconv.IntSize)
			val := int(v)
			offset = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
		}
	}
	var jwt string
	{
		jwt = package_PullJWT
	}
	v := &package_.PullPayload{}
	v.Ref = ref
	v.Type = type_
	v.Offset = offset
	v.JWT = jwt

	return v, nil
}

// BuildPushPayload builds the payload for the package push endpoint from CLI
// flags.
func BuildPushPayload(package_PushTag string, package_PushForce string, package_PushType string, package_PushDigest string, package_PushJWT string) (*package_.PushPayload, error) {
	var err error
	var tag string
	{
		tag = package_PushTag
	}
	var force *bool
	{
		if package_PushForce != "" {
			var val bool
			val, err = strconv.ParseBool(package_PushForce)
			force = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for force, must be BOOL")
			}
		}
	}
	var type_ string
	{
		type_ = package_PushType
		if !(type_ == "manifest" || type_ == "config" || type_ == "layer") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("type", type_, []any{"manifest", "config", "layer"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var digest string
	{
		digest = package_PushDigest
	}
	var jwt string
	{
		jwt = package_PushJWT
	}
	v := &package_.PushPayload{}
	v.Tag = tag
	v.Force = force
	v.Type = type_
	v.Digest = digest
	v.JWT = jwt

	return v, nil
}

// BuildPatchPayload builds the payload for the package patch endpoint from CLI
// flags.
func BuildPatchPayload(package_PatchTag string, package_PatchDigest string, package_PatchTotal string, package_PatchStart string, package_PatchEnd string, package_PatchLocation string, package_PatchJWT string) (*package_.PatchPayload, error) {
	var err error
	var tag string
	{
		tag = package_PatchTag
	}
	var digest string
	{
		digest = package_PatchDigest
	}
	var total int
	{
		var v int64
		v, err = strconv.ParseInt(package_PatchTotal, 10, strconv.IntSize)
		total = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for total, must be INT")
		}
	}
	var start int
	{
		var v int64
		v, err = strconv.ParseInt(package_PatchStart, 10, strconv.IntSize)
		start = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for start, must be INT")
		}
	}
	var end int
	{
		var v int64
		v, err = strconv.ParseInt(package_PatchEnd, 10, strconv.IntSize)
		end = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for end, must be INT")
		}
	}
	var location string
	{
		location = package_PatchLocation
	}
	var jwt string
	{
		jwt = package_PatchJWT
	}
	v := &package_.PatchPayload{}
	v.Tag = tag
	v.Digest = digest
	v.Total = total
	v.Start = start
	v.End = end
	v.Location = location
	v.JWT = jwt

	return v, nil
}

// BuildPutPayload builds the payload for the package put endpoint from CLI
// flags.
func BuildPutPayload(package_PutTag string, package_PutDigest string, package_PutLocation string, package_PutJWT string) (*package_.PutPayload, error) {
	var tag string
	{
		tag = package_PutTag
	}
	var digest string
	{
		digest = package_PutDigest
	}
	var location string
	{
		location = package_PutLocation
	}
	var jwt string
	{
		jwt = package_PutJWT
	}
	v := &package_.PutPayload{}
	v.Tag = tag
	v.Digest = digest
	v.Location = location
	v.JWT = jwt

	return v, nil
}

// BuildRemovePayload builds the payload for the package remove endpoint from
// CLI flags.
func BuildRemovePayload(package_RemoveTag string, package_RemoveJWT string) (*package_.RemovePayload, error) {
	var tag string
	{
		tag = package_RemoveTag
	}
	var jwt string
	{
		jwt = package_RemoveJWT
	}
	v := &package_.RemovePayload{}
	v.Tag = tag
	v.JWT = jwt

	return v, nil
}
