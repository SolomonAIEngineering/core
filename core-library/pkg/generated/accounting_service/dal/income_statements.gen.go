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

func newIncomeStatementORM(db *gorm.DB, opts ...gen.DOOption) incomeStatementORM {
	_incomeStatementORM := incomeStatementORM{}

	_incomeStatementORM.incomeStatementORMDo.UseDB(db, opts...)
	_incomeStatementORM.incomeStatementORMDo.UseModel(&accounting_servicev1.IncomeStatementORM{})

	tableName := _incomeStatementORM.incomeStatementORMDo.TableName()
	_incomeStatementORM.ALL = field.NewAsterisk(tableName)
	_incomeStatementORM.Company = field.NewString(tableName, "company")
	_incomeStatementORM.CreatedAt = field.NewTime(tableName, "created_at")
	_incomeStatementORM.Currency = field.NewString(tableName, "currency")
	_incomeStatementORM.EndPeriod = field.NewTime(tableName, "end_period")
	_incomeStatementORM.GrossProfit = field.NewFloat64(tableName, "gross_profit")
	_incomeStatementORM.Id = field.NewUint64(tableName, "id")
	_incomeStatementORM.LinkedAccountingAccountId = field.NewUint64(tableName, "linked_accounting_account_id")
	_incomeStatementORM.MergeRecordId = field.NewString(tableName, "merge_record_id")
	_incomeStatementORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_incomeStatementORM.Name = field.NewString(tableName, "name")
	_incomeStatementORM.NetIncome = field.NewFloat64(tableName, "net_income")
	_incomeStatementORM.NetOperatingIncome = field.NewFloat64(tableName, "net_operating_income")
	_incomeStatementORM.RemoteId = field.NewString(tableName, "remote_id")
	_incomeStatementORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_incomeStatementORM.StartPeriod = field.NewTime(tableName, "start_period")
	_incomeStatementORM.CostOfSales = incomeStatementORMHasManyCostOfSales{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("CostOfSales", "accounting_servicev1.ReportItemORM"),
	}

	_incomeStatementORM.Income = incomeStatementORMHasManyIncome{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Income", "accounting_servicev1.ReportItemORM"),
	}

	_incomeStatementORM.NonOperatingExpenses = incomeStatementORMHasManyNonOperatingExpenses{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("NonOperatingExpenses", "accounting_servicev1.ReportItemORM"),
	}

	_incomeStatementORM.OperatingExpenses = incomeStatementORMHasManyOperatingExpenses{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("OperatingExpenses", "accounting_servicev1.ReportItemORM"),
	}

	_incomeStatementORM.fillFieldMap()

	return _incomeStatementORM
}

type incomeStatementORM struct {
	incomeStatementORMDo

	ALL                       field.Asterisk
	Company                   field.String
	CreatedAt                 field.Time
	Currency                  field.String
	EndPeriod                 field.Time
	GrossProfit               field.Float64
	Id                        field.Uint64
	LinkedAccountingAccountId field.Uint64
	MergeRecordId             field.String
	ModifiedAt                field.Time
	Name                      field.String
	NetIncome                 field.Float64
	NetOperatingIncome        field.Float64
	RemoteId                  field.String
	RemoteWasDeleted          field.Bool
	StartPeriod               field.Time
	CostOfSales               incomeStatementORMHasManyCostOfSales

	Income incomeStatementORMHasManyIncome

	NonOperatingExpenses incomeStatementORMHasManyNonOperatingExpenses

	OperatingExpenses incomeStatementORMHasManyOperatingExpenses

	fieldMap map[string]field.Expr
}

