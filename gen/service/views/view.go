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

// ServiceListRT is the viewed result type that is projected based on a view.
type ServiceListRT struct {
	// Type to project
	Projected *ServiceListRTView
	// View to render
	View string
}

// JobListRT is the viewed result type that is projected based on a view.
type JobListRT struct {
	// Type to project
	Projected *JobListRTView
	// View to render
	View string
}

// ServiceListRTView is a type that runs validations on a projected type.
type ServiceListRTView struct {
	// Services
	Items []*ServiceListItemTView
	// Time at which this list was valid
	AtTime *string
	Links  []*LinkTView
}

// ServiceListItemTView is a type that runs validations on a projected type.
type ServiceListItemTView struct {
	// ID
	ID *string
	// Optional customer provided name
	Name *string
	// Optional description of the service
	Description *string
	// Optional tags defined for service to help in categorising them
	Tags []string
	// type of controller used for this service
	ControllerSchema *string
	// time this service has been available from
	ValidFrom *string
	// time this service has been available to
	ValidTo *string
	Href    *string `json:"href,omitempty"`
}

// LinkTView is a type that runs validations on a projected type.
type LinkTView struct {
	// relation type
	Rel *string
	// mime type
	Type *string
	// web link
	Href *string
}

// JobListRTView is a type that runs validations on a projected type.
type JobListRTView struct {
	// Jobs
	Items []*JobListItemView
	// Time at which this list was valid
	AtTime *string
	Links  []*LinkTView
}

// JobListItemView is a type that runs validations on a projected type.
type JobListItemView struct {
	// ID
	ID *string
	// Optional customer provided name
	Name *string
	// Job status
	Status *string
	// DateTime job processing started
	StartedAt *string
	// DateTime job processing finished
	FinishedAt *string
	// Reference to service requested
	Service *string
	// Reference to order
	Order *string
	Href  *string `json:"href,omitempty"`
}

var (
	// ServiceListRTMap is a map indexing the attribute names of ServiceListRT by
	// view name.
	ServiceListRTMap = map[string][]string{
		"default": {
			"items",
			"at-time",
			"links",
		},
	}
	// JobListRTMap is a map indexing the attribute names of JobListRT by view name.
	JobListRTMap = map[string][]string{
		"default": {
			"items",
			"at-time",
			"links",
		},
	}
)

// ValidateServiceListRT runs the validations defined on the viewed result type
// ServiceListRT.
func ValidateServiceListRT(result *ServiceListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateServiceListRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateJobListRT runs the validations defined on the viewed result type
// JobListRT.
func ValidateJobListRT(result *JobListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJobListRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateServiceListRTView runs the validations defined on ServiceListRTView
// using the "default" view.
func ValidateServiceListRTView(result *ServiceListRTView) (err error) {
	if result.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "result"))
	}
	if result.AtTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("at-time", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	for _, e := range result.Items {
		if e != nil {
			if err2 := ValidateServiceListItemTView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	for _, e := range result.Links {
		if e != nil {
			if err2 := ValidateLinkTView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateServiceListItemTView runs the validations defined on
// ServiceListItemTView.
func ValidateServiceListItemTView(result *ServiceListItemTView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.ControllerSchema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("controller-schema", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	if result.ValidFrom != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.valid-from", *result.ValidFrom, goa.FormatDateTime))
	}
	if result.ValidTo != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.valid-to", *result.ValidTo, goa.FormatDateTime))
	}
	return
}

// ValidateLinkTView runs the validations defined on LinkTView.
func ValidateLinkTView(result *LinkTView) (err error) {
	if result.Rel == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rel", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	return
}

// ValidateJobListRTView runs the validations defined on JobListRTView using
// the "default" view.
func ValidateJobListRTView(result *JobListRTView) (err error) {
	if result.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "result"))
	}
	if result.AtTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("at-time", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	for _, e := range result.Items {
		if e != nil {
			if err2 := ValidateJobListItemView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	for _, e := range result.Links {
		if e != nil {
			if err2 := ValidateLinkTView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateJobListItemView runs the validations defined on JobListItemView.
func ValidateJobListItemView(result *JobListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "pending" || *result.Status == "scheduled" || *result.Status == "executing" || *result.Status == "succeeded" || *result.Status == "failed" || *result.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started-at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.finished-at", *result.FinishedAt, goa.FormatDateTime))
	}
	if result.Service != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.service", *result.Service, goa.FormatURI))
	}
	if result.Order != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.order", *result.Order, goa.FormatURI))
	}
	return
}
