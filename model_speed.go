/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Speed struct for Speed
type Speed struct {
	Label string `json:"label"`
	Value int32 `json:"value"`
}

// NewSpeed instantiates a new Speed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpeed(label string, value int32) *Speed {
	this := Speed{}
	this.Label = label
	this.Value = value
	return &this
}

// NewSpeedWithDefaults instantiates a new Speed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpeedWithDefaults() *Speed {
	this := Speed{}
	return &this
}

// GetLabel returns the Label field value
func (o *Speed) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Speed) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Speed) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Speed) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Speed) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Speed) SetValue(v int32) {
	o.Value = v
}

func (o Speed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSpeed struct {
	value *Speed
	isSet bool
}

func (v NullableSpeed) Get() *Speed {
	return v.value
}

func (v *NullableSpeed) Set(val *Speed) {
	v.value = val
	v.isSet = true
}

func (v NullableSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpeed(val *Speed) *NullableSpeed {
	return &NullableSpeed{value: val, isSet: true}
}

func (v NullableSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

