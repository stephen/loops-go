# \ContactsAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsCreatePost**](ContactsAPI.md#ContactsCreatePost) | **Post** /contacts/create | Create a contact
[**ContactsDeletePost**](ContactsAPI.md#ContactsDeletePost) | **Post** /contacts/delete | Delete a contact
[**ContactsFindGet**](ContactsAPI.md#ContactsFindGet) | **Get** /contacts/find | Find a contact
[**ContactsUpdatePut**](ContactsAPI.md#ContactsUpdatePut) | **Put** /contacts/update | Update a contact



## ContactsCreatePost

> ContactSuccessResponse ContactsCreatePost(ctx).ContactRequest(contactRequest).Execute()

Create a contact



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
	contactRequest := *openapiclient.NewContactRequest("Email_example") // ContactRequest | You can add custom contact properties as keys in this request (of type `string`, `number`, `boolean` or `date` ([see available date formats](https://loops.so/docs/contacts/properties#dates))).<br>Make sure to create the properties in Loops before using them in API calls.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.ContactsCreatePost(context.Background()).ContactRequest(contactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.ContactsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsCreatePost`: ContactSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.ContactsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactRequest** | [**ContactRequest**](ContactRequest.md) | You can add custom contact properties as keys in this request (of type &#x60;string&#x60;, &#x60;number&#x60;, &#x60;boolean&#x60; or &#x60;date&#x60; ([see available date formats](https://loops.so/docs/contacts/properties#dates))).&lt;br&gt;Make sure to create the properties in Loops before using them in API calls. | 

### Return type

[**ContactSuccessResponse**](ContactSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsDeletePost

> ContactDeleteResponse ContactsDeletePost(ctx).ContactDeleteRequest(contactDeleteRequest).Execute()

Delete a contact



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
	contactDeleteRequest := *openapiclient.NewContactDeleteRequest() // ContactDeleteRequest | Include only one of `email` or `userId`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.ContactsDeletePost(context.Background()).ContactDeleteRequest(contactDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.ContactsDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsDeletePost`: ContactDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.ContactsDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactDeleteRequest** | [**ContactDeleteRequest**](ContactDeleteRequest.md) | Include only one of &#x60;email&#x60; or &#x60;userId&#x60;. | 

### Return type

[**ContactDeleteResponse**](ContactDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsFindGet

> []Contact ContactsFindGet(ctx).Email(email).Execute()

Find a contact



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
	email := "email_example" // string | Email address (URI-encoded)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.ContactsFindGet(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.ContactsFindGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsFindGet`: []Contact
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.ContactsFindGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsFindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email address (URI-encoded) | 

### Return type

[**[]Contact**](Contact.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsUpdatePut

> ContactSuccessResponse ContactsUpdatePut(ctx).ContactRequest(contactRequest).Execute()

Update a contact



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
	contactRequest := *openapiclient.NewContactRequest("Email_example") // ContactRequest | You can add custom contact properties as keys in this request (of type `string`, `number`, `boolean` or `date` ([see available date formats](https://loops.so/docs/contacts/properties#dates))).<br>Make sure to create the properties in Loops before using them in API calls.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.ContactsUpdatePut(context.Background()).ContactRequest(contactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.ContactsUpdatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactsUpdatePut`: ContactSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.ContactsUpdatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactRequest** | [**ContactRequest**](ContactRequest.md) | You can add custom contact properties as keys in this request (of type &#x60;string&#x60;, &#x60;number&#x60;, &#x60;boolean&#x60; or &#x60;date&#x60; ([see available date formats](https://loops.so/docs/contacts/properties#dates))).&lt;br&gt;Make sure to create the properties in Loops before using them in API calls. | 

### Return type

[**ContactSuccessResponse**](ContactSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

