# ImageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**ContentType** | **string** |  | 
**ObjectId** | **int32** |  | 
**Parent** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] [readonly] 
**ImageHeight** | **int32** |  | 
**ImageWidth** | **int32** |  | 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewImageAttachment

`func NewImageAttachment(contentType string, objectId int32, imageHeight int32, imageWidth int32, ) *ImageAttachment`

NewImageAttachment instantiates a new ImageAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAttachmentWithDefaults

`func NewImageAttachmentWithDefaults() *ImageAttachment`

NewImageAttachmentWithDefaults instantiates a new ImageAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageAttachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageAttachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageAttachment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImageAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ImageAttachment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageAttachment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageAttachment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImageAttachment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *ImageAttachment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ImageAttachment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ImageAttachment) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ImageAttachment) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetContentType

`func (o *ImageAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ImageAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ImageAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *ImageAttachment) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ImageAttachment) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ImageAttachment) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.


### GetParent

`func (o *ImageAttachment) GetParent() map[string]string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ImageAttachment) GetParentOk() (*map[string]string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ImageAttachment) SetParent(v map[string]string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ImageAttachment) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetName

`func (o *ImageAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *ImageAttachment) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageAttachment) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageAttachment) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ImageAttachment) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageHeight

`func (o *ImageAttachment) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *ImageAttachment) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *ImageAttachment) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.


### GetImageWidth

`func (o *ImageAttachment) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *ImageAttachment) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *ImageAttachment) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.


### GetCreated

`func (o *ImageAttachment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImageAttachment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImageAttachment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ImageAttachment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


