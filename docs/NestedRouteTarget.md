# NestedRouteTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | Route target value (formatted in accordance with RFC 4360) | 

## Methods

### NewNestedRouteTarget

`func NewNestedRouteTarget(name string, ) *NestedRouteTarget`

NewNestedRouteTarget instantiates a new NestedRouteTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedRouteTargetWithDefaults

`func NewNestedRouteTargetWithDefaults() *NestedRouteTarget`

NewNestedRouteTargetWithDefaults instantiates a new NestedRouteTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedRouteTarget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedRouteTarget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedRouteTarget) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NestedRouteTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *NestedRouteTarget) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedRouteTarget) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedRouteTarget) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NestedRouteTarget) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *NestedRouteTarget) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedRouteTarget) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedRouteTarget) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *NestedRouteTarget) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *NestedRouteTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedRouteTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedRouteTarget) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


