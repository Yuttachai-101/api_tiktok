/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affiliate_partner_v202405

import (
    "encoding/json"
    "github.com/Yuttachai-101/api_tiktok/utils"
)

            // checks if the AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts{}

// AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts struct for AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts
type AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts struct {
    Category *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsCategory `json:"category,omitempty"`
    // The creator commission rate in hundredths of a percent (0.01%) units.
    CreatorCommissionRate *int64 `json:"creator_commission_rate,omitempty"`
    HighestPrice *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsHighestPrice `json:"highest_price,omitempty"`
    // The product identifier.
    Id *string `json:"id,omitempty"`
    // The number of in-stock units of all SKUs for this product.
    Inventory *int64 `json:"inventory,omitempty"`
    // Set to `true` if a product URL is available. Set to `false` if no product URL is available.
    IsAvailable *bool `json:"is_available,omitempty"`
    LowestPrice *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsLowestPrice `json:"lowest_price,omitempty"`
    // The product image URL.
    MainImageUrl *string `json:"main_image_url,omitempty"`
    // The product name.
    Name *string `json:"name,omitempty"`
    // The product open collaboration commission rate in hundredths of a percent (0.01%) units.  
    OpenCollaborationCommissionRate *int64 `json:"open_collaboration_commission_rate,omitempty"`
    // The partner commission rate in hundredths of a percent (0.01%) units.
    PartnerCommissionRate *int64 `json:"partner_commission_rate,omitempty"`
    // Return product description
    ProductDescription *string `json:"product_description,omitempty"`
    // The total number of units sold of this product for this campaign.
    ProductSales *int64 `json:"product_sales,omitempty"`
    // The product review status. This an enumerated type with values: - PENDING - APPROVED - REJECTED - PENDING_CLOSED - CLOSED
    ReviewStatus *string `json:"review_status,omitempty"`
    // The total amount of sample inventory available for allocation to creators by the seller.
    SampleQuota *int64 `json:"sample_quota,omitempty"`
    // The TikTok Shop name.
    ShopName *string `json:"shop_name,omitempty"`
    // A list of Stock Keeping Units (SKUs) used to identify distinct variants of the product.
    SkuInformationList []AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsSkuInformationList `json:"sku_information_list,omitempty"`
    // The total commission rate in hundredths of a percent (0.01%) units.
    TotalCommissionRate *int64 `json:"total_commission_rate,omitempty"`
}

// NewAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts instantiates a new AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts() *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts {
    this := AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts{}
    return &this
}

// NewAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsWithDefaults instantiates a new AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsWithDefaults() *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts {
    this := AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts{}
    return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetCategory() AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsCategory {
    if o == nil || utils.IsNil(o.Category) {
        var ret AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsCategory
        return ret
    }
    return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetCategoryOk() (*AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsCategory, bool) {
    if o == nil || utils.IsNil(o.Category) {
        return nil, false
    }
    return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasCategory() bool {
    if o != nil && !utils.IsNil(o.Category) {
        return true
    }

    return false
}

// SetCategory gets a reference to the given AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsCategory and assigns it to the Category field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetCategory(v AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsCategory) {
    o.Category = &v
}

// GetCreatorCommissionRate returns the CreatorCommissionRate field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetCreatorCommissionRate() int64 {
    if o == nil || utils.IsNil(o.CreatorCommissionRate) {
        var ret int64
        return ret
    }
    return *o.CreatorCommissionRate
}

// GetCreatorCommissionRateOk returns a tuple with the CreatorCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetCreatorCommissionRateOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreatorCommissionRate) {
        return nil, false
    }
    return o.CreatorCommissionRate, true
}

// HasCreatorCommissionRate returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasCreatorCommissionRate() bool {
    if o != nil && !utils.IsNil(o.CreatorCommissionRate) {
        return true
    }

    return false
}

// SetCreatorCommissionRate gets a reference to the given int64 and assigns it to the CreatorCommissionRate field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetCreatorCommissionRate(v int64) {
    o.CreatorCommissionRate = &v
}

// GetHighestPrice returns the HighestPrice field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetHighestPrice() AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsHighestPrice {
    if o == nil || utils.IsNil(o.HighestPrice) {
        var ret AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsHighestPrice
        return ret
    }
    return *o.HighestPrice
}

