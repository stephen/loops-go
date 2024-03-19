# \EventsAPI

All URIs are relative to *https://app.loops.so/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsSendPost**](EventsAPI.md#EventsSendPost) | **Post** /events/send | Send an event



## EventsSendPost

> EventSuccessResponse EventsSendPost(ctx).EventRequest(eventRequest).Execute()

Send an event



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
	eventRequest := *openapiclient.NewEventRequest("EventName_example") // EventRequest | Provide either `email` or `userId` to identify the contact ([read more](https://loops.so/docs/api-reference/send-event#body)).<br>You can add custom contact properties as keys in this request (of type `string`, `number`, `boolean` or `date` ([see available date formats](https://loops.so/docs/contacts/properties#dates))).<br>Make sure to create the properties in Loops before using them in API calls.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsSendPost(context.Background()).EventRequest(eventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsSendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsSendPost`: EventSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsSendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventRequest** | [**EventRequest**](EventRequest.md) | Provide either &#x60;email&#x60; or &#x60;userId&#x60; to identify the contact ([read more](https://loops.so/docs/api-reference/send-event#body)).&lt;br&gt;You can add custom contact properties as keys in this request (of type &#x60;string&#x60;, &#x60;number&#x60;, &#x60;boolean&#x60; or &#x60;date&#x60; ([see available date formats](https://loops.so/docs/contacts/properties#dates))).&lt;br&gt;Make sure to create the properties in Loops before using them in API calls. | 

### Return type

[**EventSuccessResponse**](EventSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

