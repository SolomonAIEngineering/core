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

func newPaymentORM(db *gorm.DB, opts ...gen.DOOption) paymentORM {
	_paymentORM := paymentORM{}

	_paymentORM.paymentORMDo.UseDB(db, opts...)
	_paymentORM.paymentORMDo.UseModel(&accounting_servicev1.PaymentORM{})

	tableName := _paymentORM.paymentORMDo.TableName()
	_paymentORM.ALL = field.NewAsterisk(tableName)
	_paymentORM.Account = field.NewString(tableName, "account")
	_paymentORM.AccountingPeriod = field.NewString(tableName, "accounting_period")
	_paymentORM.Company = field.NewString(tableName, "company")
	_paymentORM.Contact = field.NewString(tableName, "contact")
	_paymentORM.CreatedAt = field.NewTime(tableName, "created_at")
	_paymentORM.Currency = field.NewString(tableName, "currency")
	_paymentORM.ExchangeRate = field.NewString(tableName, "exchange_rate")
	_paymentORM.Id = field.NewUint64(tableName, "id")
	_paymentORM.LinkedAccountingAccountId = field.NewUint64(tableName, "linked_accounting_account_id")
	_paymentORM.MergeRecordId = field.NewString(tableName, "merge_record_id")
	_paymentORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_paymentORM.RemoteId = field.NewString(tableName, "remote_id")
	_paymentORM.RemoteUpdatedAt = field.NewTime(tableName, "remote_updated_at")
	_paymentORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_paymentORM.TotalAmount = field.NewFloat32(tableName, "total_amount")
	_paymentORM.TrackingCategories = field.NewField(tableName, "tracking_categories")
	_paymentORM.TransactionDate = field.NewTime(tableName, "transaction_date")

	_paymentORM.fillFieldMap()

	return _paymentORM
}

type paymentORM struct {
	paymentORMDo

	ALL                       field.Asterisk
	Account                   field.String
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
	RemoteId                  field.String
	RemoteUpdatedAt           field.Time
	RemoteWasDeleted          field.Bool
	TotalAmount               field.Float32
	TrackingCategories        field.Field
	TransactionDate           field.Time

	fieldMap map[string]field.Expr
}

func (p paymentORM) Table(newTableName string) *paymentORM {
	p.paymentORMDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p paymentORM) As(alias string) *paymentORM {
	p.paymentORMDo.DO = *(p.paymentORMDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *paymentORM) updateTableName(table string) *paymentORM {
	p.ALL = field.NewAsterisk(table)
	p.Account = field.NewString(table, "account")
	p.AccountingPeriod = field.NewString(table, "accounting_period")
	p.Company = field.NewString(table, "company")
	p.Contact = field.NewString(table, "contact")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.Currency = field.NewString(table, "currency")
	p.ExchangeRate = field.NewString(table, "exchange_rate")
	p.Id = field.NewUint64(table, "id")
	p.LinkedAccountingAccountId = field.NewUint64(table, "linked_accounting_account_id")
	p.MergeRecordId = field.NewString(table, "merge_record_id")
	p.ModifiedAt = field.NewTime(table, "modified_at")
	p.RemoteId = field.NewString(table, "remote_id")
	p.RemoteUpdatedAt = field.NewTime(table, "remote_updated_at")
	p.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	p.TotalAmount = field.NewFloat32(table, "total_amount")
	p.TrackingCategories = field.NewField(table, "tracking_categories")
	p.TransactionDate = field.NewTime(table, "transaction_date")

	p.fillFieldMap()

	return p
}

func (p *paymentORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *paymentORM) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 17)
	p.fieldMap["account"] = p.Account
	p.fieldMap["accounting_period"] = p.AccountingPeriod
	p.fieldMap["company"] = p.Company
	p.fieldMap["contact"] = p.Contact
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["currency"] = p.Currency
	p.fieldMap["exchange_rate"] = p.ExchangeRate
	p.fieldMap["id"] = p.Id
	p.fieldMap["linked_accounting_account_id"] = p.LinkedAccountingAccountId
	p.fieldMap["merge_record_id"] = p.MergeRecordId
	p.fieldMap["modified_at"] = p.ModifiedAt
	p.fieldMap["remote_id"] = p.RemoteId
	p.fieldMap["remote_updated_at"] = p.RemoteUpdatedAt
	p.fieldMap["remote_was_deleted"] = p.RemoteWasDeleted
	p.fieldMap["total_amount"] = p.TotalAmount
	p.fieldMap["tracking_categories"] = p.TrackingCategories
	p.fieldMap["transaction_date"] = p.TransactionDate
}

func (p paymentORM) clone(db *gorm.DB) paymentORM {
	p.paymentORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p paymentORM) replaceDB(db *gorm.DB) paymentORM {
	p.paymentORMDo.ReplaceDB(db)
	return p
}

type paymentORMDo struct{ gen.DO }

