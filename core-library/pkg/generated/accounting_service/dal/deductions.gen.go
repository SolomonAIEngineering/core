// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	accounting_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/accounting_service/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newDeductionORM(db *gorm.DB, opts ...gen.DOOption) deductionORM {
	_deductionORM := deductionORM{}

	_deductionORM.deductionORMDo.UseDB(db, opts...)
	_deductionORM.deductionORMDo.UseModel(&accounting_servicev1.DeductionORM{})

	tableName := _deductionORM.deductionORMDo.TableName()
	_deductionORM.ALL = field.NewAsterisk(tableName)
	_deductionORM.CompanyDeduction = field.NewFloat64(tableName, "company_deduction")
	_deductionORM.CreatedAt = field.NewTime(tableName, "created_at")
	_deductionORM.EmployeeDeduction = field.NewFloat64(tableName, "employee_deduction")
	_deductionORM.EmployeePayrollRunId = field.NewUint64(tableName, "employee_payroll_run_id")
	_deductionORM.Id = field.NewUint64(tableName, "id")
	_deductionORM.MergeAccountId = field.NewString(tableName, "merge_account_id")
	_deductionORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_deductionORM.Name = field.NewString(tableName, "name")
	_deductionORM.PayrollRunMergeAccountId = field.NewString(tableName, "payroll_run_merge_account_id")
	_deductionORM.RemoteId = field.NewString(tableName, "remote_id")
	_deductionORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")

	_deductionORM.fillFieldMap()

	return _deductionORM
}

type deductionORM struct {
	deductionORMDo

	ALL                      field.Asterisk
	CompanyDeduction         field.Float64
	CreatedAt                field.Time
	EmployeeDeduction        field.Float64
	EmployeePayrollRunId     field.Uint64
	Id                       field.Uint64
	MergeAccountId           field.String
	ModifiedAt               field.Time
	Name                     field.String
	PayrollRunMergeAccountId field.String
	RemoteId                 field.String
	RemoteWasDeleted         field.Bool

	fieldMap map[string]field.Expr
}

func (d deductionORM) Table(newTableName string) *deductionORM {
	d.deductionORMDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deductionORM) As(alias string) *deductionORM {
	d.deductionORMDo.DO = *(d.deductionORMDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deductionORM) updateTableName(table string) *deductionORM {
	d.ALL = field.NewAsterisk(table)
	d.CompanyDeduction = field.NewFloat64(table, "company_deduction")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.EmployeeDeduction = field.NewFloat64(table, "employee_deduction")
	d.EmployeePayrollRunId = field.NewUint64(table, "employee_payroll_run_id")
	d.Id = field.NewUint64(table, "id")
	d.MergeAccountId = field.NewString(table, "merge_account_id")
	d.ModifiedAt = field.NewTime(table, "modified_at")
	d.Name = field.NewString(table, "name")
	d.PayrollRunMergeAccountId = field.NewString(table, "payroll_run_merge_account_id")
	d.RemoteId = field.NewString(table, "remote_id")
	d.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")

	d.fillFieldMap()

	return d
}

func (d *deductionORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deductionORM) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["company_deduction"] = d.CompanyDeduction
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["employee_deduction"] = d.EmployeeDeduction
	d.fieldMap["employee_payroll_run_id"] = d.EmployeePayrollRunId
	d.fieldMap["id"] = d.Id
	d.fieldMap["merge_account_id"] = d.MergeAccountId
	d.fieldMap["modified_at"] = d.ModifiedAt
	d.fieldMap["name"] = d.Name
	d.fieldMap["payroll_run_merge_account_id"] = d.PayrollRunMergeAccountId
	d.fieldMap["remote_id"] = d.RemoteId
	d.fieldMap["remote_was_deleted"] = d.RemoteWasDeleted
}

func (d deductionORM) clone(db *gorm.DB) deductionORM {
	d.deductionORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deductionORM) replaceDB(db *gorm.DB) deductionORM {
	d.deductionORMDo.ReplaceDB(db)
	return d
}

type deductionORMDo struct{ gen.DO }

