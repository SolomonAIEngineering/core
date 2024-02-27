// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	financial_servicev1 "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/financial_service/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&financial_servicev1.CategoryORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&financial_servicev1.CategoryORM{}) fail: %s", err)
	}
}

func Test_categoryORMQuery(t *testing.T) {
	categoryORM := newCategoryORM(db)
	categoryORM = *categoryORM.As(categoryORM.TableName())
	_do := categoryORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(categoryORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <categories> fail:", err)
		return
	}

	_, ok := categoryORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from categoryORM success")
	}

	err = _do.Create(&financial_servicev1.CategoryORM{})
	if err != nil {
		t.Error("create item in table <categories> fail:", err)
	}

	err = _do.Save(&financial_servicev1.CategoryORM{})
	if err != nil {
		t.Error("create item in table <categories> fail:", err)
	}

	err = _do.CreateInBatches([]*financial_servicev1.CategoryORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <categories> fail:", err)
	}

	_, err = _do.Select(categoryORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <categories> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <categories> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <categories> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <categories> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*financial_servicev1.CategoryORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <categories> fail:", err)
	}

	_, err = _do.Select(categoryORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <categories> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <categories> fail:", err)
	}

	_, err = _do.Select(categoryORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <categories> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <categories> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <categories> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <categories> fail:", err)
	}

	_, err = _do.ScanByPage(&financial_servicev1.CategoryORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <categories> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <categories> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <categories> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <categories> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <categories> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <categories> fail:", err)
	}
}

var CategoryORMGetRecordByIDTestCase = []TestCase{}

func Test_categoryORM_GetRecordByID(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var CategoryORMGetRecordByIDsTestCase = []TestCase{}

func Test_categoryORM_GetRecordByIDs(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var CategoryORMCreateRecordTestCase = []TestCase{}

func Test_categoryORM_CreateRecord(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(financial_servicev1.CategoryORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var CategoryORMUpdateRecordByIDTestCase = []TestCase{}

func Test_categoryORM_UpdateRecordByID(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(financial_servicev1.CategoryORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var CategoryORMDeleteRecordByIDTestCase = []TestCase{}

func Test_categoryORM_DeleteRecordByID(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var CategoryORMGetAllRecordsTestCase = []TestCase{}

func Test_categoryORM_GetAllRecords(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var CategoryORMCountAllTestCase = []TestCase{}

func Test_categoryORM_CountAll(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var CategoryORMGetByIDTestCase = []TestCase{}

func Test_categoryORM_GetByID(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var CategoryORMGetByIDsTestCase = []TestCase{}

func Test_categoryORM_GetByIDs(t *testing.T) {
	categoryORM := newCategoryORM(db)
	do := categoryORM.WithContext(context.Background()).Debug()

	for i, tt := range CategoryORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
