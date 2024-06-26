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

func newEmployerPayrollRunORM(db *gorm.DB, opts ...gen.DOOption) employerPayrollRunORM {
	_employerPayrollRunORM := employerPayrollRunORM{}

	_employerPayrollRunORM.employerPayrollRunORMDo.UseDB(db, opts...)
	_employerPayrollRunORM.employerPayrollRunORMDo.UseModel(&accounting_servicev1.EmployerPayrollRunORM{})

	tableName := _employerPayrollRunORM.employerPayrollRunORMDo.TableName()
	_employerPayrollRunORM.ALL = field.NewAsterisk(tableName)
	_employerPayrollRunORM.CheckDate = field.NewTime(tableName, "check_date")
	_employerPayrollRunORM.CreatedAt = field.NewTime(tableName, "created_at")
	_employerPayrollRunORM.EndDate = field.NewTime(tableName, "end_date")
	_employerPayrollRunORM.HrisLinkedAccountId = field.NewUint64(tableName, "hris_linked_account_id")
	_employerPayrollRunORM.Id = field.NewUint64(tableName, "id")
	_employerPayrollRunORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_employerPayrollRunORM.RemoteId = field.NewString(tableName, "remote_id")
	_employerPayrollRunORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_employerPayrollRunORM.RunState = field.NewString(tableName, "run_state")
	_employerPayrollRunORM.RunType = field.NewString(tableName, "run_type")
	_employerPayrollRunORM.StartDate = field.NewTime(tableName, "start_date")
	_employerPayrollRunORM.PayrollRuns = employerPayrollRunORMHasManyPayrollRuns{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PayrollRuns", "accounting_servicev1.EmployeePayrollRunORM"),
		Deductions: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PayrollRuns.Deductions", "accounting_servicev1.DeductionORM"),
		},
		Earnings: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PayrollRuns.Earnings", "accounting_servicev1.EarningORM"),
		},
		Taxes: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PayrollRuns.Taxes", "accounting_servicev1.TaxORM"),
		},
	}

	_employerPayrollRunORM.fillFieldMap()

	return _employerPayrollRunORM
}

type employerPayrollRunORM struct {
	employerPayrollRunORMDo

	ALL                 field.Asterisk
	CheckDate           field.Time
	CreatedAt           field.Time
	EndDate             field.Time
	HrisLinkedAccountId field.Uint64
	Id                  field.Uint64
	ModifiedAt          field.Time
	RemoteId            field.String
	RemoteWasDeleted    field.Bool
	RunState            field.String
	RunType             field.String
	StartDate           field.Time
	PayrollRuns         employerPayrollRunORMHasManyPayrollRuns

	fieldMap map[string]field.Expr
}

func (e employerPayrollRunORM) Table(newTableName string) *employerPayrollRunORM {
	e.employerPayrollRunORMDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e employerPayrollRunORM) As(alias string) *employerPayrollRunORM {
	e.employerPayrollRunORMDo.DO = *(e.employerPayrollRunORMDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *employerPayrollRunORM) updateTableName(table string) *employerPayrollRunORM {
	e.ALL = field.NewAsterisk(table)
	e.CheckDate = field.NewTime(table, "check_date")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.EndDate = field.NewTime(table, "end_date")
	e.HrisLinkedAccountId = field.NewUint64(table, "hris_linked_account_id")
	e.Id = field.NewUint64(table, "id")
	e.ModifiedAt = field.NewTime(table, "modified_at")
	e.RemoteId = field.NewString(table, "remote_id")
	e.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	e.RunState = field.NewString(table, "run_state")
	e.RunType = field.NewString(table, "run_type")
	e.StartDate = field.NewTime(table, "start_date")

	e.fillFieldMap()

	return e
}

func (e *employerPayrollRunORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *employerPayrollRunORM) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 12)
	e.fieldMap["check_date"] = e.CheckDate
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["end_date"] = e.EndDate
	e.fieldMap["hris_linked_account_id"] = e.HrisLinkedAccountId
	e.fieldMap["id"] = e.Id
	e.fieldMap["modified_at"] = e.ModifiedAt
	e.fieldMap["remote_id"] = e.RemoteId
	e.fieldMap["remote_was_deleted"] = e.RemoteWasDeleted
	e.fieldMap["run_state"] = e.RunState
	e.fieldMap["run_type"] = e.RunType
	e.fieldMap["start_date"] = e.StartDate

}

func (e employerPayrollRunORM) clone(db *gorm.DB) employerPayrollRunORM {
	e.employerPayrollRunORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e employerPayrollRunORM) replaceDB(db *gorm.DB) employerPayrollRunORM {
	e.employerPayrollRunORMDo.ReplaceDB(db)
	return e
}

type employerPayrollRunORMHasManyPayrollRuns struct {
	db *gorm.DB

	field.RelationField

	Deductions struct {
		field.RelationField
	}
	Earnings struct {
		field.RelationField
	}
	Taxes struct {
		field.RelationField
	}
}

