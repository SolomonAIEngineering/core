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

func newCashFlowStatementORM(db *gorm.DB, opts ...gen.DOOption) cashFlowStatementORM {
	_cashFlowStatementORM := cashFlowStatementORM{}

	_cashFlowStatementORM.cashFlowStatementORMDo.UseDB(db, opts...)
	_cashFlowStatementORM.cashFlowStatementORMDo.UseModel(&accounting_servicev1.CashFlowStatementORM{})

	tableName := _cashFlowStatementORM.cashFlowStatementORMDo.TableName()
	_cashFlowStatementORM.ALL = field.NewAsterisk(tableName)
	_cashFlowStatementORM.CashAtBeginningOfPeriod = field.NewFloat64(tableName, "cash_at_beginning_of_period")
	_cashFlowStatementORM.CashAtEndOfPeriod = field.NewFloat64(tableName, "cash_at_end_of_period")
	_cashFlowStatementORM.Company = field.NewString(tableName, "company")
	_cashFlowStatementORM.Currency = field.NewString(tableName, "currency")
	_cashFlowStatementORM.EndPeriod = field.NewTime(tableName, "end_period")
	_cashFlowStatementORM.Id = field.NewUint64(tableName, "id")
	_cashFlowStatementORM.LinkedAccountingAccountId = field.NewUint64(tableName, "linked_accounting_account_id")
	_cashFlowStatementORM.MergeRecordId = field.NewString(tableName, "merge_record_id")
	_cashFlowStatementORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_cashFlowStatementORM.Name = field.NewString(tableName, "name")
	_cashFlowStatementORM.RemoteGeneratedAt = field.NewTime(tableName, "remote_generated_at")
	_cashFlowStatementORM.RemoteId = field.NewString(tableName, "remote_id")
	_cashFlowStatementORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_cashFlowStatementORM.StartPeriod = field.NewTime(tableName, "start_period")
	_cashFlowStatementORM.FinancingActivities = cashFlowStatementORMHasManyFinancingActivities{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("FinancingActivities", "accounting_servicev1.ReportItemORM"),
	}

	_cashFlowStatementORM.InvestingActivities = cashFlowStatementORMHasManyInvestingActivities{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("InvestingActivities", "accounting_servicev1.ReportItemORM"),
	}

	_cashFlowStatementORM.OperatingActivities = cashFlowStatementORMHasManyOperatingActivities{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("OperatingActivities", "accounting_servicev1.ReportItemORM"),
	}

	_cashFlowStatementORM.fillFieldMap()

	return _cashFlowStatementORM
}

type cashFlowStatementORM struct {
	cashFlowStatementORMDo

	ALL                       field.Asterisk
	CashAtBeginningOfPeriod   field.Float64
	CashAtEndOfPeriod         field.Float64
	Company                   field.String
	Currency                  field.String
	EndPeriod                 field.Time
	Id                        field.Uint64
	LinkedAccountingAccountId field.Uint64
	MergeRecordId             field.String
	ModifiedAt                field.Time
	Name                      field.String
	RemoteGeneratedAt         field.Time
	RemoteId                  field.String
	RemoteWasDeleted          field.Bool
	StartPeriod               field.Time
	FinancingActivities       cashFlowStatementORMHasManyFinancingActivities

	InvestingActivities cashFlowStatementORMHasManyInvestingActivities

	OperatingActivities cashFlowStatementORMHasManyOperatingActivities

	fieldMap map[string]field.Expr
}

