// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	financial_servicev1 "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/financial_service/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newPlaidAccountTransactionORM(db *gorm.DB, opts ...gen.DOOption) plaidAccountTransactionORM {
	_plaidAccountTransactionORM := plaidAccountTransactionORM{}

	_plaidAccountTransactionORM.plaidAccountTransactionORMDo.UseDB(db, opts...)
	_plaidAccountTransactionORM.plaidAccountTransactionORMDo.UseModel(&financial_servicev1.PlaidAccountTransactionORM{})

	tableName := _plaidAccountTransactionORM.plaidAccountTransactionORMDo.TableName()
	_plaidAccountTransactionORM.ALL = field.NewAsterisk(tableName)
	_plaidAccountTransactionORM.AccountId = field.NewString(tableName, "account_id")
	_plaidAccountTransactionORM.AccountOwner = field.NewString(tableName, "account_owner")
	_plaidAccountTransactionORM.Amount = field.NewFloat64(tableName, "amount")
	_plaidAccountTransactionORM.AuthorizedDate = field.NewTime(tableName, "authorized_date")
	_plaidAccountTransactionORM.AuthorizedDatetime = field.NewTime(tableName, "authorized_datetime")
	_plaidAccountTransactionORM.BankAccountId = field.NewUint64(tableName, "bank_account_id")
	_plaidAccountTransactionORM.Categories = field.NewField(tableName, "categories")
	_plaidAccountTransactionORM.CategoryId = field.NewString(tableName, "category_id")
	_plaidAccountTransactionORM.CheckNumber = field.NewString(tableName, "check_number")
	_plaidAccountTransactionORM.CreditAccountId = field.NewUint64(tableName, "credit_account_id")
	_plaidAccountTransactionORM.CurrentDate = field.NewTime(tableName, "current_date")
	_plaidAccountTransactionORM.CurrentDatetime = field.NewTime(tableName, "current_datetime")
	_plaidAccountTransactionORM.HideTransaction = field.NewBool(tableName, "hide_transaction")
	_plaidAccountTransactionORM.Id = field.NewUint64(tableName, "id")
	_plaidAccountTransactionORM.IsoCurrencyCode = field.NewString(tableName, "iso_currency_code")
	_plaidAccountTransactionORM.LinkId = field.NewUint64(tableName, "link_id")
	_plaidAccountTransactionORM.LocationAddress = field.NewString(tableName, "location_address")
	_plaidAccountTransactionORM.LocationCity = field.NewString(tableName, "location_city")
	_plaidAccountTransactionORM.LocationCountry = field.NewString(tableName, "location_country")
	_plaidAccountTransactionORM.LocationLat = field.NewFloat64(tableName, "location_lat")
	_plaidAccountTransactionORM.LocationLon = field.NewFloat64(tableName, "location_lon")
	_plaidAccountTransactionORM.LocationPostalCode = field.NewString(tableName, "location_postal_code")
	_plaidAccountTransactionORM.LocationRegion = field.NewString(tableName, "location_region")
	_plaidAccountTransactionORM.LocationStoreNumber = field.NewString(tableName, "location_store_number")
	_plaidAccountTransactionORM.MerchantName = field.NewString(tableName, "merchant_name")
	_plaidAccountTransactionORM.NeedsReview = field.NewBool(tableName, "needs_review")
	_plaidAccountTransactionORM.PaymentChannel = field.NewString(tableName, "payment_channel")
	_plaidAccountTransactionORM.PaymentMetaByOrderOf = field.NewString(tableName, "payment_meta_by_order_of")
	_plaidAccountTransactionORM.PaymentMetaPayee = field.NewString(tableName, "payment_meta_payee")
	_plaidAccountTransactionORM.PaymentMetaPayer = field.NewString(tableName, "payment_meta_payer")
	_plaidAccountTransactionORM.PaymentMetaPaymentMethod = field.NewString(tableName, "payment_meta_payment_method")
	_plaidAccountTransactionORM.PaymentMetaPaymentProcessor = field.NewString(tableName, "payment_meta_payment_processor")
	_plaidAccountTransactionORM.PaymentMetaPpdId = field.NewString(tableName, "payment_meta_ppd_id")
	_plaidAccountTransactionORM.PaymentMetaReason = field.NewString(tableName, "payment_meta_reason")
	_plaidAccountTransactionORM.PaymentMetaReferenceNumber = field.NewString(tableName, "payment_meta_reference_number")
	_plaidAccountTransactionORM.Pending = field.NewBool(tableName, "pending")
	_plaidAccountTransactionORM.PendingTransactionId = field.NewString(tableName, "pending_transaction_id")
	_plaidAccountTransactionORM.PersonalFinanceCategoryDetailed = field.NewString(tableName, "personal_finance_category_detailed")
	_plaidAccountTransactionORM.PersonalFinanceCategoryPrimary = field.NewString(tableName, "personal_finance_category_primary")
	_plaidAccountTransactionORM.Tags = field.NewField(tableName, "tags")
	_plaidAccountTransactionORM.Time = field.NewTime(tableName, "time")
	_plaidAccountTransactionORM.TransactionCode = field.NewString(tableName, "transaction_code")
	_plaidAccountTransactionORM.TransactionId = field.NewString(tableName, "transaction_id")
	_plaidAccountTransactionORM.TransactionName = field.NewString(tableName, "transaction_name")
	_plaidAccountTransactionORM.UnofficialCurrencyCode = field.NewString(tableName, "unofficial_currency_code")
	_plaidAccountTransactionORM.UserId = field.NewString(tableName, "user_id")
	_plaidAccountTransactionORM.Notes = plaidAccountTransactionORMHasManyNotes{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Notes", "financial_servicev1.SmartNoteORM"),
	}

	_plaidAccountTransactionORM.Splits = plaidAccountTransactionORMHasManySplits{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Splits", "financial_servicev1.TransactionSplitORM"),
	}

	_plaidAccountTransactionORM.fillFieldMap()

	return _plaidAccountTransactionORM
}

