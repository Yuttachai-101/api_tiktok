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

            // checks if the Finance202501GetTransactionsbyStatementResponseDataTransactions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Finance202501GetTransactionsbyStatementResponseDataTransactions{}

// Finance202501GetTransactionsbyStatementResponseDataTransactions struct for Finance202501GetTransactionsbyStatementResponseDataTransactions
type Finance202501GetTransactionsbyStatementResponseDataTransactions struct {
    // The adjustment amount based on TikTok Shop policy.  Refer to `transactions.type` for the list of adjustment-related policies.
    AdjustmentAmount *string `json:"adjustment_amount,omitempty"`
    // The adjustment ID if the transaction is an adjustment. Each transaction can only be associated with an order ID or an adjustment ID.
    AdjustmentId *string `json:"adjustment_id,omitempty"`
    // The order ID associated with the adjustment, if any.
    AdjustmentOrderId *string `json:"adjustment_order_id,omitempty"`
    // The order ID associated with the reserve transaction.
    AssociatedOrderId *string `json:"associated_order_id,omitempty"`
    // The estimated date when the reserve funds will be released and paid out to the seller. Unix timestamp.  Returns an empty value if the reserve funds have already been released.
    EstimatedReleaseTime *string `json:"estimated_release_time,omitempty"`
    // The total fees and taxes charged by TikTok Shop at the time of order settlement. Shipping-related costs are excluded. This is equivalent to the sum of all contributory amounts in `fee_tax_breakdown`. 
    FeeTaxAmount *string `json:"fee_tax_amount,omitempty"`
    FeeTaxBreakdown *Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown `json:"fee_tax_breakdown,omitempty"`
    // The transaction ID.
    Id *string `json:"id,omitempty"`
    // The creation time of the order. Unix timestamp.
    OrderCreateTime *int64 `json:"order_create_time,omitempty"`
    // The order ID. Each transaction can only be associated with an order ID or an adjustment ID.
    OrderId *string `json:"order_id,omitempty"`
    // The amount withheld from settlement based on TikTok Shop Reserve Policy. Refer to TikTok Shop Academy for more information. - A positive amount indicates funds that have been released. - A negative amount indicates funds being withheld from the settlement.
    ReserveAmount *string `json:"reserve_amount,omitempty"`
    // The ID of a reserve transaction.
    ReserveId *string `json:"reserve_id,omitempty"`
    // The status of the reserve funds. Possible values: - COLLECTED: A portion of the order's settlement amount has been withheld as reserve funds. - RELEASED: The previously reserved funds have been released and paid out to the seller.
    ReserveStatus *string `json:"reserve_status,omitempty"`
    // The revenue amount at the time of order settlement. This is equivalent to the sum of all amounts in `revenue_breakdown`.
    RevenueAmount *string `json:"revenue_amount,omitempty"`
    RevenueBreakdown *Finance202501GetTransactionsbyStatementResponseDataTransactionsRevenueBreakdown `json:"revenue_breakdown,omitempty"`
    // The settlement amount for the order.  Formula: revenue_amount - shipping_cost_amount - fee_tax_amount - adjustment_amount
    SettlementAmount *string `json:"settlement_amount,omitempty"`
    // The shipping costs at the time of order settlement. This is equivalent to the sum of all contributory amounts in `shipping_cost_breakdown`.
    ShippingCostAmount *string `json:"shipping_cost_amount,omitempty"`
    ShippingCostBreakdown *Finance202501GetTransactionsbyStatementResponseDataTransactionsShippingCostBreakdown `json:"shipping_cost_breakdown,omitempty"`
    SupplementaryComponent *Finance202501GetTransactionsbyStatementResponseDataTransactionsSupplementaryComponent `json:"supplementary_component,omitempty"`
    // The transaction type.  **Standard transactions** - `ORDER`: A transaction related to an order settlement. - `RESERVE`: A transaction involving collection or release of reserve funds. - If the transaction is an adjustment, it returns one of the following values:  **Platform-related adjustments** - `CHARGE_BACK`: Charges returned to a payment card after a customer has successfully disputed an item on their account statement or transactions report. - `CUSTOMER_SERVICE_COMPENSATION`: Extra compensation or compensation paid to a customer after the after-sales period by customer service. - `DEDUCTIONS_INCURRED_BY_SELLER`: Deduction arising from customer dissatisfaction as a result of the seller's responsibility. This includes issues such as fraud, empty packages, items that do not match the product display page, or items of lower value than advertised. - `GMV_PAYMENT_FOR_ADS`: Amount used to pay for your advertisement if you are enabled \"auto pay ads with shop GMV\", or to pay for Tiktok Promote ads orders. - `PLATFORM_COMMISSION_ADJUSTMENT`: Adjustment when there are differences in the platform commission paid by the seller. - `PLATFORM_COMMISSION_COMPENSATION`: Compensation paid to the seller when there are differences in the platform commission paid by the seller. - `PLATFORM_PENALTY`: Penalty imposed for a violation of TikTok Shop policies (the corresponding amount has been deducted from the seller's account). For details, please refer to the email notification sent to the seller. - `PROMOTION_ADJUSTMENT`: Adjustment when a seller takes part in a platform promotion and there are differences between the promotion price and the actual amount paid by the seller. - `REBATE`: A discount on referral fees offered by TikTok Shop to eligible sellers. - `PLATFORM_COMPENSATION`: Compensation paid to the seller after the seller successfully appealed for a customer dispute. - `PLATFORM_REIMBURSEMENT`: Reimbursement paid by TikTok Shop for an order refunded under TikTok's refund without return policy (the seller is not responsible). - `COFUNDED_CREATOR_REWARDS`: Fees charged for joining the co-funded creator rewards program to reward creator activities.  **Logistics-related adjustments** - `FBT_WAREHOUSE_SERVICE_FEE`: Amount charged by TikTok Fulfillment Portal (Pipak) for warehousing-related bills incurred by the seller under the Fulfilled by TikTok (FBT) service. - `LOGISTICS_REIMBURSEMENT`: Reimbursement paid by TikTok Shop for an order refunded due to logistics-related issues (e.g. lost or damaged order). - `SHIPPING_FEE_ADJUSTMENT`: Adjustment when there are differences or mistakes with the shipping fee paid by the seller. - `SHIPPING_FEE_COMPENSATION`: Compensation given to sellers due to differences between the actual shipping fee and the pre-paid shipping fee. - `SHIPPING_FEE_REBATE`: Shipping fee rebate provided to the seller as part of their participation in a platform campaign. - `SAMPLE_SHIPPING_FEE`: Fees charged for sending samples using the TikTok logistics provider. - `SELLER_MISSION_REWARD`: Platform provided cash reward for seller mission completion.  **Miscellaneous adjustments** `OTHER_ADJUSTMENT`: Adjustment for other reasons.
    Type *string `json:"type,omitempty"`
}

