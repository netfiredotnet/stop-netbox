# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**ContentTypes** | **[]string** |  | 
**Type** | [**Type7**](Type7.md) |  | 
**Name** | **string** | Internal field name | 
**Label** | Pointer to **string** | Name of the field as displayed to users (if not provided, the field&#39;s name will be used) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** | If true, this field is required when creating new objects or editing an existing object. | [optional] 
**FilterLogic** | Pointer to [**FilterLogic**](FilterLogic.md) |  | [optional] 
**Default** | Pointer to **NullableString** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**Weight** | Pointer to **int32** | Fields with higher weights appear lower in a form. | [optional] 
**ValidationMinimum** | Pointer to **NullableInt32** | Minimum allowed value (for numeric fields) | [optional] 
**ValidationMaximum** | Pointer to **NullableInt32** | Maximum allowed value (for numeric fields) | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] 
**Choices** | Pointer to **[]string** | Comma-separated list of available choices (for selection fields) | [optional] 

## Methods

### NewCustomField

`func NewCustomField(contentTypes []string, type_ Type7, name string, ) *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *CustomField) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomField) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomField) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CustomField) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *CustomField) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CustomField) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CustomField) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CustomField) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetContentTypes

`func (o *CustomField) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *CustomField) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *CustomField) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetType

`func (o *CustomField) GetType() Type7`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomField) GetTypeOk() (*Type7, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomField) SetType(v Type7)`

SetType sets Type field to given value.


### GetName

`func (o *CustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomField) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CustomField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *CustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *CustomField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CustomField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetFilterLogic

`func (o *CustomField) GetFilterLogic() FilterLogic`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *CustomField) GetFilterLogicOk() (*FilterLogic, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *CustomField) SetFilterLogic(v FilterLogic)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *CustomField) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetDefault

`func (o *CustomField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *CustomField) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *CustomField) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetWeight

`func (o *CustomField) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomField) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomField) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CustomField) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *CustomField) GetValidationMinimum() int32`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *CustomField) GetValidationMinimumOk() (*int32, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *CustomField) SetValidationMinimum(v int32)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *CustomField) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *CustomField) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *CustomField) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *CustomField) GetValidationMaximum() int32`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *CustomField) GetValidationMaximumOk() (*int32, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *CustomField) SetValidationMaximum(v int32)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *CustomField) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *CustomField) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *CustomField) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *CustomField) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *CustomField) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *CustomField) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *CustomField) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetChoices

`func (o *CustomField) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *CustomField) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *CustomField) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *CustomField) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *CustomField) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *CustomField) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


