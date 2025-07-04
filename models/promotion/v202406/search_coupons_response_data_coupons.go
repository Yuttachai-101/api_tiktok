/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package promotion_v202406

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Promotion202406SearchCouponsResponseDataCoupons type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Promotion202406SearchCouponsResponseDataCoupons{}

// Promotion202406SearchCouponsResponseDataCoupons struct for Promotion202406SearchCouponsResponseDataCoupons
type Promotion202406SearchCouponsResponseDataCoupons struct {
    ClaimDuration *Promotion202406SearchCouponsResponseDataCouponsClaimDuration `json:"claim_duration,omitempty"`
    // The UNIX timestamp of when the coupon was created.
    CreateTime *int64 `json:"create_time,omitempty"`
    // The system where the coupon is created: - `SELLER_CENTER`: Created via the Promotions section of TikTok Seller Center. - `SELLER_APP`: Created via the TikTok Seller Mobile App. - `TTS_CRM`: Created via the TikTok Shop CRM.
    CreationSource *string `json:"creation_source,omitempty"`
    Discount *Promotion202406SearchCouponsResponseDataCouponsDiscount `json:"discount,omitempty"`
    // The display type of coupons. Values: - `REGULAR`: Regular coupons which are displayed to TikTok users across all display locations available in TTS, including PLPs, PDPs, TikTok Videos, LIVE Rooms, Creator Showcases, and may be shared via TTS Customer Support instant messages. Includes coupons which target select customer segments. - `LIVE`: Coupons which are only displayed to TikTok users in LIVE Rooms. - `CREATOR_EXCLUSIVE`: Coupons that can be claimed through the display channels of the specified creator, such as their LIVE Rooms, and their Creator Showcase. - `CHAT`: Exclusive coupons that may be sent to customers via TTS customer support chat messages, but are not displayed in other display locations. - `PROMO_CODE`: Promo codes do not have dedicated display locations on TikTok, but may be shared with customers in LIVEs, in TikTok videos, or other social media platforms. They can be claimed by customers entering a claim code at checkout or by customers that use a custom promo code landing page URL.   
    DisplayType *string `json:"display_type,omitempty"`
    // A unique ID that identifies different coupons.
    Id *string `json:"id,omitempty"`
    // The range of the products which the coupon applies to. The possible enumerations are: - `FULL_SHOP`: All products sold in the shop are eligible for the coupon. - `SPECIFIC_PRODUCTS`: Only specified products are eligible for the coupon. The list of specified products are returned in the response of the `get_coupon` API.
    ProductScope *string `json:"product_scope,omitempty"`
    // Promotion code. The string contains only Unicode letters or digits. The length of the string is in the range of `[6, 12]`.
    PromoCode *string `json:"promo_code,omitempty"`
    RedemptionDuration *Promotion202406SearchCouponsResponseDataCouponsRedemptionDuration `json:"redemption_duration,omitempty"`
    // Coupon's promotion status. Values: - `NOT_START`: Not available to TikTok users until the coupon's configured start time. - `ONGOING`: Available to TikTok users. - `EXPIRED`: Not available to TikTok users because it has expired. - `DEACTIVATED`: Deactivated by the seller and is not available to TikTok users.
    Status *string `json:"status,omitempty"`
    // The target buyer segment of the coupon. Possible enumerations are: - `ALL`: May be discovered and claimed by all TTS buyers. - `NEW`:Customers who have never purchased from your shop. - `REPEAT_CUSTOMERS`: People who have previously placed orders in your shop and made another purchase within certain days(30 days for non-US and 90 days for US). - `RECENT_CUSTOMERS`: People who have made their first purchase in your shop in the past certain days(30 days for non-US and 90 days for US). - `FREQUENT_CUSTOMERS`: Customers with more than 1 purchase within the last 90 days. - `LAPSED_CUSTOMERS`: Customers with at least 1 purchase in the past 365 days but no purchases within the last certain days(90 days for non-US and 30 days for US). - `NEW_FOLLOWERS`: People who started following the TikTok account of your shop in the past 30 days. - `EXISTING_ACTIVE_FOLLOWERS`: People who followed the TikTok account of your shop and engaged with your shop through LIVE, short videos or product cards in the past 30 days. 
    TargetBuyerSegment *string `json:"target_buyer_segment,omitempty"`
    Threshold *Promotion202406SearchCouponsResponseDataCouponsThreshold `json:"threshold,omitempty"`
    // Seller-specified title of the coupon.
    Title *string `json:"title,omitempty"`
    // The UNIX timestamp of when the coupon was updated.
    UpdateTime *int64 `json:"update_time,omitempty"`
    UsageLimits *Promotion202406SearchCouponsResponseDataCouponsUsageLimits `json:"usage_limits,omitempty"`
}

