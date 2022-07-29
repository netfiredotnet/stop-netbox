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

// Interface struct for Interface
type Interface struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Device NestedDevice `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type Type2 `json:"type"`
	Enabled *bool `json:"enabled,omitempty"`
	Parent *NestedInterface `json:"parent,omitempty"`
	Lag *NestedInterface `json:"lag,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	MacAddress NullableString `json:"mac_address,omitempty"`
	// This interface is used only for out-of-band management
	MgmtOnly *bool `json:"mgmt_only,omitempty"`
	Description *string `json:"description,omitempty"`
	Mode *Mode `json:"mode,omitempty"`
	UntaggedVlan NullableNestedVLAN `json:"untagged_vlan,omitempty"`
	TaggedVlans []NestedVLAN `json:"tagged_vlans,omitempty"`
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
	CountIpaddresses *int32 `json:"count_ipaddresses,omitempty"`
	Occupied *bool `json:"_occupied,omitempty"`
}

// NewInterface instantiates a new Interface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterface(device NestedDevice, name string, type_ Type2) *Interface {
	this := Interface{}
	this.Device = device
	this.Name = name
	this.Type = type_
	return &this
}

// NewInterfaceWithDefaults instantiates a new Interface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceWithDefaults() *Interface {
	this := Interface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Interface) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Interface) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Interface) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Interface) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Interface) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Interface) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *Interface) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *Interface) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *Interface) SetDisplay(v string) {
	o.Display = &v
}

// GetDevice returns the Device field value
func (o *Interface) GetDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *Interface) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *Interface) SetDevice(v NestedDevice) {
	o.Device = v
}

// GetName returns the Name field value
func (o *Interface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Interface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Interface) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Interface) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Interface) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Interface) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *Interface) GetType() Type2 {
	if o == nil {
		var ret Type2
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Interface) GetTypeOk() (*Type2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Interface) SetType(v Type2) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Interface) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Interface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Interface) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Interface) GetParent() NestedInterface {
	if o == nil || o.Parent == nil {
		var ret NestedInterface
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetParentOk() (*NestedInterface, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Interface) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given NestedInterface and assigns it to the Parent field.
func (o *Interface) SetParent(v NestedInterface) {
	o.Parent = &v
}

// GetLag returns the Lag field value if set, zero value otherwise.
func (o *Interface) GetLag() NestedInterface {
	if o == nil || o.Lag == nil {
		var ret NestedInterface
		return ret
	}
	return *o.Lag
}

// GetLagOk returns a tuple with the Lag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetLagOk() (*NestedInterface, bool) {
	if o == nil || o.Lag == nil {
		return nil, false
	}
	return o.Lag, true
}

// HasLag returns a boolean if a field has been set.
func (o *Interface) HasLag() bool {
	if o != nil && o.Lag != nil {
		return true
	}

	return false
}

// SetLag gets a reference to the given NestedInterface and assigns it to the Lag field.
func (o *Interface) SetLag(v NestedInterface) {
	o.Lag = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interface) GetMtu() int32 {
	if o == nil || o.Mtu.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interface) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *Interface) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *Interface) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *Interface) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *Interface) UnsetMtu() {
	o.Mtu.Unset()
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interface) GetMacAddress() string {
	if o == nil || o.MacAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.MacAddress.Get()
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interface) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddress.Get(), o.MacAddress.IsSet()
}

