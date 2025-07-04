/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202409

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Product202409SearchManufacturersRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202409SearchManufacturersRequestBody{}

// Product202409SearchManufacturersRequestBody struct for Product202409SearchManufacturersRequestBody
type Product202409SearchManufacturersRequestBody struct {
    // Filter results to show those that contain this keyword. Search scope: name, registered trade name, local_number, email Max length: 200 characters  **Note**: Provide either the `manufacturer_ids` or `keyword`; if both are provided, `manufacturer_ids` will take priority.
    Keyword *string `json:"keyword,omitempty"`
    // Filter results by these manufacturer IDs. Max IDs: The value of `page_size`
    ManufacturerIds []string `json:"manufacturer_ids,omitempty"`
}

// NewProduct202409SearchManufacturersRequestBody instantiates a new Product202409SearchManufacturersRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202409SearchManufacturersRequestBody() *Product202409SearchManufacturersRequestBody {
    this := Product202409SearchManufacturersRequestBody{}
    return &this
}

// NewProduct202409SearchManufacturersRequestBodyWithDefaults instantiates a new Product202409SearchManufacturersRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202409SearchManufacturersRequestBodyWithDefaults() *Product202409SearchManufacturersRequestBody {
    this := Product202409SearchManufacturersRequestBody{}
    return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *Product202409SearchManufacturersRequestBody) GetKeyword() string {
    if o == nil || utils.IsNil(o.Keyword) {
        var ret string
        return ret
    }
    return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409SearchManufacturersRequestBody) GetKeywordOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Keyword) {
        return nil, false
    }
    return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *Product202409SearchManufacturersRequestBody) HasKeyword() bool {
    if o != nil && !utils.IsNil(o.Keyword) {
        return true
    }

    return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *Product202409SearchManufacturersRequestBody) SetKeyword(v string) {
    o.Keyword = &v
}

// GetManufacturerIds returns the ManufacturerIds field value if set, zero value otherwise.
func (o *Product202409SearchManufacturersRequestBody) GetManufacturerIds() []string {
    if o == nil || utils.IsNil(o.ManufacturerIds) {
        var ret []string
        return ret
    }
    return o.ManufacturerIds
}

// GetManufacturerIdsOk returns a tuple with the ManufacturerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409SearchManufacturersRequestBody) GetManufacturerIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.ManufacturerIds) {
        return nil, false
    }
    return o.ManufacturerIds, true
}

// HasManufacturerIds returns a boolean if a field has been set.
func (o *Product202409SearchManufacturersRequestBody) HasManufacturerIds() bool {
    if o != nil && !utils.IsNil(o.ManufacturerIds) {
        return true
    }

    return false
}

// SetManufacturerIds gets a reference to the given []string and assigns it to the ManufacturerIds field.
func (o *Product202409SearchManufacturersRequestBody) SetManufacturerIds(v []string) {
    o.ManufacturerIds = v
}

func (o Product202409SearchManufacturersRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202409SearchManufacturersRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Keyword) {
        toSerialize["keyword"] = o.Keyword
    }
    if !utils.IsNil(o.ManufacturerIds) {
        toSerialize["manufacturer_ids"] = o.ManufacturerIds
    }
    return toSerialize, nil
}

type NullableProduct202409SearchManufacturersRequestBody struct {
	value *Product202409SearchManufacturersRequestBody
	isSet bool
}

func (v NullableProduct202409SearchManufacturersRequestBody) Get() *Product202409SearchManufacturersRequestBody {
	return v.value
}

func (v *NullableProduct202409SearchManufacturersRequestBody) Set(val *Product202409SearchManufacturersRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202409SearchManufacturersRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202409SearchManufacturersRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202409SearchManufacturersRequestBody(val *Product202409SearchManufacturersRequestBody) *NullableProduct202409SearchManufacturersRequestBody {
	return &NullableProduct202409SearchManufacturersRequestBody{value: val, isSet: true}
}

func (v NullableProduct202409SearchManufacturersRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202409SearchManufacturersRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