// NewPromotion202406SearchCouponsResponseDataCoupons instantiates a new Promotion202406SearchCouponsResponseDataCoupons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion202406SearchCouponsResponseDataCoupons() *Promotion202406SearchCouponsResponseDataCoupons {
    this := Promotion202406SearchCouponsResponseDataCoupons{}
    return &this
}

// NewPromotion202406SearchCouponsResponseDataCouponsWithDefaults instantiates a new Promotion202406SearchCouponsResponseDataCoupons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotion202406SearchCouponsResponseDataCouponsWithDefaults() *Promotion202406SearchCouponsResponseDataCoupons {
    this := Promotion202406SearchCouponsResponseDataCoupons{}
    return &this
}

// GetClaimDuration returns the ClaimDuration field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetClaimDuration() Promotion202406SearchCouponsResponseDataCouponsClaimDuration {
    if o == nil || utils.IsNil(o.ClaimDuration) {
        var ret Promotion202406SearchCouponsResponseDataCouponsClaimDuration
        return ret
    }
    return *o.ClaimDuration
}

// GetClaimDurationOk returns a tuple with the ClaimDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetClaimDurationOk() (*Promotion202406SearchCouponsResponseDataCouponsClaimDuration, bool) {
    if o == nil || utils.IsNil(o.ClaimDuration) {
        return nil, false
    }
    return o.ClaimDuration, true
}

// HasClaimDuration returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasClaimDuration() bool {
    if o != nil && !utils.IsNil(o.ClaimDuration) {
        return true
    }

    return false
}

// SetClaimDuration gets a reference to the given Promotion202406SearchCouponsResponseDataCouponsClaimDuration and assigns it to the ClaimDuration field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetClaimDuration(v Promotion202406SearchCouponsResponseDataCouponsClaimDuration) {
    o.ClaimDuration = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetCreateTime() int64 {
    if o == nil || utils.IsNil(o.CreateTime) {
        var ret int64
        return ret
    }
    return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetCreateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTime) {
        return nil, false
    }
    return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasCreateTime() bool {
    if o != nil && !utils.IsNil(o.CreateTime) {
        return true
    }

    return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetCreateTime(v int64) {
    o.CreateTime = &v
}

// GetCreationSource returns the CreationSource field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetCreationSource() string {
    if o == nil || utils.IsNil(o.CreationSource) {
        var ret string
        return ret
    }
    return *o.CreationSource
}

// GetCreationSourceOk returns a tuple with the CreationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetCreationSourceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CreationSource) {
        return nil, false
    }
    return o.CreationSource, true
}

// HasCreationSource returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasCreationSource() bool {
    if o != nil && !utils.IsNil(o.CreationSource) {
        return true
    }

    return false
}

// SetCreationSource gets a reference to the given string and assigns it to the CreationSource field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetCreationSource(v string) {
    o.CreationSource = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetDiscount() Promotion202406SearchCouponsResponseDataCouponsDiscount {
    if o == nil || utils.IsNil(o.Discount) {
        var ret Promotion202406SearchCouponsResponseDataCouponsDiscount
        return ret
    }
    return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetDiscountOk() (*Promotion202406SearchCouponsResponseDataCouponsDiscount, bool) {
    if o == nil || utils.IsNil(o.Discount) {
        return nil, false
    }
    return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasDiscount() bool {
    if o != nil && !utils.IsNil(o.Discount) {
        return true
    }

    return false
}

// SetDiscount gets a reference to the given Promotion202406SearchCouponsResponseDataCouponsDiscount and assigns it to the Discount field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetDiscount(v Promotion202406SearchCouponsResponseDataCouponsDiscount) {
    o.Discount = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetDisplayType() string {
    if o == nil || utils.IsNil(o.DisplayType) {
        var ret string
        return ret
    }
    return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetDisplayTypeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.DisplayType) {
        return nil, false
    }
    return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasDisplayType() bool {
    if o != nil && !utils.IsNil(o.DisplayType) {
        return true
    }

    return false
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetDisplayType(v string) {
    o.DisplayType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetId(v string) {
    o.Id = &v
}

// GetProductScope returns the ProductScope field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetProductScope() string {
    if o == nil || utils.IsNil(o.ProductScope) {
        var ret string
        return ret
    }
    return *o.ProductScope
}

// GetProductScopeOk returns a tuple with the ProductScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetProductScopeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ProductScope) {
        return nil, false
    }
    return o.ProductScope, true
}

// HasProductScope returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasProductScope() bool {
    if o != nil && !utils.IsNil(o.ProductScope) {
        return true
    }

    return false
}

