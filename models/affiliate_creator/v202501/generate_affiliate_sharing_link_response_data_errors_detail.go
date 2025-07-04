/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_creator_v202501

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail{}

// AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail struct for AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail
type AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail struct {
    // Detail fail reason for specific tag
    FailReason *string `json:"fail_reason,omitempty"`
    // Same as description in request params
    Tag *string `json:"tag,omitempty"`
}

// NewAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail instantiates a new AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail() *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail {
    this := AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail{}
    return &this
}

// NewAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetailWithDefaults instantiates a new AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetailWithDefaults() *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail {
    this := AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail{}
    return &this
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) GetFailReason() string {
    if o == nil || utils.IsNil(o.FailReason) {
        var ret string
        return ret
    }
    return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) GetFailReasonOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FailReason) {
        return nil, false
    }
    return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) HasFailReason() bool {
    if o != nil && !utils.IsNil(o.FailReason) {
        return true
    }

    return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) SetFailReason(v string) {
    o.FailReason = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) GetTag() string {
    if o == nil || utils.IsNil(o.Tag) {
        var ret string
        return ret
    }
    return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) GetTagOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Tag) {
        return nil, false
    }
    return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) HasTag() bool {
    if o != nil && !utils.IsNil(o.Tag) {
        return true
    }

    return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) SetTag(v string) {
    o.Tag = &v
}

func (o AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.FailReason) {
        toSerialize["fail_reason"] = o.FailReason
    }
    if !utils.IsNil(o.Tag) {
        toSerialize["tag"] = o.Tag
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail struct {
	value *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail
	isSet bool
}

func (v NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) Get() *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail {
	return v.value
}

func (v *NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) Set(val *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail(val *AffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) *NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail {
	return &NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail{value: val, isSet: true}
}

func (v NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202501GenerateAffiliateSharingLinkResponseDataErrorsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


