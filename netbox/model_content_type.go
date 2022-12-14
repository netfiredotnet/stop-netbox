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

// ContentType struct for ContentType
type ContentType struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	AppLabel string `json:"app_label"`
	Model string `json:"model"`
	DisplayName *string `json:"display_name,omitempty"`
}

// NewContentType instantiates a new ContentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentType(appLabel string, model string) *ContentType {
	this := ContentType{}
	this.AppLabel = appLabel
	this.Model = model
	return &this
}

// NewContentTypeWithDefaults instantiates a new ContentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTypeWithDefaults() *ContentType {
	this := ContentType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContentType) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContentType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContentType) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ContentType) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ContentType) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ContentType) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ContentType) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ContentType) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ContentType) SetDisplay(v string) {
	o.Display = &v
}

// GetAppLabel returns the AppLabel field value
func (o *ContentType) GetAppLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value
// and a boolean to check if the value has been set.
func (o *ContentType) GetAppLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabel, true
}

// SetAppLabel sets field value
func (o *ContentType) SetAppLabel(v string) {
	o.AppLabel = v
}

// GetModel returns the Model field value
func (o *ContentType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ContentType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ContentType) SetModel(v string) {
	o.Model = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ContentType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ContentType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ContentType) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o ContentType) MarshalJSON() ([]byte, error) {
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
		toSerialize["app_label"] = o.AppLabel
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableContentType struct {
	value *ContentType
	isSet bool
}

func (v NullableContentType) Get() *ContentType {
	return v.value
}

func (v *NullableContentType) Set(val *ContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentType(val *ContentType) *NullableContentType {
	return &NullableContentType{value: val, isSet: true}
}

func (v NullableContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