type plaidAccountTransactionORM struct {
	plaidAccountTransactionORMDo

	ALL                             field.Asterisk
	AccountId                       field.String
	AccountOwner                    field.String
	Amount                          field.Float64
	AuthorizedDate                  field.Time
	AuthorizedDatetime              field.Time
	BankAccountId                   field.Uint64
	Categories                      field.Field
	CategoryId                      field.String
	CheckNumber                     field.String
	CreditAccountId                 field.Uint64
	CurrentDate                     field.Time
	CurrentDatetime                 field.Time
	HideTransaction                 field.Bool
	Id                              field.Uint64
	IsoCurrencyCode                 field.String
	LinkId                          field.Uint64
	LocationAddress                 field.String
	LocationCity                    field.String
	LocationCountry                 field.String
	LocationLat                     field.Float64
	LocationLon                     field.Float64
	LocationPostalCode              field.String
	LocationRegion                  field.String
	LocationStoreNumber             field.String
	MerchantName                    field.String
	NeedsReview                     field.Bool
	PaymentChannel                  field.String
	PaymentMetaByOrderOf            field.String
	PaymentMetaPayee                field.String
	PaymentMetaPayer                field.String
	PaymentMetaPaymentMethod        field.String
	PaymentMetaPaymentProcessor     field.String
	PaymentMetaPpdId                field.String
	PaymentMetaReason               field.String
	PaymentMetaReferenceNumber      field.String
	Pending                         field.Bool
	PendingTransactionId            field.String
	PersonalFinanceCategoryDetailed field.String
	PersonalFinanceCategoryPrimary  field.String
	Tags                            field.Field
	Time                            field.Time
	TransactionCode                 field.String
	TransactionId                   field.String
	TransactionName                 field.String
	UnofficialCurrencyCode          field.String
	UserId                          field.String
	Notes                           plaidAccountTransactionORMHasManyNotes

	Splits plaidAccountTransactionORMHasManySplits

	fieldMap map[string]field.Expr
}

