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
	err := db.AutoMigrate(&accounting_servicev1.BusinessActionableInsightORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&accounting_servicev1.BusinessActionableInsightORM{}) fail: %s", err)
	}
}

func Test_businessActionableInsightORMQuery(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	businessActionableInsightORM = *businessActionableInsightORM.As(businessActionableInsightORM.TableName())
	_do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(businessActionableInsightORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <business_actionable_insights> fail:", err)
		return
	}

	_, ok := businessActionableInsightORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from businessActionableInsightORM success")
	}

	err = _do.Create(&accounting_servicev1.BusinessActionableInsightORM{})
	if err != nil {
		t.Error("create item in table <business_actionable_insights> fail:", err)
	}

	err = _do.Save(&accounting_servicev1.BusinessActionableInsightORM{})
	if err != nil {
		t.Error("create item in table <business_actionable_insights> fail:", err)
	}

	err = _do.CreateInBatches([]*accounting_servicev1.BusinessActionableInsightORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Select(businessActionableInsightORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <business_actionable_insights> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*accounting_servicev1.BusinessActionableInsightORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Select(businessActionableInsightORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Select(businessActionableInsightORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <business_actionable_insights> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.ScanByPage(&accounting_servicev1.BusinessActionableInsightORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <business_actionable_insights> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <business_actionable_insights> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <business_actionable_insights> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <business_actionable_insights> fail:", err)
	}
}

var BusinessActionableInsightORMGetRecordByIDTestCase = []TestCase{}

func Test_businessActionableInsightORM_GetRecordByID(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var BusinessActionableInsightORMGetRecordByIDsTestCase = []TestCase{}

func Test_businessActionableInsightORM_GetRecordByIDs(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var BusinessActionableInsightORMCreateRecordTestCase = []TestCase{}

func Test_businessActionableInsightORM_CreateRecord(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(accounting_servicev1.BusinessActionableInsightORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var BusinessActionableInsightORMUpdateRecordByIDTestCase = []TestCase{}

func Test_businessActionableInsightORM_UpdateRecordByID(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(accounting_servicev1.BusinessActionableInsightORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var BusinessActionableInsightORMDeleteRecordByIDTestCase = []TestCase{}

func Test_businessActionableInsightORM_DeleteRecordByID(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var BusinessActionableInsightORMGetAllRecordsTestCase = []TestCase{}

func Test_businessActionableInsightORM_GetAllRecords(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var BusinessActionableInsightORMCountAllTestCase = []TestCase{}

func Test_businessActionableInsightORM_CountAll(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var BusinessActionableInsightORMGetByIDTestCase = []TestCase{}

func Test_businessActionableInsightORM_GetByID(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var BusinessActionableInsightORMGetByIDsTestCase = []TestCase{}

func Test_businessActionableInsightORM_GetByIDs(t *testing.T) {
	businessActionableInsightORM := newBusinessActionableInsightORM(db)
	do := businessActionableInsightORM.WithContext(context.Background()).Debug()

	for i, tt := range BusinessActionableInsightORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
