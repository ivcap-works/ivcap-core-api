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

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ListMetaRT is the viewed result type that is projected based on a view.
type ListMetaRT struct {
	// Type to project
	Projected *ListMetaRTView
	// View to render
	View string
}

// MetadataRecordRT is the viewed result type that is projected based on a view.
type MetadataRecordRT struct {
	// Type to project
	Projected *MetadataRecordRTView
	// View to render
	View string
}

// AddMetaRT is the viewed result type that is projected based on a view.
type AddMetaRT struct {
	// Type to project
	Projected *AddMetaRTView
	// View to render
	View string
}

// ListMetaRTView is a type that runs validations on a projected type.
type ListMetaRTView struct {
	// List of metadata records
	Records []*MetadataListItemRTView
	// Entity for which to request metadata
	EntityID *string
	// Optional schema to filter on
	Schema *string
	// Optional json path to further filter on returned list
	AspectPath *string
	// Time at which this list was valid
	AtTime *string
	// Navigation links
	Links *NavTView
}

// MetadataListItemRTView is a type that runs validations on a projected type.
type MetadataListItemRTView struct {
	// Record ID
	RecordID *string
	// Entity ID
	Entity *string
	// Schema ID
	Schema *string
	// Attached metadata aspect
	Aspect interface{}
	// If aspectPath was defined, this is what matched the query
	AspectContext interface{}
}

// NavTView is a type that runs validations on a projected type.
type NavTView struct {
	Self  *string
	First *string
	Next  *string
}

// MetadataRecordRTView is a type that runs validations on a projected type.
type MetadataRecordRTView struct {
	// Record ID
	RecordID *string
	// Entity ID
	Entity *string
	// Schema ID
	Schema *string
	// Attached metadata aspect
	Aspect interface{}
	// Time this record was asserted
	ValidFrom *string
	// Time this record was revoked
	ValidTo *string
	// Entity asserting this metadata record at 'valid-from'
	Asserter *string
	// Entity revoking this record at 'valid-to'
	Revoker *string
}

// AddMetaRTView is a type that runs validations on a projected type.
type AddMetaRTView struct {
	// Reference to record created
	RecordID *string
}

var (
	// ListMetaRTMap is a map indexing the attribute names of ListMetaRT by view
	// name.
	ListMetaRTMap = map[string][]string{
		"default": {
			"records",
			"entity-id",
			"schema",
			"aspect-path",
			"at-time",
			"links",
		},
	}
	// MetadataRecordRTMap is a map indexing the attribute names of
	// MetadataRecordRT by view name.
	MetadataRecordRTMap = map[string][]string{
		"default": {
			"record-id",
			"entity",
			"schema",
			"aspect",
			"valid-from",
			"valid-to",
			"asserter",
			"revoker",
		},
	}
	// AddMetaRTMap is a map indexing the attribute names of AddMetaRT by view name.
	AddMetaRTMap = map[string][]string{
		"default": {
			"record-id",
		},
	}
	// MetadataListItemRTMap is a map indexing the attribute names of
	// MetadataListItemRT by view name.
	MetadataListItemRTMap = map[string][]string{
		"default": {
			"record-id",
			"entity",
			"schema",
			"aspect",
			"aspectContext",
		},
	}
)

// ValidateListMetaRT runs the validations defined on the viewed result type
// ListMetaRT.
func ValidateListMetaRT(result *ListMetaRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateListMetaRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateMetadataRecordRT runs the validations defined on the viewed result
// type MetadataRecordRT.
func ValidateMetadataRecordRT(result *MetadataRecordRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateMetadataRecordRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateAddMetaRT runs the validations defined on the viewed result type
// AddMetaRT.
func ValidateAddMetaRT(result *AddMetaRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateAddMetaRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateListMetaRTView runs the validations defined on ListMetaRTView using
// the "default" view.
func ValidateListMetaRTView(result *ListMetaRTView) (err error) {
	if result.Records == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("records", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	for _, e := range result.Records {
		if e != nil {
			if err2 := ValidateMetadataListItemRTView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.EntityID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.entity-id", *result.EntityID, goa.FormatURI))
	}
	if result.Schema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.schema", *result.Schema, goa.FormatURI))
	}
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	if result.Links != nil {
		if err2 := ValidateNavTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateMetadataListItemRTView runs the validations defined on
// MetadataListItemRTView using the "default" view.
func ValidateMetadataListItemRTView(result *MetadataListItemRTView) (err error) {
	if result.RecordID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.record-id", *result.RecordID, goa.FormatURI))
	}
	if result.Entity != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.entity", *result.Entity, goa.FormatURI))
	}
	if result.Schema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.schema", *result.Schema, goa.FormatURI))
	}
	return
}

// ValidateNavTView runs the validations defined on NavTView.
func ValidateNavTView(result *NavTView) (err error) {
	if result.Self != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.self", *result.Self, goa.FormatURI))
	}
	if result.First != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.first", *result.First, goa.FormatURI))
	}
	if result.Next != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.next", *result.Next, goa.FormatURI))
	}
	return
}

// ValidateMetadataRecordRTView runs the validations defined on
// MetadataRecordRTView using the "default" view.
func ValidateMetadataRecordRTView(result *MetadataRecordRTView) (err error) {
	if result.RecordID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.record-id", *result.RecordID, goa.FormatURI))
	}
	if result.Entity != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.entity", *result.Entity, goa.FormatURI))
	}
	if result.Schema != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.schema", *result.Schema, goa.FormatURI))
	}
	if result.ValidFrom != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.valid-from", *result.ValidFrom, goa.FormatDateTime))
	}
	if result.ValidTo != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.valid-to", *result.ValidTo, goa.FormatDateTime))
	}
	if result.Revoker != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.revoker", *result.Revoker, goa.FormatDateTime))
	}
	return
}

// ValidateAddMetaRTView runs the validations defined on AddMetaRTView using
// the "default" view.
func ValidateAddMetaRTView(result *AddMetaRTView) (err error) {
	if result.RecordID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("record-id", "result"))
	}
	if result.RecordID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.record-id", *result.RecordID, goa.FormatURI))
	}
	return
}
