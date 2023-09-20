# Work

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
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
**BriefDescription** | Pointer to **string** |  | [optional] 
**ConstellationsPrimary** | Pointer to [**[]ConstellationSummary**](ConstellationSummary.md) |  | [optional] 
**ConstellationsOther** | Pointer to [**[]ConstellationSummary**](ConstellationSummary.md) |  | [optional] 
**Recommendations** | Pointer to [**[]WorkRecommendationsInner**](WorkRecommendationsInner.md) |  | [optional] 
**TitleForLabel** | Pointer to **string** |  | [optional] 
**CreatorCreditForLabel** | Pointer to **string** |  | [optional] 
**HeadlineCreditForLabel** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionForLabel** | Pointer to **string** |  | [optional] 
**CreditLineForLabel** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]WorkDetailsInner**](WorkDetailsInner.md) |  | [optional] 
**Stats** | Pointer to [**WorkStats**](WorkStats.md) |  | [optional] 
**Links** | Pointer to [**[]WorkLinksInner**](WorkLinksInner.md) |  | [optional] 
**CreatorsPrimary** | Pointer to [**[]CreatorSummary**](CreatorSummary.md) |  | [optional] 
**CreatorsOther** | Pointer to [**[]CreatorSummary**](CreatorSummary.md) |  | [optional] 
**VideoLinks** | Pointer to [**[]WorkVideoLinksInner**](WorkVideoLinksInner.md) |  | [optional] 
**MediaNote** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]WorkImagesInner**](WorkImagesInner.md) |  | [optional] 
**Videos** | Pointer to [**[]WorkVideosInner**](WorkVideosInner.md) |  | [optional] 
**Holdings** | Pointer to [**[]WorkHoldingsInner**](WorkHoldingsInner.md) |  | [optional] 
**PartOf** | Pointer to [**WorkSummary**](WorkSummary.md) |  | [optional] 
**Parts** | Pointer to [**[]WorkSummary**](WorkSummary.md) |  | [optional] 
**PartSiblings** | Pointer to [**[]WorkSummary**](WorkSummary.md) |  | [optional] 
**Group** | Pointer to [**WorkSummary**](WorkSummary.md) |  | [optional] 
**GroupWorks** | Pointer to [**[]WorkSummary**](WorkSummary.md) |  | [optional] 
**GroupSiblings** | Pointer to [**[]WorkSummary**](WorkSummary.md) |  | [optional] 
**Source** | Pointer to [**WorkSource**](WorkSource.md) |  | [optional] 
**SourceIdentifier** | Pointer to **string** |  | [optional] 
**ProductionPlaces** | Pointer to [**[]WorkProductionPlacesInner**](WorkProductionPlacesInner.md) |  | [optional] 
**ProductionDates** | Pointer to [**[]WorkProductionDatesInner**](WorkProductionDatesInner.md) |  | [optional] 
**Labels** | Pointer to **[]int32** |  | [optional] 
**EaasEnvironmentId** | Pointer to **string** |  | [optional] 
**ExternalReferences** | Pointer to [**[]WorkExternalReferencesInner**](WorkExternalReferencesInner.md) |  | [optional] 

## Methods

### NewWork

`func NewWork() *Work`

NewWork instantiates a new Work object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkWithDefaults

`func NewWorkWithDefaults() *Work`

NewWorkWithDefaults instantiates a new Work object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Work) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Work) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Work) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Work) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAcmiId

`func (o *Work) GetAcmiId() string`

GetAcmiId returns the AcmiId field if non-nil, zero value otherwise.

### GetAcmiIdOk

`func (o *Work) GetAcmiIdOk() (*string, bool)`

GetAcmiIdOk returns a tuple with the AcmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmiId

`func (o *Work) SetAcmiId(v string)`

SetAcmiId sets AcmiId field to given value.

### HasAcmiId

`func (o *Work) HasAcmiId() bool`

HasAcmiId returns a boolean if a field has been set.

### GetTitle

`func (o *Work) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Work) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Work) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Work) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAnnotation

`func (o *Work) GetTitleAnnotation() string`

GetTitleAnnotation returns the TitleAnnotation field if non-nil, zero value otherwise.

### GetTitleAnnotationOk

`func (o *Work) GetTitleAnnotationOk() (*string, bool)`

GetTitleAnnotationOk returns a tuple with the TitleAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAnnotation

`func (o *Work) SetTitleAnnotation(v string)`

SetTitleAnnotation sets TitleAnnotation field to given value.

### HasTitleAnnotation

`func (o *Work) HasTitleAnnotation() bool`

HasTitleAnnotation returns a boolean if a field has been set.

### GetSlug

`func (o *Work) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Work) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Work) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Work) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetCreatorCredit

`func (o *Work) GetCreatorCredit() string`

GetCreatorCredit returns the CreatorCredit field if non-nil, zero value otherwise.

### GetCreatorCreditOk

`func (o *Work) GetCreatorCreditOk() (*string, bool)`

GetCreatorCreditOk returns a tuple with the CreatorCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCredit

`func (o *Work) SetCreatorCredit(v string)`

SetCreatorCredit sets CreatorCredit field to given value.

### HasCreatorCredit

`func (o *Work) HasCreatorCredit() bool`

HasCreatorCredit returns a boolean if a field has been set.

### GetCreditLine

`func (o *Work) GetCreditLine() string`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *Work) GetCreditLineOk() (*string, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *Work) SetCreditLine(v string)`

SetCreditLine sets CreditLine field to given value.

### HasCreditLine

`func (o *Work) HasCreditLine() bool`

HasCreditLine returns a boolean if a field has been set.

### GetHeadlineCredit

`func (o *Work) GetHeadlineCredit() string`

GetHeadlineCredit returns the HeadlineCredit field if non-nil, zero value otherwise.

### GetHeadlineCreditOk

`func (o *Work) GetHeadlineCreditOk() (*string, bool)`

GetHeadlineCreditOk returns a tuple with the HeadlineCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlineCredit

`func (o *Work) SetHeadlineCredit(v string)`

SetHeadlineCredit sets HeadlineCredit field to given value.

### HasHeadlineCredit

`func (o *Work) HasHeadlineCredit() bool`

HasHeadlineCredit returns a boolean if a field has been set.

### GetThumbnail

`func (o *Work) GetThumbnail() WorkThumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Work) GetThumbnailOk() (*WorkThumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Work) SetThumbnail(v WorkThumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Work) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetRecordType

`func (o *Work) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Work) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Work) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Work) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetType

`func (o *Work) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Work) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Work) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Work) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsOnDisplay

`func (o *Work) GetIsOnDisplay() bool`

GetIsOnDisplay returns the IsOnDisplay field if non-nil, zero value otherwise.

### GetIsOnDisplayOk

`func (o *Work) GetIsOnDisplayOk() (*bool, bool)`

GetIsOnDisplayOk returns a tuple with the IsOnDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnDisplay

`func (o *Work) SetIsOnDisplay(v bool)`

SetIsOnDisplay sets IsOnDisplay field to given value.

### HasIsOnDisplay

`func (o *Work) HasIsOnDisplay() bool`

HasIsOnDisplay returns a boolean if a field has been set.

### GetLastOnDisplayPlace

`func (o *Work) GetLastOnDisplayPlace() string`

GetLastOnDisplayPlace returns the LastOnDisplayPlace field if non-nil, zero value otherwise.

### GetLastOnDisplayPlaceOk

`func (o *Work) GetLastOnDisplayPlaceOk() (*string, bool)`

GetLastOnDisplayPlaceOk returns a tuple with the LastOnDisplayPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnDisplayPlace

`func (o *Work) SetLastOnDisplayPlace(v string)`

SetLastOnDisplayPlace sets LastOnDisplayPlace field to given value.

### HasLastOnDisplayPlace

`func (o *Work) HasLastOnDisplayPlace() bool`

HasLastOnDisplayPlace returns a boolean if a field has been set.

### GetLastOnDisplayDate

`func (o *Work) GetLastOnDisplayDate() string`

GetLastOnDisplayDate returns the LastOnDisplayDate field if non-nil, zero value otherwise.

### GetLastOnDisplayDateOk

`func (o *Work) GetLastOnDisplayDateOk() (*string, bool)`

GetLastOnDisplayDateOk returns a tuple with the LastOnDisplayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnDisplayDate

`func (o *Work) SetLastOnDisplayDate(v string)`

SetLastOnDisplayDate sets LastOnDisplayDate field to given value.

### HasLastOnDisplayDate

`func (o *Work) HasLastOnDisplayDate() bool`

HasLastOnDisplayDate returns a boolean if a field has been set.

### GetIsContextIndigenous

