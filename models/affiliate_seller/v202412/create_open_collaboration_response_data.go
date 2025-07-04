/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202412

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202412CreateOpenCollaborationResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202412CreateOpenCollaborationResponseData{}

// AffiliateSeller202412CreateOpenCollaborationResponseData struct for AffiliateSeller202412CreateOpenCollaborationResponseData
type AffiliateSeller202412CreateOpenCollaborationResponseData struct {
    OpenCollaboration *AffiliateSeller202412CreateOpenCollaborationResponseDataOpenCollaboration `json:"open_collaboration,omitempty"`
}

// NewAffiliateSeller202412CreateOpenCollaborationResponseData instantiates a new AffiliateSeller202412CreateOpenCollaborationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202412CreateOpenCollaborationResponseData() *AffiliateSeller202412CreateOpenCollaborationResponseData {
    this := AffiliateSeller202412CreateOpenCollaborationResponseData{}
    return &this
}

// NewAffiliateSeller202412CreateOpenCollaborationResponseDataWithDefaults instantiates a new AffiliateSeller202412CreateOpenCollaborationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202412CreateOpenCollaborationResponseDataWithDefaults() *AffiliateSeller202412CreateOpenCollaborationResponseData {
    this := AffiliateSeller202412CreateOpenCollaborationResponseData{}
    return &this
}

// GetOpenCollaboration returns the OpenCollaboration field value if set, zero value otherwise.
func (o *AffiliateSeller202412CreateOpenCollaborationResponseData) GetOpenCollaboration() AffiliateSeller202412CreateOpenCollaborationResponseDataOpenCollaboration {
    if o == nil || utils.IsNil(o.OpenCollaboration) {
        var ret AffiliateSeller202412CreateOpenCollaborationResponseDataOpenCollaboration
        return ret
    }
    return *o.OpenCollaboration
}

// GetOpenCollaborationOk returns a tuple with the OpenCollaboration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412CreateOpenCollaborationResponseData) GetOpenCollaborationOk() (*AffiliateSeller202412CreateOpenCollaborationResponseDataOpenCollaboration, bool) {
    if o == nil || utils.IsNil(o.OpenCollaboration) {
        return nil, false
    }
    return o.OpenCollaboration, true
}

// HasOpenCollaboration returns a boolean if a field has been set.
func (o *AffiliateSeller202412CreateOpenCollaborationResponseData) HasOpenCollaboration() bool {
    if o != nil && !utils.IsNil(o.OpenCollaboration) {
        return true
    }

    return false
}

// SetOpenCollaboration gets a reference to the given AffiliateSeller202412CreateOpenCollaborationResponseDataOpenCollaboration and assigns it to the OpenCollaboration field.
func (o *AffiliateSeller202412CreateOpenCollaborationResponseData) SetOpenCollaboration(v AffiliateSeller202412CreateOpenCollaborationResponseDataOpenCollaboration) {
    o.OpenCollaboration = &v
}

func (o AffiliateSeller202412CreateOpenCollaborationResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202412CreateOpenCollaborationResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.OpenCollaboration) {
        toSerialize["open_collaboration"] = o.OpenCollaboration
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202412CreateOpenCollaborationResponseData struct {
	value *AffiliateSeller202412CreateOpenCollaborationResponseData
	isSet bool
}

func (v NullableAffiliateSeller202412CreateOpenCollaborationResponseData) Get() *AffiliateSeller202412CreateOpenCollaborationResponseData {
	return v.value
}

func (v *NullableAffiliateSeller202412CreateOpenCollaborationResponseData) Set(val *AffiliateSeller202412CreateOpenCollaborationResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202412CreateOpenCollaborationResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202412CreateOpenCollaborationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202412CreateOpenCollaborationResponseData(val *AffiliateSeller202412CreateOpenCollaborationResponseData) *NullableAffiliateSeller202412CreateOpenCollaborationResponseData {
	return &NullableAffiliateSeller202412CreateOpenCollaborationResponseData{value: val, isSet: true}
}

func (v NullableAffiliateSeller202412CreateOpenCollaborationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202412CreateOpenCollaborationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


