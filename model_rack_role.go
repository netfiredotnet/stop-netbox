/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// RackRole struct for RackRole
type RackRole struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	RackCount *int32 `json:"rack_count,omitempty"`
}

// NewRackRole instantiates a new RackRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackRole(name string, slug string) *RackRole {
	this := RackRole{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewRackRoleWithDefaults instantiates a new RackRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackRoleWithDefaults() *RackRole {
	this := RackRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RackRole) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RackRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RackRole) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RackRole) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RackRole) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RackRole) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *RackRole) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *RackRole) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *RackRole) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *RackRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RackRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *RackRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *RackRole) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *RackRole) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *RackRole) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *RackRole) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RackRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RackRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RackRole) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackRole) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackRole) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackRole) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RackRole) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RackRole) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *RackRole) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RackRole) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RackRole) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RackRole) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRackCount returns the RackCount field value if set, zero value otherwise.
func (o *RackRole) GetRackCount() int32 {
	if o == nil || o.RackCount == nil {
		var ret int32
		return ret
	}
	return *o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetRackCountOk() (*int32, bool) {
	if o == nil || o.RackCount == nil {
		return nil, false
	}
	return o.RackCount, true
}

// HasRackCount returns a boolean if a field has been set.
func (o *RackRole) HasRackCount() bool {
	if o != nil && o.RackCount != nil {
		return true
	}

	return false
}

// SetRackCount gets a reference to the given int32 and assigns it to the RackCount field.
func (o *RackRole) SetRackCount(v int32) {
	o.RackCount = &v
}

func (o RackRole) MarshalJSON() ([]byte, error) {
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
	if o.Color != nil {
		toSerialize["color"] = o.Color
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
	if o.RackCount != nil {
		toSerialize["rack_count"] = o.RackCount
	}
	return json.Marshal(toSerialize)
}

type NullableRackRole struct {
	value *RackRole
	isSet bool
}

func (v NullableRackRole) Get() *RackRole {
	return v.value
}

func (v *NullableRackRole) Set(val *RackRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRackRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRackRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackRole(val *RackRole) *NullableRackRole {
	return &NullableRackRole{value: val, isSet: true}
}

func (v NullableRackRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

