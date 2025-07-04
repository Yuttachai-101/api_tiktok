/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_creator_v202412

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList{}

// AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList struct for AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList
type AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList struct {
    // The SKU identifier.
    Id *string `json:"id,omitempty"`
    // If this SKU is available.
    IsAvailable *bool `json:"is_available,omitempty"`
    Price *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListPrice `json:"price,omitempty"`
    // The SKU property information.
    SaleProperties []AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListSaleProperties `json:"sale_properties,omitempty"`
    // The combination of SKU properties for this SKU.
    SalePropertyValueIds *string `json:"sale_property_value_ids,omitempty"`
    // The reason why the SKU is unavailable:  - IS_PREORDER : this product is a preorder product which does not support free sample - IS_GIFT: this product is a gift product which does not support free sample - OUT_OF_STOCK: product sold out - EXCEED_CB_PRICE_THRESHOLD:  - ALREADY_APPLYED: creator has already applied this SKU.
    UnavailableReason *string `json:"unavailable_reason,omitempty"`
}

// NewAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList instantiates a new AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList() *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList {
    this := AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList{}
    return &this
}

// NewAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListWithDefaults instantiates a new AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListWithDefaults() *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList {
    this := AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) SetId(v string) {
    o.Id = &v
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetIsAvailable() bool {
    if o == nil || utils.IsNil(o.IsAvailable) {
        var ret bool
        return ret
    }
    return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetIsAvailableOk() (*bool, bool) {
    if o == nil || utils.IsNil(o.IsAvailable) {
        return nil, false
    }
    return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) HasIsAvailable() bool {
    if o != nil && !utils.IsNil(o.IsAvailable) {
        return true
    }

    return false
}

// SetIsAvailable gets a reference to the given bool and assigns it to the IsAvailable field.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) SetIsAvailable(v bool) {
    o.IsAvailable = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetPrice() AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListPrice {
    if o == nil || utils.IsNil(o.Price) {
        var ret AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListPrice
        return ret
    }
    return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetPriceOk() (*AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListPrice, bool) {
    if o == nil || utils.IsNil(o.Price) {
        return nil, false
    }
    return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) HasPrice() bool {
    if o != nil && !utils.IsNil(o.Price) {
        return true
    }

    return false
}

// SetPrice gets a reference to the given AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListPrice and assigns it to the Price field.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) SetPrice(v AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListPrice) {
    o.Price = &v
}

// GetSaleProperties returns the SaleProperties field value if set, zero value otherwise.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetSaleProperties() []AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListSaleProperties {
    if o == nil || utils.IsNil(o.SaleProperties) {
        var ret []AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListSaleProperties
        return ret
    }
    return o.SaleProperties
}

// GetSalePropertiesOk returns a tuple with the SaleProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetSalePropertiesOk() ([]AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListSaleProperties, bool) {
    if o == nil || utils.IsNil(o.SaleProperties) {
        return nil, false
    }
    return o.SaleProperties, true
}

// HasSaleProperties returns a boolean if a field has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) HasSaleProperties() bool {
    if o != nil && !utils.IsNil(o.SaleProperties) {
        return true
    }

    return false
}

// SetSaleProperties gets a reference to the given []AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListSaleProperties and assigns it to the SaleProperties field.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) SetSaleProperties(v []AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuListSaleProperties) {
    o.SaleProperties = v
}

// GetSalePropertyValueIds returns the SalePropertyValueIds field value if set, zero value otherwise.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetSalePropertyValueIds() string {
    if o == nil || utils.IsNil(o.SalePropertyValueIds) {
        var ret string
        return ret
    }
    return *o.SalePropertyValueIds
}

// GetSalePropertyValueIdsOk returns a tuple with the SalePropertyValueIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetSalePropertyValueIdsOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SalePropertyValueIds) {
        return nil, false
    }
    return o.SalePropertyValueIds, true
}

// HasSalePropertyValueIds returns a boolean if a field has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) HasSalePropertyValueIds() bool {
    if o != nil && !utils.IsNil(o.SalePropertyValueIds) {
        return true
    }

    return false
}

// SetSalePropertyValueIds gets a reference to the given string and assigns it to the SalePropertyValueIds field.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) SetSalePropertyValueIds(v string) {
    o.SalePropertyValueIds = &v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetUnavailableReason() string {
    if o == nil || utils.IsNil(o.UnavailableReason) {
        var ret string
        return ret
    }
    return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) GetUnavailableReasonOk() (*string, bool) {
    if o == nil || utils.IsNil(o.UnavailableReason) {
        return nil, false
    }
    return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) HasUnavailableReason() bool {
    if o != nil && !utils.IsNil(o.UnavailableReason) {
        return true
    }

    return false
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) SetUnavailableReason(v string) {
    o.UnavailableReason = &v
}

func (o AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.IsAvailable) {
        toSerialize["is_available"] = o.IsAvailable
    }
    if !utils.IsNil(o.Price) {
        toSerialize["price"] = o.Price
    }
    if !utils.IsNil(o.SaleProperties) {
        toSerialize["sale_properties"] = o.SaleProperties
    }
    if !utils.IsNil(o.SalePropertyValueIds) {
        toSerialize["sale_property_value_ids"] = o.SalePropertyValueIds
    }
    if !utils.IsNil(o.UnavailableReason) {
        toSerialize["unavailable_reason"] = o.UnavailableReason
    }
    return toSerialize, nil
}

type NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList struct {
	value *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList
	isSet bool
}

func (v NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) Get() *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList {
	return v.value
}

func (v *NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) Set(val *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList(val *AffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) *NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList {
	return &NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList{value: val, isSet: true}
}

func (v NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateCreator202412GetCreatorApplicableSampleLabelResponseDataLabelSampleProductSampleSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