// NewFinance202501GetTransactionsbyStatementResponseDataTransactions instantiates a new Finance202501GetTransactionsbyStatementResponseDataTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinance202501GetTransactionsbyStatementResponseDataTransactions() *Finance202501GetTransactionsbyStatementResponseDataTransactions {
    this := Finance202501GetTransactionsbyStatementResponseDataTransactions{}
    return &this
}

// NewFinance202501GetTransactionsbyStatementResponseDataTransactionsWithDefaults instantiates a new Finance202501GetTransactionsbyStatementResponseDataTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinance202501GetTransactionsbyStatementResponseDataTransactionsWithDefaults() *Finance202501GetTransactionsbyStatementResponseDataTransactions {
    this := Finance202501GetTransactionsbyStatementResponseDataTransactions{}
    return &this
}

// GetAdjustmentAmount returns the AdjustmentAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAdjustmentAmount() string {
    if o == nil || utils.IsNil(o.AdjustmentAmount) {
        var ret string
        return ret
    }
    return *o.AdjustmentAmount
}

// GetAdjustmentAmountOk returns a tuple with the AdjustmentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAdjustmentAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.AdjustmentAmount) {
        return nil, false
    }
    return o.AdjustmentAmount, true
}

// HasAdjustmentAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasAdjustmentAmount() bool {
    if o != nil && !utils.IsNil(o.AdjustmentAmount) {
        return true
    }

    return false
}

