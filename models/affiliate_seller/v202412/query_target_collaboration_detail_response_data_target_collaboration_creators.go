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

            // checks if the AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators{}

// AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators struct for AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators
type AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators struct {
    Avatar *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsAvatar `json:"avatar,omitempty"`
    // The status of the creator in the current target cooperation. Field values: - NORMAL: The status of the creator in the current target collaboration is normal. - DELETING: The status of the creator in the current target collaboration is deleting.The creator's product promotion relationship in the DELETING status will end the delayed effectiveness status and flow to the DELETED status at 00:00 the next day. - DELETED: The status of the creator in the current target collaboration is deleted.
    CollaborationStatus *string `json:"collaboration_status,omitempty"`
    // This field indicates the number of products creator has posted video or live from target collaboration. The count includes NORMAL and DELETING states.
    ContentProductCount *int64 `json:"content_product_count,omitempty"`
    // The TikTok nick name.
    Nickname *string `json:"nickname,omitempty"`
    // The effectiveness of the creators' commissions and products. Field values: - EFFECTIVE_ALL: The current product commission is effective for all creators. - EFFECTIVE_PARTIALLY: The current product commission are effective for some creators. - EFFECTIVE_NONE: The current product commission is not effective for all creators. Normally, the commission rate for all products under Target Collaboration is effective. If the merchant participates in TOP_CREATOR_PROGRAM, and TOP_CREATOR_PROGRAM includes the current creator and some products, the commission rate will be based on TOP_CREATOR_PROGRAM. The products' commission under the current Target Collaboration will only be partially effective for the creator, corresponding to EFFECTIVE_PARTIALLY status.
    ProductEffectiveStatus *string `json:"product_effective_status,omitempty"`
    // The regions in which the creator is eligible to promote products in showcases, videos, and live streams.
    SelectionRegion *string `json:"selection_region,omitempty"`
    // This field indicates the number of products creator has added  to the showcase from target collaboration. The count includes NORMAL and DELETING states.
    ShowcaseProductCount *int64 `json:"showcase_product_count,omitempty"`
    // The TikTok user name.
    Username *string `json:"username,omitempty"`
}

// NewAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators instantiates a new AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators() *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators {
    this := AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators{}
    return &this
}

// NewAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsWithDefaults instantiates a new AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsWithDefaults() *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators {
    this := AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators{}
    return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetAvatar() AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsAvatar {
    if o == nil || utils.IsNil(o.Avatar) {
        var ret AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsAvatar
        return ret
    }
    return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetAvatarOk() (*AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsAvatar, bool) {
    if o == nil || utils.IsNil(o.Avatar) {
        return nil, false
    }
    return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasAvatar() bool {
    if o != nil && !utils.IsNil(o.Avatar) {
        return true
    }

    return false
}

// SetAvatar gets a reference to the given AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsAvatar and assigns it to the Avatar field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetAvatar(v AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreatorsAvatar) {
    o.Avatar = &v
}

// GetCollaborationStatus returns the CollaborationStatus field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetCollaborationStatus() string {
    if o == nil || utils.IsNil(o.CollaborationStatus) {
        var ret string
        return ret
    }
    return *o.CollaborationStatus
}

// GetCollaborationStatusOk returns a tuple with the CollaborationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetCollaborationStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CollaborationStatus) {
        return nil, false
    }
    return o.CollaborationStatus, true
}

// HasCollaborationStatus returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasCollaborationStatus() bool {
    if o != nil && !utils.IsNil(o.CollaborationStatus) {
        return true
    }

    return false
}

// SetCollaborationStatus gets a reference to the given string and assigns it to the CollaborationStatus field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetCollaborationStatus(v string) {
    o.CollaborationStatus = &v
}

// GetContentProductCount returns the ContentProductCount field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetContentProductCount() int64 {
    if o == nil || utils.IsNil(o.ContentProductCount) {
        var ret int64
        return ret
    }
    return *o.ContentProductCount
}

// GetContentProductCountOk returns a tuple with the ContentProductCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetContentProductCountOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.ContentProductCount) {
        return nil, false
    }
    return o.ContentProductCount, true
}

// HasContentProductCount returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasContentProductCount() bool {
    if o != nil && !utils.IsNil(o.ContentProductCount) {
        return true
    }

    return false
}

