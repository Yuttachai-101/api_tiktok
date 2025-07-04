/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finance_v202309

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the Finance202309GetWithdrawalsResponseDataWithdrawals type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Finance202309GetWithdrawalsResponseDataWithdrawals{}

// Finance202309GetWithdrawalsResponseDataWithdrawals struct for Finance202309GetWithdrawalsResponseDataWithdrawals
type Finance202309GetWithdrawalsResponseDataWithdrawals struct {
    // Withdraw amount
    Amount *string `json:"amount,omitempty"`
    // Withdraw create time
    CreateTime *int64 `json:"create_time,omitempty"`
    // The three-digit currency code in ISO 4217 format.   
    Currency *string `json:"currency,omitempty"`
    // A unique transaction id generated by TTS for the withdrawal.
    Id *string `json:"id,omitempty"`
    // The processing status of the withdrawal indicates whether the withdrawal is transferred. Possible values: - PROCESSING：the withdrawal is currently processing. If the withdrawal is successful, the status will transition to PAID. If not, it will be FAILED. - SUCCESS：the withdrawal has been transferred to the Seller - FAILED：the withdrawal failed
    Status *string `json:"status,omitempty"`
    // WITHDRAW：The action of the seller to receive the settlement amount to the bank card through the action of withdrawal SETTLE：The platform settles the amount to the seller TRANSFER：Platform subsidies or deductions due to platform policies REVERSE：Withdrawal failure due to incorrect bank card 
    Type *string `json:"type,omitempty"`
}

// NewFinance202309GetWithdrawalsResponseDataWithdrawals instantiates a new Finance202309GetWithdrawalsResponseDataWithdrawals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinance202309GetWithdrawalsResponseDataWithdrawals() *Finance202309GetWithdrawalsResponseDataWithdrawals {
    this := Finance202309GetWithdrawalsResponseDataWithdrawals{}
    return &this
}

// NewFinance202309GetWithdrawalsResponseDataWithdrawalsWithDefaults instantiates a new Finance202309GetWithdrawalsResponseDataWithdrawals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinance202309GetWithdrawalsResponseDataWithdrawalsWithDefaults() *Finance202309GetWithdrawalsResponseDataWithdrawals {
    this := Finance202309GetWithdrawalsResponseDataWithdrawals{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) SetAmount(v string) {
    o.Amount = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetCreateTime() int64 {
    if o == nil || utils.IsNil(o.CreateTime) {
        var ret int64
        return ret
    }
    return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetCreateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTime) {
        return nil, false
    }
    return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) HasCreateTime() bool {
    if o != nil && !utils.IsNil(o.CreateTime) {
        return true
    }

    return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) SetCreateTime(v int64) {
    o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) SetCurrency(v string) {
    o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) SetId(v string) {
    o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) SetStatus(v string) {
    o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetType() string {
    if o == nil || utils.IsNil(o.Type) {
        var ret string
        return ret
    }
    return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) GetTypeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Type) {
        return nil, false
    }
    return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) HasType() bool {
    if o != nil && !utils.IsNil(o.Type) {
        return true
    }

    return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Finance202309GetWithdrawalsResponseDataWithdrawals) SetType(v string) {
    o.Type = &v
}

func (o Finance202309GetWithdrawalsResponseDataWithdrawals) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Finance202309GetWithdrawalsResponseDataWithdrawals) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.CreateTime) {
        toSerialize["create_time"] = o.CreateTime
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.Type) {
        toSerialize["type"] = o.Type
    }
    return toSerialize, nil
}

type NullableFinance202309GetWithdrawalsResponseDataWithdrawals struct {
	value *Finance202309GetWithdrawalsResponseDataWithdrawals
	isSet bool
}

func (v NullableFinance202309GetWithdrawalsResponseDataWithdrawals) Get() *Finance202309GetWithdrawalsResponseDataWithdrawals {
	return v.value
}

func (v *NullableFinance202309GetWithdrawalsResponseDataWithdrawals) Set(val *Finance202309GetWithdrawalsResponseDataWithdrawals) {
	v.value = val
	v.isSet = true
}

func (v NullableFinance202309GetWithdrawalsResponseDataWithdrawals) IsSet() bool {
	return v.isSet
}

func (v *NullableFinance202309GetWithdrawalsResponseDataWithdrawals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinance202309GetWithdrawalsResponseDataWithdrawals(val *Finance202309GetWithdrawalsResponseDataWithdrawals) *NullableFinance202309GetWithdrawalsResponseDataWithdrawals {
	return &NullableFinance202309GetWithdrawalsResponseDataWithdrawals{value: val, isSet: true}
}

func (v NullableFinance202309GetWithdrawalsResponseDataWithdrawals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinance202309GetWithdrawalsResponseDataWithdrawals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