func (i incomeStatementORM) Table(newTableName string) *incomeStatementORM {
	i.incomeStatementORMDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i incomeStatementORM) As(alias string) *incomeStatementORM {
	i.incomeStatementORMDo.DO = *(i.incomeStatementORMDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *incomeStatementORM) updateTableName(table string) *incomeStatementORM {
	i.ALL = field.NewAsterisk(table)
	i.Company = field.NewString(table, "company")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.Currency = field.NewString(table, "currency")
	i.EndPeriod = field.NewTime(table, "end_period")
	i.GrossProfit = field.NewFloat64(table, "gross_profit")
	i.Id = field.NewUint64(table, "id")
	i.LinkedAccountingAccountId = field.NewUint64(table, "linked_accounting_account_id")
	i.MergeRecordId = field.NewString(table, "merge_record_id")
	i.ModifiedAt = field.NewTime(table, "modified_at")
	i.Name = field.NewString(table, "name")
	i.NetIncome = field.NewFloat64(table, "net_income")
	i.NetOperatingIncome = field.NewFloat64(table, "net_operating_income")
	i.RemoteId = field.NewString(table, "remote_id")
	i.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	i.StartPeriod = field.NewTime(table, "start_period")

	i.fillFieldMap()

	return i
}

func (i *incomeStatementORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *incomeStatementORM) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 19)
	i.fieldMap["company"] = i.Company
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["currency"] = i.Currency
	i.fieldMap["end_period"] = i.EndPeriod
	i.fieldMap["gross_profit"] = i.GrossProfit
	i.fieldMap["id"] = i.Id
	i.fieldMap["linked_accounting_account_id"] = i.LinkedAccountingAccountId
	i.fieldMap["merge_record_id"] = i.MergeRecordId
	i.fieldMap["modified_at"] = i.ModifiedAt
	i.fieldMap["name"] = i.Name
	i.fieldMap["net_income"] = i.NetIncome
	i.fieldMap["net_operating_income"] = i.NetOperatingIncome
	i.fieldMap["remote_id"] = i.RemoteId
	i.fieldMap["remote_was_deleted"] = i.RemoteWasDeleted
	i.fieldMap["start_period"] = i.StartPeriod

}

func (i incomeStatementORM) clone(db *gorm.DB) incomeStatementORM {
	i.incomeStatementORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i incomeStatementORM) replaceDB(db *gorm.DB) incomeStatementORM {
	i.incomeStatementORMDo.ReplaceDB(db)
	return i
}

type incomeStatementORMHasManyCostOfSales struct {
	db *gorm.DB

	field.RelationField
}

