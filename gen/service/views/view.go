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

// ServiceListRTView is a type that runs validations on a projected type.
type ServiceListRTView struct {
	// Services
	Items []*ServiceListItemView
	// Time at which this list was valid
	AtTime *string
	Links  []*LinkTView
}

// ServiceListItemView is a type that runs validations on a projected type.
type ServiceListItemView struct {
	// ID
	ID *string
	// Optional customer provided name
	Name *string
	// Optional description of the service
	Description *string
	// Reference to billable account
	Account *string `json:"account"`
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
			if err2 := ValidateServiceListItemView(e); err2 != nil {
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

// ValidateServiceListItemView runs the validations defined on
// ServiceListItemView.
func ValidateServiceListItemView(result *ServiceListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
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
