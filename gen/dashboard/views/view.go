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

// DashboardListRT is the viewed result type that is projected based on a view.
type DashboardListRT struct {
	// Type to project
	Projected *DashboardListRTView
	// View to render
	View string
}

// DashboardListRTView is a type that runs validations on a projected type.
type DashboardListRTView struct {
	// Dashboards
	Items []*DashboardListItemView
}

// DashboardListItemView is a type that runs validations on a projected type.
type DashboardListItemView struct {
	// dashboard id
	ID *int
	// dashboard uid
	UID *string
	// Dashboard title
	Title *string
	// Dashboard url
	URL *string
}

var (
	// DashboardListRTMap is a map indexing the attribute names of DashboardListRT
	// by view name.
	DashboardListRTMap = map[string][]string{
		"default": {
			"items",
		},
	}
)

// ValidateDashboardListRT runs the validations defined on the viewed result
// type DashboardListRT.
func ValidateDashboardListRT(result *DashboardListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateDashboardListRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateDashboardListRTView runs the validations defined on
// DashboardListRTView using the "default" view.
func ValidateDashboardListRTView(result *DashboardListRTView) (err error) {
	if result.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "result"))
	}
	for _, e := range result.Items {
		if e != nil {
			if err2 := ValidateDashboardListItemView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateDashboardListItemView runs the validations defined on
// DashboardListItemView.
func ValidateDashboardListItemView(result *DashboardListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.UID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("uid", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	return
}