func (a incomeStatementORMHasManyCostOfSales) Where(conds ...field.Expr) *incomeStatementORMHasManyCostOfSales {
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

func (a incomeStatementORMHasManyCostOfSales) WithContext(ctx context.Context) *incomeStatementORMHasManyCostOfSales {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a incomeStatementORMHasManyCostOfSales) Session(session *gorm.Session) *incomeStatementORMHasManyCostOfSales {
	a.db = a.db.Session(session)
	return &a
}

func (a incomeStatementORMHasManyCostOfSales) Model(m *accounting_servicev1.IncomeStatementORM) *incomeStatementORMHasManyCostOfSalesTx {
	return &incomeStatementORMHasManyCostOfSalesTx{a.db.Model(m).Association(a.Name())}
}

type incomeStatementORMHasManyCostOfSalesTx struct{ tx *gorm.Association }

func (a incomeStatementORMHasManyCostOfSalesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a incomeStatementORMHasManyCostOfSalesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a incomeStatementORMHasManyCostOfSalesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a incomeStatementORMHasManyCostOfSalesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a incomeStatementORMHasManyCostOfSalesTx) Clear() error {
	return a.tx.Clear()
}

func (a incomeStatementORMHasManyCostOfSalesTx) Count() int64 {
	return a.tx.Count()
}

type incomeStatementORMHasManyIncome struct {
	db *gorm.DB

	field.RelationField
}

func (a incomeStatementORMHasManyIncome) Where(conds ...field.Expr) *incomeStatementORMHasManyIncome {
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

func (a incomeStatementORMHasManyIncome) WithContext(ctx context.Context) *incomeStatementORMHasManyIncome {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a incomeStatementORMHasManyIncome) Session(session *gorm.Session) *incomeStatementORMHasManyIncome {
	a.db = a.db.Session(session)
	return &a
}

func (a incomeStatementORMHasManyIncome) Model(m *accounting_servicev1.IncomeStatementORM) *incomeStatementORMHasManyIncomeTx {
	return &incomeStatementORMHasManyIncomeTx{a.db.Model(m).Association(a.Name())}
}

type incomeStatementORMHasManyIncomeTx struct{ tx *gorm.Association }

func (a incomeStatementORMHasManyIncomeTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a incomeStatementORMHasManyIncomeTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a incomeStatementORMHasManyIncomeTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a incomeStatementORMHasManyIncomeTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a incomeStatementORMHasManyIncomeTx) Clear() error {
	return a.tx.Clear()
}

func (a incomeStatementORMHasManyIncomeTx) Count() int64 {
	return a.tx.Count()
}

type incomeStatementORMHasManyNonOperatingExpenses struct {
	db *gorm.DB

	field.RelationField
}

func (a incomeStatementORMHasManyNonOperatingExpenses) Where(conds ...field.Expr) *incomeStatementORMHasManyNonOperatingExpenses {
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

func (a incomeStatementORMHasManyNonOperatingExpenses) WithContext(ctx context.Context) *incomeStatementORMHasManyNonOperatingExpenses {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a incomeStatementORMHasManyNonOperatingExpenses) Session(session *gorm.Session) *incomeStatementORMHasManyNonOperatingExpenses {
	a.db = a.db.Session(session)
	return &a
}

func (a incomeStatementORMHasManyNonOperatingExpenses) Model(m *accounting_servicev1.IncomeStatementORM) *incomeStatementORMHasManyNonOperatingExpensesTx {
	return &incomeStatementORMHasManyNonOperatingExpensesTx{a.db.Model(m).Association(a.Name())}
}

type incomeStatementORMHasManyNonOperatingExpensesTx struct{ tx *gorm.Association }

func (a incomeStatementORMHasManyNonOperatingExpensesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a incomeStatementORMHasManyNonOperatingExpensesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a incomeStatementORMHasManyNonOperatingExpensesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a incomeStatementORMHasManyNonOperatingExpensesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a incomeStatementORMHasManyNonOperatingExpensesTx) Clear() error {
	return a.tx.Clear()
}

func (a incomeStatementORMHasManyNonOperatingExpensesTx) Count() int64 {
	return a.tx.Count()
}

type incomeStatementORMHasManyOperatingExpenses struct {
	db *gorm.DB

	field.RelationField
}

func (a incomeStatementORMHasManyOperatingExpenses) Where(conds ...field.Expr) *incomeStatementORMHasManyOperatingExpenses {
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

func (a incomeStatementORMHasManyOperatingExpenses) WithContext(ctx context.Context) *incomeStatementORMHasManyOperatingExpenses {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a incomeStatementORMHasManyOperatingExpenses) Session(session *gorm.Session) *incomeStatementORMHasManyOperatingExpenses {
	a.db = a.db.Session(session)
	return &a
}

func (a incomeStatementORMHasManyOperatingExpenses) Model(m *accounting_servicev1.IncomeStatementORM) *incomeStatementORMHasManyOperatingExpensesTx {
	return &incomeStatementORMHasManyOperatingExpensesTx{a.db.Model(m).Association(a.Name())}
}

type incomeStatementORMHasManyOperatingExpensesTx struct{ tx *gorm.Association }

func (a incomeStatementORMHasManyOperatingExpensesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a incomeStatementORMHasManyOperatingExpensesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a incomeStatementORMHasManyOperatingExpensesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a incomeStatementORMHasManyOperatingExpensesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a incomeStatementORMHasManyOperatingExpensesTx) Clear() error {
	return a.tx.Clear()
}

func (a incomeStatementORMHasManyOperatingExpensesTx) Count() int64 {
	return a.tx.Count()
}

type incomeStatementORMDo struct{ gen.DO }

type IIncomeStatementORMDo interface {
	gen.SubQuery
	Debug() IIncomeStatementORMDo
	WithContext(ctx context.Context) IIncomeStatementORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIncomeStatementORMDo
	WriteDB() IIncomeStatementORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIncomeStatementORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIncomeStatementORMDo
	Not(conds ...gen.Condition) IIncomeStatementORMDo
	Or(conds ...gen.Condition) IIncomeStatementORMDo
	Select(conds ...field.Expr) IIncomeStatementORMDo
	Where(conds ...gen.Condition) IIncomeStatementORMDo
	Order(conds ...field.Expr) IIncomeStatementORMDo
	Distinct(cols ...field.Expr) IIncomeStatementORMDo
	Omit(cols ...field.Expr) IIncomeStatementORMDo
	Join(table schema.Tabler, on ...field.Expr) IIncomeStatementORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIncomeStatementORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIncomeStatementORMDo
	Group(cols ...field.Expr) IIncomeStatementORMDo
	Having(conds ...gen.Condition) IIncomeStatementORMDo
	Limit(limit int) IIncomeStatementORMDo
	Offset(offset int) IIncomeStatementORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIncomeStatementORMDo
	Unscoped() IIncomeStatementORMDo
	Create(values ...*accounting_servicev1.IncomeStatementORM) error
	CreateInBatches(values []*accounting_servicev1.IncomeStatementORM, batchSize int) error
	Save(values ...*accounting_servicev1.IncomeStatementORM) error
	First() (*accounting_servicev1.IncomeStatementORM, error)
	Take() (*accounting_servicev1.IncomeStatementORM, error)
	Last() (*accounting_servicev1.IncomeStatementORM, error)
	Find() ([]*accounting_servicev1.IncomeStatementORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.IncomeStatementORM, err error)
	FindInBatches(result *[]*accounting_servicev1.IncomeStatementORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.IncomeStatementORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIncomeStatementORMDo
	Assign(attrs ...field.AssignExpr) IIncomeStatementORMDo
	Joins(fields ...field.RelationField) IIncomeStatementORMDo
	Preload(fields ...field.RelationField) IIncomeStatementORMDo
	FirstOrInit() (*accounting_servicev1.IncomeStatementORM, error)
	FirstOrCreate() (*accounting_servicev1.IncomeStatementORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.IncomeStatementORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIncomeStatementORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.IncomeStatementORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.IncomeStatementORM, err error)
	CreateRecord(item accounting_servicev1.IncomeStatementORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.IncomeStatementORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.IncomeStatementORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.IncomeStatementORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.IncomeStatementORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (i incomeStatementORMDo) GetRecordByID(id int) (result accounting_servicev1.IncomeStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM income_statements ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (i incomeStatementORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.IncomeStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM income_statements ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (i incomeStatementORMDo) CreateRecord(item accounting_servicev1.IncomeStatementORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO income_statements (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (i incomeStatementORMDo) UpdateRecordByID(id int, item accounting_servicev1.IncomeStatementORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE income_statements SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (i incomeStatementORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM income_statements ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (i incomeStatementORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.IncomeStatementORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM income_statements ORDER BY " + i.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (i incomeStatementORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM income_statements ")

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (i incomeStatementORMDo) GetByID(id uint64) (result accounting_servicev1.IncomeStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM income_statements ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (i incomeStatementORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.IncomeStatementORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM income_statements ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (i incomeStatementORMDo) Debug() IIncomeStatementORMDo {
	return i.withDO(i.DO.Debug())
}

func (i incomeStatementORMDo) WithContext(ctx context.Context) IIncomeStatementORMDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i incomeStatementORMDo) ReadDB() IIncomeStatementORMDo {
	return i.Clauses(dbresolver.Read)
}

func (i incomeStatementORMDo) WriteDB() IIncomeStatementORMDo {
	return i.Clauses(dbresolver.Write)
}

func (i incomeStatementORMDo) Session(config *gorm.Session) IIncomeStatementORMDo {
	return i.withDO(i.DO.Session(config))
}

func (i incomeStatementORMDo) Clauses(conds ...clause.Expression) IIncomeStatementORMDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i incomeStatementORMDo) Returning(value interface{}, columns ...string) IIncomeStatementORMDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i incomeStatementORMDo) Not(conds ...gen.Condition) IIncomeStatementORMDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i incomeStatementORMDo) Or(conds ...gen.Condition) IIncomeStatementORMDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i incomeStatementORMDo) Select(conds ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i incomeStatementORMDo) Where(conds ...gen.Condition) IIncomeStatementORMDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i incomeStatementORMDo) Order(conds ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i incomeStatementORMDo) Distinct(cols ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i incomeStatementORMDo) Omit(cols ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i incomeStatementORMDo) Join(table schema.Tabler, on ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i incomeStatementORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i incomeStatementORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i incomeStatementORMDo) Group(cols ...field.Expr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i incomeStatementORMDo) Having(conds ...gen.Condition) IIncomeStatementORMDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i incomeStatementORMDo) Limit(limit int) IIncomeStatementORMDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i incomeStatementORMDo) Offset(offset int) IIncomeStatementORMDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i incomeStatementORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIncomeStatementORMDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i incomeStatementORMDo) Unscoped() IIncomeStatementORMDo {
	return i.withDO(i.DO.Unscoped())
}

func (i incomeStatementORMDo) Create(values ...*accounting_servicev1.IncomeStatementORM) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i incomeStatementORMDo) CreateInBatches(values []*accounting_servicev1.IncomeStatementORM, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i incomeStatementORMDo) Save(values ...*accounting_servicev1.IncomeStatementORM) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i incomeStatementORMDo) First() (*accounting_servicev1.IncomeStatementORM, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.IncomeStatementORM), nil
	}
}

func (i incomeStatementORMDo) Take() (*accounting_servicev1.IncomeStatementORM, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.IncomeStatementORM), nil
	}
}

