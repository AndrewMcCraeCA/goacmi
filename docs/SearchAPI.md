# \SearchAPI

All URIs are relative to *https://api.acmi.net.au*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchAPI.md#Search) | **Get** /search/ | 



## Search

> []Work Search(ctx).Query(query).Field(field).Size(size).Page(page).Raw(raw).Execute()



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
    query := "query_example" // string | Search query (optional)
    field := "field_example" // string | Limit search queries to a Work field. e.g. title (optional)
    size := int32(56) // int32 | Limit search result page size. (optional) (default to 20)
    page := int32(56) // int32 | Retrieve a parcitular page of serach results. (optional)
    raw := "raw_example" // string | Return the raw Elasticsearch results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.Search(context.Background()).Query(query).Field(field).Size(size).Page(page).Raw(raw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: []Work
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Search query | 
 **field** | **string** | Limit search queries to a Work field. e.g. title | 
 **size** | **int32** | Limit search result page size. | [default to 20]
 **page** | **int32** | Retrieve a parcitular page of serach results. | 
 **raw** | **string** | Return the raw Elasticsearch results. | 

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

