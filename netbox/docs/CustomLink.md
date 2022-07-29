# CustomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**ContentType** | **string** |  | 
**Name** | **string** |  | 
**LinkText** | **string** | Jinja2 template code for link text | 
**LinkUrl** | **string** | Jinja2 template code for link URL | 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to **string** | The class of the first link in a group will be used for the dropdown button | [optional] 
**NewWindow** | Pointer to **bool** | Force link to open in a new window | [optional] 

## Methods

### NewCustomLink

`func NewCustomLink(contentType string, name string, linkText string, linkUrl string, ) *CustomLink`

NewCustomLink instantiates a new CustomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLinkWithDefaults

`func NewCustomLinkWithDefaults() *CustomLink`

NewCustomLinkWithDefaults instantiates a new CustomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomLink) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *CustomLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CustomLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *CustomLink) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CustomLink) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CustomLink) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CustomLink) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetContentType

`func (o *CustomLink) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CustomLink) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CustomLink) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *CustomLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomLink) SetName(v string)`

SetName sets Name field to given value.


### GetLinkText

`func (o *CustomLink) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *CustomLink) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *CustomLink) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.


### GetLinkUrl

`func (o *CustomLink) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *CustomLink) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *CustomLink) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.


### GetWeight

`func (o *CustomLink) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomLink) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomLink) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CustomLink) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *CustomLink) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CustomLink) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CustomLink) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CustomLink) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *CustomLink) GetButtonClass() string`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *CustomLink) GetButtonClassOk() (*string, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *CustomLink) SetButtonClass(v string)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *CustomLink) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetNewWindow

`func (o *CustomLink) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *CustomLink) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *CustomLink) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.

### HasNewWindow

`func (o *CustomLink) HasNewWindow() bool`

HasNewWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


