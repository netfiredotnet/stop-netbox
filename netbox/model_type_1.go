/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// Type1 struct for Type1
type Type1 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// NewType1 instantiates a new Type1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewType1(label string, value string) *Type1 {
	this := Type1{}
	this.Label = label
	this.Value = value
	return &this
}

// NewType1WithDefaults instantiates a new Type1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewType1WithDefaults() *Type1 {
	this := Type1{}
	return &this
}

// GetLabel returns the Label field value
func (o *Type1) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Type1) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Type1) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Type1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Type1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Type1) SetValue(v string) {
	o.Value = v
}

func (o Type1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableType1 struct {
	value *Type1
	isSet bool
}

func (v NullableType1) Get() *Type1 {
	return v.value
}

func (v *NullableType1) Set(val *Type1) {
	v.value = val
	v.isSet = true
}

func (v NullableType1) IsSet() bool {
	return v.isSet
}

func (v *NullableType1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableType1(val *Type1) *NullableType1 {
	return &NullableType1{value: val, isSet: true}
}

func (v NullableType1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableType1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