`func (o *Work) GetIsContextIndigenous() bool`

GetIsContextIndigenous returns the IsContextIndigenous field if non-nil, zero value otherwise.

### GetIsContextIndigenousOk

`func (o *Work) GetIsContextIndigenousOk() (*bool, bool)`

GetIsContextIndigenousOk returns a tuple with the IsContextIndigenous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContextIndigenous

`func (o *Work) SetIsContextIndigenous(v bool)`

SetIsContextIndigenous sets IsContextIndigenous field to given value.

### HasIsContextIndigenous

`func (o *Work) HasIsContextIndigenous() bool`

HasIsContextIndigenous returns a boolean if a field has been set.

### GetMaterialDescription

`func (o *Work) GetMaterialDescription() string`

GetMaterialDescription returns the MaterialDescription field if non-nil, zero value otherwise.

### GetMaterialDescriptionOk

`func (o *Work) GetMaterialDescriptionOk() (*string, bool)`

GetMaterialDescriptionOk returns a tuple with the MaterialDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialDescription

`func (o *Work) SetMaterialDescription(v string)`

SetMaterialDescription sets MaterialDescription field to given value.

### HasMaterialDescription

`func (o *Work) HasMaterialDescription() bool`

HasMaterialDescription returns a boolean if a field has been set.

### GetUnpublished

`func (o *Work) GetUnpublished() bool`

GetUnpublished returns the Unpublished field if non-nil, zero value otherwise.

### GetUnpublishedOk

`func (o *Work) GetUnpublishedOk() (*bool, bool)`

GetUnpublishedOk returns a tuple with the Unpublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublished

`func (o *Work) SetUnpublished(v bool)`

SetUnpublished sets Unpublished field to given value.

### HasUnpublished

`func (o *Work) HasUnpublished() bool`

HasUnpublished returns a boolean if a field has been set.

### GetFirstProductionDate

`func (o *Work) GetFirstProductionDate() string`

GetFirstProductionDate returns the FirstProductionDate field if non-nil, zero value otherwise.

### GetFirstProductionDateOk

`func (o *Work) GetFirstProductionDateOk() (*string, bool)`

GetFirstProductionDateOk returns a tuple with the FirstProductionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstProductionDate

`func (o *Work) SetFirstProductionDate(v string)`

SetFirstProductionDate sets FirstProductionDate field to given value.

### HasFirstProductionDate

`func (o *Work) HasFirstProductionDate() bool`

HasFirstProductionDate returns a boolean if a field has been set.

### GetBriefDescription

`func (o *Work) GetBriefDescription() string`

GetBriefDescription returns the BriefDescription field if non-nil, zero value otherwise.

### GetBriefDescriptionOk

`func (o *Work) GetBriefDescriptionOk() (*string, bool)`

GetBriefDescriptionOk returns a tuple with the BriefDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefDescription

`func (o *Work) SetBriefDescription(v string)`

SetBriefDescription sets BriefDescription field to given value.

### HasBriefDescription

`func (o *Work) HasBriefDescription() bool`

HasBriefDescription returns a boolean if a field has been set.

### GetConstellationsPrimary

`func (o *Work) GetConstellationsPrimary() []ConstellationSummary`

GetConstellationsPrimary returns the ConstellationsPrimary field if non-nil, zero value otherwise.

### GetConstellationsPrimaryOk

`func (o *Work) GetConstellationsPrimaryOk() (*[]ConstellationSummary, bool)`

GetConstellationsPrimaryOk returns a tuple with the ConstellationsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstellationsPrimary

`func (o *Work) SetConstellationsPrimary(v []ConstellationSummary)`

SetConstellationsPrimary sets ConstellationsPrimary field to given value.

### HasConstellationsPrimary

`func (o *Work) HasConstellationsPrimary() bool`

HasConstellationsPrimary returns a boolean if a field has been set.

### GetConstellationsOther

`func (o *Work) GetConstellationsOther() []ConstellationSummary`

GetConstellationsOther returns the ConstellationsOther field if non-nil, zero value otherwise.

### GetConstellationsOtherOk

`func (o *Work) GetConstellationsOtherOk() (*[]ConstellationSummary, bool)`

GetConstellationsOtherOk returns a tuple with the ConstellationsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstellationsOther

`func (o *Work) SetConstellationsOther(v []ConstellationSummary)`

SetConstellationsOther sets ConstellationsOther field to given value.

### HasConstellationsOther