// SetAdjustmentAmount gets a reference to the given string and assigns it to the AdjustmentAmount field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetAdjustmentAmount(v string) {
    o.AdjustmentAmount = &v
}

// GetAdjustmentId returns the AdjustmentId field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAdjustmentId() string {
    if o == nil || utils.IsNil(o.AdjustmentId) {
        var ret string
        return ret
    }
    return *o.AdjustmentId
}

// GetAdjustmentIdOk returns a tuple with the AdjustmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAdjustmentIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.AdjustmentId) {
        return nil, false
    }
    return o.AdjustmentId, true
}

// HasAdjustmentId returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasAdjustmentId() bool {
    if o != nil && !utils.IsNil(o.AdjustmentId) {
        return true
    }

    return false
}

// SetAdjustmentId gets a reference to the given string and assigns it to the AdjustmentId field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetAdjustmentId(v string) {
    o.AdjustmentId = &v
}

// GetAdjustmentOrderId returns the AdjustmentOrderId field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAdjustmentOrderId() string {
    if o == nil || utils.IsNil(o.AdjustmentOrderId) {
        var ret string
        return ret
    }
    return *o.AdjustmentOrderId
}

// GetAdjustmentOrderIdOk returns a tuple with the AdjustmentOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAdjustmentOrderIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.AdjustmentOrderId) {
        return nil, false
    }
    return o.AdjustmentOrderId, true
}

// HasAdjustmentOrderId returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasAdjustmentOrderId() bool {
    if o != nil && !utils.IsNil(o.AdjustmentOrderId) {
        return true
    }

    return false
}

// SetAdjustmentOrderId gets a reference to the given string and assigns it to the AdjustmentOrderId field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetAdjustmentOrderId(v string) {
    o.AdjustmentOrderId = &v
}

// GetAssociatedOrderId returns the AssociatedOrderId field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAssociatedOrderId() string {
    if o == nil || utils.IsNil(o.AssociatedOrderId) {
        var ret string
        return ret
    }
    return *o.AssociatedOrderId
}

// GetAssociatedOrderIdOk returns a tuple with the AssociatedOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetAssociatedOrderIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.AssociatedOrderId) {
        return nil, false
    }
    return o.AssociatedOrderId, true
}

// HasAssociatedOrderId returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasAssociatedOrderId() bool {
    if o != nil && !utils.IsNil(o.AssociatedOrderId) {
        return true
    }

    return false
}

// SetAssociatedOrderId gets a reference to the given string and assigns it to the AssociatedOrderId field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetAssociatedOrderId(v string) {
    o.AssociatedOrderId = &v
}

// GetEstimatedReleaseTime returns the EstimatedReleaseTime field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetEstimatedReleaseTime() string {
    if o == nil || utils.IsNil(o.EstimatedReleaseTime) {
        var ret string
        return ret
    }
    return *o.EstimatedReleaseTime
}

// GetEstimatedReleaseTimeOk returns a tuple with the EstimatedReleaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetEstimatedReleaseTimeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.EstimatedReleaseTime) {
        return nil, false
    }
    return o.EstimatedReleaseTime, true
}

// HasEstimatedReleaseTime returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasEstimatedReleaseTime() bool {
    if o != nil && !utils.IsNil(o.EstimatedReleaseTime) {
        return true
    }

    return false
}

// SetEstimatedReleaseTime gets a reference to the given string and assigns it to the EstimatedReleaseTime field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetEstimatedReleaseTime(v string) {
    o.EstimatedReleaseTime = &v
}

