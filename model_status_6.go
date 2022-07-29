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

// Status6 struct for Status6
type Status6 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// NewStatus6 instantiates a new Status6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus6(label string, value string) *Status6 {
	this := Status6{}
	this.Label = label
	this.Value = value
	return &this
}

// NewStatus6WithDefaults instantiates a new Status6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatus6WithDefaults() *Status6 {
	this := Status6{}
	return &this
}

// GetLabel returns the Label field value
func (o *Status6) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Status6) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Status6) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Status6) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Status6) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Status6) SetValue(v string) {
	o.Value = v
}

func (o Status6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatus6 struct {
	value *Status6
	isSet bool
}

func (v NullableStatus6) Get() *Status6 {
	return v.value
}

func (v *NullableStatus6) Set(val *Status6) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus6) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus6(val *Status6) *NullableStatus6 {
	return &NullableStatus6{value: val, isSet: true}
}

func (v NullableStatus6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


