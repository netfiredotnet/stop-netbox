/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CircuitTermination struct for CircuitTermination
type CircuitTermination struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Circuit NestedCircuit `json:"circuit"`
	TermSide string `json:"term_side"`
	Site *NestedSite `json:"site,omitempty"`
	ProviderNetwork *NestedProviderNetwork `json:"provider_network,omitempty"`
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	XconnectId *string `json:"xconnect_id,omitempty"`
	PpInfo *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Cable *NestedCable `json:"cable,omitempty"`
	//  Return the appropriate serializer for the cable termination model. 
	CablePeer *map[string]string `json:"cable_peer,omitempty"`
	CablePeerType *string `json:"cable_peer_type,omitempty"`
	Occupied *bool `json:"_occupied,omitempty"`
}

// NewCircuitTermination instantiates a new CircuitTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitTermination(circuit NestedCircuit, termSide string) *CircuitTermination {
	this := CircuitTermination{}
	this.Circuit = circuit
	this.TermSide = termSide
	return &this
}

// NewCircuitTerminationWithDefaults instantiates a new CircuitTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitTerminationWithDefaults() *CircuitTermination {
	this := CircuitTermination{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CircuitTermination) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CircuitTermination) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CircuitTermination) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CircuitTermination) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CircuitTermination) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CircuitTermination) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *CircuitTermination) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *CircuitTermination) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *CircuitTermination) SetDisplay(v string) {
	o.Display = &v
}

// GetCircuit returns the Circuit field value
func (o *CircuitTermination) GetCircuit() NestedCircuit {
	if o == nil {
		var ret NestedCircuit
		return ret
	}

	return o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetCircuitOk() (*NestedCircuit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Circuit, true
}

// SetCircuit sets field value
func (o *CircuitTermination) SetCircuit(v NestedCircuit) {
	o.Circuit = v
}

// GetTermSide returns the TermSide field value
func (o *CircuitTermination) GetTermSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TermSide
}

// GetTermSideOk returns a tuple with the TermSide field value
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetTermSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermSide, true
}

// SetTermSide sets field value
func (o *CircuitTermination) SetTermSide(v string) {
	o.TermSide = v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *CircuitTermination) GetSite() NestedSite {
	if o == nil || o.Site == nil {
		var ret NestedSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetSiteOk() (*NestedSite, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *CircuitTermination) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given NestedSite and assigns it to the Site field.
func (o *CircuitTermination) SetSite(v NestedSite) {
	o.Site = &v
}

// GetProviderNetwork returns the ProviderNetwork field value if set, zero value otherwise.
func (o *CircuitTermination) GetProviderNetwork() NestedProviderNetwork {
	if o == nil || o.ProviderNetwork == nil {
		var ret NestedProviderNetwork
		return ret
	}
	return *o.ProviderNetwork
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetProviderNetworkOk() (*NestedProviderNetwork, bool) {
	if o == nil || o.ProviderNetwork == nil {
		return nil, false
	}
	return o.ProviderNetwork, true
}

// HasProviderNetwork returns a boolean if a field has been set.
func (o *CircuitTermination) HasProviderNetwork() bool {
	if o != nil && o.ProviderNetwork != nil {
		return true
	}

	return false
}

// SetProviderNetwork gets a reference to the given NestedProviderNetwork and assigns it to the ProviderNetwork field.
func (o *CircuitTermination) SetProviderNetwork(v NestedProviderNetwork) {
	o.ProviderNetwork = &v
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitTermination) GetPortSpeed() int32 {
	if o == nil || o.PortSpeed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitTermination) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *CircuitTermination) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *CircuitTermination) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}
// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *CircuitTermination) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *CircuitTermination) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitTermination) GetUpstreamSpeed() int32 {
	if o == nil || o.UpstreamSpeed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitTermination) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *CircuitTermination) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *CircuitTermination) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}
// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *CircuitTermination) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *CircuitTermination) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *CircuitTermination) GetXconnectId() string {
	if o == nil || o.XconnectId == nil {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetXconnectIdOk() (*string, bool) {
	if o == nil || o.XconnectId == nil {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *CircuitTermination) HasXconnectId() bool {
	if o != nil && o.XconnectId != nil {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *CircuitTermination) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *CircuitTermination) GetPpInfo() string {
	if o == nil || o.PpInfo == nil {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetPpInfoOk() (*string, bool) {
	if o == nil || o.PpInfo == nil {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *CircuitTermination) HasPpInfo() bool {
	if o != nil && o.PpInfo != nil {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *CircuitTermination) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CircuitTermination) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CircuitTermination) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CircuitTermination) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *CircuitTermination) GetMarkConnected() bool {
	if o == nil || o.MarkConnected == nil {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || o.MarkConnected == nil {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *CircuitTermination) HasMarkConnected() bool {
	if o != nil && o.MarkConnected != nil {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *CircuitTermination) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *CircuitTermination) GetCable() NestedCable {
	if o == nil || o.Cable == nil {
		var ret NestedCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetCableOk() (*NestedCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *CircuitTermination) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given NestedCable and assigns it to the Cable field.
func (o *CircuitTermination) SetCable(v NestedCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise.
func (o *CircuitTermination) GetCablePeer() map[string]string {
	if o == nil || o.CablePeer == nil {
		var ret map[string]string
		return ret
	}
	return *o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetCablePeerOk() (*map[string]string, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *CircuitTermination) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]string and assigns it to the CablePeer field.
func (o *CircuitTermination) SetCablePeer(v map[string]string) {
	o.CablePeer = &v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise.
func (o *CircuitTermination) GetCablePeerType() string {
	if o == nil || o.CablePeerType == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetCablePeerTypeOk() (*string, bool) {
	if o == nil || o.CablePeerType == nil {
		return nil, false
	}
	return o.CablePeerType, true
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *CircuitTermination) HasCablePeerType() bool {
	if o != nil && o.CablePeerType != nil {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given string and assigns it to the CablePeerType field.
func (o *CircuitTermination) SetCablePeerType(v string) {
	o.CablePeerType = &v
}

// GetOccupied returns the Occupied field value if set, zero value otherwise.
func (o *CircuitTermination) GetOccupied() bool {
	if o == nil || o.Occupied == nil {
		var ret bool
		return ret
	}
	return *o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTermination) GetOccupiedOk() (*bool, bool) {
	if o == nil || o.Occupied == nil {
		return nil, false
	}
	return o.Occupied, true
}

// HasOccupied returns a boolean if a field has been set.
func (o *CircuitTermination) HasOccupied() bool {
	if o != nil && o.Occupied != nil {
		return true
	}

	return false
}

// SetOccupied gets a reference to the given bool and assigns it to the Occupied field.
func (o *CircuitTermination) SetOccupied(v bool) {
	o.Occupied = &v
}

func (o CircuitTermination) MarshalJSON() ([]byte, error) {
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
		toSerialize["circuit"] = o.Circuit
	}
	if true {
		toSerialize["term_side"] = o.TermSide
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.ProviderNetwork != nil {
		toSerialize["provider_network"] = o.ProviderNetwork
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if o.XconnectId != nil {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if o.PpInfo != nil {
		toSerialize["pp_info"] = o.PpInfo
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
	if o.Occupied != nil {
		toSerialize["_occupied"] = o.Occupied
	}
	return json.Marshal(toSerialize)
}

type NullableCircuitTermination struct {
	value *CircuitTermination
	isSet bool
}

func (v NullableCircuitTermination) Get() *CircuitTermination {
	return v.value
}

func (v *NullableCircuitTermination) Set(val *CircuitTermination) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitTermination(val *CircuitTermination) *NullableCircuitTermination {
	return &NullableCircuitTermination{value: val, isSet: true}
}

func (v NullableCircuitTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


