/*
Loops OpenAPI Spec

This is the OpenAPI Spec for the [Loops API](https://loops.so/docs/api).

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loops

import (
	"encoding/json"
)

// checks if the ContactsDeletePost500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactsDeletePost500Response{}

// ContactsDeletePost500Response struct for ContactsDeletePost500Response
type ContactsDeletePost500Response struct {
	Success *bool `json:"success,omitempty"`
}

// NewContactsDeletePost500Response instantiates a new ContactsDeletePost500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactsDeletePost500Response() *ContactsDeletePost500Response {
	this := ContactsDeletePost500Response{}
	return &this
}

// NewContactsDeletePost500ResponseWithDefaults instantiates a new ContactsDeletePost500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactsDeletePost500ResponseWithDefaults() *ContactsDeletePost500Response {
	this := ContactsDeletePost500Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ContactsDeletePost500Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsDeletePost500Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ContactsDeletePost500Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ContactsDeletePost500Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o ContactsDeletePost500Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactsDeletePost500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableContactsDeletePost500Response struct {
	value *ContactsDeletePost500Response
	isSet bool
}

func (v NullableContactsDeletePost500Response) Get() *ContactsDeletePost500Response {
	return v.value
}

func (v *NullableContactsDeletePost500Response) Set(val *ContactsDeletePost500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableContactsDeletePost500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableContactsDeletePost500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactsDeletePost500Response(val *ContactsDeletePost500Response) *NullableContactsDeletePost500Response {
	return &NullableContactsDeletePost500Response{value: val, isSet: true}
}

func (v NullableContactsDeletePost500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactsDeletePost500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


