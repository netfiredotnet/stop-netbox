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

// WritablePowerPortTemplate struct for WritablePowerPortTemplate
type WritablePowerPortTemplate struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	DeviceType int32 `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *string `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw NullableInt32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw NullableInt32 `json:"allocated_draw,omitempty"`
	Description *string `json:"description,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewWritablePowerPortTemplate instantiates a new WritablePowerPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritablePowerPortTemplate(deviceType int32, name string) *WritablePowerPortTemplate {
	this := WritablePowerPortTemplate{}
	this.DeviceType = deviceType
	this.Name = name
	return &this
}

// NewWritablePowerPortTemplateWithDefaults instantiates a new WritablePowerPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritablePowerPortTemplateWithDefaults() *WritablePowerPortTemplate {
	this := WritablePowerPortTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritablePowerPortTemplate) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritablePowerPortTemplate) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritablePowerPortTemplate) SetDisplay(v string) {
	o.Display = &v
}

// GetDeviceType returns the DeviceType field value
func (o *WritablePowerPortTemplate) GetDeviceType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetDeviceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *WritablePowerPortTemplate) SetDeviceType(v int32) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *WritablePowerPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritablePowerPortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritablePowerPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WritablePowerPortTemplate) SetType(v string) {
	o.Type = &v
}

// GetMaximumDraw returns the MaximumDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePowerPortTemplate) GetMaximumDraw() int32 {
	if o == nil || o.MaximumDraw.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaximumDraw.Get()
}

// GetMaximumDrawOk returns a tuple with the MaximumDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerPortTemplate) GetMaximumDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumDraw.Get(), o.MaximumDraw.IsSet()
}

// HasMaximumDraw returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasMaximumDraw() bool {
	if o != nil && o.MaximumDraw.IsSet() {
		return true
	}

	return false
}

// SetMaximumDraw gets a reference to the given NullableInt32 and assigns it to the MaximumDraw field.
func (o *WritablePowerPortTemplate) SetMaximumDraw(v int32) {
	o.MaximumDraw.Set(&v)
}
// SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil
func (o *WritablePowerPortTemplate) SetMaximumDrawNil() {
	o.MaximumDraw.Set(nil)
}

// UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
func (o *WritablePowerPortTemplate) UnsetMaximumDraw() {
	o.MaximumDraw.Unset()
}

// GetAllocatedDraw returns the AllocatedDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePowerPortTemplate) GetAllocatedDraw() int32 {
	if o == nil || o.AllocatedDraw.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AllocatedDraw.Get()
}

// GetAllocatedDrawOk returns a tuple with the AllocatedDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerPortTemplate) GetAllocatedDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocatedDraw.Get(), o.AllocatedDraw.IsSet()
}

// HasAllocatedDraw returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasAllocatedDraw() bool {
	if o != nil && o.AllocatedDraw.IsSet() {
		return true
	}

	return false
}

// SetAllocatedDraw gets a reference to the given NullableInt32 and assigns it to the AllocatedDraw field.
func (o *WritablePowerPortTemplate) SetAllocatedDraw(v int32) {
	o.AllocatedDraw.Set(&v)
}
// SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil
func (o *WritablePowerPortTemplate) SetAllocatedDrawNil() {
	o.AllocatedDraw.Set(nil)
}

// UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
func (o *WritablePowerPortTemplate) UnsetAllocatedDraw() {
	o.AllocatedDraw.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritablePowerPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritablePowerPortTemplate) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritablePowerPortTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerPortTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritablePowerPortTemplate) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritablePowerPortTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o WritablePowerPortTemplate) MarshalJSON() ([]byte, error) {
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
		toSerialize["device_type"] = o.DeviceType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.MaximumDraw.IsSet() {
		toSerialize["maximum_draw"] = o.MaximumDraw.Get()
	}
	if o.AllocatedDraw.IsSet() {
		toSerialize["allocated_draw"] = o.AllocatedDraw.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableWritablePowerPortTemplate struct {
	value *WritablePowerPortTemplate
	isSet bool
}

func (v NullableWritablePowerPortTemplate) Get() *WritablePowerPortTemplate {
	return v.value
}

func (v *NullableWritablePowerPortTemplate) Set(val *WritablePowerPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePowerPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePowerPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePowerPortTemplate(val *WritablePowerPortTemplate) *NullableWritablePowerPortTemplate {
	return &NullableWritablePowerPortTemplate{value: val, isSet: true}
}

func (v NullableWritablePowerPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePowerPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

