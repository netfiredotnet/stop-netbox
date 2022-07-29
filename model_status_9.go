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

// Status9 struct for Status9
type Status9 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// NewStatus9 instantiates a new Status9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus9(label string, value string) *Status9 {
	this := Status9{}
	this.Label = label
	this.Value = value
	return &this
}

// NewStatus9WithDefaults instantiates a new Status9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatus9WithDefaults() *Status9 {
	this := Status9{}
	return &this
}

// GetLabel returns the Label field value
func (o *Status9) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Status9) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Status9) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Status9) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Status9) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Status9) SetValue(v string) {
	o.Value = v
}

func (o Status9) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatus9 struct {
	value *Status9
	isSet bool
}

func (v NullableStatus9) Get() *Status9 {
	return v.value
}

func (v *NullableStatus9) Set(val *Status9) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus9) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus9(val *Status9) *NullableStatus9 {
	return &NullableStatus9{value: val, isSet: true}
}

func (v NullableStatus9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


