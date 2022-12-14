/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// VLANGroup struct for VLANGroup
type VLANGroup struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	ScopeType *string `json:"scope_type,omitempty"`
	ScopeId NullableInt32 `json:"scope_id,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	VlanCount *int32 `json:"vlan_count,omitempty"`
}

// NewVLANGroup instantiates a new VLANGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVLANGroup(name string, slug string) *VLANGroup {
	this := VLANGroup{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewVLANGroupWithDefaults instantiates a new VLANGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVLANGroupWithDefaults() *VLANGroup {
	this := VLANGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VLANGroup) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VLANGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VLANGroup) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VLANGroup) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VLANGroup) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VLANGroup) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *VLANGroup) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *VLANGroup) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *VLANGroup) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *VLANGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VLANGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *VLANGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *VLANGroup) SetSlug(v string) {
	o.Slug = v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *VLANGroup) GetScopeType() string {
	if o == nil || o.ScopeType == nil {
		var ret string
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetScopeTypeOk() (*string, bool) {
	if o == nil || o.ScopeType == nil {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *VLANGroup) HasScopeType() bool {
	if o != nil && o.ScopeType != nil {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given string and assigns it to the ScopeType field.
func (o *VLANGroup) SetScopeType(v string) {
	o.ScopeType = &v
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLANGroup) GetScopeId() int32 {
	if o == nil || o.ScopeId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ScopeId.Get()
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroup) GetScopeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeId.Get(), o.ScopeId.IsSet()
}

// HasScopeId returns a boolean if a field has been set.
func (o *VLANGroup) HasScopeId() bool {
	if o != nil && o.ScopeId.IsSet() {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given NullableInt32 and assigns it to the ScopeId field.
func (o *VLANGroup) SetScopeId(v int32) {
	o.ScopeId.Set(&v)
}
// SetScopeIdNil sets the value for ScopeId to be an explicit nil
func (o *VLANGroup) SetScopeIdNil() {
	o.ScopeId.Set(nil)
}

// UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
func (o *VLANGroup) UnsetScopeId() {
	o.ScopeId.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *VLANGroup) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *VLANGroup) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *VLANGroup) SetScope(v string) {
	o.Scope = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VLANGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VLANGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VLANGroup) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VLANGroup) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VLANGroup) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VLANGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *VLANGroup) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *VLANGroup) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *VLANGroup) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *VLANGroup) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VLANGroup) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *VLANGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetVlanCount returns the VlanCount field value if set, zero value otherwise.
func (o *VLANGroup) GetVlanCount() int32 {
	if o == nil || o.VlanCount == nil {
		var ret int32
		return ret
	}
	return *o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetVlanCountOk() (*int32, bool) {
	if o == nil || o.VlanCount == nil {
		return nil, false
	}
	return o.VlanCount, true
}

// HasVlanCount returns a boolean if a field has been set.
func (o *VLANGroup) HasVlanCount() bool {
	if o != nil && o.VlanCount != nil {
		return true
	}

	return false
}

// SetVlanCount gets a reference to the given int32 and assigns it to the VlanCount field.
func (o *VLANGroup) SetVlanCount(v int32) {
	o.VlanCount = &v
}

func (o VLANGroup) MarshalJSON() ([]byte, error) {
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
	if o.ScopeType != nil {
		toSerialize["scope_type"] = o.ScopeType
	}
	if o.ScopeId.IsSet() {
		toSerialize["scope_id"] = o.ScopeId.Get()
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.VlanCount != nil {
		toSerialize["vlan_count"] = o.VlanCount
	}
	return json.Marshal(toSerialize)
}

type NullableVLANGroup struct {
	value *VLANGroup
	isSet bool
}

func (v NullableVLANGroup) Get() *VLANGroup {
	return v.value
}

func (v *NullableVLANGroup) Set(val *VLANGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVLANGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVLANGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLANGroup(val *VLANGroup) *NullableVLANGroup {
	return &NullableVLANGroup{value: val, isSet: true}
}

func (v NullableVLANGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLANGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


