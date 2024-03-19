# EventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**EventName** | **string** |  | 

## Methods

### NewEventRequest

`func NewEventRequest(eventName string, ) *EventRequest`

NewEventRequest instantiates a new EventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRequestWithDefaults

`func NewEventRequestWithDefaults() *EventRequest`

NewEventRequestWithDefaults instantiates a new EventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EventRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EventRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EventRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EventRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUserId

`func (o *EventRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEventName

`func (o *EventRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