`func (o *Work) HasConstellationsOther() bool`

HasConstellationsOther returns a boolean if a field has been set.

### GetRecommendations

`func (o *Work) GetRecommendations() []WorkRecommendationsInner`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *Work) GetRecommendationsOk() (*[]WorkRecommendationsInner, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *Work) SetRecommendations(v []WorkRecommendationsInner)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *Work) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetTitleForLabel

`func (o *Work) GetTitleForLabel() string`

GetTitleForLabel returns the TitleForLabel field if non-nil, zero value otherwise.

### GetTitleForLabelOk

`func (o *Work) GetTitleForLabelOk() (*string, bool)`

GetTitleForLabelOk returns a tuple with the TitleForLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleForLabel

`func (o *Work) SetTitleForLabel(v string)`

SetTitleForLabel sets TitleForLabel field to given value.

### HasTitleForLabel

`func (o *Work) HasTitleForLabel() bool`

HasTitleForLabel returns a boolean if a field has been set.

### GetCreatorCreditForLabel

`func (o *Work) GetCreatorCreditForLabel() string`

GetCreatorCreditForLabel returns the CreatorCreditForLabel field if non-nil, zero value otherwise.

### GetCreatorCreditForLabelOk

`func (o *Work) GetCreatorCreditForLabelOk() (*string, bool)`

GetCreatorCreditForLabelOk returns a tuple with the CreatorCreditForLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCreditForLabel

`func (o *Work) SetCreatorCreditForLabel(v string)`

SetCreatorCreditForLabel sets CreatorCreditForLabel field to given value.

### HasCreatorCreditForLabel

`func (o *Work) HasCreatorCreditForLabel() bool`

HasCreatorCreditForLabel returns a boolean if a field has been set.

### GetHeadlineCreditForLabel

`func (o *Work) GetHeadlineCreditForLabel() string`

GetHeadlineCreditForLabel returns the HeadlineCreditForLabel field if non-nil, zero value otherwise.

### GetHeadlineCreditForLabelOk

`func (o *Work) GetHeadlineCreditForLabelOk() (*string, bool)`

GetHeadlineCreditForLabelOk returns a tuple with the HeadlineCreditForLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlineCreditForLabel

`func (o *Work) SetHeadlineCreditForLabel(v string)`

SetHeadlineCreditForLabel sets HeadlineCreditForLabel field to given value.

### HasHeadlineCreditForLabel

`func (o *Work) HasHeadlineCreditForLabel() bool`

HasHeadlineCreditForLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Work) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Work) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Work) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Work) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionForLabel

`func (o *Work) GetDescriptionForLabel() string`

GetDescriptionForLabel returns the DescriptionForLabel field if non-nil, zero value otherwise.

### GetDescriptionForLabelOk

`func (o *Work) GetDescriptionForLabelOk() (*string, bool)`

GetDescriptionForLabelOk returns a tuple with the DescriptionForLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionForLabel

`func (o *Work) SetDescriptionForLabel(v string)`

SetDescriptionForLabel sets DescriptionForLabel field to given value.

### HasDescriptionForLabel

`func (o *Work) HasDescriptionForLabel() bool`

HasDescriptionForLabel returns a boolean if a field has been set.

### GetCreditLineForLabel

`func (o *Work) GetCreditLineForLabel() string`

GetCreditLineForLabel returns the CreditLineForLabel field if non-nil, zero value otherwise.

### GetCreditLineForLabelOk

`func (o *Work) GetCreditLineForLabelOk() (*string, bool)`

GetCreditLineForLabelOk returns a tuple with the CreditLineForLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLineForLabel

`func (o *Work) SetCreditLineForLabel(v string)`

SetCreditLineForLabel sets CreditLineForLabel field to given value.

### HasCreditLineForLabel

`func (o *Work) HasCreditLineForLabel() bool`

HasCreditLineForLabel returns a boolean if a field has been set.

### GetDetails

`func (o *Work) GetDetails() []WorkDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Work) GetDetailsOk() (*[]WorkDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Work) SetDetails(v []WorkDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Work) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetStats

`func (o *Work) GetStats() WorkStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Work) GetStatsOk() (*WorkStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Work) SetStats(v WorkStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Work) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetLinks

