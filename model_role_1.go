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

// Role1 struct for Role1
type Role1 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// NewRole1 instantiates a new Role1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole1(label string, value string) *Role1 {
	this := Role1{}
	this.Label = label
	this.Value = value
	return &this
}

// NewRole1WithDefaults instantiates a new Role1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRole1WithDefaults() *Role1 {
	this := Role1{}
	return &this
}

// GetLabel returns the Label field value
func (o *Role1) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Role1) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Role1) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Role1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Role1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Role1) SetValue(v string) {
	o.Value = v
}

func (o Role1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRole1 struct {
	value *Role1
	isSet bool
}

func (v NullableRole1) Get() *Role1 {
	return v.value
}

func (v *NullableRole1) Set(val *Role1) {
	v.value = val
	v.isSet = true
}

func (v NullableRole1) IsSet() bool {
	return v.isSet
}

func (v *NullableRole1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole1(val *Role1) *NullableRole1 {
	return &NullableRole1{value: val, isSet: true}
}

func (v NullableRole1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


