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

// Family struct for Family
type Family struct {
	Label string `json:"label"`
	Value int32 `json:"value"`
}

// NewFamily instantiates a new Family object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFamily(label string, value int32) *Family {
	this := Family{}
	this.Label = label
	this.Value = value
	return &this
}

// NewFamilyWithDefaults instantiates a new Family object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFamilyWithDefaults() *Family {
	this := Family{}
	return &this
}

// GetLabel returns the Label field value
func (o *Family) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Family) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Family) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Family) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Family) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Family) SetValue(v int32) {
	o.Value = v
}

func (o Family) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFamily struct {
	value *Family
	isSet bool
}

func (v NullableFamily) Get() *Family {
	return v.value
}

func (v *NullableFamily) Set(val *Family) {
	v.value = val
	v.isSet = true
}

func (v NullableFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFamily(val *Family) *NullableFamily {
	return &NullableFamily{value: val, isSet: true}
}

func (v NullableFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

