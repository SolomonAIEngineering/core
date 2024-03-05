/* tslint:disable */
/* eslint-disable */
/**
 * User Service
 * Solomon AI User Service API
 *
 * The version of the OpenAPI document: 0.1
 * Contact: yoanyomba@solomon-ai.co
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BankInfo } from './BankInfo';
import {
    BankInfoFromJSON,
    BankInfoFromJSONTyped,
    BankInfoToJSON,
} from './BankInfo';
import type { Dependents } from './Dependents';
import {
    DependentsFromJSON,
    DependentsFromJSONTyped,
    DependentsToJSON,
} from './Dependents';
import type { EmployeTimeOffBalance } from './EmployeTimeOffBalance';
import {
    EmployeTimeOffBalanceFromJSON,
    EmployeTimeOffBalanceFromJSONTyped,
    EmployeTimeOffBalanceToJSON,
} from './EmployeTimeOffBalance';
import type { EmployeeBenefits } from './EmployeeBenefits';
import {
    EmployeeBenefitsFromJSON,
    EmployeeBenefitsFromJSONTyped,
    EmployeeBenefitsToJSON,
} from './EmployeeBenefits';
import type { EmployeeJobPositionAtCompany } from './EmployeeJobPositionAtCompany';
import {
    EmployeeJobPositionAtCompanyFromJSON,
    EmployeeJobPositionAtCompanyFromJSONTyped,
    EmployeeJobPositionAtCompanyToJSON,
} from './EmployeeJobPositionAtCompany';
import type { EmployeePayrollRun } from './EmployeePayrollRun';
import {
    EmployeePayrollRunFromJSON,
    EmployeePayrollRunFromJSONTyped,
    EmployeePayrollRunToJSON,
} from './EmployeePayrollRun';
import type { EmploymentStatus } from './EmploymentStatus';
import {
    EmploymentStatusFromJSON,
    EmploymentStatusFromJSONTyped,
    EmploymentStatusToJSON,
} from './EmploymentStatus';
import type { Ethnicity } from './Ethnicity';
import {
    EthnicityFromJSON,
    EthnicityFromJSONTyped,
    EthnicityToJSON,
} from './Ethnicity';
import type { Gender } from './Gender';
import {
    GenderFromJSON,
    GenderFromJSONTyped,
    GenderToJSON,
} from './Gender';
import type { Group } from './Group';
import {
    GroupFromJSON,
    GroupFromJSONTyped,
    GroupToJSON,
} from './Group';
import type { LocationAddress } from './LocationAddress';
import {
    LocationAddressFromJSON,
    LocationAddressFromJSONTyped,
    LocationAddressToJSON,
} from './LocationAddress';
import type { MaritalStatus } from './MaritalStatus';
import {
    MaritalStatusFromJSON,
    MaritalStatusFromJSONTyped,
    MaritalStatusToJSON,
} from './MaritalStatus';

/**
 * The Employee object is used to represent any person who has been employed by a company.
 * @export
 * @interface Employee
 */
