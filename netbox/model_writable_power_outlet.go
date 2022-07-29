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

// WritablePowerOutlet struct for WritablePowerOutlet
type WritablePowerOutlet struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Device int32 `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Physical port type
	Type *string `json:"type,omitempty"`
	PowerPort NullableInt32 `json:"power_port,omitempty"`
	// Phase (for three-phase feeds)
	FeedLeg *string `json:"feed_leg,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Cable *NestedCable `json:"cable,omitempty"`
	//  Return the appropriate serializer for the cable termination model. 
	CablePeer *map[string]string `json:"cable_peer,omitempty"`
	CablePeerType *string `json:"cable_peer_type,omitempty"`
	//  Return the appropriate serializer for the type of connected object. 
	ConnectedEndpoint *map[string]string `json:"connected_endpoint,omitempty"`
	ConnectedEndpointType *string `json:"connected_endpoint_type,omitempty"`
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	Occupied *bool `json:"_occupied,omitempty"`
}

// NewWritablePowerOutlet instantiates a new WritablePowerOutlet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritablePowerOutlet(device int32, name string) *WritablePowerOutlet {
	this := WritablePowerOutlet{}
	this.Device = device
	this.Name = name
	return &this
}

// NewWritablePowerOutletWithDefaults instantiates a new WritablePowerOutlet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritablePowerOutletWithDefaults() *WritablePowerOutlet {
	this := WritablePowerOutlet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritablePowerOutlet) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritablePowerOutlet) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritablePowerOutlet) SetDisplay(v string) {
	o.Display = &v
}

// GetDevice returns the Device field value
func (o *WritablePowerOutlet) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritablePowerOutlet) SetDevice(v int32) {
	o.Device = v
}

// GetName returns the Name field value
func (o *WritablePowerOutlet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritablePowerOutlet) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritablePowerOutlet) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WritablePowerOutlet) SetType(v string) {
	o.Type = &v
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePowerOutlet) GetPowerPort() int32 {
	if o == nil || o.PowerPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerOutlet) GetPowerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableInt32 and assigns it to the PowerPort field.
func (o *WritablePowerOutlet) SetPowerPort(v int32) {
	o.PowerPort.Set(&v)
}
// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *WritablePowerOutlet) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *WritablePowerOutlet) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetFeedLeg() string {
	if o == nil || o.FeedLeg == nil {
		var ret string
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetFeedLegOk() (*string, bool) {
	if o == nil || o.FeedLeg == nil {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasFeedLeg() bool {
	if o != nil && o.FeedLeg != nil {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given string and assigns it to the FeedLeg field.
func (o *WritablePowerOutlet) SetFeedLeg(v string) {
	o.FeedLeg = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritablePowerOutlet) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetMarkConnected() bool {
	if o == nil || o.MarkConnected == nil {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || o.MarkConnected == nil {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasMarkConnected() bool {
	if o != nil && o.MarkConnected != nil {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *WritablePowerOutlet) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetCable() NestedCable {
	if o == nil || o.Cable == nil {
		var ret NestedCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetCableOk() (*NestedCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given NestedCable and assigns it to the Cable field.
func (o *WritablePowerOutlet) SetCable(v NestedCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetCablePeer() map[string]string {
	if o == nil || o.CablePeer == nil {
		var ret map[string]string
		return ret
	}
	return *o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetCablePeerOk() (*map[string]string, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]string and assigns it to the CablePeer field.
func (o *WritablePowerOutlet) SetCablePeer(v map[string]string) {
	o.CablePeer = &v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetCablePeerType() string {
	if o == nil || o.CablePeerType == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetCablePeerTypeOk() (*string, bool) {
	if o == nil || o.CablePeerType == nil {
		return nil, false
	}
	return o.CablePeerType, true
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasCablePeerType() bool {
	if o != nil && o.CablePeerType != nil {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given string and assigns it to the CablePeerType field.
func (o *WritablePowerOutlet) SetCablePeerType(v string) {
	o.CablePeerType = &v
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetConnectedEndpoint() map[string]string {
	if o == nil || o.ConnectedEndpoint == nil {
		var ret map[string]string
		return ret
	}
	return *o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetConnectedEndpointOk() (*map[string]string, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// HasConnectedEndpoint returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasConnectedEndpoint() bool {
	if o != nil && o.ConnectedEndpoint != nil {
		return true
	}

	return false
}

// SetConnectedEndpoint gets a reference to the given map[string]string and assigns it to the ConnectedEndpoint field.
func (o *WritablePowerOutlet) SetConnectedEndpoint(v map[string]string) {
	o.ConnectedEndpoint = &v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType == nil {
		var ret string
		return ret
	}
	return *o.ConnectedEndpointType
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil || o.ConnectedEndpointType == nil {
		return nil, false
	}
	return o.ConnectedEndpointType, true
}

// HasConnectedEndpointType returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasConnectedEndpointType() bool {
	if o != nil && o.ConnectedEndpointType != nil {
		return true
	}

	return false
}

// SetConnectedEndpointType gets a reference to the given string and assigns it to the ConnectedEndpointType field.
func (o *WritablePowerOutlet) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType = &v
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable == nil {
		var ret bool
		return ret
	}
	return *o.ConnectedEndpointReachable
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil || o.ConnectedEndpointReachable == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable, true
}

// HasConnectedEndpointReachable returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasConnectedEndpointReachable() bool {
	if o != nil && o.ConnectedEndpointReachable != nil {
		return true
	}

	return false
}

// SetConnectedEndpointReachable gets a reference to the given bool and assigns it to the ConnectedEndpointReachable field.
func (o *WritablePowerOutlet) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WritablePowerOutlet) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritablePowerOutlet) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritablePowerOutlet) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritablePowerOutlet) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetOccupied returns the Occupied field value if set, zero value otherwise.
func (o *WritablePowerOutlet) GetOccupied() bool {
	if o == nil || o.Occupied == nil {
		var ret bool
		return ret
	}
	return *o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerOutlet) GetOccupiedOk() (*bool, bool) {
	if o == nil || o.Occupied == nil {
		return nil, false
	}
	return o.Occupied, true
}

// HasOccupied returns a boolean if a field has been set.
func (o *WritablePowerOutlet) HasOccupied() bool {
	if o != nil && o.Occupied != nil {
		return true
	}

	return false
}

// SetOccupied gets a reference to the given bool and assigns it to the Occupied field.
func (o *WritablePowerOutlet) SetOccupied(v bool) {
	o.Occupied = &v
}

func (o WritablePowerOutlet) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if o.FeedLeg != nil {
		toSerialize["feed_leg"] = o.FeedLeg
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
	if o.ConnectedEndpoint != nil {
		toSerialize["connected_endpoint"] = o.ConnectedEndpoint
	}
	if o.ConnectedEndpointType != nil {
		toSerialize["connected_endpoint_type"] = o.ConnectedEndpointType
	}
	if o.ConnectedEndpointReachable != nil {
		toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable
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

type NullableWritablePowerOutlet struct {
	value *WritablePowerOutlet
	isSet bool
}

func (v NullableWritablePowerOutlet) Get() *WritablePowerOutlet {
	return v.value
}

func (v *NullableWritablePowerOutlet) Set(val *WritablePowerOutlet) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePowerOutlet) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePowerOutlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePowerOutlet(val *WritablePowerOutlet) *NullableWritablePowerOutlet {
	return &NullableWritablePowerOutlet{value: val, isSet: true}
}

func (v NullableWritablePowerOutlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePowerOutlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

