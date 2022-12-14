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

// NestedCluster struct for NestedCluster
type NestedCluster struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	VirtualmachineCount *int32 `json:"virtualmachine_count,omitempty"`
}

// NewNestedCluster instantiates a new NestedCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCluster(name string) *NestedCluster {
	this := NestedCluster{}
	this.Name = name
	return &this
}

// NewNestedClusterWithDefaults instantiates a new NestedCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedClusterWithDefaults() *NestedCluster {
	this := NestedCluster{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedCluster) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCluster) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedCluster) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedCluster) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedCluster) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCluster) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedCluster) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedCluster) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedCluster) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCluster) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedCluster) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedCluster) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedCluster) SetName(v string) {
	o.Name = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *NestedCluster) GetVirtualmachineCount() int32 {
	if o == nil || o.VirtualmachineCount == nil {
		var ret int32
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCluster) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil || o.VirtualmachineCount == nil {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *NestedCluster) HasVirtualmachineCount() bool {
	if o != nil && o.VirtualmachineCount != nil {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int32 and assigns it to the VirtualmachineCount field.
func (o *NestedCluster) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = &v
}

func (o NestedCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.VirtualmachineCount != nil {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedCluster struct {
	value *NestedCluster
	isSet bool
}

func (v NullableNestedCluster) Get() *NestedCluster {
	return v.value
}

func (v *NullableNestedCluster) Set(val *NestedCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCluster(val *NestedCluster) *NullableNestedCluster {
	return &NullableNestedCluster{value: val, isSet: true}
}

func (v NullableNestedCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


