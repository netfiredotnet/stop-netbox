# InventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Device** | [**NestedDevice**](NestedDevice.md) |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Manufacturer** | Pointer to [**NestedManufacturer**](NestedManufacturer.md) |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this item | [optional] 
**Discovered** | Pointer to **bool** | This item was automatically discovered | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewInventoryItem

`func NewInventoryItem(device NestedDevice, name string, ) *InventoryItem`

NewInventoryItem instantiates a new InventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemWithDefaults

`func NewInventoryItemWithDefaults() *InventoryItem`

NewInventoryItemWithDefaults instantiates a new InventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *InventoryItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InventoryItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InventoryItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InventoryItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *InventoryItem) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InventoryItem) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InventoryItem) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *InventoryItem) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDevice

`func (o *InventoryItem) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InventoryItem) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InventoryItem) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.


### GetParent

`func (o *InventoryItem) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *InventoryItem) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *InventoryItem) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *InventoryItem) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *InventoryItem) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *InventoryItem) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *InventoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryItem) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *InventoryItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InventoryItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InventoryItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InventoryItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetManufacturer

`func (o *InventoryItem) GetManufacturer() NestedManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InventoryItem) GetManufacturerOk() (*NestedManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InventoryItem) SetManufacturer(v NestedManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InventoryItem) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetPartId

`func (o *InventoryItem) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *InventoryItem) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *InventoryItem) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *InventoryItem) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetSerial

`func (o *InventoryItem) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InventoryItem) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InventoryItem) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InventoryItem) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *InventoryItem) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *InventoryItem) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *InventoryItem) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *InventoryItem) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *InventoryItem) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *InventoryItem) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDiscovered

`func (o *InventoryItem) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *InventoryItem) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *InventoryItem) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *InventoryItem) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDescription

`func (o *InventoryItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InventoryItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *InventoryItem) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryItem) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryItem) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InventoryItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *InventoryItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InventoryItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InventoryItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InventoryItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *InventoryItem) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryItem) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InventoryItem) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InventoryItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InventoryItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InventoryItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InventoryItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InventoryItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDepth

`func (o *InventoryItem) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *InventoryItem) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *InventoryItem) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *InventoryItem) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


