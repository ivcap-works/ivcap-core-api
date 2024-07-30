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
func BuildPushPayload(package_PushTag string, package_PushForce string, package_PushType string, package_PushDigest string, package_PushTotal string, package_PushStart string, package_PushEnd string, package_PushJWT string) (*package_.PushPayload, error) {
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
	var total *int
	{
		if package_PushTotal != "" {
			var v int64
			v, err = strconv.ParseInt(package_PushTotal, 10, strconv.IntSize)
			val := int(v)
			total = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for total, must be INT")
			}
		}
	}
	var start *int
	{
		if package_PushStart != "" {
			var v int64
			v, err = strconv.ParseInt(package_PushStart, 10, strconv.IntSize)
			val := int(v)
			start = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for start, must be INT")
			}
		}
	}
	var end *int
	{
		if package_PushEnd != "" {
			var v int64
			v, err = strconv.ParseInt(package_PushEnd, 10, strconv.IntSize)
			val := int(v)
			end = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for end, must be INT")
			}
		}
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
	v.Total = total
	v.Start = start
	v.End = end
	v.JWT = jwt

	return v, nil
}

// BuildStatusPayload builds the payload for the package status endpoint from
// CLI flags.
func BuildStatusPayload(package_StatusTag string, package_StatusDigest string, package_StatusJWT string) (*package_.StatusPayload, error) {
	var tag string
	{
		tag = package_StatusTag
	}
	var digest string
	{
		digest = package_StatusDigest
	}
	var jwt string
	{
		jwt = package_StatusJWT
	}
	v := &package_.StatusPayload{}
	v.Tag = tag
	v.Digest = digest
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
