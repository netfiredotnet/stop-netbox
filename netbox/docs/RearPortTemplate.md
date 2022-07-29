# RearPortTemplate

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
**Positions** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewRearPortTemplate

`func NewRearPortTemplate(deviceType NestedDeviceType, name string, type_ Type1, ) *RearPortTemplate`

NewRearPortTemplate instantiates a new RearPortTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRearPortTemplateWithDefaults

`func NewRearPortTemplateWithDefaults() *RearPortTemplate`

NewRearPortTemplateWithDefaults instantiates a new RearPortTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RearPortTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RearPortTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RearPortTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RearPortTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *RearPortTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RearPortTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RearPortTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RearPortTemplate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *RearPortTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RearPortTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RearPortTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *RearPortTemplate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDeviceType

`func (o *RearPortTemplate) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *RearPortTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *RearPortTemplate) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetName

`func (o *RearPortTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RearPortTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RearPortTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *RearPortTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RearPortTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RearPortTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RearPortTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *RearPortTemplate) GetType() Type1`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RearPortTemplate) GetTypeOk() (*Type1, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RearPortTemplate) SetType(v Type1)`

SetType sets Type field to given value.


### GetPositions

`func (o *RearPortTemplate) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *RearPortTemplate) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *RearPortTemplate) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *RearPortTemplate) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetDescription

`func (o *RearPortTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RearPortTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RearPortTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RearPortTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *RearPortTemplate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RearPortTemplate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RearPortTemplate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RearPortTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RearPortTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RearPortTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RearPortTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RearPortTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


