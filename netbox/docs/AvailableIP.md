# AvailableIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **int32** |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] [readonly] 
**Vrf** | Pointer to [**NullableNestedVRF**](NestedVRF.md) |  | [optional] 

## Methods

### NewAvailableIP

`func NewAvailableIP() *AvailableIP`

NewAvailableIP instantiates a new AvailableIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableIPWithDefaults

`func NewAvailableIPWithDefaults() *AvailableIP`

NewAvailableIPWithDefaults instantiates a new AvailableIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *AvailableIP) GetFamily() int32`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AvailableIP) GetFamilyOk() (*int32, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AvailableIP) SetFamily(v int32)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AvailableIP) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetAddress

`func (o *AvailableIP) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AvailableIP) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AvailableIP) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AvailableIP) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVrf

`func (o *AvailableIP) GetVrf() NestedVRF`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *AvailableIP) GetVrfOk() (*NestedVRF, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *AvailableIP) SetVrf(v NestedVRF)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *AvailableIP) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *AvailableIP) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *AvailableIP) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