// GetHighestPriceOk returns a tuple with the HighestPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetHighestPriceOk() (*AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsHighestPrice, bool) {
    if o == nil || utils.IsNil(o.HighestPrice) {
        return nil, false
    }
    return o.HighestPrice, true
}

// HasHighestPrice returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasHighestPrice() bool {
    if o != nil && !utils.IsNil(o.HighestPrice) {
        return true
    }

    return false
}

// SetHighestPrice gets a reference to the given AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsHighestPrice and assigns it to the HighestPrice field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetHighestPrice(v AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsHighestPrice) {
    o.HighestPrice = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetId(v string) {
    o.Id = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetInventory() int64 {
    if o == nil || utils.IsNil(o.Inventory) {
        var ret int64
        return ret
    }
    return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetInventoryOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.Inventory) {
        return nil, false
    }
    return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasInventory() bool {
    if o != nil && !utils.IsNil(o.Inventory) {
        return true
    }

    return false
}

// SetInventory gets a reference to the given int64 and assigns it to the Inventory field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetInventory(v int64) {
    o.Inventory = &v
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetIsAvailable() bool {
    if o == nil || utils.IsNil(o.IsAvailable) {
        var ret bool
        return ret
    }
    return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetIsAvailableOk() (*bool, bool) {
    if o == nil || utils.IsNil(o.IsAvailable) {
        return nil, false
    }
    return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasIsAvailable() bool {
    if o != nil && !utils.IsNil(o.IsAvailable) {
        return true
    }

    return false
}

// SetIsAvailable gets a reference to the given bool and assigns it to the IsAvailable field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetIsAvailable(v bool) {
    o.IsAvailable = &v
}

// GetLowestPrice returns the LowestPrice field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetLowestPrice() AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsLowestPrice {
    if o == nil || utils.IsNil(o.LowestPrice) {
        var ret AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsLowestPrice
        return ret
    }
    return *o.LowestPrice
}

// GetLowestPriceOk returns a tuple with the LowestPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetLowestPriceOk() (*AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsLowestPrice, bool) {
    if o == nil || utils.IsNil(o.LowestPrice) {
        return nil, false
    }
    return o.LowestPrice, true
}

// HasLowestPrice returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasLowestPrice() bool {
    if o != nil && !utils.IsNil(o.LowestPrice) {
        return true
    }

    return false
}

// SetLowestPrice gets a reference to the given AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsLowestPrice and assigns it to the LowestPrice field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetLowestPrice(v AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsLowestPrice) {
    o.LowestPrice = &v
}

// GetMainImageUrl returns the MainImageUrl field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetMainImageUrl() string {
    if o == nil || utils.IsNil(o.MainImageUrl) {
        var ret string
        return ret
    }
    return *o.MainImageUrl
}

// GetMainImageUrlOk returns a tuple with the MainImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetMainImageUrlOk() (*string, bool) {
    if o == nil || utils.IsNil(o.MainImageUrl) {
        return nil, false
    }
    return o.MainImageUrl, true
}

// HasMainImageUrl returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasMainImageUrl() bool {
    if o != nil && !utils.IsNil(o.MainImageUrl) {
        return true
    }

    return false
}

// SetMainImageUrl gets a reference to the given string and assigns it to the MainImageUrl field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetMainImageUrl(v string) {
    o.MainImageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetName() string {
    if o == nil || utils.IsNil(o.Name) {
        var ret string
        return ret
    }
    return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Name) {
        return nil, false
    }
    return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasName() bool {
    if o != nil && !utils.IsNil(o.Name) {
        return true
    }

    return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetName(v string) {
    o.Name = &v
}

// GetOpenCollaborationCommissionRate returns the OpenCollaborationCommissionRate field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetOpenCollaborationCommissionRate() int64 {
    if o == nil || utils.IsNil(o.OpenCollaborationCommissionRate) {
        var ret int64
        return ret
    }
    return *o.OpenCollaborationCommissionRate
}

// GetOpenCollaborationCommissionRateOk returns a tuple with the OpenCollaborationCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetOpenCollaborationCommissionRateOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.OpenCollaborationCommissionRate) {
        return nil, false
    }
    return o.OpenCollaborationCommissionRate, true
}

