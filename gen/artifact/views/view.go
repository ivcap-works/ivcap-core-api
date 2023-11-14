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

// ArtifactListRT is the viewed result type that is projected based on a view.
type ArtifactListRT struct {
	// Type to project
	Projected *ArtifactListRTView
	// View to render
	View string
}

// ArtifactStatusRT is the viewed result type that is projected based on a view.
type ArtifactStatusRT struct {
	// Type to project
	Projected *ArtifactStatusRTView
	// View to render
	View string
}

// ArtifactListRTView is a type that runs validations on a projected type.
type ArtifactListRTView struct {
	// Artifacts
	Artifacts []*ArtifactListItemView
	// Time at which this list was valid
	AtTime *string
	// Navigation links
	Links *NavTView
}

// ArtifactListItemView is a type that runs validations on a projected type.
type ArtifactListItemView struct {
	// Artifact ID
	ID *string
	// Optional name
	Name *string
	// Artifact status
	Status *string
	// Size of aritfact in bytes
	Size *int64
	// Mime (content) type of artifact
	MimeType *string
	Links    *SelfTView
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

// ArtifactStatusRTView is a type that runs validations on a projected type.
type ArtifactStatusRTView struct {
	// Artifact ID
	ID *string
	// Optional name
	Name *string
	// Artifact status
	Status *string
	// Mime-type of data
	MimeType *string
	// Size of data
	Size *int64
	// URL of object this artifact is caching
	CacheOf *string
	// ETAG of artifact
	Etag *string `json:"etag,omitempty"`
	// DateTime artifact was created
	CreatedAt *string
	// DateTime artifact was last modified
	LastModifiedAt *string
	// Reference to policy controlling access
	Policy *RefTView
	// Reference to billable account
	Account *RefTView
	// Link to retrieve the artifact data
	Data  *SelfTView
	Links *SelfTView
	// link back to record
	Location *string
	// indicate version of TUS supported
	TusResumable *string
	// TUS offset for partially uploaded content
	TusOffset *int64
}

// RefTView is a type that runs validations on a projected type.
type RefTView struct {
	ID    *string
	Links *SelfTView
}

var (
	// ArtifactListRTMap is a map indexing the attribute names of ArtifactListRT by
	// view name.
	ArtifactListRTMap = map[string][]string{
		"default": {
			"artifacts",
			"at-time",
			"links",
		},
	}
	// ArtifactStatusRTMap is a map indexing the attribute names of
	// ArtifactStatusRT by view name.
	ArtifactStatusRTMap = map[string][]string{
		"default": {
			"id",
			"name",
			"status",
			"mime-type",
			"size",
			"cache-of",
			"etag",
			"created-at",
			"last-modified-at",
			"policy",
			"account",
			"data",
			"links",
			"location",
			"tus-resumable",
			"tus-offset",
		},
	}
)

// ValidateArtifactListRT runs the validations defined on the viewed result
// type ArtifactListRT.
func ValidateArtifactListRT(result *ArtifactListRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateArtifactListRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateArtifactStatusRT runs the validations defined on the viewed result
// type ArtifactStatusRT.
func ValidateArtifactStatusRT(result *ArtifactStatusRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateArtifactStatusRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateArtifactListRTView runs the validations defined on
// ArtifactListRTView using the "default" view.
func ValidateArtifactListRTView(result *ArtifactListRTView) (err error) {
	if result.Artifacts == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("artifacts", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	for _, e := range result.Artifacts {
		if e != nil {
			if err2 := ValidateArtifactListItemView(e); err2 != nil {
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

// ValidateArtifactListItemView runs the validations defined on
// ArtifactListItemView.
func ValidateArtifactListItemView(result *ArtifactListItemView) (err error) {
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "pending" || *result.Status == "partial" || *result.Status == "ready" || *result.Status == "error" || *result.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"pending", "partial", "ready", "error", "unknown"}))
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

// ValidateArtifactStatusRTView runs the validations defined on
// ArtifactStatusRTView using the "default" view.
func ValidateArtifactStatusRTView(result *ArtifactStatusRTView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Links == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("links", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "pending" || *result.Status == "partial" || *result.Status == "ready" || *result.Status == "error" || *result.Status == "unknown") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"pending", "partial", "ready", "error", "unknown"}))
		}
	}
	if result.Policy != nil {
		if err2 := ValidateRefTView(result.Policy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Account != nil {
		if err2 := ValidateRefTView(result.Account); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Data != nil {
		if err2 := ValidateSelfTView(result.Data); err2 != nil {
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