// GetFeeTaxAmount returns the FeeTaxAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetFeeTaxAmount() string {
    if o == nil || utils.IsNil(o.FeeTaxAmount) {
        var ret string
        return ret
    }
    return *o.FeeTaxAmount
}

// GetFeeTaxAmountOk returns a tuple with the FeeTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetFeeTaxAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.FeeTaxAmount) {
        return nil, false
    }
    return o.FeeTaxAmount, true
}

// HasFeeTaxAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasFeeTaxAmount() bool {
    if o != nil && !utils.IsNil(o.FeeTaxAmount) {
        return true
    }

    return false
}

// SetFeeTaxAmount gets a reference to the given string and assigns it to the FeeTaxAmount field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetFeeTaxAmount(v string) {
    o.FeeTaxAmount = &v
}

// GetFeeTaxBreakdown returns the FeeTaxBreakdown field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetFeeTaxBreakdown() Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown {
    if o == nil || utils.IsNil(o.FeeTaxBreakdown) {
        var ret Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown
        return ret
    }
    return *o.FeeTaxBreakdown
}

// GetFeeTaxBreakdownOk returns a tuple with the FeeTaxBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetFeeTaxBreakdownOk() (*Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown, bool) {
    if o == nil || utils.IsNil(o.FeeTaxBreakdown) {
        return nil, false
    }
    return o.FeeTaxBreakdown, true
}

// HasFeeTaxBreakdown returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasFeeTaxBreakdown() bool {
    if o != nil && !utils.IsNil(o.FeeTaxBreakdown) {
        return true
    }

    return false
}

// SetFeeTaxBreakdown gets a reference to the given Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown and assigns it to the FeeTaxBreakdown field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetFeeTaxBreakdown(v Finance202501GetTransactionsbyStatementResponseDataTransactionsFeeTaxBreakdown) {
    o.FeeTaxBreakdown = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetId(v string) {
    o.Id = &v
}

// GetOrderCreateTime returns the OrderCreateTime field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetOrderCreateTime() int64 {
    if o == nil || utils.IsNil(o.OrderCreateTime) {
        var ret int64
        return ret
    }
    return *o.OrderCreateTime
}

// GetOrderCreateTimeOk returns a tuple with the OrderCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetOrderCreateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.OrderCreateTime) {
        return nil, false
    }
    return o.OrderCreateTime, true
}

// HasOrderCreateTime returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasOrderCreateTime() bool {
    if o != nil && !utils.IsNil(o.OrderCreateTime) {
        return true
    }

    return false
}

// SetOrderCreateTime gets a reference to the given int64 and assigns it to the OrderCreateTime field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetOrderCreateTime(v int64) {
    o.OrderCreateTime = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetOrderId() string {
    if o == nil || utils.IsNil(o.OrderId) {
        var ret string
        return ret
    }
    return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetOrderIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.OrderId) {
        return nil, false
    }
    return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasOrderId() bool {
    if o != nil && !utils.IsNil(o.OrderId) {
        return true
    }

    return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetOrderId(v string) {
    o.OrderId = &v
}

// GetReserveAmount returns the ReserveAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetReserveAmount() string {
    if o == nil || utils.IsNil(o.ReserveAmount) {
        var ret string
        return ret
    }
    return *o.ReserveAmount
}

// GetReserveAmountOk returns a tuple with the ReserveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetReserveAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReserveAmount) {
        return nil, false
    }
    return o.ReserveAmount, true
}

// HasReserveAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasReserveAmount() bool {
    if o != nil && !utils.IsNil(o.ReserveAmount) {
        return true
    }

    return false
}

// SetReserveAmount gets a reference to the given string and assigns it to the ReserveAmount field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetReserveAmount(v string) {
    o.ReserveAmount = &v
}

