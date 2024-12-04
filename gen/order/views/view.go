// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// OrderListRT is the viewed result type that is projected based on a view.
type OrderListRT struct {
	// Type to project
	Projected *OrderListRTView
	// View to render
	View string
}

// OrderTopResultItemCollection is the viewed result type that is projected
// based on a view.
type OrderTopResultItemCollection struct {
	// Type to project
	Projected OrderTopResultItemCollectionView
	// View to render
	View string
}

// OrderListRTView is a type that runs validations on a projected type.
type OrderListRTView struct {
	// Orders
	Items []*OrderListItemView
	// Time at which this list was valid
	AtTime *string
	Links  []*LinkTView
}

// OrderListItemView is a type that runs validations on a projected type.
type OrderListItemView struct {
	// ID
	ID *string
	// Optional customer provided name
	Name *string
	// Order status
	Status *string
	// DateTime order was placed
	OrderedAt *string
	// DateTime order processing started
	StartedAt *string
	// DateTime order processing finished
	FinishedAt *string
	// Reference to service requested
	Service *string
	// Reference to billable account
	Account *string
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

// OrderTopResultItemCollectionView is a type that runs validations on a
// projected type.
type OrderTopResultItemCollectionView []*OrderTopResultItemView

// OrderTopResultItemView is a type that runs validations on a projected type.
type OrderTopResultItemView struct {
	// container
	Container *string
	// cpu
	CPU *string
	// memory
	Memory *string
	// storage
	Storage *string
	// ephemeral-storage
	EphemeralStorage *string
}

var (
	// OrderListRTMap is a map indexing the attribute names of OrderListRT by view
	// name.
	OrderListRTMap = map[string][]string{
		"default": {
			"items",
			"at-time",
			"links",
		},
	}
	// OrderTopResultItemCollectionMap is a map indexing the attribute names of
	// OrderTopResultItemCollection by view name.
	OrderTopResultItemCollectionMap = map[string][]string{
		"default": {
			"container",
			"cpu",
			"memory",
			"storage",
			"ephemeral-storage",
		},
	}
	// OrderTopResultItemMap is a map indexing the attribute names of
	// OrderTopResultItem by view name.
	OrderTopResultItemMap = map[string][]string{
		"default": {
			"container",
			"cpu",
			"memory",
			"storage",
			"ephemeral-storage",
		},
	}
)

// ValidateOrderListRT runs the validations defined on the viewed result type
// OrderListRT.
func ValidateOrderListRT(result *OrderListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateOrderListRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateOrderTopResultItemCollection runs the validations defined on the
// viewed result type OrderTopResultItemCollection.
func ValidateOrderTopResultItemCollection(result OrderTopResultItemCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateOrderTopResultItemCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateOrderListRTView runs the validations defined on OrderListRTView
// using the "default" view.
func ValidateOrderListRTView(result *OrderListRTView) (err error) {
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
			if err2 := ValidateOrderListItemView(e); err2 != nil {
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

// ValidateOrderListItemView runs the validations defined on OrderListItemView.
func ValidateOrderListItemView(result *OrderListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "result"))
	}
	if result.Account == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("account", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatUUID))
	}
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "pending" || *result.Status == "scheduled" || *result.Status == "executing" || *result.Status == "succeeded" || *result.Status == "failed" || *result.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if result.OrderedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.ordered-at", *result.OrderedAt, goa.FormatDateTime))
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
	if result.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.account", *result.Account, goa.FormatURI))
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

// ValidateOrderTopResultItemCollectionView runs the validations defined on
// OrderTopResultItemCollectionView using the "default" view.
func ValidateOrderTopResultItemCollectionView(result OrderTopResultItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateOrderTopResultItemView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateOrderTopResultItemView runs the validations defined on
// OrderTopResultItemView using the "default" view.
func ValidateOrderTopResultItemView(result *OrderTopResultItemView) (err error) {
	if result.Container == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("container", "result"))
	}
	if result.CPU == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cpu", "result"))
	}
	if result.Memory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("memory", "result"))
	}
	if result.Storage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("storage", "result"))
	}
	if result.EphemeralStorage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ephemeral-storage", "result"))
	}
	return
}
