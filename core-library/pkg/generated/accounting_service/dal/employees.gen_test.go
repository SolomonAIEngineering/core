// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	accounting_servicev1 "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/accounting_service/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&accounting_servicev1.EmployeeORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&accounting_servicev1.EmployeeORM{}) fail: %s", err)
	}
}

func Test_employeeORMQuery(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	employeeORM = *employeeORM.As(employeeORM.TableName())
	_do := employeeORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(employeeORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <employees> fail:", err)
		return
	}

	_, ok := employeeORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from employeeORM success")
	}

	err = _do.Create(&accounting_servicev1.EmployeeORM{})
	if err != nil {
		t.Error("create item in table <employees> fail:", err)
	}

	err = _do.Save(&accounting_servicev1.EmployeeORM{})
	if err != nil {
		t.Error("create item in table <employees> fail:", err)
	}

	err = _do.CreateInBatches([]*accounting_servicev1.EmployeeORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <employees> fail:", err)
	}

	_, err = _do.Select(employeeORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <employees> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <employees> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <employees> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <employees> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*accounting_servicev1.EmployeeORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <employees> fail:", err)
	}

	_, err = _do.Select(employeeORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <employees> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <employees> fail:", err)
	}

	_, err = _do.Select(employeeORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <employees> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <employees> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <employees> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <employees> fail:", err)
	}

	_, err = _do.ScanByPage(&accounting_servicev1.EmployeeORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <employees> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <employees> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <employees> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <employees> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <employees> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <employees> fail:", err)
	}
}

var EmployeeORMGetRecordByIDTestCase = []TestCase{}

func Test_employeeORM_GetRecordByID(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var EmployeeORMGetRecordByIDsTestCase = []TestCase{}

func Test_employeeORM_GetRecordByIDs(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var EmployeeORMCreateRecordTestCase = []TestCase{}

func Test_employeeORM_CreateRecord(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(accounting_servicev1.EmployeeORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var EmployeeORMUpdateRecordByIDTestCase = []TestCase{}

func Test_employeeORM_UpdateRecordByID(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(accounting_servicev1.EmployeeORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var EmployeeORMDeleteRecordByIDTestCase = []TestCase{}

func Test_employeeORM_DeleteRecordByID(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var EmployeeORMGetAllRecordsTestCase = []TestCase{}

func Test_employeeORM_GetAllRecords(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var EmployeeORMCountAllTestCase = []TestCase{}

func Test_employeeORM_CountAll(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var EmployeeORMGetByIDTestCase = []TestCase{}

func Test_employeeORM_GetByID(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var EmployeeORMGetByIDsTestCase = []TestCase{}

func Test_employeeORM_GetByIDs(t *testing.T) {
	employeeORM := newEmployeeORM(db)
	do := employeeORM.WithContext(context.Background()).Debug()

	for i, tt := range EmployeeORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
