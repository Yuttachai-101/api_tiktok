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

            // checks if the Product202309GetProductResponseDataSkusPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetProductResponseDataSkusPrice{}

// Product202309GetProductResponseDataSkusPrice struct for Product202309GetProductResponseDataSkusPrice
type Product202309GetProductResponseDataSkusPrice struct {
    // The currency. Possible values based on the region: - BRL:  Brazil - EUR: France, Germany, Ireland, Italy, Spain - GBP: United Kingdom - IDR: Indonesia - JPY: Japan - MXN: Mexico - MYR: Malaysia - PHP: Philippines - SGD: Singapore - THB: Thailand - USD: United States - VND: Vietnam
    Currency *string `json:"currency,omitempty"`
    // **All sellers** The SKU's **local display price** shown on the product page before any discounts.
    SalePrice *string `json:"sale_price,omitempty"`
    // **Global sellers** The SKU's **local pre-tax price**. This excludes any applicable charges such as cross-border shipping costs, taxes, and other fees, and therefore does not appear on the product page.  **Note**: Tax-exclusive pricing does not apply to JP and US shops using China warehouses, therefore this value is the same as `sale_price`.
    TaxExclusivePrice *string `json:"tax_exclusive_price,omitempty"`
    // The unit price of the SKU.  You can display the unit price to facilitate easier price comparisons across different products and packaging sizes. Applicable only for the EU market.  **Note**:  - This value is available only if you have defined the elements used to calculate this price when creating the product. - Unit price = Selling price/(SKU unit count/base unit count)
    UnitPrice *string `json:"unit_price,omitempty"`
}

// NewProduct202309GetProductResponseDataSkusPrice instantiates a new Product202309GetProductResponseDataSkusPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetProductResponseDataSkusPrice() *Product202309GetProductResponseDataSkusPrice {
    this := Product202309GetProductResponseDataSkusPrice{}
    return &this
}

// NewProduct202309GetProductResponseDataSkusPriceWithDefaults instantiates a new Product202309GetProductResponseDataSkusPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetProductResponseDataSkusPriceWithDefaults() *Product202309GetProductResponseDataSkusPrice {
    this := Product202309GetProductResponseDataSkusPrice{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Product202309GetProductResponseDataSkusPrice) SetCurrency(v string) {
    o.Currency = &v
}

// GetSalePrice returns the SalePrice field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusPrice) GetSalePrice() string {
    if o == nil || utils.IsNil(o.SalePrice) {
        var ret string
        return ret
    }
    return *o.SalePrice
}

// GetSalePriceOk returns a tuple with the SalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusPrice) GetSalePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SalePrice) {
        return nil, false
    }
    return o.SalePrice, true
}

// HasSalePrice returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusPrice) HasSalePrice() bool {
    if o != nil && !utils.IsNil(o.SalePrice) {
        return true
    }

    return false
}

// SetSalePrice gets a reference to the given string and assigns it to the SalePrice field.
func (o *Product202309GetProductResponseDataSkusPrice) SetSalePrice(v string) {
    o.SalePrice = &v
}

// GetTaxExclusivePrice returns the TaxExclusivePrice field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusPrice) GetTaxExclusivePrice() string {
    if o == nil || utils.IsNil(o.TaxExclusivePrice) {
        var ret string
        return ret
    }
    return *o.TaxExclusivePrice
}

// GetTaxExclusivePriceOk returns a tuple with the TaxExclusivePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusPrice) GetTaxExclusivePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TaxExclusivePrice) {
        return nil, false
    }
    return o.TaxExclusivePrice, true
}

// HasTaxExclusivePrice returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusPrice) HasTaxExclusivePrice() bool {
    if o != nil && !utils.IsNil(o.TaxExclusivePrice) {
        return true
    }

    return false
}

// SetTaxExclusivePrice gets a reference to the given string and assigns it to the TaxExclusivePrice field.
func (o *Product202309GetProductResponseDataSkusPrice) SetTaxExclusivePrice(v string) {
    o.TaxExclusivePrice = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataSkusPrice) GetUnitPrice() string {
    if o == nil || utils.IsNil(o.UnitPrice) {
        var ret string
        return ret
    }
    return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataSkusPrice) GetUnitPriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.UnitPrice) {
        return nil, false
    }
    return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataSkusPrice) HasUnitPrice() bool {
    if o != nil && !utils.IsNil(o.UnitPrice) {
        return true
    }

    return false
}

// SetUnitPrice gets a reference to the given string and assigns it to the UnitPrice field.
func (o *Product202309GetProductResponseDataSkusPrice) SetUnitPrice(v string) {
    o.UnitPrice = &v
}

func (o Product202309GetProductResponseDataSkusPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetProductResponseDataSkusPrice) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    if !utils.IsNil(o.SalePrice) {
        toSerialize["sale_price"] = o.SalePrice
    }
    if !utils.IsNil(o.TaxExclusivePrice) {
        toSerialize["tax_exclusive_price"] = o.TaxExclusivePrice
    }
    if !utils.IsNil(o.UnitPrice) {
        toSerialize["unit_price"] = o.UnitPrice
    }
    return toSerialize, nil
}

type NullableProduct202309GetProductResponseDataSkusPrice struct {
	value *Product202309GetProductResponseDataSkusPrice
	isSet bool
}

func (v NullableProduct202309GetProductResponseDataSkusPrice) Get() *Product202309GetProductResponseDataSkusPrice {
	return v.value
}

func (v *NullableProduct202309GetProductResponseDataSkusPrice) Set(val *Product202309GetProductResponseDataSkusPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetProductResponseDataSkusPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetProductResponseDataSkusPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetProductResponseDataSkusPrice(val *Product202309GetProductResponseDataSkusPrice) *NullableProduct202309GetProductResponseDataSkusPrice {
	return &NullableProduct202309GetProductResponseDataSkusPrice{value: val, isSet: true}
}

func (v NullableProduct202309GetProductResponseDataSkusPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetProductResponseDataSkusPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


