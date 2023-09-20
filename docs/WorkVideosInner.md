# WorkVideosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**WebResource** | Pointer to **string** |  | [optional] 
**DurationSecs** | Pointer to **float32** |  | [optional] 
**Subtitles** | Pointer to **string** |  | [optional] 
**MasterMetadata** | Pointer to [**VideoMetadata**](VideoMetadata.md) |  | [optional] 
**AccessMetadata** | Pointer to [**VideoMetadata**](VideoMetadata.md) |  | [optional] 
**WebMetadata** | Pointer to [**VideoMetadata**](VideoMetadata.md) |  | [optional] 
**AccessLevel** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to **string** |  | [optional] 
**AnimatedSnapshot** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkVideosInner

`func NewWorkVideosInner() *WorkVideosInner`

NewWorkVideosInner instantiates a new WorkVideosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkVideosInnerWithDefaults

`func NewWorkVideosInnerWithDefaults() *WorkVideosInner`

NewWorkVideosInnerWithDefaults instantiates a new WorkVideosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkVideosInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkVideosInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkVideosInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkVideosInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *WorkVideosInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkVideosInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkVideosInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkVideosInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetResource

`func (o *WorkVideosInner) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WorkVideosInner) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WorkVideosInner) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *WorkVideosInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetWebResource

`func (o *WorkVideosInner) GetWebResource() string`

GetWebResource returns the WebResource field if non-nil, zero value otherwise.

### GetWebResourceOk

`func (o *WorkVideosInner) GetWebResourceOk() (*string, bool)`

GetWebResourceOk returns a tuple with the WebResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebResource

`func (o *WorkVideosInner) SetWebResource(v string)`

SetWebResource sets WebResource field to given value.

### HasWebResource

`func (o *WorkVideosInner) HasWebResource() bool`

HasWebResource returns a boolean if a field has been set.

### GetDurationSecs

`func (o *WorkVideosInner) GetDurationSecs() float32`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *WorkVideosInner) GetDurationSecsOk() (*float32, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *WorkVideosInner) SetDurationSecs(v float32)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *WorkVideosInner) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### GetSubtitles

`func (o *WorkVideosInner) GetSubtitles() string`

GetSubtitles returns the Subtitles field if non-nil, zero value otherwise.

### GetSubtitlesOk

`func (o *WorkVideosInner) GetSubtitlesOk() (*string, bool)`

GetSubtitlesOk returns a tuple with the Subtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitles

`func (o *WorkVideosInner) SetSubtitles(v string)`

SetSubtitles sets Subtitles field to given value.

### HasSubtitles

`func (o *WorkVideosInner) HasSubtitles() bool`

HasSubtitles returns a boolean if a field has been set.

### GetMasterMetadata

`func (o *WorkVideosInner) GetMasterMetadata() VideoMetadata`

GetMasterMetadata returns the MasterMetadata field if non-nil, zero value otherwise.

### GetMasterMetadataOk

`func (o *WorkVideosInner) GetMasterMetadataOk() (*VideoMetadata, bool)`

GetMasterMetadataOk returns a tuple with the MasterMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMetadata

`func (o *WorkVideosInner) SetMasterMetadata(v VideoMetadata)`

SetMasterMetadata sets MasterMetadata field to given value.

### HasMasterMetadata

`func (o *WorkVideosInner) HasMasterMetadata() bool`

HasMasterMetadata returns a boolean if a field has been set.

### GetAccessMetadata

`func (o *WorkVideosInner) GetAccessMetadata() VideoMetadata`

GetAccessMetadata returns the AccessMetadata field if non-nil, zero value otherwise.

### GetAccessMetadataOk

`func (o *WorkVideosInner) GetAccessMetadataOk() (*VideoMetadata, bool)`

GetAccessMetadataOk returns a tuple with the AccessMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMetadata

`func (o *WorkVideosInner) SetAccessMetadata(v VideoMetadata)`

SetAccessMetadata sets AccessMetadata field to given value.

### HasAccessMetadata

`func (o *WorkVideosInner) HasAccessMetadata() bool`

HasAccessMetadata returns a boolean if a field has been set.

### GetWebMetadata

`func (o *WorkVideosInner) GetWebMetadata() VideoMetadata`

GetWebMetadata returns the WebMetadata field if non-nil, zero value otherwise.

### GetWebMetadataOk

`func (o *WorkVideosInner) GetWebMetadataOk() (*VideoMetadata, bool)`

GetWebMetadataOk returns a tuple with the WebMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMetadata

`func (o *WorkVideosInner) SetWebMetadata(v VideoMetadata)`

SetWebMetadata sets WebMetadata field to given value.

### HasWebMetadata

`func (o *WorkVideosInner) HasWebMetadata() bool`

HasWebMetadata returns a boolean if a field has been set.

### GetAccessLevel

`func (o *WorkVideosInner) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *WorkVideosInner) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *WorkVideosInner) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *WorkVideosInner) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetSnapshot

`func (o *WorkVideosInner) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *WorkVideosInner) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *WorkVideosInner) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *WorkVideosInner) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetAnimatedSnapshot

`func (o *WorkVideosInner) GetAnimatedSnapshot() string`

GetAnimatedSnapshot returns the AnimatedSnapshot field if non-nil, zero value otherwise.

### GetAnimatedSnapshotOk

`func (o *WorkVideosInner) GetAnimatedSnapshotOk() (*string, bool)`

GetAnimatedSnapshotOk returns a tuple with the AnimatedSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimatedSnapshot

`func (o *WorkVideosInner) SetAnimatedSnapshot(v string)`

SetAnimatedSnapshot sets AnimatedSnapshot field to given value.

### HasAnimatedSnapshot

`func (o *WorkVideosInner) HasAnimatedSnapshot() bool`

HasAnimatedSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


