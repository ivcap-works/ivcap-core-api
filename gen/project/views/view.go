// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ProjectListRT is the viewed result type that is projected based on a view.
type ProjectListRT struct {
	// Type to project
	Projected *ProjectListRTView
	// View to render
	View string
}

// ProjectStatusRT is the viewed result type that is projected based on a view.
type ProjectStatusRT struct {
	// Type to project
	Projected *ProjectStatusRTView
	// View to render
	View string
}

// ProjectListRTView is a type that runs validations on a projected type.
type ProjectListRTView struct {
	// Projects
	Projects ProjectListItemCollectionView
	// Time at which this list was valid
	AtTime *string
	// A pagination token to retrieve the next set of results. Empty if there are
	// no more results
	Page *string
}

// ProjectListItemCollectionView is a type that runs validations on a projected
// type.
type ProjectListItemCollectionView []*ProjectListItemView

// ProjectListItemView is a type that runs validations on a projected type.
type ProjectListItemView struct {
	// Project Name
	Name *string
	// User Role
	Role *string
	// Project URN
	Urn *string
	// DateTime project was created
	CreatedAt *string
	// DateTime project last modified
	ModifiedAt *string
	// Time at which this list was valid
	AtTime *string
}

// ProjectStatusRTView is a type that runs validations on a projected type.
type ProjectStatusRTView struct {
	// Project status
	Status *string
	// DateTime project was created
	CreatedAt *string
	// DateTime project last modified
	ModifiedAt *string
	// Account URN
	Account *string
	// Parent Project URN
	Parent *string
	// Additional Metadata
	Properties *ProjectPropertiesView
	// Project URN
	Urn *string
	// Project name
	Name *string
}

// ProjectPropertiesView is a type that runs validations on a projected type.
type ProjectPropertiesView struct {
	// String metadata for detailing the use of this project
	Details *string
}

var (
	// ProjectListRTMap is a map indexing the attribute names of ProjectListRT by
	// view name.
	ProjectListRTMap = map[string][]string{
		"default": {
			"projects",
			"at-time",
			"page",
		},
		"tiny": {
			"projects",
			"at-time",
			"page",
		},
	}
	// ProjectStatusRTMap is a map indexing the attribute names of ProjectStatusRT
	// by view name.
	ProjectStatusRTMap = map[string][]string{
		"default": {
			"urn",
			"name",
			"account",
			"parent",
			"status",
			"created_at",
			"modified_at",
			"properties",
		},
		"tiny": {
			"name",
			"status",
		},
	}
	// ProjectListItemCollectionMap is a map indexing the attribute names of
	// ProjectListItemCollection by view name.
	ProjectListItemCollectionMap = map[string][]string{
		"default": {
			"name",
			"role",
			"urn",
			"created_at",
			"modified_at",
			"at-time",
		},
		"tiny": {
			"urn",
		},
	}
	// ProjectListItemMap is a map indexing the attribute names of ProjectListItem
	// by view name.
	ProjectListItemMap = map[string][]string{
		"default": {
			"name",
			"role",
			"urn",
			"created_at",
			"modified_at",
			"at-time",
		},
		"tiny": {
			"urn",
		},
	}
)

// ValidateProjectListRT runs the validations defined on the viewed result type
// ProjectListRT.
func ValidateProjectListRT(result *ProjectListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProjectListRTView(result.Projected)
	case "tiny":
		err = ValidateProjectListRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default", "tiny"})
	}
	return
}

// ValidateProjectStatusRT runs the validations defined on the viewed result
// type ProjectStatusRT.
func ValidateProjectStatusRT(result *ProjectStatusRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProjectStatusRTView(result.Projected)
	case "tiny":
		err = ValidateProjectStatusRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default", "tiny"})
	}
	return
}

// ValidateProjectListRTView runs the validations defined on ProjectListRTView
// using the "default" view.
func ValidateProjectListRTView(result *ProjectListRTView) (err error) {
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	if result.Projects != nil {
		if err2 := ValidateProjectListItemCollectionView(result.Projects); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectListRTViewTiny runs the validations defined on
// ProjectListRTView using the "tiny" view.
func ValidateProjectListRTViewTiny(result *ProjectListRTView) (err error) {
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	if result.Projects != nil {
		if err2 := ValidateProjectListItemCollectionViewTiny(result.Projects); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectListItemCollectionView runs the validations defined on
// ProjectListItemCollectionView using the "default" view.
func ValidateProjectListItemCollectionView(result ProjectListItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateProjectListItemView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectListItemCollectionViewTiny runs the validations defined on
// ProjectListItemCollectionView using the "tiny" view.
func ValidateProjectListItemCollectionViewTiny(result ProjectListItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateProjectListItemViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectListItemView runs the validations defined on
// ProjectListItemView using the "default" view.
func ValidateProjectListItemView(result *ProjectListItemView) (err error) {
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.ModifiedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.modified_at", *result.ModifiedAt, goa.FormatDateTime))
	}
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	return
}

// ValidateProjectListItemViewTiny runs the validations defined on
// ProjectListItemView using the "tiny" view.
func ValidateProjectListItemViewTiny(result *ProjectListItemView) (err error) {

	return
}

// ValidateProjectStatusRTView runs the validations defined on
// ProjectStatusRTView using the "default" view.
func ValidateProjectStatusRTView(result *ProjectStatusRTView) (err error) {
	if result.Urn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("urn", "result"))
	}
	if result.Urn != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.urn", *result.Urn, goa.FormatURI))
	}
	if result.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.account", *result.Account, goa.FormatURI))
	}
	if result.Parent != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.parent", *result.Parent, goa.FormatURI))
	}
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "active" || *result.Status == "disabled" || *result.Status == "deleted") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unknown", "active", "disabled", "deleted"}))
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.ModifiedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.modified_at", *result.ModifiedAt, goa.FormatDateTime))
	}
	return
}

// ValidateProjectStatusRTViewTiny runs the validations defined on
// ProjectStatusRTView using the "tiny" view.
func ValidateProjectStatusRTViewTiny(result *ProjectStatusRTView) (err error) {
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "active" || *result.Status == "disabled" || *result.Status == "deleted") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unknown", "active", "disabled", "deleted"}))
		}
	}
	return
}

// ValidateProjectPropertiesView runs the validations defined on
// ProjectPropertiesView.
func ValidateProjectPropertiesView(result *ProjectPropertiesView) (err error) {

	return
}
