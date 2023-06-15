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

// OrderListRT is the viewed result type that is projected based on a view.
type OrderListRT struct {
	// Type to project
	Projected *OrderListRTView
	// View to render
	View string
}

// OrderStatusRT is the viewed result type that is projected based on a view.
type OrderStatusRT struct {
	// Type to project
	Projected *OrderStatusRTView
	// View to render
	View string
}

// OrderListRTView is a type that runs validations on a projected type.
type OrderListRTView struct {
	// Orders
	Orders []*OrderListItemView
	// Time at which this list was valid
	AtTime *string
	// Navigation links
	Links *NavTView
}

// OrderListItemView is a type that runs validations on a projected type.
type OrderListItemView struct {
	// Order ID
	ID *string
	// Optional customer provided name
	Name *string
	// Order status
	Status *string
	// DateTime order was placed
	OrderedAt *string
	// DateTime processing of order started
	StartedAt *string
	// DateTime order was finished
	FinishedAt *string
	// ID of ordered service
	ServiceID *string
	// ID of ordered service
	AccountID *string
	Links     *SelfTView
}

// SelfTView is a type that runs validations on a projected type.
type SelfTView struct {
	Self        *string
	DescribedBy *DescribedByTView
}

// DescribedByTView is a type that runs validations on a projected type.
type DescribedByTView struct {
	Href *string
	Type *string
}

// NavTView is a type that runs validations on a projected type.
type NavTView struct {
	Self  *string
	First *string
	Next  *string
}

// OrderStatusRTView is a type that runs validations on a projected type.
type OrderStatusRTView struct {
	// Order ID
	ID *string
	// Order status
	Status *string
	// DateTime order was placed
	OrderedAt *string
	// DateTime order processing started
	StartedAt *string
	// DateTime order processing finished
	FinishedAt *string
	// Products delivered for this order
	Products []*ProductTView
	// Reference to service requested
	Service *RefTView
	// Reference to billable account
	Account *RefTView
	Links   *SelfTView
	// Optional customer provided name
	Name *string
	// Service parameters
	Parameters []*ParameterTView
}

// ProductTView is a type that runs validations on a projected type.
type ProductTView struct {
	ID       *string
	Name     *string
	Status   *string
	MimeType *string
	Size     *int64
	Links    *SelfWithDataTView
}

// SelfWithDataTView is a type that runs validations on a projected type.
type SelfWithDataTView struct {
	Self        *string
	DescribedBy *DescribedByTView
	Data        *string
}

// RefTView is a type that runs validations on a projected type.
type RefTView struct {
	ID    *string
	Links *SelfTView
}

// ParameterTView is a type that runs validations on a projected type.
type ParameterTView struct {
	Name  *string
	Value *string
}

var (
	// OrderListRTMap is a map indexing the attribute names of OrderListRT by view
	// name.
	OrderListRTMap = map[string][]string{
		"default": {
			"orders",
			"at-time",
			"links",
		},
	}
	// OrderStatusRTMap is a map indexing the attribute names of OrderStatusRT by
	// view name.
	OrderStatusRTMap = map[string][]string{
		"default": {
			"id",
			"name",
			"status",
			"ordered_at",
			"started_at",
			"finished_at",
			"parameters",
			"products",
			"service",
			"account",
			"links",
		},
		"tiny": {
			"name",
			"status",
			"links",
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
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateOrderStatusRT runs the validations defined on the viewed result type
// OrderStatusRT.
func ValidateOrderStatusRT(result *OrderStatusRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateOrderStatusRTView(result.Projected)
	case "tiny":
		err = ValidateOrderStatusRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateOrderListRTView runs the validations defined on OrderListRTView
// using the "default" view.
func ValidateOrderListRTView(result *OrderListRTView) (err error) {
	if result.Orders == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("orders", "result"))
	}
	if result.AtTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("at-time", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	for _, e := range result.Orders {
		if e != nil {
			if err2 := ValidateOrderListItemView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
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

// ValidateOrderListItemView runs the validations defined on OrderListItemView.
func ValidateOrderListItemView(result *OrderListItemView) (err error) {
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "pending" || *result.Status == "scheduled" || *result.Status == "executing" || *result.Status == "succeeded" || *result.Status == "failed" || *result.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if result.Links != nil {
		if err2 := ValidateSelfTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSelfTView runs the validations defined on SelfTView.
func ValidateSelfTView(result *SelfTView) (err error) {
	if result.DescribedBy != nil {
		if err2 := ValidateDescribedByTView(result.DescribedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDescribedByTView runs the validations defined on DescribedByTView.
func ValidateDescribedByTView(result *DescribedByTView) (err error) {
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.href", *result.Href, goa.FormatURI))
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

// ValidateOrderStatusRTView runs the validations defined on OrderStatusRTView
// using the "default" view.
func ValidateOrderStatusRTView(result *OrderStatusRTView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatUUID))
	}
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "pending" || *result.Status == "scheduled" || *result.Status == "executing" || *result.Status == "succeeded" || *result.Status == "failed" || *result.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	for _, e := range result.Products {
		if e != nil {
			if err2 := ValidateProductTView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Service != nil {
		if err2 := ValidateRefTView(result.Service); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Account != nil {
		if err2 := ValidateRefTView(result.Account); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Links != nil {
		if err2 := ValidateSelfTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateOrderStatusRTViewTiny runs the validations defined on
// OrderStatusRTView using the "tiny" view.
func ValidateOrderStatusRTViewTiny(result *OrderStatusRTView) (err error) {
	if result.Status != nil {
		if !(*result.Status == "unknown" || *result.Status == "pending" || *result.Status == "scheduled" || *result.Status == "executing" || *result.Status == "succeeded" || *result.Status == "failed" || *result.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"unknown", "pending", "scheduled", "executing", "succeeded", "failed", "error"}))
		}
	}
	if result.Links != nil {
		if err2 := ValidateSelfTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProductTView runs the validations defined on ProductTView.
func ValidateProductTView(result *ProductTView) (err error) {
	if result.Links != nil {
		if err2 := ValidateSelfWithDataTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSelfWithDataTView runs the validations defined on SelfWithDataTView.
func ValidateSelfWithDataTView(result *SelfWithDataTView) (err error) {
	if result.DescribedBy != nil {
		if err2 := ValidateDescribedByTView(result.DescribedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRefTView runs the validations defined on RefTView.
func ValidateRefTView(result *RefTView) (err error) {
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	if result.Links != nil {
		if err2 := ValidateSelfTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateParameterTView runs the validations defined on ParameterTView.
func ValidateParameterTView(result *ParameterTView) (err error) {

	return
}
