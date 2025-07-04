/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202406

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{}

// AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution struct for AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution
type AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution struct {
    // The top-level category identifier. 
    CategoryId *string `json:"category_id,omitempty"`
    // GMV associated with the category in hundredths of a percent.
    Value *string `json:"value,omitempty"`
}

// NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution instantiates a new AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution() *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
    this := AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{}
    return &this
}

// NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistributionWithDefaults instantiates a new AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistributionWithDefaults() *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
    this := AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{}
    return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetCategoryId() string {
    if o == nil || utils.IsNil(o.CategoryId) {
        var ret string
        return ret
    }
    return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetCategoryIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CategoryId) {
        return nil, false
    }
    return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) HasCategoryId() bool {
    if o != nil && !utils.IsNil(o.CategoryId) {
        return true
    }

    return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) SetCategoryId(v string) {
    o.CategoryId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetValue() string {
    if o == nil || utils.IsNil(o.Value) {
        var ret string
        return ret
    }
    return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) GetValueOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Value) {
        return nil, false
    }
    return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) HasValue() bool {
    if o != nil && !utils.IsNil(o.Value) {
        return true
    }

    return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) SetValue(v string) {
    o.Value = &v
}

func (o AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CategoryId) {
        toSerialize["category_id"] = o.CategoryId
    }
    if !utils.IsNil(o.Value) {
        toSerialize["value"] = o.Value
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution struct {
	value *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution
	isSet bool
}

func (v NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) Get() *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
	return v.value
}

func (v *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) Set(val *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution(val *AffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution {
	return &NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution{value: val, isSet: true}
}

func (v NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202406GetMarketplaceCreatorPerformanceResponseDataCreatorCategoryGmvDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


