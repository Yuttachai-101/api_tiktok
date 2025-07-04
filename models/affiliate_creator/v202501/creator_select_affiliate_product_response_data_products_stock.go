/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_creator_v202501

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock{}

// AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock struct for AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock
type AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock struct {
    // The detailed stock quantity of the product
    Quantity *int32 `json:"quantity,omitempty"`
}

// NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock instantiates a new AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock() *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock {
    this := AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock{}
    return &this
}

// NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStockWithDefaults instantiates a new AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStockWithDefaults() *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock {
    this := AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock{}
    return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) GetQuantity() int32 {
    if o == nil || utils.IsNil(o.Quantity) {
        var ret int32
        return ret
    }
    return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) GetQuantityOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Quantity) {
        return nil, false
    }
    return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) HasQuantity() bool {
    if o != nil && !utils.IsNil(o.Quantity) {
        return true
    }

    return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) SetQuantity(v int32) {
    o.Quantity = &v
}

func (o AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Quantity) {
        toSerialize["quantity"] = o.Quantity
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock struct {
	value *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock
	isSet bool
}

func (v NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) Get() *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock {
	return v.value
}

func (v *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) Set(val *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock(val *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock {
	return &NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock{value: val, isSet: true}
}

func (v NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsStock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


