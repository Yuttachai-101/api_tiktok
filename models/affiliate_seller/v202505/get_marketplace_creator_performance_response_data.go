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

            // checks if the AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData{}

// AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData struct for AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData
type AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData struct {
    Creator *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreator `json:"creator,omitempty"`
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData{}
    return &this
}

// NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataWithDefaults instantiates a new AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataWithDefaults() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData {
    this := AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData{}
    return &this
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) GetCreator() AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreator {
    if o == nil || utils.IsNil(o.Creator) {
        var ret AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreator
        return ret
    }
    return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) GetCreatorOk() (*AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreator, bool) {
    if o == nil || utils.IsNil(o.Creator) {
        return nil, false
    }
    return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) HasCreator() bool {
    if o != nil && !utils.IsNil(o.Creator) {
        return true
    }

    return false
}

// SetCreator gets a reference to the given AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreator and assigns it to the Creator field.
func (o *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) SetCreator(v AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseDataCreator) {
    o.Creator = &v
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Creator) {
        toSerialize["creator"] = o.Creator
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData struct {
	value *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData
	isSet bool
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) Get() *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData {
	return v.value
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) Set(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData(val *AffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData {
	return &NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData{value: val, isSet: true}
}

func (v NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202505GetMarketplaceCreatorPerformanceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


