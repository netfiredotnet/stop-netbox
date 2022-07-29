# WritableVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Domain** | Pointer to **string** |  | [optional] 
**Master** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewWritableVirtualChassis

`func NewWritableVirtualChassis(name string, ) *WritableVirtualChassis`

NewWritableVirtualChassis instantiates a new WritableVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVirtualChassisWithDefaults

`func NewWritableVirtualChassisWithDefaults() *WritableVirtualChassis`

NewWritableVirtualChassisWithDefaults instantiates a new WritableVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableVirtualChassis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableVirtualChassis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableVirtualChassis) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableVirtualChassis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableVirtualChassis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableVirtualChassis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableVirtualChassis) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableVirtualChassis) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableVirtualChassis) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableVirtualChassis) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableVirtualChassis) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableVirtualChassis) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *WritableVirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableVirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableVirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *WritableVirtualChassis) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WritableVirtualChassis) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WritableVirtualChassis) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WritableVirtualChassis) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMaster

`func (o *WritableVirtualChassis) GetMaster() int32`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *WritableVirtualChassis) GetMasterOk() (*int32, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *WritableVirtualChassis) SetMaster(v int32)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *WritableVirtualChassis) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *WritableVirtualChassis) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *WritableVirtualChassis) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetTags

`func (o *WritableVirtualChassis) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVirtualChassis) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVirtualChassis) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVirtualChassis) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVirtualChassis) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVirtualChassis) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVirtualChassis) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVirtualChassis) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMemberCount

`func (o *WritableVirtualChassis) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *WritableVirtualChassis) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *WritableVirtualChassis) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *WritableVirtualChassis) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