func (p plaidAccountTransactionORM) Table(newTableName string) *plaidAccountTransactionORM {
	p.plaidAccountTransactionORMDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p plaidAccountTransactionORM) As(alias string) *plaidAccountTransactionORM {
	p.plaidAccountTransactionORMDo.DO = *(p.plaidAccountTransactionORMDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *plaidAccountTransactionORM) updateTableName(table string) *plaidAccountTransactionORM {
	p.ALL = field.NewAsterisk(table)
	p.AccountId = field.NewString(table, "account_id")
	p.AccountOwner = field.NewString(table, "account_owner")
	p.Amount = field.NewFloat64(table, "amount")
	p.AuthorizedDate = field.NewTime(table, "authorized_date")
	p.AuthorizedDatetime = field.NewTime(table, "authorized_datetime")
	p.BankAccountId = field.NewUint64(table, "bank_account_id")
	p.Categories = field.NewField(table, "categories")
	p.CategoryId = field.NewString(table, "category_id")
	p.CheckNumber = field.NewString(table, "check_number")
	p.CreditAccountId = field.NewUint64(table, "credit_account_id")
	p.CurrentDate = field.NewTime(table, "current_date")
	p.CurrentDatetime = field.NewTime(table, "current_datetime")
	p.HideTransaction = field.NewBool(table, "hide_transaction")
	p.Id = field.NewUint64(table, "id")
	p.IsoCurrencyCode = field.NewString(table, "iso_currency_code")
	p.LinkId = field.NewUint64(table, "link_id")
	p.LocationAddress = field.NewString(table, "location_address")
	p.LocationCity = field.NewString(table, "location_city")
	p.LocationCountry = field.NewString(table, "location_country")
	p.LocationLat = field.NewFloat64(table, "location_lat")
	p.LocationLon = field.NewFloat64(table, "location_lon")
	p.LocationPostalCode = field.NewString(table, "location_postal_code")
	p.LocationRegion = field.NewString(table, "location_region")
	p.LocationStoreNumber = field.NewString(table, "location_store_number")
	p.MerchantName = field.NewString(table, "merchant_name")
	p.NeedsReview = field.NewBool(table, "needs_review")
	p.PaymentChannel = field.NewString(table, "payment_channel")
	p.PaymentMetaByOrderOf = field.NewString(table, "payment_meta_by_order_of")
	p.PaymentMetaPayee = field.NewString(table, "payment_meta_payee")
	p.PaymentMetaPayer = field.NewString(table, "payment_meta_payer")
	p.PaymentMetaPaymentMethod = field.NewString(table, "payment_meta_payment_method")
	p.PaymentMetaPaymentProcessor = field.NewString(table, "payment_meta_payment_processor")
	p.PaymentMetaPpdId = field.NewString(table, "payment_meta_ppd_id")
	p.PaymentMetaReason = field.NewString(table, "payment_meta_reason")
	p.PaymentMetaReferenceNumber = field.NewString(table, "payment_meta_reference_number")
	p.Pending = field.NewBool(table, "pending")
	p.PendingTransactionId = field.NewString(table, "pending_transaction_id")
	p.PersonalFinanceCategoryDetailed = field.NewString(table, "personal_finance_category_detailed")
	p.PersonalFinanceCategoryPrimary = field.NewString(table, "personal_finance_category_primary")
	p.Tags = field.NewField(table, "tags")
	p.Time = field.NewTime(table, "time")
	p.TransactionCode = field.NewString(table, "transaction_code")
	p.TransactionId = field.NewString(table, "transaction_id")
	p.TransactionName = field.NewString(table, "transaction_name")
	p.UnofficialCurrencyCode = field.NewString(table, "unofficial_currency_code")
	p.UserId = field.NewString(table, "user_id")

	p.fillFieldMap()

	return p
}

func (p *plaidAccountTransactionORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *plaidAccountTransactionORM) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 48)
	p.fieldMap["account_id"] = p.AccountId
	p.fieldMap["account_owner"] = p.AccountOwner
	p.fieldMap["amount"] = p.Amount
	p.fieldMap["authorized_date"] = p.AuthorizedDate
	p.fieldMap["authorized_datetime"] = p.AuthorizedDatetime
	p.fieldMap["bank_account_id"] = p.BankAccountId
	p.fieldMap["categories"] = p.Categories
	p.fieldMap["category_id"] = p.CategoryId
	p.fieldMap["check_number"] = p.CheckNumber
	p.fieldMap["credit_account_id"] = p.CreditAccountId
	p.fieldMap["current_date"] = p.CurrentDate
	p.fieldMap["current_datetime"] = p.CurrentDatetime
	p.fieldMap["hide_transaction"] = p.HideTransaction
	p.fieldMap["id"] = p.Id
	p.fieldMap["iso_currency_code"] = p.IsoCurrencyCode
	p.fieldMap["link_id"] = p.LinkId
	p.fieldMap["location_address"] = p.LocationAddress
	p.fieldMap["location_city"] = p.LocationCity
	p.fieldMap["location_country"] = p.LocationCountry
	p.fieldMap["location_lat"] = p.LocationLat
	p.fieldMap["location_lon"] = p.LocationLon
	p.fieldMap["location_postal_code"] = p.LocationPostalCode
	p.fieldMap["location_region"] = p.LocationRegion
	p.fieldMap["location_store_number"] = p.LocationStoreNumber
	p.fieldMap["merchant_name"] = p.MerchantName
	p.fieldMap["needs_review"] = p.NeedsReview
	p.fieldMap["payment_channel"] = p.PaymentChannel
	p.fieldMap["payment_meta_by_order_of"] = p.PaymentMetaByOrderOf
	p.fieldMap["payment_meta_payee"] = p.PaymentMetaPayee
	p.fieldMap["payment_meta_payer"] = p.PaymentMetaPayer
	p.fieldMap["payment_meta_payment_method"] = p.PaymentMetaPaymentMethod
	p.fieldMap["payment_meta_payment_processor"] = p.PaymentMetaPaymentProcessor
	p.fieldMap["payment_meta_ppd_id"] = p.PaymentMetaPpdId
	p.fieldMap["payment_meta_reason"] = p.PaymentMetaReason
	p.fieldMap["payment_meta_reference_number"] = p.PaymentMetaReferenceNumber
	p.fieldMap["pending"] = p.Pending
	p.fieldMap["pending_transaction_id"] = p.PendingTransactionId
	p.fieldMap["personal_finance_category_detailed"] = p.PersonalFinanceCategoryDetailed
	p.fieldMap["personal_finance_category_primary"] = p.PersonalFinanceCategoryPrimary
	p.fieldMap["tags"] = p.Tags
	p.fieldMap["time"] = p.Time
	p.fieldMap["transaction_code"] = p.TransactionCode
	p.fieldMap["transaction_id"] = p.TransactionId
	p.fieldMap["transaction_name"] = p.TransactionName
	p.fieldMap["unofficial_currency_code"] = p.UnofficialCurrencyCode
	p.fieldMap["user_id"] = p.UserId

}

