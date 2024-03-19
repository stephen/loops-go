# TransactionalFailureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionalFailureResponse

`func NewTransactionalFailureResponse() *TransactionalFailureResponse`

NewTransactionalFailureResponse instantiates a new TransactionalFailureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalFailureResponseWithDefaults

`func NewTransactionalFailureResponseWithDefaults() *TransactionalFailureResponse`

NewTransactionalFailureResponseWithDefaults instantiates a new TransactionalFailureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TransactionalFailureResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TransactionalFailureResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TransactionalFailureResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TransactionalFailureResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetPath

`func (o *TransactionalFailureResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TransactionalFailureResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TransactionalFailureResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TransactionalFailureResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMessage

`func (o *TransactionalFailureResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionalFailureResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionalFailureResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionalFailureResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


