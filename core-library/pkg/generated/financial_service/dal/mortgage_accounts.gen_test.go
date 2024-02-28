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
	err := db.AutoMigrate(&financial_servicev1.MortgageAccountORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&financial_servicev1.MortgageAccountORM{}) fail: %s", err)
	}
}

func Test_mortgageAccountORMQuery(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	mortgageAccountORM = *mortgageAccountORM.As(mortgageAccountORM.TableName())
	_do := mortgageAccountORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(mortgageAccountORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <mortgage_accounts> fail:", err)
		return
	}

	_, ok := mortgageAccountORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from mortgageAccountORM success")
	}

	err = _do.Create(&financial_servicev1.MortgageAccountORM{})
	if err != nil {
		t.Error("create item in table <mortgage_accounts> fail:", err)
	}

	err = _do.Save(&financial_servicev1.MortgageAccountORM{})
	if err != nil {
		t.Error("create item in table <mortgage_accounts> fail:", err)
	}

	err = _do.CreateInBatches([]*financial_servicev1.MortgageAccountORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Select(mortgageAccountORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <mortgage_accounts> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*financial_servicev1.MortgageAccountORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Select(mortgageAccountORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Select(mortgageAccountORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <mortgage_accounts> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.ScanByPage(&financial_servicev1.MortgageAccountORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <mortgage_accounts> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <mortgage_accounts> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <mortgage_accounts> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <mortgage_accounts> fail:", err)
	}
}

var MortgageAccountORMGetRecordByIDTestCase = []TestCase{}

func Test_mortgageAccountORM_GetRecordByID(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var MortgageAccountORMGetRecordByIDsTestCase = []TestCase{}

func Test_mortgageAccountORM_GetRecordByIDs(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var MortgageAccountORMCreateRecordTestCase = []TestCase{}

func Test_mortgageAccountORM_CreateRecord(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(financial_servicev1.MortgageAccountORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var MortgageAccountORMUpdateRecordByIDTestCase = []TestCase{}

func Test_mortgageAccountORM_UpdateRecordByID(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(financial_servicev1.MortgageAccountORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var MortgageAccountORMDeleteRecordByIDTestCase = []TestCase{}

func Test_mortgageAccountORM_DeleteRecordByID(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var MortgageAccountORMGetAllRecordsTestCase = []TestCase{}

func Test_mortgageAccountORM_GetAllRecords(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var MortgageAccountORMCountAllTestCase = []TestCase{}

func Test_mortgageAccountORM_CountAll(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var MortgageAccountORMGetByIDTestCase = []TestCase{}

func Test_mortgageAccountORM_GetByID(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var MortgageAccountORMGetByIDsTestCase = []TestCase{}

func Test_mortgageAccountORM_GetByIDs(t *testing.T) {
	mortgageAccountORM := newMortgageAccountORM(db)
	do := mortgageAccountORM.WithContext(context.Background()).Debug()

	for i, tt := range MortgageAccountORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
