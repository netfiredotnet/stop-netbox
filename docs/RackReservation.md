# RackReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Rack** | [**NullableNestedRack**](NestedRack.md) |  | 
**Units** | **[]int32** |  | 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**User** | [**NestedUser**](NestedUser.md) |  | 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Description** | **string** |  | 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRackReservation

`func NewRackReservation(rack NullableNestedRack, units []int32, user NestedUser, description string, ) *RackReservation`

NewRackReservation instantiates a new RackReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackReservationWithDefaults

`func NewRackReservationWithDefaults() *RackReservation`

NewRackReservationWithDefaults instantiates a new RackReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RackReservation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RackReservation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RackReservation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RackReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *RackReservation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RackReservation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RackReservation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RackReservation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *RackReservation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RackReservation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RackReservation) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *RackReservation) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetRack

`func (o *RackReservation) GetRack() NestedRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *RackReservation) GetRackOk() (*NestedRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *RackReservation) SetRack(v NestedRack)`

SetRack sets Rack field to given value.


### SetRackNil

`func (o *RackReservation) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *RackReservation) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetUnits

`func (o *RackReservation) GetUnits() []int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *RackReservation) GetUnitsOk() (*[]int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *RackReservation) SetUnits(v []int32)`

SetUnits sets Units field to given value.


### GetCreated

`func (o *RackReservation) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RackReservation) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RackReservation) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RackReservation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUser

`func (o *RackReservation) GetUser() NestedUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RackReservation) GetUserOk() (*NestedUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RackReservation) SetUser(v NestedUser)`

SetUser sets User field to given value.


### GetTenant

`func (o *RackReservation) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RackReservation) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RackReservation) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *RackReservation) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *RackReservation) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *RackReservation) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *RackReservation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RackReservation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RackReservation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *RackReservation) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RackReservation) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RackReservation) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RackReservation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *RackReservation) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RackReservation) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RackReservation) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RackReservation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


