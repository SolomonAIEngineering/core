// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"

	accounting_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/accounting_service/v1"
)

func newTimeOffORM(db *gorm.DB, opts ...gen.DOOption) timeOffORM {
	_timeOffORM := timeOffORM{}

	_timeOffORM.timeOffORMDo.UseDB(db, opts...)
	_timeOffORM.timeOffORMDo.UseModel(&accounting_servicev1.TimeOffORM{})

	tableName := _timeOffORM.timeOffORMDo.TableName()
	_timeOffORM.ALL = field.NewAsterisk(tableName)
	_timeOffORM.Amount = field.NewFloat64(tableName, "amount")
	_timeOffORM.ApproverMergeAccountId = field.NewString(tableName, "approver_merge_account_id")
	_timeOffORM.CreatedAt = field.NewTime(tableName, "created_at")
	_timeOffORM.EmployeeMergeAccountId = field.NewString(tableName, "employee_merge_account_id")
	_timeOffORM.EmploymentNote = field.NewString(tableName, "employment_note")
	_timeOffORM.EndTime = field.NewTime(tableName, "end_time")
	_timeOffORM.Id = field.NewUint64(tableName, "id")
	_timeOffORM.MergeAccountId = field.NewString(tableName, "merge_account_id")
	_timeOffORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_timeOffORM.RemoteId = field.NewString(tableName, "remote_id")
	_timeOffORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_timeOffORM.RequestType = field.NewString(tableName, "request_type")
	_timeOffORM.StartTime = field.NewTime(tableName, "start_time")
	_timeOffORM.Status = field.NewString(tableName, "status")
	_timeOffORM.Units = field.NewString(tableName, "units")
	_timeOffORM.Approver = timeOffORMHasOneApprover{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Approver", "accounting_servicev1.EmployeeORM"),
		Manager: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.Manager", "accounting_servicev1.EmployeeORM"),
		},
		Group: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.Group", "accounting_servicev1.GroupORM"),
		},
		HomeLocation: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.HomeLocation", "accounting_servicev1.LocationAddressORM"),
		},
		PayTimeOffBalance: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.PayTimeOffBalance", "accounting_servicev1.EmployeTimeOffBalanceORM"),
		},
		WorkLocation: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.WorkLocation", "accounting_servicev1.LocationAddressORM"),
		},
		BankAccounts: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.BankAccounts", "accounting_servicev1.BankInfoORM"),
		},
		Benefits: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.Benefits", "accounting_servicev1.EmployeeBenefitsORM"),
		},
		Dependents: struct {
			field.RelationField
			HomeLocation struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Approver.Dependents", "accounting_servicev1.DependentsORM"),
			HomeLocation: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Approver.Dependents.HomeLocation", "accounting_servicev1.LocationAddressORM"),
			},
		},
		Employments: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Approver.Employments", "accounting_servicev1.EmployeeJobPositionAtCompanyORM"),
		},
		PayrollRuns: struct {
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
		}{
			RelationField: field.NewRelation("Approver.PayrollRuns", "accounting_servicev1.EmployeePayrollRunORM"),
			Deductions: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Approver.PayrollRuns.Deductions", "accounting_servicev1.DeductionORM"),
			},
			Earnings: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Approver.PayrollRuns.Earnings", "accounting_servicev1.EarningORM"),
			},
			Taxes: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Approver.PayrollRuns.Taxes", "accounting_servicev1.TaxORM"),
			},
		},
	}

	_timeOffORM.Employee = timeOffORMHasOneEmployee{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Employee", "accounting_servicev1.EmployeeORM"),
	}

	_timeOffORM.fillFieldMap()

	return _timeOffORM
}

type timeOffORM struct {
	timeOffORMDo

	ALL                    field.Asterisk
	Amount                 field.Float64
	ApproverMergeAccountId field.String
	CreatedAt              field.Time
	EmployeeMergeAccountId field.String
	EmploymentNote         field.String
	EndTime                field.Time
	Id                     field.Uint64
	MergeAccountId         field.String
	ModifiedAt             field.Time
	RemoteId               field.String
	RemoteWasDeleted       field.Bool
	RequestType            field.String
	StartTime              field.Time
	Status                 field.String
	Units                  field.String
	Approver               timeOffORMHasOneApprover

	Employee timeOffORMHasOneEmployee

	fieldMap map[string]field.Expr
}