`func (o *Work) GetLinks() []WorkLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Work) GetLinksOk() (*[]WorkLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Work) SetLinks(v []WorkLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Work) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatorsPrimary

`func (o *Work) GetCreatorsPrimary() []CreatorSummary`

GetCreatorsPrimary returns the CreatorsPrimary field if non-nil, zero value otherwise.

### GetCreatorsPrimaryOk

`func (o *Work) GetCreatorsPrimaryOk() (*[]CreatorSummary, bool)`

GetCreatorsPrimaryOk returns a tuple with the CreatorsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorsPrimary

`func (o *Work) SetCreatorsPrimary(v []CreatorSummary)`

SetCreatorsPrimary sets CreatorsPrimary field to given value.

### HasCreatorsPrimary

`func (o *Work) HasCreatorsPrimary() bool`

HasCreatorsPrimary returns a boolean if a field has been set.

### GetCreatorsOther

`func (o *Work) GetCreatorsOther() []CreatorSummary`

GetCreatorsOther returns the CreatorsOther field if non-nil, zero value otherwise.

### GetCreatorsOtherOk

`func (o *Work) GetCreatorsOtherOk() (*[]CreatorSummary, bool)`

GetCreatorsOtherOk returns a tuple with the CreatorsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorsOther

`func (o *Work) SetCreatorsOther(v []CreatorSummary)`

SetCreatorsOther sets CreatorsOther field to given value.

### HasCreatorsOther

`func (o *Work) HasCreatorsOther() bool`

HasCreatorsOther returns a boolean if a field has been set.

### GetVideoLinks

`func (o *Work) GetVideoLinks() []WorkVideoLinksInner`

GetVideoLinks returns the VideoLinks field if non-nil, zero value otherwise.

### GetVideoLinksOk

`func (o *Work) GetVideoLinksOk() (*[]WorkVideoLinksInner, bool)`

GetVideoLinksOk returns a tuple with the VideoLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLinks

`func (o *Work) SetVideoLinks(v []WorkVideoLinksInner)`

SetVideoLinks sets VideoLinks field to given value.

### HasVideoLinks

`func (o *Work) HasVideoLinks() bool`

HasVideoLinks returns a boolean if a field has been set.

### GetMediaNote

`func (o *Work) GetMediaNote() string`

GetMediaNote returns the MediaNote field if non-nil, zero value otherwise.

### GetMediaNoteOk

`func (o *Work) GetMediaNoteOk() (*string, bool)`

GetMediaNoteOk returns a tuple with the MediaNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaNote

`func (o *Work) SetMediaNote(v string)`

SetMediaNote sets MediaNote field to given value.

### HasMediaNote

`func (o *Work) HasMediaNote() bool`

HasMediaNote returns a boolean if a field has been set.

### GetImages

`func (o *Work) GetImages() []WorkImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Work) GetImagesOk() (*[]WorkImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Work) SetImages(v []WorkImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *Work) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetVideos

`func (o *Work) GetVideos() []WorkVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *Work) GetVideosOk() (*[]WorkVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *Work) SetVideos(v []WorkVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *Work) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetHoldings

`func (o *Work) GetHoldings() []WorkHoldingsInner`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *Work) GetHoldingsOk() (*[]WorkHoldingsInner, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *Work) SetHoldings(v []WorkHoldingsInner)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *Work) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetPartOf

`func (o *Work) GetPartOf() WorkSummary`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *Work) GetPartOfOk() (*WorkSummary, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *Work) SetPartOf(v WorkSummary)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *Work) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.

### GetParts

`func (o *Work) GetParts() []WorkSummary`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *Work) GetPartsOk() (*[]WorkSummary, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *Work) SetParts(v []WorkSummary)`

SetParts sets Parts field to given value.

### HasParts

