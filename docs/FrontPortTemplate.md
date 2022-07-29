# FrontPortTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | [**NestedDeviceType**](NestedDeviceType.md) |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**Type1**](Type1.md) |  | 
**RearPort** | [**NestedRearPortTemplate**](NestedRearPortTemplate.md) |  | 
**RearPortPosition** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewFrontPortTemplate

`func NewFrontPortTemplate(deviceType NestedDeviceType, name string, type_ Type1, rearPort NestedRearPortTemplate, ) *FrontPortTemplate`

NewFrontPortTemplate instantiates a new FrontPortTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontPortTemplateWithDefaults

`func NewFrontPortTemplateWithDefaults() *FrontPortTemplate`

NewFrontPortTemplateWithDefaults instantiates a new FrontPortTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FrontPortTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FrontPortTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FrontPortTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FrontPortTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *FrontPortTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FrontPortTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FrontPortTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FrontPortTemplate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *FrontPortTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *FrontPortTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *FrontPortTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *FrontPortTemplate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDeviceType

`func (o *FrontPortTemplate) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *FrontPortTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *FrontPortTemplate) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetName

`func (o *FrontPortTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrontPortTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrontPortTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *FrontPortTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FrontPortTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FrontPortTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FrontPortTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *FrontPortTemplate) GetType() Type1`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FrontPortTemplate) GetTypeOk() (*Type1, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FrontPortTemplate) SetType(v Type1)`

SetType sets Type field to given value.


### GetRearPort

`func (o *FrontPortTemplate) GetRearPort() NestedRearPortTemplate`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *FrontPortTemplate) GetRearPortOk() (*NestedRearPortTemplate, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *FrontPortTemplate) SetRearPort(v NestedRearPortTemplate)`

SetRearPort sets RearPort field to given value.


### GetRearPortPosition

`func (o *FrontPortTemplate) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *FrontPortTemplate) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *FrontPortTemplate) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *FrontPortTemplate) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDescription

`func (o *FrontPortTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FrontPortTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FrontPortTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FrontPortTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *FrontPortTemplate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FrontPortTemplate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FrontPortTemplate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FrontPortTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *FrontPortTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FrontPortTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FrontPortTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *FrontPortTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


