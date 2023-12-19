// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package client

import (
	"fmt"
)

// ReadOrderPath returns the URL path to the order service read HTTP endpoint.
func ReadOrderPath(id string) string {
	return fmt.Sprintf("/1/orders/%v", id)
}

// ListOrderPath returns the URL path to the order service list HTTP endpoint.
func ListOrderPath() string {
	return "/1/orders"
}

// CreateOrderPath returns the URL path to the order service create HTTP endpoint.
func CreateOrderPath() string {
	return "/1/orders"
}

// LogsOrderPath returns the URL path to the order service logs HTTP endpoint.
func LogsOrderPath() string {
	return "/1/orders/logs"
}

// TopOrderPath returns the URL path to the order service top HTTP endpoint.
func TopOrderPath() string {
	return "/1/orders/top"
}
