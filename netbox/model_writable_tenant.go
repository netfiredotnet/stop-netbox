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

// WritableTenant struct for WritableTenant
type WritableTenant struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Group NullableInt32 `json:"group,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	CircuitCount *int32 `json:"circuit_count,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`
	PrefixCount *int32 `json:"prefix_count,omitempty"`
	RackCount *int32 `json:"rack_count,omitempty"`
	SiteCount *int32 `json:"site_count,omitempty"`
	VirtualmachineCount *int32 `json:"virtualmachine_count,omitempty"`
	VlanCount *int32 `json:"vlan_count,omitempty"`
	VrfCount *int32 `json:"vrf_count,omitempty"`
	ClusterCount *int32 `json:"cluster_count,omitempty"`
}

// NewWritableTenant instantiates a new WritableTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableTenant(name string, slug string) *WritableTenant {
	this := WritableTenant{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewWritableTenantWithDefaults instantiates a new WritableTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableTenantWithDefaults() *WritableTenant {
	this := WritableTenant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritableTenant) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritableTenant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritableTenant) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritableTenant) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritableTenant) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritableTenant) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritableTenant) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritableTenant) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritableTenant) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value
func (o *WritableTenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableTenant) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritableTenant) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritableTenant) SetSlug(v string) {
	o.Slug = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableTenant) GetGroup() int32 {
	if o == nil || o.Group.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableTenant) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *WritableTenant) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *WritableTenant) SetGroup(v int32) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *WritableTenant) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *WritableTenant) UnsetGroup() {
	o.Group.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableTenant) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableTenant) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableTenant) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableTenant) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableTenant) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableTenant) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableTenant) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableTenant) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WritableTenant) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableTenant) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableTenant) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableTenant) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritableTenant) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritableTenant) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritableTenant) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritableTenant) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritableTenant) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritableTenant) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCircuitCount returns the CircuitCount field value if set, zero value otherwise.