// SetProductScope gets a reference to the given string and assigns it to the ProductScope field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetProductScope(v string) {
    o.ProductScope = &v
}

// GetPromoCode returns the PromoCode field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetPromoCode() string {
    if o == nil || utils.IsNil(o.PromoCode) {
        var ret string
        return ret
    }
    return *o.PromoCode
}

// GetPromoCodeOk returns a tuple with the PromoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetPromoCodeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PromoCode) {
        return nil, false
    }
    return o.PromoCode, true
}

// HasPromoCode returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasPromoCode() bool {
    if o != nil && !utils.IsNil(o.PromoCode) {
        return true
    }

    return false
}

// SetPromoCode gets a reference to the given string and assigns it to the PromoCode field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetPromoCode(v string) {
    o.PromoCode = &v
}

// GetRedemptionDuration returns the RedemptionDuration field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetRedemptionDuration() Promotion202406SearchCouponsResponseDataCouponsRedemptionDuration {
    if o == nil || utils.IsNil(o.RedemptionDuration) {
        var ret Promotion202406SearchCouponsResponseDataCouponsRedemptionDuration
        return ret
    }
    return *o.RedemptionDuration
}

// GetRedemptionDurationOk returns a tuple with the RedemptionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetRedemptionDurationOk() (*Promotion202406SearchCouponsResponseDataCouponsRedemptionDuration, bool) {
    if o == nil || utils.IsNil(o.RedemptionDuration) {
        return nil, false
    }
    return o.RedemptionDuration, true
}

// HasRedemptionDuration returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasRedemptionDuration() bool {
    if o != nil && !utils.IsNil(o.RedemptionDuration) {
        return true
    }

    return false
}

// SetRedemptionDuration gets a reference to the given Promotion202406SearchCouponsResponseDataCouponsRedemptionDuration and assigns it to the RedemptionDuration field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetRedemptionDuration(v Promotion202406SearchCouponsResponseDataCouponsRedemptionDuration) {
    o.RedemptionDuration = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetStatus(v string) {
    o.Status = &v
}

// GetTargetBuyerSegment returns the TargetBuyerSegment field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetTargetBuyerSegment() string {
    if o == nil || utils.IsNil(o.TargetBuyerSegment) {
        var ret string
        return ret
    }
    return *o.TargetBuyerSegment
}

// GetTargetBuyerSegmentOk returns a tuple with the TargetBuyerSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetTargetBuyerSegmentOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TargetBuyerSegment) {
        return nil, false
    }
    return o.TargetBuyerSegment, true
}

// HasTargetBuyerSegment returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasTargetBuyerSegment() bool {
    if o != nil && !utils.IsNil(o.TargetBuyerSegment) {
        return true
    }

    return false
}

