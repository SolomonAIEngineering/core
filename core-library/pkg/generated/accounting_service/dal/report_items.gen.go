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

func newReportItemORM(db *gorm.DB, opts ...gen.DOOption) reportItemORM {
	_reportItemORM := reportItemORM{}

	_reportItemORM.reportItemORMDo.UseDB(db, opts...)
	_reportItemORM.reportItemORMDo.UseModel(&accounting_servicev1.ReportItemORM{})

	tableName := _reportItemORM.reportItemORMDo.TableName()
	_reportItemORM.ALL = field.NewAsterisk(tableName)
	_reportItemORM.AssetsBalanceSheetId = field.NewUint64(tableName, "assets_balance_sheet_id")
	_reportItemORM.Company = field.NewString(tableName, "company")
	_reportItemORM.CostOfSalesIncomeStatementId = field.NewUint64(tableName, "cost_of_sales_income_statement_id")
	_reportItemORM.EquityBalanceSheetId = field.NewUint64(tableName, "equity_balance_sheet_id")
	_reportItemORM.FinancingActivitiesCashFlowStatementId = field.NewUint64(tableName, "financing_activities_cash_flow_statement_id")
	_reportItemORM.Id = field.NewUint64(tableName, "id")
	_reportItemORM.IncomeIncomeStatementId = field.NewUint64(tableName, "income_income_statement_id")
	_reportItemORM.InvestingActivitiesCashFlowStatementId = field.NewUint64(tableName, "investing_activities_cash_flow_statement_id")
	_reportItemORM.LiabilitiesBalanceSheetId = field.NewUint64(tableName, "liabilities_balance_sheet_id")
	_reportItemORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_reportItemORM.Name = field.NewString(tableName, "name")
	_reportItemORM.NonOperatingExpensesIncomeStatementId = field.NewUint64(tableName, "non_operating_expenses_income_statement_id")
	_reportItemORM.OperatingActivitiesCashFlowStatementId = field.NewUint64(tableName, "operating_activities_cash_flow_statement_id")
	_reportItemORM.OperatingExpensesIncomeStatementId = field.NewUint64(tableName, "operating_expenses_income_statement_id")
	_reportItemORM.RemoteId = field.NewString(tableName, "remote_id")
	_reportItemORM.Value = field.NewInt64(tableName, "value")

	_reportItemORM.fillFieldMap()

	return _reportItemORM
}

type reportItemORM struct {
	reportItemORMDo

	ALL                                    field.Asterisk
	AssetsBalanceSheetId                   field.Uint64
	Company                                field.String
	CostOfSalesIncomeStatementId           field.Uint64
	EquityBalanceSheetId                   field.Uint64
	FinancingActivitiesCashFlowStatementId field.Uint64
	Id                                     field.Uint64
	IncomeIncomeStatementId                field.Uint64
	InvestingActivitiesCashFlowStatementId field.Uint64
	LiabilitiesBalanceSheetId              field.Uint64
	ModifiedAt                             field.Time
	Name                                   field.String
	NonOperatingExpensesIncomeStatementId  field.Uint64
	OperatingActivitiesCashFlowStatementId field.Uint64
	OperatingExpensesIncomeStatementId     field.Uint64
	RemoteId                               field.String
	Value                                  field.Int64

	fieldMap map[string]field.Expr
}