func (p plaidAccountTransactionORM) clone(db *gorm.DB) plaidAccountTransactionORM {
	p.plaidAccountTransactionORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p plaidAccountTransactionORM) replaceDB(db *gorm.DB) plaidAccountTransactionORM {
	p.plaidAccountTransactionORMDo.ReplaceDB(db)
	return p
}

type plaidAccountTransactionORMHasManyNotes struct {
	db *gorm.DB

	field.RelationField
}

func (a plaidAccountTransactionORMHasManyNotes) Where(conds ...field.Expr) *plaidAccountTransactionORMHasManyNotes {
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

func (a plaidAccountTransactionORMHasManyNotes) WithContext(ctx context.Context) *plaidAccountTransactionORMHasManyNotes {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a plaidAccountTransactionORMHasManyNotes) Session(session *gorm.Session) *plaidAccountTransactionORMHasManyNotes {
	a.db = a.db.Session(session)
	return &a
}

func (a plaidAccountTransactionORMHasManyNotes) Model(m *financial_servicev1.PlaidAccountTransactionORM) *plaidAccountTransactionORMHasManyNotesTx {
	return &plaidAccountTransactionORMHasManyNotesTx{a.db.Model(m).Association(a.Name())}
}

type plaidAccountTransactionORMHasManyNotesTx struct{ tx *gorm.Association }

func (a plaidAccountTransactionORMHasManyNotesTx) Find() (result []*financial_servicev1.SmartNoteORM, err error) {
	return result, a.tx.Find(&result)
}

func (a plaidAccountTransactionORMHasManyNotesTx) Append(values ...*financial_servicev1.SmartNoteORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a plaidAccountTransactionORMHasManyNotesTx) Replace(values ...*financial_servicev1.SmartNoteORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a plaidAccountTransactionORMHasManyNotesTx) Delete(values ...*financial_servicev1.SmartNoteORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a plaidAccountTransactionORMHasManyNotesTx) Clear() error {
	return a.tx.Clear()
}

func (a plaidAccountTransactionORMHasManyNotesTx) Count() int64 {
	return a.tx.Count()
}

type plaidAccountTransactionORMHasManySplits struct {
	db *gorm.DB

	field.RelationField
}

func (a plaidAccountTransactionORMHasManySplits) Where(conds ...field.Expr) *plaidAccountTransactionORMHasManySplits {
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

func (a plaidAccountTransactionORMHasManySplits) WithContext(ctx context.Context) *plaidAccountTransactionORMHasManySplits {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a plaidAccountTransactionORMHasManySplits) Session(session *gorm.Session) *plaidAccountTransactionORMHasManySplits {
	a.db = a.db.Session(session)
	return &a
}

func (a plaidAccountTransactionORMHasManySplits) Model(m *financial_servicev1.PlaidAccountTransactionORM) *plaidAccountTransactionORMHasManySplitsTx {
	return &plaidAccountTransactionORMHasManySplitsTx{a.db.Model(m).Association(a.Name())}
}

type plaidAccountTransactionORMHasManySplitsTx struct{ tx *gorm.Association }

func (a plaidAccountTransactionORMHasManySplitsTx) Find() (result []*financial_servicev1.TransactionSplitORM, err error) {
	return result, a.tx.Find(&result)
}

func (a plaidAccountTransactionORMHasManySplitsTx) Append(values ...*financial_servicev1.TransactionSplitORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a plaidAccountTransactionORMHasManySplitsTx) Replace(values ...*financial_servicev1.TransactionSplitORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a plaidAccountTransactionORMHasManySplitsTx) Delete(values ...*financial_servicev1.TransactionSplitORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a plaidAccountTransactionORMHasManySplitsTx) Clear() error {
	return a.tx.Clear()
}

func (a plaidAccountTransactionORMHasManySplitsTx) Count() int64 {
	return a.tx.Count()
}

type plaidAccountTransactionORMDo struct{ gen.DO }

type IPlaidAccountTransactionORMDo interface {
	gen.SubQuery
	Debug() IPlaidAccountTransactionORMDo
	WithContext(ctx context.Context) IPlaidAccountTransactionORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPlaidAccountTransactionORMDo
	WriteDB() IPlaidAccountTransactionORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPlaidAccountTransactionORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPlaidAccountTransactionORMDo
	Not(conds ...gen.Condition) IPlaidAccountTransactionORMDo
	Or(conds ...gen.Condition) IPlaidAccountTransactionORMDo
	Select(conds ...field.Expr) IPlaidAccountTransactionORMDo
	Where(conds ...gen.Condition) IPlaidAccountTransactionORMDo
	Order(conds ...field.Expr) IPlaidAccountTransactionORMDo
	Distinct(cols ...field.Expr) IPlaidAccountTransactionORMDo
	Omit(cols ...field.Expr) IPlaidAccountTransactionORMDo
	Join(table schema.Tabler, on ...field.Expr) IPlaidAccountTransactionORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPlaidAccountTransactionORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPlaidAccountTransactionORMDo
	Group(cols ...field.Expr) IPlaidAccountTransactionORMDo
	Having(conds ...gen.Condition) IPlaidAccountTransactionORMDo
	Limit(limit int) IPlaidAccountTransactionORMDo
	Offset(offset int) IPlaidAccountTransactionORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPlaidAccountTransactionORMDo
	Unscoped() IPlaidAccountTransactionORMDo
	Create(values ...*financial_servicev1.PlaidAccountTransactionORM) error
	CreateInBatches(values []*financial_servicev1.PlaidAccountTransactionORM, batchSize int) error
	Save(values ...*financial_servicev1.PlaidAccountTransactionORM) error
	First() (*financial_servicev1.PlaidAccountTransactionORM, error)
	Take() (*financial_servicev1.PlaidAccountTransactionORM, error)
	Last() (*financial_servicev1.PlaidAccountTransactionORM, error)
	Find() ([]*financial_servicev1.PlaidAccountTransactionORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_servicev1.PlaidAccountTransactionORM, err error)
	FindInBatches(result *[]*financial_servicev1.PlaidAccountTransactionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_servicev1.PlaidAccountTransactionORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPlaidAccountTransactionORMDo
	Assign(attrs ...field.AssignExpr) IPlaidAccountTransactionORMDo
	Joins(fields ...field.RelationField) IPlaidAccountTransactionORMDo
	Preload(fields ...field.RelationField) IPlaidAccountTransactionORMDo
	FirstOrInit() (*financial_servicev1.PlaidAccountTransactionORM, error)
	FirstOrCreate() (*financial_servicev1.PlaidAccountTransactionORM, error)
	FindByPage(offset int, limit int) (result []*financial_servicev1.PlaidAccountTransactionORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPlaidAccountTransactionORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result financial_servicev1.PlaidAccountTransactionORM, err error)
	GetRecordByIDs(ids []int) (result []financial_servicev1.PlaidAccountTransactionORM, err error)
	CreateRecord(item financial_servicev1.PlaidAccountTransactionORM) (err error)
	UpdateRecordByID(id int, item financial_servicev1.PlaidAccountTransactionORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []financial_servicev1.PlaidAccountTransactionORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result financial_servicev1.PlaidAccountTransactionORM, err error)
	GetByIDs(ids []uint64) (result []financial_servicev1.PlaidAccountTransactionORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (p plaidAccountTransactionORMDo) GetRecordByID(id int) (result financial_servicev1.PlaidAccountTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM plaid_account_transactions ")
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
func (p plaidAccountTransactionORMDo) GetRecordByIDs(ids []int) (result []financial_servicev1.PlaidAccountTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM plaid_account_transactions ")
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
func (p plaidAccountTransactionORMDo) CreateRecord(item financial_servicev1.PlaidAccountTransactionORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO plaid_account_transactions (columns) VALUES (values) ")

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
func (p plaidAccountTransactionORMDo) UpdateRecordByID(id int, item financial_servicev1.PlaidAccountTransactionORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE plaid_account_transactions SET columns=values ")
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
func (p plaidAccountTransactionORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM plaid_account_transactions ")
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
func (p plaidAccountTransactionORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []financial_servicev1.PlaidAccountTransactionORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM plaid_account_transactions ORDER BY " + p.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (p plaidAccountTransactionORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM plaid_account_transactions ")

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
func (p plaidAccountTransactionORMDo) GetByID(id uint64) (result financial_servicev1.PlaidAccountTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM plaid_account_transactions ")
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
func (p plaidAccountTransactionORMDo) GetByIDs(ids []uint64) (result []financial_servicev1.PlaidAccountTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM plaid_account_transactions ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p plaidAccountTransactionORMDo) Debug() IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Debug())
}

func (p plaidAccountTransactionORMDo) WithContext(ctx context.Context) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p plaidAccountTransactionORMDo) ReadDB() IPlaidAccountTransactionORMDo {
	return p.Clauses(dbresolver.Read)
}

func (p plaidAccountTransactionORMDo) WriteDB() IPlaidAccountTransactionORMDo {
	return p.Clauses(dbresolver.Write)
}

func (p plaidAccountTransactionORMDo) Session(config *gorm.Session) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Session(config))
}

func (p plaidAccountTransactionORMDo) Clauses(conds ...clause.Expression) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p plaidAccountTransactionORMDo) Returning(value interface{}, columns ...string) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p plaidAccountTransactionORMDo) Not(conds ...gen.Condition) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p plaidAccountTransactionORMDo) Or(conds ...gen.Condition) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p plaidAccountTransactionORMDo) Select(conds ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p plaidAccountTransactionORMDo) Where(conds ...gen.Condition) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p plaidAccountTransactionORMDo) Order(conds ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p plaidAccountTransactionORMDo) Distinct(cols ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p plaidAccountTransactionORMDo) Omit(cols ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p plaidAccountTransactionORMDo) Join(table schema.Tabler, on ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p plaidAccountTransactionORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p plaidAccountTransactionORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p plaidAccountTransactionORMDo) Group(cols ...field.Expr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p plaidAccountTransactionORMDo) Having(conds ...gen.Condition) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p plaidAccountTransactionORMDo) Limit(limit int) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p plaidAccountTransactionORMDo) Offset(offset int) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p plaidAccountTransactionORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p plaidAccountTransactionORMDo) Unscoped() IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Unscoped())
}

