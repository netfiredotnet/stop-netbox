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

// NestedCircuitType struct for NestedCircuitType
type NestedCircuitType struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	CircuitCount *int32 `json:"circuit_count,omitempty"`
}

// NewNestedCircuitType instantiates a new NestedCircuitType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCircuitType(name string, slug string) *NestedCircuitType {
	this := NestedCircuitType{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedCircuitTypeWithDefaults instantiates a new NestedCircuitType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCircuitTypeWithDefaults() *NestedCircuitType {
	this := NestedCircuitType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedCircuitType) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedCircuitType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedCircuitType) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedCircuitType) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedCircuitType) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedCircuitType) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedCircuitType) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedCircuitType) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedCircuitType) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedCircuitType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedCircuitType) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedCircuitType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedCircuitType) SetSlug(v string) {
	o.Slug = v
}

// GetCircuitCount returns the CircuitCount field value if set, zero value otherwise.
func (o *NestedCircuitType) GetCircuitCount() int32 {
	if o == nil || o.CircuitCount == nil {
		var ret int32
		return ret
	}
	return *o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCircuitType) GetCircuitCountOk() (*int32, bool) {
	if o == nil || o.CircuitCount == nil {
		return nil, false
	}
	return o.CircuitCount, true
}

// HasCircuitCount returns a boolean if a field has been set.
func (o *NestedCircuitType) HasCircuitCount() bool {
	if o != nil && o.CircuitCount != nil {
		return true
	}

	return false
}

// SetCircuitCount gets a reference to the given int32 and assigns it to the CircuitCount field.
func (o *NestedCircuitType) SetCircuitCount(v int32) {
	o.CircuitCount = &v
}

func (o NestedCircuitType) MarshalJSON() ([]byte, error) {
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
	if o.CircuitCount != nil {
		toSerialize["circuit_count"] = o.CircuitCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedCircuitType struct {
	value *NestedCircuitType
	isSet bool
}

func (v NullableNestedCircuitType) Get() *NestedCircuitType {
	return v.value
}

func (v *NullableNestedCircuitType) Set(val *NestedCircuitType) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCircuitType) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCircuitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCircuitType(val *NestedCircuitType) *NullableNestedCircuitType {
	return &NullableNestedCircuitType{value: val, isSet: true}
}

func (v NullableNestedCircuitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCircuitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

