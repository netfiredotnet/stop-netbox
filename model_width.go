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

// Width struct for Width
type Width struct {
	Label string `json:"label"`
	Value int32 `json:"value"`
}

// NewWidth instantiates a new Width object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidth(label string, value int32) *Width {
	this := Width{}
	this.Label = label
	this.Value = value
	return &this
}

// NewWidthWithDefaults instantiates a new Width object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidthWithDefaults() *Width {
	this := Width{}
	return &this
}

// GetLabel returns the Label field value
func (o *Width) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Width) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Width) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Width) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Width) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Width) SetValue(v int32) {
	o.Value = v
}

func (o Width) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWidth struct {
	value *Width
	isSet bool
}

func (v NullableWidth) Get() *Width {
	return v.value
}

func (v *NullableWidth) Set(val *Width) {
	v.value = val
	v.isSet = true
}

func (v NullableWidth) IsSet() bool {
	return v.isSet
}

func (v *NullableWidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidth(val *Width) *NullableWidth {
	return &NullableWidth{value: val, isSet: true}
}

func (v NullableWidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

