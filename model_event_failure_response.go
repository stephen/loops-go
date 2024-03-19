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

// checks if the EventFailureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFailureResponse{}

// EventFailureResponse struct for EventFailureResponse
type EventFailureResponse struct {
	Success *bool `json:"success,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewEventFailureResponse instantiates a new EventFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFailureResponse() *EventFailureResponse {
	this := EventFailureResponse{}
	return &this
}

// NewEventFailureResponseWithDefaults instantiates a new EventFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFailureResponseWithDefaults() *EventFailureResponse {
	this := EventFailureResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *EventFailureResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFailureResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *EventFailureResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *EventFailureResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventFailureResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFailureResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventFailureResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventFailureResponse) SetMessage(v string) {
	o.Message = &v
}

func (o EventFailureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventFailureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableEventFailureResponse struct {
	value *EventFailureResponse
	isSet bool
}

func (v NullableEventFailureResponse) Get() *EventFailureResponse {
	return v.value
}

func (v *NullableEventFailureResponse) Set(val *EventFailureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFailureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFailureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFailureResponse(val *EventFailureResponse) *NullableEventFailureResponse {
	return &NullableEventFailureResponse{value: val, isSet: true}
}

func (v NullableEventFailureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFailureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


