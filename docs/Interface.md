# Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Device** | [**NestedDevice**](NestedDevice.md) |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**Type2**](Type2.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**NestedInterface**](NestedInterface.md) |  | [optional] 
**Lag** | Pointer to [**NestedInterface**](NestedInterface.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**Mode**](Mode.md) |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableNestedVLAN**](NestedVLAN.md) |  | [optional] 
**TaggedVlans** | Pointer to [**[]NestedVLAN**](NestedVLAN.md) |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | Pointer to [**NestedCable**](NestedCable.md) |  | [optional] 
**CablePeer** | Pointer to **map[string]string** |  Return the appropriate serializer for the cable termination model.  | [optional] [readonly] 
**CablePeerType** | Pointer to **string** |  | [optional] [readonly] 
**ConnectedEndpoint** | Pointer to **map[string]string** |  Return the appropriate serializer for the type of connected object.  | [optional] [readonly] 
**ConnectedEndpointType** | Pointer to **string** |  | [optional] [readonly] 
**ConnectedEndpointReachable** | Pointer to **bool** |  | [optional] [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**CountIpaddresses** | Pointer to **int32** |  | [optional] [readonly] 
**Occupied** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewInterface

`func NewInterface(device NestedDevice, name string, type_ Type2, ) *Interface`

NewInterface instantiates a new Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceWithDefaults

`func NewInterfaceWithDefaults() *Interface`

NewInterfaceWithDefaults instantiates a new Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Interface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Interface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Interface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Interface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Interface) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Interface) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Interface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Interface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Interface) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Interface) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *Interface) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Interface) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Interface) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *Interface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interface) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Interface) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Interface) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Interface) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Interface) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *Interface) GetType() Type2`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interface) GetTypeOk() (*Type2, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interface) SetType(v Type2)`

SetType sets Type field to given value.


### GetEnabled

`func (o *Interface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Interface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Interface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Interface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *Interface) GetParent() NestedInterface`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Interface) GetParentOk() (*NestedInterface, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Interface) SetParent(v NestedInterface)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Interface) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetLag

`func (o *Interface) GetLag() NestedInterface`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *Interface) GetLagOk() (*NestedInterface, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *Interface) SetLag(v NestedInterface)`

SetLag sets Lag field to given value.

### HasLag

`func (o *Interface) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetMtu

`func (o *Interface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Interface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Interface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *Interface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *Interface) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *Interface) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMacAddress

`func (o *Interface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Interface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Interface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Interface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *Interface) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *Interface) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetMgmtOnly

`func (o *Interface) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *Interface) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *Interface) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *Interface) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *Interface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Interface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Interface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Interface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *Interface) GetMode() Mode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Interface) GetModeOk() (*Mode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Interface) SetMode(v Mode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Interface) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetUntaggedVlan

`func (o *Interface) GetUntaggedVlan() NestedVLAN`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *Interface) GetUntaggedVlanOk() (*NestedVLAN, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *Interface) SetUntaggedVlan(v NestedVLAN)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *Interface) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *Interface) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *Interface) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *Interface) GetTaggedVlans() []NestedVLAN`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *Interface) GetTaggedVlansOk() (*[]NestedVLAN, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *Interface) SetTaggedVlans(v []NestedVLAN)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *Interface) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetMarkConnected

`func (o *Interface) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *Interface) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *Interface) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *Interface) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *Interface) GetCable() NestedCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *Interface) GetCableOk() (*NestedCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *Interface) SetCable(v NestedCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *Interface) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *Interface) GetCablePeer() map[string]string`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *Interface) GetCablePeerOk() (*map[string]string, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *Interface) SetCablePeer(v map[string]string)`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *Interface) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### GetCablePeerType

`func (o *Interface) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *Interface) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *Interface) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *Interface) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### GetConnectedEndpoint

`func (o *Interface) GetConnectedEndpoint() map[string]string`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *Interface) GetConnectedEndpointOk() (*map[string]string, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *Interface) SetConnectedEndpoint(v map[string]string)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.

### HasConnectedEndpoint

`func (o *Interface) HasConnectedEndpoint() bool`

HasConnectedEndpoint returns a boolean if a field has been set.

### GetConnectedEndpointType

`func (o *Interface) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *Interface) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *Interface) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.

### HasConnectedEndpointType

`func (o *Interface) HasConnectedEndpointType() bool`

HasConnectedEndpointType returns a boolean if a field has been set.

### GetConnectedEndpointReachable

`func (o *Interface) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *Interface) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *Interface) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.

### HasConnectedEndpointReachable

`func (o *Interface) HasConnectedEndpointReachable() bool`

HasConnectedEndpointReachable returns a boolean if a field has been set.

### GetTags

`func (o *Interface) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Interface) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Interface) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Interface) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Interface) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Interface) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Interface) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Interface) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Interface) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Interface) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Interface) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Interface) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Interface) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Interface) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Interface) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Interface) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCountIpaddresses

`func (o *Interface) GetCountIpaddresses() int32`

GetCountIpaddresses returns the CountIpaddresses field if non-nil, zero value otherwise.

### GetCountIpaddressesOk

`func (o *Interface) GetCountIpaddressesOk() (*int32, bool)`

GetCountIpaddressesOk returns a tuple with the CountIpaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountIpaddresses

`func (o *Interface) SetCountIpaddresses(v int32)`

SetCountIpaddresses sets CountIpaddresses field to given value.

### HasCountIpaddresses

`func (o *Interface) HasCountIpaddresses() bool`

HasCountIpaddresses returns a boolean if a field has been set.

### GetOccupied

`func (o *Interface) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *Interface) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *Interface) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.

### HasOccupied

`func (o *Interface) HasOccupied() bool`

HasOccupied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


