// Code generated by goa v3.10.2, DO NOT EDIT.
//
// HTTP request path constructors for the order service.
//
// Command:
// $ goa gen cayp/api_gateway/design

package client

import (
	"fmt"
)

// ListOrderPath returns the URL path to the order service list HTTP endpoint.
func ListOrderPath() string {
	return "/1/orders"
}

// CreateOrderPath returns the URL path to the order service create HTTP endpoint.
func CreateOrderPath() string {
	return "/1/orders"
}

// ReadOrderPath returns the URL path to the order service read HTTP endpoint.
func ReadOrderPath(id string) string {
	return fmt.Sprintf("/1/orders/%v", id)
}