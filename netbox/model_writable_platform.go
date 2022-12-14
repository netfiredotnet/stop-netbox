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

// WritablePlatform struct for WritablePlatform
type WritablePlatform struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	// Optionally limit this platform to devices of a certain manufacturer
	Manufacturer NullableInt32 `json:"manufacturer,omitempty"`
	// The name of the NAPALM driver to use when interacting with devices
	NapalmDriver *string `json:"napalm_driver,omitempty"`
	// Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs NullableString `json:"napalm_args,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	VirtualmachineCount *int32 `json:"virtualmachine_count,omitempty"`
}

// NewWritablePlatform instantiates a new WritablePlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritablePlatform(name string, slug string) *WritablePlatform {
	this := WritablePlatform{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewWritablePlatformWithDefaults instantiates a new WritablePlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritablePlatformWithDefaults() *WritablePlatform {
	this := WritablePlatform{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritablePlatform) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritablePlatform) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritablePlatform) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritablePlatform) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritablePlatform) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritablePlatform) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritablePlatform) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritablePlatform) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritablePlatform) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *WritablePlatform) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritablePlatform) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritablePlatform) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritablePlatform) SetSlug(v string) {
	o.Slug = v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePlatform) GetManufacturer() int32 {
	if o == nil || o.Manufacturer.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePlatform) GetManufacturerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *WritablePlatform) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableInt32 and assigns it to the Manufacturer field.
func (o *WritablePlatform) SetManufacturer(v int32) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *WritablePlatform) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *WritablePlatform) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetNapalmDriver returns the NapalmDriver field value if set, zero value otherwise.
func (o *WritablePlatform) GetNapalmDriver() string {
	if o == nil || o.NapalmDriver == nil {
		var ret string
		return ret
	}
	return *o.NapalmDriver
}

// GetNapalmDriverOk returns a tuple with the NapalmDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetNapalmDriverOk() (*string, bool) {
	if o == nil || o.NapalmDriver == nil {
		return nil, false
	}
	return o.NapalmDriver, true
}

// HasNapalmDriver returns a boolean if a field has been set.
func (o *WritablePlatform) HasNapalmDriver() bool {
	if o != nil && o.NapalmDriver != nil {
		return true
	}

	return false
}

// SetNapalmDriver gets a reference to the given string and assigns it to the NapalmDriver field.
func (o *WritablePlatform) SetNapalmDriver(v string) {
	o.NapalmDriver = &v
}

// GetNapalmArgs returns the NapalmArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePlatform) GetNapalmArgs() string {
	if o == nil || o.NapalmArgs.Get() == nil {
		var ret string
		return ret
	}
	return *o.NapalmArgs.Get()
}

// GetNapalmArgsOk returns a tuple with the NapalmArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePlatform) GetNapalmArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NapalmArgs.Get(), o.NapalmArgs.IsSet()
}

// HasNapalmArgs returns a boolean if a field has been set.
func (o *WritablePlatform) HasNapalmArgs() bool {
	if o != nil && o.NapalmArgs.IsSet() {
		return true
	}

	return false
}

// SetNapalmArgs gets a reference to the given NullableString and assigns it to the NapalmArgs field.
func (o *WritablePlatform) SetNapalmArgs(v string) {
	o.NapalmArgs.Set(&v)
}
// SetNapalmArgsNil sets the value for NapalmArgs to be an explicit nil
func (o *WritablePlatform) SetNapalmArgsNil() {
	o.NapalmArgs.Set(nil)
}

// UnsetNapalmArgs ensures that no value is present for NapalmArgs, not even an explicit nil
func (o *WritablePlatform) UnsetNapalmArgs() {
	o.NapalmArgs.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritablePlatform) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritablePlatform) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritablePlatform) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritablePlatform) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritablePlatform) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritablePlatform) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritablePlatform) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritablePlatform) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritablePlatform) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritablePlatform) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritablePlatform) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritablePlatform) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *WritablePlatform) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetDeviceCountOk() (*int32, bool) {
	if o == nil || o.DeviceCount == nil {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *WritablePlatform) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *WritablePlatform) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *WritablePlatform) GetVirtualmachineCount() int32 {
	if o == nil || o.VirtualmachineCount == nil {
		var ret int32
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePlatform) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil || o.VirtualmachineCount == nil {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *WritablePlatform) HasVirtualmachineCount() bool {
	if o != nil && o.VirtualmachineCount != nil {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int32 and assigns it to the VirtualmachineCount field.
func (o *WritablePlatform) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = &v
}

func (o WritablePlatform) MarshalJSON() ([]byte, error) {
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
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.NapalmDriver != nil {
		toSerialize["napalm_driver"] = o.NapalmDriver
	}
	if o.NapalmArgs.IsSet() {
		toSerialize["napalm_args"] = o.NapalmArgs.Get()
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
	if o.DeviceCount != nil {
		toSerialize["device_count"] = o.DeviceCount
	}
	if o.VirtualmachineCount != nil {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	return json.Marshal(toSerialize)
}

type NullableWritablePlatform struct {
	value *WritablePlatform
	isSet bool
}

func (v NullableWritablePlatform) Get() *WritablePlatform {
	return v.value
}

func (v *NullableWritablePlatform) Set(val *WritablePlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePlatform(val *WritablePlatform) *NullableWritablePlatform {
	return &NullableWritablePlatform{value: val, isSet: true}
}

func (v NullableWritablePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


