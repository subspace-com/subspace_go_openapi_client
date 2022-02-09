# \WebRtcCdnServiceApi

All URIs are relative to *https://api.subspace.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebRtcCdnServiceGetWebRtcCdn**](WebRtcCdnServiceApi.md#WebRtcCdnServiceGetWebRtcCdn) | **Post** /v1/webrtc-cdn | 



## WebRtcCdnServiceGetWebRtcCdn

> V1WebRtcCdnResponse WebRtcCdnServiceGetWebRtcCdn(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebRtcCdnServiceApi.WebRtcCdnServiceGetWebRtcCdn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebRtcCdnServiceApi.WebRtcCdnServiceGetWebRtcCdn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebRtcCdnServiceGetWebRtcCdn`: V1WebRtcCdnResponse
    fmt.Fprintf(os.Stdout, "Response from `WebRtcCdnServiceApi.WebRtcCdnServiceGetWebRtcCdn`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebRtcCdnServiceGetWebRtcCdnRequest struct via the builder pattern


### Return type

[**V1WebRtcCdnResponse**](V1WebRtcCdnResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