// HasOpenCollaborationCommissionRate returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasOpenCollaborationCommissionRate() bool {
    if o != nil && !utils.IsNil(o.OpenCollaborationCommissionRate) {
        return true
    }

    return false
}

// SetOpenCollaborationCommissionRate gets a reference to the given int64 and assigns it to the OpenCollaborationCommissionRate field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetOpenCollaborationCommissionRate(v int64) {
    o.OpenCollaborationCommissionRate = &v
}

// GetPartnerCommissionRate returns the PartnerCommissionRate field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetPartnerCommissionRate() int64 {
    if o == nil || utils.IsNil(o.PartnerCommissionRate) {
        var ret int64
        return ret
    }
    return *o.PartnerCommissionRate
}

// GetPartnerCommissionRateOk returns a tuple with the PartnerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetPartnerCommissionRateOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.PartnerCommissionRate) {
        return nil, false
    }
    return o.PartnerCommissionRate, true
}

// HasPartnerCommissionRate returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasPartnerCommissionRate() bool {
    if o != nil && !utils.IsNil(o.PartnerCommissionRate) {
        return true
    }

    return false
}

// SetPartnerCommissionRate gets a reference to the given int64 and assigns it to the PartnerCommissionRate field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetPartnerCommissionRate(v int64) {
    o.PartnerCommissionRate = &v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetProductDescription() string {
    if o == nil || utils.IsNil(o.ProductDescription) {
        var ret string
        return ret
    }
    return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetProductDescriptionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ProductDescription) {
        return nil, false
    }
    return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasProductDescription() bool {
    if o != nil && !utils.IsNil(o.ProductDescription) {
        return true
    }

    return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetProductDescription(v string) {
    o.ProductDescription = &v
}

// GetProductSales returns the ProductSales field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetProductSales() int64 {
    if o == nil || utils.IsNil(o.ProductSales) {
        var ret int64
        return ret
    }
    return *o.ProductSales
}

// GetProductSalesOk returns a tuple with the ProductSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetProductSalesOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.ProductSales) {
        return nil, false
    }
    return o.ProductSales, true
}

// HasProductSales returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasProductSales() bool {
    if o != nil && !utils.IsNil(o.ProductSales) {
        return true
    }

    return false
}

// SetProductSales gets a reference to the given int64 and assigns it to the ProductSales field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetProductSales(v int64) {
    o.ProductSales = &v
}

// GetReviewStatus returns the ReviewStatus field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetReviewStatus() string {
    if o == nil || utils.IsNil(o.ReviewStatus) {
        var ret string
        return ret
    }
    return *o.ReviewStatus
}

// GetReviewStatusOk returns a tuple with the ReviewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetReviewStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReviewStatus) {
        return nil, false
    }
    return o.ReviewStatus, true
}

// HasReviewStatus returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasReviewStatus() bool {
    if o != nil && !utils.IsNil(o.ReviewStatus) {
        return true
    }

    return false
}

// SetReviewStatus gets a reference to the given string and assigns it to the ReviewStatus field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetReviewStatus(v string) {
    o.ReviewStatus = &v
}

// GetSampleQuota returns the SampleQuota field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetSampleQuota() int64 {
    if o == nil || utils.IsNil(o.SampleQuota) {
        var ret int64
        return ret
    }
    return *o.SampleQuota
}

// GetSampleQuotaOk returns a tuple with the SampleQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetSampleQuotaOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.SampleQuota) {
        return nil, false
    }
    return o.SampleQuota, true
}

// HasSampleQuota returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasSampleQuota() bool {
    if o != nil && !utils.IsNil(o.SampleQuota) {
        return true
    }

    return false
}

// SetSampleQuota gets a reference to the given int64 and assigns it to the SampleQuota field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetSampleQuota(v int64) {
    o.SampleQuota = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetShopName() string {
    if o == nil || utils.IsNil(o.ShopName) {
        var ret string
        return ret
    }
    return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetShopNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ShopName) {
        return nil, false
    }
    return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasShopName() bool {
    if o != nil && !utils.IsNil(o.ShopName) {
        return true
    }

    return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetShopName(v string) {
    o.ShopName = &v
}

// GetSkuInformationList returns the SkuInformationList field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetSkuInformationList() []AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsSkuInformationList {
    if o == nil || utils.IsNil(o.SkuInformationList) {
        var ret []AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsSkuInformationList
        return ret
    }
    return o.SkuInformationList
}

