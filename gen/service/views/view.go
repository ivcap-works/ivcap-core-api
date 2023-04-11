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

// ServiceListRT is the viewed result type that is projected based on a view.
type ServiceListRT struct {
	// Type to project
	Projected *ServiceListRTView
	// View to render
	View string
}

// ServiceStatusRT is the viewed result type that is projected based on a view.
type ServiceStatusRT struct {
	// Type to project
	Projected *ServiceStatusRTView
	// View to render
	View string
}

// OrderListRT is the viewed result type that is projected based on a view.
type OrderListRT struct {
	// Type to project
	Projected *OrderListRTView
	// View to render
	View string
}

// ServiceListRTView is a type that runs validations on a projected type.
type ServiceListRTView struct {
	// Services
	Services []*ServiceListItemView
	// Navigation links
	Links *NavTView
}

// ServiceListItemView is a type that runs validations on a projected type.
type ServiceListItemView struct {
	// Service ID
	ID *string
	// Optional customer provided name
	Name *string
	// Optional description of the service
	Description *string
	// Optional provider link
	Provider *RefTView
	Links    *SelfTView
}

// RefTView is a type that runs validations on a projected type.
type RefTView struct {
	ID    *string
	Links *SelfTView
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

// ServiceStatusRTView is a type that runs validations on a projected type.
type ServiceStatusRTView struct {
	// Service ID
	ID *string
	// Provider provided ID. Needs to be a single string with punctuations allowed.
	// Might have been changed
	ProviderRef *string
	// More detailed description of the service
	Description *string
	// Service status
	Status *string
	// Optional provider provided meta tags
	Metadata []*ParameterTView
	// Reference to service provider
	Provider *RefTView
	// Reference to billable account
	Account *RefTView
	Links   *SelfTView
	// Optional provider provided name
	Name *string
	// Optional provider provided tags
	Tags []string
	// Service parameter definitions
	Parameters []*ParameterDefTView
}

// ParameterTView is a type that runs validations on a projected type.
type ParameterTView struct {
	Name  *string
	Value *string
}

// ParameterDefTView is a type that runs validations on a projected type.
type ParameterDefTView struct {
	Name        *string
	Label       *string
	Type        *string
	Description *string
	Unit        *string
	Constant    *bool
	Optional    *bool
	Default     *string
	Options     []*ParameterOptTView
}

// ParameterOptTView is a type that runs validations on a projected type.
type ParameterOptTView struct {
	Value       *string
	Description *string
}

// OrderListRTView is a type that runs validations on a projected type.
type OrderListRTView struct {
	// Orders
	Orders []*OrderListItemView
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

var (
	// ServiceListRTMap is a map indexing the attribute names of ServiceListRT by
	// view name.
	ServiceListRTMap = map[string][]string{
		"default": {
			"services",
			"links",
		},
	}
	// ServiceStatusRTMap is a map indexing the attribute names of ServiceStatusRT
	// by view name.
	ServiceStatusRTMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"tags",
			"metadata",
			"parameters",
			"provider",
			"account",
			"links",
		},
		"tiny": {
			"name",
			"links",
		},
	}
	// OrderListRTMap is a map indexing the attribute names of OrderListRT by view
	// name.
	OrderListRTMap = map[string][]string{
		"default": {
			"orders",
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
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateServiceStatusRT runs the validations defined on the viewed result
// type ServiceStatusRT.
func ValidateServiceStatusRT(result *ServiceStatusRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateServiceStatusRTView(result.Projected)
	case "tiny":
		err = ValidateServiceStatusRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

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

// ValidateServiceListRTView runs the validations defined on ServiceListRTView
// using the "default" view.
func ValidateServiceListRTView(result *ServiceListRTView) (err error) {
	if result.Services == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("services", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	for _, e := range result.Services {
		if e != nil {
			if err2 := ValidateServiceListItemView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Links != nil {
		if err2 := ValidateNavTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateServiceListItemView runs the validations defined on
// ServiceListItemView.
func ValidateServiceListItemView(result *ServiceListItemView) (err error) {
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	if result.Provider != nil {
		if err2 := ValidateRefTView(result.Provider); err2 != nil {
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

// ValidateServiceStatusRTView runs the validations defined on
// ServiceStatusRTView using the "default" view.
func ValidateServiceStatusRTView(result *ServiceStatusRTView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	if result.Parameters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parameters", "result"))
	}
	if result.Provider != nil {
		if err2 := ValidateRefTView(result.Provider); err2 != nil {
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

// ValidateServiceStatusRTViewTiny runs the validations defined on
// ServiceStatusRTView using the "tiny" view.
func ValidateServiceStatusRTViewTiny(result *ServiceStatusRTView) (err error) {
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
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

// ValidateParameterDefTView runs the validations defined on ParameterDefTView.
func ValidateParameterDefTView(result *ParameterDefTView) (err error) {

	return
}

// ValidateParameterOptTView runs the validations defined on ParameterOptTView.
func ValidateParameterOptTView(result *ParameterOptTView) (err error) {

	return
}

// ValidateOrderListRTView runs the validations defined on OrderListRTView
// using the "default" view.
func ValidateOrderListRTView(result *OrderListRTView) (err error) {
	if result.Orders == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("orders", "result"))
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
		if !(*result.Status == "pending" || *result.Status == "executing" || *result.Status == "finished" || *result.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"pending", "executing", "finished", "error"}))
		}
	}
	if result.Links != nil {
		if err2 := ValidateSelfTView(result.Links); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
