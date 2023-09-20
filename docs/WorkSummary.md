# WorkSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AcmiId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleAnnotation** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**CreatorCredit** | Pointer to **string** |  | [optional] 
**CreditLine** | Pointer to **string** |  | [optional] 
**HeadlineCredit** | Pointer to **string** |  | [optional] 
**Thumbnail** | Pointer to [**WorkThumbnail**](WorkThumbnail.md) |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IsOnDisplay** | Pointer to **bool** |  | [optional] 
**LastOnDisplayPlace** | Pointer to **string** |  | [optional] 
**LastOnDisplayDate** | Pointer to **string** |  | [optional] 
**IsContextIndigenous** | Pointer to **bool** |  | [optional] 
**MaterialDescription** | Pointer to **string** |  | [optional] 
**Unpublished** | Pointer to **bool** |  | [optional] 
**FirstProductionDate** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkSummary

`func NewWorkSummary() *WorkSummary`

NewWorkSummary instantiates a new WorkSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkSummaryWithDefaults

`func NewWorkSummaryWithDefaults() *WorkSummary`

NewWorkSummaryWithDefaults instantiates a new WorkSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAcmiId

`func (o *WorkSummary) GetAcmiId() string`

GetAcmiId returns the AcmiId field if non-nil, zero value otherwise.

### GetAcmiIdOk

`func (o *WorkSummary) GetAcmiIdOk() (*string, bool)`

GetAcmiIdOk returns a tuple with the AcmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmiId

`func (o *WorkSummary) SetAcmiId(v string)`

SetAcmiId sets AcmiId field to given value.

### HasAcmiId

`func (o *WorkSummary) HasAcmiId() bool`

HasAcmiId returns a boolean if a field has been set.

### GetTitle

`func (o *WorkSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAnnotation

`func (o *WorkSummary) GetTitleAnnotation() string`

GetTitleAnnotation returns the TitleAnnotation field if non-nil, zero value otherwise.

### GetTitleAnnotationOk

`func (o *WorkSummary) GetTitleAnnotationOk() (*string, bool)`

GetTitleAnnotationOk returns a tuple with the TitleAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAnnotation

`func (o *WorkSummary) SetTitleAnnotation(v string)`

SetTitleAnnotation sets TitleAnnotation field to given value.

### HasTitleAnnotation

`func (o *WorkSummary) HasTitleAnnotation() bool`

HasTitleAnnotation returns a boolean if a field has been set.

### GetSlug

`func (o *WorkSummary) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WorkSummary) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WorkSummary) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WorkSummary) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetCreatorCredit

`func (o *WorkSummary) GetCreatorCredit() string`

GetCreatorCredit returns the CreatorCredit field if non-nil, zero value otherwise.

### GetCreatorCreditOk

`func (o *WorkSummary) GetCreatorCreditOk() (*string, bool)`

GetCreatorCreditOk returns a tuple with the CreatorCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCredit

`func (o *WorkSummary) SetCreatorCredit(v string)`

SetCreatorCredit sets CreatorCredit field to given value.

### HasCreatorCredit

`func (o *WorkSummary) HasCreatorCredit() bool`

HasCreatorCredit returns a boolean if a field has been set.

### GetCreditLine

`func (o *WorkSummary) GetCreditLine() string`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *WorkSummary) GetCreditLineOk() (*string, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *WorkSummary) SetCreditLine(v string)`

SetCreditLine sets CreditLine field to given value.

### HasCreditLine

`func (o *WorkSummary) HasCreditLine() bool`

HasCreditLine returns a boolean if a field has been set.

### GetHeadlineCredit

`func (o *WorkSummary) GetHeadlineCredit() string`

GetHeadlineCredit returns the HeadlineCredit field if non-nil, zero value otherwise.

### GetHeadlineCreditOk

`func (o *WorkSummary) GetHeadlineCreditOk() (*string, bool)`

GetHeadlineCreditOk returns a tuple with the HeadlineCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlineCredit

`func (o *WorkSummary) SetHeadlineCredit(v string)`

SetHeadlineCredit sets HeadlineCredit field to given value.

### HasHeadlineCredit

`func (o *WorkSummary) HasHeadlineCredit() bool`

HasHeadlineCredit returns a boolean if a field has been set.

### GetThumbnail

`func (o *WorkSummary) GetThumbnail() WorkThumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *WorkSummary) GetThumbnailOk() (*WorkThumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *WorkSummary) SetThumbnail(v WorkThumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *WorkSummary) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetRecordType

