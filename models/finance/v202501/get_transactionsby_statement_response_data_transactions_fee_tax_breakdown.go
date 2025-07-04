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

            // checks if the Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown{}

// Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown struct for Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown
type Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown struct {
    Fee *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownFee `json:"fee,omitempty"`
    Tax *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownTax `json:"tax,omitempty"`
}

// NewFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown instantiates a new Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown() *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown {
    this := Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown{}
    return &this
}

// NewFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownWithDefaults instantiates a new Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownWithDefaults() *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown {
    this := Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown{}
    return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) GetFee() Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownFee {
    if o == nil || utils.IsNil(o.Fee) {
        var ret Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownFee
        return ret
    }
    return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) GetFeeOk() (*Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownFee, bool) {
    if o == nil || utils.IsNil(o.Fee) {
        return nil, false
    }
    return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) HasFee() bool {
    if o != nil && !utils.IsNil(o.Fee) {
        return true
    }

    return false
}

// SetFee gets a reference to the given Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownFee and assigns it to the Fee field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) SetFee(v Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownFee) {
    o.Fee = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) GetTax() Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownTax {
    if o == nil || utils.IsNil(o.Tax) {
        var ret Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownTax
        return ret
    }
    return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) GetTaxOk() (*Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownTax, bool) {
    if o == nil || utils.IsNil(o.Tax) {
        return nil, false
    }
    return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) HasTax() bool {
    if o != nil && !utils.IsNil(o.Tax) {
        return true
    }

    return false
}

// SetTax gets a reference to the given Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownTax and assigns it to the Tax field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) SetTax(v Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdownTax) {
    o.Tax = &v
}

func (o Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Fee) {
        toSerialize["fee"] = o.Fee
    }
    if !utils.IsNil(o.Tax) {
        toSerialize["tax"] = o.Tax
    }
    return toSerialize, nil
}

type NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown struct {
	value *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown
	isSet bool
}

func (v NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) Get() *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown {
	return v.value
}

func (v *NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) Set(val *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown(val *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) *NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown {
	return &NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown{value: val, isSet: true}
}

func (v NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


