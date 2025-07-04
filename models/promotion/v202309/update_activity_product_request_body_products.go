/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package promotion_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Promotion202309UpdateActivityProductRequestBodyProducts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Promotion202309UpdateActivityProductRequestBodyProducts{}

// Promotion202309UpdateActivityProductRequestBodyProducts struct for Promotion202309UpdateActivityProductRequestBodyProducts
type Promotion202309UpdateActivityProductRequestBodyProducts struct {
    // Deal price.  You must specify the value when `product_level==PRODUCT` and `activity_type==FIXED_PRICE / FLASHSALE`. The currency is the same between activity price and product price.
    ActivityPriceAmount *string `json:"activity_price_amount,omitempty"`
    // Discount value. If the product is 10% off, the value is `10`. You must specify the value when `product_level==PRODUCT` and `activity_type==DIRECT_DISCOUNT`; and you must not specify it when it's not.
    Discount *string `json:"discount,omitempty"`
    // TikTok Shop product ID
    Id *string `json:"id,omitempty"`
    // The quantity limit of the products involved in the activity. The range is `[1, 99]`, or you can use `-1` for unlimited. If you are updating the value of an existing product, the value cannot be less than the current value. When `product_level==VARIATION`, you must specify `-1`.
    QuantityLimit *int32 `json:"quantity_limit,omitempty"`
    // Limit of product purchase per buyer. The range is `[1, 99]`, or you can use `-1` for unlimited. If you are updating the value of an existing product, the value cannot be less than the current value. When `product_level==VARIATION`, you must specify `-1`.
    QuantityPerUser *int32 `json:"quantity_per_user,omitempty"`
    // The SKUs to add to the list or to edit. The number of the SKUs across all products must not exceed 300 in an API call. You must specify the value to `[]` when `product_level==PRODUCT`.
    Skus []Promotion202309UpdateActivityProductRequestBodyProductsSkus `json:"skus,omitempty"`
}

// NewPromotion202309UpdateActivityProductRequestBodyProducts instantiates a new Promotion202309UpdateActivityProductRequestBodyProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion202309UpdateActivityProductRequestBodyProducts() *Promotion202309UpdateActivityProductRequestBodyProducts {
    this := Promotion202309UpdateActivityProductRequestBodyProducts{}
    return &this
}

// NewPromotion202309UpdateActivityProductRequestBodyProductsWithDefaults instantiates a new Promotion202309UpdateActivityProductRequestBodyProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotion202309UpdateActivityProductRequestBodyProductsWithDefaults() *Promotion202309UpdateActivityProductRequestBodyProducts {
    this := Promotion202309UpdateActivityProductRequestBodyProducts{}
    return &this
}

// GetActivityPriceAmount returns the ActivityPriceAmount field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetActivityPriceAmount() string {
    if o == nil || utils.IsNil(o.ActivityPriceAmount) {
        var ret string
        return ret
    }
    return *o.ActivityPriceAmount
}

// GetActivityPriceAmountOk returns a tuple with the ActivityPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetActivityPriceAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ActivityPriceAmount) {
        return nil, false
    }
    return o.ActivityPriceAmount, true
}

// HasActivityPriceAmount returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) HasActivityPriceAmount() bool {
    if o != nil && !utils.IsNil(o.ActivityPriceAmount) {
        return true
    }

    return false
}

// SetActivityPriceAmount gets a reference to the given string and assigns it to the ActivityPriceAmount field.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) SetActivityPriceAmount(v string) {
    o.ActivityPriceAmount = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetDiscount() string {
    if o == nil || utils.IsNil(o.Discount) {
        var ret string
        return ret
    }
    return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetDiscountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Discount) {
        return nil, false
    }
    return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) HasDiscount() bool {
    if o != nil && !utils.IsNil(o.Discount) {
        return true
    }

    return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) SetDiscount(v string) {
    o.Discount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) SetId(v string) {
    o.Id = &v
}

// GetQuantityLimit returns the QuantityLimit field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetQuantityLimit() int32 {
    if o == nil || utils.IsNil(o.QuantityLimit) {
        var ret int32
        return ret
    }
    return *o.QuantityLimit
}

