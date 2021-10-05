# \GlobalTurnServiceApi

All URIs are relative to *https://api.subspace.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalTurnServiceGetGlobalTurn**](GlobalTurnServiceApi.md#GlobalTurnServiceGetGlobalTurn) | **Post** /v1/globalturn | 



## GlobalTurnServiceGetGlobalTurn

> V1GlobalTurnResponse GlobalTurnServiceGetGlobalTurn(ctx).Execute()



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalTurnServiceApi.GlobalTurnServiceGetGlobalTurn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTurnServiceApi.GlobalTurnServiceGetGlobalTurn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalTurnServiceGetGlobalTurn`: V1GlobalTurnResponse
    fmt.Fprintf(os.Stdout, "Response from `GlobalTurnServiceApi.GlobalTurnServiceGetGlobalTurn`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalTurnServiceGetGlobalTurnRequest struct via the builder pattern


### Return type

[**V1GlobalTurnResponse**](V1GlobalTurnResponse.md)

### Authorization

[accessCode](../README.md#accessCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