func (c cashFlowStatementORM) Table(newTableName string) *cashFlowStatementORM {
	c.cashFlowStatementORMDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cashFlowStatementORM) As(alias string) *cashFlowStatementORM {
	c.cashFlowStatementORMDo.DO = *(c.cashFlowStatementORMDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cashFlowStatementORM) updateTableName(table string) *cashFlowStatementORM {
	c.ALL = field.NewAsterisk(table)
	c.CashAtBeginningOfPeriod = field.NewFloat64(table, "cash_at_beginning_of_period")
	c.CashAtEndOfPeriod = field.NewFloat64(table, "cash_at_end_of_period")
	c.Company = field.NewString(table, "company")
	c.Currency = field.NewString(table, "currency")
	c.EndPeriod = field.NewTime(table, "end_period")
	c.Id = field.NewUint64(table, "id")
	c.LinkedAccountingAccountId = field.NewUint64(table, "linked_accounting_account_id")
	c.MergeRecordId = field.NewString(table, "merge_record_id")
	c.ModifiedAt = field.NewTime(table, "modified_at")
	c.Name = field.NewString(table, "name")
	c.RemoteGeneratedAt = field.NewTime(table, "remote_generated_at")
	c.RemoteId = field.NewString(table, "remote_id")
	c.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	c.StartPeriod = field.NewTime(table, "start_period")

	c.fillFieldMap()

	return c
}

func (c *cashFlowStatementORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cashFlowStatementORM) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 17)
	c.fieldMap["cash_at_beginning_of_period"] = c.CashAtBeginningOfPeriod
	c.fieldMap["cash_at_end_of_period"] = c.CashAtEndOfPeriod
	c.fieldMap["company"] = c.Company
	c.fieldMap["currency"] = c.Currency
	c.fieldMap["end_period"] = c.EndPeriod
	c.fieldMap["id"] = c.Id
	c.fieldMap["linked_accounting_account_id"] = c.LinkedAccountingAccountId
	c.fieldMap["merge_record_id"] = c.MergeRecordId
	c.fieldMap["modified_at"] = c.ModifiedAt
	c.fieldMap["name"] = c.Name
	c.fieldMap["remote_generated_at"] = c.RemoteGeneratedAt
	c.fieldMap["remote_id"] = c.RemoteId
	c.fieldMap["remote_was_deleted"] = c.RemoteWasDeleted
	c.fieldMap["start_period"] = c.StartPeriod

}

func (c cashFlowStatementORM) clone(db *gorm.DB) cashFlowStatementORM {
	c.cashFlowStatementORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cashFlowStatementORM) replaceDB(db *gorm.DB) cashFlowStatementORM {
	c.cashFlowStatementORMDo.ReplaceDB(db)
	return c
}

type cashFlowStatementORMHasManyFinancingActivities struct {
	db *gorm.DB

	field.RelationField
}

