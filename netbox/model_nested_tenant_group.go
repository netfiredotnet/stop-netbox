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

// NestedTenantGroup struct for NestedTenantGroup
type NestedTenantGroup struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	TenantCount *int32 `json:"tenant_count,omitempty"`
	Depth *int32 `json:"_depth,omitempty"`
}

// NewNestedTenantGroup instantiates a new NestedTenantGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTenantGroup(name string, slug string) *NestedTenantGroup {
	this := NestedTenantGroup{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedTenantGroupWithDefaults instantiates a new NestedTenantGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTenantGroupWithDefaults() *NestedTenantGroup {
	this := NestedTenantGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedTenantGroup) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedTenantGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedTenantGroup) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedTenantGroup) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedTenantGroup) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedTenantGroup) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedTenantGroup) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedTenantGroup) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedTenantGroup) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedTenantGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTenantGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedTenantGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedTenantGroup) SetSlug(v string) {
	o.Slug = v
}

// GetTenantCount returns the TenantCount field value if set, zero value otherwise.
func (o *NestedTenantGroup) GetTenantCount() int32 {
	if o == nil || o.TenantCount == nil {
		var ret int32
		return ret
	}
	return *o.TenantCount
}

// GetTenantCountOk returns a tuple with the TenantCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetTenantCountOk() (*int32, bool) {
	if o == nil || o.TenantCount == nil {
		return nil, false
	}
	return o.TenantCount, true
}

// HasTenantCount returns a boolean if a field has been set.
func (o *NestedTenantGroup) HasTenantCount() bool {
	if o != nil && o.TenantCount != nil {
		return true
	}

	return false
}

// SetTenantCount gets a reference to the given int32 and assigns it to the TenantCount field.
func (o *NestedTenantGroup) SetTenantCount(v int32) {
	o.TenantCount = &v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *NestedTenantGroup) GetDepth() int32 {
	if o == nil || o.Depth == nil {
		var ret int32
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDepthOk() (*int32, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *NestedTenantGroup) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int32 and assigns it to the Depth field.
func (o *NestedTenantGroup) SetDepth(v int32) {
	o.Depth = &v
}

func (o NestedTenantGroup) MarshalJSON() ([]byte, error) {
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
	if o.TenantCount != nil {
		toSerialize["tenant_count"] = o.TenantCount
	}
	if o.Depth != nil {
		toSerialize["_depth"] = o.Depth
	}
	return json.Marshal(toSerialize)
}

type NullableNestedTenantGroup struct {
	value *NestedTenantGroup
	isSet bool
}

func (v NullableNestedTenantGroup) Get() *NestedTenantGroup {
	return v.value
}

func (v *NullableNestedTenantGroup) Set(val *NestedTenantGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTenantGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTenantGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTenantGroup(val *NestedTenantGroup) *NullableNestedTenantGroup {
	return &NullableNestedTenantGroup{value: val, isSet: true}
}

func (v NullableNestedTenantGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTenantGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

