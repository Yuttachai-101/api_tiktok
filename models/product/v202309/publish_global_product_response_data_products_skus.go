/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Product202309PublishGlobalProductResponseDataProductsSkus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309PublishGlobalProductResponseDataProductsSkus{}

// Product202309PublishGlobalProductResponseDataProductsSkus struct for Product202309PublishGlobalProductResponseDataProductsSkus
type Product202309PublishGlobalProductResponseDataProductsSkus struct {
    // The newly generated local SKU ID.
    Id *string `json:"id,omitempty"`
    // The associated global SKU ID.
    RelatedGlobalSkuId *string `json:"related_global_sku_id,omitempty"`
    // A list of attributes  (e.g. size, color, length) that define each variant of a product.
    SaleAttributes []Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes `json:"sale_attributes,omitempty"`
    // An internal code/name for managing SKUs, not visible to buyers.
    SellerSku *string `json:"seller_sku,omitempty"`
}

// NewProduct202309PublishGlobalProductResponseDataProductsSkus instantiates a new Product202309PublishGlobalProductResponseDataProductsSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309PublishGlobalProductResponseDataProductsSkus() *Product202309PublishGlobalProductResponseDataProductsSkus {
    this := Product202309PublishGlobalProductResponseDataProductsSkus{}
    return &this
}

// NewProduct202309PublishGlobalProductResponseDataProductsSkusWithDefaults instantiates a new Product202309PublishGlobalProductResponseDataProductsSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309PublishGlobalProductResponseDataProductsSkusWithDefaults() *Product202309PublishGlobalProductResponseDataProductsSkus {
    this := Product202309PublishGlobalProductResponseDataProductsSkus{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) SetId(v string) {
    o.Id = &v
}

// GetRelatedGlobalSkuId returns the RelatedGlobalSkuId field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetRelatedGlobalSkuId() string {
    if o == nil || utils.IsNil(o.RelatedGlobalSkuId) {
        var ret string
        return ret
    }
    return *o.RelatedGlobalSkuId
}

// GetRelatedGlobalSkuIdOk returns a tuple with the RelatedGlobalSkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetRelatedGlobalSkuIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RelatedGlobalSkuId) {
        return nil, false
    }
    return o.RelatedGlobalSkuId, true
}

// HasRelatedGlobalSkuId returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) HasRelatedGlobalSkuId() bool {
    if o != nil && !utils.IsNil(o.RelatedGlobalSkuId) {
        return true
    }

    return false
}

// SetRelatedGlobalSkuId gets a reference to the given string and assigns it to the RelatedGlobalSkuId field.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) SetRelatedGlobalSkuId(v string) {
    o.RelatedGlobalSkuId = &v
}

// GetSaleAttributes returns the SaleAttributes field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetSaleAttributes() []Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes {
    if o == nil || utils.IsNil(o.SaleAttributes) {
        var ret []Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes
        return ret
    }
    return o.SaleAttributes
}

// GetSaleAttributesOk returns a tuple with the SaleAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetSaleAttributesOk() ([]Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes, bool) {
    if o == nil || utils.IsNil(o.SaleAttributes) {
        return nil, false
    }
    return o.SaleAttributes, true
}

// HasSaleAttributes returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) HasSaleAttributes() bool {
    if o != nil && !utils.IsNil(o.SaleAttributes) {
        return true
    }

    return false
}

// SetSaleAttributes gets a reference to the given []Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes and assigns it to the SaleAttributes field.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) SetSaleAttributes(v []Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) {
    o.SaleAttributes = v
}

// GetSellerSku returns the SellerSku field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetSellerSku() string {
    if o == nil || utils.IsNil(o.SellerSku) {
        var ret string
        return ret
    }
    return *o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) GetSellerSkuOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SellerSku) {
        return nil, false
    }
    return o.SellerSku, true
}

// HasSellerSku returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) HasSellerSku() bool {
    if o != nil && !utils.IsNil(o.SellerSku) {
        return true
    }

    return false
}

// SetSellerSku gets a reference to the given string and assigns it to the SellerSku field.
func (o *Product202309PublishGlobalProductResponseDataProductsSkus) SetSellerSku(v string) {
    o.SellerSku = &v
}

func (o Product202309PublishGlobalProductResponseDataProductsSkus) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309PublishGlobalProductResponseDataProductsSkus) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.RelatedGlobalSkuId) {
        toSerialize["related_global_sku_id"] = o.RelatedGlobalSkuId
    }
    if !utils.IsNil(o.SaleAttributes) {
        toSerialize["sale_attributes"] = o.SaleAttributes
    }
    if !utils.IsNil(o.SellerSku) {
        toSerialize["seller_sku"] = o.SellerSku
    }
    return toSerialize, nil
}

type NullableProduct202309PublishGlobalProductResponseDataProductsSkus struct {
	value *Product202309PublishGlobalProductResponseDataProductsSkus
	isSet bool
}

func (v NullableProduct202309PublishGlobalProductResponseDataProductsSkus) Get() *Product202309PublishGlobalProductResponseDataProductsSkus {
	return v.value
}

func (v *NullableProduct202309PublishGlobalProductResponseDataProductsSkus) Set(val *Product202309PublishGlobalProductResponseDataProductsSkus) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309PublishGlobalProductResponseDataProductsSkus) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309PublishGlobalProductResponseDataProductsSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309PublishGlobalProductResponseDataProductsSkus(val *Product202309PublishGlobalProductResponseDataProductsSkus) *NullableProduct202309PublishGlobalProductResponseDataProductsSkus {
	return &NullableProduct202309PublishGlobalProductResponseDataProductsSkus{value: val, isSet: true}
}

func (v NullableProduct202309PublishGlobalProductResponseDataProductsSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309PublishGlobalProductResponseDataProductsSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


