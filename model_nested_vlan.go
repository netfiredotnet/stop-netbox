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

// NestedVLAN struct for NestedVLAN
type NestedVLAN struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	DisplayName *string `json:"display_name,omitempty"`
}

// NewNestedVLAN instantiates a new NestedVLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVLAN(vid int32, name string) *NestedVLAN {
	this := NestedVLAN{}
	this.Vid = vid
	this.Name = name
	return &this
}

// NewNestedVLANWithDefaults instantiates a new NestedVLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVLANWithDefaults() *NestedVLAN {
	this := NestedVLAN{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedVLAN) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedVLAN) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedVLAN) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedVLAN) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedVLAN) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedVLAN) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedVLAN) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedVLAN) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedVLAN) SetDisplay(v string) {
	o.Display = &v
}

// GetVid returns the Vid field value
func (o *NestedVLAN) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *NestedVLAN) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *NestedVLAN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVLAN) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NestedVLAN) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedVLAN) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NestedVLAN) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NestedVLAN) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o NestedVLAN) MarshalJSON() ([]byte, error) {
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
		toSerialize["vid"] = o.Vid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableNestedVLAN struct {
	value *NestedVLAN
	isSet bool
}

func (v NullableNestedVLAN) Get() *NestedVLAN {
	return v.value
}

func (v *NullableNestedVLAN) Set(val *NestedVLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVLAN(val *NestedVLAN) *NullableNestedVLAN {
	return &NullableNestedVLAN{value: val, isSet: true}
}

func (v NullableNestedVLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

