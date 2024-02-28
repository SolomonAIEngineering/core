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
	err := db.AutoMigrate(&accounting_servicev1.DeductionORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&accounting_servicev1.DeductionORM{}) fail: %s", err)
	}
}

func Test_deductionORMQuery(t *testing.T) {
	deductionORM := newDeductionORM(db)
	deductionORM = *deductionORM.As(deductionORM.TableName())
	_do := deductionORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(deductionORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <deductions> fail:", err)
		return
	}

	_, ok := deductionORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from deductionORM success")
	}

	err = _do.Create(&accounting_servicev1.DeductionORM{})
	if err != nil {
		t.Error("create item in table <deductions> fail:", err)
	}

	err = _do.Save(&accounting_servicev1.DeductionORM{})
	if err != nil {
		t.Error("create item in table <deductions> fail:", err)
	}

	err = _do.CreateInBatches([]*accounting_servicev1.DeductionORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <deductions> fail:", err)
	}

	_, err = _do.Select(deductionORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <deductions> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <deductions> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <deductions> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <deductions> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*accounting_servicev1.DeductionORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <deductions> fail:", err)
	}

	_, err = _do.Select(deductionORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <deductions> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <deductions> fail:", err)
	}

	_, err = _do.Select(deductionORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <deductions> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <deductions> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <deductions> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <deductions> fail:", err)
	}

	_, err = _do.ScanByPage(&accounting_servicev1.DeductionORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <deductions> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <deductions> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <deductions> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <deductions> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <deductions> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <deductions> fail:", err)
	}
}

var DeductionORMGetRecordByIDTestCase = []TestCase{}

func Test_deductionORM_GetRecordByID(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var DeductionORMGetRecordByIDsTestCase = []TestCase{}

func Test_deductionORM_GetRecordByIDs(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var DeductionORMCreateRecordTestCase = []TestCase{}

func Test_deductionORM_CreateRecord(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(accounting_servicev1.DeductionORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var DeductionORMUpdateRecordByIDTestCase = []TestCase{}

func Test_deductionORM_UpdateRecordByID(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(accounting_servicev1.DeductionORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var DeductionORMDeleteRecordByIDTestCase = []TestCase{}

func Test_deductionORM_DeleteRecordByID(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var DeductionORMGetAllRecordsTestCase = []TestCase{}

func Test_deductionORM_GetAllRecords(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var DeductionORMCountAllTestCase = []TestCase{}

func Test_deductionORM_CountAll(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var DeductionORMGetByIDTestCase = []TestCase{}

func Test_deductionORM_GetByID(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var DeductionORMGetByIDsTestCase = []TestCase{}

func Test_deductionORM_GetByIDs(t *testing.T) {
	deductionORM := newDeductionORM(db)
	do := deductionORM.WithContext(context.Background()).Debug()

	for i, tt := range DeductionORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
