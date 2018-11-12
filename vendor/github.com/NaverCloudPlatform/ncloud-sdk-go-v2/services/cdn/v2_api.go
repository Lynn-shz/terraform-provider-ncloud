/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-10-18T06:18:35Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"reflect"
	"strings"
	"bytes"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type V2ApiService service


/* V2ApiService 
 CDN+인스턴스리스트조회
 @param getCdnPlusInstanceListRequest getCdnPlusInstanceListRequest
 @return *GetCdnPlusInstanceListResponse*/
func (a *V2ApiService) GetCdnPlusInstanceList(getCdnPlusInstanceListRequest *GetCdnPlusInstanceListRequest) (*GetCdnPlusInstanceListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetCdnPlusInstanceListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getCdnPlusInstanceList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getCdnPlusInstanceListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if !strings.Contains(string(bodyBytes), `{"error"`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}


	return &successPayload, err
}

/* V2ApiService 
 CDN+퍼지기록조회
 @param getCdnPlusPurgeHistoryListRequest getCdnPlusPurgeHistoryListRequest
 @return *GetCdnPlusPurgeHistoryListResponse*/
func (a *V2ApiService) GetCdnPlusPurgeHistoryList(getCdnPlusPurgeHistoryListRequest *GetCdnPlusPurgeHistoryListRequest) (*GetCdnPlusPurgeHistoryListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetCdnPlusPurgeHistoryListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getCdnPlusPurgeHistoryList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getCdnPlusPurgeHistoryListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if !strings.Contains(string(bodyBytes), `{"error"`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}


	return &successPayload, err
}

/* V2ApiService 
 Global CDN 인스턴스리스트조회
 @param getGlobalCdnInstanceListRequest getGlobalCdnInstanceListRequest
 @return *GetGlobalCdnInstanceListResponse*/
func (a *V2ApiService) GetGlobalCdnInstanceList(getGlobalCdnInstanceListRequest *GetGlobalCdnInstanceListRequest) (*GetGlobalCdnInstanceListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetGlobalCdnInstanceListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getGlobalCdnInstanceList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getGlobalCdnInstanceListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if !strings.Contains(string(bodyBytes), `{"error"`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}


	return &successPayload, err
}

/* V2ApiService 
 Global CDN퍼지기록조회
 @param getGlobalCdnPurgeHistoryListRequest getGlobalCdnPurgeHistoryListRequest
 @return *GetGlobalCdnPurgeHistoryListResponse*/
func (a *V2ApiService) GetGlobalCdnPurgeHistoryList(getGlobalCdnPurgeHistoryListRequest *GetGlobalCdnPurgeHistoryListRequest) (*GetGlobalCdnPurgeHistoryListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetGlobalCdnPurgeHistoryListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getGlobalCdnPurgeHistoryList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getGlobalCdnPurgeHistoryListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if !strings.Contains(string(bodyBytes), `{"error"`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}


	return &successPayload, err
}

/* V2ApiService 
 CDN+퍼지요청
 @param requestCdnPlusPurgeRequest requestCdnPlusPurgeRequest
 @return *RequestCdnPlusPurgeResponse*/
func (a *V2ApiService) RequestCdnPlusPurge(requestCdnPlusPurgeRequest *RequestCdnPlusPurgeRequest) (*RequestCdnPlusPurgeResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  RequestCdnPlusPurgeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/requestCdnPlusPurge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = requestCdnPlusPurgeRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if !strings.Contains(string(bodyBytes), `{"error"`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}


	return &successPayload, err
}

/* V2ApiService 
 Global CDN퍼지요청
 @param requestGlobalCdnPurgeRequest requestGlobalCdnPurgeRequest
 @return *RequestGlobalCdnPurgeResponse*/
func (a *V2ApiService) RequestGlobalCdnPurge(requestGlobalCdnPurgeRequest *RequestGlobalCdnPurgeRequest) (*RequestGlobalCdnPurgeResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  RequestGlobalCdnPurgeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/requestGlobalCdnPurge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = requestGlobalCdnPurgeRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if !strings.Contains(string(bodyBytes), `{"error"`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}


	return &successPayload, err
}

