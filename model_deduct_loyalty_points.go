/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// DeductLoyaltyPoints Points to deduct.
type DeductLoyaltyPoints struct {
	// Amount of loyalty points.
	Points float32 `json:"points"`
	// Name / reason for the point deduction.
	Name *string `json:"name,omitempty"`
	// ID of the subledger the points are deducted from.
	SubledgerId *string `json:"subledgerId,omitempty"`
	// ID of the Application that is connected to the loyalty program.
	ApplicationId *int32 `json:"applicationId,omitempty"`
}

// GetPoints returns the Points field value
func (o *DeductLoyaltyPoints) GetPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Points
}

// SetPoints sets field value
func (o *DeductLoyaltyPoints) SetPoints(v float32) {
	o.Points = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeductLoyaltyPoints) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPoints) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeductLoyaltyPoints) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeductLoyaltyPoints) SetName(v string) {
	o.Name = &v
}

// GetSubledgerId returns the SubledgerId field value if set, zero value otherwise.
func (o *DeductLoyaltyPoints) GetSubledgerId() string {
	if o == nil || o.SubledgerId == nil {
		var ret string
		return ret
	}
	return *o.SubledgerId
}

// GetSubledgerIdOk returns a tuple with the SubledgerId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPoints) GetSubledgerIdOk() (string, bool) {
	if o == nil || o.SubledgerId == nil {
		var ret string
		return ret, false
	}
	return *o.SubledgerId, true
}

// HasSubledgerId returns a boolean if a field has been set.
func (o *DeductLoyaltyPoints) HasSubledgerId() bool {
	if o != nil && o.SubledgerId != nil {
		return true
	}

	return false
}

// SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.
func (o *DeductLoyaltyPoints) SetSubledgerId(v string) {
	o.SubledgerId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *DeductLoyaltyPoints) GetApplicationId() int32 {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeductLoyaltyPoints) GetApplicationIdOk() (int32, bool) {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret, false
	}
	return *o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *DeductLoyaltyPoints) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *DeductLoyaltyPoints) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

type NullableDeductLoyaltyPoints struct {
	Value        DeductLoyaltyPoints
	ExplicitNull bool
}

func (v NullableDeductLoyaltyPoints) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDeductLoyaltyPoints) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}