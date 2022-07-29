# WritableCable

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
**Status** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**LengthUnit** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableCable

`func NewWritableCable(terminationAType string, terminationAId int32, terminationBType string, terminationBId int32, ) *WritableCable`

NewWritableCable instantiates a new WritableCable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCableWithDefaults

`func NewWritableCableWithDefaults() *WritableCable`

NewWritableCableWithDefaults instantiates a new WritableCable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCable) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCable) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCable) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WritableCable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *WritableCable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableCable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableCable) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WritableCable) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableCable) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableCable) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableCable) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *WritableCable) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTerminationAType

`func (o *WritableCable) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *WritableCable) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *WritableCable) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.


### GetTerminationAId

`func (o *WritableCable) GetTerminationAId() int32`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *WritableCable) GetTerminationAIdOk() (*int32, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *WritableCable) SetTerminationAId(v int32)`

SetTerminationAId sets TerminationAId field to given value.


### GetTerminationA

`func (o *WritableCable) GetTerminationA() map[string]string`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *WritableCable) GetTerminationAOk() (*map[string]string, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *WritableCable) SetTerminationA(v map[string]string)`

SetTerminationA sets TerminationA field to given value.

### HasTerminationA

`func (o *WritableCable) HasTerminationA() bool`

HasTerminationA returns a boolean if a field has been set.

### GetTerminationBType

`func (o *WritableCable) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *WritableCable) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *WritableCable) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.


### GetTerminationBId

`func (o *WritableCable) GetTerminationBId() int32`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *WritableCable) GetTerminationBIdOk() (*int32, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *WritableCable) SetTerminationBId(v int32)`

SetTerminationBId sets TerminationBId field to given value.


### GetTerminationB

`func (o *WritableCable) GetTerminationB() map[string]string`

GetTerminationB returns the TerminationB field if non-nil, zero value otherwise.

### GetTerminationBOk

`func (o *WritableCable) GetTerminationBOk() (*map[string]string, bool)`

GetTerminationBOk returns a tuple with the TerminationB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationB

`func (o *WritableCable) SetTerminationB(v map[string]string)`

SetTerminationB sets TerminationB field to given value.

### HasTerminationB

`func (o *WritableCable) HasTerminationB() bool`

HasTerminationB returns a boolean if a field has been set.

### GetType

`func (o *WritableCable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WritableCable) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableCable) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableCable) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableCable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLabel

`func (o *WritableCable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableCable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableCable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableCable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *WritableCable) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WritableCable) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WritableCable) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WritableCable) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *WritableCable) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WritableCable) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WritableCable) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *WritableCable) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *WritableCable) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *WritableCable) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *WritableCable) GetLengthUnit() string`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *WritableCable) GetLengthUnitOk() (*string, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *WritableCable) SetLengthUnit(v string)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *WritableCable) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetTags

`func (o *WritableCable) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCable) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCable) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCable) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCable) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCable) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCable) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


