// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// AspectRT is the viewed result type that is projected based on a view.
type AspectRT struct {
	// Type to project
	Projected *AspectRTView
	// View to render
	View string
}

// AspectListRT is the viewed result type that is projected based on a view.
type AspectListRT struct {
	// Type to project
	Projected *AspectListRTView
	// View to render
	View string
}

// AspectIDRT is the viewed result type that is projected based on a view.
type AspectIDRT struct {
	// Type to project
	Projected *AspectIDRTView
	// View to render
	View string
}

// AspectRTView is a type that runs validations on a projected type.
type AspectRTView struct {
	// Record URN
	ID *string
	// Entity URN
	Entity *string
	// Schema URN
	Schema *string
	// Description of aspect encoded as 'content-type'
	Content any
	// Content-Type header, MUST be of application/json.
	ContentType *string `json:"content-type,omitempty"`
	// Time this aspect was asserted
	ValidFrom *string
	// Time this aspect was retractd
	ValidTo *string
	// Entity asserting this aspect aspect at 'valid-from'
	Asserter *string
	// Entity retracting this aspect at 'valid-to'
	Retracter *string
}

// AspectListRTView is a type that runs validations on a projected type.
type AspectListRTView struct {
	// List of aspect descriptions
	Items []*AspectListItemRTView
	// Entity for which to request aspect
	Entity *string
	// Optional schema to filter on
	Schema *string
	// Optional json path to further filter on returned list
	AspectPath *string
	// Time at which this list was valid
	AtTime *string
	// Navigation links
	Links *NavTView
}

// AspectListItemRTView is a type that runs validations on a projected type.
type AspectListItemRTView struct {
	// Record URN
	ID *string
	// Entity URN
	Entity *string
	// Schema URN
	Schema *string
	// Attached aspect aspect
	Content any
	// Content-Type header, MUST be of application/json.
	ContentType *string `json:"content-type,omitempty"`
}

// NavTView is a type that runs validations on a projected type.
type NavTView struct {
	Self  *string
	First *string
	Next  *string
}

// AspectIDRTView is a type that runs validations on a projected type.
type AspectIDRTView struct {
	// ID to specific aspect
	ID *string
}

var (
	// AspectRTMap is a map indexing the attribute names of AspectRT by view name.
	AspectRTMap = map[string][]string{
		"default": {
			"id",
			"entity",
			"schema",
			"content",
			"content-type",
			"valid-from",
			"valid-to",
			"asserter",
			"retracter",
		},
	}
	// AspectListRTMap is a map indexing the attribute names of AspectListRT by
	// view name.
	AspectListRTMap = map[string][]string{
		"default": {
			"items",
			"entity",
			"schema",
			"aspect-path",
			"at-time",
			"links",
		},
	}
	// AspectIDRTMap is a map indexing the attribute names of AspectIDRT by view
	// name.
	AspectIDRTMap = map[string][]string{
		"default": {
			"id",
		},
	}
	// AspectListItemRTMap is a map indexing the attribute names of
	// AspectListItemRT by view name.
	AspectListItemRTMap = map[string][]string{
		"default": {
			"id",
			"entity",
			"schema",
			"content",
			"content-type",
		},
	}
)

// ValidateAspectRT runs the validations defined on the viewed result type
// AspectRT.
func ValidateAspectRT(result *AspectRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateAspectRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateAspectListRT runs the validations defined on the viewed result type
// AspectListRT.
func ValidateAspectListRT(result *AspectListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateAspectListRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateAspectIDRT runs the validations defined on the viewed result type
// AspectIDRT.
func ValidateAspectIDRT(result *AspectIDRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateAspectIDRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateAspectRTView runs the validations defined on AspectRTView using the
// "default" view.
func ValidateAspectRTView(result *AspectRTView) (err error) {
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
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
	return
}

// ValidateAspectListRTView runs the validations defined on AspectListRTView
// using the "default" view.
func ValidateAspectListRTView(result *AspectListRTView) (err error) {
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
			if err2 := ValidateAspectListItemRTView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Entity != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.entity", *result.Entity, goa.FormatURI))
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

// ValidateAspectListItemRTView runs the validations defined on
// AspectListItemRTView using the "default" view.
func ValidateAspectListItemRTView(result *AspectListItemRTView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Entity == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("entity", "result"))
	}
	if result.Schema == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("schema", "result"))
	}
	if result.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "result"))
	}
	if result.ContentType == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content-type", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
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

// ValidateAspectIDRTView runs the validations defined on AspectIDRTView using
// the "default" view.
func ValidateAspectIDRTView(result *AspectIDRTView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	return
}