export interface Employee {
    /**
     * 
     * @type {string}
     * @memberof Employee
     */
    id?: string;
    /**
     * The third-party API ID of the matching object.
     * @type {string}
     * @memberof Employee
     */
    remoteId?: string;
    /**
     * 
     * @type {string}
     * @memberof Employee
     */
    employeeNumber?: string;
    /**
     * The ID of the employee's company.
     * @type {string}
     * @memberof Employee
     */
    companyId?: string;
    /**
     * The employee's first name.
     * @type {string}
     * @memberof Employee
     */
    firstName?: string;
    /**
     * The employee's last name.
     * @type {string}
     * @memberof Employee
     */
    lastName?: string;
    /**
     * The employee's preferred name.
     * @type {string}
     * @memberof Employee
     */
    employeesPreferredName?: string;
    /**
     * 
     * @type {string}
     * @memberof Employee
     */
    displayFullName?: string;
    /**
     * The employee's username that appears in the remote UI.
     * @type {string}
     * @memberof Employee
     */
    employeeUserNameAsSeenInRemoteUi?: string;
    /**
     * The employee's work email.
     * @type {string}
     * @memberof Employee
     */
    workEmail?: string;
    /**
     * 
     * @type {string}
     * @memberof Employee
     */
    personalEmail?: string;
    /**
     * The employee's mobile phone number.
     * @type {string}
     * @memberof Employee
     */
    mobilePhoneNumber?: string;
    /**
     * 
     * @type {Array<EmployeeJobPositionAtCompany>}
     * @memberof Employee
     */
    employments?: Array<EmployeeJobPositionAtCompany>;
    /**
     * UUID fields
     * @type {string}
     * @memberof Employee
     */
    employmentType?: string;
    /**
     * 
     * @type {LocationAddress}
     * @memberof Employee
     */
    homeLocation?: LocationAddress;
    /**
     * 
     * @type {LocationAddress}
     * @memberof Employee
     */
    workLocation?: LocationAddress;
    /**
     * 
     * @type {Employee}
     * @memberof Employee
     */
    manager?: Employee;
    /**
     * 
     * @type {Group}
     * @memberof Employee
     */
    group?: Group;
    /**
     * 
     * @type {string}
     * @memberof Employee
     */
    ssn?: string;
    /**
     * 
     * @type {Gender}
     * @memberof Employee
     */
    gender?: Gender;
    /**
     * 
     * @type {Ethnicity}
     * @memberof Employee
     */
    ethnicity?: Ethnicity;
    /**
     * 
     * @type {MaritalStatus}
     * @memberof Employee
     */
    maritalStatus?: MaritalStatus;
    /**
     * The employee's date of birth.
     * 
     * Use string for ISO 8601 datetime
     * @type {string}
     * @memberof Employee
     */
    dateOfBirth?: string;
    /**
     * The date that the employee started working. 
     * If an employee was rehired, the most recent start date will be returned.
     * @type {Date}
     * @memberof Employee
     */
    startDate?: Date;
    /**
     * When the third party's employee was created.
     * @type {Date}
     * @memberof Employee
     */
    remoteCreatedAt?: Date;
    /**
     * 
     * @type {EmploymentStatus}
     * @memberof Employee
     */
    employmentStatus?: EmploymentStatus;
    /**
     * The employee's termination date.
     * @type {Date}
     * @memberof Employee
     */
    terminationDate?: Date;
    /**
     * The URL of the employee's avatar image.
     * @type {string}
     * @memberof Employee
     */
    avatar?: string;
    /**
     * 
     * @type {Array<BankInfo>}
     * @memberof Employee
     */
    bankAccounts?: Array<BankInfo>;
    /**
     * 
     * @type {Array<Dependents>}
     * @memberof Employee
     */
    dependents?: Array<Dependents>;
    /**
     * Represent an employee's pay statement for a specific payroll run.
     * @type {Array<EmployeePayrollRun>}
     * @memberof Employee
     */
    payrollRuns?: Array<EmployeePayrollRun>;
    /**
     * 
     * @type {EmployeTimeOffBalance}
     * @memberof Employee
     */
    payTimeOffBalance?: EmployeTimeOffBalance;
    /**
     * the benefits associated with the employee.
     * @type {Array<EmployeeBenefits>}
     * @memberof Employee
     */
    benefits?: Array<EmployeeBenefits>;
    /**
     * 
     * @type {string}
     * @memberof Employee
     */
    mergeAccountId?: string;
    /**
     * 
     * @type {Date}
     * @memberof Employee
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Employee
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof Employee
     */
    remoteWasDeleted?: boolean;
}

/**
 * Check if a given object implements the Employee interface.
 */
