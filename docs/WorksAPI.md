# \WorksAPI

All URIs are relative to *https://api.acmi.net.au*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorks**](WorksAPI.md#GetWorks) | **Get** /works/ | 



## GetWorks

> []Work GetWorks(ctx).Page(page).Execute()



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
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorksAPI.GetWorks(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorksAPI.GetWorks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorks`: []Work
    fmt.Fprintf(os.Stdout, "Response from `WorksAPI.GetWorks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 

### Return type

[**[]Work**](Work.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

