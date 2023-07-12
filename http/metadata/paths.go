// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	"fmt"
)

// ListMetadataPath returns the URL path to the metadata service list HTTP endpoint.
func ListMetadataPath() string {
	return "/1/metadata"
}

// ReadMetadataPath returns the URL path to the metadata service read HTTP endpoint.
func ReadMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}

// AddMetadataPath returns the URL path to the metadata service add HTTP endpoint.
func AddMetadataPath() string {
	return "/1/metadata"
}

// UpdateMetadataPath returns the URL path to the metadata service update HTTP endpoint.
func UpdateMetadataPath() string {
	return "/1/metadata"
}

// RevokeMetadataPath returns the URL path to the metadata service revoke HTTP endpoint.
func RevokeMetadataPath(id string) string {
	return fmt.Sprintf("/1/metadata/%v", id)
}