func (a cashFlowStatementORMHasManyFinancingActivities) Where(conds ...field.Expr) *cashFlowStatementORMHasManyFinancingActivities {
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

func (a cashFlowStatementORMHasManyFinancingActivities) WithContext(ctx context.Context) *cashFlowStatementORMHasManyFinancingActivities {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a cashFlowStatementORMHasManyFinancingActivities) Session(session *gorm.Session) *cashFlowStatementORMHasManyFinancingActivities {
	a.db = a.db.Session(session)
	return &a
}

func (a cashFlowStatementORMHasManyFinancingActivities) Model(m *accounting_servicev1.CashFlowStatementORM) *cashFlowStatementORMHasManyFinancingActivitiesTx {
	return &cashFlowStatementORMHasManyFinancingActivitiesTx{a.db.Model(m).Association(a.Name())}
}

type cashFlowStatementORMHasManyFinancingActivitiesTx struct{ tx *gorm.Association }

func (a cashFlowStatementORMHasManyFinancingActivitiesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a cashFlowStatementORMHasManyFinancingActivitiesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a cashFlowStatementORMHasManyFinancingActivitiesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a cashFlowStatementORMHasManyFinancingActivitiesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a cashFlowStatementORMHasManyFinancingActivitiesTx) Clear() error {
	return a.tx.Clear()
}

func (a cashFlowStatementORMHasManyFinancingActivitiesTx) Count() int64 {
	return a.tx.Count()
}

type cashFlowStatementORMHasManyInvestingActivities struct {
	db *gorm.DB

	field.RelationField
}

func (a cashFlowStatementORMHasManyInvestingActivities) Where(conds ...field.Expr) *cashFlowStatementORMHasManyInvestingActivities {
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

func (a cashFlowStatementORMHasManyInvestingActivities) WithContext(ctx context.Context) *cashFlowStatementORMHasManyInvestingActivities {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a cashFlowStatementORMHasManyInvestingActivities) Session(session *gorm.Session) *cashFlowStatementORMHasManyInvestingActivities {
	a.db = a.db.Session(session)
	return &a
}

func (a cashFlowStatementORMHasManyInvestingActivities) Model(m *accounting_servicev1.CashFlowStatementORM) *cashFlowStatementORMHasManyInvestingActivitiesTx {
	return &cashFlowStatementORMHasManyInvestingActivitiesTx{a.db.Model(m).Association(a.Name())}
}

type cashFlowStatementORMHasManyInvestingActivitiesTx struct{ tx *gorm.Association }

func (a cashFlowStatementORMHasManyInvestingActivitiesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a cashFlowStatementORMHasManyInvestingActivitiesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a cashFlowStatementORMHasManyInvestingActivitiesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a cashFlowStatementORMHasManyInvestingActivitiesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a cashFlowStatementORMHasManyInvestingActivitiesTx) Clear() error {
	return a.tx.Clear()
}

func (a cashFlowStatementORMHasManyInvestingActivitiesTx) Count() int64 {
	return a.tx.Count()
}

type cashFlowStatementORMHasManyOperatingActivities struct {
	db *gorm.DB

	field.RelationField
}

func (a cashFlowStatementORMHasManyOperatingActivities) Where(conds ...field.Expr) *cashFlowStatementORMHasManyOperatingActivities {
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

func (a cashFlowStatementORMHasManyOperatingActivities) WithContext(ctx context.Context) *cashFlowStatementORMHasManyOperatingActivities {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a cashFlowStatementORMHasManyOperatingActivities) Session(session *gorm.Session) *cashFlowStatementORMHasManyOperatingActivities {
	a.db = a.db.Session(session)
	return &a
}

func (a cashFlowStatementORMHasManyOperatingActivities) Model(m *accounting_servicev1.CashFlowStatementORM) *cashFlowStatementORMHasManyOperatingActivitiesTx {
	return &cashFlowStatementORMHasManyOperatingActivitiesTx{a.db.Model(m).Association(a.Name())}
}

type cashFlowStatementORMHasManyOperatingActivitiesTx struct{ tx *gorm.Association }

func (a cashFlowStatementORMHasManyOperatingActivitiesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a cashFlowStatementORMHasManyOperatingActivitiesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a cashFlowStatementORMHasManyOperatingActivitiesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a cashFlowStatementORMHasManyOperatingActivitiesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a cashFlowStatementORMHasManyOperatingActivitiesTx) Clear() error {
	return a.tx.Clear()
}

func (a cashFlowStatementORMHasManyOperatingActivitiesTx) Count() int64 {
	return a.tx.Count()
}

type cashFlowStatementORMDo struct{ gen.DO }

type ICashFlowStatementORMDo interface {
	gen.SubQuery
	Debug() ICashFlowStatementORMDo
	WithContext(ctx context.Context) ICashFlowStatementORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICashFlowStatementORMDo
	WriteDB() ICashFlowStatementORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICashFlowStatementORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICashFlowStatementORMDo
	Not(conds ...gen.Condition) ICashFlowStatementORMDo
	Or(conds ...gen.Condition) ICashFlowStatementORMDo
	Select(conds ...field.Expr) ICashFlowStatementORMDo
	Where(conds ...gen.Condition) ICashFlowStatementORMDo
	Order(conds ...field.Expr) ICashFlowStatementORMDo
	Distinct(cols ...field.Expr) ICashFlowStatementORMDo
	Omit(cols ...field.Expr) ICashFlowStatementORMDo
	Join(table schema.Tabler, on ...field.Expr) ICashFlowStatementORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICashFlowStatementORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICashFlowStatementORMDo
	Group(cols ...field.Expr) ICashFlowStatementORMDo
	Having(conds ...gen.Condition) ICashFlowStatementORMDo
	Limit(limit int) ICashFlowStatementORMDo
	Offset(offset int) ICashFlowStatementORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICashFlowStatementORMDo
	Unscoped() ICashFlowStatementORMDo
	Create(values ...*accounting_servicev1.CashFlowStatementORM) error
	CreateInBatches(values []*accounting_servicev1.CashFlowStatementORM, batchSize int) error
	Save(values ...*accounting_servicev1.CashFlowStatementORM) error
	First() (*accounting_servicev1.CashFlowStatementORM, error)
	Take() (*accounting_servicev1.CashFlowStatementORM, error)
	Last() (*accounting_servicev1.CashFlowStatementORM, error)
	Find() ([]*accounting_servicev1.CashFlowStatementORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.CashFlowStatementORM, err error)
	FindInBatches(result *[]*accounting_servicev1.CashFlowStatementORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.CashFlowStatementORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICashFlowStatementORMDo
	Assign(attrs ...field.AssignExpr) ICashFlowStatementORMDo
	Joins(fields ...field.RelationField) ICashFlowStatementORMDo
	Preload(fields ...field.RelationField) ICashFlowStatementORMDo
	FirstOrInit() (*accounting_servicev1.CashFlowStatementORM, error)
	FirstOrCreate() (*accounting_servicev1.CashFlowStatementORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.CashFlowStatementORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICashFlowStatementORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.CashFlowStatementORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.CashFlowStatementORM, err error)
	CreateRecord(item accounting_servicev1.CashFlowStatementORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.CashFlowStatementORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.CashFlowStatementORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.CashFlowStatementORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.CashFlowStatementORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c cashFlowStatementORMDo) GetRecordByID(id int) (result accounting_servicev1.CashFlowStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM cash_flow_statements ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (c cashFlowStatementORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.CashFlowStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM cash_flow_statements ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (c cashFlowStatementORMDo) CreateRecord(item accounting_servicev1.CashFlowStatementORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO cash_flow_statements (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (c cashFlowStatementORMDo) UpdateRecordByID(id int, item accounting_servicev1.CashFlowStatementORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE cash_flow_statements SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c cashFlowStatementORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM cash_flow_statements ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (c cashFlowStatementORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.CashFlowStatementORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM cash_flow_statements ORDER BY " + c.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (c cashFlowStatementORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM cash_flow_statements ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c cashFlowStatementORMDo) GetByID(id uint64) (result accounting_servicev1.CashFlowStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM cash_flow_statements ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (c cashFlowStatementORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.CashFlowStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM cash_flow_statements ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c cashFlowStatementORMDo) Debug() ICashFlowStatementORMDo {
	return c.withDO(c.DO.Debug())
}

func (c cashFlowStatementORMDo) WithContext(ctx context.Context) ICashFlowStatementORMDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cashFlowStatementORMDo) ReadDB() ICashFlowStatementORMDo {
	return c.Clauses(dbresolver.Read)
}

func (c cashFlowStatementORMDo) WriteDB() ICashFlowStatementORMDo {
	return c.Clauses(dbresolver.Write)
}

func (c cashFlowStatementORMDo) Session(config *gorm.Session) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Session(config))
}

func (c cashFlowStatementORMDo) Clauses(conds ...clause.Expression) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cashFlowStatementORMDo) Returning(value interface{}, columns ...string) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cashFlowStatementORMDo) Not(conds ...gen.Condition) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cashFlowStatementORMDo) Or(conds ...gen.Condition) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cashFlowStatementORMDo) Select(conds ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cashFlowStatementORMDo) Where(conds ...gen.Condition) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cashFlowStatementORMDo) Order(conds ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cashFlowStatementORMDo) Distinct(cols ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cashFlowStatementORMDo) Omit(cols ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cashFlowStatementORMDo) Join(table schema.Tabler, on ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cashFlowStatementORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cashFlowStatementORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cashFlowStatementORMDo) Group(cols ...field.Expr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cashFlowStatementORMDo) Having(conds ...gen.Condition) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cashFlowStatementORMDo) Limit(limit int) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cashFlowStatementORMDo) Offset(offset int) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cashFlowStatementORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cashFlowStatementORMDo) Unscoped() ICashFlowStatementORMDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cashFlowStatementORMDo) Create(values ...*accounting_servicev1.CashFlowStatementORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cashFlowStatementORMDo) CreateInBatches(values []*accounting_servicev1.CashFlowStatementORM, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cashFlowStatementORMDo) Save(values ...*accounting_servicev1.CashFlowStatementORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cashFlowStatementORMDo) First() (*accounting_servicev1.CashFlowStatementORM, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CashFlowStatementORM), nil
	}
}

