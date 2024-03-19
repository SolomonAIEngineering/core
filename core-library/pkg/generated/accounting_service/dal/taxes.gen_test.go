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
	err := db.AutoMigrate(&accounting_servicev1.TaxORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&accounting_servicev1.TaxORM{}) fail: %s", err)
	}
}

func Test_taxORMQuery(t *testing.T) {
	taxORM := newTaxORM(db)
	taxORM = *taxORM.As(taxORM.TableName())
	_do := taxORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(taxORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <taxes> fail:", err)
		return
	}

	_, ok := taxORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from taxORM success")
	}

	err = _do.Create(&accounting_servicev1.TaxORM{})
	if err != nil {
		t.Error("create item in table <taxes> fail:", err)
	}

	err = _do.Save(&accounting_servicev1.TaxORM{})
	if err != nil {
		t.Error("create item in table <taxes> fail:", err)
	}

	err = _do.CreateInBatches([]*accounting_servicev1.TaxORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <taxes> fail:", err)
	}

	_, err = _do.Select(taxORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <taxes> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <taxes> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <taxes> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <taxes> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*accounting_servicev1.TaxORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <taxes> fail:", err)
	}

	_, err = _do.Select(taxORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <taxes> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <taxes> fail:", err)
	}

	_, err = _do.Select(taxORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <taxes> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <taxes> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <taxes> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <taxes> fail:", err)
	}

	_, err = _do.ScanByPage(&accounting_servicev1.TaxORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <taxes> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <taxes> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <taxes> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <taxes> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <taxes> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <taxes> fail:", err)
	}
}

var TaxORMGetRecordByIDTestCase = []TestCase{}

func Test_taxORM_GetRecordByID(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxORMGetRecordByIDsTestCase = []TestCase{}

func Test_taxORM_GetRecordByIDs(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxORMCreateRecordTestCase = []TestCase{}

func Test_taxORM_CreateRecord(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(accounting_servicev1.TaxORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var TaxORMUpdateRecordByIDTestCase = []TestCase{}

func Test_taxORM_UpdateRecordByID(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(accounting_servicev1.TaxORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var TaxORMDeleteRecordByIDTestCase = []TestCase{}

func Test_taxORM_DeleteRecordByID(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var TaxORMGetAllRecordsTestCase = []TestCase{}

func Test_taxORM_GetAllRecords(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxORMCountAllTestCase = []TestCase{}

func Test_taxORM_CountAll(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxORMGetByIDTestCase = []TestCase{}

func Test_taxORM_GetByID(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxORMGetByIDsTestCase = []TestCase{}

func Test_taxORM_GetByIDs(t *testing.T) {
	taxORM := newTaxORM(db)
	do := taxORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}