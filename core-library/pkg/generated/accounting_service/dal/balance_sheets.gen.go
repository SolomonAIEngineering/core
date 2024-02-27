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

func newBalanceSheetORM(db *gorm.DB, opts ...gen.DOOption) balanceSheetORM {
	_balanceSheetORM := balanceSheetORM{}

	_balanceSheetORM.balanceSheetORMDo.UseDB(db, opts...)
	_balanceSheetORM.balanceSheetORMDo.UseModel(&accounting_servicev1.BalanceSheetORM{})

	tableName := _balanceSheetORM.balanceSheetORMDo.TableName()
	_balanceSheetORM.ALL = field.NewAsterisk(tableName)
	_balanceSheetORM.Company = field.NewString(tableName, "company")
	_balanceSheetORM.CreatedAt = field.NewTime(tableName, "created_at")
	_balanceSheetORM.Currency = field.NewString(tableName, "currency")
	_balanceSheetORM.Date = field.NewTime(tableName, "date")
	_balanceSheetORM.Id = field.NewUint64(tableName, "id")
	_balanceSheetORM.LinkedAccountingAccountId = field.NewUint64(tableName, "linked_accounting_account_id")
	_balanceSheetORM.MergeRecordId = field.NewString(tableName, "merge_record_id")
	_balanceSheetORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_balanceSheetORM.Name = field.NewString(tableName, "name")
	_balanceSheetORM.NetAssets = field.NewFloat64(tableName, "net_assets")
	_balanceSheetORM.RemoteGeneratedAt = field.NewTime(tableName, "remote_generated_at")
	_balanceSheetORM.RemoteId = field.NewString(tableName, "remote_id")
	_balanceSheetORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_balanceSheetORM.Assets = balanceSheetORMHasManyAssets{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Assets", "accounting_servicev1.ReportItemORM"),
	}

	_balanceSheetORM.Equity = balanceSheetORMHasManyEquity{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Equity", "accounting_servicev1.ReportItemORM"),
	}

	_balanceSheetORM.Liabilities = balanceSheetORMHasManyLiabilities{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Liabilities", "accounting_servicev1.ReportItemORM"),
	}

	_balanceSheetORM.fillFieldMap()

	return _balanceSheetORM
}

type balanceSheetORM struct {
	balanceSheetORMDo

	ALL                       field.Asterisk
	Company                   field.String
	CreatedAt                 field.Time
	Currency                  field.String
	Date                      field.Time
	Id                        field.Uint64
	LinkedAccountingAccountId field.Uint64
	MergeRecordId             field.String
	ModifiedAt                field.Time
	Name                      field.String
	NetAssets                 field.Float64
	RemoteGeneratedAt         field.Time
	RemoteId                  field.String
	RemoteWasDeleted          field.Bool
	Assets                    balanceSheetORMHasManyAssets

	Equity balanceSheetORMHasManyEquity

	Liabilities balanceSheetORMHasManyLiabilities

	fieldMap map[string]field.Expr
}

func (b balanceSheetORM) Table(newTableName string) *balanceSheetORM {
	b.balanceSheetORMDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b balanceSheetORM) As(alias string) *balanceSheetORM {
	b.balanceSheetORMDo.DO = *(b.balanceSheetORMDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *balanceSheetORM) updateTableName(table string) *balanceSheetORM {
	b.ALL = field.NewAsterisk(table)
	b.Company = field.NewString(table, "company")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.Currency = field.NewString(table, "currency")
	b.Date = field.NewTime(table, "date")
	b.Id = field.NewUint64(table, "id")
	b.LinkedAccountingAccountId = field.NewUint64(table, "linked_accounting_account_id")
	b.MergeRecordId = field.NewString(table, "merge_record_id")
	b.ModifiedAt = field.NewTime(table, "modified_at")
	b.Name = field.NewString(table, "name")
	b.NetAssets = field.NewFloat64(table, "net_assets")
	b.RemoteGeneratedAt = field.NewTime(table, "remote_generated_at")
	b.RemoteId = field.NewString(table, "remote_id")
	b.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")

	b.fillFieldMap()

	return b
}

func (b *balanceSheetORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *balanceSheetORM) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 16)
	b.fieldMap["company"] = b.Company
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["currency"] = b.Currency
	b.fieldMap["date"] = b.Date
	b.fieldMap["id"] = b.Id
	b.fieldMap["linked_accounting_account_id"] = b.LinkedAccountingAccountId
	b.fieldMap["merge_record_id"] = b.MergeRecordId
	b.fieldMap["modified_at"] = b.ModifiedAt
	b.fieldMap["name"] = b.Name
	b.fieldMap["net_assets"] = b.NetAssets
	b.fieldMap["remote_generated_at"] = b.RemoteGeneratedAt
	b.fieldMap["remote_id"] = b.RemoteId
	b.fieldMap["remote_was_deleted"] = b.RemoteWasDeleted

}

