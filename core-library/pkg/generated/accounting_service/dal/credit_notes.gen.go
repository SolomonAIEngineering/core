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

func newCreditNoteORM(db *gorm.DB, opts ...gen.DOOption) creditNoteORM {
	_creditNoteORM := creditNoteORM{}

	_creditNoteORM.creditNoteORMDo.UseDB(db, opts...)
	_creditNoteORM.creditNoteORMDo.UseModel(&accounting_servicev1.CreditNoteORM{})

	tableName := _creditNoteORM.creditNoteORMDo.TableName()
	_creditNoteORM.ALL = field.NewAsterisk(tableName)
	_creditNoteORM.AccountingPeriod = field.NewString(tableName, "accounting_period")
	_creditNoteORM.Company = field.NewString(tableName, "company")
	_creditNoteORM.Contact = field.NewString(tableName, "contact")
	_creditNoteORM.CreatedAt = field.NewTime(tableName, "created_at")
	_creditNoteORM.Currency = field.NewString(tableName, "currency")
	_creditNoteORM.ExchangeRate = field.NewString(tableName, "exchange_rate")
	_creditNoteORM.Id = field.NewUint64(tableName, "id")
	_creditNoteORM.LinkedAccountingAccountId = field.NewUint64(tableName, "linked_accounting_account_id")
	_creditNoteORM.MergeRecordId = field.NewString(tableName, "merge_record_id")
	_creditNoteORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_creditNoteORM.Number = field.NewString(tableName, "number")
	_creditNoteORM.PaymentIds = field.NewField(tableName, "payment_ids")
	_creditNoteORM.RemainingCredit = field.NewFloat64(tableName, "remaining_credit")
	_creditNoteORM.RemoteCreatedAt = field.NewTime(tableName, "remote_created_at")
	_creditNoteORM.RemoteId = field.NewString(tableName, "remote_id")
	_creditNoteORM.RemoteUpdatedAt = field.NewTime(tableName, "remote_updated_at")
	_creditNoteORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_creditNoteORM.Status = field.NewString(tableName, "status")
	_creditNoteORM.TotalAmount = field.NewFloat64(tableName, "total_amount")
	_creditNoteORM.TrackingCategories = field.NewField(tableName, "tracking_categories")
	_creditNoteORM.TransactionDate = field.NewTime(tableName, "transaction_date")
	_creditNoteORM.LineItems = creditNoteORMHasManyLineItems{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("LineItems", "accounting_servicev1.CreditNoteLineItemORM"),
	}

	_creditNoteORM.fillFieldMap()

	return _creditNoteORM
}

type creditNoteORM struct {
	creditNoteORMDo

	ALL                       field.Asterisk
	AccountingPeriod          field.String
	Company                   field.String
	Contact                   field.String
	CreatedAt                 field.Time
	Currency                  field.String
	ExchangeRate              field.String
	Id                        field.Uint64
	LinkedAccountingAccountId field.Uint64
	MergeRecordId             field.String
	ModifiedAt                field.Time
	Number                    field.String
	PaymentIds                field.Field
	RemainingCredit           field.Float64
	RemoteCreatedAt           field.Time
	RemoteId                  field.String
	RemoteUpdatedAt           field.Time
	RemoteWasDeleted          field.Bool
	Status                    field.String
	TotalAmount               field.Float64
	TrackingCategories        field.Field
	TransactionDate           field.Time
	LineItems                 creditNoteORMHasManyLineItems

	fieldMap map[string]field.Expr
}

