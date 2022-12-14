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

// NestedIPAddress struct for NestedIPAddress
type NestedIPAddress struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Family *int32 `json:"family,omitempty"`
	// IPv4 or IPv6 address (with mask)
	Address string `json:"address"`
}

// NewNestedIPAddress instantiates a new NestedIPAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedIPAddress(address string) *NestedIPAddress {
	this := NestedIPAddress{}
	this.Address = address
	return &this
}

// NewNestedIPAddressWithDefaults instantiates a new NestedIPAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedIPAddressWithDefaults() *NestedIPAddress {
	this := NestedIPAddress{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedIPAddress) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedIPAddress) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedIPAddress) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedIPAddress) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedIPAddress) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedIPAddress) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedIPAddress) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedIPAddress) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedIPAddress) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedIPAddress) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedIPAddress) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedIPAddress) SetDisplay(v string) {
	o.Display = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *NestedIPAddress) GetFamily() int32 {
	if o == nil || o.Family == nil {
		var ret int32
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedIPAddress) GetFamilyOk() (*int32, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *NestedIPAddress) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given int32 and assigns it to the Family field.
func (o *NestedIPAddress) SetFamily(v int32) {
	o.Family = &v
}

// GetAddress returns the Address field value
func (o *NestedIPAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NestedIPAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NestedIPAddress) SetAddress(v string) {
	o.Address = v
}

func (o NestedIPAddress) MarshalJSON() ([]byte, error) {
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
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNestedIPAddress struct {
	value *NestedIPAddress
	isSet bool
}

func (v NullableNestedIPAddress) Get() *NestedIPAddress {
	return v.value
}

func (v *NullableNestedIPAddress) Set(val *NestedIPAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedIPAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedIPAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedIPAddress(val *NestedIPAddress) *NullableNestedIPAddress {
	return &NullableNestedIPAddress{value: val, isSet: true}
}

func (v NullableNestedIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedIPAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


