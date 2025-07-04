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

            // checks if the AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory{}

// AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory struct for AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory
type AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory struct {
    // The category identifier. Note that only first-level categories are supported.
    Id *string `json:"id,omitempty"`
}

// NewAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory instantiates a new AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory() *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory {
    this := AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory{}
    return &this
}

// NewAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategoryWithDefaults instantiates a new AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategoryWithDefaults() *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory {
    this := AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) SetId(v string) {
    o.Id = &v
}

func (o AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory struct {
	value *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory
	isSet bool
}

func (v NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) Get() *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory {
	return v.value
}

func (v *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) Set(val *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory(val *AffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory {
	return &NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory{value: val, isSet: true}
}

func (v NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202405CreatorSearchOpenCollaborationProductRequestBodyCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


