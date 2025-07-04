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

            // checks if the AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange{}

// AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange struct for AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange
type AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange struct {
    // The minimum value of follower count.  The value passed in must be greater than or equal to 0
    CountGe *int64 `json:"count_ge,omitempty"`
    // The maximum value of follower count.  - Generally, a value greater than or equal to 0 needs to be passed. If a value less than 0 is passed, it means that the field will not be filtered.
    CountLe *int64 `json:"count_le,omitempty"`
}

// NewAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange instantiates a new AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange() *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange {
    this := AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange{}
    return &this
}

// NewAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRangeWithDefaults instantiates a new AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRangeWithDefaults() *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange {
    this := AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange{}
    return &this
}

// GetCountGe returns the CountGe field value if set, zero value otherwise.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) GetCountGe() int64 {
    if o == nil || utils.IsNil(o.CountGe) {
        var ret int64
        return ret
    }
    return *o.CountGe
}

// GetCountGeOk returns a tuple with the CountGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) GetCountGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CountGe) {
        return nil, false
    }
    return o.CountGe, true
}

// HasCountGe returns a boolean if a field has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) HasCountGe() bool {
    if o != nil && !utils.IsNil(o.CountGe) {
        return true
    }

    return false
}

// SetCountGe gets a reference to the given int64 and assigns it to the CountGe field.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) SetCountGe(v int64) {
    o.CountGe = &v
}

// GetCountLe returns the CountLe field value if set, zero value otherwise.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) GetCountLe() int64 {
    if o == nil || utils.IsNil(o.CountLe) {
        var ret int64
        return ret
    }
    return *o.CountLe
}

// GetCountLeOk returns a tuple with the CountLe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) GetCountLeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CountLe) {
        return nil, false
    }
    return o.CountLe, true
}

// HasCountLe returns a boolean if a field has been set.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) HasCountLe() bool {
    if o != nil && !utils.IsNil(o.CountLe) {
        return true
    }

    return false
}

// SetCountLe gets a reference to the given int64 and assigns it to the CountLe field.
func (o *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) SetCountLe(v int64) {
    o.CountLe = &v
}

func (o AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CountGe) {
        toSerialize["count_ge"] = o.CountGe
    }
    if !utils.IsNil(o.CountLe) {
        toSerialize["count_le"] = o.CountLe
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange struct {
	value *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange
	isSet bool
}

func (v NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) Get() *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange {
	return v.value
}

func (v *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) Set(val *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange(val *AffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange {
	return &NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505SellerSearchCreatoronMarketplaceRequestBodyFollowerDemographicsCountRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


