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

// DcimPowerOutletTemplatesList200Response struct for DcimPowerOutletTemplatesList200Response
type DcimPowerOutletTemplatesList200Response struct {
	Count int32 `json:"count"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []PowerOutletTemplate `json:"results"`
}

// NewDcimPowerOutletTemplatesList200Response instantiates a new DcimPowerOutletTemplatesList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDcimPowerOutletTemplatesList200Response(count int32, results []PowerOutletTemplate) *DcimPowerOutletTemplatesList200Response {
	this := DcimPowerOutletTemplatesList200Response{}
	this.Count = count
	this.Results = results
	return &this
}

// NewDcimPowerOutletTemplatesList200ResponseWithDefaults instantiates a new DcimPowerOutletTemplatesList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDcimPowerOutletTemplatesList200ResponseWithDefaults() *DcimPowerOutletTemplatesList200Response {
	this := DcimPowerOutletTemplatesList200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *DcimPowerOutletTemplatesList200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DcimPowerOutletTemplatesList200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DcimPowerOutletTemplatesList200Response) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DcimPowerOutletTemplatesList200Response) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DcimPowerOutletTemplatesList200Response) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *DcimPowerOutletTemplatesList200Response) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *DcimPowerOutletTemplatesList200Response) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *DcimPowerOutletTemplatesList200Response) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *DcimPowerOutletTemplatesList200Response) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DcimPowerOutletTemplatesList200Response) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DcimPowerOutletTemplatesList200Response) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *DcimPowerOutletTemplatesList200Response) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *DcimPowerOutletTemplatesList200Response) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *DcimPowerOutletTemplatesList200Response) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *DcimPowerOutletTemplatesList200Response) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *DcimPowerOutletTemplatesList200Response) GetResults() []PowerOutletTemplate {
	if o == nil {
		var ret []PowerOutletTemplate
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *DcimPowerOutletTemplatesList200Response) GetResultsOk() ([]PowerOutletTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *DcimPowerOutletTemplatesList200Response) SetResults(v []PowerOutletTemplate) {
	o.Results = v
}

func (o DcimPowerOutletTemplatesList200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableDcimPowerOutletTemplatesList200Response struct {
	value *DcimPowerOutletTemplatesList200Response
	isSet bool
}

func (v NullableDcimPowerOutletTemplatesList200Response) Get() *DcimPowerOutletTemplatesList200Response {
	return v.value
}

func (v *NullableDcimPowerOutletTemplatesList200Response) Set(val *DcimPowerOutletTemplatesList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimPowerOutletTemplatesList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimPowerOutletTemplatesList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimPowerOutletTemplatesList200Response(val *DcimPowerOutletTemplatesList200Response) *NullableDcimPowerOutletTemplatesList200Response {
	return &NullableDcimPowerOutletTemplatesList200Response{value: val, isSet: true}
}

func (v NullableDcimPowerOutletTemplatesList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimPowerOutletTemplatesList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