func (o *WritableTenant) GetCircuitCount() int32 {
	if o == nil || o.CircuitCount == nil {
		var ret int32
		return ret
	}
	return *o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetCircuitCountOk() (*int32, bool) {
	if o == nil || o.CircuitCount == nil {
		return nil, false
	}
	return o.CircuitCount, true
}

// HasCircuitCount returns a boolean if a field has been set.
func (o *WritableTenant) HasCircuitCount() bool {
	if o != nil && o.CircuitCount != nil {
		return true
	}

	return false
}

// SetCircuitCount gets a reference to the given int32 and assigns it to the CircuitCount field.
func (o *WritableTenant) SetCircuitCount(v int32) {
	o.CircuitCount = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *WritableTenant) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetDeviceCountOk() (*int32, bool) {
	if o == nil || o.DeviceCount == nil {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *WritableTenant) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *WritableTenant) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetIpaddressCount returns the IpaddressCount field value if set, zero value otherwise.
func (o *WritableTenant) GetIpaddressCount() int32 {
	if o == nil || o.IpaddressCount == nil {
		var ret int32
		return ret
	}
	return *o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetIpaddressCountOk() (*int32, bool) {
	if o == nil || o.IpaddressCount == nil {
		return nil, false
	}
	return o.IpaddressCount, true
}

// HasIpaddressCount returns a boolean if a field has been set.
func (o *WritableTenant) HasIpaddressCount() bool {
	if o != nil && o.IpaddressCount != nil {
		return true
	}

	return false
}

// SetIpaddressCount gets a reference to the given int32 and assigns it to the IpaddressCount field.
func (o *WritableTenant) SetIpaddressCount(v int32) {
	o.IpaddressCount = &v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *WritableTenant) GetPrefixCount() int32 {
	if o == nil || o.PrefixCount == nil {
		var ret int32
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetPrefixCountOk() (*int32, bool) {
	if o == nil || o.PrefixCount == nil {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *WritableTenant) HasPrefixCount() bool {
	if o != nil && o.PrefixCount != nil {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int32 and assigns it to the PrefixCount field.
func (o *WritableTenant) SetPrefixCount(v int32) {
	o.PrefixCount = &v
}

// GetRackCount returns the RackCount field value if set, zero value otherwise.
func (o *WritableTenant) GetRackCount() int32 {
	if o == nil || o.RackCount == nil {
		var ret int32
		return ret
	}
	return *o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetRackCountOk() (*int32, bool) {
	if o == nil || o.RackCount == nil {
		return nil, false
	}
	return o.RackCount, true
}

// HasRackCount returns a boolean if a field has been set.
func (o *WritableTenant) HasRackCount() bool {
	if o != nil && o.RackCount != nil {
		return true
	}

	return false
}

// SetRackCount gets a reference to the given int32 and assigns it to the RackCount field.
func (o *WritableTenant) SetRackCount(v int32) {
	o.RackCount = &v
}

// GetSiteCount returns the SiteCount field value if set, zero value otherwise.
func (o *WritableTenant) GetSiteCount() int32 {
	if o == nil || o.SiteCount == nil {
		var ret int32
		return ret
	}
	return *o.SiteCount
}

// GetSiteCountOk returns a tuple with the SiteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetSiteCountOk() (*int32, bool) {
	if o == nil || o.SiteCount == nil {
		return nil, false
	}
	return o.SiteCount, true
}

// HasSiteCount returns a boolean if a field has been set.
func (o *WritableTenant) HasSiteCount() bool {
	if o != nil && o.SiteCount != nil {
		return true
	}

	return false
}

// SetSiteCount gets a reference to the given int32 and assigns it to the SiteCount field.
func (o *WritableTenant) SetSiteCount(v int32) {
	o.SiteCount = &v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *WritableTenant) GetVirtualmachineCount() int32 {
	if o == nil || o.VirtualmachineCount == nil {
		var ret int32
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil || o.VirtualmachineCount == nil {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *WritableTenant) HasVirtualmachineCount() bool {
	if o != nil && o.VirtualmachineCount != nil {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int32 and assigns it to the VirtualmachineCount field.
func (o *WritableTenant) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = &v
}

// GetVlanCount returns the VlanCount field value if set, zero value otherwise.
func (o *WritableTenant) GetVlanCount() int32 {
	if o == nil || o.VlanCount == nil {
		var ret int32
		return ret
	}
	return *o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetVlanCountOk() (*int32, bool) {
	if o == nil || o.VlanCount == nil {
		return nil, false
	}
	return o.VlanCount, true
}

// HasVlanCount returns a boolean if a field has been set.
func (o *WritableTenant) HasVlanCount() bool {
	if o != nil && o.VlanCount != nil {
		return true
	}

	return false
}

// SetVlanCount gets a reference to the given int32 and assigns it to the VlanCount field.
func (o *WritableTenant) SetVlanCount(v int32) {
	o.VlanCount = &v
}

// GetVrfCount returns the VrfCount field value if set, zero value otherwise.
func (o *WritableTenant) GetVrfCount() int32 {
	if o == nil || o.VrfCount == nil {
		var ret int32
		return ret
	}
	return *o.VrfCount
}

// GetVrfCountOk returns a tuple with the VrfCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetVrfCountOk() (*int32, bool) {
	if o == nil || o.VrfCount == nil {
		return nil, false
	}
	return o.VrfCount, true
}

// HasVrfCount returns a boolean if a field has been set.
func (o *WritableTenant) HasVrfCount() bool {
	if o != nil && o.VrfCount != nil {
		return true
	}

	return false
}

// SetVrfCount gets a reference to the given int32 and assigns it to the VrfCount field.
func (o *WritableTenant) SetVrfCount(v int32) {
	o.VrfCount = &v
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *WritableTenant) GetClusterCount() int32 {
	if o == nil || o.ClusterCount == nil {
		var ret int32
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTenant) GetClusterCountOk() (*int32, bool) {
	if o == nil || o.ClusterCount == nil {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *WritableTenant) HasClusterCount() bool {
	if o != nil && o.ClusterCount != nil {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int32 and assigns it to the ClusterCount field.
func (o *WritableTenant) SetClusterCount(v int32) {
	o.ClusterCount = &v
}

func (o WritableTenant) MarshalJSON() ([]byte, error) {
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
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
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
	if o.CircuitCount != nil {
		toSerialize["circuit_count"] = o.CircuitCount
	}
	if o.DeviceCount != nil {
		toSerialize["device_count"] = o.DeviceCount
	}
	if o.IpaddressCount != nil {
		toSerialize["ipaddress_count"] = o.IpaddressCount
	}
	if o.PrefixCount != nil {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if o.RackCount != nil {
		toSerialize["rack_count"] = o.RackCount
	}
	if o.SiteCount != nil {
		toSerialize["site_count"] = o.SiteCount
	}
	if o.VirtualmachineCount != nil {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	if o.VlanCount != nil {
		toSerialize["vlan_count"] = o.VlanCount
	}
	if o.VrfCount != nil {
		toSerialize["vrf_count"] = o.VrfCount
	}
	if o.ClusterCount != nil {
		toSerialize["cluster_count"] = o.ClusterCount
	}
	return json.Marshal(toSerialize)
}

type NullableWritableTenant struct {
	value *WritableTenant
	isSet bool
}

func (v NullableWritableTenant) Get() *WritableTenant {
	return v.value
}

func (v *NullableWritableTenant) Set(val *WritableTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableTenant(val *WritableTenant) *NullableWritableTenant {
	return &NullableWritableTenant{value: val, isSet: true}
}

func (v NullableWritableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


