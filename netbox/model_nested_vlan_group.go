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

// NestedVLANGroup struct for NestedVLANGroup
type NestedVLANGroup struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	VlanCount *int32 `json:"vlan_count,omitempty"`
}

// NewNestedVLANGroup instantiates a new NestedVLANGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVLANGroup(name string, slug string) *NestedVLANGroup {
	this := NestedVLANGroup{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedVLANGroupWithDefaults instantiates a new NestedVLANGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVLANGroupWithDefaults() *NestedVLANGroup {
	this := NestedVLANGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedVLANGroup) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedVLANGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedVLANGroup) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedVLANGroup) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedVLANGroup) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedVLANGroup) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedVLANGroup) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedVLANGroup) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedVLANGroup) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedVLANGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVLANGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedVLANGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedVLANGroup) SetSlug(v string) {
	o.Slug = v
}

// GetVlanCount returns the VlanCount field value if set, zero value otherwise.
func (o *NestedVLANGroup) GetVlanCount() int32 {
	if o == nil || o.VlanCount == nil {
		var ret int32
		return ret
	}
	return *o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLANGroup) GetVlanCountOk() (*int32, bool) {
	if o == nil || o.VlanCount == nil {
		return nil, false
	}
	return o.VlanCount, true
}

// HasVlanCount returns a boolean if a field has been set.
func (o *NestedVLANGroup) HasVlanCount() bool {
	if o != nil && o.VlanCount != nil {
		return true
	}

	return false
}

// SetVlanCount gets a reference to the given int32 and assigns it to the VlanCount field.
func (o *NestedVLANGroup) SetVlanCount(v int32) {
	o.VlanCount = &v
}

func (o NestedVLANGroup) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.VlanCount != nil {
		toSerialize["vlan_count"] = o.VlanCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedVLANGroup struct {
	value *NestedVLANGroup
	isSet bool
}

func (v NullableNestedVLANGroup) Get() *NestedVLANGroup {
	return v.value
}

func (v *NullableNestedVLANGroup) Set(val *NestedVLANGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVLANGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVLANGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVLANGroup(val *NestedVLANGroup) *NullableNestedVLANGroup {
	return &NullableNestedVLANGroup{value: val, isSet: true}
}

func (v NullableNestedVLANGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVLANGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


