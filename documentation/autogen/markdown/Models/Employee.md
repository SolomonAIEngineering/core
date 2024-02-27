# Employee
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **employeeNumber** | **String** |  | [optional] [default to null] |
| **companyId** | **String** | The ID of the employee&#39;s company. | [optional] [default to null] |
| **firstName** | **String** | The employee&#39;s first name. | [optional] [default to null] |
| **lastName** | **String** | The employee&#39;s last name. | [optional] [default to null] |
| **employeesPreferredName** | **String** | The employee&#39;s preferred name. | [optional] [default to null] |
| **displayFullName** | **String** |  | [optional] [default to null] |
| **employeeUserNameAsSeenInRemoteUi** | **String** | The employee&#39;s username that appears in the remote UI. | [optional] [default to null] |
| **workEmail** | **String** | The employee&#39;s work email. | [optional] [default to null] |
| **personalEmail** | **String** |  | [optional] [default to null] |
| **mobilePhoneNumber** | **String** | The employee&#39;s mobile phone number. | [optional] [default to null] |
| **employments** | [**List**](EmployeeJobPositionAtCompany.md) |  | [optional] [default to null] |
| **employmentType** | **String** | UUID fields | [optional] [default to null] |
| **homeLocation** | [**LocationAddress**](LocationAddress.md) |  | [optional] [default to null] |
| **workLocation** | [**LocationAddress**](LocationAddress.md) |  | [optional] [default to null] |
| **manager** | [**Employee**](Employee.md) |  | [optional] [default to null] |
| **group** | [**Group**](Group.md) |  | [optional] [default to null] |
| **ssn** | **String** |  | [optional] [default to null] |
| **gender** | [**Gender**](Gender.md) |  | [optional] [default to null] |
| **ethnicity** | [**Ethnicity**](Ethnicity.md) |  | [optional] [default to null] |
| **maritalStatus** | [**MaritalStatus**](MaritalStatus.md) |  | [optional] [default to null] |
| **dateOfBirth** | **String** | The employee&#39;s date of birth.  Use string for ISO 8601 datetime | [optional] [default to null] |
| **startDate** | **Date** | The date that the employee started working.  If an employee was rehired, the most recent start date will be returned. | [optional] [default to null] |
| **remoteCreatedAt** | **Date** | When the third party&#39;s employee was created. | [optional] [default to null] |
| **employmentStatus** | [**EmploymentStatus**](EmploymentStatus.md) |  | [optional] [default to null] |
| **terminationDate** | **Date** | The employee&#39;s termination date. | [optional] [default to null] |
| **avatar** | **String** | The URL of the employee&#39;s avatar image. | [optional] [default to null] |
| **bankAccounts** | [**List**](BankInfo.md) |  | [optional] [default to null] |
| **dependents** | [**List**](Dependents.md) |  | [optional] [default to null] |
| **payrollRuns** | [**List**](EmployeePayrollRun.md) | Represent an employee&#39;s pay statement for a specific payroll run. | [optional] [default to null] |
| **payTimeOffBalance** | [**EmployeTimeOffBalance**](EmployeTimeOffBalance.md) |  | [optional] [default to null] |
| **benefits** | [**List**](EmployeeBenefits.md) | the benefits associated with the employee. | [optional] [default to null] |
| **mergeAccountId** | **String** |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