func (c cashFlowStatementORMDo) Take() (*accounting_servicev1.CashFlowStatementORM, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CashFlowStatementORM), nil
	}
}

func (c cashFlowStatementORMDo) Last() (*accounting_servicev1.CashFlowStatementORM, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CashFlowStatementORM), nil
	}
}

func (c cashFlowStatementORMDo) Find() ([]*accounting_servicev1.CashFlowStatementORM, error) {
	result, err := c.DO.Find()
	return result.([]*accounting_servicev1.CashFlowStatementORM), err
}

func (c cashFlowStatementORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.CashFlowStatementORM, err error) {
	buf := make([]*accounting_servicev1.CashFlowStatementORM, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cashFlowStatementORMDo) FindInBatches(result *[]*accounting_servicev1.CashFlowStatementORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cashFlowStatementORMDo) Attrs(attrs ...field.AssignExpr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cashFlowStatementORMDo) Assign(attrs ...field.AssignExpr) ICashFlowStatementORMDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cashFlowStatementORMDo) Joins(fields ...field.RelationField) ICashFlowStatementORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cashFlowStatementORMDo) Preload(fields ...field.RelationField) ICashFlowStatementORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cashFlowStatementORMDo) FirstOrInit() (*accounting_servicev1.CashFlowStatementORM, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CashFlowStatementORM), nil
	}
}

func (c cashFlowStatementORMDo) FirstOrCreate() (*accounting_servicev1.CashFlowStatementORM, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CashFlowStatementORM), nil
	}
}

func (c cashFlowStatementORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.CashFlowStatementORM, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cashFlowStatementORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cashFlowStatementORMDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cashFlowStatementORMDo) Delete(models ...*accounting_servicev1.CashFlowStatementORM) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cashFlowStatementORMDo) withDO(do gen.Dao) *cashFlowStatementORMDo {
	c.DO = *do.(*gen.DO)
	return c
}
