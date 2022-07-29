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

// RackUnit struct for RackUnit
type RackUnit struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Face *Face1 `json:"face,omitempty"`
	Device *NestedDevice `json:"device,omitempty"`
	Occupied *bool `json:"occupied,omitempty"`
	Display *string `json:"display,omitempty"`
}

// NewRackUnit instantiates a new RackUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackUnit() *RackUnit {
	this := RackUnit{}
	return &this
}

// NewRackUnitWithDefaults instantiates a new RackUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackUnitWithDefaults() *RackUnit {
	this := RackUnit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RackUnit) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnit) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RackUnit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RackUnit) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RackUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RackUnit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RackUnit) SetName(v string) {
	o.Name = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *RackUnit) GetFace() Face1 {
	if o == nil || o.Face == nil {
		var ret Face1
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnit) GetFaceOk() (*Face1, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *RackUnit) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given Face1 and assigns it to the Face field.
func (o *RackUnit) SetFace(v Face1) {
	o.Face = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *RackUnit) GetDevice() NestedDevice {
	if o == nil || o.Device == nil {
		var ret NestedDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnit) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *RackUnit) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NestedDevice and assigns it to the Device field.
func (o *RackUnit) SetDevice(v NestedDevice) {
	o.Device = &v
}

// GetOccupied returns the Occupied field value if set, zero value otherwise.
func (o *RackUnit) GetOccupied() bool {
	if o == nil || o.Occupied == nil {
		var ret bool
		return ret
	}
	return *o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnit) GetOccupiedOk() (*bool, bool) {
	if o == nil || o.Occupied == nil {
		return nil, false
	}
	return o.Occupied, true
}

// HasOccupied returns a boolean if a field has been set.
func (o *RackUnit) HasOccupied() bool {
	if o != nil && o.Occupied != nil {
		return true
	}

	return false
}

// SetOccupied gets a reference to the given bool and assigns it to the Occupied field.
func (o *RackUnit) SetOccupied(v bool) {
	o.Occupied = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *RackUnit) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnit) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *RackUnit) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *RackUnit) SetDisplay(v string) {
	o.Display = &v
}

func (o RackUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Face != nil {
		toSerialize["face"] = o.Face
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Occupied != nil {
		toSerialize["occupied"] = o.Occupied
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableRackUnit struct {
	value *RackUnit
	isSet bool
}

func (v NullableRackUnit) Get() *RackUnit {
	return v.value
}

func (v *NullableRackUnit) Set(val *RackUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableRackUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableRackUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackUnit(val *RackUnit) *NullableRackUnit {
	return &NullableRackUnit{value: val, isSet: true}
}

func (v NullableRackUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