// GetSkuInformationListOk returns a tuple with the SkuInformationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetSkuInformationListOk() ([]AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsSkuInformationList, bool) {
    if o == nil || utils.IsNil(o.SkuInformationList) {
        return nil, false
    }
    return o.SkuInformationList, true
}

// HasSkuInformationList returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasSkuInformationList() bool {
    if o != nil && !utils.IsNil(o.SkuInformationList) {
        return true
    }

    return false
}

// SetSkuInformationList gets a reference to the given []AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsSkuInformationList and assigns it to the SkuInformationList field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetSkuInformationList(v []AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProductsSkuInformationList) {
    o.SkuInformationList = v
}

// GetTotalCommissionRate returns the TotalCommissionRate field value if set, zero value otherwise.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetTotalCommissionRate() int64 {
    if o == nil || utils.IsNil(o.TotalCommissionRate) {
        var ret int64
        return ret
    }
    return *o.TotalCommissionRate
}

// GetTotalCommissionRateOk returns a tuple with the TotalCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) GetTotalCommissionRateOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.TotalCommissionRate) {
        return nil, false
    }
    return o.TotalCommissionRate, true
}

// HasTotalCommissionRate returns a boolean if a field has been set.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) HasTotalCommissionRate() bool {
    if o != nil && !utils.IsNil(o.TotalCommissionRate) {
        return true
    }

    return false
}

// SetTotalCommissionRate gets a reference to the given int64 and assigns it to the TotalCommissionRate field.
func (o *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) SetTotalCommissionRate(v int64) {
    o.TotalCommissionRate = &v
}

func (o AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Category) {
        toSerialize["category"] = o.Category
    }
    if !utils.IsNil(o.CreatorCommissionRate) {
        toSerialize["creator_commission_rate"] = o.CreatorCommissionRate
    }
    if !utils.IsNil(o.HighestPrice) {
        toSerialize["highest_price"] = o.HighestPrice
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Inventory) {
        toSerialize["inventory"] = o.Inventory
    }
    if !utils.IsNil(o.IsAvailable) {
        toSerialize["is_available"] = o.IsAvailable
    }
    if !utils.IsNil(o.LowestPrice) {
        toSerialize["lowest_price"] = o.LowestPrice
    }
    if !utils.IsNil(o.MainImageUrl) {
        toSerialize["main_image_url"] = o.MainImageUrl
    }
    if !utils.IsNil(o.Name) {
        toSerialize["name"] = o.Name
    }
    if !utils.IsNil(o.OpenCollaborationCommissionRate) {
        toSerialize["open_collaboration_commission_rate"] = o.OpenCollaborationCommissionRate
    }
    if !utils.IsNil(o.PartnerCommissionRate) {
        toSerialize["partner_commission_rate"] = o.PartnerCommissionRate
    }
    if !utils.IsNil(o.ProductDescription) {
        toSerialize["product_description"] = o.ProductDescription
    }
    if !utils.IsNil(o.ProductSales) {
        toSerialize["product_sales"] = o.ProductSales
    }
    if !utils.IsNil(o.ReviewStatus) {
        toSerialize["review_status"] = o.ReviewStatus
    }
    if !utils.IsNil(o.SampleQuota) {
        toSerialize["sample_quota"] = o.SampleQuota
    }
    if !utils.IsNil(o.ShopName) {
        toSerialize["shop_name"] = o.ShopName
    }
    if !utils.IsNil(o.SkuInformationList) {
        toSerialize["sku_information_list"] = o.SkuInformationList
    }
    if !utils.IsNil(o.TotalCommissionRate) {
        toSerialize["total_commission_rate"] = o.TotalCommissionRate
    }
    return toSerialize, nil
}

type NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts struct {
	value *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts
	isSet bool
}

func (v NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) Get() *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts {
	return v.value
}

func (v *NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) Set(val *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts(val *AffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) *NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts {
	return &NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts{value: val, isSet: true}
}

func (v NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffiliatePartner202405GetAffiliatePartnerCampaignProductListResponseDataProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


