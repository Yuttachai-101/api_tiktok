/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apis

import (
    "bytes"
    "context"
    "io"
    "net/http"
    "net/url"

    "github.com/Yuttachai-101/api_tiktok/models/affiliate_partner/v202411"
)


// AffiliatePartnerV202411APIService AffiliatePartnerV202411API service
type AffiliatePartnerV202411APIService service

type ApiAffiliatePartner202411OrdersSearchPostRequest struct {
    ctx context.Context
    ApiService *AffiliatePartnerV202411APIService
    pageSize *int64
    categoryAssetCipher *string
    xTtsAccessToken *string
    contentType *string
    pageToken *string
    affiliatePartner202411SearchTapAffiliateOrdersRequestBody *affiliate_partner_v202411.AffiliatePartner202411SearchTapAffiliateOrdersRequestBody
}

// The default is 20, it must be positive integer, the range is 1-100
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) PageSize(pageSize int64) ApiAffiliatePartner202411OrdersSearchPostRequest {
    r.pageSize = &pageSize
    return r
}
// The partner identifier used in API requests.  Retrieve this value by using the [Get Authorized Category Assets API] (https://partner.tiktokshop.com/docv2/page/666012dd609d4402cc3be995). 
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) CategoryAssetCipher(categoryAssetCipher string) ApiAffiliatePartner202411OrdersSearchPostRequest {
    r.categoryAssetCipher = &categoryAssetCipher
    return r
}
// 
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) XTtsAccessToken(xTtsAccessToken string) ApiAffiliatePartner202411OrdersSearchPostRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) ContentType(contentType string) ApiAffiliatePartner202411OrdersSearchPostRequest {
    r.contentType = &contentType
    return r
}
// The next page token
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) PageToken(pageToken string) ApiAffiliatePartner202411OrdersSearchPostRequest {
    r.pageToken = &pageToken
    return r
}
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) AffiliatePartner202411SearchTapAffiliateOrdersRequestBody(affiliatePartner202411SearchTapAffiliateOrdersRequestBody affiliate_partner_v202411.AffiliatePartner202411SearchTapAffiliateOrdersRequestBody) ApiAffiliatePartner202411OrdersSearchPostRequest {
    r.affiliatePartner202411SearchTapAffiliateOrdersRequestBody = &affiliatePartner202411SearchTapAffiliateOrdersRequestBody
    return r
}
func (r ApiAffiliatePartner202411OrdersSearchPostRequest) Execute() (*affiliate_partner_v202411.AffiliatePartner202411SearchTapAffiliateOrdersResponse, *http.Response, error) {
    return r.ApiService.AffiliatePartner202411OrdersSearchPostExecute(r)
}

/*
AffiliatePartner202411OrdersSearchPost SearchTapAffiliateOrders
TAP can use this API to retrieve a list of affiliate orders and track the affiliate conversions.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiAffiliatePartner202411OrdersSearchPostRequest
*/
func (a *AffiliatePartnerV202411APIService) AffiliatePartner202411OrdersSearchPost(ctx context.Context) ApiAffiliatePartner202411OrdersSearchPostRequest {
    return ApiAffiliatePartner202411OrdersSearchPostRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return AffiliatePartner202411SearchTapAffiliateOrdersResponse
func (a *AffiliatePartnerV202411APIService) AffiliatePartner202411OrdersSearchPostExecute(r ApiAffiliatePartner202411OrdersSearchPostRequest) (*affiliate_partner_v202411.AffiliatePartner202411SearchTapAffiliateOrdersResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodPost
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *affiliate_partner_v202411.AffiliatePartner202411SearchTapAffiliateOrdersResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AffiliatePartnerV202411APIService.AffiliatePartner202411OrdersSearchPost")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/affiliate_partner/202411/orders/search"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.pageSize == nil {
        return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
    }
    if r.categoryAssetCipher == nil {
        return localVarReturnValue, nil, reportError("categoryAssetCipher is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    parameterAddToHeaderOrQuery(localVarQueryParams, "category_asset_cipher", r.categoryAssetCipher, "")
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{"application/json"}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    // body params
    localVarPostBody = r.affiliatePartner202411SearchTapAffiliateOrdersRequestBody
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}