func (t timeOffORM) Table(newTableName string) *timeOffORM {
	t.timeOffORMDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timeOffORM) As(alias string) *timeOffORM {
	t.timeOffORMDo.DO = *(t.timeOffORMDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timeOffORM) updateTableName(table string) *timeOffORM {
	t.ALL = field.NewAsterisk(table)
	t.Amount = field.NewFloat64(table, "amount")
	t.ApproverMergeAccountId = field.NewString(table, "approver_merge_account_id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.EmployeeMergeAccountId = field.NewString(table, "employee_merge_account_id")
	t.EmploymentNote = field.NewString(table, "employment_note")
	t.EndTime = field.NewTime(table, "end_time")
	t.Id = field.NewUint64(table, "id")
	t.MergeAccountId = field.NewString(table, "merge_account_id")
	t.ModifiedAt = field.NewTime(table, "modified_at")
	t.RemoteId = field.NewString(table, "remote_id")
	t.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	t.RequestType = field.NewString(table, "request_type")
	t.StartTime = field.NewTime(table, "start_time")
	t.Status = field.NewString(table, "status")
	t.Units = field.NewString(table, "units")

	t.fillFieldMap()

	return t
}

func (t *timeOffORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timeOffORM) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 17)
	t.fieldMap["amount"] = t.Amount
	t.fieldMap["approver_merge_account_id"] = t.ApproverMergeAccountId
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["employee_merge_account_id"] = t.EmployeeMergeAccountId
	t.fieldMap["employment_note"] = t.EmploymentNote
	t.fieldMap["end_time"] = t.EndTime
	t.fieldMap["id"] = t.Id
	t.fieldMap["merge_account_id"] = t.MergeAccountId
	t.fieldMap["modified_at"] = t.ModifiedAt
	t.fieldMap["remote_id"] = t.RemoteId
	t.fieldMap["remote_was_deleted"] = t.RemoteWasDeleted
	t.fieldMap["request_type"] = t.RequestType
	t.fieldMap["start_time"] = t.StartTime
	t.fieldMap["status"] = t.Status
	t.fieldMap["units"] = t.Units

}

func (t timeOffORM) clone(db *gorm.DB) timeOffORM {
	t.timeOffORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timeOffORM) replaceDB(db *gorm.DB) timeOffORM {
	t.timeOffORMDo.ReplaceDB(db)
	return t
}

