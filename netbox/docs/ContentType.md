# ContentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**AppLabel** | **string** |  | 
**Model** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewContentType

`func NewContentType(appLabel string, model string, ) *ContentType`

NewContentType instantiates a new ContentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTypeWithDefaults

`func NewContentTypeWithDefaults() *ContentType`

NewContentTypeWithDefaults instantiates a new ContentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContentType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ContentType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentType) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ContentType) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *ContentType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ContentType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ContentType) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ContentType) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetAppLabel

`func (o *ContentType) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *ContentType) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *ContentType) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.


### GetModel

`func (o *ContentType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ContentType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ContentType) SetModel(v string)`

SetModel sets Model field to given value.


### GetDisplayName

`func (o *ContentType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContentType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContentType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ContentType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


