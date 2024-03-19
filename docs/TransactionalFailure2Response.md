# TransactionalFailure2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to [**TransactionalFailure2ResponseError**](TransactionalFailure2ResponseError.md) |  | [optional] 

## Methods

### NewTransactionalFailure2Response

`func NewTransactionalFailure2Response() *TransactionalFailure2Response`

NewTransactionalFailure2Response instantiates a new TransactionalFailure2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalFailure2ResponseWithDefaults

`func NewTransactionalFailure2ResponseWithDefaults() *TransactionalFailure2Response`

NewTransactionalFailure2ResponseWithDefaults instantiates a new TransactionalFailure2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TransactionalFailure2Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TransactionalFailure2Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TransactionalFailure2Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TransactionalFailure2Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *TransactionalFailure2Response) GetError() TransactionalFailure2ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TransactionalFailure2Response) GetErrorOk() (*TransactionalFailure2ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TransactionalFailure2Response) SetError(v TransactionalFailure2ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *TransactionalFailure2Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


