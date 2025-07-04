/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finance_v202501

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown{}

// Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown struct for Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown
type Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown struct {
    // The cash on delivery service fees charged to buyers. Applicable only for Saudi Arabia.
    CodServiceFeeAmount *string `json:"cod_service_fee_amount,omitempty"`
    // The refund for cash on delivery service fees. Applicable only for Saudi Arabia.
    RefundCodServiceFeeAmount *string `json:"refund_cod_service_fee_amount,omitempty"`
    // The total price of all refunded items before any seller discounts. This is equivalent to the shop's gross sales refund.
    RefundSubtotalBeforeDiscountAmount *string `json:"refund_subtotal_before_discount_amount,omitempty"`
    // The total amount of discounts funded by the seller, including: - Seller promotions (Product Discount, Flash Deal, Buy More Save More, Voucher and Bundle Deal) - Seller's portion of a co-funded voucher discount in co-funding campaigns - Seller discounts during a campaign
    SellerDiscountAmount *string `json:"seller_discount_amount,omitempty"`
    // Discounts returned to the sellers due to returns or refunds.
    SellerDiscountRefundAmount *string `json:"seller_discount_refund_amount,omitempty"`
    // The total price of all order items before any seller discounts and platform discounts are deducted. This is equivalent to the shop's gross sales.
    SubtotalBeforeDiscountAmount *string `json:"subtotal_before_discount_amount,omitempty"`
}

// NewFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown instantiates a new Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown() *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown {
    this := Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown{}
    return &this
}

// NewFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdownWithDefaults instantiates a new Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdownWithDefaults() *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown {
    this := Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown{}
    return &this
}

// GetCodServiceFeeAmount returns the CodServiceFeeAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetCodServiceFeeAmount() string {
    if o == nil || utils.IsNil(o.CodServiceFeeAmount) {
        var ret string
        return ret
    }
    return *o.CodServiceFeeAmount
}

// GetCodServiceFeeAmountOk returns a tuple with the CodServiceFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetCodServiceFeeAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CodServiceFeeAmount) {
        return nil, false
    }
    return o.CodServiceFeeAmount, true
}

// HasCodServiceFeeAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) HasCodServiceFeeAmount() bool {
    if o != nil && !utils.IsNil(o.CodServiceFeeAmount) {
        return true
    }

    return false
}

// SetCodServiceFeeAmount gets a reference to the given string and assigns it to the CodServiceFeeAmount field.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) SetCodServiceFeeAmount(v string) {
    o.CodServiceFeeAmount = &v
}

// GetRefundCodServiceFeeAmount returns the RefundCodServiceFeeAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetRefundCodServiceFeeAmount() string {
    if o == nil || utils.IsNil(o.RefundCodServiceFeeAmount) {
        var ret string
        return ret
    }
    return *o.RefundCodServiceFeeAmount
}

// GetRefundCodServiceFeeAmountOk returns a tuple with the RefundCodServiceFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetRefundCodServiceFeeAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RefundCodServiceFeeAmount) {
        return nil, false
    }
    return o.RefundCodServiceFeeAmount, true
}

// HasRefundCodServiceFeeAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) HasRefundCodServiceFeeAmount() bool {
    if o != nil && !utils.IsNil(o.RefundCodServiceFeeAmount) {
        return true
    }

    return false
}

// SetRefundCodServiceFeeAmount gets a reference to the given string and assigns it to the RefundCodServiceFeeAmount field.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) SetRefundCodServiceFeeAmount(v string) {
    o.RefundCodServiceFeeAmount = &v
}

// GetRefundSubtotalBeforeDiscountAmount returns the RefundSubtotalBeforeDiscountAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetRefundSubtotalBeforeDiscountAmount() string {
    if o == nil || utils.IsNil(o.RefundSubtotalBeforeDiscountAmount) {
        var ret string
        return ret
    }
    return *o.RefundSubtotalBeforeDiscountAmount
}

// GetRefundSubtotalBeforeDiscountAmountOk returns a tuple with the RefundSubtotalBeforeDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetRefundSubtotalBeforeDiscountAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RefundSubtotalBeforeDiscountAmount) {
        return nil, false
    }
    return o.RefundSubtotalBeforeDiscountAmount, true
}

// HasRefundSubtotalBeforeDiscountAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) HasRefundSubtotalBeforeDiscountAmount() bool {
    if o != nil && !utils.IsNil(o.RefundSubtotalBeforeDiscountAmount) {
        return true
    }

    return false
}