export function instanceOfEmployee(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EmployeeFromJSON(json: any): Employee {
    return EmployeeFromJSONTyped(json, false);
}

export function EmployeeFromJSONTyped(json: any, ignoreDiscriminator: boolean): Employee {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'employeeNumber': !exists(json, 'employeeNumber') ? undefined : json['employeeNumber'],
        'companyId': !exists(json, 'companyId') ? undefined : json['companyId'],
        'firstName': !exists(json, 'firstName') ? undefined : json['firstName'],
        'lastName': !exists(json, 'lastName') ? undefined : json['lastName'],
        'employeesPreferredName': !exists(json, 'employeesPreferredName') ? undefined : json['employeesPreferredName'],
        'displayFullName': !exists(json, 'displayFullName') ? undefined : json['displayFullName'],
        'employeeUserNameAsSeenInRemoteUi': !exists(json, 'employeeUserNameAsSeenInRemoteUi') ? undefined : json['employeeUserNameAsSeenInRemoteUi'],
        'workEmail': !exists(json, 'workEmail') ? undefined : json['workEmail'],
        'personalEmail': !exists(json, 'personalEmail') ? undefined : json['personalEmail'],
        'mobilePhoneNumber': !exists(json, 'mobilePhoneNumber') ? undefined : json['mobilePhoneNumber'],
        'employments': !exists(json, 'employments') ? undefined : ((json['employments'] as Array<any>).map(EmployeeJobPositionAtCompanyFromJSON)),
        'employmentType': !exists(json, 'employmentType') ? undefined : json['employmentType'],
        'homeLocation': !exists(json, 'homeLocation') ? undefined : LocationAddressFromJSON(json['homeLocation']),
        'workLocation': !exists(json, 'workLocation') ? undefined : LocationAddressFromJSON(json['workLocation']),
        'manager': !exists(json, 'manager') ? undefined : EmployeeFromJSON(json['manager']),
        'group': !exists(json, 'group') ? undefined : GroupFromJSON(json['group']),
        'ssn': !exists(json, 'ssn') ? undefined : json['ssn'],
        'gender': !exists(json, 'gender') ? undefined : GenderFromJSON(json['gender']),
        'ethnicity': !exists(json, 'ethnicity') ? undefined : EthnicityFromJSON(json['ethnicity']),
        'maritalStatus': !exists(json, 'maritalStatus') ? undefined : MaritalStatusFromJSON(json['maritalStatus']),
        'dateOfBirth': !exists(json, 'dateOfBirth') ? undefined : json['dateOfBirth'],
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'remoteCreatedAt': !exists(json, 'remoteCreatedAt') ? undefined : (new Date(json['remoteCreatedAt'])),
        'employmentStatus': !exists(json, 'employmentStatus') ? undefined : EmploymentStatusFromJSON(json['employmentStatus']),
        'terminationDate': !exists(json, 'terminationDate') ? undefined : (new Date(json['terminationDate'])),
        'avatar': !exists(json, 'avatar') ? undefined : json['avatar'],
        'bankAccounts': !exists(json, 'bankAccounts') ? undefined : ((json['bankAccounts'] as Array<any>).map(BankInfoFromJSON)),
        'dependents': !exists(json, 'dependents') ? undefined : ((json['dependents'] as Array<any>).map(DependentsFromJSON)),
        'payrollRuns': !exists(json, 'payrollRuns') ? undefined : ((json['payrollRuns'] as Array<any>).map(EmployeePayrollRunFromJSON)),
        'payTimeOffBalance': !exists(json, 'payTimeOffBalance') ? undefined : EmployeTimeOffBalanceFromJSON(json['payTimeOffBalance']),
        'benefits': !exists(json, 'benefits') ? undefined : ((json['benefits'] as Array<any>).map(EmployeeBenefitsFromJSON)),
        'mergeAccountId': !exists(json, 'mergeAccountId') ? undefined : json['mergeAccountId'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
    };
}

export function EmployeeToJSON(value?: Employee | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'employeeNumber': value.employeeNumber,
        'companyId': value.companyId,
        'firstName': value.firstName,
        'lastName': value.lastName,
        'employeesPreferredName': value.employeesPreferredName,
        'displayFullName': value.displayFullName,
        'employeeUserNameAsSeenInRemoteUi': value.employeeUserNameAsSeenInRemoteUi,
        'workEmail': value.workEmail,
        'personalEmail': value.personalEmail,
        'mobilePhoneNumber': value.mobilePhoneNumber,
        'employments': value.employments === undefined ? undefined : ((value.employments as Array<any>).map(EmployeeJobPositionAtCompanyToJSON)),
        'employmentType': value.employmentType,
        'homeLocation': LocationAddressToJSON(value.homeLocation),
        'workLocation': LocationAddressToJSON(value.workLocation),
        'manager': EmployeeToJSON(value.manager),
        'group': GroupToJSON(value.group),
        'ssn': value.ssn,
        'gender': GenderToJSON(value.gender),
        'ethnicity': EthnicityToJSON(value.ethnicity),
        'maritalStatus': MaritalStatusToJSON(value.maritalStatus),
        'dateOfBirth': value.dateOfBirth,
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString()),
        'remoteCreatedAt': value.remoteCreatedAt === undefined ? undefined : (value.remoteCreatedAt.toISOString()),
        'employmentStatus': EmploymentStatusToJSON(value.employmentStatus),
        'terminationDate': value.terminationDate === undefined ? undefined : (value.terminationDate.toISOString()),
        'avatar': value.avatar,
        'bankAccounts': value.bankAccounts === undefined ? undefined : ((value.bankAccounts as Array<any>).map(BankInfoToJSON)),
        'dependents': value.dependents === undefined ? undefined : ((value.dependents as Array<any>).map(DependentsToJSON)),
        'payrollRuns': value.payrollRuns === undefined ? undefined : ((value.payrollRuns as Array<any>).map(EmployeePayrollRunToJSON)),
        'payTimeOffBalance': EmployeTimeOffBalanceToJSON(value.payTimeOffBalance),
        'benefits': value.benefits === undefined ? undefined : ((value.benefits as Array<any>).map(EmployeeBenefitsToJSON)),
        'mergeAccountId': value.mergeAccountId,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'remoteWasDeleted': value.remoteWasDeleted,
    };
}