type IDeductionORMDo interface {
	gen.SubQuery
	Debug() IDeductionORMDo
	WithContext(ctx context.Context) IDeductionORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeductionORMDo
	WriteDB() IDeductionORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeductionORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeductionORMDo
	Not(conds ...gen.Condition) IDeductionORMDo
	Or(conds ...gen.Condition) IDeductionORMDo
	Select(conds ...field.Expr) IDeductionORMDo
	Where(conds ...gen.Condition) IDeductionORMDo
	Order(conds ...field.Expr) IDeductionORMDo
	Distinct(cols ...field.Expr) IDeductionORMDo
	Omit(cols ...field.Expr) IDeductionORMDo
	Join(table schema.Tabler, on ...field.Expr) IDeductionORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeductionORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeductionORMDo
	Group(cols ...field.Expr) IDeductionORMDo
	Having(conds ...gen.Condition) IDeductionORMDo
	Limit(limit int) IDeductionORMDo
	Offset(offset int) IDeductionORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeductionORMDo
	Unscoped() IDeductionORMDo
	Create(values ...*accounting_servicev1.DeductionORM) error
	CreateInBatches(values []*accounting_servicev1.DeductionORM, batchSize int) error
	Save(values ...*accounting_servicev1.DeductionORM) error
	First() (*accounting_servicev1.DeductionORM, error)
	Take() (*accounting_servicev1.DeductionORM, error)
	Last() (*accounting_servicev1.DeductionORM, error)
	Find() ([]*accounting_servicev1.DeductionORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.DeductionORM, err error)
	FindInBatches(result *[]*accounting_servicev1.DeductionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.DeductionORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeductionORMDo
	Assign(attrs ...field.AssignExpr) IDeductionORMDo
	Joins(fields ...field.RelationField) IDeductionORMDo
	Preload(fields ...field.RelationField) IDeductionORMDo
	FirstOrInit() (*accounting_servicev1.DeductionORM, error)
	FirstOrCreate() (*accounting_servicev1.DeductionORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.DeductionORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeductionORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.DeductionORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.DeductionORM, err error)
	CreateRecord(item accounting_servicev1.DeductionORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.DeductionORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.DeductionORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.DeductionORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.DeductionORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (d deductionORMDo) GetRecordByID(id int) (result accounting_servicev1.DeductionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM deductions ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (d deductionORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.DeductionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM deductions ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (d deductionORMDo) CreateRecord(item accounting_servicev1.DeductionORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO deductions (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (d deductionORMDo) UpdateRecordByID(id int, item accounting_servicev1.DeductionORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE deductions SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (d deductionORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM deductions ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (d deductionORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.DeductionORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM deductions ORDER BY " + d.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (d deductionORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM deductions ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (d deductionORMDo) GetByID(id uint64) (result accounting_servicev1.DeductionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM deductions ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (d deductionORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.DeductionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM deductions ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (d deductionORMDo) Debug() IDeductionORMDo {
	return d.withDO(d.DO.Debug())
}

func (d deductionORMDo) WithContext(ctx context.Context) IDeductionORMDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deductionORMDo) ReadDB() IDeductionORMDo {
	return d.Clauses(dbresolver.Read)
}

func (d deductionORMDo) WriteDB() IDeductionORMDo {
	return d.Clauses(dbresolver.Write)
}

func (d deductionORMDo) Session(config *gorm.Session) IDeductionORMDo {
	return d.withDO(d.DO.Session(config))
}

func (d deductionORMDo) Clauses(conds ...clause.Expression) IDeductionORMDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deductionORMDo) Returning(value interface{}, columns ...string) IDeductionORMDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deductionORMDo) Not(conds ...gen.Condition) IDeductionORMDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deductionORMDo) Or(conds ...gen.Condition) IDeductionORMDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deductionORMDo) Select(conds ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deductionORMDo) Where(conds ...gen.Condition) IDeductionORMDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deductionORMDo) Order(conds ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deductionORMDo) Distinct(cols ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deductionORMDo) Omit(cols ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deductionORMDo) Join(table schema.Tabler, on ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deductionORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deductionORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deductionORMDo) Group(cols ...field.Expr) IDeductionORMDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deductionORMDo) Having(conds ...gen.Condition) IDeductionORMDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deductionORMDo) Limit(limit int) IDeductionORMDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deductionORMDo) Offset(offset int) IDeductionORMDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deductionORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeductionORMDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deductionORMDo) Unscoped() IDeductionORMDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deductionORMDo) Create(values ...*accounting_servicev1.DeductionORM) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deductionORMDo) CreateInBatches(values []*accounting_servicev1.DeductionORM, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deductionORMDo) Save(values ...*accounting_servicev1.DeductionORM) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deductionORMDo) First() (*accounting_servicev1.DeductionORM, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DeductionORM), nil
	}
}

func (d deductionORMDo) Take() (*accounting_servicev1.DeductionORM, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DeductionORM), nil
	}
}

func (d deductionORMDo) Last() (*accounting_servicev1.DeductionORM, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DeductionORM), nil
	}
}

func (d deductionORMDo) Find() ([]*accounting_servicev1.DeductionORM, error) {
	result, err := d.DO.Find()
	return result.([]*accounting_servicev1.DeductionORM), err
}

func (d deductionORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.DeductionORM, err error) {
	buf := make([]*accounting_servicev1.DeductionORM, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deductionORMDo) FindInBatches(result *[]*accounting_servicev1.DeductionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deductionORMDo) Attrs(attrs ...field.AssignExpr) IDeductionORMDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deductionORMDo) Assign(attrs ...field.AssignExpr) IDeductionORMDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deductionORMDo) Joins(fields ...field.RelationField) IDeductionORMDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deductionORMDo) Preload(fields ...field.RelationField) IDeductionORMDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deductionORMDo) FirstOrInit() (*accounting_servicev1.DeductionORM, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DeductionORM), nil
	}
}

func (d deductionORMDo) FirstOrCreate() (*accounting_servicev1.DeductionORM, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DeductionORM), nil
	}
}

func (d deductionORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.DeductionORM, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deductionORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deductionORMDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deductionORMDo) Delete(models ...*accounting_servicev1.DeductionORM) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deductionORMDo) withDO(do gen.Dao) *deductionORMDo {
	d.DO = *do.(*gen.DO)
	return d
}