# WritableIPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Family** | Pointer to **string** |  | [optional] [readonly] 
**Address** | **string** | IPv4 or IPv6 address (with mask) | 
**Vrf** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **string** | The operational status of this IP | [optional] 
**Role** | Pointer to **string** | The functional role of this IP | [optional] 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableInt32** |  | [optional] 
**AssignedObject** | Pointer to **map[string]string** |  | [optional] [readonly] 
**NatInside** | Pointer to **NullableInt32** | The IP for which this address is the \&quot;outside\&quot; IP | [optional] 
**NatOutside** | Pointer to **string** |  | [optional] [readonly] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewWritableIPAddress

`func NewWritableIPAddress(address string, ) *WritableIPAddress`

NewWritableIPAddress instantiates a new WritableIPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableIPAddressWithDefaults

`func NewWritableIPAddressWithDefaults() *WritableIPAddress`

NewWritableIPAddressWithDefaults instantiates a new WritableIPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableIPAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableIPAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableIPAddress) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableIPAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableIPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableIPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableIPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableIPAddress) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableIPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableIPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableIPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableIPAddress) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFamily

`func (o *WritableIPAddress) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *WritableIPAddress) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *WritableIPAddress) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *WritableIPAddress) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetAddress

`func (o *WritableIPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WritableIPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WritableIPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVrf

`func (o *WritableIPAddress) GetVrf() int32`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *WritableIPAddress) GetVrfOk() (*int32, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *WritableIPAddress) SetVrf(v int32)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *WritableIPAddress) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *WritableIPAddress) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *WritableIPAddress) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *WritableIPAddress) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableIPAddress) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableIPAddress) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableIPAddress) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableIPAddress) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableIPAddress) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *WritableIPAddress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableIPAddress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableIPAddress) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableIPAddress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *WritableIPAddress) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableIPAddress) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableIPAddress) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritableIPAddress) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *WritableIPAddress) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *WritableIPAddress) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *WritableIPAddress) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *WritableIPAddress) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *WritableIPAddress) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *WritableIPAddress) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *WritableIPAddress) GetAssignedObjectId() int32`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *WritableIPAddress) GetAssignedObjectIdOk() (*int32, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *WritableIPAddress) SetAssignedObjectId(v int32)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *WritableIPAddress) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *WritableIPAddress) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *WritableIPAddress) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetAssignedObject

`func (o *WritableIPAddress) GetAssignedObject() map[string]string`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *WritableIPAddress) GetAssignedObjectOk() (*map[string]string, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *WritableIPAddress) SetAssignedObject(v map[string]string)`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *WritableIPAddress) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### GetNatInside

`func (o *WritableIPAddress) GetNatInside() int32`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *WritableIPAddress) GetNatInsideOk() (*int32, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *WritableIPAddress) SetNatInside(v int32)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *WritableIPAddress) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *WritableIPAddress) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *WritableIPAddress) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetNatOutside

`func (o *WritableIPAddress) GetNatOutside() string`

GetNatOutside returns the NatOutside field if non-nil, zero value otherwise.

### GetNatOutsideOk

`func (o *WritableIPAddress) GetNatOutsideOk() (*string, bool)`

GetNatOutsideOk returns a tuple with the NatOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutside

`func (o *WritableIPAddress) SetNatOutside(v string)`

SetNatOutside sets NatOutside field to given value.

### HasNatOutside

`func (o *WritableIPAddress) HasNatOutside() bool`

HasNatOutside returns a boolean if a field has been set.

### GetDnsName

`func (o *WritableIPAddress) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *WritableIPAddress) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *WritableIPAddress) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *WritableIPAddress) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *WritableIPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableIPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableIPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableIPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *WritableIPAddress) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableIPAddress) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableIPAddress) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableIPAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableIPAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableIPAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableIPAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableIPAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableIPAddress) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableIPAddress) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableIPAddress) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableIPAddress) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableIPAddress) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableIPAddress) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableIPAddress) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableIPAddress) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


