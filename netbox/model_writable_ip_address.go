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

// WritableIPAddress struct for WritableIPAddress
type WritableIPAddress struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Family *string `json:"family,omitempty"`
	// IPv4 or IPv6 address (with mask)
	Address string `json:"address"`
	Vrf NullableInt32 `json:"vrf,omitempty"`
	Tenant NullableInt32 `json:"tenant,omitempty"`
	// The operational status of this IP
	Status *string `json:"status,omitempty"`
	// The functional role of this IP
	Role *string `json:"role,omitempty"`
	AssignedObjectType NullableString `json:"assigned_object_type,omitempty"`
	AssignedObjectId NullableInt32 `json:"assigned_object_id,omitempty"`
	AssignedObject *map[string]string `json:"assigned_object,omitempty"`
	// The IP for which this address is the \"outside\" IP
	NatInside NullableInt32 `json:"nat_inside,omitempty"`
	NatOutside *string `json:"nat_outside,omitempty"`
	// Hostname or FQDN (not case-sensitive)
	DnsName *string `json:"dns_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewWritableIPAddress instantiates a new WritableIPAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableIPAddress(address string) *WritableIPAddress {
	this := WritableIPAddress{}
	this.Address = address
	return &this
}

// NewWritableIPAddressWithDefaults instantiates a new WritableIPAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableIPAddressWithDefaults() *WritableIPAddress {
	this := WritableIPAddress{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WritableIPAddress) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WritableIPAddress) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WritableIPAddress) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WritableIPAddress) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WritableIPAddress) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WritableIPAddress) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *WritableIPAddress) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *WritableIPAddress) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *WritableIPAddress) SetDisplay(v string) {
	o.Display = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *WritableIPAddress) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *WritableIPAddress) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *WritableIPAddress) SetFamily(v string) {
	o.Family = &v
}

// GetAddress returns the Address field value
func (o *WritableIPAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *WritableIPAddress) SetAddress(v string) {
	o.Address = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddress) GetVrf() int32 {
	if o == nil || o.Vrf.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddress) GetVrfOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *WritableIPAddress) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableInt32 and assigns it to the Vrf field.
func (o *WritableIPAddress) SetVrf(v int32) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *WritableIPAddress) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *WritableIPAddress) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddress) GetTenant() int32 {
	if o == nil || o.Tenant.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddress) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableIPAddress) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableIPAddress) SetTenant(v int32) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableIPAddress) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableIPAddress) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableIPAddress) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableIPAddress) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WritableIPAddress) SetStatus(v string) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *WritableIPAddress) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *WritableIPAddress) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *WritableIPAddress) SetRole(v string) {
	o.Role = &v
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddress) GetAssignedObjectType() string {
	if o == nil || o.AssignedObjectType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssignedObjectType.Get()
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddress) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectType.Get(), o.AssignedObjectType.IsSet()
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *WritableIPAddress) HasAssignedObjectType() bool {
	if o != nil && o.AssignedObjectType.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given NullableString and assigns it to the AssignedObjectType field.
func (o *WritableIPAddress) SetAssignedObjectType(v string) {
	o.AssignedObjectType.Set(&v)
}
// SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil
func (o *WritableIPAddress) SetAssignedObjectTypeNil() {
	o.AssignedObjectType.Set(nil)
}

// UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
func (o *WritableIPAddress) UnsetAssignedObjectType() {
	o.AssignedObjectType.Unset()
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddress) GetAssignedObjectId() int32 {
	if o == nil || o.AssignedObjectId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AssignedObjectId.Get()
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddress) GetAssignedObjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectId.Get(), o.AssignedObjectId.IsSet()
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *WritableIPAddress) HasAssignedObjectId() bool {
	if o != nil && o.AssignedObjectId.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given NullableInt32 and assigns it to the AssignedObjectId field.
func (o *WritableIPAddress) SetAssignedObjectId(v int32) {
	o.AssignedObjectId.Set(&v)
}
// SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil
func (o *WritableIPAddress) SetAssignedObjectIdNil() {
	o.AssignedObjectId.Set(nil)
}

// UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
func (o *WritableIPAddress) UnsetAssignedObjectId() {
	o.AssignedObjectId.Unset()
}

// GetAssignedObject returns the AssignedObject field value if set, zero value otherwise.
func (o *WritableIPAddress) GetAssignedObject() map[string]string {
	if o == nil || o.AssignedObject == nil {
		var ret map[string]string
		return ret
	}
	return *o.AssignedObject
}

// GetAssignedObjectOk returns a tuple with the AssignedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetAssignedObjectOk() (*map[string]string, bool) {
	if o == nil || o.AssignedObject == nil {
		return nil, false
	}
	return o.AssignedObject, true
}

// HasAssignedObject returns a boolean if a field has been set.
func (o *WritableIPAddress) HasAssignedObject() bool {
	if o != nil && o.AssignedObject != nil {
		return true
	}

	return false
}

// SetAssignedObject gets a reference to the given map[string]string and assigns it to the AssignedObject field.
func (o *WritableIPAddress) SetAssignedObject(v map[string]string) {
	o.AssignedObject = &v
}

// GetNatInside returns the NatInside field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddress) GetNatInside() int32 {
	if o == nil || o.NatInside.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NatInside.Get()
}

// GetNatInsideOk returns a tuple with the NatInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddress) GetNatInsideOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatInside.Get(), o.NatInside.IsSet()
}

// HasNatInside returns a boolean if a field has been set.
func (o *WritableIPAddress) HasNatInside() bool {
	if o != nil && o.NatInside.IsSet() {
		return true
	}

	return false
}

// SetNatInside gets a reference to the given NullableInt32 and assigns it to the NatInside field.
func (o *WritableIPAddress) SetNatInside(v int32) {
	o.NatInside.Set(&v)
}
// SetNatInsideNil sets the value for NatInside to be an explicit nil
func (o *WritableIPAddress) SetNatInsideNil() {
	o.NatInside.Set(nil)
}

// UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
func (o *WritableIPAddress) UnsetNatInside() {
	o.NatInside.Unset()
}

// GetNatOutside returns the NatOutside field value if set, zero value otherwise.
func (o *WritableIPAddress) GetNatOutside() string {
	if o == nil || o.NatOutside == nil {
		var ret string
		return ret
	}
	return *o.NatOutside
}

// GetNatOutsideOk returns a tuple with the NatOutside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetNatOutsideOk() (*string, bool) {
	if o == nil || o.NatOutside == nil {
		return nil, false
	}
	return o.NatOutside, true
}

// HasNatOutside returns a boolean if a field has been set.
func (o *WritableIPAddress) HasNatOutside() bool {
	if o != nil && o.NatOutside != nil {
		return true
	}

	return false
}

// SetNatOutside gets a reference to the given string and assigns it to the NatOutside field.
func (o *WritableIPAddress) SetNatOutside(v string) {
	o.NatOutside = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *WritableIPAddress) GetDnsName() string {
	if o == nil || o.DnsName == nil {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetDnsNameOk() (*string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *WritableIPAddress) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *WritableIPAddress) SetDnsName(v string) {
	o.DnsName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableIPAddress) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableIPAddress) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableIPAddress) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableIPAddress) GetTags() []NestedTag {
	if o == nil || o.Tags == nil {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableIPAddress) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WritableIPAddress) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableIPAddress) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableIPAddress) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableIPAddress) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WritableIPAddress) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WritableIPAddress) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *WritableIPAddress) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WritableIPAddress) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddress) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WritableIPAddress) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WritableIPAddress) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o WritableIPAddress) MarshalJSON() ([]byte, error) {
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
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.AssignedObjectType.IsSet() {
		toSerialize["assigned_object_type"] = o.AssignedObjectType.Get()
	}
	if o.AssignedObjectId.IsSet() {
		toSerialize["assigned_object_id"] = o.AssignedObjectId.Get()
	}
	if o.AssignedObject != nil {
		toSerialize["assigned_object"] = o.AssignedObject
	}
	if o.NatInside.IsSet() {
		toSerialize["nat_inside"] = o.NatInside.Get()
	}
	if o.NatOutside != nil {
		toSerialize["nat_outside"] = o.NatOutside
	}
	if o.DnsName != nil {
		toSerialize["dns_name"] = o.DnsName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	return json.Marshal(toSerialize)
}

type NullableWritableIPAddress struct {
	value *WritableIPAddress
	isSet bool
}

func (v NullableWritableIPAddress) Get() *WritableIPAddress {
	return v.value
}

func (v *NullableWritableIPAddress) Set(val *WritableIPAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableIPAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableIPAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableIPAddress(val *WritableIPAddress) *NullableWritableIPAddress {
	return &NullableWritableIPAddress{value: val, isSet: true}
}

func (v NullableWritableIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableIPAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


