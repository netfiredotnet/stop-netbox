/*
NetBox API

API to access NetBox

API version: 2.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// JobResult struct for JobResult
type JobResult struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Display *string `json:"display,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Completed NullableTime `json:"completed,omitempty"`
	Name string `json:"name"`
	ObjType *string `json:"obj_type,omitempty"`
	Status *Status6 `json:"status,omitempty"`
	User *NestedUser `json:"user,omitempty"`
	Data NullableString `json:"data,omitempty"`
	JobId string `json:"job_id"`
}

// NewJobResult instantiates a new JobResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResult(name string, jobId string) *JobResult {
	this := JobResult{}
	this.Name = name
	this.JobId = jobId
	return &this
}

// NewJobResultWithDefaults instantiates a new JobResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResultWithDefaults() *JobResult {
	this := JobResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JobResult) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JobResult) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *JobResult) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *JobResult) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *JobResult) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *JobResult) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *JobResult) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *JobResult) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *JobResult) SetDisplay(v string) {
	o.Display = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *JobResult) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *JobResult) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *JobResult) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetCompleted() time.Time {
	if o == nil || o.Completed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *JobResult) HasCompleted() bool {
	if o != nil && o.Completed.IsSet() {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given NullableTime and assigns it to the Completed field.
func (o *JobResult) SetCompleted(v time.Time) {
	o.Completed.Set(&v)
}
// SetCompletedNil sets the value for Completed to be an explicit nil
func (o *JobResult) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
func (o *JobResult) UnsetCompleted() {
	o.Completed.Unset()
}

// GetName returns the Name field value
func (o *JobResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobResult) SetName(v string) {
	o.Name = v
}

// GetObjType returns the ObjType field value if set, zero value otherwise.
func (o *JobResult) GetObjType() string {
	if o == nil || o.ObjType == nil {
		var ret string
		return ret
	}
	return *o.ObjType
}

// GetObjTypeOk returns a tuple with the ObjType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetObjTypeOk() (*string, bool) {
	if o == nil || o.ObjType == nil {
		return nil, false
	}
	return o.ObjType, true
}

// HasObjType returns a boolean if a field has been set.
func (o *JobResult) HasObjType() bool {
	if o != nil && o.ObjType != nil {
		return true
	}

	return false
}

// SetObjType gets a reference to the given string and assigns it to the ObjType field.
func (o *JobResult) SetObjType(v string) {
	o.ObjType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JobResult) GetStatus() Status6 {
	if o == nil || o.Status == nil {
		var ret Status6
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetStatusOk() (*Status6, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JobResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status6 and assigns it to the Status field.
func (o *JobResult) SetStatus(v Status6) {
	o.Status = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *JobResult) GetUser() NestedUser {
	if o == nil || o.User == nil {
		var ret NestedUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetUserOk() (*NestedUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *JobResult) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given NestedUser and assigns it to the User field.
func (o *JobResult) SetUser(v NestedUser) {
	o.User = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *JobResult) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *JobResult) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *JobResult) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *JobResult) UnsetData() {
	o.Data.Unset()
}

// GetJobId returns the JobId field value
func (o *JobResult) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *JobResult) SetJobId(v string) {
	o.JobId = v
}

func (o JobResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ObjType != nil {
		toSerialize["obj_type"] = o.ObjType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if true {
		toSerialize["job_id"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableJobResult struct {
	value *JobResult
	isSet bool
}

func (v NullableJobResult) Get() *JobResult {
	return v.value
}

func (v *NullableJobResult) Set(val *JobResult) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResult) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResult(val *JobResult) *NullableJobResult {
	return &NullableJobResult{value: val, isSet: true}
}

func (v NullableJobResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