func (a employerPayrollRunORMHasManyPayrollRuns) Where(conds ...field.Expr) *employerPayrollRunORMHasManyPayrollRuns {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a employerPayrollRunORMHasManyPayrollRuns) WithContext(ctx context.Context) *employerPayrollRunORMHasManyPayrollRuns {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a employerPayrollRunORMHasManyPayrollRuns) Session(session *gorm.Session) *employerPayrollRunORMHasManyPayrollRuns {
	a.db = a.db.Session(session)
	return &a
}

func (a employerPayrollRunORMHasManyPayrollRuns) Model(m *accounting_servicev1.EmployerPayrollRunORM) *employerPayrollRunORMHasManyPayrollRunsTx {
	return &employerPayrollRunORMHasManyPayrollRunsTx{a.db.Model(m).Association(a.Name())}
}

type employerPayrollRunORMHasManyPayrollRunsTx struct{ tx *gorm.Association }

func (a employerPayrollRunORMHasManyPayrollRunsTx) Find() (result []*accounting_servicev1.EmployeePayrollRunORM, err error) {
	return result, a.tx.Find(&result)
}

func (a employerPayrollRunORMHasManyPayrollRunsTx) Append(values ...*accounting_servicev1.EmployeePayrollRunORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a employerPayrollRunORMHasManyPayrollRunsTx) Replace(values ...*accounting_servicev1.EmployeePayrollRunORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a employerPayrollRunORMHasManyPayrollRunsTx) Delete(values ...*accounting_servicev1.EmployeePayrollRunORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a employerPayrollRunORMHasManyPayrollRunsTx) Clear() error {
	return a.tx.Clear()
}

func (a employerPayrollRunORMHasManyPayrollRunsTx) Count() int64 {
	return a.tx.Count()
}

type employerPayrollRunORMDo struct{ gen.DO }

type IEmployerPayrollRunORMDo interface {
	gen.SubQuery
	Debug() IEmployerPayrollRunORMDo
	WithContext(ctx context.Context) IEmployerPayrollRunORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEmployerPayrollRunORMDo
	WriteDB() IEmployerPayrollRunORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEmployerPayrollRunORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEmployerPayrollRunORMDo
	Not(conds ...gen.Condition) IEmployerPayrollRunORMDo
	Or(conds ...gen.Condition) IEmployerPayrollRunORMDo
	Select(conds ...field.Expr) IEmployerPayrollRunORMDo
	Where(conds ...gen.Condition) IEmployerPayrollRunORMDo
	Order(conds ...field.Expr) IEmployerPayrollRunORMDo
	Distinct(cols ...field.Expr) IEmployerPayrollRunORMDo
	Omit(cols ...field.Expr) IEmployerPayrollRunORMDo
	Join(table schema.Tabler, on ...field.Expr) IEmployerPayrollRunORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEmployerPayrollRunORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEmployerPayrollRunORMDo
	Group(cols ...field.Expr) IEmployerPayrollRunORMDo
	Having(conds ...gen.Condition) IEmployerPayrollRunORMDo
	Limit(limit int) IEmployerPayrollRunORMDo
	Offset(offset int) IEmployerPayrollRunORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployerPayrollRunORMDo
	Unscoped() IEmployerPayrollRunORMDo
	Create(values ...*accounting_servicev1.EmployerPayrollRunORM) error
	CreateInBatches(values []*accounting_servicev1.EmployerPayrollRunORM, batchSize int) error
	Save(values ...*accounting_servicev1.EmployerPayrollRunORM) error
	First() (*accounting_servicev1.EmployerPayrollRunORM, error)
	Take() (*accounting_servicev1.EmployerPayrollRunORM, error)
	Last() (*accounting_servicev1.EmployerPayrollRunORM, error)
	Find() ([]*accounting_servicev1.EmployerPayrollRunORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.EmployerPayrollRunORM, err error)
	FindInBatches(result *[]*accounting_servicev1.EmployerPayrollRunORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.EmployerPayrollRunORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEmployerPayrollRunORMDo
	Assign(attrs ...field.AssignExpr) IEmployerPayrollRunORMDo
	Joins(fields ...field.RelationField) IEmployerPayrollRunORMDo
	Preload(fields ...field.RelationField) IEmployerPayrollRunORMDo
	FirstOrInit() (*accounting_servicev1.EmployerPayrollRunORM, error)
	FirstOrCreate() (*accounting_servicev1.EmployerPayrollRunORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.EmployerPayrollRunORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEmployerPayrollRunORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.EmployerPayrollRunORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.EmployerPayrollRunORM, err error)
	CreateRecord(item accounting_servicev1.EmployerPayrollRunORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.EmployerPayrollRunORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.EmployerPayrollRunORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.EmployerPayrollRunORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.EmployerPayrollRunORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (e employerPayrollRunORMDo) GetRecordByID(id int) (result accounting_servicev1.EmployerPayrollRunORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employer_payroll_runs ")
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
func (e employerPayrollRunORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.EmployerPayrollRunORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employer_payroll_runs ")
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
func (e employerPayrollRunORMDo) CreateRecord(item accounting_servicev1.EmployerPayrollRunORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO employer_payroll_runs (columns) VALUES (values) ")

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
func (e employerPayrollRunORMDo) UpdateRecordByID(id int, item accounting_servicev1.EmployerPayrollRunORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE employer_payroll_runs SET columns=values ")
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
func (e employerPayrollRunORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM employer_payroll_runs ")
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
func (e employerPayrollRunORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.EmployerPayrollRunORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employer_payroll_runs ORDER BY " + e.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (e employerPayrollRunORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM employer_payroll_runs ")

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
func (e employerPayrollRunORMDo) GetByID(id uint64) (result accounting_servicev1.EmployerPayrollRunORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employer_payroll_runs ")
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
func (e employerPayrollRunORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.EmployerPayrollRunORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM employer_payroll_runs ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = e.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (e employerPayrollRunORMDo) Debug() IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Debug())
}

func (e employerPayrollRunORMDo) WithContext(ctx context.Context) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e employerPayrollRunORMDo) ReadDB() IEmployerPayrollRunORMDo {
	return e.Clauses(dbresolver.Read)
}

func (e employerPayrollRunORMDo) WriteDB() IEmployerPayrollRunORMDo {
	return e.Clauses(dbresolver.Write)
}

func (e employerPayrollRunORMDo) Session(config *gorm.Session) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Session(config))
}

func (e employerPayrollRunORMDo) Clauses(conds ...clause.Expression) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e employerPayrollRunORMDo) Returning(value interface{}, columns ...string) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e employerPayrollRunORMDo) Not(conds ...gen.Condition) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e employerPayrollRunORMDo) Or(conds ...gen.Condition) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e employerPayrollRunORMDo) Select(conds ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e employerPayrollRunORMDo) Where(conds ...gen.Condition) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e employerPayrollRunORMDo) Order(conds ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e employerPayrollRunORMDo) Distinct(cols ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e employerPayrollRunORMDo) Omit(cols ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e employerPayrollRunORMDo) Join(table schema.Tabler, on ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e employerPayrollRunORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e employerPayrollRunORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e employerPayrollRunORMDo) Group(cols ...field.Expr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e employerPayrollRunORMDo) Having(conds ...gen.Condition) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e employerPayrollRunORMDo) Limit(limit int) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e employerPayrollRunORMDo) Offset(offset int) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e employerPayrollRunORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e employerPayrollRunORMDo) Unscoped() IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Unscoped())
}

