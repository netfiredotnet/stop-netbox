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

// InterfaceConnection struct for InterfaceConnection
type InterfaceConnection struct {
	InterfaceA *NestedInterface `json:"interface_a,omitempty"`
	InterfaceB NestedInterface `json:"interface_b"`
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`
}

// NewInterfaceConnection instantiates a new InterfaceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceConnection(interfaceB NestedInterface) *InterfaceConnection {
	this := InterfaceConnection{}
	this.InterfaceB = interfaceB
	return &this
}

// NewInterfaceConnectionWithDefaults instantiates a new InterfaceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceConnectionWithDefaults() *InterfaceConnection {
	this := InterfaceConnection{}
	return &this
}

// GetInterfaceA returns the InterfaceA field value if set, zero value otherwise.
func (o *InterfaceConnection) GetInterfaceA() NestedInterface {
	if o == nil || o.InterfaceA == nil {
		var ret NestedInterface
		return ret
	}
	return *o.InterfaceA
}

// GetInterfaceAOk returns a tuple with the InterfaceA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetInterfaceAOk() (*NestedInterface, bool) {
	if o == nil || o.InterfaceA == nil {
		return nil, false
	}
	return o.InterfaceA, true
}

// HasInterfaceA returns a boolean if a field has been set.
func (o *InterfaceConnection) HasInterfaceA() bool {
	if o != nil && o.InterfaceA != nil {
		return true
	}

	return false
}

// SetInterfaceA gets a reference to the given NestedInterface and assigns it to the InterfaceA field.
func (o *InterfaceConnection) SetInterfaceA(v NestedInterface) {
	o.InterfaceA = &v
}

// GetInterfaceB returns the InterfaceB field value
func (o *InterfaceConnection) GetInterfaceB() NestedInterface {
	if o == nil {
		var ret NestedInterface
		return ret
	}

	return o.InterfaceB
}

// GetInterfaceBOk returns a tuple with the InterfaceB field value
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetInterfaceBOk() (*NestedInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceB, true
}

// SetInterfaceB sets field value
func (o *InterfaceConnection) SetInterfaceB(v NestedInterface) {
	o.InterfaceB = v
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value if set, zero value otherwise.
func (o *InterfaceConnection) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable == nil {
		var ret bool
		return ret
	}
	return *o.ConnectedEndpointReachable
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil || o.ConnectedEndpointReachable == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable, true
}

// HasConnectedEndpointReachable returns a boolean if a field has been set.
func (o *InterfaceConnection) HasConnectedEndpointReachable() bool {
	if o != nil && o.ConnectedEndpointReachable != nil {
		return true
	}

	return false
}

// SetConnectedEndpointReachable gets a reference to the given bool and assigns it to the ConnectedEndpointReachable field.
func (o *InterfaceConnection) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable = &v
}

func (o InterfaceConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InterfaceA != nil {
		toSerialize["interface_a"] = o.InterfaceA
	}
	if true {
		toSerialize["interface_b"] = o.InterfaceB
	}
	if o.ConnectedEndpointReachable != nil {
		toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceConnection struct {
	value *InterfaceConnection
	isSet bool
}

func (v NullableInterfaceConnection) Get() *InterfaceConnection {
	return v.value
}

func (v *NullableInterfaceConnection) Set(val *InterfaceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceConnection(val *InterfaceConnection) *NullableInterfaceConnection {
	return &NullableInterfaceConnection{value: val, isSet: true}
}

func (v NullableInterfaceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

