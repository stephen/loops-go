# \TransactionalEmailsAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionalPost**](TransactionalEmailsAPI.md#TransactionalPost) | **Post** /transactional | Send a transactional email



## TransactionalPost

> TransactionalSuccessResponse TransactionalPost(ctx).TransactionalRequest(transactionalRequest).Execute()

Send a transactional email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/stephen/loops-go"
)

func main() {
	transactionalRequest := *openapiclient.NewTransactionalRequest("Email_example", "TransactionalId_example") // TransactionalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionalEmailsAPI.TransactionalPost(context.Background()).TransactionalRequest(transactionalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionalEmailsAPI.TransactionalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionalPost`: TransactionalSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionalEmailsAPI.TransactionalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionalRequest** | [**TransactionalRequest**](TransactionalRequest.md) |  | 

### Return type

[**TransactionalSuccessResponse**](TransactionalSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

