/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory{}

// AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory struct for AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory
type AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory struct {
    // The category identifier. Note that only first-level categories are supported.
    Id *string `json:"id,omitempty"`
}

// NewAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory instantiates a new AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory() *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory {
    this := AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory{}
    return &this
}

// NewAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategoryWithDefaults instantiates a new AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategoryWithDefaults() *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory {
    this := AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) SetId(v string) {
    o.Id = &v
}

func (o AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory struct {
	value *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory
	isSet bool
}

func (v NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) Get() *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory {
	return v.value
}

func (v *NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) Set(val *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory(val *AffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) *NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory {
	return &NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory{value: val, isSet: true}
}

func (v NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202405SellerSearchAffiliateOpenCollaborationProductRequestBodyCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


