/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_seller_v202409

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct{}

// AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct struct for AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct
type AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct struct {
    // Product's unique id.
    Id *string `json:"id,omitempty"`
    // The inventory of this product.
    Inventory *int64 `json:"inventory,omitempty"`
    // The product image url.
    MainImageUrl *string `json:"main_image_url,omitempty"`
    OriginalPrice *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductOriginalPrice `json:"original_price,omitempty"`
    // Product's status. Field values: - LIVE: When the product is normal for sale, return to the LIVE status - OUT_OF_STOCK: When the product is out of stock for the consumer, the OUT_OF_STOCK state is returned - SELLER_DEACTIVATE:  When the product is deactivated by the merchant, the SELLER_DEACTIVATE status is returned - PLATFORM_DEACTIVATE: When the product is deactivated by the platform or is not available for sale, the PLATFORM_DEACTIVATE status is returned - GNE_REJECT: When the product is governed or the open collaboration is dismissed, the GNE_REJECT state is returned - DELETE: When the product is deleted, the DELETE status is returned - OTHER: When the product is in an unsaleable state, such as draft, frozen, review, etc, the OTHER status is returned
    Status *string `json:"status,omitempty"`
    // Product's name.
    Title *string `json:"title,omitempty"`
}

// NewAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct instantiates a new AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct() *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct {
    this := AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct{}
    return &this
}

// NewAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductWithDefaults instantiates a new AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductWithDefaults() *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct {
    this := AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) SetId(v string) {
    o.Id = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetInventory() int64 {
    if o == nil || utils.IsNil(o.Inventory) {
        var ret int64
        return ret
    }
    return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetInventoryOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.Inventory) {
        return nil, false
    }
    return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) HasInventory() bool {
    if o != nil && !utils.IsNil(o.Inventory) {
        return true
    }

    return false
}

// SetInventory gets a reference to the given int64 and assigns it to the Inventory field.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) SetInventory(v int64) {
    o.Inventory = &v
}

// GetMainImageUrl returns the MainImageUrl field value if set, zero value otherwise.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetMainImageUrl() string {
    if o == nil || utils.IsNil(o.MainImageUrl) {
        var ret string
        return ret
    }
    return *o.MainImageUrl
}

// GetMainImageUrlOk returns a tuple with the MainImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetMainImageUrlOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MainImageUrl) {
        return nil, false
    }
    return o.MainImageUrl, true
}

// HasMainImageUrl returns a boolean if a field has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) HasMainImageUrl() bool {
    if o != nil && !utils.IsNil(o.MainImageUrl) {
        return true
    }

    return false
}

// SetMainImageUrl gets a reference to the given string and assigns it to the MainImageUrl field.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) SetMainImageUrl(v string) {
    o.MainImageUrl = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetOriginalPrice() AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductOriginalPrice {
    if o == nil || utils.IsNil(o.OriginalPrice) {
        var ret AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductOriginalPrice
        return ret
    }
    return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetOriginalPriceOk() (*AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductOriginalPrice, bool) {
    if o == nil || utils.IsNil(o.OriginalPrice) {
        return nil, false
    }
    return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) HasOriginalPrice() bool {
    if o != nil && !utils.IsNil(o.OriginalPrice) {
        return true
    }

    return false
}

// SetOriginalPrice gets a reference to the given AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductOriginalPrice and assigns it to the OriginalPrice field.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) SetOriginalPrice(v AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProductOriginalPrice) {
    o.OriginalPrice = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) SetStatus(v string) {
    o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetTitle() string {
    if o == nil || utils.IsNil(o.Title) {
        var ret string
        return ret
    }
    return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) GetTitleOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Title) {
        return nil, false
    }
    return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) HasTitle() bool {
    if o != nil && !utils.IsNil(o.Title) {
        return true
    }

    return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) SetTitle(v string) {
    o.Title = &v
}

func (o AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Inventory) {
        toSerialize["inventory"] = o.Inventory
    }
    if !utils.IsNil(o.MainImageUrl) {
        toSerialize["main_image_url"] = o.MainImageUrl
    }
    if !utils.IsNil(o.OriginalPrice) {
        toSerialize["original_price"] = o.OriginalPrice
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.Title) {
        toSerialize["title"] = o.Title
    }
    return toSerialize, nil
}

type NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct struct {
	value *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct
	isSet bool
}

func (v NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) Get() *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct {
	return v.value
}

func (v *NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) Set(val *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct(val *AffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) *NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct {
	return &NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct{value: val, isSet: true}
}

func (v NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliateSeller202409SearchOpenCollaborationResponseDataOpenCollaborationsProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