func (p plaidAccountTransactionORMDo) Create(values ...*financial_servicev1.PlaidAccountTransactionORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p plaidAccountTransactionORMDo) CreateInBatches(values []*financial_servicev1.PlaidAccountTransactionORM, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p plaidAccountTransactionORMDo) Save(values ...*financial_servicev1.PlaidAccountTransactionORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p plaidAccountTransactionORMDo) First() (*financial_servicev1.PlaidAccountTransactionORM, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.PlaidAccountTransactionORM), nil
	}
}

func (p plaidAccountTransactionORMDo) Take() (*financial_servicev1.PlaidAccountTransactionORM, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.PlaidAccountTransactionORM), nil
	}
}

func (p plaidAccountTransactionORMDo) Last() (*financial_servicev1.PlaidAccountTransactionORM, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.PlaidAccountTransactionORM), nil
	}
}

func (p plaidAccountTransactionORMDo) Find() ([]*financial_servicev1.PlaidAccountTransactionORM, error) {
	result, err := p.DO.Find()
	return result.([]*financial_servicev1.PlaidAccountTransactionORM), err
}

func (p plaidAccountTransactionORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_servicev1.PlaidAccountTransactionORM, err error) {
	buf := make([]*financial_servicev1.PlaidAccountTransactionORM, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p plaidAccountTransactionORMDo) FindInBatches(result *[]*financial_servicev1.PlaidAccountTransactionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p plaidAccountTransactionORMDo) Attrs(attrs ...field.AssignExpr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p plaidAccountTransactionORMDo) Assign(attrs ...field.AssignExpr) IPlaidAccountTransactionORMDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p plaidAccountTransactionORMDo) Joins(fields ...field.RelationField) IPlaidAccountTransactionORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p plaidAccountTransactionORMDo) Preload(fields ...field.RelationField) IPlaidAccountTransactionORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p plaidAccountTransactionORMDo) FirstOrInit() (*financial_servicev1.PlaidAccountTransactionORM, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.PlaidAccountTransactionORM), nil
	}
}

func (p plaidAccountTransactionORMDo) FirstOrCreate() (*financial_servicev1.PlaidAccountTransactionORM, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.PlaidAccountTransactionORM), nil
	}
}

func (p plaidAccountTransactionORMDo) FindByPage(offset int, limit int) (result []*financial_servicev1.PlaidAccountTransactionORM, count int64, err error) {
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

func (p plaidAccountTransactionORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p plaidAccountTransactionORMDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p plaidAccountTransactionORMDo) Delete(models ...*financial_servicev1.PlaidAccountTransactionORM) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *plaidAccountTransactionORMDo) withDO(do gen.Dao) *plaidAccountTransactionORMDo {
	p.DO = *do.(*gen.DO)
	return p
}