// HasMacAddress returns a boolean if a field has been set.
func (o *Interface) HasMacAddress() bool {
	if o != nil && o.MacAddress.IsSet() {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given NullableString and assigns it to the MacAddress field.
func (o *Interface) SetMacAddress(v string) {
	o.MacAddress.Set(&v)
}
// SetMacAddressNil sets the value for MacAddress to be an explicit nil
func (o *Interface) SetMacAddressNil() {
	o.MacAddress.Set(nil)
}

// UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
func (o *Interface) UnsetMacAddress() {
	o.MacAddress.Unset()
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *Interface) GetMgmtOnly() bool {
	if o == nil || o.MgmtOnly == nil {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || o.MgmtOnly == nil {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *Interface) HasMgmtOnly() bool {
	if o != nil && o.MgmtOnly != nil {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *Interface) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Interface) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Interface) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Interface) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Interface) GetMode() Mode {
	if o == nil || o.Mode == nil {
		var ret Mode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetModeOk() (*Mode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Interface) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given Mode and assigns it to the Mode field.
func (o *Interface) SetMode(v Mode) {
	o.Mode = &v
}

// GetUntaggedVlan returns the UntaggedVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interface) GetUntaggedVlan() NestedVLAN {
	if o == nil || o.UntaggedVlan.Get() == nil {
		var ret NestedVLAN
		return ret
	}
	return *o.UntaggedVlan.Get()
}

// GetUntaggedVlanOk returns a tuple with the UntaggedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interface) GetUntaggedVlanOk() (*NestedVLAN, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntaggedVlan.Get(), o.UntaggedVlan.IsSet()
}

// HasUntaggedVlan returns a boolean if a field has been set.
func (o *Interface) HasUntaggedVlan() bool {
	if o != nil && o.UntaggedVlan.IsSet() {
		return true
	}

	return false
}

// SetUntaggedVlan gets a reference to the given NullableNestedVLAN and assigns it to the UntaggedVlan field.
func (o *Interface) SetUntaggedVlan(v NestedVLAN) {
	o.UntaggedVlan.Set(&v)
}
// SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil
func (o *Interface) SetUntaggedVlanNil() {
	o.UntaggedVlan.Set(nil)
}

// UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
func (o *Interface) UnsetUntaggedVlan() {
	o.UntaggedVlan.Unset()
}

// GetTaggedVlans returns the TaggedVlans field value if set, zero value otherwise.
func (o *Interface) GetTaggedVlans() []NestedVLAN {
	if o == nil || o.TaggedVlans == nil {
		var ret []NestedVLAN
		return ret
	}
	return o.TaggedVlans
}

// GetTaggedVlansOk returns a tuple with the TaggedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetTaggedVlansOk() ([]NestedVLAN, bool) {
	if o == nil || o.TaggedVlans == nil {
		return nil, false
	}
	return o.TaggedVlans, true
}

// HasTaggedVlans returns a boolean if a field has been set.
func (o *Interface) HasTaggedVlans() bool {
	if o != nil && o.TaggedVlans != nil {
		return true
	}

	return false
}

// SetTaggedVlans gets a reference to the given []NestedVLAN and assigns it to the TaggedVlans field.
func (o *Interface) SetTaggedVlans(v []NestedVLAN) {
	o.TaggedVlans = v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *Interface) GetMarkConnected() bool {
	if o == nil || o.MarkConnected == nil {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || o.MarkConnected == nil {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *Interface) HasMarkConnected() bool {
	if o != nil && o.MarkConnected != nil {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *Interface) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *Interface) GetCable() NestedCable {
	if o == nil || o.Cable == nil {
		var ret NestedCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetCableOk() (*NestedCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *Interface) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given NestedCable and assigns it to the Cable field.
func (o *Interface) SetCable(v NestedCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise.
func (o *Interface) GetCablePeer() map[string]string {
	if o == nil || o.CablePeer == nil {
		var ret map[string]string
		return ret
	}
	return *o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetCablePeerOk() (*map[string]string, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *Interface) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]string and assigns it to the CablePeer field.
func (o *Interface) SetCablePeer(v map[string]string) {
	o.CablePeer = &v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise.
func (o *Interface) GetCablePeerType() string {
	if o == nil || o.CablePeerType == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetCablePeerTypeOk() (*string, bool) {
	if o == nil || o.CablePeerType == nil {
		return nil, false
	}
	return o.CablePeerType, true
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *Interface) HasCablePeerType() bool {
	if o != nil && o.CablePeerType != nil {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given string and assigns it to the CablePeerType field.
func (o *Interface) SetCablePeerType(v string) {
	o.CablePeerType = &v
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value if set, zero value otherwise.
func (o *Interface) GetConnectedEndpoint() map[string]string {
	if o == nil || o.ConnectedEndpoint == nil {
		var ret map[string]string
		return ret
	}
	return *o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetConnectedEndpointOk() (*map[string]string, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// HasConnectedEndpoint returns a boolean if a field has been set.
func (o *Interface) HasConnectedEndpoint() bool {
	if o != nil && o.ConnectedEndpoint != nil {
		return true
	}

	return false
}

// SetConnectedEndpoint gets a reference to the given map[string]string and assigns it to the ConnectedEndpoint field.
func (o *Interface) SetConnectedEndpoint(v map[string]string) {
	o.ConnectedEndpoint = &v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value if set, zero value otherwise.
func (o *Interface) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType == nil {
		var ret string
		return ret
	}
	return *o.ConnectedEndpointType
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil || o.ConnectedEndpointType == nil {
		return nil, false
	}
	return o.ConnectedEndpointType, true
}

// HasConnectedEndpointType returns a boolean if a field has been set.
func (o *Interface) HasConnectedEndpointType() bool {
	if o != nil && o.ConnectedEndpointType != nil {
		return true
	}

	return false
}

// SetConnectedEndpointType gets a reference to the given string and assigns it to the ConnectedEndpointType field.
func (o *Interface) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType = &v
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value if set, zero value otherwise.
func (o *Interface) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable == nil {
		var ret bool
		return ret
	}
	return *o.ConnectedEndpointReachable
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil || o.ConnectedEndpointReachable == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable, true
}

// HasConnectedEndpointReachable returns a boolean if a field has been set.
func (o *Interface) HasConnectedEndpointReachable() bool {
	if o != nil && o.ConnectedEndpointReachable != nil {
		return true
	}

	return false
}

// SetConnectedEndpointReachable gets a reference to the given bool and assigns it to the ConnectedEndpointReachable field.
func (o *Interface) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Interface) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Interface) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Interface) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Interface) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Interface) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Interface) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Interface) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Interface) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Interface) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Interface) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Interface) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Interface) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCountIpaddresses returns the CountIpaddresses field value if set, zero value otherwise.
func (o *Interface) GetCountIpaddresses() int32 {
	if o == nil || o.CountIpaddresses == nil {
		var ret int32
		return ret
	}
	return *o.CountIpaddresses
}

