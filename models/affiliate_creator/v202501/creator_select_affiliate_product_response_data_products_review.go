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

            // checks if the AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview{}

// AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview struct for AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview
type AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview struct {
    // The count of reviews
    Count *int32 `json:"count,omitempty"`
    // The average score of the product, the range is (0,5]
    OverallScore *string `json:"overall_score,omitempty"`
}

// NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview instantiates a new AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview() *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview {
    this := AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview{}
    return &this
}

// NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReviewWithDefaults instantiates a new AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReviewWithDefaults() *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview {
    this := AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview{}
    return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) GetCount() int32 {
    if o == nil || utils.IsNil(o.Count) {
        var ret int32
        return ret
    }
    return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) GetCountOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Count) {
        return nil, false
    }
    return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) HasCount() bool {
    if o != nil && !utils.IsNil(o.Count) {
        return true
    }

    return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) SetCount(v int32) {
    o.Count = &v
}

// GetOverallScore returns the OverallScore field value if set, zero value otherwise.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) GetOverallScore() string {
    if o == nil || utils.IsNil(o.OverallScore) {
        var ret string
        return ret
    }
    return *o.OverallScore
}

// GetOverallScoreOk returns a tuple with the OverallScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) GetOverallScoreOk() (*string, bool) {
    if o == nil || utils.IsNil(o.OverallScore) {
        return nil, false
    }
    return o.OverallScore, true
}

// HasOverallScore returns a boolean if a field has been set.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) HasOverallScore() bool {
    if o != nil && !utils.IsNil(o.OverallScore) {
        return true
    }

    return false
}

// SetOverallScore gets a reference to the given string and assigns it to the OverallScore field.
func (o *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) SetOverallScore(v string) {
    o.OverallScore = &v
}

func (o AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Count) {
        toSerialize["count"] = o.Count
    }
    if !utils.IsNil(o.OverallScore) {
        toSerialize["overall_score"] = o.OverallScore
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview struct {
	value *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview
	isSet bool
}

func (v NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) Get() *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview {
	return v.value
}

func (v *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) Set(val *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview(val *AffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview {
	return &NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview{value: val, isSet: true}
}

func (v NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202501CreatorSelectAffiliateProductResponseDataProductsReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


