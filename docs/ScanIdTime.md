# ScanIdTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Scan ID | [optional] 
**Time** | Pointer to **time.Time** | Date when the scan was performed | [optional] 

## Methods

### NewScanIdTime

`func NewScanIdTime() *ScanIdTime`

NewScanIdTime instantiates a new ScanIdTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanIdTimeWithDefaults

`func NewScanIdTimeWithDefaults() *ScanIdTime`

NewScanIdTimeWithDefaults instantiates a new ScanIdTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScanIdTime) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScanIdTime) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScanIdTime) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ScanIdTime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *ScanIdTime) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScanIdTime) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ScanIdTime) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ScanIdTime) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