func (b balanceSheetORM) clone(db *gorm.DB) balanceSheetORM {
	b.balanceSheetORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b balanceSheetORM) replaceDB(db *gorm.DB) balanceSheetORM {
	b.balanceSheetORMDo.ReplaceDB(db)
	return b
}

type balanceSheetORMHasManyAssets struct {
	db *gorm.DB

	field.RelationField
}

func (a balanceSheetORMHasManyAssets) Where(conds ...field.Expr) *balanceSheetORMHasManyAssets {
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

func (a balanceSheetORMHasManyAssets) WithContext(ctx context.Context) *balanceSheetORMHasManyAssets {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a balanceSheetORMHasManyAssets) Session(session *gorm.Session) *balanceSheetORMHasManyAssets {
	a.db = a.db.Session(session)
	return &a
}

func (a balanceSheetORMHasManyAssets) Model(m *accounting_servicev1.BalanceSheetORM) *balanceSheetORMHasManyAssetsTx {
	return &balanceSheetORMHasManyAssetsTx{a.db.Model(m).Association(a.Name())}
}

type balanceSheetORMHasManyAssetsTx struct{ tx *gorm.Association }

func (a balanceSheetORMHasManyAssetsTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a balanceSheetORMHasManyAssetsTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a balanceSheetORMHasManyAssetsTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a balanceSheetORMHasManyAssetsTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a balanceSheetORMHasManyAssetsTx) Clear() error {
	return a.tx.Clear()
}

func (a balanceSheetORMHasManyAssetsTx) Count() int64 {
	return a.tx.Count()
}

type balanceSheetORMHasManyEquity struct {
	db *gorm.DB

	field.RelationField
}

func (a balanceSheetORMHasManyEquity) Where(conds ...field.Expr) *balanceSheetORMHasManyEquity {
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

func (a balanceSheetORMHasManyEquity) WithContext(ctx context.Context) *balanceSheetORMHasManyEquity {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a balanceSheetORMHasManyEquity) Session(session *gorm.Session) *balanceSheetORMHasManyEquity {
	a.db = a.db.Session(session)
	return &a
}

func (a balanceSheetORMHasManyEquity) Model(m *accounting_servicev1.BalanceSheetORM) *balanceSheetORMHasManyEquityTx {
	return &balanceSheetORMHasManyEquityTx{a.db.Model(m).Association(a.Name())}
}

type balanceSheetORMHasManyEquityTx struct{ tx *gorm.Association }

func (a balanceSheetORMHasManyEquityTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a balanceSheetORMHasManyEquityTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a balanceSheetORMHasManyEquityTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a balanceSheetORMHasManyEquityTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a balanceSheetORMHasManyEquityTx) Clear() error {
	return a.tx.Clear()
}

func (a balanceSheetORMHasManyEquityTx) Count() int64 {
	return a.tx.Count()
}

type balanceSheetORMHasManyLiabilities struct {
	db *gorm.DB

	field.RelationField
}

func (a balanceSheetORMHasManyLiabilities) Where(conds ...field.Expr) *balanceSheetORMHasManyLiabilities {
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

func (a balanceSheetORMHasManyLiabilities) WithContext(ctx context.Context) *balanceSheetORMHasManyLiabilities {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a balanceSheetORMHasManyLiabilities) Session(session *gorm.Session) *balanceSheetORMHasManyLiabilities {
	a.db = a.db.Session(session)
	return &a
}

func (a balanceSheetORMHasManyLiabilities) Model(m *accounting_servicev1.BalanceSheetORM) *balanceSheetORMHasManyLiabilitiesTx {
	return &balanceSheetORMHasManyLiabilitiesTx{a.db.Model(m).Association(a.Name())}
}

type balanceSheetORMHasManyLiabilitiesTx struct{ tx *gorm.Association }

func (a balanceSheetORMHasManyLiabilitiesTx) Find() (result []*accounting_servicev1.ReportItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a balanceSheetORMHasManyLiabilitiesTx) Append(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a balanceSheetORMHasManyLiabilitiesTx) Replace(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a balanceSheetORMHasManyLiabilitiesTx) Delete(values ...*accounting_servicev1.ReportItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a balanceSheetORMHasManyLiabilitiesTx) Clear() error {
	return a.tx.Clear()
}

func (a balanceSheetORMHasManyLiabilitiesTx) Count() int64 {
	return a.tx.Count()
}

type balanceSheetORMDo struct{ gen.DO }

type IBalanceSheetORMDo interface {
	gen.SubQuery
	Debug() IBalanceSheetORMDo
	WithContext(ctx context.Context) IBalanceSheetORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBalanceSheetORMDo
	WriteDB() IBalanceSheetORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBalanceSheetORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBalanceSheetORMDo
	Not(conds ...gen.Condition) IBalanceSheetORMDo
	Or(conds ...gen.Condition) IBalanceSheetORMDo
	Select(conds ...field.Expr) IBalanceSheetORMDo
	Where(conds ...gen.Condition) IBalanceSheetORMDo
	Order(conds ...field.Expr) IBalanceSheetORMDo
	Distinct(cols ...field.Expr) IBalanceSheetORMDo
	Omit(cols ...field.Expr) IBalanceSheetORMDo
	Join(table schema.Tabler, on ...field.Expr) IBalanceSheetORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBalanceSheetORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBalanceSheetORMDo
	Group(cols ...field.Expr) IBalanceSheetORMDo
	Having(conds ...gen.Condition) IBalanceSheetORMDo
	Limit(limit int) IBalanceSheetORMDo
	Offset(offset int) IBalanceSheetORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBalanceSheetORMDo
	Unscoped() IBalanceSheetORMDo
	Create(values ...*accounting_servicev1.BalanceSheetORM) error
	CreateInBatches(values []*accounting_servicev1.BalanceSheetORM, batchSize int) error
	Save(values ...*accounting_servicev1.BalanceSheetORM) error
	First() (*accounting_servicev1.BalanceSheetORM, error)
	Take() (*accounting_servicev1.BalanceSheetORM, error)
	Last() (*accounting_servicev1.BalanceSheetORM, error)
	Find() ([]*accounting_servicev1.BalanceSheetORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.BalanceSheetORM, err error)
	FindInBatches(result *[]*accounting_servicev1.BalanceSheetORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.BalanceSheetORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBalanceSheetORMDo
	Assign(attrs ...field.AssignExpr) IBalanceSheetORMDo
	Joins(fields ...field.RelationField) IBalanceSheetORMDo
	Preload(fields ...field.RelationField) IBalanceSheetORMDo
	FirstOrInit() (*accounting_servicev1.BalanceSheetORM, error)
	FirstOrCreate() (*accounting_servicev1.BalanceSheetORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.BalanceSheetORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBalanceSheetORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.BalanceSheetORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.BalanceSheetORM, err error)
	CreateRecord(item accounting_servicev1.BalanceSheetORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.BalanceSheetORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.BalanceSheetORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.BalanceSheetORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.BalanceSheetORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (b balanceSheetORMDo) GetRecordByID(id int) (result accounting_servicev1.BalanceSheetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM balance_sheets ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (b balanceSheetORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.BalanceSheetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM balance_sheets ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (b balanceSheetORMDo) CreateRecord(item accounting_servicev1.BalanceSheetORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO balance_sheets (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (b balanceSheetORMDo) UpdateRecordByID(id int, item accounting_servicev1.BalanceSheetORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE balance_sheets SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (b balanceSheetORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM balance_sheets ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (b balanceSheetORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.BalanceSheetORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM balance_sheets ORDER BY " + b.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (b balanceSheetORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM balance_sheets ")

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (b balanceSheetORMDo) GetByID(id uint64) (result accounting_servicev1.BalanceSheetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM balance_sheets ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (b balanceSheetORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.BalanceSheetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM balance_sheets ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (b balanceSheetORMDo) Debug() IBalanceSheetORMDo {
	return b.withDO(b.DO.Debug())
}

func (b balanceSheetORMDo) WithContext(ctx context.Context) IBalanceSheetORMDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b balanceSheetORMDo) ReadDB() IBalanceSheetORMDo {
	return b.Clauses(dbresolver.Read)
}

func (b balanceSheetORMDo) WriteDB() IBalanceSheetORMDo {
	return b.Clauses(dbresolver.Write)
}

func (b balanceSheetORMDo) Session(config *gorm.Session) IBalanceSheetORMDo {
	return b.withDO(b.DO.Session(config))
}

func (b balanceSheetORMDo) Clauses(conds ...clause.Expression) IBalanceSheetORMDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b balanceSheetORMDo) Returning(value interface{}, columns ...string) IBalanceSheetORMDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b balanceSheetORMDo) Not(conds ...gen.Condition) IBalanceSheetORMDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b balanceSheetORMDo) Or(conds ...gen.Condition) IBalanceSheetORMDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b balanceSheetORMDo) Select(conds ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b balanceSheetORMDo) Where(conds ...gen.Condition) IBalanceSheetORMDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b balanceSheetORMDo) Order(conds ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b balanceSheetORMDo) Distinct(cols ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b balanceSheetORMDo) Omit(cols ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b balanceSheetORMDo) Join(table schema.Tabler, on ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b balanceSheetORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b balanceSheetORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b balanceSheetORMDo) Group(cols ...field.Expr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b balanceSheetORMDo) Having(conds ...gen.Condition) IBalanceSheetORMDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b balanceSheetORMDo) Limit(limit int) IBalanceSheetORMDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b balanceSheetORMDo) Offset(offset int) IBalanceSheetORMDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b balanceSheetORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBalanceSheetORMDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b balanceSheetORMDo) Unscoped() IBalanceSheetORMDo {
	return b.withDO(b.DO.Unscoped())
}

func (b balanceSheetORMDo) Create(values ...*accounting_servicev1.BalanceSheetORM) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b balanceSheetORMDo) CreateInBatches(values []*accounting_servicev1.BalanceSheetORM, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b balanceSheetORMDo) Save(values ...*accounting_servicev1.BalanceSheetORM) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b balanceSheetORMDo) First() (*accounting_servicev1.BalanceSheetORM, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.BalanceSheetORM), nil
	}
}

func (b balanceSheetORMDo) Take() (*accounting_servicev1.BalanceSheetORM, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.BalanceSheetORM), nil
	}
}