// GetReserveId returns the ReserveId field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetReserveId() string {
    if o == nil || utils.IsNil(o.ReserveId) {
        var ret string
        return ret
    }
    return *o.ReserveId
}

// GetReserveIdOk returns a tuple with the ReserveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetReserveIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReserveId) {
        return nil, false
    }
    return o.ReserveId, true
}

// HasReserveId returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasReserveId() bool {
    if o != nil && !utils.IsNil(o.ReserveId) {
        return true
    }

    return false
}

// SetReserveId gets a reference to the given string and assigns it to the ReserveId field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetReserveId(v string) {
    o.ReserveId = &v
}

// GetReserveStatus returns the ReserveStatus field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetReserveStatus() string {
    if o == nil || utils.IsNil(o.ReserveStatus) {
        var ret string
        return ret
    }
    return *o.ReserveStatus
}

// GetReserveStatusOk returns a tuple with the ReserveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetReserveStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReserveStatus) {
        return nil, false
    }
    return o.ReserveStatus, true
}

// HasReserveStatus returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasReserveStatus() bool {
    if o != nil && !utils.IsNil(o.ReserveStatus) {
        return true
    }

    return false
}

// SetReserveStatus gets a reference to the given string and assigns it to the ReserveStatus field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetReserveStatus(v string) {
    o.ReserveStatus = &v
}

// GetRevenueAmount returns the RevenueAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetRevenueAmount() string {
    if o == nil || utils.IsNil(o.RevenueAmount) {
        var ret string
        return ret
    }
    return *o.RevenueAmount
}

// GetRevenueAmountOk returns a tuple with the RevenueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetRevenueAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RevenueAmount) {
        return nil, false
    }
    return o.RevenueAmount, true
}

// HasRevenueAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasRevenueAmount() bool {
    if o != nil && !utils.IsNil(o.RevenueAmount) {
        return true
    }

    return false
}

// SetRevenueAmount gets a reference to the given string and assigns it to the RevenueAmount field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetRevenueAmount(v string) {
    o.RevenueAmount = &v
}

// GetRevenueBreakdown returns the RevenueBreakdown field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetRevenueBreakdown() Finance202501GetTransactionsbyStatementResponseDataTransactionsRevenueBreakdown {
    if o == nil || utils.IsNil(o.RevenueBreakdown) {
        var ret Finance202501GetTransactionsbyStatementResponseDataTransactionsRevenueBreakdown
        return ret
    }
    return *o.RevenueBreakdown
}

// GetRevenueBreakdownOk returns a tuple with the RevenueBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetRevenueBreakdownOk() (*Finance202501GetTransactionsbyStatementResponseDataTransactionsRevenueBreakdown, bool) {
    if o == nil || utils.IsNil(o.RevenueBreakdown) {
        return nil, false
    }
    return o.RevenueBreakdown, true
}

// HasRevenueBreakdown returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasRevenueBreakdown() bool {
    if o != nil && !utils.IsNil(o.RevenueBreakdown) {
        return true
    }

    return false
}

// SetRevenueBreakdown gets a reference to the given Finance202501GetTransactionsbyStatementResponseDataTransactionsRevenueBreakdown and assigns it to the RevenueBreakdown field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetRevenueBreakdown(v Finance202501GetTransactionsbyStatementResponseDataTransactionsRevenueBreakdown) {
    o.RevenueBreakdown = &v
}

// GetSettlementAmount returns the SettlementAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetSettlementAmount() string {
    if o == nil || utils.IsNil(o.SettlementAmount) {
        var ret string
        return ret
    }
    return *o.SettlementAmount
}

// GetSettlementAmountOk returns a tuple with the SettlementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetSettlementAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SettlementAmount) {
        return nil, false
    }
    return o.SettlementAmount, true
}

// HasSettlementAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasSettlementAmount() bool {
    if o != nil && !utils.IsNil(o.SettlementAmount) {
        return true
    }

    return false
}