func (c creditNoteORM) Table(newTableName string) *creditNoteORM {
	c.creditNoteORMDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c creditNoteORM) As(alias string) *creditNoteORM {
	c.creditNoteORMDo.DO = *(c.creditNoteORMDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *creditNoteORM) updateTableName(table string) *creditNoteORM {
	c.ALL = field.NewAsterisk(table)
	c.AccountingPeriod = field.NewString(table, "accounting_period")
	c.Company = field.NewString(table, "company")
	c.Contact = field.NewString(table, "contact")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.Currency = field.NewString(table, "currency")
	c.ExchangeRate = field.NewString(table, "exchange_rate")
	c.Id = field.NewUint64(table, "id")
	c.LinkedAccountingAccountId = field.NewUint64(table, "linked_accounting_account_id")
	c.MergeRecordId = field.NewString(table, "merge_record_id")
	c.ModifiedAt = field.NewTime(table, "modified_at")
	c.Number = field.NewString(table, "number")
	c.PaymentIds = field.NewField(table, "payment_ids")
	c.RemainingCredit = field.NewFloat64(table, "remaining_credit")
	c.RemoteCreatedAt = field.NewTime(table, "remote_created_at")
	c.RemoteId = field.NewString(table, "remote_id")
	c.RemoteUpdatedAt = field.NewTime(table, "remote_updated_at")
	c.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	c.Status = field.NewString(table, "status")
	c.TotalAmount = field.NewFloat64(table, "total_amount")
	c.TrackingCategories = field.NewField(table, "tracking_categories")
	c.TransactionDate = field.NewTime(table, "transaction_date")

	c.fillFieldMap()

	return c
}

func (c *creditNoteORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *creditNoteORM) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 22)
	c.fieldMap["accounting_period"] = c.AccountingPeriod
	c.fieldMap["company"] = c.Company
	c.fieldMap["contact"] = c.Contact
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["currency"] = c.Currency
	c.fieldMap["exchange_rate"] = c.ExchangeRate
	c.fieldMap["id"] = c.Id
	c.fieldMap["linked_accounting_account_id"] = c.LinkedAccountingAccountId
	c.fieldMap["merge_record_id"] = c.MergeRecordId
	c.fieldMap["modified_at"] = c.ModifiedAt
	c.fieldMap["number"] = c.Number
	c.fieldMap["payment_ids"] = c.PaymentIds
	c.fieldMap["remaining_credit"] = c.RemainingCredit
	c.fieldMap["remote_created_at"] = c.RemoteCreatedAt
	c.fieldMap["remote_id"] = c.RemoteId
	c.fieldMap["remote_updated_at"] = c.RemoteUpdatedAt
	c.fieldMap["remote_was_deleted"] = c.RemoteWasDeleted
	c.fieldMap["status"] = c.Status
	c.fieldMap["total_amount"] = c.TotalAmount
	c.fieldMap["tracking_categories"] = c.TrackingCategories
	c.fieldMap["transaction_date"] = c.TransactionDate

}

func (c creditNoteORM) clone(db *gorm.DB) creditNoteORM {
	c.creditNoteORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c creditNoteORM) replaceDB(db *gorm.DB) creditNoteORM {
	c.creditNoteORMDo.ReplaceDB(db)
	return c
}

type creditNoteORMHasManyLineItems struct {
	db *gorm.DB

	field.RelationField
}

