# TransactionalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**TransactionalId** | **string** | The ID of the transactional email to send. | 
**DataVariables** | Pointer to **map[string]interface{}** | An object containing contact data as defined by the data variables added to the transactional email template. | [optional] 

## Methods

### NewTransactionalRequest

`func NewTransactionalRequest(email string, transactionalId string, ) *TransactionalRequest`

NewTransactionalRequest instantiates a new TransactionalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalRequestWithDefaults

`func NewTransactionalRequestWithDefaults() *TransactionalRequest`

NewTransactionalRequestWithDefaults instantiates a new TransactionalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TransactionalRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TransactionalRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TransactionalRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTransactionalId

`func (o *TransactionalRequest) GetTransactionalId() string`

GetTransactionalId returns the TransactionalId field if non-nil, zero value otherwise.

### GetTransactionalIdOk

`func (o *TransactionalRequest) GetTransactionalIdOk() (*string, bool)`

GetTransactionalIdOk returns a tuple with the TransactionalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionalId

`func (o *TransactionalRequest) SetTransactionalId(v string)`

SetTransactionalId sets TransactionalId field to given value.


### GetDataVariables

`func (o *TransactionalRequest) GetDataVariables() map[string]interface{}`

GetDataVariables returns the DataVariables field if non-nil, zero value otherwise.

### GetDataVariablesOk

`func (o *TransactionalRequest) GetDataVariablesOk() (*map[string]interface{}, bool)`

GetDataVariablesOk returns a tuple with the DataVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVariables

`func (o *TransactionalRequest) SetDataVariables(v map[string]interface{})`

SetDataVariables sets DataVariables field to given value.

### HasDataVariables

`func (o *TransactionalRequest) HasDataVariables() bool`

HasDataVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


