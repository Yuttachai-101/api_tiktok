/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems{}

// AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems struct for AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems
type AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems struct {
    // Line item ID
    Id *string `json:"id,omitempty"`
    // Product ID in the order. We return the product ID, because within an order, there could be multiple products (one, some, or all) which are affiliate commission earning products.
    ProductId *string `json:"product_id,omitempty"`
}

// NewAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems instantiates a new AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems() *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems {
    this := AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems{}
    return &this
}

// NewAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItemsWithDefaults instantiates a new AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItemsWithDefaults() *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems {
    this := AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) SetId(v string) {
    o.Id = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) GetProductId() string {
    if o == nil || utils.IsNil(o.ProductId) {
        var ret string
        return ret
    }
    return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) GetProductIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ProductId) {
        return nil, false
    }
    return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) HasProductId() bool {
    if o != nil && !utils.IsNil(o.ProductId) {
        return true
    }

    return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) SetProductId(v string) {
    o.ProductId = &v
}

func (o AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.ProductId) {
        toSerialize["product_id"] = o.ProductId
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems struct {
	value *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems
	isSet bool
}

func (v NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) Get() *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems {
	return v.value
}

func (v *NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) Set(val *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems(val *AffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) *NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems {
	return &NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems{value: val, isSet: true}
}

func (v NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202405SearchSellerAffiliateOrdersResponseDataOrdersLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


