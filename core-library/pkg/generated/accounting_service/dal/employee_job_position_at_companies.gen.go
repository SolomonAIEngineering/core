// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	accounting_servicev1 "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/accounting_service/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newEmployeeJobPositionAtCompanyORM(db *gorm.DB, opts ...gen.DOOption) employeeJobPositionAtCompanyORM {
	_employeeJobPositionAtCompanyORM := employeeJobPositionAtCompanyORM{}

	_employeeJobPositionAtCompanyORM.employeeJobPositionAtCompanyORMDo.UseDB(db, opts...)
	_employeeJobPositionAtCompanyORM.employeeJobPositionAtCompanyORMDo.UseModel(&accounting_servicev1.EmployeeJobPositionAtCompanyORM{})

	tableName := _employeeJobPositionAtCompanyORM.employeeJobPositionAtCompanyORMDo.TableName()
	_employeeJobPositionAtCompanyORM.ALL = field.NewAsterisk(tableName)
	_employeeJobPositionAtCompanyORM.CreatedAt = field.NewTime(tableName, "created_at")
	_employeeJobPositionAtCompanyORM.EffectiveDate = field.NewTime(tableName, "effective_date")
	_employeeJobPositionAtCompanyORM.EmployeeId = field.NewUint64(tableName, "employee_id")
	_employeeJobPositionAtCompanyORM.FlsaStatus = field.NewString(tableName, "flsa_status")
	_employeeJobPositionAtCompanyORM.Id = field.NewUint64(tableName, "id")
	_employeeJobPositionAtCompanyORM.JobTitle = field.NewString(tableName, "job_title")
	_employeeJobPositionAtCompanyORM.MergeAccountId = field.NewString(tableName, "merge_account_id")
	_employeeJobPositionAtCompanyORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_employeeJobPositionAtCompanyORM.PayCurrency = field.NewString(tableName, "pay_currency")
	_employeeJobPositionAtCompanyORM.PayFrequency = field.NewString(tableName, "pay_frequency")
	_employeeJobPositionAtCompanyORM.PayPeriod = field.NewString(tableName, "pay_period")
	_employeeJobPositionAtCompanyORM.PayRate = field.NewFloat64(tableName, "pay_rate")
	_employeeJobPositionAtCompanyORM.RemoteId = field.NewString(tableName, "remote_id")
	_employeeJobPositionAtCompanyORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")

	_employeeJobPositionAtCompanyORM.fillFieldMap()

	return _employeeJobPositionAtCompanyORM
}

type employeeJobPositionAtCompanyORM struct {
	employeeJobPositionAtCompanyORMDo

	ALL              field.Asterisk
	CreatedAt        field.Time
	EffectiveDate    field.Time
	EmployeeId       field.Uint64
	FlsaStatus       field.String
	Id               field.Uint64
	JobTitle         field.String
	MergeAccountId   field.String
	ModifiedAt       field.Time
	PayCurrency      field.String
	PayFrequency     field.String
	PayPeriod        field.String
	PayRate          field.Float64
	RemoteId         field.String
	RemoteWasDeleted field.Bool

	fieldMap map[string]field.Expr
}

func (e employeeJobPositionAtCompanyORM) Table(newTableName string) *employeeJobPositionAtCompanyORM {
	e.employeeJobPositionAtCompanyORMDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e employeeJobPositionAtCompanyORM) As(alias string) *employeeJobPositionAtCompanyORM {
	e.employeeJobPositionAtCompanyORMDo.DO = *(e.employeeJobPositionAtCompanyORMDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *employeeJobPositionAtCompanyORM) updateTableName(table string) *employeeJobPositionAtCompanyORM {
	e.ALL = field.NewAsterisk(table)
	e.CreatedAt = field.NewTime(table, "created_at")
	e.EffectiveDate = field.NewTime(table, "effective_date")
	e.EmployeeId = field.NewUint64(table, "employee_id")
	e.FlsaStatus = field.NewString(table, "flsa_status")
	e.Id = field.NewUint64(table, "id")
	e.JobTitle = field.NewString(table, "job_title")
	e.MergeAccountId = field.NewString(table, "merge_account_id")
	e.ModifiedAt = field.NewTime(table, "modified_at")
	e.PayCurrency = field.NewString(table, "pay_currency")
	e.PayFrequency = field.NewString(table, "pay_frequency")
	e.PayPeriod = field.NewString(table, "pay_period")
	e.PayRate = field.NewFloat64(table, "pay_rate")
	e.RemoteId = field.NewString(table, "remote_id")
	e.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")

	e.fillFieldMap()

	return e
}

func (e *employeeJobPositionAtCompanyORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *employeeJobPositionAtCompanyORM) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 14)
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["effective_date"] = e.EffectiveDate
	e.fieldMap["employee_id"] = e.EmployeeId
	e.fieldMap["flsa_status"] = e.FlsaStatus
	e.fieldMap["id"] = e.Id
	e.fieldMap["job_title"] = e.JobTitle
	e.fieldMap["merge_account_id"] = e.MergeAccountId
	e.fieldMap["modified_at"] = e.ModifiedAt
	e.fieldMap["pay_currency"] = e.PayCurrency
	e.fieldMap["pay_frequency"] = e.PayFrequency
	e.fieldMap["pay_period"] = e.PayPeriod
	e.fieldMap["pay_rate"] = e.PayRate
	e.fieldMap["remote_id"] = e.RemoteId
	e.fieldMap["remote_was_deleted"] = e.RemoteWasDeleted
}

