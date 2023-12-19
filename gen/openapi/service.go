// $ goa gen github.com/reinventingscience/ivcap-core-api/design

package openapi

// The openapi service serves the OpenAPI definition.
type Service interface {
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "openapi"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [0]string{}
