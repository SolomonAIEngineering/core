// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	user_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/user_service/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&user_servicev1.SettingsORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&user_servicev1.SettingsORM{}) fail: %s", err)
	}
}

func Test_settingsORMQuery(t *testing.T) {
	settingsORM := newSettingsORM(db)
	settingsORM = *settingsORM.As(settingsORM.TableName())
	_do := settingsORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(settingsORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <settings> fail:", err)
		return
	}

	_, ok := settingsORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from settingsORM success")
	}

	err = _do.Create(&user_servicev1.SettingsORM{})
	if err != nil {
		t.Error("create item in table <settings> fail:", err)
	}

	err = _do.Save(&user_servicev1.SettingsORM{})
	if err != nil {
		t.Error("create item in table <settings> fail:", err)
	}

	err = _do.CreateInBatches([]*user_servicev1.SettingsORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <settings> fail:", err)
	}

	_, err = _do.Select(settingsORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <settings> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <settings> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <settings> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <settings> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*user_servicev1.SettingsORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <settings> fail:", err)
	}

	_, err = _do.Select(settingsORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <settings> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <settings> fail:", err)
	}

	_, err = _do.Select(settingsORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <settings> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <settings> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <settings> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <settings> fail:", err)
	}

	_, err = _do.ScanByPage(&user_servicev1.SettingsORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <settings> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <settings> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <settings> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <settings> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <settings> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <settings> fail:", err)
	}
}

var SettingsORMGetRecordByIDTestCase = []TestCase{}

func Test_settingsORM_GetRecordByID(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var SettingsORMGetRecordByIDsTestCase = []TestCase{}

func Test_settingsORM_GetRecordByIDs(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var SettingsORMCreateRecordTestCase = []TestCase{}

func Test_settingsORM_CreateRecord(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(user_servicev1.SettingsORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var SettingsORMUpdateRecordByIDTestCase = []TestCase{}

func Test_settingsORM_UpdateRecordByID(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(user_servicev1.SettingsORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var SettingsORMDeleteRecordByIDTestCase = []TestCase{}

func Test_settingsORM_DeleteRecordByID(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var SettingsORMGetAllRecordsTestCase = []TestCase{}

func Test_settingsORM_GetAllRecords(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var SettingsORMCountAllTestCase = []TestCase{}

func Test_settingsORM_CountAll(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var SettingsORMGetByIDTestCase = []TestCase{}

func Test_settingsORM_GetByID(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var SettingsORMGetByIDsTestCase = []TestCase{}

func Test_settingsORM_GetByIDs(t *testing.T) {
	settingsORM := newSettingsORM(db)
	do := settingsORM.WithContext(context.Background()).Debug()

	for i, tt := range SettingsORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}