// GetQuantityLimitOk returns a tuple with the QuantityLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetQuantityLimitOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.QuantityLimit) {
        return nil, false
    }
    return o.QuantityLimit, true
}

// HasQuantityLimit returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) HasQuantityLimit() bool {
    if o != nil && !utils.IsNil(o.QuantityLimit) {
        return true
    }

    return false
}

// SetQuantityLimit gets a reference to the given int32 and assigns it to the QuantityLimit field.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) SetQuantityLimit(v int32) {
    o.QuantityLimit = &v
}

// GetQuantityPerUser returns the QuantityPerUser field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetQuantityPerUser() int32 {
    if o == nil || utils.IsNil(o.QuantityPerUser) {
        var ret int32
        return ret
    }
    return *o.QuantityPerUser
}

// GetQuantityPerUserOk returns a tuple with the QuantityPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetQuantityPerUserOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.QuantityPerUser) {
        return nil, false
    }
    return o.QuantityPerUser, true
}

// HasQuantityPerUser returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) HasQuantityPerUser() bool {
    if o != nil && !utils.IsNil(o.QuantityPerUser) {
        return true
    }

    return false
}

// SetQuantityPerUser gets a reference to the given int32 and assigns it to the QuantityPerUser field.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) SetQuantityPerUser(v int32) {
    o.QuantityPerUser = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetSkus() []Promotion202309UpdateActivityProductRequestBodyProductsSkus {
    if o == nil || utils.IsNil(o.Skus) {
        var ret []Promotion202309UpdateActivityProductRequestBodyProductsSkus
        return ret
    }
    return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) GetSkusOk() ([]Promotion202309UpdateActivityProductRequestBodyProductsSkus, bool) {
    if o == nil || utils.IsNil(o.Skus) {
        return nil, false
    }
    return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) HasSkus() bool {
    if o != nil && !utils.IsNil(o.Skus) {
        return true
    }

    return false
}

// SetSkus gets a reference to the given []Promotion202309UpdateActivityProductRequestBodyProductsSkus and assigns it to the Skus field.
func (o *Promotion202309UpdateActivityProductRequestBodyProducts) SetSkus(v []Promotion202309UpdateActivityProductRequestBodyProductsSkus) {
    o.Skus = v
}

func (o Promotion202309UpdateActivityProductRequestBodyProducts) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Promotion202309UpdateActivityProductRequestBodyProducts) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ActivityPriceAmount) {
        toSerialize["activity_price_amount"] = o.ActivityPriceAmount
    }
    if !utils.IsNil(o.Discount) {
        toSerialize["discount"] = o.Discount
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.QuantityLimit) {
        toSerialize["quantity_limit"] = o.QuantityLimit
    }
    if !utils.IsNil(o.QuantityPerUser) {
        toSerialize["quantity_per_user"] = o.QuantityPerUser
    }
    if !utils.IsNil(o.Skus) {
        toSerialize["skus"] = o.Skus
    }
    return toSerialize, nil
}

type NullablePromotion202309UpdateActivityProductRequestBodyProducts struct {
	value *Promotion202309UpdateActivityProductRequestBodyProducts
	isSet bool
}

func (v NullablePromotion202309UpdateActivityProductRequestBodyProducts) Get() *Promotion202309UpdateActivityProductRequestBodyProducts {
	return v.value
}

func (v *NullablePromotion202309UpdateActivityProductRequestBodyProducts) Set(val *Promotion202309UpdateActivityProductRequestBodyProducts) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion202309UpdateActivityProductRequestBodyProducts) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion202309UpdateActivityProductRequestBodyProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion202309UpdateActivityProductRequestBodyProducts(val *Promotion202309UpdateActivityProductRequestBodyProducts) *NullablePromotion202309UpdateActivityProductRequestBodyProducts {
	return &NullablePromotion202309UpdateActivityProductRequestBodyProducts{value: val, isSet: true}
}

func (v NullablePromotion202309UpdateActivityProductRequestBodyProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion202309UpdateActivityProductRequestBodyProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


