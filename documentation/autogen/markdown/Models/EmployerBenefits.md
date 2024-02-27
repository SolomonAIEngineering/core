# EmployerBenefits
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object.  External system identifier (integration) | [optional] [default to null] |
| **benefitPlanType** | [**BenefitPlanType**](BenefitPlanType.md) |  | [optional] [default to null] |
| **name** | **String** | The employer benefit&#39;s name - typically the carrier or network name.  Name of the benefit plan | [optional] [default to null] |
| **description** | **String** | The employer benefit&#39;s description.  Description of the plan | [optional] [default to null] |
| **deductionCode** | **String** | The employer benefit&#39;s deduction code.  Code for payroll deduction | [optional] [default to null] |
| **remoteWasDeleted** | **Date** | Indicates whether or not this object has been deleted in the third party platform.  Flag for deletion status | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |
| **modifiedAt** | **Date** | Last modification date in ISO 8601 format | [optional] [default to null] |
| **mergeAccountId** | **String** | Merge record UUID | [optional] [default to null] |
| **employeeBenefits** | [**List**](EmployeeBenefits.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

