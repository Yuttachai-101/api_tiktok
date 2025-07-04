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

            // checks if the AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop{}

// AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop struct for AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop
type AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop struct {
    // The TikTok Shop name.
    Name *string `json:"name,omitempty"`
}

// NewAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop instantiates a new AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop() *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop {
    this := AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop{}
    return &this
}

// NewAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShopWithDefaults instantiates a new AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShopWithDefaults() *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop {
    this := AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop{}
    return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) GetName() string {
    if o == nil || utils.IsNil(o.Name) {
        var ret string
        return ret
    }
    return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) GetNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Name) {
        return nil, false
    }
    return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) HasName() bool {
    if o != nil && !utils.IsNil(o.Name) {
        return true
    }

    return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) SetName(v string) {
    o.Name = &v
}

func (o AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Name) {
        toSerialize["name"] = o.Name
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop struct {
	value *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop
	isSet bool
}

func (v NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) Get() *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop {
	return v.value
}

func (v *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) Set(val *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop(val *AffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop {
	return &NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop{value: val, isSet: true}
}

func (v NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductResponseDataProductsShop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


