# NestedVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Master** | [**NestedDevice**](NestedDevice.md) |  | 
**MemberCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewNestedVirtualChassis

`func NewNestedVirtualChassis(name string, master NestedDevice, ) *NestedVirtualChassis`

NewNestedVirtualChassis instantiates a new NestedVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedVirtualChassisWithDefaults

`func NewNestedVirtualChassisWithDefaults() *NestedVirtualChassis`

NewNestedVirtualChassisWithDefaults instantiates a new NestedVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedVirtualChassis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedVirtualChassis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedVirtualChassis) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NestedVirtualChassis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NestedVirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedVirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedVirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *NestedVirtualChassis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedVirtualChassis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedVirtualChassis) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NestedVirtualChassis) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMaster

`func (o *NestedVirtualChassis) GetMaster() NestedDevice`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *NestedVirtualChassis) GetMasterOk() (*NestedDevice, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *NestedVirtualChassis) SetMaster(v NestedDevice)`

SetMaster sets Master field to given value.


### GetMemberCount

`func (o *NestedVirtualChassis) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *NestedVirtualChassis) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *NestedVirtualChassis) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *NestedVirtualChassis) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


