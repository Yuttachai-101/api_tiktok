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

    "github.com/Yuttachai-101/api_tiktok/models/affiliate/v202402"
)


// AffiliateV202402APIService AffiliateV202402API service
type AffiliateV202402APIService service

type ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest struct {
    ctx context.Context
    ApiService *AffiliateV202402APIService
    xTtsAccessToken *string
    contentType *string
    affiliate202402CheckAnchorPrerequisitesRequestBody *affiliate_v202402.Affiliate202402CheckAnchorPrerequisitesRequestBody
}

// 
func (r ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest) XTtsAccessToken(xTtsAccessToken string) ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest) ContentType(contentType string) ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest {
    r.contentType = &contentType
    return r
}
func (r ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest) Affiliate202402CheckAnchorPrerequisitesRequestBody(affiliate202402CheckAnchorPrerequisitesRequestBody affiliate_v202402.Affiliate202402CheckAnchorPrerequisitesRequestBody) ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest {
    r.affiliate202402CheckAnchorPrerequisitesRequestBody = &affiliate202402CheckAnchorPrerequisitesRequestBody
    return r
}
func (r ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest) Execute() (*affiliate_v202402.Affiliate202402CheckAnchorPrerequisitesResponse, *http.Response, error) {
    return r.ApiService.Affiliate202402AnchorsPrerequisiteCheckPostExecute(r)
}

/*
Affiliate202402AnchorsPrerequisiteCheckPost CheckAnchorPrerequisites
The is  a pre-verification interface for creator adding products to video. This interface will verify the creator's permissions and product status, etc.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest
*/
func (a *AffiliateV202402APIService) Affiliate202402AnchorsPrerequisiteCheckPost(ctx context.Context) ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest {
    return ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Affiliate202402CheckAnchorPrerequisitesResponse
func (a *AffiliateV202402APIService) Affiliate202402AnchorsPrerequisiteCheckPostExecute(r ApiAffiliate202402AnchorsPrerequisiteCheckPostRequest) (*affiliate_v202402.Affiliate202402CheckAnchorPrerequisitesResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodPost
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *affiliate_v202402.Affiliate202402CheckAnchorPrerequisitesResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AffiliateV202402APIService.Affiliate202402AnchorsPrerequisiteCheckPost")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/affiliate/202402/anchors/prerequisite_check"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

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
    localVarPostBody = r.affiliate202402CheckAnchorPrerequisitesRequestBody
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