`func (o *WorkSummary) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WorkSummary) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WorkSummary) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WorkSummary) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetType

`func (o *WorkSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsOnDisplay

`func (o *WorkSummary) GetIsOnDisplay() bool`

GetIsOnDisplay returns the IsOnDisplay field if non-nil, zero value otherwise.

### GetIsOnDisplayOk

`func (o *WorkSummary) GetIsOnDisplayOk() (*bool, bool)`

GetIsOnDisplayOk returns a tuple with the IsOnDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnDisplay

`func (o *WorkSummary) SetIsOnDisplay(v bool)`

SetIsOnDisplay sets IsOnDisplay field to given value.

### HasIsOnDisplay

`func (o *WorkSummary) HasIsOnDisplay() bool`

HasIsOnDisplay returns a boolean if a field has been set.

### GetLastOnDisplayPlace

`func (o *WorkSummary) GetLastOnDisplayPlace() string`

GetLastOnDisplayPlace returns the LastOnDisplayPlace field if non-nil, zero value otherwise.

### GetLastOnDisplayPlaceOk

`func (o *WorkSummary) GetLastOnDisplayPlaceOk() (*string, bool)`

GetLastOnDisplayPlaceOk returns a tuple with the LastOnDisplayPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnDisplayPlace

`func (o *WorkSummary) SetLastOnDisplayPlace(v string)`

SetLastOnDisplayPlace sets LastOnDisplayPlace field to given value.

### HasLastOnDisplayPlace

`func (o *WorkSummary) HasLastOnDisplayPlace() bool`

HasLastOnDisplayPlace returns a boolean if a field has been set.

### GetLastOnDisplayDate

`func (o *WorkSummary) GetLastOnDisplayDate() string`

GetLastOnDisplayDate returns the LastOnDisplayDate field if non-nil, zero value otherwise.

### GetLastOnDisplayDateOk

`func (o *WorkSummary) GetLastOnDisplayDateOk() (*string, bool)`

GetLastOnDisplayDateOk returns a tuple with the LastOnDisplayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnDisplayDate

`func (o *WorkSummary) SetLastOnDisplayDate(v string)`

SetLastOnDisplayDate sets LastOnDisplayDate field to given value.

### HasLastOnDisplayDate

`func (o *WorkSummary) HasLastOnDisplayDate() bool`

HasLastOnDisplayDate returns a boolean if a field has been set.

### GetIsContextIndigenous

`func (o *WorkSummary) GetIsContextIndigenous() bool`

GetIsContextIndigenous returns the IsContextIndigenous field if non-nil, zero value otherwise.

### GetIsContextIndigenousOk

`func (o *WorkSummary) GetIsContextIndigenousOk() (*bool, bool)`

GetIsContextIndigenousOk returns a tuple with the IsContextIndigenous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContextIndigenous

`func (o *WorkSummary) SetIsContextIndigenous(v bool)`

SetIsContextIndigenous sets IsContextIndigenous field to given value.

### HasIsContextIndigenous

`func (o *WorkSummary) HasIsContextIndigenous() bool`

HasIsContextIndigenous returns a boolean if a field has been set.

### GetMaterialDescription

`func (o *WorkSummary) GetMaterialDescription() string`

GetMaterialDescription returns the MaterialDescription field if non-nil, zero value otherwise.

### GetMaterialDescriptionOk

`func (o *WorkSummary) GetMaterialDescriptionOk() (*string, bool)`

GetMaterialDescriptionOk returns a tuple with the MaterialDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialDescription

`func (o *WorkSummary) SetMaterialDescription(v string)`

SetMaterialDescription sets MaterialDescription field to given value.

### HasMaterialDescription

`func (o *WorkSummary) HasMaterialDescription() bool`

HasMaterialDescription returns a boolean if a field has been set.

### GetUnpublished

`func (o *WorkSummary) GetUnpublished() bool`

GetUnpublished returns the Unpublished field if non-nil, zero value otherwise.

### GetUnpublishedOk

`func (o *WorkSummary) GetUnpublishedOk() (*bool, bool)`

GetUnpublishedOk returns a tuple with the Unpublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublished

`func (o *WorkSummary) SetUnpublished(v bool)`

SetUnpublished sets Unpublished field to given value.

### HasUnpublished

`func (o *WorkSummary) HasUnpublished() bool`

HasUnpublished returns a boolean if a field has been set.

### GetFirstProductionDate

`func (o *WorkSummary) GetFirstProductionDate() string`

GetFirstProductionDate returns the FirstProductionDate field if non-nil, zero value otherwise.

### GetFirstProductionDateOk

`func (o *WorkSummary) GetFirstProductionDateOk() (*string, bool)`

GetFirstProductionDateOk returns a tuple with the FirstProductionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstProductionDate

`func (o *WorkSummary) SetFirstProductionDate(v string)`

SetFirstProductionDate sets FirstProductionDate field to given value.

### HasFirstProductionDate

`func (o *WorkSummary) HasFirstProductionDate() bool`

HasFirstProductionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


