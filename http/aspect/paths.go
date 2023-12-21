// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	"fmt"
)

// ReadAspectPath returns the URL path to the aspect service read HTTP endpoint.
func ReadAspectPath(id string) string {
	return fmt.Sprintf("/1/aspects/%v", id)
}

// ListAspectPath returns the URL path to the aspect service list HTTP endpoint.
func ListAspectPath() string {
	return "/1/aspects"
}

// CreateAspectPath returns the URL path to the aspect service create HTTP endpoint.
func CreateAspectPath() string {
	return "/1/aspects"
}

// UpdateAspectPath returns the URL path to the aspect service update HTTP endpoint.
func UpdateAspectPath(id string) string {
	return fmt.Sprintf("/1/aspects/%v", id)
}

// RetractAspectPath returns the URL path to the aspect service retract HTTP endpoint.
func RetractAspectPath(id string) string {
	return fmt.Sprintf("/1/aspects/%v", id)
}
