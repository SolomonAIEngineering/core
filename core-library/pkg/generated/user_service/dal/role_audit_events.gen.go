// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	user_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/user_service/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newRoleAuditEventsORM(db *gorm.DB, opts ...gen.DOOption) roleAuditEventsORM {
	_roleAuditEventsORM := roleAuditEventsORM{}

	_roleAuditEventsORM.roleAuditEventsORMDo.UseDB(db, opts...)
	_roleAuditEventsORM.roleAuditEventsORMDo.UseModel(&user_servicev1.RoleAuditEventsORM{})

	tableName := _roleAuditEventsORM.roleAuditEventsORMDo.TableName()
	_roleAuditEventsORM.ALL = field.NewAsterisk(tableName)
	_roleAuditEventsORM.Action = field.NewString(tableName, "action")
	_roleAuditEventsORM.AffectedFields = field.NewField(tableName, "affected_fields")
	_roleAuditEventsORM.ClientIp = field.NewString(tableName, "client_ip")
	_roleAuditEventsORM.Context = field.NewString(tableName, "context")
	_roleAuditEventsORM.Id = field.NewInt64(tableName, "id")
	_roleAuditEventsORM.PerformedBy = field.NewString(tableName, "performed_by")
	_roleAuditEventsORM.PreviousValues = field.NewField(tableName, "previous_values")
	_roleAuditEventsORM.RoleId = field.NewInt64(tableName, "role_id")
	_roleAuditEventsORM.Timestamp = field.NewTime(tableName, "timestamp")
	_roleAuditEventsORM.UserAgent = field.NewString(tableName, "user_agent")

	_roleAuditEventsORM.fillFieldMap()

	return _roleAuditEventsORM
}

type roleAuditEventsORM struct {
	roleAuditEventsORMDo

	ALL            field.Asterisk
	Action         field.String
	AffectedFields field.Field
	ClientIp       field.String
	Context        field.String
	Id             field.Int64
	PerformedBy    field.String
	PreviousValues field.Field
	RoleId         field.Int64
	Timestamp      field.Time
	UserAgent      field.String

	fieldMap map[string]field.Expr
}

func (r roleAuditEventsORM) Table(newTableName string) *roleAuditEventsORM {
	r.roleAuditEventsORMDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r roleAuditEventsORM) As(alias string) *roleAuditEventsORM {
	r.roleAuditEventsORMDo.DO = *(r.roleAuditEventsORMDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *roleAuditEventsORM) updateTableName(table string) *roleAuditEventsORM {
	r.ALL = field.NewAsterisk(table)
	r.Action = field.NewString(table, "action")
	r.AffectedFields = field.NewField(table, "affected_fields")
	r.ClientIp = field.NewString(table, "client_ip")
	r.Context = field.NewString(table, "context")
	r.Id = field.NewInt64(table, "id")
	r.PerformedBy = field.NewString(table, "performed_by")
	r.PreviousValues = field.NewField(table, "previous_values")
	r.RoleId = field.NewInt64(table, "role_id")
	r.Timestamp = field.NewTime(table, "timestamp")
	r.UserAgent = field.NewString(table, "user_agent")

	r.fillFieldMap()

	return r
}

func (r *roleAuditEventsORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *roleAuditEventsORM) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 10)
	r.fieldMap["action"] = r.Action
	r.fieldMap["affected_fields"] = r.AffectedFields
	r.fieldMap["client_ip"] = r.ClientIp
	r.fieldMap["context"] = r.Context
	r.fieldMap["id"] = r.Id
	r.fieldMap["performed_by"] = r.PerformedBy
	r.fieldMap["previous_values"] = r.PreviousValues
	r.fieldMap["role_id"] = r.RoleId
	r.fieldMap["timestamp"] = r.Timestamp
	r.fieldMap["user_agent"] = r.UserAgent
}

func (r roleAuditEventsORM) clone(db *gorm.DB) roleAuditEventsORM {
	r.roleAuditEventsORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r roleAuditEventsORM) replaceDB(db *gorm.DB) roleAuditEventsORM {
	r.roleAuditEventsORMDo.ReplaceDB(db)
	return r
}

type roleAuditEventsORMDo struct{ gen.DO }