func (a creditNoteORMHasManyLineItems) Where(conds ...field.Expr) *creditNoteORMHasManyLineItems {
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

func (a creditNoteORMHasManyLineItems) WithContext(ctx context.Context) *creditNoteORMHasManyLineItems {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a creditNoteORMHasManyLineItems) Session(session *gorm.Session) *creditNoteORMHasManyLineItems {
	a.db = a.db.Session(session)
	return &a
}

func (a creditNoteORMHasManyLineItems) Model(m *accounting_servicev1.CreditNoteORM) *creditNoteORMHasManyLineItemsTx {
	return &creditNoteORMHasManyLineItemsTx{a.db.Model(m).Association(a.Name())}
}

type creditNoteORMHasManyLineItemsTx struct{ tx *gorm.Association }

func (a creditNoteORMHasManyLineItemsTx) Find() (result []*accounting_servicev1.CreditNoteLineItemORM, err error) {
	return result, a.tx.Find(&result)
}

func (a creditNoteORMHasManyLineItemsTx) Append(values ...*accounting_servicev1.CreditNoteLineItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a creditNoteORMHasManyLineItemsTx) Replace(values ...*accounting_servicev1.CreditNoteLineItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a creditNoteORMHasManyLineItemsTx) Delete(values ...*accounting_servicev1.CreditNoteLineItemORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a creditNoteORMHasManyLineItemsTx) Clear() error {
	return a.tx.Clear()
}

func (a creditNoteORMHasManyLineItemsTx) Count() int64 {
	return a.tx.Count()
}

type creditNoteORMDo struct{ gen.DO }

type ICreditNoteORMDo interface {
	gen.SubQuery
	Debug() ICreditNoteORMDo
	WithContext(ctx context.Context) ICreditNoteORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICreditNoteORMDo
	WriteDB() ICreditNoteORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICreditNoteORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICreditNoteORMDo
	Not(conds ...gen.Condition) ICreditNoteORMDo
	Or(conds ...gen.Condition) ICreditNoteORMDo
	Select(conds ...field.Expr) ICreditNoteORMDo
	Where(conds ...gen.Condition) ICreditNoteORMDo
	Order(conds ...field.Expr) ICreditNoteORMDo
	Distinct(cols ...field.Expr) ICreditNoteORMDo
	Omit(cols ...field.Expr) ICreditNoteORMDo
	Join(table schema.Tabler, on ...field.Expr) ICreditNoteORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICreditNoteORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICreditNoteORMDo
	Group(cols ...field.Expr) ICreditNoteORMDo
	Having(conds ...gen.Condition) ICreditNoteORMDo
	Limit(limit int) ICreditNoteORMDo
	Offset(offset int) ICreditNoteORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICreditNoteORMDo
	Unscoped() ICreditNoteORMDo
	Create(values ...*accounting_servicev1.CreditNoteORM) error
	CreateInBatches(values []*accounting_servicev1.CreditNoteORM, batchSize int) error
	Save(values ...*accounting_servicev1.CreditNoteORM) error
	First() (*accounting_servicev1.CreditNoteORM, error)
	Take() (*accounting_servicev1.CreditNoteORM, error)
	Last() (*accounting_servicev1.CreditNoteORM, error)
	Find() ([]*accounting_servicev1.CreditNoteORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.CreditNoteORM, err error)
	FindInBatches(result *[]*accounting_servicev1.CreditNoteORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.CreditNoteORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICreditNoteORMDo
	Assign(attrs ...field.AssignExpr) ICreditNoteORMDo
	Joins(fields ...field.RelationField) ICreditNoteORMDo
	Preload(fields ...field.RelationField) ICreditNoteORMDo
	FirstOrInit() (*accounting_servicev1.CreditNoteORM, error)
	FirstOrCreate() (*accounting_servicev1.CreditNoteORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.CreditNoteORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICreditNoteORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.CreditNoteORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.CreditNoteORM, err error)
	CreateRecord(item accounting_servicev1.CreditNoteORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.CreditNoteORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.CreditNoteORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.CreditNoteORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.CreditNoteORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c creditNoteORMDo) GetRecordByID(id int) (result accounting_servicev1.CreditNoteORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_notes ")
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
func (c creditNoteORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.CreditNoteORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_notes ")
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
func (c creditNoteORMDo) CreateRecord(item accounting_servicev1.CreditNoteORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO credit_notes (columns) VALUES (values) ")

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
func (c creditNoteORMDo) UpdateRecordByID(id int, item accounting_servicev1.CreditNoteORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE credit_notes SET columns=values ")
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
func (c creditNoteORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM credit_notes ")
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
func (c creditNoteORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.CreditNoteORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_notes ORDER BY " + c.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (c creditNoteORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM credit_notes ")

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
func (c creditNoteORMDo) GetByID(id uint64) (result accounting_servicev1.CreditNoteORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_notes ")
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
func (c creditNoteORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.CreditNoteORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_notes ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c creditNoteORMDo) Debug() ICreditNoteORMDo {
	return c.withDO(c.DO.Debug())
}

func (c creditNoteORMDo) WithContext(ctx context.Context) ICreditNoteORMDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c creditNoteORMDo) ReadDB() ICreditNoteORMDo {
	return c.Clauses(dbresolver.Read)
}

func (c creditNoteORMDo) WriteDB() ICreditNoteORMDo {
	return c.Clauses(dbresolver.Write)
}

func (c creditNoteORMDo) Session(config *gorm.Session) ICreditNoteORMDo {
	return c.withDO(c.DO.Session(config))
}

func (c creditNoteORMDo) Clauses(conds ...clause.Expression) ICreditNoteORMDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c creditNoteORMDo) Returning(value interface{}, columns ...string) ICreditNoteORMDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c creditNoteORMDo) Not(conds ...gen.Condition) ICreditNoteORMDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c creditNoteORMDo) Or(conds ...gen.Condition) ICreditNoteORMDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c creditNoteORMDo) Select(conds ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c creditNoteORMDo) Where(conds ...gen.Condition) ICreditNoteORMDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c creditNoteORMDo) Order(conds ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c creditNoteORMDo) Distinct(cols ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c creditNoteORMDo) Omit(cols ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c creditNoteORMDo) Join(table schema.Tabler, on ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c creditNoteORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c creditNoteORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c creditNoteORMDo) Group(cols ...field.Expr) ICreditNoteORMDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c creditNoteORMDo) Having(conds ...gen.Condition) ICreditNoteORMDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c creditNoteORMDo) Limit(limit int) ICreditNoteORMDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c creditNoteORMDo) Offset(offset int) ICreditNoteORMDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c creditNoteORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICreditNoteORMDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c creditNoteORMDo) Unscoped() ICreditNoteORMDo {
	return c.withDO(c.DO.Unscoped())
}

