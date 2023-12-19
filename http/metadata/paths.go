// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	"fmt"
)

// ReadMetadataPath returns the URL path to the metadata service read HTTP endpoint.
func ReadMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}

// ListMetadataPath returns the URL path to the metadata service list HTTP endpoint.
func ListMetadataPath() string {
	return "/1/metadata"
}

// AddMetadataPath returns the URL path to the metadata service add HTTP endpoint.
func AddMetadataPath() string {
	return "/1/metadata"
}

// UpdateOneMetadataPath returns the URL path to the metadata service update_one HTTP endpoint.
func UpdateOneMetadataPath() string {
	return "/1/metadata"
}

// UpdateRecordMetadataPath returns the URL path to the metadata service update_record HTTP endpoint.
func UpdateRecordMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}

// RevokeMetadataPath returns the URL path to the metadata service revoke HTTP endpoint.
func RevokeMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}
