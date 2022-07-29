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

// Status1 struct for Status1
type Status1 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// NewStatus1 instantiates a new Status1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus1(label string, value string) *Status1 {
	this := Status1{}
	this.Label = label
	this.Value = value
	return &this
}

// NewStatus1WithDefaults instantiates a new Status1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatus1WithDefaults() *Status1 {
	this := Status1{}
	return &this
}

// GetLabel returns the Label field value
func (o *Status1) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Status1) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Status1) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Status1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Status1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Status1) SetValue(v string) {
	o.Value = v
}

func (o Status1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatus1 struct {
	value *Status1
	isSet bool
}

func (v NullableStatus1) Get() *Status1 {
	return v.value
}

func (v *NullableStatus1) Set(val *Status1) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus1) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus1(val *Status1) *NullableStatus1 {
	return &NullableStatus1{value: val, isSet: true}
}

func (v NullableStatus1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