type timeOffORMHasOneApprover struct {
	db *gorm.DB

	field.RelationField

	Manager struct {
		field.RelationField
	}
	Group struct {
		field.RelationField
	}
	HomeLocation struct {
		field.RelationField
	}
	PayTimeOffBalance struct {
		field.RelationField
	}
	WorkLocation struct {
		field.RelationField
	}
	BankAccounts struct {
		field.RelationField
	}
	Benefits struct {
		field.RelationField
	}
	Dependents struct {
		field.RelationField
		HomeLocation struct {
			field.RelationField
		}
	}
	Employments struct {
		field.RelationField
	}
	PayrollRuns struct {
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
}

func (a timeOffORMHasOneApprover) Where(conds ...field.Expr) *timeOffORMHasOneApprover {
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

func (a timeOffORMHasOneApprover) WithContext(ctx context.Context) *timeOffORMHasOneApprover {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a timeOffORMHasOneApprover) Session(session *gorm.Session) *timeOffORMHasOneApprover {
	a.db = a.db.Session(session)
	return &a
}

func (a timeOffORMHasOneApprover) Model(m *accounting_servicev1.TimeOffORM) *timeOffORMHasOneApproverTx {
	return &timeOffORMHasOneApproverTx{a.db.Model(m).Association(a.Name())}
}

type timeOffORMHasOneApproverTx struct{ tx *gorm.Association }

func (a timeOffORMHasOneApproverTx) Find() (result *accounting_servicev1.EmployeeORM, err error) {
	return result, a.tx.Find(&result)
}

func (a timeOffORMHasOneApproverTx) Append(values ...*accounting_servicev1.EmployeeORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a timeOffORMHasOneApproverTx) Replace(values ...*accounting_servicev1.EmployeeORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a timeOffORMHasOneApproverTx) Delete(values ...*accounting_servicev1.EmployeeORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a timeOffORMHasOneApproverTx) Clear() error {
	return a.tx.Clear()
}

func (a timeOffORMHasOneApproverTx) Count() int64 {
	return a.tx.Count()
}

type timeOffORMHasOneEmployee struct {
	db *gorm.DB

	field.RelationField
}

func (a timeOffORMHasOneEmployee) Where(conds ...field.Expr) *timeOffORMHasOneEmployee {
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

func (a timeOffORMHasOneEmployee) WithContext(ctx context.Context) *timeOffORMHasOneEmployee {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a timeOffORMHasOneEmployee) Session(session *gorm.Session) *timeOffORMHasOneEmployee {
	a.db = a.db.Session(session)
	return &a
}

func (a timeOffORMHasOneEmployee) Model(m *accounting_servicev1.TimeOffORM) *timeOffORMHasOneEmployeeTx {
	return &timeOffORMHasOneEmployeeTx{a.db.Model(m).Association(a.Name())}
}

type timeOffORMHasOneEmployeeTx struct{ tx *gorm.Association }

func (a timeOffORMHasOneEmployeeTx) Find() (result *accounting_servicev1.EmployeeORM, err error) {
	return result, a.tx.Find(&result)
}

func (a timeOffORMHasOneEmployeeTx) Append(values ...*accounting_servicev1.EmployeeORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a timeOffORMHasOneEmployeeTx) Replace(values ...*accounting_servicev1.EmployeeORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a timeOffORMHasOneEmployeeTx) Delete(values ...*accounting_servicev1.EmployeeORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a timeOffORMHasOneEmployeeTx) Clear() error {
	return a.tx.Clear()
}

func (a timeOffORMHasOneEmployeeTx) Count() int64 {
	return a.tx.Count()
}

type timeOffORMDo struct{ gen.DO }

type ITimeOffORMDo interface {
	gen.SubQuery
	Debug() ITimeOffORMDo
	WithContext(ctx context.Context) ITimeOffORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITimeOffORMDo
	WriteDB() ITimeOffORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITimeOffORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITimeOffORMDo
	Not(conds ...gen.Condition) ITimeOffORMDo
	Or(conds ...gen.Condition) ITimeOffORMDo
	Select(conds ...field.Expr) ITimeOffORMDo
	Where(conds ...gen.Condition) ITimeOffORMDo
	Order(conds ...field.Expr) ITimeOffORMDo
	Distinct(cols ...field.Expr) ITimeOffORMDo
	Omit(cols ...field.Expr) ITimeOffORMDo
	Join(table schema.Tabler, on ...field.Expr) ITimeOffORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITimeOffORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITimeOffORMDo
	Group(cols ...field.Expr) ITimeOffORMDo
	Having(conds ...gen.Condition) ITimeOffORMDo
	Limit(limit int) ITimeOffORMDo
	Offset(offset int) ITimeOffORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITimeOffORMDo
	Unscoped() ITimeOffORMDo
	Create(values ...*accounting_servicev1.TimeOffORM) error
	CreateInBatches(values []*accounting_servicev1.TimeOffORM, batchSize int) error
	Save(values ...*accounting_servicev1.TimeOffORM) error
	First() (*accounting_servicev1.TimeOffORM, error)
	Take() (*accounting_servicev1.TimeOffORM, error)
	Last() (*accounting_servicev1.TimeOffORM, error)
	Find() ([]*accounting_servicev1.TimeOffORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.TimeOffORM, err error)
	FindInBatches(result *[]*accounting_servicev1.TimeOffORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.TimeOffORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITimeOffORMDo
	Assign(attrs ...field.AssignExpr) ITimeOffORMDo
	Joins(fields ...field.RelationField) ITimeOffORMDo
	Preload(fields ...field.RelationField) ITimeOffORMDo
	FirstOrInit() (*accounting_servicev1.TimeOffORM, error)
	FirstOrCreate() (*accounting_servicev1.TimeOffORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.TimeOffORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITimeOffORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.TimeOffORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.TimeOffORM, err error)
	CreateRecord(item accounting_servicev1.TimeOffORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.TimeOffORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.TimeOffORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.TimeOffORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.TimeOffORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (t timeOffORMDo) GetRecordByID(id int) (result accounting_servicev1.TimeOffORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM time_offs ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (t timeOffORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.TimeOffORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM time_offs ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (t timeOffORMDo) CreateRecord(item accounting_servicev1.TimeOffORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO time_offs (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (t timeOffORMDo) UpdateRecordByID(id int, item accounting_servicev1.TimeOffORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE time_offs SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (t timeOffORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM time_offs ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (t timeOffORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.TimeOffORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM time_offs ORDER BY " + t.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (t timeOffORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM time_offs ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (t timeOffORMDo) GetByID(id uint64) (result accounting_servicev1.TimeOffORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM time_offs ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (t timeOffORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.TimeOffORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM time_offs ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (t timeOffORMDo) Debug() ITimeOffORMDo {
	return t.withDO(t.DO.Debug())
}

func (t timeOffORMDo) WithContext(ctx context.Context) ITimeOffORMDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeOffORMDo) ReadDB() ITimeOffORMDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeOffORMDo) WriteDB() ITimeOffORMDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeOffORMDo) Session(config *gorm.Session) ITimeOffORMDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeOffORMDo) Clauses(conds ...clause.Expression) ITimeOffORMDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeOffORMDo) Returning(value interface{}, columns ...string) ITimeOffORMDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeOffORMDo) Not(conds ...gen.Condition) ITimeOffORMDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeOffORMDo) Or(conds ...gen.Condition) ITimeOffORMDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeOffORMDo) Select(conds ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeOffORMDo) Where(conds ...gen.Condition) ITimeOffORMDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeOffORMDo) Order(conds ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeOffORMDo) Distinct(cols ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeOffORMDo) Omit(cols ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeOffORMDo) Join(table schema.Tabler, on ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeOffORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeOffORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeOffORMDo) Group(cols ...field.Expr) ITimeOffORMDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeOffORMDo) Having(conds ...gen.Condition) ITimeOffORMDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeOffORMDo) Limit(limit int) ITimeOffORMDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeOffORMDo) Offset(offset int) ITimeOffORMDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeOffORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITimeOffORMDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeOffORMDo) Unscoped() ITimeOffORMDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeOffORMDo) Create(values ...*accounting_servicev1.TimeOffORM) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeOffORMDo) CreateInBatches(values []*accounting_servicev1.TimeOffORM, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeOffORMDo) Save(values ...*accounting_servicev1.TimeOffORM) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeOffORMDo) First() (*accounting_servicev1.TimeOffORM, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.TimeOffORM), nil
	}
}

