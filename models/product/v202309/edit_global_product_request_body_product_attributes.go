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

            // checks if the Product202309EditGlobalProductRequestBodyProductAttributes type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309EditGlobalProductRequestBodyProductAttributes{}

// Product202309EditGlobalProductRequestBodyProductAttributes struct for Product202309EditGlobalProductRequestBodyProductAttributes
type Product202309EditGlobalProductRequestBodyProductAttributes struct {
    // The ID of the product attribute, retrieved from the [Get Global Attributes API](https://partner.tiktokshop.com/docv2/page/650a0483c16ffe02b8dfc80a).
    Id *string `json:"id,omitempty"`
    // A list of selectable values for the product attribute. Max count: 300 for US; 100 for other regions.  **Note**: Provide either a built-in ID or a custom name; if both are provided, the ID takes priority.
    Values []Product202309EditGlobalProductRequestBodyProductAttributesValues `json:"values,omitempty"`
}

// NewProduct202309EditGlobalProductRequestBodyProductAttributes instantiates a new Product202309EditGlobalProductRequestBodyProductAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309EditGlobalProductRequestBodyProductAttributes() *Product202309EditGlobalProductRequestBodyProductAttributes {
    this := Product202309EditGlobalProductRequestBodyProductAttributes{}
    return &this
}

// NewProduct202309EditGlobalProductRequestBodyProductAttributesWithDefaults instantiates a new Product202309EditGlobalProductRequestBodyProductAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309EditGlobalProductRequestBodyProductAttributesWithDefaults() *Product202309EditGlobalProductRequestBodyProductAttributes {
    this := Product202309EditGlobalProductRequestBodyProductAttributes{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) SetId(v string) {
    o.Id = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) GetValues() []Product202309EditGlobalProductRequestBodyProductAttributesValues {
    if o == nil || utils.IsNil(o.Values) {
        var ret []Product202309EditGlobalProductRequestBodyProductAttributesValues
        return ret
    }
    return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) GetValuesOk() ([]Product202309EditGlobalProductRequestBodyProductAttributesValues, bool) {
    if o == nil || utils.IsNil(o.Values) {
        return nil, false
    }
    return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) HasValues() bool {
    if o != nil && !utils.IsNil(o.Values) {
        return true
    }

    return false
}

// SetValues gets a reference to the given []Product202309EditGlobalProductRequestBodyProductAttributesValues and assigns it to the Values field.
func (o *Product202309EditGlobalProductRequestBodyProductAttributes) SetValues(v []Product202309EditGlobalProductRequestBodyProductAttributesValues) {
    o.Values = v
}

func (o Product202309EditGlobalProductRequestBodyProductAttributes) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309EditGlobalProductRequestBodyProductAttributes) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Values) {
        toSerialize["values"] = o.Values
    }
    return toSerialize, nil
}

type NullableProduct202309EditGlobalProductRequestBodyProductAttributes struct {
	value *Product202309EditGlobalProductRequestBodyProductAttributes
	isSet bool
}

func (v NullableProduct202309EditGlobalProductRequestBodyProductAttributes) Get() *Product202309EditGlobalProductRequestBodyProductAttributes {
	return v.value
}

func (v *NullableProduct202309EditGlobalProductRequestBodyProductAttributes) Set(val *Product202309EditGlobalProductRequestBodyProductAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309EditGlobalProductRequestBodyProductAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309EditGlobalProductRequestBodyProductAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309EditGlobalProductRequestBodyProductAttributes(val *Product202309EditGlobalProductRequestBodyProductAttributes) *NullableProduct202309EditGlobalProductRequestBodyProductAttributes {
	return &NullableProduct202309EditGlobalProductRequestBodyProductAttributes{value: val, isSet: true}
}

func (v NullableProduct202309EditGlobalProductRequestBodyProductAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309EditGlobalProductRequestBodyProductAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


