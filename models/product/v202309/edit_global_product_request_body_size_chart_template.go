/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Product202309EditGlobalProductRequestBodySizeChartTemplate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309EditGlobalProductRequestBodySizeChartTemplate{}

// Product202309EditGlobalProductRequestBodySizeChartTemplate struct for Product202309EditGlobalProductRequestBodySizeChartTemplate
type Product202309EditGlobalProductRequestBodySizeChartTemplate struct {
    // The size chart template ID.
    Id *string `json:"id,omitempty"`
}

// NewProduct202309EditGlobalProductRequestBodySizeChartTemplate instantiates a new Product202309EditGlobalProductRequestBodySizeChartTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309EditGlobalProductRequestBodySizeChartTemplate() *Product202309EditGlobalProductRequestBodySizeChartTemplate {
    this := Product202309EditGlobalProductRequestBodySizeChartTemplate{}
    return &this
}

// NewProduct202309EditGlobalProductRequestBodySizeChartTemplateWithDefaults instantiates a new Product202309EditGlobalProductRequestBodySizeChartTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309EditGlobalProductRequestBodySizeChartTemplateWithDefaults() *Product202309EditGlobalProductRequestBodySizeChartTemplate {
    this := Product202309EditGlobalProductRequestBodySizeChartTemplate{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309EditGlobalProductRequestBodySizeChartTemplate) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditGlobalProductRequestBodySizeChartTemplate) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309EditGlobalProductRequestBodySizeChartTemplate) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309EditGlobalProductRequestBodySizeChartTemplate) SetId(v string) {
    o.Id = &v
}

func (o Product202309EditGlobalProductRequestBodySizeChartTemplate) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309EditGlobalProductRequestBodySizeChartTemplate) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    return toSerialize, nil
}

type NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate struct {
	value *Product202309EditGlobalProductRequestBodySizeChartTemplate
	isSet bool
}

func (v NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate) Get() *Product202309EditGlobalProductRequestBodySizeChartTemplate {
	return v.value
}

func (v *NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate) Set(val *Product202309EditGlobalProductRequestBodySizeChartTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309EditGlobalProductRequestBodySizeChartTemplate(val *Product202309EditGlobalProductRequestBodySizeChartTemplate) *NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate {
	return &NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate{value: val, isSet: true}
}

func (v NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309EditGlobalProductRequestBodySizeChartTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


