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

// NestedDeviceType struct for NestedDeviceType
type NestedDeviceType struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty"`
	Model string `json:"model"`
	Slug string `json:"slug"`
	DisplayName *string `json:"display_name,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
}

// NewNestedDeviceType instantiates a new NestedDeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDeviceType(model string, slug string) *NestedDeviceType {
	this := NestedDeviceType{}
	this.Model = model
	this.Slug = slug
	return &this
}

// NewNestedDeviceTypeWithDefaults instantiates a new NestedDeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeviceTypeWithDefaults() *NestedDeviceType {
	this := NestedDeviceType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedDeviceType) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedDeviceType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedDeviceType) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedDeviceType) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedDeviceType) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedDeviceType) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedDeviceType) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedDeviceType) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedDeviceType) SetDisplay(v string) {
	o.Display = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *NestedDeviceType) GetManufacturer() NestedManufacturer {
	if o == nil || o.Manufacturer == nil {
		var ret NestedManufacturer
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetManufacturerOk() (*NestedManufacturer, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *NestedDeviceType) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NestedManufacturer and assigns it to the Manufacturer field.
func (o *NestedDeviceType) SetManufacturer(v NestedManufacturer) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value
func (o *NestedDeviceType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *NestedDeviceType) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *NestedDeviceType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedDeviceType) SetSlug(v string) {
	o.Slug = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NestedDeviceType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NestedDeviceType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NestedDeviceType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *NestedDeviceType) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetDeviceCountOk() (*int32, bool) {
	if o == nil || o.DeviceCount == nil {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *NestedDeviceType) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *NestedDeviceType) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

func (o NestedDeviceType) MarshalJSON() ([]byte, error) {
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
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.DeviceCount != nil {
		toSerialize["device_count"] = o.DeviceCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedDeviceType struct {
	value *NestedDeviceType
	isSet bool
}

func (v NullableNestedDeviceType) Get() *NestedDeviceType {
	return v.value
}

func (v *NullableNestedDeviceType) Set(val *NestedDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDeviceType(val *NestedDeviceType) *NullableNestedDeviceType {
	return &NullableNestedDeviceType{value: val, isSet: true}
}

func (v NullableNestedDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


