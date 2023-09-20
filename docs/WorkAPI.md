# \WorkAPI

All URIs are relative to *https://api.acmi.net.au*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWork**](WorkAPI.md#GetWork) | **Get** /works/{id}/ | 



## GetWork

> Work GetWork(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/AndrewMcCraeCA/goacmi"
)

func main() {
    id := int32(56) // int32 | Work ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkAPI.GetWork(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkAPI.GetWork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWork`: Work
    fmt.Fprintf(os.Stdout, "Response from `WorkAPI.GetWork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Work ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Work**](Work.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