func (b balanceSheetORMDo) Last() (*accounting_servicev1.BalanceSheetORM, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.BalanceSheetORM), nil
	}
}

func (b balanceSheetORMDo) Find() ([]*accounting_servicev1.BalanceSheetORM, error) {
	result, err := b.DO.Find()
	return result.([]*accounting_servicev1.BalanceSheetORM), err
}

func (b balanceSheetORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.BalanceSheetORM, err error) {
	buf := make([]*accounting_servicev1.BalanceSheetORM, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b balanceSheetORMDo) FindInBatches(result *[]*accounting_servicev1.BalanceSheetORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b balanceSheetORMDo) Attrs(attrs ...field.AssignExpr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b balanceSheetORMDo) Assign(attrs ...field.AssignExpr) IBalanceSheetORMDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b balanceSheetORMDo) Joins(fields ...field.RelationField) IBalanceSheetORMDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b balanceSheetORMDo) Preload(fields ...field.RelationField) IBalanceSheetORMDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b balanceSheetORMDo) FirstOrInit() (*accounting_servicev1.BalanceSheetORM, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.BalanceSheetORM), nil
	}
}

func (b balanceSheetORMDo) FirstOrCreate() (*accounting_servicev1.BalanceSheetORM, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.BalanceSheetORM), nil
	}
}

func (b balanceSheetORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.BalanceSheetORM, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b balanceSheetORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b balanceSheetORMDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b balanceSheetORMDo) Delete(models ...*accounting_servicev1.BalanceSheetORM) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *balanceSheetORMDo) withDO(do gen.Dao) *balanceSheetORMDo {
	b.DO = *do.(*gen.DO)
	return b
}