func (i incomeStatementORMDo) Last() (*accounting_servicev1.IncomeStatementORM, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.IncomeStatementORM), nil
	}
}

func (i incomeStatementORMDo) Find() ([]*accounting_servicev1.IncomeStatementORM, error) {
	result, err := i.DO.Find()
	return result.([]*accounting_servicev1.IncomeStatementORM), err
}

func (i incomeStatementORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.IncomeStatementORM, err error) {
	buf := make([]*accounting_servicev1.IncomeStatementORM, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i incomeStatementORMDo) FindInBatches(result *[]*accounting_servicev1.IncomeStatementORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i incomeStatementORMDo) Attrs(attrs ...field.AssignExpr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i incomeStatementORMDo) Assign(attrs ...field.AssignExpr) IIncomeStatementORMDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i incomeStatementORMDo) Joins(fields ...field.RelationField) IIncomeStatementORMDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i incomeStatementORMDo) Preload(fields ...field.RelationField) IIncomeStatementORMDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i incomeStatementORMDo) FirstOrInit() (*accounting_servicev1.IncomeStatementORM, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.IncomeStatementORM), nil
	}
}

func (i incomeStatementORMDo) FirstOrCreate() (*accounting_servicev1.IncomeStatementORM, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.IncomeStatementORM), nil
	}
}

func (i incomeStatementORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.IncomeStatementORM, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i incomeStatementORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i incomeStatementORMDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i incomeStatementORMDo) Delete(models ...*accounting_servicev1.IncomeStatementORM) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *incomeStatementORMDo) withDO(do gen.Dao) *incomeStatementORMDo {
	i.DO = *do.(*gen.DO)
	return i
}
