// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	social_servicev2 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/social_service/v2"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&social_servicev2.TopicORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&social_servicev2.TopicORM{}) fail: %s", err)
	}
}

func Test_topicORMQuery(t *testing.T) {
	topicORM := newTopicORM(db)
	topicORM = *topicORM.As(topicORM.TableName())
	_do := topicORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(topicORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <topics> fail:", err)
		return
	}

	_, ok := topicORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from topicORM success")
	}

	err = _do.Create(&social_servicev2.TopicORM{})
	if err != nil {
		t.Error("create item in table <topics> fail:", err)
	}

	err = _do.Save(&social_servicev2.TopicORM{})
	if err != nil {
		t.Error("create item in table <topics> fail:", err)
	}

	err = _do.CreateInBatches([]*social_servicev2.TopicORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <topics> fail:", err)
	}

	_, err = _do.Select(topicORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <topics> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <topics> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <topics> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <topics> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*social_servicev2.TopicORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <topics> fail:", err)
	}

	_, err = _do.Select(topicORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <topics> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <topics> fail:", err)
	}

	_, err = _do.Select(topicORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <topics> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <topics> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <topics> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <topics> fail:", err)
	}

	_, err = _do.ScanByPage(&social_servicev2.TopicORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <topics> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <topics> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <topics> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <topics> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <topics> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <topics> fail:", err)
	}
}

var TopicORMGetRecordByIDTestCase = []TestCase{}

func Test_topicORM_GetRecordByID(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMGetRecordByIDTestCase {
		t.Run("GetRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByID(tt.Input.Args[0].(int))
			assert(t, "GetRecordByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TopicORMGetRecordByIDsTestCase = []TestCase{}

func Test_topicORM_GetRecordByIDs(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMGetRecordByIDsTestCase {
		t.Run("GetRecordByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetRecordByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetRecordByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetRecordByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}

var TopicORMCreateRecordTestCase = []TestCase{}

func Test_topicORM_CreateRecord(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMCreateRecordTestCase {
		t.Run("CreateRecord_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.CreateRecord(tt.Input.Args[0].(social_servicev2.TopicORM))
			assert(t, "CreateRecord", res1, tt.Expectation.Ret[0])
		})
	}
}

var TopicORMUpdateRecordByIDTestCase = []TestCase{}

func Test_topicORM_UpdateRecordByID(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMUpdateRecordByIDTestCase {
		t.Run("UpdateRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.UpdateRecordByID(tt.Input.Args[0].(int), tt.Input.Args[1].(social_servicev2.TopicORM))
			assert(t, "UpdateRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var TopicORMDeleteRecordByIDTestCase = []TestCase{}

func Test_topicORM_DeleteRecordByID(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMDeleteRecordByIDTestCase {
		t.Run("DeleteRecordByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1 := do.DeleteRecordByID(tt.Input.Args[0].(int))
			assert(t, "DeleteRecordByID", res1, tt.Expectation.Ret[0])
		})
	}
}

var TopicORMGetAllRecordsTestCase = []TestCase{}

func Test_topicORM_GetAllRecords(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMGetAllRecordsTestCase {
		t.Run("GetAllRecords_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetAllRecords(tt.Input.Args[0].(string), tt.Input.Args[1].(int), tt.Input.Args[2].(int))
			assert(t, "GetAllRecords", res1, tt.Expectation.Ret[0])
			assert(t, "GetAllRecords", res2, tt.Expectation.Ret[1])
		})
	}
}

var TopicORMCountAllTestCase = []TestCase{}

func Test_topicORM_CountAll(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMCountAllTestCase {
		t.Run("CountAll_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.CountAll()
			assert(t, "CountAll", res1, tt.Expectation.Ret[0])
			assert(t, "CountAll", res2, tt.Expectation.Ret[1])
		})
	}
}

var TopicORMGetByIDTestCase = []TestCase{}

func Test_topicORM_GetByID(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(uint64))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TopicORMGetByIDsTestCase = []TestCase{}

func Test_topicORM_GetByIDs(t *testing.T) {
	topicORM := newTopicORM(db)
	do := topicORM.WithContext(context.Background()).Debug()

	for i, tt := range TopicORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]uint64))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}