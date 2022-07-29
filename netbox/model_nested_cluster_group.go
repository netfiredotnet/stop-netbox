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

// NestedClusterGroup struct for NestedClusterGroup
type NestedClusterGroup struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	ClusterCount *int32 `json:"cluster_count,omitempty"`
}

// NewNestedClusterGroup instantiates a new NestedClusterGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedClusterGroup(name string, slug string) *NestedClusterGroup {
	this := NestedClusterGroup{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedClusterGroupWithDefaults instantiates a new NestedClusterGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedClusterGroupWithDefaults() *NestedClusterGroup {
	this := NestedClusterGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedClusterGroup) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedClusterGroup) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedClusterGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedClusterGroup) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedClusterGroup) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedClusterGroup) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedClusterGroup) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedClusterGroup) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedClusterGroup) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedClusterGroup) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedClusterGroup) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedClusterGroup) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedClusterGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedClusterGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedClusterGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedClusterGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedClusterGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedClusterGroup) SetSlug(v string) {
	o.Slug = v
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *NestedClusterGroup) GetClusterCount() int32 {
	if o == nil || o.ClusterCount == nil {
		var ret int32
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedClusterGroup) GetClusterCountOk() (*int32, bool) {
	if o == nil || o.ClusterCount == nil {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *NestedClusterGroup) HasClusterCount() bool {
	if o != nil && o.ClusterCount != nil {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int32 and assigns it to the ClusterCount field.
func (o *NestedClusterGroup) SetClusterCount(v int32) {
	o.ClusterCount = &v
}

func (o NestedClusterGroup) MarshalJSON() ([]byte, error) {
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
	if o.ClusterCount != nil {
		toSerialize["cluster_count"] = o.ClusterCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedClusterGroup struct {
	value *NestedClusterGroup
	isSet bool
}

func (v NullableNestedClusterGroup) Get() *NestedClusterGroup {
	return v.value
}

func (v *NullableNestedClusterGroup) Set(val *NestedClusterGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedClusterGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedClusterGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedClusterGroup(val *NestedClusterGroup) *NullableNestedClusterGroup {
	return &NullableNestedClusterGroup{value: val, isSet: true}
}

func (v NullableNestedClusterGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedClusterGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