func (r reportItemORM) Table(newTableName string) *reportItemORM {
	r.reportItemORMDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reportItemORM) As(alias string) *reportItemORM {
	r.reportItemORMDo.DO = *(r.reportItemORMDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reportItemORM) updateTableName(table string) *reportItemORM {
	r.ALL = field.NewAsterisk(table)
	r.AssetsBalanceSheetId = field.NewUint64(table, "assets_balance_sheet_id")
	r.Company = field.NewString(table, "company")
	r.CostOfSalesIncomeStatementId = field.NewUint64(table, "cost_of_sales_income_statement_id")
	r.EquityBalanceSheetId = field.NewUint64(table, "equity_balance_sheet_id")
	r.FinancingActivitiesCashFlowStatementId = field.NewUint64(table, "financing_activities_cash_flow_statement_id")
	r.Id = field.NewUint64(table, "id")
	r.IncomeIncomeStatementId = field.NewUint64(table, "income_income_statement_id")
	r.InvestingActivitiesCashFlowStatementId = field.NewUint64(table, "investing_activities_cash_flow_statement_id")
	r.LiabilitiesBalanceSheetId = field.NewUint64(table, "liabilities_balance_sheet_id")
	r.ModifiedAt = field.NewTime(table, "modified_at")
	r.Name = field.NewString(table, "name")
	r.NonOperatingExpensesIncomeStatementId = field.NewUint64(table, "non_operating_expenses_income_statement_id")
	r.OperatingActivitiesCashFlowStatementId = field.NewUint64(table, "operating_activities_cash_flow_statement_id")
	r.OperatingExpensesIncomeStatementId = field.NewUint64(table, "operating_expenses_income_statement_id")
	r.RemoteId = field.NewString(table, "remote_id")
	r.Value = field.NewInt64(table, "value")

	r.fillFieldMap()

	return r
}

func (r *reportItemORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reportItemORM) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 16)
	r.fieldMap["assets_balance_sheet_id"] = r.AssetsBalanceSheetId
	r.fieldMap["company"] = r.Company
	r.fieldMap["cost_of_sales_income_statement_id"] = r.CostOfSalesIncomeStatementId
	r.fieldMap["equity_balance_sheet_id"] = r.EquityBalanceSheetId
	r.fieldMap["financing_activities_cash_flow_statement_id"] = r.FinancingActivitiesCashFlowStatementId
	r.fieldMap["id"] = r.Id
	r.fieldMap["income_income_statement_id"] = r.IncomeIncomeStatementId
	r.fieldMap["investing_activities_cash_flow_statement_id"] = r.InvestingActivitiesCashFlowStatementId
	r.fieldMap["liabilities_balance_sheet_id"] = r.LiabilitiesBalanceSheetId
	r.fieldMap["modified_at"] = r.ModifiedAt
	r.fieldMap["name"] = r.Name
	r.fieldMap["non_operating_expenses_income_statement_id"] = r.NonOperatingExpensesIncomeStatementId
	r.fieldMap["operating_activities_cash_flow_statement_id"] = r.OperatingActivitiesCashFlowStatementId
	r.fieldMap["operating_expenses_income_statement_id"] = r.OperatingExpensesIncomeStatementId
	r.fieldMap["remote_id"] = r.RemoteId
	r.fieldMap["value"] = r.Value
}

func (r reportItemORM) clone(db *gorm.DB) reportItemORM {
	r.reportItemORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reportItemORM) replaceDB(db *gorm.DB) reportItemORM {
	r.reportItemORMDo.ReplaceDB(db)
	return r
}

type reportItemORMDo struct{ gen.DO }