// SetSettlementAmount gets a reference to the given string and assigns it to the SettlementAmount field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetSettlementAmount(v string) {
    o.SettlementAmount = &v
}

// GetShippingCostAmount returns the ShippingCostAmount field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetShippingCostAmount() string {
    if o == nil || utils.IsNil(o.ShippingCostAmount) {
        var ret string
        return ret
    }
    return *o.ShippingCostAmount
}

// GetShippingCostAmountOk returns a tuple with the ShippingCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetShippingCostAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ShippingCostAmount) {
        return nil, false
    }
    return o.ShippingCostAmount, true
}

// HasShippingCostAmount returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasShippingCostAmount() bool {
    if o != nil && !utils.IsNil(o.ShippingCostAmount) {
        return true
    }

    return false
}

// SetShippingCostAmount gets a reference to the given string and assigns it to the ShippingCostAmount field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetShippingCostAmount(v string) {
    o.ShippingCostAmount = &v
}

// GetShippingCostBreakdown returns the ShippingCostBreakdown field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetShippingCostBreakdown() Finance202501GetTransactionsbyStatementResponseDataTransactionsShippingCostBreakdown {
    if o == nil || utils.IsNil(o.ShippingCostBreakdown) {
        var ret Finance202501GetTransactionsbyStatementResponseDataTransactionsShippingCostBreakdown
        return ret
    }
    return *o.ShippingCostBreakdown
}

// GetShippingCostBreakdownOk returns a tuple with the ShippingCostBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetShippingCostBreakdownOk() (*Finance202501GetTransactionsbyStatementResponseDataTransactionsShippingCostBreakdown, bool) {
    if o == nil || utils.IsNil(o.ShippingCostBreakdown) {
        return nil, false
    }
    return o.ShippingCostBreakdown, true
}

// HasShippingCostBreakdown returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasShippingCostBreakdown() bool {
    if o != nil && !utils.IsNil(o.ShippingCostBreakdown) {
        return true
    }

    return false
}

// SetShippingCostBreakdown gets a reference to the given Finance202501GetTransactionsbyStatementResponseDataTransactionsShippingCostBreakdown and assigns it to the ShippingCostBreakdown field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetShippingCostBreakdown(v Finance202501GetTransactionsbyStatementResponseDataTransactionsShippingCostBreakdown) {
    o.ShippingCostBreakdown = &v
}

// GetSupplementaryComponent returns the SupplementaryComponent field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetSupplementaryComponent() Finance202501GetTransactionsbyStatementResponseDataTransactionsSupplementaryComponent {
    if o == nil || utils.IsNil(o.SupplementaryComponent) {
        var ret Finance202501GetTransactionsbyStatementResponseDataTransactionsSupplementaryComponent
        return ret
    }
    return *o.SupplementaryComponent
}

// GetSupplementaryComponentOk returns a tuple with the SupplementaryComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetSupplementaryComponentOk() (*Finance202501GetTransactionsbyStatementResponseDataTransactionsSupplementaryComponent, bool) {
    if o == nil || utils.IsNil(o.SupplementaryComponent) {
        return nil, false
    }
    return o.SupplementaryComponent, true
}

// HasSupplementaryComponent returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasSupplementaryComponent() bool {
    if o != nil && !utils.IsNil(o.SupplementaryComponent) {
        return true
    }

    return false
}

// SetSupplementaryComponent gets a reference to the given Finance202501GetTransactionsbyStatementResponseDataTransactionsSupplementaryComponent and assigns it to the SupplementaryComponent field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetSupplementaryComponent(v Finance202501GetTransactionsbyStatementResponseDataTransactionsSupplementaryComponent) {
    o.SupplementaryComponent = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetType() string {
    if o == nil || utils.IsNil(o.Type) {
        var ret string
        return ret
    }
    return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) GetTypeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Type) {
        return nil, false
    }
    return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) HasType() bool {
    if o != nil && !utils.IsNil(o.Type) {
        return true
    }

    return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Finance202501GetTransactionsbyStatementResponseDataTransactions) SetType(v string) {
    o.Type = &v
}