func (e employeeJobPositionAtCompanyORM) clone(db *gorm.DB) employeeJobPositionAtCompanyORM {
	e.employeeJobPositionAtCompanyORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e employeeJobPositionAtCompanyORM) replaceDB(db *gorm.DB) employeeJobPositionAtCompanyORM {
	e.employeeJobPositionAtCompanyORMDo.ReplaceDB(db)
	return e
}

type employeeJobPositionAtCompanyORMDo struct{ gen.DO }

type IEmployeeJobPositionAtCompanyORMDo interface {
	gen.SubQuery
	Debug() IEmployeeJobPositionAtCompanyORMDo
	WithContext(ctx context.Context) IEmployeeJobPositionAtCompanyORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEmployeeJobPositionAtCompanyORMDo
	WriteDB() IEmployeeJobPositionAtCompanyORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEmployeeJobPositionAtCompanyORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEmployeeJobPositionAtCompanyORMDo
	Not(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo
	Or(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo
	Select(conds ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	Where(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo
	Order(conds ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	Distinct(cols ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	Omit(cols ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	Join(table schema.Tabler, on ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	Group(cols ...field.Expr) IEmployeeJobPositionAtCompanyORMDo
	Having(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo
	Limit(limit int) IEmployeeJobPositionAtCompanyORMDo
	Offset(offset int) IEmployeeJobPositionAtCompanyORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeJobPositionAtCompanyORMDo
	Unscoped() IEmployeeJobPositionAtCompanyORMDo
	Create(values ...*accounting_servicev1.EmployeeJobPositionAtCompanyORM) error
	CreateInBatches(values []*accounting_servicev1.EmployeeJobPositionAtCompanyORM, batchSize int) error
	Save(values ...*accounting_servicev1.EmployeeJobPositionAtCompanyORM) error
	First() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error)
	Take() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error)
	Last() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error)
	Find() ([]*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error)
	FindInBatches(result *[]*accounting_servicev1.EmployeeJobPositionAtCompanyORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.EmployeeJobPositionAtCompanyORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEmployeeJobPositionAtCompanyORMDo
	Assign(attrs ...field.AssignExpr) IEmployeeJobPositionAtCompanyORMDo
	Joins(fields ...field.RelationField) IEmployeeJobPositionAtCompanyORMDo
	Preload(fields ...field.RelationField) IEmployeeJobPositionAtCompanyORMDo
	FirstOrInit() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error)
	FirstOrCreate() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.EmployeeJobPositionAtCompanyORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEmployeeJobPositionAtCompanyORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error)
	CreateRecord(item accounting_servicev1.EmployeeJobPositionAtCompanyORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.EmployeeJobPositionAtCompanyORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (e employeeJobPositionAtCompanyORMDo) GetRecordByID(id int) (result accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employee_job_position_at_companies ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (e employeeJobPositionAtCompanyORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employee_job_position_at_companies ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (e employeeJobPositionAtCompanyORMDo) CreateRecord(item accounting_servicev1.EmployeeJobPositionAtCompanyORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO employee_job_position_at_companies (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (e employeeJobPositionAtCompanyORMDo) UpdateRecordByID(id int, item accounting_servicev1.EmployeeJobPositionAtCompanyORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE employee_job_position_at_companies SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (e employeeJobPositionAtCompanyORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM employee_job_position_at_companies ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (e employeeJobPositionAtCompanyORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employee_job_position_at_companies ORDER BY " + e.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (e employeeJobPositionAtCompanyORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM employee_job_position_at_companies ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (e employeeJobPositionAtCompanyORMDo) GetByID(id uint64) (result accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employee_job_position_at_companies ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (e employeeJobPositionAtCompanyORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employee_job_position_at_companies ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (e employeeJobPositionAtCompanyORMDo) Debug() IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Debug())
}

func (e employeeJobPositionAtCompanyORMDo) WithContext(ctx context.Context) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e employeeJobPositionAtCompanyORMDo) ReadDB() IEmployeeJobPositionAtCompanyORMDo {
	return e.Clauses(dbresolver.Read)
}

func (e employeeJobPositionAtCompanyORMDo) WriteDB() IEmployeeJobPositionAtCompanyORMDo {
	return e.Clauses(dbresolver.Write)
}

func (e employeeJobPositionAtCompanyORMDo) Session(config *gorm.Session) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Session(config))
}

func (e employeeJobPositionAtCompanyORMDo) Clauses(conds ...clause.Expression) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Returning(value interface{}, columns ...string) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e employeeJobPositionAtCompanyORMDo) Not(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Or(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Select(conds ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Where(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Order(conds ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Distinct(cols ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e employeeJobPositionAtCompanyORMDo) Omit(cols ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e employeeJobPositionAtCompanyORMDo) Join(table schema.Tabler, on ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e employeeJobPositionAtCompanyORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e employeeJobPositionAtCompanyORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e employeeJobPositionAtCompanyORMDo) Group(cols ...field.Expr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e employeeJobPositionAtCompanyORMDo) Having(conds ...gen.Condition) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e employeeJobPositionAtCompanyORMDo) Limit(limit int) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e employeeJobPositionAtCompanyORMDo) Offset(offset int) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e employeeJobPositionAtCompanyORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e employeeJobPositionAtCompanyORMDo) Unscoped() IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Unscoped())
}

func (e employeeJobPositionAtCompanyORMDo) Create(values ...*accounting_servicev1.EmployeeJobPositionAtCompanyORM) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e employeeJobPositionAtCompanyORMDo) CreateInBatches(values []*accounting_servicev1.EmployeeJobPositionAtCompanyORM, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e employeeJobPositionAtCompanyORMDo) Save(values ...*accounting_servicev1.EmployeeJobPositionAtCompanyORM) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e employeeJobPositionAtCompanyORMDo) First() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployeeJobPositionAtCompanyORM), nil
	}
}

func (e employeeJobPositionAtCompanyORMDo) Take() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployeeJobPositionAtCompanyORM), nil
	}
}

func (e employeeJobPositionAtCompanyORMDo) Last() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployeeJobPositionAtCompanyORM), nil
	}
}