type IReportItemORMDo interface {
	gen.SubQuery
	Debug() IReportItemORMDo
	WithContext(ctx context.Context) IReportItemORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReportItemORMDo
	WriteDB() IReportItemORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReportItemORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReportItemORMDo
	Not(conds ...gen.Condition) IReportItemORMDo
	Or(conds ...gen.Condition) IReportItemORMDo
	Select(conds ...field.Expr) IReportItemORMDo
	Where(conds ...gen.Condition) IReportItemORMDo
	Order(conds ...field.Expr) IReportItemORMDo
	Distinct(cols ...field.Expr) IReportItemORMDo
	Omit(cols ...field.Expr) IReportItemORMDo
	Join(table schema.Tabler, on ...field.Expr) IReportItemORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReportItemORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReportItemORMDo
	Group(cols ...field.Expr) IReportItemORMDo
	Having(conds ...gen.Condition) IReportItemORMDo
	Limit(limit int) IReportItemORMDo
	Offset(offset int) IReportItemORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReportItemORMDo
	Unscoped() IReportItemORMDo
	Create(values ...*accounting_servicev1.ReportItemORM) error
	CreateInBatches(values []*accounting_servicev1.ReportItemORM, batchSize int) error
	Save(values ...*accounting_servicev1.ReportItemORM) error
	First() (*accounting_servicev1.ReportItemORM, error)
	Take() (*accounting_servicev1.ReportItemORM, error)
	Last() (*accounting_servicev1.ReportItemORM, error)
	Find() ([]*accounting_servicev1.ReportItemORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.ReportItemORM, err error)
	FindInBatches(result *[]*accounting_servicev1.ReportItemORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.ReportItemORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReportItemORMDo
	Assign(attrs ...field.AssignExpr) IReportItemORMDo
	Joins(fields ...field.RelationField) IReportItemORMDo
	Preload(fields ...field.RelationField) IReportItemORMDo
	FirstOrInit() (*accounting_servicev1.ReportItemORM, error)
	FirstOrCreate() (*accounting_servicev1.ReportItemORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.ReportItemORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReportItemORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.ReportItemORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.ReportItemORM, err error)
	CreateRecord(item accounting_servicev1.ReportItemORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.ReportItemORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.ReportItemORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.ReportItemORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.ReportItemORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (r reportItemORMDo) GetRecordByID(id int) (result accounting_servicev1.ReportItemORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM report_items ")
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
func (r reportItemORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.ReportItemORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM report_items ")
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
func (r reportItemORMDo) CreateRecord(item accounting_servicev1.ReportItemORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO report_items (columns) VALUES (values) ")

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
func (r reportItemORMDo) UpdateRecordByID(id int, item accounting_servicev1.ReportItemORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE report_items SET columns=values ")
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
func (r reportItemORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM report_items ")
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
func (r reportItemORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.ReportItemORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM report_items ORDER BY " + r.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (r reportItemORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM report_items ")

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
func (r reportItemORMDo) GetByID(id uint64) (result accounting_servicev1.ReportItemORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM report_items ")
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
func (r reportItemORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.ReportItemORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM report_items ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (r reportItemORMDo) Debug() IReportItemORMDo {
	return r.withDO(r.DO.Debug())
}

func (r reportItemORMDo) WithContext(ctx context.Context) IReportItemORMDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reportItemORMDo) ReadDB() IReportItemORMDo {
	return r.Clauses(dbresolver.Read)
}

func (r reportItemORMDo) WriteDB() IReportItemORMDo {
	return r.Clauses(dbresolver.Write)
}

func (r reportItemORMDo) Session(config *gorm.Session) IReportItemORMDo {
	return r.withDO(r.DO.Session(config))
}

func (r reportItemORMDo) Clauses(conds ...clause.Expression) IReportItemORMDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reportItemORMDo) Returning(value interface{}, columns ...string) IReportItemORMDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reportItemORMDo) Not(conds ...gen.Condition) IReportItemORMDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reportItemORMDo) Or(conds ...gen.Condition) IReportItemORMDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reportItemORMDo) Select(conds ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reportItemORMDo) Where(conds ...gen.Condition) IReportItemORMDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reportItemORMDo) Order(conds ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reportItemORMDo) Distinct(cols ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reportItemORMDo) Omit(cols ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reportItemORMDo) Join(table schema.Tabler, on ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reportItemORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reportItemORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reportItemORMDo) Group(cols ...field.Expr) IReportItemORMDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reportItemORMDo) Having(conds ...gen.Condition) IReportItemORMDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reportItemORMDo) Limit(limit int) IReportItemORMDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reportItemORMDo) Offset(offset int) IReportItemORMDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reportItemORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReportItemORMDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reportItemORMDo) Unscoped() IReportItemORMDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reportItemORMDo) Create(values ...*accounting_servicev1.ReportItemORM) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reportItemORMDo) CreateInBatches(values []*accounting_servicev1.ReportItemORM, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reportItemORMDo) Save(values ...*accounting_servicev1.ReportItemORM) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reportItemORMDo) First() (*accounting_servicev1.ReportItemORM, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.ReportItemORM), nil
	}
}

func (r reportItemORMDo) Take() (*accounting_servicev1.ReportItemORM, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.ReportItemORM), nil
	}
}

func (r reportItemORMDo) Last() (*accounting_servicev1.ReportItemORM, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.ReportItemORM), nil
	}
}

func (r reportItemORMDo) Find() ([]*accounting_servicev1.ReportItemORM, error) {
	result, err := r.DO.Find()
	return result.([]*accounting_servicev1.ReportItemORM), err
}

func (r reportItemORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.ReportItemORM, err error) {
	buf := make([]*accounting_servicev1.ReportItemORM, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reportItemORMDo) FindInBatches(result *[]*accounting_servicev1.ReportItemORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reportItemORMDo) Attrs(attrs ...field.AssignExpr) IReportItemORMDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reportItemORMDo) Assign(attrs ...field.AssignExpr) IReportItemORMDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reportItemORMDo) Joins(fields ...field.RelationField) IReportItemORMDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reportItemORMDo) Preload(fields ...field.RelationField) IReportItemORMDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reportItemORMDo) FirstOrInit() (*accounting_servicev1.ReportItemORM, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.ReportItemORM), nil
	}
}

func (r reportItemORMDo) FirstOrCreate() (*accounting_servicev1.ReportItemORM, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.ReportItemORM), nil
	}
}

func (r reportItemORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.ReportItemORM, count int64, err error) {
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

func (r reportItemORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reportItemORMDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reportItemORMDo) Delete(models ...*accounting_servicev1.ReportItemORM) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reportItemORMDo) withDO(do gen.Dao) *reportItemORMDo {
	r.DO = *do.(*gen.DO)
	return r
}
