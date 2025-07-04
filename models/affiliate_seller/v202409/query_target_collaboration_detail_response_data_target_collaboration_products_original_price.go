/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202409

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice{}

// AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice struct for AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice
type AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice struct {
    // The currency code.
    Currency *string `json:"currency,omitempty"`
    // The highest original price of all SKUs of the product.
    MaximumAmount *string `json:"maximum_amount,omitempty"`
    // The lowest original price of all SKUs of the product.
    MinimumAmount *string `json:"minimum_amount,omitempty"`
}

// NewAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice instantiates a new AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice() *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice {
    this := AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice{}
    return &this
}

// NewAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPriceWithDefaults instantiates a new AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPriceWithDefaults() *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice {
    this := AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) SetCurrency(v string) {
    o.Currency = &v
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) GetMaximumAmount() string {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        var ret string
        return ret
    }
    return *o.MaximumAmount
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) GetMaximumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        return nil, false
    }
    return o.MaximumAmount, true
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) HasMaximumAmount() bool {
    if o != nil && !utils.IsNil(o.MaximumAmount) {
        return true
    }

    return false
}

// SetMaximumAmount gets a reference to the given string and assigns it to the MaximumAmount field.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) SetMaximumAmount(v string) {
    o.MaximumAmount = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) GetMinimumAmount() string {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        var ret string
        return ret
    }
    return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) GetMinimumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        return nil, false
    }
    return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) HasMinimumAmount() bool {
    if o != nil && !utils.IsNil(o.MinimumAmount) {
        return true
    }

    return false
}

// SetMinimumAmount gets a reference to the given string and assigns it to the MinimumAmount field.
func (o *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) SetMinimumAmount(v string) {
    o.MinimumAmount = &v
}

func (o AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    if !utils.IsNil(o.MaximumAmount) {
        toSerialize["maximum_amount"] = o.MaximumAmount
    }
    if !utils.IsNil(o.MinimumAmount) {
        toSerialize["minimum_amount"] = o.MinimumAmount
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice struct {
	value *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice
	isSet bool
}

func (v NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) Get() *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice {
	return v.value
}

func (v *NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) Set(val *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice(val *AffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) *NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice {
	return &NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice{value: val, isSet: true}
}

func (v NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202409QueryTargetCollaborationDetailResponseDataTargetCollaborationProductsOriginalPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


