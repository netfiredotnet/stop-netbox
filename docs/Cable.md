# Cable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**TerminationAType** | **string** |  | 
**TerminationAId** | **int32** |  | 
**TerminationA** | Pointer to **map[string]string** |  | [optional] [readonly] 
**TerminationBType** | **string** |  | 
**TerminationBId** | **int32** |  | 
**TerminationB** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**Status1**](Status1.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**LengthUnit** | Pointer to [**LengthUnit**](LengthUnit.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCable

`func NewCable(terminationAType string, terminationAId int32, terminationBType string, terminationBId int32, ) *Cable`

NewCable instantiates a new Cable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableWithDefaults

`func NewCableWithDefaults() *Cable`

NewCableWithDefaults instantiates a new Cable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cable) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cable) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cable) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Cable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Cable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Cable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Cable) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Cable) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Cable) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Cable) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Cable) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Cable) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTerminationAType

`func (o *Cable) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *Cable) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *Cable) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.


### GetTerminationAId

`func (o *Cable) GetTerminationAId() int32`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *Cable) GetTerminationAIdOk() (*int32, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *Cable) SetTerminationAId(v int32)`

SetTerminationAId sets TerminationAId field to given value.


### GetTerminationA

`func (o *Cable) GetTerminationA() map[string]string`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *Cable) GetTerminationAOk() (*map[string]string, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *Cable) SetTerminationA(v map[string]string)`

SetTerminationA sets TerminationA field to given value.

### HasTerminationA

`func (o *Cable) HasTerminationA() bool`

HasTerminationA returns a boolean if a field has been set.

### GetTerminationBType

`func (o *Cable) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *Cable) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *Cable) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.


### GetTerminationBId

`func (o *Cable) GetTerminationBId() int32`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *Cable) GetTerminationBIdOk() (*int32, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *Cable) SetTerminationBId(v int32)`

SetTerminationBId sets TerminationBId field to given value.


### GetTerminationB

`func (o *Cable) GetTerminationB() map[string]string`

GetTerminationB returns the TerminationB field if non-nil, zero value otherwise.

### GetTerminationBOk

`func (o *Cable) GetTerminationBOk() (*map[string]string, bool)`

GetTerminationBOk returns a tuple with the TerminationB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationB

`func (o *Cable) SetTerminationB(v map[string]string)`

SetTerminationB sets TerminationB field to given value.

### HasTerminationB

`func (o *Cable) HasTerminationB() bool`

HasTerminationB returns a boolean if a field has been set.

### GetType

`func (o *Cable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Cable) GetStatus() Status1`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cable) GetStatusOk() (*Status1, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cable) SetStatus(v Status1)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLabel

`func (o *Cable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Cable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Cable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Cable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *Cable) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Cable) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Cable) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Cable) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *Cable) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Cable) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Cable) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Cable) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *Cable) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *Cable) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *Cable) GetLengthUnit() LengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *Cable) GetLengthUnitOk() (*LengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *Cable) SetLengthUnit(v LengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *Cable) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetTags

`func (o *Cable) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Cable) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Cable) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Cable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Cable) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Cable) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Cable) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Cable) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


