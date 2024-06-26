// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	financial_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/financial_service/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&financial_servicev1.AccountStatementsORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&financial_servicev1.AccountStatementsORM{}) fail: %s", err)
	}
}

func Test_accountStatementsORMQuery(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	accountStatementsORM = *accountStatementsORM.As(accountStatementsORM.TableName())
	_do := accountStatementsORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(accountStatementsORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <account_statements> fail:", err)
		return
	}

	_, ok := accountStatementsORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from accountStatementsORM success")
	}

	err = _do.Create(&financial_servicev1.AccountStatementsORM{})
	if err != nil {
		t.Error("create item in table <account_statements> fail:", err)
	}

	err = _do.Save(&financial_servicev1.AccountStatementsORM{})
	if err != nil {
		t.Error("create item in table <account_statements> fail:", err)
	}

	err = _do.CreateInBatches([]*financial_servicev1.AccountStatementsORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <account_statements> fail:", err)
	}

	_, err = _do.Select(accountStatementsORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <account_statements> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <account_statements> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <account_statements> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <account_statements> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*financial_servicev1.AccountStatementsORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <account_statements> fail:", err)
	}

	_, err = _do.Select(accountStatementsORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <account_statements> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <account_statements> fail:", err)
	}

	_, err = _do.Select(accountStatementsORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <account_statements> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <account_statements> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <account_statements> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <account_statements> fail:", err)
	}

	_, err = _do.ScanByPage(&financial_servicev1.AccountStatementsORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <account_statements> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <account_statements> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <account_statements> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <account_statements> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <account_statements> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <account_statements> fail:", err)
	}
}

var AccountStatementsORMGetRecordByIDTestCase = []TestCase{}

func Test_accountStatementsORM_GetRecordByID(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountStatementsORMGetRecordByIDsTestCase = []TestCase{}

func Test_accountStatementsORM_GetRecordByIDs(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountStatementsORMCreateRecordTestCase = []TestCase{}

func Test_accountStatementsORM_CreateRecord(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(financial_servicev1.AccountStatementsORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var AccountStatementsORMUpdateRecordByIDTestCase = []TestCase{}

func Test_accountStatementsORM_UpdateRecordByID(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(financial_servicev1.AccountStatementsORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var AccountStatementsORMDeleteRecordByIDTestCase = []TestCase{}

func Test_accountStatementsORM_DeleteRecordByID(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var AccountStatementsORMGetAllRecordsTestCase = []TestCase{}

func Test_accountStatementsORM_GetAllRecords(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountStatementsORMCountAllTestCase = []TestCase{}

func Test_accountStatementsORM_CountAll(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountStatementsORMGetByIDTestCase = []TestCase{}

func Test_accountStatementsORM_GetByID(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var AccountStatementsORMGetByIDsTestCase = []TestCase{}

func Test_accountStatementsORM_GetByIDs(t *testing.T) {
	accountStatementsORM := newAccountStatementsORM(db)
	do := accountStatementsORM.WithContext(context.Background()).Debug()

	for i, tt := range AccountStatementsORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
