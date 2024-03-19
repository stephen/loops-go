# \CustomFieldsAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsCustomFieldsGet**](CustomFieldsAPI.md#ContactsCustomFieldsGet) | **Get** /contacts/customFields | Get a list of custom contact properties



## ContactsCustomFieldsGet

> []CustomField ContactsCustomFieldsGet(ctx).Execute()

Get a list of custom contact properties



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.ContactsCustomFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.ContactsCustomFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsCustomFieldsGet`: []CustomField
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.ContactsCustomFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiContactsCustomFieldsGetRequest struct via the builder pattern


### Return type

[**[]CustomField**](CustomField.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

