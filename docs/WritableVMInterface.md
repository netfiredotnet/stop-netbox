# WritableVMInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**VirtualMachine** | **int32** |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**UntaggedVlan** | Pointer to **NullableInt32** |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**CountIpaddresses** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableVMInterface

`func NewWritableVMInterface(virtualMachine int32, name string, ) *WritableVMInterface`

NewWritableVMInterface instantiates a new WritableVMInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVMInterfaceWithDefaults

`func NewWritableVMInterfaceWithDefaults() *WritableVMInterface`

NewWritableVMInterfaceWithDefaults instantiates a new WritableVMInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableVMInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableVMInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableVMInterface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableVMInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableVMInterface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableVMInterface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableVMInterface) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableVMInterface) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableVMInterface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableVMInterface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableVMInterface) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableVMInterface) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *WritableVMInterface) GetVirtualMachine() int32`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *WritableVMInterface) GetVirtualMachineOk() (*int32, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *WritableVMInterface) SetVirtualMachine(v int32)`

SetVirtualMachine sets VirtualMachine field to given value.


### GetName

`func (o *WritableVMInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableVMInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableVMInterface) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *WritableVMInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WritableVMInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WritableVMInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WritableVMInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *WritableVMInterface) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WritableVMInterface) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WritableVMInterface) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WritableVMInterface) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *WritableVMInterface) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WritableVMInterface) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetMtu

`func (o *WritableVMInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *WritableVMInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *WritableVMInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *WritableVMInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *WritableVMInterface) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *WritableVMInterface) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMacAddress

`func (o *WritableVMInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *WritableVMInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *WritableVMInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *WritableVMInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *WritableVMInterface) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *WritableVMInterface) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetDescription

`func (o *WritableVMInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableVMInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableVMInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableVMInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *WritableVMInterface) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WritableVMInterface) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WritableVMInterface) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WritableVMInterface) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetUntaggedVlan

`func (o *WritableVMInterface) GetUntaggedVlan() int32`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *WritableVMInterface) GetUntaggedVlanOk() (*int32, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *WritableVMInterface) SetUntaggedVlan(v int32)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *WritableVMInterface) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *WritableVMInterface) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *WritableVMInterface) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *WritableVMInterface) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *WritableVMInterface) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *WritableVMInterface) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *WritableVMInterface) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetTags

`func (o *WritableVMInterface) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVMInterface) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVMInterface) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVMInterface) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVMInterface) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVMInterface) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVMInterface) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVMInterface) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableVMInterface) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableVMInterface) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableVMInterface) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WritableVMInterface) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WritableVMInterface) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableVMInterface) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableVMInterface) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WritableVMInterface) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCountIpaddresses

`func (o *WritableVMInterface) GetCountIpaddresses() int32`

GetCountIpaddresses returns the CountIpaddresses field if non-nil, zero value otherwise.

### GetCountIpaddressesOk

`func (o *WritableVMInterface) GetCountIpaddressesOk() (*int32, bool)`

GetCountIpaddressesOk returns a tuple with the CountIpaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountIpaddresses

`func (o *WritableVMInterface) SetCountIpaddresses(v int32)`

SetCountIpaddresses sets CountIpaddresses field to given value.

### HasCountIpaddresses

`func (o *WritableVMInterface) HasCountIpaddresses() bool`

HasCountIpaddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


