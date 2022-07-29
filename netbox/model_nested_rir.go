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

// NestedRIR struct for NestedRIR
type NestedRIR struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	AggregateCount *int32 `json:"aggregate_count,omitempty"`
}

// NewNestedRIR instantiates a new NestedRIR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRIR(name string, slug string) *NestedRIR {
	this := NestedRIR{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedRIRWithDefaults instantiates a new NestedRIR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRIRWithDefaults() *NestedRIR {
	this := NestedRIR{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedRIR) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRIR) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedRIR) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NestedRIR) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NestedRIR) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRIR) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NestedRIR) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NestedRIR) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *NestedRIR) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRIR) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *NestedRIR) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *NestedRIR) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *NestedRIR) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRIR) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRIR) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedRIR) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedRIR) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedRIR) SetSlug(v string) {
	o.Slug = v
}

// GetAggregateCount returns the AggregateCount field value if set, zero value otherwise.
func (o *NestedRIR) GetAggregateCount() int32 {
	if o == nil || o.AggregateCount == nil {
		var ret int32
		return ret
	}
	return *o.AggregateCount
}

// GetAggregateCountOk returns a tuple with the AggregateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRIR) GetAggregateCountOk() (*int32, bool) {
	if o == nil || o.AggregateCount == nil {
		return nil, false
	}
	return o.AggregateCount, true
}

// HasAggregateCount returns a boolean if a field has been set.
func (o *NestedRIR) HasAggregateCount() bool {
	if o != nil && o.AggregateCount != nil {
		return true
	}

	return false
}

// SetAggregateCount gets a reference to the given int32 and assigns it to the AggregateCount field.
func (o *NestedRIR) SetAggregateCount(v int32) {
	o.AggregateCount = &v
}

func (o NestedRIR) MarshalJSON() ([]byte, error) {
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
	if o.AggregateCount != nil {
		toSerialize["aggregate_count"] = o.AggregateCount
	}
	return json.Marshal(toSerialize)
}

type NullableNestedRIR struct {
	value *NestedRIR
	isSet bool
}

func (v NullableNestedRIR) Get() *NestedRIR {
	return v.value
}

func (v *NullableNestedRIR) Set(val *NestedRIR) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRIR) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRIR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRIR(val *NestedRIR) *NullableNestedRIR {
	return &NullableNestedRIR{value: val, isSet: true}
}

func (v NullableNestedRIR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRIR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

