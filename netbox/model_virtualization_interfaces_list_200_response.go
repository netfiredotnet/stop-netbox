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

// VirtualizationInterfacesList200Response struct for VirtualizationInterfacesList200Response
type VirtualizationInterfacesList200Response struct {
	Count int32 `json:"count"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []VMInterface `json:"results"`
}

// NewVirtualizationInterfacesList200Response instantiates a new VirtualizationInterfacesList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationInterfacesList200Response(count int32, results []VMInterface) *VirtualizationInterfacesList200Response {
	this := VirtualizationInterfacesList200Response{}
	this.Count = count
	this.Results = results
	return &this
}

// NewVirtualizationInterfacesList200ResponseWithDefaults instantiates a new VirtualizationInterfacesList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationInterfacesList200ResponseWithDefaults() *VirtualizationInterfacesList200Response {
	this := VirtualizationInterfacesList200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *VirtualizationInterfacesList200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *VirtualizationInterfacesList200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *VirtualizationInterfacesList200Response) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationInterfacesList200Response) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationInterfacesList200Response) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *VirtualizationInterfacesList200Response) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *VirtualizationInterfacesList200Response) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *VirtualizationInterfacesList200Response) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *VirtualizationInterfacesList200Response) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationInterfacesList200Response) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationInterfacesList200Response) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *VirtualizationInterfacesList200Response) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *VirtualizationInterfacesList200Response) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *VirtualizationInterfacesList200Response) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *VirtualizationInterfacesList200Response) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *VirtualizationInterfacesList200Response) GetResults() []VMInterface {
	if o == nil {
		var ret []VMInterface
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *VirtualizationInterfacesList200Response) GetResultsOk() ([]VMInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *VirtualizationInterfacesList200Response) SetResults(v []VMInterface) {
	o.Results = v
}

func (o VirtualizationInterfacesList200Response) MarshalJSON() ([]byte, error) {
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

type NullableVirtualizationInterfacesList200Response struct {
	value *VirtualizationInterfacesList200Response
	isSet bool
}

func (v NullableVirtualizationInterfacesList200Response) Get() *VirtualizationInterfacesList200Response {
	return v.value
}

func (v *NullableVirtualizationInterfacesList200Response) Set(val *VirtualizationInterfacesList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationInterfacesList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationInterfacesList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationInterfacesList200Response(val *VirtualizationInterfacesList200Response) *NullableVirtualizationInterfacesList200Response {
	return &NullableVirtualizationInterfacesList200Response{value: val, isSet: true}
}

func (v NullableVirtualizationInterfacesList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationInterfacesList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


