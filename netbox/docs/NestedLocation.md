# NestedLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**RackCount** | Pointer to **int32** |  | [optional] [readonly] 
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewNestedLocation

`func NewNestedLocation(name string, slug string, ) *NestedLocation`

NewNestedLocation instantiates a new NestedLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedLocationWithDefaults

`func NewNestedLocationWithDefaults() *NestedLocation`

NewNestedLocationWithDefaults instantiates a new NestedLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedLocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NestedLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *NestedLocation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedLocation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedLocation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NestedLocation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *NestedLocation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedLocation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedLocation) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *NestedLocation) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *NestedLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedLocation) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedLocation) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedLocation) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedLocation) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetRackCount

`func (o *NestedLocation) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *NestedLocation) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *NestedLocation) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.

### HasRackCount

`func (o *NestedLocation) HasRackCount() bool`

HasRackCount returns a boolean if a field has been set.

### GetDepth

`func (o *NestedLocation) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *NestedLocation) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *NestedLocation) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *NestedLocation) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


