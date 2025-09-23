// Copyright 2025 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
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

// Createqueueresponse is the viewed result type that is projected based on a
// view.
type Createqueueresponse struct {
	// Type to project
	Projected *CreatequeueresponseView
	// View to render
	View string
}

// Readqueueresponse is the viewed result type that is projected based on a
// view.
type Readqueueresponse struct {
	// Type to project
	Projected *ReadqueueresponseView
	// View to render
	View string
}

// Messagestatus is the viewed result type that is projected based on a view.
type Messagestatus struct {
	// Type to project
	Projected *MessagestatusView
	// View to render
	View string
}

// CreatequeueresponseView is a type that runs validations on a projected type.
type CreatequeueresponseView struct {
	// queue
	ID *string
	// Name of the created queue.
	Name *string
	// Description of the created queue.
	Description *string
	// Reference to billable account
	Account *string
}

// ReadqueueresponseView is a type that runs validations on a projected type.
type ReadqueueresponseView struct {
	// ID
	ID *string
	// Name of the queue.
	Name *string
	// Description of the queue.
	Description *string
	// Number of messages sent to the queue
	TotalMessages *uint64
	// Number of bytes in the queue
	Bytes *uint64
	// Timestamp of the first message in the queue
	FirstTime *string
	// Timestamp of the last message in the queue
	LastTime *string
	// Number of consumers
	ConsumerCount *int
	// Timestamp when the queue was created
	CreatedAt *string
}

// MessagestatusView is a type that runs validations on a projected type.
type MessagestatusView struct {
	// queue
	ID *string
}

// MessageListView is a type that runs validations on a projected type.
type MessageListView struct {
	// Messages in the queue
	Messages []*PublishedmessageView
	// Time at which this list was valid
	AtTime *string
}

// PublishedmessageView is a type that runs validations on a projected type.
type PublishedmessageView struct {
	// Message identifier
	ID *string
	// Message content in JSON format.
	Content any
	// Schema used for message
	Schema *string
	// Encoding type of message content (defaults to 'application/json')
	ContentType *string
}

var (
	// CreatequeueresponseMap is a map indexing the attribute names of
	// Createqueueresponse by view name.
	CreatequeueresponseMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"account",
		},
	}
	// ReadqueueresponseMap is a map indexing the attribute names of
	// Readqueueresponse by view name.
	ReadqueueresponseMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"total-messages",
			"bytes",
			"first-time",
			"last-time",
			"consumer-count",
			"created-at",
		},
	}
	// MessagestatusMap is a map indexing the attribute names of Messagestatus by
	// view name.
	MessagestatusMap = map[string][]string{
		"default": {
			"id",
		},
	}
	// PublishedmessageMap is a map indexing the attribute names of
	// Publishedmessage by view name.
	PublishedmessageMap = map[string][]string{
		"default": {
			"id",
			"content",
			"schema",
			"content-type",
		},
	}
)

// ValidateCreatequeueresponse runs the validations defined on the viewed
// result type Createqueueresponse.
func ValidateCreatequeueresponse(result *Createqueueresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateCreatequeueresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateReadqueueresponse runs the validations defined on the viewed result
// type Readqueueresponse.
func ValidateReadqueueresponse(result *Readqueueresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateReadqueueresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateMessagestatus runs the validations defined on the viewed result type
// Messagestatus.
func ValidateMessagestatus(result *Messagestatus) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateMessagestatusView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateCreatequeueresponseView runs the validations defined on
// CreatequeueresponseView using the "default" view.
func ValidateCreatequeueresponseView(result *CreatequeueresponseView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	if result.Account != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.account", *result.Account, goa.FormatURI))
	}
	return
}

// ValidateReadqueueresponseView runs the validations defined on
// ReadqueueresponseView using the "default" view.
func ValidateReadqueueresponseView(result *ReadqueueresponseView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created-at", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	if result.FirstTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.first-time", *result.FirstTime, goa.FormatDateTime))
	}
	if result.LastTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.last-time", *result.LastTime, goa.FormatDateTime))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created-at", *result.CreatedAt, goa.FormatDateTime))
	}
	return
}

// ValidateMessagestatusView runs the validations defined on MessagestatusView
// using the "default" view.
func ValidateMessagestatusView(result *MessagestatusView) (err error) {
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	return
}

// ValidateMessageListView runs the validations defined on MessageListView.
func ValidateMessageListView(result *MessageListView) (err error) {
	if result.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("messages", "result"))
	}
	for _, e := range result.Messages {
		if e != nil {
			if err2 := ValidatePublishedmessageView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.AtTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.at-time", *result.AtTime, goa.FormatDateTime))
	}
	return
}

// ValidatePublishedmessageView runs the validations defined on
// PublishedmessageView using the "default" view.
func ValidatePublishedmessageView(result *PublishedmessageView) (err error) {
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatURI))
	}
	return
}
