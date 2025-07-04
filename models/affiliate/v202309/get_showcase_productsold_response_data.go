/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Affiliate202309GetShowcaseProductsoldResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Affiliate202309GetShowcaseProductsoldResponseData{}

// Affiliate202309GetShowcaseProductsoldResponseData struct for Affiliate202309GetShowcaseProductsoldResponseData
type Affiliate202309GetShowcaseProductsoldResponseData struct {
    // The product IDs in the livebag among the products returned in this response.
    LiveProductIds []string `json:"live_product_ids,omitempty"`
    // Returns the token to get the next page of products if the response does not return all the products, otherwise returns an empty string.
    NextPageToken *string `json:"next_page_token,omitempty"`
    // The product's detailed info fields
    Products []Affiliate202309GetShowcaseProductsoldResponseDataProducts `json:"products,omitempty"`
    // Returns the total number of products in the showcase.
    TotalCount *int32 `json:"total_count,omitempty"`
}

// NewAffiliate202309GetShowcaseProductsoldResponseData instantiates a new Affiliate202309GetShowcaseProductsoldResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliate202309GetShowcaseProductsoldResponseData() *Affiliate202309GetShowcaseProductsoldResponseData {
    this := Affiliate202309GetShowcaseProductsoldResponseData{}
    return &this
}

// NewAffiliate202309GetShowcaseProductsoldResponseDataWithDefaults instantiates a new Affiliate202309GetShowcaseProductsoldResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliate202309GetShowcaseProductsoldResponseDataWithDefaults() *Affiliate202309GetShowcaseProductsoldResponseData {
    this := Affiliate202309GetShowcaseProductsoldResponseData{}
    return &this
}

// GetLiveProductIds returns the LiveProductIds field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetLiveProductIds() []string {
    if o == nil || utils.IsNil(o.LiveProductIds) {
        var ret []string
        return ret
    }
    return o.LiveProductIds
}

// GetLiveProductIdsOk returns a tuple with the LiveProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetLiveProductIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.LiveProductIds) {
        return nil, false
    }
    return o.LiveProductIds, true
}

// HasLiveProductIds returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) HasLiveProductIds() bool {
    if o != nil && !utils.IsNil(o.LiveProductIds) {
        return true
    }

    return false
}

// SetLiveProductIds gets a reference to the given []string and assigns it to the LiveProductIds field.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) SetLiveProductIds(v []string) {
    o.LiveProductIds = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetNextPageToken() string {
    if o == nil || utils.IsNil(o.NextPageToken) {
        var ret string
        return ret
    }
    return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetNextPageTokenOk() (*string, bool) {
    if o == nil || utils.IsNil(o.NextPageToken) {
        return nil, false
    }
    return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) HasNextPageToken() bool {
    if o != nil && !utils.IsNil(o.NextPageToken) {
        return true
    }

    return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) SetNextPageToken(v string) {
    o.NextPageToken = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetProducts() []Affiliate202309GetShowcaseProductsoldResponseDataProducts {
    if o == nil || utils.IsNil(o.Products) {
        var ret []Affiliate202309GetShowcaseProductsoldResponseDataProducts
        return ret
    }
    return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetProductsOk() ([]Affiliate202309GetShowcaseProductsoldResponseDataProducts, bool) {
    if o == nil || utils.IsNil(o.Products) {
        return nil, false
    }
    return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) HasProducts() bool {
    if o != nil && !utils.IsNil(o.Products) {
        return true
    }

    return false
}

// SetProducts gets a reference to the given []Affiliate202309GetShowcaseProductsoldResponseDataProducts and assigns it to the Products field.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) SetProducts(v []Affiliate202309GetShowcaseProductsoldResponseDataProducts) {
    o.Products = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetTotalCount() int32 {
    if o == nil || utils.IsNil(o.TotalCount) {
        var ret int32
        return ret
    }
    return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) GetTotalCountOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.TotalCount) {
        return nil, false
    }
    return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) HasTotalCount() bool {
    if o != nil && !utils.IsNil(o.TotalCount) {
        return true
    }

    return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *Affiliate202309GetShowcaseProductsoldResponseData) SetTotalCount(v int32) {
    o.TotalCount = &v
}

func (o Affiliate202309GetShowcaseProductsoldResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Affiliate202309GetShowcaseProductsoldResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.LiveProductIds) {
        toSerialize["live_product_ids"] = o.LiveProductIds
    }
    if !utils.IsNil(o.NextPageToken) {
        toSerialize["next_page_token"] = o.NextPageToken
    }
    if !utils.IsNil(o.Products) {
        toSerialize["products"] = o.Products
    }
    if !utils.IsNil(o.TotalCount) {
        toSerialize["total_count"] = o.TotalCount
    }
    return toSerialize, nil
}

type NullableAffiliate202309GetShowcaseProductsoldResponseData struct {
	value *Affiliate202309GetShowcaseProductsoldResponseData
	isSet bool
}

func (v NullableAffiliate202309GetShowcaseProductsoldResponseData) Get() *Affiliate202309GetShowcaseProductsoldResponseData {
	return v.value
}

func (v *NullableAffiliate202309GetShowcaseProductsoldResponseData) Set(val *Affiliate202309GetShowcaseProductsoldResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliate202309GetShowcaseProductsoldResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliate202309GetShowcaseProductsoldResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliate202309GetShowcaseProductsoldResponseData(val *Affiliate202309GetShowcaseProductsoldResponseData) *NullableAffiliate202309GetShowcaseProductsoldResponseData {
	return &NullableAffiliate202309GetShowcaseProductsoldResponseData{value: val, isSet: true}
}

func (v NullableAffiliate202309GetShowcaseProductsoldResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliate202309GetShowcaseProductsoldResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


