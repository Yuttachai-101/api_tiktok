/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202505

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{}

// AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution struct for AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution
type AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution struct {
    // The top-level category identifier. 
    CategoryId *string `json:"category_id,omitempty"`
    // GMV associated with the category in hundredths of a percent.
    Value *string `json:"value,omitempty"`
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{}
    return &this
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistributionWithDefaults instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistributionWithDefaults() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{}
    return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetCategoryId() string {
    if o == nil || utils.IsNil(o.CategoryId) {
        var ret string
        return ret
    }
    return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetCategoryIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CategoryId) {
        return nil, false
    }
    return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) HasCategoryId() bool {
    if o != nil && !utils.IsNil(o.CategoryId) {
        return true
    }

    return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) SetCategoryId(v string) {
    o.CategoryId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetValue() string {
    if o == nil || utils.IsNil(o.Value) {
        var ret string
        return ret
    }
    return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetValueOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Value) {
        return nil, false
    }
    return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) HasValue() bool {
    if o != nil && !utils.IsNil(o.Value) {
        return true
    }

    return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) SetValue(v string) {
    o.Value = &v
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CategoryId) {
        toSerialize["category_id"] = o.CategoryId
    }
    if !utils.IsNil(o.Value) {
        toSerialize["value"] = o.Value
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution struct {
	value *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution
	isSet bool
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) Get() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
	return v.value
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) Set(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
	return &NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


