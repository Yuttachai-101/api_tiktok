/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_creator_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice{}

// AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice struct for AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice
type AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice struct {
    // The currency code.
    Currency *string `json:"currency,omitempty"`
    // The highest product platform discount price.
    MaximumAmount *string `json:"maximum_amount,omitempty"`
    // The lowest product platform discount price.
    MinimumAmount *string `json:"minimum_amount,omitempty"`
}

// NewAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice instantiates a new AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice() *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice {
    this := AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice{}
    return &this
}

// NewAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPriceWithDefaults instantiates a new AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPriceWithDefaults() *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice {
    this := AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) SetCurrency(v string) {
    o.Currency = &v
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) GetMaximumAmount() string {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        var ret string
        return ret
    }
    return *o.MaximumAmount
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) GetMaximumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MaximumAmount) {
        return nil, false
    }
    return o.MaximumAmount, true
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) HasMaximumAmount() bool {
    if o != nil && !utils.IsNil(o.MaximumAmount) {
        return true
    }

    return false
}

// SetMaximumAmount gets a reference to the given string and assigns it to the MaximumAmount field.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) SetMaximumAmount(v string) {
    o.MaximumAmount = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) GetMinimumAmount() string {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        var ret string
        return ret
    }
    return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) GetMinimumAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MinimumAmount) {
        return nil, false
    }
    return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) HasMinimumAmount() bool {
    if o != nil && !utils.IsNil(o.MinimumAmount) {
        return true
    }

    return false
}

// SetMinimumAmount gets a reference to the given string and assigns it to the MinimumAmount field.
func (o *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) SetMinimumAmount(v string) {
    o.MinimumAmount = &v
}

func (o AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) ToMap() (map[string]interface{}, error) {
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

type NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice struct {
	value *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice
	isSet bool
}

func (v NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) Get() *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice {
	return v.value
}

func (v *NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) Set(val *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice(val *AffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) *NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice {
	return &NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice{value: val, isSet: true}
}

func (v NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202405GetShowcaseProductsResponseDataProductsPricePlatformDiscountPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