type IPaymentORMDo interface {
	gen.SubQuery
	Debug() IPaymentORMDo
	WithContext(ctx context.Context) IPaymentORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPaymentORMDo
	WriteDB() IPaymentORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPaymentORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPaymentORMDo
	Not(conds ...gen.Condition) IPaymentORMDo
	Or(conds ...gen.Condition) IPaymentORMDo
	Select(conds ...field.Expr) IPaymentORMDo
	Where(conds ...gen.Condition) IPaymentORMDo
	Order(conds ...field.Expr) IPaymentORMDo
	Distinct(cols ...field.Expr) IPaymentORMDo
	Omit(cols ...field.Expr) IPaymentORMDo
	Join(table schema.Tabler, on ...field.Expr) IPaymentORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPaymentORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPaymentORMDo
	Group(cols ...field.Expr) IPaymentORMDo
	Having(conds ...gen.Condition) IPaymentORMDo
	Limit(limit int) IPaymentORMDo
	Offset(offset int) IPaymentORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPaymentORMDo
	Unscoped() IPaymentORMDo
	Create(values ...*accounting_servicev1.PaymentORM) error
	CreateInBatches(values []*accounting_servicev1.PaymentORM, batchSize int) error
	Save(values ...*accounting_servicev1.PaymentORM) error
	First() (*accounting_servicev1.PaymentORM, error)
	Take() (*accounting_servicev1.PaymentORM, error)
	Last() (*accounting_servicev1.PaymentORM, error)
	Find() ([]*accounting_servicev1.PaymentORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.PaymentORM, err error)
	FindInBatches(result *[]*accounting_servicev1.PaymentORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.PaymentORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPaymentORMDo
	Assign(attrs ...field.AssignExpr) IPaymentORMDo
	Joins(fields ...field.RelationField) IPaymentORMDo
	Preload(fields ...field.RelationField) IPaymentORMDo
	FirstOrInit() (*accounting_servicev1.PaymentORM, error)
	FirstOrCreate() (*accounting_servicev1.PaymentORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.PaymentORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPaymentORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.PaymentORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.PaymentORM, err error)
	CreateRecord(item accounting_servicev1.PaymentORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.PaymentORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.PaymentORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.PaymentORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.PaymentORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (p paymentORMDo) GetRecordByID(id int) (result accounting_servicev1.PaymentORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM payments ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (p paymentORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.PaymentORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM payments ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (p paymentORMDo) CreateRecord(item accounting_servicev1.PaymentORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO payments (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (p paymentORMDo) UpdateRecordByID(id int, item accounting_servicev1.PaymentORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE payments SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (p paymentORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM payments ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (p paymentORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.PaymentORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM payments ORDER BY " + p.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (p paymentORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM payments ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (p paymentORMDo) GetByID(id uint64) (result accounting_servicev1.PaymentORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM payments ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (p paymentORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.PaymentORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM payments ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p paymentORMDo) Debug() IPaymentORMDo {
	return p.withDO(p.DO.Debug())
}

func (p paymentORMDo) WithContext(ctx context.Context) IPaymentORMDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paymentORMDo) ReadDB() IPaymentORMDo {
	return p.Clauses(dbresolver.Read)
}

func (p paymentORMDo) WriteDB() IPaymentORMDo {
	return p.Clauses(dbresolver.Write)
}

func (p paymentORMDo) Session(config *gorm.Session) IPaymentORMDo {
	return p.withDO(p.DO.Session(config))
}

func (p paymentORMDo) Clauses(conds ...clause.Expression) IPaymentORMDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paymentORMDo) Returning(value interface{}, columns ...string) IPaymentORMDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paymentORMDo) Not(conds ...gen.Condition) IPaymentORMDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paymentORMDo) Or(conds ...gen.Condition) IPaymentORMDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paymentORMDo) Select(conds ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paymentORMDo) Where(conds ...gen.Condition) IPaymentORMDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paymentORMDo) Order(conds ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paymentORMDo) Distinct(cols ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paymentORMDo) Omit(cols ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paymentORMDo) Join(table schema.Tabler, on ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paymentORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paymentORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paymentORMDo) Group(cols ...field.Expr) IPaymentORMDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paymentORMDo) Having(conds ...gen.Condition) IPaymentORMDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paymentORMDo) Limit(limit int) IPaymentORMDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paymentORMDo) Offset(offset int) IPaymentORMDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paymentORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPaymentORMDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paymentORMDo) Unscoped() IPaymentORMDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paymentORMDo) Create(values ...*accounting_servicev1.PaymentORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paymentORMDo) CreateInBatches(values []*accounting_servicev1.PaymentORM, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paymentORMDo) Save(values ...*accounting_servicev1.PaymentORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paymentORMDo) First() (*accounting_servicev1.PaymentORM, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.PaymentORM), nil
	}
}

func (p paymentORMDo) Take() (*accounting_servicev1.PaymentORM, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.PaymentORM), nil
	}
}

func (p paymentORMDo) Last() (*accounting_servicev1.PaymentORM, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.PaymentORM), nil
	}
}

func (p paymentORMDo) Find() ([]*accounting_servicev1.PaymentORM, error) {
	result, err := p.DO.Find()
	return result.([]*accounting_servicev1.PaymentORM), err
}

func (p paymentORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.PaymentORM, err error) {
	buf := make([]*accounting_servicev1.PaymentORM, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paymentORMDo) FindInBatches(result *[]*accounting_servicev1.PaymentORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paymentORMDo) Attrs(attrs ...field.AssignExpr) IPaymentORMDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paymentORMDo) Assign(attrs ...field.AssignExpr) IPaymentORMDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paymentORMDo) Joins(fields ...field.RelationField) IPaymentORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paymentORMDo) Preload(fields ...field.RelationField) IPaymentORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paymentORMDo) FirstOrInit() (*accounting_servicev1.PaymentORM, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.PaymentORM), nil
	}
}

func (p paymentORMDo) FirstOrCreate() (*accounting_servicev1.PaymentORM, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.PaymentORM), nil
	}
}

func (p paymentORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.PaymentORM, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p paymentORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paymentORMDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paymentORMDo) Delete(models ...*accounting_servicev1.PaymentORM) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paymentORMDo) withDO(do gen.Dao) *paymentORMDo {
	p.DO = *do.(*gen.DO)
	return p
}
