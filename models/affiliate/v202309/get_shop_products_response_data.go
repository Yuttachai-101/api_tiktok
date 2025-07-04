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

            // checks if the Affiliate202309GetShopProductsResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Affiliate202309GetShopProductsResponseData{}

// Affiliate202309GetShopProductsResponseData struct for Affiliate202309GetShopProductsResponseData
type Affiliate202309GetShopProductsResponseData struct {
    // The pagination token is a cursor used for pagination. The token is returned in the previous pagination query to determine the current position. It will be empty when there aren't any products to search for.
    NextPageToken *string `json:"next_page_token,omitempty"`
    // The searched product list. It will be empty when there are no search results.
    Products []Affiliate202309GetShopProductsResponseDataProducts `json:"products,omitempty"`
}

// NewAffiliate202309GetShopProductsResponseData instantiates a new Affiliate202309GetShopProductsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliate202309GetShopProductsResponseData() *Affiliate202309GetShopProductsResponseData {
    this := Affiliate202309GetShopProductsResponseData{}
    return &this
}

// NewAffiliate202309GetShopProductsResponseDataWithDefaults instantiates a new Affiliate202309GetShopProductsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliate202309GetShopProductsResponseDataWithDefaults() *Affiliate202309GetShopProductsResponseData {
    this := Affiliate202309GetShopProductsResponseData{}
    return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseData) GetNextPageToken() string {
    if o == nil || utils.IsNil(o.NextPageToken) {
        var ret string
        return ret
    }
    return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseData) GetNextPageTokenOk() (*string, bool) {
    if o == nil || utils.IsNil(o.NextPageToken) {
        return nil, false
    }
    return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseData) HasNextPageToken() bool {
    if o != nil && !utils.IsNil(o.NextPageToken) {
        return true
    }

    return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Affiliate202309GetShopProductsResponseData) SetNextPageToken(v string) {
    o.NextPageToken = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Affiliate202309GetShopProductsResponseData) GetProducts() []Affiliate202309GetShopProductsResponseDataProducts {
    if o == nil || utils.IsNil(o.Products) {
        var ret []Affiliate202309GetShopProductsResponseDataProducts
        return ret
    }
    return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affiliate202309GetShopProductsResponseData) GetProductsOk() ([]Affiliate202309GetShopProductsResponseDataProducts, bool) {
    if o == nil || utils.IsNil(o.Products) {
        return nil, false
    }
    return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Affiliate202309GetShopProductsResponseData) HasProducts() bool {
    if o != nil && !utils.IsNil(o.Products) {
        return true
    }

    return false
}

// SetProducts gets a reference to the given []Affiliate202309GetShopProductsResponseDataProducts and assigns it to the Products field.
func (o *Affiliate202309GetShopProductsResponseData) SetProducts(v []Affiliate202309GetShopProductsResponseDataProducts) {
    o.Products = v
}

func (o Affiliate202309GetShopProductsResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Affiliate202309GetShopProductsResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.NextPageToken) {
        toSerialize["next_page_token"] = o.NextPageToken
    }
    if !utils.IsNil(o.Products) {
        toSerialize["products"] = o.Products
    }
    return toSerialize, nil
}

type NullableAffiliate202309GetShopProductsResponseData struct {
	value *Affiliate202309GetShopProductsResponseData
	isSet bool
}

func (v NullableAffiliate202309GetShopProductsResponseData) Get() *Affiliate202309GetShopProductsResponseData {
	return v.value
}

func (v *NullableAffiliate202309GetShopProductsResponseData) Set(val *Affiliate202309GetShopProductsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliate202309GetShopProductsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliate202309GetShopProductsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliate202309GetShopProductsResponseData(val *Affiliate202309GetShopProductsResponseData) *NullableAffiliate202309GetShopProductsResponseData {
	return &NullableAffiliate202309GetShopProductsResponseData{value: val, isSet: true}
}

func (v NullableAffiliate202309GetShopProductsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliate202309GetShopProductsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