`func (o *Work) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetPartSiblings

`func (o *Work) GetPartSiblings() []WorkSummary`

GetPartSiblings returns the PartSiblings field if non-nil, zero value otherwise.

### GetPartSiblingsOk

`func (o *Work) GetPartSiblingsOk() (*[]WorkSummary, bool)`

GetPartSiblingsOk returns a tuple with the PartSiblings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartSiblings

`func (o *Work) SetPartSiblings(v []WorkSummary)`

SetPartSiblings sets PartSiblings field to given value.

### HasPartSiblings

`func (o *Work) HasPartSiblings() bool`

HasPartSiblings returns a boolean if a field has been set.

### GetGroup

`func (o *Work) GetGroup() WorkSummary`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Work) GetGroupOk() (*WorkSummary, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Work) SetGroup(v WorkSummary)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Work) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupWorks

`func (o *Work) GetGroupWorks() []WorkSummary`

GetGroupWorks returns the GroupWorks field if non-nil, zero value otherwise.

### GetGroupWorksOk

`func (o *Work) GetGroupWorksOk() (*[]WorkSummary, bool)`

GetGroupWorksOk returns a tuple with the GroupWorks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupWorks

`func (o *Work) SetGroupWorks(v []WorkSummary)`

SetGroupWorks sets GroupWorks field to given value.

### HasGroupWorks

`func (o *Work) HasGroupWorks() bool`

HasGroupWorks returns a boolean if a field has been set.

### GetGroupSiblings

`func (o *Work) GetGroupSiblings() []WorkSummary`

GetGroupSiblings returns the GroupSiblings field if non-nil, zero value otherwise.

### GetGroupSiblingsOk

`func (o *Work) GetGroupSiblingsOk() (*[]WorkSummary, bool)`

GetGroupSiblingsOk returns a tuple with the GroupSiblings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSiblings

`func (o *Work) SetGroupSiblings(v []WorkSummary)`

SetGroupSiblings sets GroupSiblings field to given value.

### HasGroupSiblings

`func (o *Work) HasGroupSiblings() bool`

HasGroupSiblings returns a boolean if a field has been set.

### GetSource

`func (o *Work) GetSource() WorkSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Work) GetSourceOk() (*WorkSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Work) SetSource(v WorkSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Work) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *Work) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *Work) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *Work) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *Work) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetProductionPlaces

`func (o *Work) GetProductionPlaces() []WorkProductionPlacesInner`

GetProductionPlaces returns the ProductionPlaces field if non-nil, zero value otherwise.

### GetProductionPlacesOk

`func (o *Work) GetProductionPlacesOk() (*[]WorkProductionPlacesInner, bool)`

GetProductionPlacesOk returns a tuple with the ProductionPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPlaces

`func (o *Work) SetProductionPlaces(v []WorkProductionPlacesInner)`

SetProductionPlaces sets ProductionPlaces field to given value.

### HasProductionPlaces

`func (o *Work) HasProductionPlaces() bool`

HasProductionPlaces returns a boolean if a field has been set.

### GetProductionDates

`func (o *Work) GetProductionDates() []WorkProductionDatesInner`

GetProductionDates returns the ProductionDates field if non-nil, zero value otherwise.

### GetProductionDatesOk

`func (o *Work) GetProductionDatesOk() (*[]WorkProductionDatesInner, bool)`

GetProductionDatesOk returns a tuple with the ProductionDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionDates

`func (o *Work) SetProductionDates(v []WorkProductionDatesInner)`

SetProductionDates sets ProductionDates field to given value.

### HasProductionDates

`func (o *Work) HasProductionDates() bool`

HasProductionDates returns a boolean if a field has been set.

### GetLabels

`func (o *Work) GetLabels() []int32`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Work) GetLabelsOk() (*[]int32, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Work) SetLabels(v []int32)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Work) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEaasEnvironmentId

`func (o *Work) GetEaasEnvironmentId() string`

GetEaasEnvironmentId returns the EaasEnvironmentId field if non-nil, zero value otherwise.

### GetEaasEnvironmentIdOk

`func (o *Work) GetEaasEnvironmentIdOk() (*string, bool)`

GetEaasEnvironmentIdOk returns a tuple with the EaasEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaasEnvironmentId

`func (o *Work) SetEaasEnvironmentId(v string)`

SetEaasEnvironmentId sets EaasEnvironmentId field to given value.

### HasEaasEnvironmentId

`func (o *Work) HasEaasEnvironmentId() bool`

HasEaasEnvironmentId returns a boolean if a field has been set.

### GetExternalReferences

`func (o *Work) GetExternalReferences() []WorkExternalReferencesInner`

GetExternalReferences returns the ExternalReferences field if non-nil, zero value otherwise.

### GetExternalReferencesOk

`func (o *Work) GetExternalReferencesOk() (*[]WorkExternalReferencesInner, bool)`

GetExternalReferencesOk returns a tuple with the ExternalReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferences

`func (o *Work) SetExternalReferences(v []WorkExternalReferencesInner)`

SetExternalReferences sets ExternalReferences field to given value.

### HasExternalReferences

`func (o *Work) HasExternalReferences() bool`

HasExternalReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


