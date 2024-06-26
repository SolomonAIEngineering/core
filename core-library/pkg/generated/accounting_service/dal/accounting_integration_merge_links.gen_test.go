// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	accounting_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/accounting_service/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&accounting_servicev1.AccountingIntegrationMergeLinkORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&accounting_servicev1.AccountingIntegrationMergeLinkORM{}) fail: %s", err)
	}
}

func Test_accountingIntegrationMergeLinkORMQuery(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	accountingIntegrationMergeLinkORM = *accountingIntegrationMergeLinkORM.As(accountingIntegrationMergeLinkORM.TableName())
	_do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(accountingIntegrationMergeLinkORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <accounting_integration_merge_links> fail:", err)
		return
	}

	_, ok := accountingIntegrationMergeLinkORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from accountingIntegrationMergeLinkORM success")
	}

	err = _do.Create(&accounting_servicev1.AccountingIntegrationMergeLinkORM{})
	if err != nil {
		t.Error("create item in table <accounting_integration_merge_links> fail:", err)
	}

	err = _do.Save(&accounting_servicev1.AccountingIntegrationMergeLinkORM{})
	if err != nil {
		t.Error("create item in table <accounting_integration_merge_links> fail:", err)
	}

	err = _do.CreateInBatches([]*accounting_servicev1.AccountingIntegrationMergeLinkORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Select(accountingIntegrationMergeLinkORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <accounting_integration_merge_links> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*accounting_servicev1.AccountingIntegrationMergeLinkORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Select(accountingIntegrationMergeLinkORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Select(accountingIntegrationMergeLinkORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <accounting_integration_merge_links> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.ScanByPage(&accounting_servicev1.AccountingIntegrationMergeLinkORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <accounting_integration_merge_links> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <accounting_integration_merge_links> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <accounting_integration_merge_links> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <accounting_integration_merge_links> fail:", err)
	}
}

var AccountingIntegrationMergeLinkORMGetRecordByIDTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_GetRecordByID(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountingIntegrationMergeLinkORMGetRecordByIDsTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_GetRecordByIDs(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountingIntegrationMergeLinkORMCreateRecordTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_CreateRecord(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(accounting_servicev1.AccountingIntegrationMergeLinkORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var AccountingIntegrationMergeLinkORMUpdateRecordByIDTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_UpdateRecordByID(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(accounting_servicev1.AccountingIntegrationMergeLinkORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var AccountingIntegrationMergeLinkORMDeleteRecordByIDTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_DeleteRecordByID(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var AccountingIntegrationMergeLinkORMGetAllRecordsTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_GetAllRecords(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountingIntegrationMergeLinkORMCountAllTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_CountAll(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountingIntegrationMergeLinkORMGetByIDTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_GetByID(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountingIntegrationMergeLinkORMGetByIDsTestCase = []TestCase{}

func Test_accountingIntegrationMergeLinkORM_GetByIDs(t *testing.T) {
	accountingIntegrationMergeLinkORM := newAccountingIntegrationMergeLinkORM(db)
	do := accountingIntegrationMergeLinkORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountingIntegrationMergeLinkORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
