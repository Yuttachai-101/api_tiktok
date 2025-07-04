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

            // checks if the Fulfillment202309BatchShipPackagesRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309BatchShipPackagesRequestBody{}

// Fulfillment202309BatchShipPackagesRequestBody struct for Fulfillment202309BatchShipPackagesRequestBody
type Fulfillment202309BatchShipPackagesRequestBody struct {
    // Input list of packages you would like to batch ship.
    Packages []Fulfillment202309BatchShipPackagesRequestBodyPackages `json:"packages,omitempty"`
}

// NewFulfillment202309BatchShipPackagesRequestBody instantiates a new Fulfillment202309BatchShipPackagesRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309BatchShipPackagesRequestBody() *Fulfillment202309BatchShipPackagesRequestBody {
    this := Fulfillment202309BatchShipPackagesRequestBody{}
    return &this
}

// NewFulfillment202309BatchShipPackagesRequestBodyWithDefaults instantiates a new Fulfillment202309BatchShipPackagesRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309BatchShipPackagesRequestBodyWithDefaults() *Fulfillment202309BatchShipPackagesRequestBody {
    this := Fulfillment202309BatchShipPackagesRequestBody{}
    return &this
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *Fulfillment202309BatchShipPackagesRequestBody) GetPackages() []Fulfillment202309BatchShipPackagesRequestBodyPackages {
    if o == nil || utils.IsNil(o.Packages) {
        var ret []Fulfillment202309BatchShipPackagesRequestBodyPackages
        return ret
    }
    return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309BatchShipPackagesRequestBody) GetPackagesOk() ([]Fulfillment202309BatchShipPackagesRequestBodyPackages, bool) {
    if o == nil || utils.IsNil(o.Packages) {
        return nil, false
    }
    return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *Fulfillment202309BatchShipPackagesRequestBody) HasPackages() bool {
    if o != nil && !utils.IsNil(o.Packages) {
        return true
    }

    return false
}

// SetPackages gets a reference to the given []Fulfillment202309BatchShipPackagesRequestBodyPackages and assigns it to the Packages field.
func (o *Fulfillment202309BatchShipPackagesRequestBody) SetPackages(v []Fulfillment202309BatchShipPackagesRequestBodyPackages) {
    o.Packages = v
}

func (o Fulfillment202309BatchShipPackagesRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309BatchShipPackagesRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Packages) {
        toSerialize["packages"] = o.Packages
    }
    return toSerialize, nil
}

type NullableFulfillment202309BatchShipPackagesRequestBody struct {
	value *Fulfillment202309BatchShipPackagesRequestBody
	isSet bool
}

func (v NullableFulfillment202309BatchShipPackagesRequestBody) Get() *Fulfillment202309BatchShipPackagesRequestBody {
	return v.value
}

func (v *NullableFulfillment202309BatchShipPackagesRequestBody) Set(val *Fulfillment202309BatchShipPackagesRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309BatchShipPackagesRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309BatchShipPackagesRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309BatchShipPackagesRequestBody(val *Fulfillment202309BatchShipPackagesRequestBody) *NullableFulfillment202309BatchShipPackagesRequestBody {
	return &NullableFulfillment202309BatchShipPackagesRequestBody{value: val, isSet: true}
}

func (v NullableFulfillment202309BatchShipPackagesRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309BatchShipPackagesRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