type IRoleAuditEventsORMDo interface {
	gen.SubQuery
	Debug() IRoleAuditEventsORMDo
	WithContext(ctx context.Context) IRoleAuditEventsORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoleAuditEventsORMDo
	WriteDB() IRoleAuditEventsORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoleAuditEventsORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoleAuditEventsORMDo
	Not(conds ...gen.Condition) IRoleAuditEventsORMDo
	Or(conds ...gen.Condition) IRoleAuditEventsORMDo
	Select(conds ...field.Expr) IRoleAuditEventsORMDo
	Where(conds ...gen.Condition) IRoleAuditEventsORMDo
	Order(conds ...field.Expr) IRoleAuditEventsORMDo
	Distinct(cols ...field.Expr) IRoleAuditEventsORMDo
	Omit(cols ...field.Expr) IRoleAuditEventsORMDo
	Join(table schema.Tabler, on ...field.Expr) IRoleAuditEventsORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoleAuditEventsORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoleAuditEventsORMDo
	Group(cols ...field.Expr) IRoleAuditEventsORMDo
	Having(conds ...gen.Condition) IRoleAuditEventsORMDo
	Limit(limit int) IRoleAuditEventsORMDo
	Offset(offset int) IRoleAuditEventsORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleAuditEventsORMDo
	Unscoped() IRoleAuditEventsORMDo
	Create(values ...*user_servicev1.RoleAuditEventsORM) error
	CreateInBatches(values []*user_servicev1.RoleAuditEventsORM, batchSize int) error
	Save(values ...*user_servicev1.RoleAuditEventsORM) error
	First() (*user_servicev1.RoleAuditEventsORM, error)
	Take() (*user_servicev1.RoleAuditEventsORM, error)
	Last() (*user_servicev1.RoleAuditEventsORM, error)
	Find() ([]*user_servicev1.RoleAuditEventsORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*user_servicev1.RoleAuditEventsORM, err error)
	FindInBatches(result *[]*user_servicev1.RoleAuditEventsORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*user_servicev1.RoleAuditEventsORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoleAuditEventsORMDo
	Assign(attrs ...field.AssignExpr) IRoleAuditEventsORMDo
	Joins(fields ...field.RelationField) IRoleAuditEventsORMDo
	Preload(fields ...field.RelationField) IRoleAuditEventsORMDo
	FirstOrInit() (*user_servicev1.RoleAuditEventsORM, error)
	FirstOrCreate() (*user_servicev1.RoleAuditEventsORM, error)
	FindByPage(offset int, limit int) (result []*user_servicev1.RoleAuditEventsORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoleAuditEventsORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result user_servicev1.RoleAuditEventsORM, err error)
	GetRecordByIDs(ids []int) (result []user_servicev1.RoleAuditEventsORM, err error)
	CreateRecord(item user_servicev1.RoleAuditEventsORM) (err error)
	UpdateRecordByID(id int, item user_servicev1.RoleAuditEventsORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []user_servicev1.RoleAuditEventsORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result user_servicev1.RoleAuditEventsORM, err error)
	GetByIDs(ids []uint64) (result []user_servicev1.RoleAuditEventsORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (r roleAuditEventsORMDo) GetRecordByID(id int) (result user_servicev1.RoleAuditEventsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM role_audit_events ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (r roleAuditEventsORMDo) GetRecordByIDs(ids []int) (result []user_servicev1.RoleAuditEventsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM role_audit_events ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (r roleAuditEventsORMDo) CreateRecord(item user_servicev1.RoleAuditEventsORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO role_audit_events (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (r roleAuditEventsORMDo) UpdateRecordByID(id int, item user_servicev1.RoleAuditEventsORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE role_audit_events SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (r roleAuditEventsORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM role_audit_events ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (r roleAuditEventsORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []user_servicev1.RoleAuditEventsORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM role_audit_events ORDER BY " + r.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (r roleAuditEventsORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM role_audit_events ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (r roleAuditEventsORMDo) GetByID(id uint64) (result user_servicev1.RoleAuditEventsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM role_audit_events ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (r roleAuditEventsORMDo) GetByIDs(ids []uint64) (result []user_servicev1.RoleAuditEventsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM role_audit_events ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (r roleAuditEventsORMDo) Debug() IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Debug())
}

func (r roleAuditEventsORMDo) WithContext(ctx context.Context) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleAuditEventsORMDo) ReadDB() IRoleAuditEventsORMDo {
	return r.Clauses(dbresolver.Read)
}

func (r roleAuditEventsORMDo) WriteDB() IRoleAuditEventsORMDo {
	return r.Clauses(dbresolver.Write)
}

func (r roleAuditEventsORMDo) Session(config *gorm.Session) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Session(config))
}

func (r roleAuditEventsORMDo) Clauses(conds ...clause.Expression) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleAuditEventsORMDo) Returning(value interface{}, columns ...string) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roleAuditEventsORMDo) Not(conds ...gen.Condition) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleAuditEventsORMDo) Or(conds ...gen.Condition) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleAuditEventsORMDo) Select(conds ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleAuditEventsORMDo) Where(conds ...gen.Condition) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleAuditEventsORMDo) Order(conds ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleAuditEventsORMDo) Distinct(cols ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleAuditEventsORMDo) Omit(cols ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleAuditEventsORMDo) Join(table schema.Tabler, on ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleAuditEventsORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleAuditEventsORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleAuditEventsORMDo) Group(cols ...field.Expr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleAuditEventsORMDo) Having(conds ...gen.Condition) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleAuditEventsORMDo) Limit(limit int) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleAuditEventsORMDo) Offset(offset int) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleAuditEventsORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleAuditEventsORMDo) Unscoped() IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleAuditEventsORMDo) Create(values ...*user_servicev1.RoleAuditEventsORM) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleAuditEventsORMDo) CreateInBatches(values []*user_servicev1.RoleAuditEventsORM, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleAuditEventsORMDo) Save(values ...*user_servicev1.RoleAuditEventsORM) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleAuditEventsORMDo) First() (*user_servicev1.RoleAuditEventsORM, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*user_servicev1.RoleAuditEventsORM), nil
	}
}