// SetRefundSubtotalBeforeDiscountAmount gets a reference to the given string and assigns it to the RefundSubtotalBeforeDiscountAmount field.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) SetRefundSubtotalBeforeDiscountAmount(v string) {
    o.RefundSubtotalBeforeDiscountAmount = &v
}

// GetSellerDiscountAmount returns the SellerDiscountAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetSellerDiscountAmount() string {
    if o == nil || utils.IsNil(o.SellerDiscountAmount) {
        var ret string
        return ret
    }
    return *o.SellerDiscountAmount
}

// GetSellerDiscountAmountOk returns a tuple with the SellerDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetSellerDiscountAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SellerDiscountAmount) {
        return nil, false
    }
    return o.SellerDiscountAmount, true
}

// HasSellerDiscountAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) HasSellerDiscountAmount() bool {
    if o != nil && !utils.IsNil(o.SellerDiscountAmount) {
        return true
    }

    return false
}

// SetSellerDiscountAmount gets a reference to the given string and assigns it to the SellerDiscountAmount field.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) SetSellerDiscountAmount(v string) {
    o.SellerDiscountAmount = &v
}

// GetSellerDiscountRefundAmount returns the SellerDiscountRefundAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetSellerDiscountRefundAmount() string {
    if o == nil || utils.IsNil(o.SellerDiscountRefundAmount) {
        var ret string
        return ret
    }
    return *o.SellerDiscountRefundAmount
}

// GetSellerDiscountRefundAmountOk returns a tuple with the SellerDiscountRefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetSellerDiscountRefundAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SellerDiscountRefundAmount) {
        return nil, false
    }
    return o.SellerDiscountRefundAmount, true
}

// HasSellerDiscountRefundAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) HasSellerDiscountRefundAmount() bool {
    if o != nil && !utils.IsNil(o.SellerDiscountRefundAmount) {
        return true
    }

    return false
}

// SetSellerDiscountRefundAmount gets a reference to the given string and assigns it to the SellerDiscountRefundAmount field.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) SetSellerDiscountRefundAmount(v string) {
    o.SellerDiscountRefundAmount = &v
}

// GetSubtotalBeforeDiscountAmount returns the SubtotalBeforeDiscountAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetSubtotalBeforeDiscountAmount() string {
    if o == nil || utils.IsNil(o.SubtotalBeforeDiscountAmount) {
        var ret string
        return ret
    }
    return *o.SubtotalBeforeDiscountAmount
}

// GetSubtotalBeforeDiscountAmountOk returns a tuple with the SubtotalBeforeDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) GetSubtotalBeforeDiscountAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SubtotalBeforeDiscountAmount) {
        return nil, false
    }
    return o.SubtotalBeforeDiscountAmount, true
}

// HasSubtotalBeforeDiscountAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) HasSubtotalBeforeDiscountAmount() bool {
    if o != nil && !utils.IsNil(o.SubtotalBeforeDiscountAmount) {
        return true
    }

    return false
}

// SetSubtotalBeforeDiscountAmount gets a reference to the given string and assigns it to the SubtotalBeforeDiscountAmount field.
func (o *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) SetSubtotalBeforeDiscountAmount(v string) {
    o.SubtotalBeforeDiscountAmount = &v
}

func (o Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CodServiceFeeAmount) {
        toSerialize["cod_service_fee_amount"] = o.CodServiceFeeAmount
    }
    if !utils.IsNil(o.RefundCodServiceFeeAmount) {
        toSerialize["refund_cod_service_fee_amount"] = o.RefundCodServiceFeeAmount
    }
    if !utils.IsNil(o.RefundSubtotalBeforeDiscountAmount) {
        toSerialize["refund_subtotal_before_discount_amount"] = o.RefundSubtotalBeforeDiscountAmount
    }
    if !utils.IsNil(o.SellerDiscountAmount) {
        toSerialize["seller_discount_amount"] = o.SellerDiscountAmount
    }
    if !utils.IsNil(o.SellerDiscountRefundAmount) {
        toSerialize["seller_discount_refund_amount"] = o.SellerDiscountRefundAmount
    }
    if !utils.IsNil(o.SubtotalBeforeDiscountAmount) {
        toSerialize["subtotal_before_discount_amount"] = o.SubtotalBeforeDiscountAmount
    }
    return toSerialize, nil
}

type NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown struct {
	value *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown
	isSet bool
}

func (v NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) Get() *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown {
	return v.value
}

func (v *NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) Set(val *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown(val *Finance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) *NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown {
	return &NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown{value: val, isSet: true}
}

func (v NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinance202501GetTransactionsbyOrderResponseDataSkuTransactionsRevenueBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


