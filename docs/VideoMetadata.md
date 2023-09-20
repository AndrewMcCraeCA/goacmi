# VideoMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**Filetype** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**VernonId** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**DurationHms** | Pointer to **string** |  | [optional] 
**DurationSecs** | Pointer to **float32** |  | [optional] 
**AudioBitRate** | Pointer to **int32** |  | [optional] 
**AudioChannels** | Pointer to **int32** |  | [optional] 
**VideoBitRate** | Pointer to **int32** |  | [optional] 
**FileSizeBytes** | Pointer to **int32** |  | [optional] 
**OverallBitRate** | Pointer to **int32** |  | [optional] 
**VideoFrameRate** | Pointer to **float32** |  | [optional] 
**AudioSampleRate** | Pointer to **int32** |  | [optional] 
**CreationDatetime** | Pointer to **string** |  | [optional] 
**AudioMaxBitRate** | Pointer to **int32** |  | [optional] 
**VideoMaxBitRate** | Pointer to **int32** |  | [optional] 

## Methods

### NewVideoMetadata

`func NewVideoMetadata() *VideoMetadata`

NewVideoMetadata instantiates a new VideoMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoMetadataWithDefaults

`func NewVideoMetadataWithDefaults() *VideoMetadata`

NewVideoMetadataWithDefaults instantiates a new VideoMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *VideoMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VideoMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VideoMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VideoMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWidth

`func (o *VideoMetadata) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoMetadata) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoMetadata) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideoMetadata) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VideoMetadata) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoMetadata) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoMetadata) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideoMetadata) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetChecksum

`func (o *VideoMetadata) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *VideoMetadata) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *VideoMetadata) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *VideoMetadata) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetFiletype

`func (o *VideoMetadata) GetFiletype() string`

GetFiletype returns the Filetype field if non-nil, zero value otherwise.

### GetFiletypeOk

`func (o *VideoMetadata) GetFiletypeOk() (*string, bool)`

GetFiletypeOk returns a tuple with the Filetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiletype

`func (o *VideoMetadata) SetFiletype(v string)`

SetFiletype sets Filetype field to given value.

### HasFiletype

`func (o *VideoMetadata) HasFiletype() bool`

HasFiletype returns a boolean if a field has been set.

### GetMimeType

`func (o *VideoMetadata) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *VideoMetadata) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *VideoMetadata) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *VideoMetadata) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetVernonId

`func (o *VideoMetadata) GetVernonId() string`

GetVernonId returns the VernonId field if non-nil, zero value otherwise.

### GetVernonIdOk

`func (o *VideoMetadata) GetVernonIdOk() (*string, bool)`

GetVernonIdOk returns a tuple with the VernonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVernonId

`func (o *VideoMetadata) SetVernonId(v string)`

SetVernonId sets VernonId field to given value.

### HasVernonId

`func (o *VideoMetadata) HasVernonId() bool`

HasVernonId returns a boolean if a field has been set.

### GetAudioCodec

`func (o *VideoMetadata) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *VideoMetadata) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *VideoMetadata) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *VideoMetadata) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoCodec

`func (o *VideoMetadata) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *VideoMetadata) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *VideoMetadata) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *VideoMetadata) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetDurationHms

`func (o *VideoMetadata) GetDurationHms() string`

GetDurationHms returns the DurationHms field if non-nil, zero value otherwise.

### GetDurationHmsOk

`func (o *VideoMetadata) GetDurationHmsOk() (*string, bool)`

GetDurationHmsOk returns a tuple with the DurationHms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationHms

`func (o *VideoMetadata) SetDurationHms(v string)`

SetDurationHms sets DurationHms field to given value.

### HasDurationHms

`func (o *VideoMetadata) HasDurationHms() bool`

HasDurationHms returns a boolean if a field has been set.

### GetDurationSecs

`func (o *VideoMetadata) GetDurationSecs() float32`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *VideoMetadata) GetDurationSecsOk() (*float32, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *VideoMetadata) SetDurationSecs(v float32)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *VideoMetadata) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### GetAudioBitRate

`func (o *VideoMetadata) GetAudioBitRate() int32`

GetAudioBitRate returns the AudioBitRate field if non-nil, zero value otherwise.

### GetAudioBitRateOk

`func (o *VideoMetadata) GetAudioBitRateOk() (*int32, bool)`

GetAudioBitRateOk returns a tuple with the AudioBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitRate

`func (o *VideoMetadata) SetAudioBitRate(v int32)`

SetAudioBitRate sets AudioBitRate field to given value.

### HasAudioBitRate

`func (o *VideoMetadata) HasAudioBitRate() bool`

HasAudioBitRate returns a boolean if a field has been set.

### GetAudioChannels

`func (o *VideoMetadata) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *VideoMetadata) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *VideoMetadata) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *VideoMetadata) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### GetVideoBitRate

`func (o *VideoMetadata) GetVideoBitRate() int32`

