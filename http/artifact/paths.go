// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	"fmt"
)

// ListArtifactPath returns the URL path to the artifact service list HTTP endpoint.
func ListArtifactPath() string {
	return "/1/artifacts"
}

// UploadArtifactPath returns the URL path to the artifact service upload HTTP endpoint.
func UploadArtifactPath() string {
	return "/1/artifacts"
}

// ReadArtifactPath returns the URL path to the artifact service read HTTP endpoint.
func ReadArtifactPath(id string) string {
	return fmt.Sprintf("/1/artifacts/%v", id)
}