// GetCountIpaddressesOk returns a tuple with the CountIpaddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetCountIpaddressesOk() (*int32, bool) {
	if o == nil || o.CountIpaddresses == nil {
		return nil, false
	}
	return o.CountIpaddresses, true
}

// HasCountIpaddresses returns a boolean if a field has been set.
func (o *Interface) HasCountIpaddresses() bool {
	if o != nil && o.CountIpaddresses != nil {
		return true
	}

	return false
}

// SetCountIpaddresses gets a reference to the given int32 and assigns it to the CountIpaddresses field.
func (o *Interface) SetCountIpaddresses(v int32) {
	o.CountIpaddresses = &v
}

// GetOccupied returns the Occupied field value if set, zero value otherwise.
func (o *Interface) GetOccupied() bool {
	if o == nil || o.Occupied == nil {
		var ret bool
		return ret
	}
	return *o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetOccupiedOk() (*bool, bool) {
	if o == nil || o.Occupied == nil {
		return nil, false
	}
	return o.Occupied, true
}

// HasOccupied returns a boolean if a field has been set.
func (o *Interface) HasOccupied() bool {
	if o != nil && o.Occupied != nil {
		return true
	}

	return false
}

// SetOccupied gets a reference to the given bool and assigns it to the Occupied field.
func (o *Interface) SetOccupied(v bool) {
	o.Occupied = &v
}

func (o Interface) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Lag != nil {
		toSerialize["lag"] = o.Lag
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.MacAddress.IsSet() {
		toSerialize["mac_address"] = o.MacAddress.Get()
	}
	if o.MgmtOnly != nil {
		toSerialize["mgmt_only"] = o.MgmtOnly
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.UntaggedVlan.IsSet() {
		toSerialize["untagged_vlan"] = o.UntaggedVlan.Get()
	}
	if o.TaggedVlans != nil {
		toSerialize["tagged_vlans"] = o.TaggedVlans
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
	if o.CountIpaddresses != nil {
		toSerialize["count_ipaddresses"] = o.CountIpaddresses
	}
	if o.Occupied != nil {
		toSerialize["_occupied"] = o.Occupied
	}
	return json.Marshal(toSerialize)
}

type NullableInterface struct {
	value *Interface
	isSet bool
}

func (v NullableInterface) Get() *Interface {
	return v.value
}

func (v *NullableInterface) Set(val *Interface) {
	v.value = val
	v.isSet = true
}

func (v NullableInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterface(val *Interface) *NullableInterface {
	return &NullableInterface{value: val, isSet: true}
}

func (v NullableInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


