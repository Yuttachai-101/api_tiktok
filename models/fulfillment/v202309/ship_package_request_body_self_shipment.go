/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fulfillment_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Fulfillment202309ShipPackageRequestBodySelfShipment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309ShipPackageRequestBodySelfShipment{}

// Fulfillment202309ShipPackageRequestBodySelfShipment struct for Fulfillment202309ShipPackageRequestBodySelfShipment
type Fulfillment202309ShipPackageRequestBodySelfShipment struct {
    // For package with `SEND_BY_SELLER` as `delivery_option` (merchant self-shipping mode), you must input a `shipping_provider_id` to call this API. Please use [Get Shipping Providers](https://partner.tiktokshop.com/docv2/page/650aa48d4a0bb702c06d85cd?external_id=650aa48d4a0bb702c06d85cd) to obtain the `shipping_provider_id`.
    ShippingProviderId *string `json:"shipping_provider_id,omitempty"`
    // For package with `SEND_BY_SELLER` as `delivery_option` (merchant self-shipping mode), you must input a `tracking_number` to call this API.
    TrackingNumber *string `json:"tracking_number,omitempty"`
}

// NewFulfillment202309ShipPackageRequestBodySelfShipment instantiates a new Fulfillment202309ShipPackageRequestBodySelfShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309ShipPackageRequestBodySelfShipment() *Fulfillment202309ShipPackageRequestBodySelfShipment {
    this := Fulfillment202309ShipPackageRequestBodySelfShipment{}
    return &this
}

// NewFulfillment202309ShipPackageRequestBodySelfShipmentWithDefaults instantiates a new Fulfillment202309ShipPackageRequestBodySelfShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309ShipPackageRequestBodySelfShipmentWithDefaults() *Fulfillment202309ShipPackageRequestBodySelfShipment {
    this := Fulfillment202309ShipPackageRequestBodySelfShipment{}
    return &this
}

// GetShippingProviderId returns the ShippingProviderId field value if set, zero value otherwise.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) GetShippingProviderId() string {
    if o == nil || utils.IsNil(o.ShippingProviderId) {
        var ret string
        return ret
    }
    return *o.ShippingProviderId
}

// GetShippingProviderIdOk returns a tuple with the ShippingProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) GetShippingProviderIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ShippingProviderId) {
        return nil, false
    }
    return o.ShippingProviderId, true
}

// HasShippingProviderId returns a boolean if a field has been set.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) HasShippingProviderId() bool {
    if o != nil && !utils.IsNil(o.ShippingProviderId) {
        return true
    }

    return false
}

// SetShippingProviderId gets a reference to the given string and assigns it to the ShippingProviderId field.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) SetShippingProviderId(v string) {
    o.ShippingProviderId = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) GetTrackingNumber() string {
    if o == nil || utils.IsNil(o.TrackingNumber) {
        var ret string
        return ret
    }
    return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) GetTrackingNumberOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TrackingNumber) {
        return nil, false
    }
    return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) HasTrackingNumber() bool {
    if o != nil && !utils.IsNil(o.TrackingNumber) {
        return true
    }

    return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *Fulfillment202309ShipPackageRequestBodySelfShipment) SetTrackingNumber(v string) {
    o.TrackingNumber = &v
}

func (o Fulfillment202309ShipPackageRequestBodySelfShipment) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309ShipPackageRequestBodySelfShipment) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ShippingProviderId) {
        toSerialize["shipping_provider_id"] = o.ShippingProviderId
    }
    if !utils.IsNil(o.TrackingNumber) {
        toSerialize["tracking_number"] = o.TrackingNumber
    }
    return toSerialize, nil
}

type NullableFulfillment202309ShipPackageRequestBodySelfShipment struct {
	value *Fulfillment202309ShipPackageRequestBodySelfShipment
	isSet bool
}

func (v NullableFulfillment202309ShipPackageRequestBodySelfShipment) Get() *Fulfillment202309ShipPackageRequestBodySelfShipment {
	return v.value
}

func (v *NullableFulfillment202309ShipPackageRequestBodySelfShipment) Set(val *Fulfillment202309ShipPackageRequestBodySelfShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309ShipPackageRequestBodySelfShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309ShipPackageRequestBodySelfShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309ShipPackageRequestBodySelfShipment(val *Fulfillment202309ShipPackageRequestBodySelfShipment) *NullableFulfillment202309ShipPackageRequestBodySelfShipment {
	return &NullableFulfillment202309ShipPackageRequestBodySelfShipment{value: val, isSet: true}
}

func (v NullableFulfillment202309ShipPackageRequestBodySelfShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309ShipPackageRequestBodySelfShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