func (c creditNoteORMDo) Create(values ...*accounting_servicev1.CreditNoteORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c creditNoteORMDo) CreateInBatches(values []*accounting_servicev1.CreditNoteORM, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c creditNoteORMDo) Save(values ...*accounting_servicev1.CreditNoteORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c creditNoteORMDo) First() (*accounting_servicev1.CreditNoteORM, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CreditNoteORM), nil
	}
}

func (c creditNoteORMDo) Take() (*accounting_servicev1.CreditNoteORM, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CreditNoteORM), nil
	}
}

func (c creditNoteORMDo) Last() (*accounting_servicev1.CreditNoteORM, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CreditNoteORM), nil
	}
}

func (c creditNoteORMDo) Find() ([]*accounting_servicev1.CreditNoteORM, error) {
	result, err := c.DO.Find()
	return result.([]*accounting_servicev1.CreditNoteORM), err
}

func (c creditNoteORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.CreditNoteORM, err error) {
	buf := make([]*accounting_servicev1.CreditNoteORM, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c creditNoteORMDo) FindInBatches(result *[]*accounting_servicev1.CreditNoteORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c creditNoteORMDo) Attrs(attrs ...field.AssignExpr) ICreditNoteORMDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c creditNoteORMDo) Assign(attrs ...field.AssignExpr) ICreditNoteORMDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c creditNoteORMDo) Joins(fields ...field.RelationField) ICreditNoteORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c creditNoteORMDo) Preload(fields ...field.RelationField) ICreditNoteORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c creditNoteORMDo) FirstOrInit() (*accounting_servicev1.CreditNoteORM, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CreditNoteORM), nil
	}
}

func (c creditNoteORMDo) FirstOrCreate() (*accounting_servicev1.CreditNoteORM, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CreditNoteORM), nil
	}
}

func (c creditNoteORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.CreditNoteORM, count int64, err error) {
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

func (c creditNoteORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c creditNoteORMDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c creditNoteORMDo) Delete(models ...*accounting_servicev1.CreditNoteORM) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *creditNoteORMDo) withDO(do gen.Dao) *creditNoteORMDo {
	c.DO = *do.(*gen.DO)
	return c
}
