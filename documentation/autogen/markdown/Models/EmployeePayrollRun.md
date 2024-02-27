# EmployeePayrollRun
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **grossPay** | **Double** | The total earnings throughout a given period for an employee before any deductions are made. | [optional] [default to null] |
| **netPay** | **Double** | The take-home pay throughout a given period for an employee after deductions are made. | [optional] [default to null] |
| **startDate** | **Date** | The day and time the payroll run started. | [optional] [default to null] |
| **endDate** | **Date** |  | [optional] [default to null] |
| **checkDate** | **Date** |  | [optional] [default to null] |
| **earnings** | [**List**](Earning.md) |  | [optional] [default to null] |
| **deductions** | [**List**](Deduction.md) | The Deduction object is used to represent an array of the wages withheld  from total earnings for the purpose of paying taxes. | [optional] [default to null] |
| **taxes** | [**List**](Tax.md) | The Tax object is used to represent an array of the tax deductions  for a given employee&#39;s payroll run. | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted in the third party platform. | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **payrollRunMergeAccountId** | **String** | The payroll being run. | [optional] [default to null] |
| **employeeMergeAccountId** | **String** | The employee whose payroll is being run. | [optional] [default to null] |
| **mergeAccountId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