// SetContentProductCount gets a reference to the given int64 and assigns it to the ContentProductCount field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetContentProductCount(v int64) {
    o.ContentProductCount = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetNickname() string {
    if o == nil || utils.IsNil(o.Nickname) {
        var ret string
        return ret
    }
    return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetNicknameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Nickname) {
        return nil, false
    }
    return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasNickname() bool {
    if o != nil && !utils.IsNil(o.Nickname) {
        return true
    }

    return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetNickname(v string) {
    o.Nickname = &v
}

// GetProductEffectiveStatus returns the ProductEffectiveStatus field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetProductEffectiveStatus() string {
    if o == nil || utils.IsNil(o.ProductEffectiveStatus) {
        var ret string
        return ret
    }
    return *o.ProductEffectiveStatus
}

// GetProductEffectiveStatusOk returns a tuple with the ProductEffectiveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetProductEffectiveStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ProductEffectiveStatus) {
        return nil, false
    }
    return o.ProductEffectiveStatus, true
}

// HasProductEffectiveStatus returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasProductEffectiveStatus() bool {
    if o != nil && !utils.IsNil(o.ProductEffectiveStatus) {
        return true
    }

    return false
}

// SetProductEffectiveStatus gets a reference to the given string and assigns it to the ProductEffectiveStatus field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetProductEffectiveStatus(v string) {
    o.ProductEffectiveStatus = &v
}

// GetSelectionRegion returns the SelectionRegion field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetSelectionRegion() string {
    if o == nil || utils.IsNil(o.SelectionRegion) {
        var ret string
        return ret
    }
    return *o.SelectionRegion
}

// GetSelectionRegionOk returns a tuple with the SelectionRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetSelectionRegionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SelectionRegion) {
        return nil, false
    }
    return o.SelectionRegion, true
}

// HasSelectionRegion returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasSelectionRegion() bool {
    if o != nil && !utils.IsNil(o.SelectionRegion) {
        return true
    }

    return false
}

// SetSelectionRegion gets a reference to the given string and assigns it to the SelectionRegion field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetSelectionRegion(v string) {
    o.SelectionRegion = &v
}

// GetShowcaseProductCount returns the ShowcaseProductCount field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetShowcaseProductCount() int64 {
    if o == nil || utils.IsNil(o.ShowcaseProductCount) {
        var ret int64
        return ret
    }
    return *o.ShowcaseProductCount
}

// GetShowcaseProductCountOk returns a tuple with the ShowcaseProductCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetShowcaseProductCountOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.ShowcaseProductCount) {
        return nil, false
    }
    return o.ShowcaseProductCount, true
}

// HasShowcaseProductCount returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasShowcaseProductCount() bool {
    if o != nil && !utils.IsNil(o.ShowcaseProductCount) {
        return true
    }

    return false
}

// SetShowcaseProductCount gets a reference to the given int64 and assigns it to the ShowcaseProductCount field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetShowcaseProductCount(v int64) {
    o.ShowcaseProductCount = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetUsername() string {
    if o == nil || utils.IsNil(o.Username) {
        var ret string
        return ret
    }
    return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) GetUsernameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Username) {
        return nil, false
    }
    return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) HasUsername() bool {
    if o != nil && !utils.IsNil(o.Username) {
        return true
    }

    return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) SetUsername(v string) {
    o.Username = &v
}

func (o AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Avatar) {
        toSerialize["avatar"] = o.Avatar
    }
    if !utils.IsNil(o.CollaborationStatus) {
        toSerialize["collaboration_status"] = o.CollaborationStatus
    }
    if !utils.IsNil(o.ContentProductCount) {
        toSerialize["content_product_count"] = o.ContentProductCount
    }
    if !utils.IsNil(o.Nickname) {
        toSerialize["nickname"] = o.Nickname
    }
    if !utils.IsNil(o.ProductEffectiveStatus) {
        toSerialize["product_effective_status"] = o.ProductEffectiveStatus
    }
    if !utils.IsNil(o.SelectionRegion) {
        toSerialize["selection_region"] = o.SelectionRegion
    }
    if !utils.IsNil(o.ShowcaseProductCount) {
        toSerialize["showcase_product_count"] = o.ShowcaseProductCount
    }
    if !utils.IsNil(o.Username) {
        toSerialize["username"] = o.Username
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators struct {
	value *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators
	isSet bool
}

func (v NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) Get() *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators {
	return v.value
}

func (v *NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) Set(val *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators(val *AffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) *NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators {
	return &NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators{value: val, isSet: true}
}

func (v NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202412QueryTargetCollaborationDetailResponseDataTargetCollaborationCreators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


