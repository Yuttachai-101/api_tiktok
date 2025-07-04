/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Affiliate202309AddShowcaseProductsoldResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Affiliate202309AddShowcaseProductsoldResponseData{}

// Affiliate202309AddShowcaseProductsoldResponseData struct for Affiliate202309AddShowcaseProductsoldResponseData
type Affiliate202309AddShowcaseProductsoldResponseData struct {
    // The errors when adding products to showcase
    Errors []Affiliate202309AddShowcaseProductsoldResponseDataErrors `json:"errors,omitempty"`
}

// NewAffiliate202309AddShowcaseProductsoldResponseData instantiates a new Affiliate202309AddShowcaseProductsoldResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliate202309AddShowcaseProductsoldResponseData() *Affiliate202309AddShowcaseProductsoldResponseData {
    this := Affiliate202309AddShowcaseProductsoldResponseData{}
    return &this
}

// NewAffiliate202309AddShowcaseProductsoldResponseDataWithDefaults instantiates a new Affiliate202309AddShowcaseProductsoldResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliate202309AddShowcaseProductsoldResponseDataWithDefaults() *Affiliate202309AddShowcaseProductsoldResponseData {
    this := Affiliate202309AddShowcaseProductsoldResponseData{}
    return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Affiliate202309AddShowcaseProductsoldResponseData) GetErrors() []Affiliate202309AddShowcaseProductsoldResponseDataErrors {
    if o == nil || utils.IsNil(o.Errors) {
        var ret []Affiliate202309AddShowcaseProductsoldResponseDataErrors
        return ret
    }
    return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309AddShowcaseProductsoldResponseData) GetErrorsOk() ([]Affiliate202309AddShowcaseProductsoldResponseDataErrors, bool) {
    if o == nil || utils.IsNil(o.Errors) {
        return nil, false
    }
    return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Affiliate202309AddShowcaseProductsoldResponseData) HasErrors() bool {
    if o != nil && !utils.IsNil(o.Errors) {
        return true
    }

    return false
}

// SetErrors gets a reference to the given []Affiliate202309AddShowcaseProductsoldResponseDataErrors and assigns it to the Errors field.
func (o *Affiliate202309AddShowcaseProductsoldResponseData) SetErrors(v []Affiliate202309AddShowcaseProductsoldResponseDataErrors) {
    o.Errors = v
}

func (o Affiliate202309AddShowcaseProductsoldResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Affiliate202309AddShowcaseProductsoldResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Errors) {
        toSerialize["errors"] = o.Errors
    }
    return toSerialize, nil
}

type NullableAffiliate202309AddShowcaseProductsoldResponseData struct {
	value *Affiliate202309AddShowcaseProductsoldResponseData
	isSet bool
}

func (v NullableAffiliate202309AddShowcaseProductsoldResponseData) Get() *Affiliate202309AddShowcaseProductsoldResponseData {
	return v.value
}

func (v *NullableAffiliate202309AddShowcaseProductsoldResponseData) Set(val *Affiliate202309AddShowcaseProductsoldResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliate202309AddShowcaseProductsoldResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliate202309AddShowcaseProductsoldResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliate202309AddShowcaseProductsoldResponseData(val *Affiliate202309AddShowcaseProductsoldResponseData) *NullableAffiliate202309AddShowcaseProductsoldResponseData {
	return &NullableAffiliate202309AddShowcaseProductsoldResponseData{value: val, isSet: true}
}

func (v NullableAffiliate202309AddShowcaseProductsoldResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliate202309AddShowcaseProductsoldResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