func (t timeOffORMDo) Take() (*accounting_servicev1.TimeOffORM, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.TimeOffORM), nil
	}
}

func (t timeOffORMDo) Last() (*accounting_servicev1.TimeOffORM, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.TimeOffORM), nil
	}
}

func (t timeOffORMDo) Find() ([]*accounting_servicev1.TimeOffORM, error) {
	result, err := t.DO.Find()
	return result.([]*accounting_servicev1.TimeOffORM), err
}

func (t timeOffORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.TimeOffORM, err error) {
	buf := make([]*accounting_servicev1.TimeOffORM, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeOffORMDo) FindInBatches(result *[]*accounting_servicev1.TimeOffORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeOffORMDo) Attrs(attrs ...field.AssignExpr) ITimeOffORMDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeOffORMDo) Assign(attrs ...field.AssignExpr) ITimeOffORMDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeOffORMDo) Joins(fields ...field.RelationField) ITimeOffORMDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeOffORMDo) Preload(fields ...field.RelationField) ITimeOffORMDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeOffORMDo) FirstOrInit() (*accounting_servicev1.TimeOffORM, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.TimeOffORM), nil
	}
}

func (t timeOffORMDo) FirstOrCreate() (*accounting_servicev1.TimeOffORM, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.TimeOffORM), nil
	}
}

func (t timeOffORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.TimeOffORM, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t timeOffORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeOffORMDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeOffORMDo) Delete(models ...*accounting_servicev1.TimeOffORM) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeOffORMDo) withDO(do gen.Dao) *timeOffORMDo {
	t.DO = *do.(*gen.DO)
	return t
}
