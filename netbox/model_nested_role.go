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

// NestedRole struct for NestedRole
type NestedRole struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	PrefixCount *int32 `json:"prefix_count,omitempty"`
	VlanCount *int32 `json:"vlan_count,omitempty"`
}

// NewNestedRole instantiates a new NestedRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRole(name string, slug string) *NestedRole {
	this := NestedRole{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedRoleWithDefaults instantiates a new NestedRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRoleWithDefaults() *NestedRole {
	this := NestedRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedRole) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedRole) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedRole) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedRole) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedRole) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedRole) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedRole) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedRole) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedRole) SetSlug(v string) {
	o.Slug = v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *NestedRole) GetPrefixCount() int32 {
	if o == nil || o.PrefixCount == nil {
		var ret int32
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetPrefixCountOk() (*int32, bool) {
	if o == nil || o.PrefixCount == nil {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *NestedRole) HasPrefixCount() bool {
	if o != nil && o.PrefixCount != nil {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int32 and assigns it to the PrefixCount field.
func (o *NestedRole) SetPrefixCount(v int32) {
	o.PrefixCount = &v
}

// GetVlanCount returns the VlanCount field value if set, zero value otherwise.
func (o *NestedRole) GetVlanCount() int32 {
	if o == nil || o.VlanCount == nil {
		var ret int32
		return ret
	}
	return *o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetVlanCountOk() (*int32, bool) {
	if o == nil || o.VlanCount == nil {
		return nil, false
	}
	return o.VlanCount, true
}

// HasVlanCount returns a boolean if a field has been set.
func (o *NestedRole) HasVlanCount() bool {
	if o != nil && o.VlanCount != nil {
		return true
	}

	return false
}

// SetVlanCount gets a reference to the given int32 and assigns it to the VlanCount field.
func (o *NestedRole) SetVlanCount(v int32) {
	o.VlanCount = &v
}

func (o NestedRole) MarshalJSON() ([]byte, error) {
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
	if o.PrefixCount != nil {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if o.VlanCount != nil {
		toSerialize["vlan_count"] = o.VlanCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedRole struct {
	value *NestedRole
	isSet bool
}

func (v NullableNestedRole) Get() *NestedRole {
	return v.value
}

func (v *NullableNestedRole) Set(val *NestedRole) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRole) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRole(val *NestedRole) *NullableNestedRole {
	return &NullableNestedRole{value: val, isSet: true}
}

func (v NullableNestedRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


