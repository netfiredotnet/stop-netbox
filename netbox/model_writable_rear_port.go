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

// WritableRearPort struct for WritableRearPort
type WritableRearPort struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Device int32 `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type string `json:"type"`
	Positions *int32 `json:"positions,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Cable *NestedCable `json:"cable,omitempty"`
	//  Return the appropriate serializer for the cable termination model. 
	CablePeer *map[string]string `json:"cable_peer,omitempty"`
	CablePeerType *string `json:"cable_peer_type,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	Occupied *bool `json:"_occupied,omitempty"`
}

// NewWritableRearPort instantiates a new WritableRearPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableRearPort(device int32, name string, type_ string) *WritableRearPort {
	this := WritableRearPort{}
	this.Device = device
	this.Name = name
	this.Type = type_
	return &this
}

// NewWritableRearPortWithDefaults instantiates a new WritableRearPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableRearPortWithDefaults() *WritableRearPort {
	this := WritableRearPort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritableRearPort) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritableRearPort) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritableRearPort) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritableRearPort) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritableRearPort) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritableRearPort) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritableRearPort) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritableRearPort) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritableRearPort) SetDisplay(v string) {
	o.Display = &v
}

// GetDevice returns the Device field value
func (o *WritableRearPort) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritableRearPort) SetDevice(v int32) {
	o.Device = v
}

// GetName returns the Name field value
func (o *WritableRearPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableRearPort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableRearPort) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableRearPort) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableRearPort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *WritableRearPort) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WritableRearPort) SetType(v string) {
	o.Type = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *WritableRearPort) GetPositions() int32 {
	if o == nil || o.Positions == nil {
		var ret int32
		return ret
	}
	return *o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetPositionsOk() (*int32, bool) {
	if o == nil || o.Positions == nil {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *WritableRearPort) HasPositions() bool {
	if o != nil && o.Positions != nil {
		return true
	}

	return false
}

// SetPositions gets a reference to the given int32 and assigns it to the Positions field.
func (o *WritableRearPort) SetPositions(v int32) {
	o.Positions = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableRearPort) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableRearPort) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableRearPort) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *WritableRearPort) GetMarkConnected() bool {
	if o == nil || o.MarkConnected == nil {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || o.MarkConnected == nil {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *WritableRearPort) HasMarkConnected() bool {
	if o != nil && o.MarkConnected != nil {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *WritableRearPort) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *WritableRearPort) GetCable() NestedCable {
	if o == nil || o.Cable == nil {
		var ret NestedCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetCableOk() (*NestedCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *WritableRearPort) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given NestedCable and assigns it to the Cable field.
func (o *WritableRearPort) SetCable(v NestedCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise.
func (o *WritableRearPort) GetCablePeer() map[string]string {
	if o == nil || o.CablePeer == nil {
		var ret map[string]string
		return ret
	}
	return *o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetCablePeerOk() (*map[string]string, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *WritableRearPort) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]string and assigns it to the CablePeer field.
func (o *WritableRearPort) SetCablePeer(v map[string]string) {
	o.CablePeer = &v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise.
func (o *WritableRearPort) GetCablePeerType() string {
	if o == nil || o.CablePeerType == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil || o.CablePeerType == nil {
		return nil, false
	}
	return o.CablePeerType, true
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *WritableRearPort) HasCablePeerType() bool {
	if o != nil && o.CablePeerType != nil {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given string and assigns it to the CablePeerType field.
func (o *WritableRearPort) SetCablePeerType(v string) {
	o.CablePeerType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableRearPort) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableRearPort) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WritableRearPort) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableRearPort) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableRearPort) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableRearPort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritableRearPort) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritableRearPort) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritableRearPort) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritableRearPort) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritableRearPort) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritableRearPort) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetOccupied returns the Occupied field value if set, zero value otherwise.
func (o *WritableRearPort) GetOccupied() bool {
	if o == nil || o.Occupied == nil {
		var ret bool
		return ret
	}
	return *o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRearPort) GetOccupiedOk() (*bool, bool) {
	if o == nil || o.Occupied == nil {
		return nil, false
	}
	return o.Occupied, true
}

// HasOccupied returns a boolean if a field has been set.
func (o *WritableRearPort) HasOccupied() bool {
	if o != nil && o.Occupied != nil {
		return true
	}

	return false
}

// SetOccupied gets a reference to the given bool and assigns it to the Occupied field.
func (o *WritableRearPort) SetOccupied(v bool) {
	o.Occupied = &v
}

func (o WritableRearPort) MarshalJSON() ([]byte, error) {
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
		toSerialize["device"] = o.Device
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Positions != nil {
		toSerialize["positions"] = o.Positions
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MarkConnected != nil {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if o.Cable != nil {
		toSerialize["cable"] = o.Cable
	}
	if o.CablePeer != nil {
		toSerialize["cable_peer"] = o.CablePeer
	}
	if o.CablePeerType != nil {
		toSerialize["cable_peer_type"] = o.CablePeerType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
	if o.Occupied != nil {
		toSerialize["_occupied"] = o.Occupied
	}
	return json.Marshal(toSerialize)
}

type NullableWritableRearPort struct {
	value *WritableRearPort
	isSet bool
}

func (v NullableWritableRearPort) Get() *WritableRearPort {
	return v.value
}

func (v *NullableWritableRearPort) Set(val *WritableRearPort) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableRearPort) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableRearPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableRearPort(val *WritableRearPort) *NullableWritableRearPort {
	return &NullableWritableRearPort{value: val, isSet: true}
}

func (v NullableWritableRearPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableRearPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