GetVideoBitRate returns the VideoBitRate field if non-nil, zero value otherwise.

### GetVideoBitRateOk

`func (o *VideoMetadata) GetVideoBitRateOk() (*int32, bool)`

GetVideoBitRateOk returns a tuple with the VideoBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBitRate

`func (o *VideoMetadata) SetVideoBitRate(v int32)`

SetVideoBitRate sets VideoBitRate field to given value.

### HasVideoBitRate

`func (o *VideoMetadata) HasVideoBitRate() bool`

HasVideoBitRate returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *VideoMetadata) GetFileSizeBytes() int32`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *VideoMetadata) GetFileSizeBytesOk() (*int32, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *VideoMetadata) SetFileSizeBytes(v int32)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *VideoMetadata) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetOverallBitRate

`func (o *VideoMetadata) GetOverallBitRate() int32`

GetOverallBitRate returns the OverallBitRate field if non-nil, zero value otherwise.

### GetOverallBitRateOk

`func (o *VideoMetadata) GetOverallBitRateOk() (*int32, bool)`

GetOverallBitRateOk returns a tuple with the OverallBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallBitRate

`func (o *VideoMetadata) SetOverallBitRate(v int32)`

SetOverallBitRate sets OverallBitRate field to given value.

### HasOverallBitRate

`func (o *VideoMetadata) HasOverallBitRate() bool`

HasOverallBitRate returns a boolean if a field has been set.

### GetVideoFrameRate

`func (o *VideoMetadata) GetVideoFrameRate() float32`

GetVideoFrameRate returns the VideoFrameRate field if non-nil, zero value otherwise.

### GetVideoFrameRateOk

`func (o *VideoMetadata) GetVideoFrameRateOk() (*float32, bool)`

GetVideoFrameRateOk returns a tuple with the VideoFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFrameRate

`func (o *VideoMetadata) SetVideoFrameRate(v float32)`

SetVideoFrameRate sets VideoFrameRate field to given value.

### HasVideoFrameRate

`func (o *VideoMetadata) HasVideoFrameRate() bool`

HasVideoFrameRate returns a boolean if a field has been set.

### GetAudioSampleRate

`func (o *VideoMetadata) GetAudioSampleRate() int32`

GetAudioSampleRate returns the AudioSampleRate field if non-nil, zero value otherwise.

### GetAudioSampleRateOk

`func (o *VideoMetadata) GetAudioSampleRateOk() (*int32, bool)`

GetAudioSampleRateOk returns a tuple with the AudioSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSampleRate

`func (o *VideoMetadata) SetAudioSampleRate(v int32)`

SetAudioSampleRate sets AudioSampleRate field to given value.

### HasAudioSampleRate

`func (o *VideoMetadata) HasAudioSampleRate() bool`

HasAudioSampleRate returns a boolean if a field has been set.

### GetCreationDatetime

`func (o *VideoMetadata) GetCreationDatetime() string`

GetCreationDatetime returns the CreationDatetime field if non-nil, zero value otherwise.

### GetCreationDatetimeOk

`func (o *VideoMetadata) GetCreationDatetimeOk() (*string, bool)`

GetCreationDatetimeOk returns a tuple with the CreationDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDatetime

`func (o *VideoMetadata) SetCreationDatetime(v string)`

SetCreationDatetime sets CreationDatetime field to given value.

### HasCreationDatetime

`func (o *VideoMetadata) HasCreationDatetime() bool`

HasCreationDatetime returns a boolean if a field has been set.

### GetAudioMaxBitRate

`func (o *VideoMetadata) GetAudioMaxBitRate() int32`

GetAudioMaxBitRate returns the AudioMaxBitRate field if non-nil, zero value otherwise.

### GetAudioMaxBitRateOk

`func (o *VideoMetadata) GetAudioMaxBitRateOk() (*int32, bool)`

GetAudioMaxBitRateOk returns a tuple with the AudioMaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioMaxBitRate

`func (o *VideoMetadata) SetAudioMaxBitRate(v int32)`

SetAudioMaxBitRate sets AudioMaxBitRate field to given value.

### HasAudioMaxBitRate

`func (o *VideoMetadata) HasAudioMaxBitRate() bool`

HasAudioMaxBitRate returns a boolean if a field has been set.

### GetVideoMaxBitRate

`func (o *VideoMetadata) GetVideoMaxBitRate() int32`

GetVideoMaxBitRate returns the VideoMaxBitRate field if non-nil, zero value otherwise.

### GetVideoMaxBitRateOk

`func (o *VideoMetadata) GetVideoMaxBitRateOk() (*int32, bool)`

GetVideoMaxBitRateOk returns a tuple with the VideoMaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoMaxBitRate

`func (o *VideoMetadata) SetVideoMaxBitRate(v int32)`

SetVideoMaxBitRate sets VideoMaxBitRate field to given value.

### HasVideoMaxBitRate

`func (o *VideoMetadata) HasVideoMaxBitRate() bool`

HasVideoMaxBitRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