func (o Finance202501GetTransactionsbyStatementResponseDataTransactions) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Finance202501GetTransactionsbyStatementResponseDataTransactions) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AdjustmentAmount) {
        toSerialize["adjustment_amount"] = o.AdjustmentAmount
    }
    if !utils.IsNil(o.AdjustmentId) {
        toSerialize["adjustment_id"] = o.AdjustmentId
    }
    if !utils.IsNil(o.AdjustmentOrderId) {
        toSerialize["adjustment_order_id"] = o.AdjustmentOrderId
    }
    if !utils.IsNil(o.AssociatedOrderId) {
        toSerialize["associated_order_id"] = o.AssociatedOrderId
    }
    if !utils.IsNil(o.EstimatedReleaseTime) {
        toSerialize["estimated_release_time"] = o.EstimatedReleaseTime
    }
    if !utils.IsNil(o.FeeTaxAmount) {
        toSerialize["fee_tax_amount"] = o.FeeTaxAmount
    }
    if !utils.IsNil(o.FeeTaxBreakdown) {
        toSerialize["fee_tax_breakdown"] = o.FeeTaxBreakdown
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.OrderCreateTime) {
        toSerialize["order_create_time"] = o.OrderCreateTime
    }
    if !utils.IsNil(o.OrderId) {
        toSerialize["order_id"] = o.OrderId
    }
    if !utils.IsNil(o.ReserveAmount) {
        toSerialize["reserve_amount"] = o.ReserveAmount
    }
    if !utils.IsNil(o.ReserveId) {
        toSerialize["reserve_id"] = o.ReserveId
    }
    if !utils.IsNil(o.ReserveStatus) {
        toSerialize["reserve_status"] = o.ReserveStatus
    }
    if !utils.IsNil(o.RevenueAmount) {
        toSerialize["revenue_amount"] = o.RevenueAmount
    }
    if !utils.IsNil(o.RevenueBreakdown) {
        toSerialize["revenue_breakdown"] = o.RevenueBreakdown
    }
    if !utils.IsNil(o.SettlementAmount) {
        toSerialize["settlement_amount"] = o.SettlementAmount
    }
    if !utils.IsNil(o.ShippingCostAmount) {
        toSerialize["shipping_cost_amount"] = o.ShippingCostAmount
    }
    if !utils.IsNil(o.ShippingCostBreakdown) {
        toSerialize["shipping_cost_breakdown"] = o.ShippingCostBreakdown
    }
    if !utils.IsNil(o.SupplementaryComponent) {
        toSerialize["supplementary_component"] = o.SupplementaryComponent
    }
    if !utils.IsNil(o.Type) {
        toSerialize["type"] = o.Type
    }
    return toSerialize, nil
}

type NullableFinance202501GetTransactionsbyStatementResponseDataTransactions struct {
	value *Finance202501GetTransactionsbyStatementResponseDataTransactions
	isSet bool
}

func (v NullableFinance202501GetTransactionsbyStatementResponseDataTransactions) Get() *Finance202501GetTransactionsbyStatementResponseDataTransactions {
	return v.value
}

func (v *NullableFinance202501GetTransactionsbyStatementResponseDataTransactions) Set(val *Finance202501GetTransactionsbyStatementResponseDataTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullableFinance202501GetTransactionsbyStatementResponseDataTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableFinance202501GetTransactionsbyStatementResponseDataTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinance202501GetTransactionsbyStatementResponseDataTransactions(val *Finance202501GetTransactionsbyStatementResponseDataTransactions) *NullableFinance202501GetTransactionsbyStatementResponseDataTransactions {
	return &NullableFinance202501GetTransactionsbyStatementResponseDataTransactions{value: val, isSet: true}
}

func (v NullableFinance202501GetTransactionsbyStatementResponseDataTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinance202501GetTransactionsbyStatementResponseDataTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


