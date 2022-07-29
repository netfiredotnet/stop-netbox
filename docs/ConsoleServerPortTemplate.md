# ConsoleServerPortTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | [**NestedDeviceType**](NestedDeviceType.md) |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**Type**](Type.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewConsoleServerPortTemplate

`func NewConsoleServerPortTemplate(deviceType NestedDeviceType, name string, ) *ConsoleServerPortTemplate`

NewConsoleServerPortTemplate instantiates a new ConsoleServerPortTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleServerPortTemplateWithDefaults

`func NewConsoleServerPortTemplateWithDefaults() *ConsoleServerPortTemplate`

NewConsoleServerPortTemplateWithDefaults instantiates a new ConsoleServerPortTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleServerPortTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleServerPortTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleServerPortTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleServerPortTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ConsoleServerPortTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConsoleServerPortTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConsoleServerPortTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConsoleServerPortTemplate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *ConsoleServerPortTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConsoleServerPortTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConsoleServerPortTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConsoleServerPortTemplate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDeviceType

`func (o *ConsoleServerPortTemplate) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ConsoleServerPortTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ConsoleServerPortTemplate) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetName

`func (o *ConsoleServerPortTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleServerPortTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleServerPortTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConsoleServerPortTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsoleServerPortTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsoleServerPortTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsoleServerPortTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *ConsoleServerPortTemplate) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleServerPortTemplate) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleServerPortTemplate) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ConsoleServerPortTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleServerPortTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleServerPortTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleServerPortTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleServerPortTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *ConsoleServerPortTemplate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsoleServerPortTemplate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsoleServerPortTemplate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConsoleServerPortTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ConsoleServerPortTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConsoleServerPortTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConsoleServerPortTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ConsoleServerPortTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