// SetTargetBuyerSegment gets a reference to the given string and assigns it to the TargetBuyerSegment field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetTargetBuyerSegment(v string) {
    o.TargetBuyerSegment = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetThreshold() Promotion202406SearchCouponsResponseDataCouponsThreshold {
    if o == nil || utils.IsNil(o.Threshold) {
        var ret Promotion202406SearchCouponsResponseDataCouponsThreshold
        return ret
    }
    return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetThresholdOk() (*Promotion202406SearchCouponsResponseDataCouponsThreshold, bool) {
    if o == nil || utils.IsNil(o.Threshold) {
        return nil, false
    }
    return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasThreshold() bool {
    if o != nil && !utils.IsNil(o.Threshold) {
        return true
    }

    return false
}

// SetThreshold gets a reference to the given Promotion202406SearchCouponsResponseDataCouponsThreshold and assigns it to the Threshold field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetThreshold(v Promotion202406SearchCouponsResponseDataCouponsThreshold) {
    o.Threshold = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetTitle() string {
    if o == nil || utils.IsNil(o.Title) {
        var ret string
        return ret
    }
    return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetTitleOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Title) {
        return nil, false
    }
    return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasTitle() bool {
    if o != nil && !utils.IsNil(o.Title) {
        return true
    }

    return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetTitle(v string) {
    o.Title = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetUpdateTime() int64 {
    if o == nil || utils.IsNil(o.UpdateTime) {
        var ret int64
        return ret
    }
    return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetUpdateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTime) {
        return nil, false
    }
    return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasUpdateTime() bool {
    if o != nil && !utils.IsNil(o.UpdateTime) {
        return true
    }

    return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetUpdateTime(v int64) {
    o.UpdateTime = &v
}

// GetUsageLimits returns the UsageLimits field value if set, zero value otherwise.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetUsageLimits() Promotion202406SearchCouponsResponseDataCouponsUsageLimits {
    if o == nil || utils.IsNil(o.UsageLimits) {
        var ret Promotion202406SearchCouponsResponseDataCouponsUsageLimits
        return ret
    }
    return *o.UsageLimits
}

// GetUsageLimitsOk returns a tuple with the UsageLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) GetUsageLimitsOk() (*Promotion202406SearchCouponsResponseDataCouponsUsageLimits, bool) {
    if o == nil || utils.IsNil(o.UsageLimits) {
        return nil, false
    }
    return o.UsageLimits, true
}

// HasUsageLimits returns a boolean if a field has been set.
func (o *Promotion202406SearchCouponsResponseDataCoupons) HasUsageLimits() bool {
    if o != nil && !utils.IsNil(o.UsageLimits) {
        return true
    }

    return false
}

// SetUsageLimits gets a reference to the given Promotion202406SearchCouponsResponseDataCouponsUsageLimits and assigns it to the UsageLimits field.
func (o *Promotion202406SearchCouponsResponseDataCoupons) SetUsageLimits(v Promotion202406SearchCouponsResponseDataCouponsUsageLimits) {
    o.UsageLimits = &v
}

func (o Promotion202406SearchCouponsResponseDataCoupons) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Promotion202406SearchCouponsResponseDataCoupons) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ClaimDuration) {
        toSerialize["claim_duration"] = o.ClaimDuration
    }
    if !utils.IsNil(o.CreateTime) {
        toSerialize["create_time"] = o.CreateTime
    }
    if !utils.IsNil(o.CreationSource) {
        toSerialize["creation_source"] = o.CreationSource
    }
    if !utils.IsNil(o.Discount) {
        toSerialize["discount"] = o.Discount
    }
    if !utils.IsNil(o.DisplayType) {
        toSerialize["display_type"] = o.DisplayType
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.ProductScope) {
        toSerialize["product_scope"] = o.ProductScope
    }
    if !utils.IsNil(o.PromoCode) {
        toSerialize["promo_code"] = o.PromoCode
    }
    if !utils.IsNil(o.RedemptionDuration) {
        toSerialize["redemption_duration"] = o.RedemptionDuration
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.TargetBuyerSegment) {
        toSerialize["target_buyer_segment"] = o.TargetBuyerSegment
    }
    if !utils.IsNil(o.Threshold) {
        toSerialize["threshold"] = o.Threshold
    }
    if !utils.IsNil(o.Title) {
        toSerialize["title"] = o.Title
    }
    if !utils.IsNil(o.UpdateTime) {
        toSerialize["update_time"] = o.UpdateTime
    }
    if !utils.IsNil(o.UsageLimits) {
        toSerialize["usage_limits"] = o.UsageLimits
    }
    return toSerialize, nil
}

type NullablePromotion202406SearchCouponsResponseDataCoupons struct {
	value *Promotion202406SearchCouponsResponseDataCoupons
	isSet bool
}

func (v NullablePromotion202406SearchCouponsResponseDataCoupons) Get() *Promotion202406SearchCouponsResponseDataCoupons {
	return v.value
}

func (v *NullablePromotion202406SearchCouponsResponseDataCoupons) Set(val *Promotion202406SearchCouponsResponseDataCoupons) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion202406SearchCouponsResponseDataCoupons) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion202406SearchCouponsResponseDataCoupons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion202406SearchCouponsResponseDataCoupons(val *Promotion202406SearchCouponsResponseDataCoupons) *NullablePromotion202406SearchCouponsResponseDataCoupons {
	return &NullablePromotion202406SearchCouponsResponseDataCoupons{value: val, isSet: true}
}

func (v NullablePromotion202406SearchCouponsResponseDataCoupons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion202406SearchCouponsResponseDataCoupons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


