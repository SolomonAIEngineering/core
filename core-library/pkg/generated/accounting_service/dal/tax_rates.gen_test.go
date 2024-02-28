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
	err := db.AutoMigrate(&accounting_servicev1.TaxRateORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&accounting_servicev1.TaxRateORM{}) fail: %s", err)
	}
}

func Test_taxRateORMQuery(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	taxRateORM = *taxRateORM.As(taxRateORM.TableName())
	_do := taxRateORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(taxRateORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <tax_rates> fail:", err)
		return
	}

	_, ok := taxRateORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from taxRateORM success")
	}

	err = _do.Create(&accounting_servicev1.TaxRateORM{})
	if err != nil {
		t.Error("create item in table <tax_rates> fail:", err)
	}

	err = _do.Save(&accounting_servicev1.TaxRateORM{})
	if err != nil {
		t.Error("create item in table <tax_rates> fail:", err)
	}

	err = _do.CreateInBatches([]*accounting_servicev1.TaxRateORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <tax_rates> fail:", err)
	}

	_, err = _do.Select(taxRateORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <tax_rates> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <tax_rates> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <tax_rates> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <tax_rates> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*accounting_servicev1.TaxRateORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <tax_rates> fail:", err)
	}

	_, err = _do.Select(taxRateORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <tax_rates> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <tax_rates> fail:", err)
	}

	_, err = _do.Select(taxRateORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <tax_rates> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <tax_rates> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <tax_rates> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <tax_rates> fail:", err)
	}

	_, err = _do.ScanByPage(&accounting_servicev1.TaxRateORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <tax_rates> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <tax_rates> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <tax_rates> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <tax_rates> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <tax_rates> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <tax_rates> fail:", err)
	}
}

var TaxRateORMGetRecordByIDTestCase = []TestCase{}

func Test_taxRateORM_GetRecordByID(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxRateORMGetRecordByIDsTestCase = []TestCase{}

func Test_taxRateORM_GetRecordByIDs(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxRateORMCreateRecordTestCase = []TestCase{}

func Test_taxRateORM_CreateRecord(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(accounting_servicev1.TaxRateORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var TaxRateORMUpdateRecordByIDTestCase = []TestCase{}

func Test_taxRateORM_UpdateRecordByID(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(accounting_servicev1.TaxRateORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var TaxRateORMDeleteRecordByIDTestCase = []TestCase{}

func Test_taxRateORM_DeleteRecordByID(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var TaxRateORMGetAllRecordsTestCase = []TestCase{}

func Test_taxRateORM_GetAllRecords(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxRateORMCountAllTestCase = []TestCase{}

func Test_taxRateORM_CountAll(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxRateORMGetByIDTestCase = []TestCase{}

func Test_taxRateORM_GetByID(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TaxRateORMGetByIDsTestCase = []TestCase{}

func Test_taxRateORM_GetByIDs(t *testing.T) {
	taxRateORM := newTaxRateORM(db)
	do := taxRateORM.WithContext(context.Background()).Debug()

	for i, tt := range TaxRateORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