func (e employerPayrollRunORMDo) Create(values ...*accounting_servicev1.EmployerPayrollRunORM) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e employerPayrollRunORMDo) CreateInBatches(values []*accounting_servicev1.EmployerPayrollRunORM, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e employerPayrollRunORMDo) Save(values ...*accounting_servicev1.EmployerPayrollRunORM) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e employerPayrollRunORMDo) First() (*accounting_servicev1.EmployerPayrollRunORM, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployerPayrollRunORM), nil
	}
}

func (e employerPayrollRunORMDo) Take() (*accounting_servicev1.EmployerPayrollRunORM, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployerPayrollRunORM), nil
	}
}

func (e employerPayrollRunORMDo) Last() (*accounting_servicev1.EmployerPayrollRunORM, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployerPayrollRunORM), nil
	}
}

func (e employerPayrollRunORMDo) Find() ([]*accounting_servicev1.EmployerPayrollRunORM, error) {
	result, err := e.DO.Find()
	return result.([]*accounting_servicev1.EmployerPayrollRunORM), err
}

func (e employerPayrollRunORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.EmployerPayrollRunORM, err error) {
	buf := make([]*accounting_servicev1.EmployerPayrollRunORM, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e employerPayrollRunORMDo) FindInBatches(result *[]*accounting_servicev1.EmployerPayrollRunORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e employerPayrollRunORMDo) Attrs(attrs ...field.AssignExpr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e employerPayrollRunORMDo) Assign(attrs ...field.AssignExpr) IEmployerPayrollRunORMDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e employerPayrollRunORMDo) Joins(fields ...field.RelationField) IEmployerPayrollRunORMDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e employerPayrollRunORMDo) Preload(fields ...field.RelationField) IEmployerPayrollRunORMDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e employerPayrollRunORMDo) FirstOrInit() (*accounting_servicev1.EmployerPayrollRunORM, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployerPayrollRunORM), nil
	}
}

func (e employerPayrollRunORMDo) FirstOrCreate() (*accounting_servicev1.EmployerPayrollRunORM, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.EmployerPayrollRunORM), nil
	}
}

func (e employerPayrollRunORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.EmployerPayrollRunORM, count int64, err error) {
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

func (e employerPayrollRunORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e employerPayrollRunORMDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e employerPayrollRunORMDo) Delete(models ...*accounting_servicev1.EmployerPayrollRunORM) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *employerPayrollRunORMDo) withDO(do gen.Dao) *employerPayrollRunORMDo {
	e.DO = *do.(*gen.DO)
	return e
}