func (r roleAuditEventsORMDo) Take() (*user_servicev1.RoleAuditEventsORM, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*user_servicev1.RoleAuditEventsORM), nil
	}
}

func (r roleAuditEventsORMDo) Last() (*user_servicev1.RoleAuditEventsORM, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*user_servicev1.RoleAuditEventsORM), nil
	}
}

func (r roleAuditEventsORMDo) Find() ([]*user_servicev1.RoleAuditEventsORM, error) {
	result, err := r.DO.Find()
	return result.([]*user_servicev1.RoleAuditEventsORM), err
}

func (r roleAuditEventsORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*user_servicev1.RoleAuditEventsORM, err error) {
	buf := make([]*user_servicev1.RoleAuditEventsORM, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleAuditEventsORMDo) FindInBatches(result *[]*user_servicev1.RoleAuditEventsORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleAuditEventsORMDo) Attrs(attrs ...field.AssignExpr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleAuditEventsORMDo) Assign(attrs ...field.AssignExpr) IRoleAuditEventsORMDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleAuditEventsORMDo) Joins(fields ...field.RelationField) IRoleAuditEventsORMDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roleAuditEventsORMDo) Preload(fields ...field.RelationField) IRoleAuditEventsORMDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roleAuditEventsORMDo) FirstOrInit() (*user_servicev1.RoleAuditEventsORM, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*user_servicev1.RoleAuditEventsORM), nil
	}
}

func (r roleAuditEventsORMDo) FirstOrCreate() (*user_servicev1.RoleAuditEventsORM, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*user_servicev1.RoleAuditEventsORM), nil
	}
}

func (r roleAuditEventsORMDo) FindByPage(offset int, limit int) (result []*user_servicev1.RoleAuditEventsORM, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roleAuditEventsORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roleAuditEventsORMDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roleAuditEventsORMDo) Delete(models ...*user_servicev1.RoleAuditEventsORM) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roleAuditEventsORMDo) withDO(do gen.Dao) *roleAuditEventsORMDo {
	r.DO = *do.(*gen.DO)
	return r
}