func (e employeeJobPositionAtCompanyORMDo) Find() ([]*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error) {
	result, err := e.DO.Find()
	return result.([]*accounting_servicev1.EmployeeJobPositionAtCompanyORM), err
}

func (e employeeJobPositionAtCompanyORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.EmployeeJobPositionAtCompanyORM, err error) {
	buf := make([]*accounting_servicev1.EmployeeJobPositionAtCompanyORM, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e employeeJobPositionAtCompanyORMDo) FindInBatches(result *[]*accounting_servicev1.EmployeeJobPositionAtCompanyORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e employeeJobPositionAtCompanyORMDo) Attrs(attrs ...field.AssignExpr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e employeeJobPositionAtCompanyORMDo) Assign(attrs ...field.AssignExpr) IEmployeeJobPositionAtCompanyORMDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e employeeJobPositionAtCompanyORMDo) Joins(fields ...field.RelationField) IEmployeeJobPositionAtCompanyORMDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e employeeJobPositionAtCompanyORMDo) Preload(fields ...field.RelationField) IEmployeeJobPositionAtCompanyORMDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e employeeJobPositionAtCompanyORMDo) FirstOrInit() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployeeJobPositionAtCompanyORM), nil
	}
}

func (e employeeJobPositionAtCompanyORMDo) FirstOrCreate() (*accounting_servicev1.EmployeeJobPositionAtCompanyORM, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployeeJobPositionAtCompanyORM), nil
	}
}

func (e employeeJobPositionAtCompanyORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.EmployeeJobPositionAtCompanyORM, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e employeeJobPositionAtCompanyORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e employeeJobPositionAtCompanyORMDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e employeeJobPositionAtCompanyORMDo) Delete(models ...*accounting_servicev1.EmployeeJobPositionAtCompanyORM) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *employeeJobPositionAtCompanyORMDo) withDO(do gen.Dao) *employeeJobPositionAtCompanyORMDo {
	e.DO = *do.(*gen.DO)
	return e
}
