/*
 * compute
 *
 * Manage instances in cloud.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package compute

type CreatedInstance struct {
	Id string `json:"id"`

	Image string `json:"image"`

	SwapSize string `json:"swapSize,omitempty"`

	Type string `json:"type"`

	Label string `json:"label"`
}

// AssertCreatedInstanceRequired checks if the required fields are not zero-ed
func AssertCreatedInstanceRequired(obj CreatedInstance) error {
	elements := map[string]interface{}{
		"id":    obj.Id,
		"image": obj.Image,
		"type":  obj.Type,
		"label": obj.Label,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreatedInstanceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreatedInstance (e.g. [][]CreatedInstance), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreatedInstanceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreatedInstance, ok := obj.(CreatedInstance)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreatedInstanceRequired(aCreatedInstance)
	})
}