// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	financial_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/financial_service/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newCreditAccountORM(db *gorm.DB, opts ...gen.DOOption) creditAccountORM {
	_creditAccountORM := creditAccountORM{}

	_creditAccountORM.creditAccountORMDo.UseDB(db, opts...)
	_creditAccountORM.creditAccountORMDo.UseModel(&financial_servicev1.CreditAccountORM{})

	tableName := _creditAccountORM.creditAccountORMDo.TableName()
	_creditAccountORM.ALL = field.NewAsterisk(tableName)
	_creditAccountORM.Balance = field.NewFloat32(tableName, "balance")
	_creditAccountORM.BalanceLimit = field.NewUint64(tableName, "balance_limit")
	_creditAccountORM.CurrentFunds = field.NewFloat64(tableName, "current_funds")
	_creditAccountORM.Id = field.NewUint64(tableName, "id")
	_creditAccountORM.IsOverdue = field.NewBool(tableName, "is_overdue")
	_creditAccountORM.LastPaymentAmount = field.NewFloat64(tableName, "last_payment_amount")
	_creditAccountORM.LastPaymentDate = field.NewString(tableName, "last_payment_date")
	_creditAccountORM.LastStatementBalance = field.NewFloat64(tableName, "last_statement_balance")
	_creditAccountORM.LastStatementIssueDate = field.NewString(tableName, "last_statement_issue_date")
	_creditAccountORM.LinkId = field.NewUint64(tableName, "link_id")
	_creditAccountORM.MinimumAmountDueDate = field.NewFloat64(tableName, "minimum_amount_due_date")
	_creditAccountORM.MinimumPaymentAmount = field.NewFloat64(tableName, "minimum_payment_amount")
	_creditAccountORM.Name = field.NewString(tableName, "name")
	_creditAccountORM.NextPaymentDate = field.NewString(tableName, "next_payment_date")
	_creditAccountORM.NextPaymentDueDate = field.NewString(tableName, "next_payment_due_date")
	_creditAccountORM.Number = field.NewString(tableName, "number")
	_creditAccountORM.PlaidAccountId = field.NewString(tableName, "plaid_account_id")
	_creditAccountORM.Status = field.NewString(tableName, "status")
	_creditAccountORM.Subtype = field.NewString(tableName, "subtype")
	_creditAccountORM.Type = field.NewString(tableName, "type")
	_creditAccountORM.UserId = field.NewString(tableName, "user_id")
	_creditAccountORM.Aprs = creditAccountORMHasManyAprs{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Aprs", "financial_servicev1.AprORM"),
	}

	_creditAccountORM.Pockets = creditAccountORMHasManyPockets{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Pockets", "financial_servicev1.PocketORM"),
		Goals: struct {
			field.RelationField
			Forecasts struct {
				field.RelationField
			}
			Milestones struct {
				field.RelationField
				Budget struct {
					field.RelationField
					Category struct {
						field.RelationField
					}
				}
			}
			Notes struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Pockets.Goals", "financial_servicev1.SmartGoalORM"),
			Forecasts: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pockets.Goals.Forecasts", "financial_servicev1.ForecastORM"),
			},
			Milestones: struct {
				field.RelationField
				Budget struct {
					field.RelationField
					Category struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("Pockets.Goals.Milestones", "financial_servicev1.MilestoneORM"),
				Budget: struct {
					field.RelationField
					Category struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Pockets.Goals.Milestones.Budget", "financial_servicev1.BudgetORM"),
					Category: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Pockets.Goals.Milestones.Budget.Category", "financial_servicev1.CategoryORM"),
					},
				},
			},
			Notes: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pockets.Goals.Notes", "financial_servicev1.SmartNoteORM"),
			},
		},
	}

	_creditAccountORM.RecurringTransactions = creditAccountORMHasManyRecurringTransactions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("RecurringTransactions", "financial_servicev1.PlaidAccountRecurringTransactionORM"),
		Notes: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("RecurringTransactions.Notes", "financial_servicev1.SmartNoteORM"),
		},
	}

	_creditAccountORM.Transactions = creditAccountORMHasManyTransactions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Transactions", "financial_servicev1.PlaidAccountTransactionORM"),
		Notes: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Transactions.Notes", "financial_servicev1.SmartNoteORM"),
		},
		Splits: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Transactions.Splits", "financial_servicev1.TransactionSplitORM"),
		},
	}

	_creditAccountORM.fillFieldMap()

	return _creditAccountORM
}

type creditAccountORM struct {
	creditAccountORMDo

	ALL                    field.Asterisk
	Balance                field.Float32
	BalanceLimit           field.Uint64
	CurrentFunds           field.Float64
	Id                     field.Uint64
	IsOverdue              field.Bool
	LastPaymentAmount      field.Float64
	LastPaymentDate        field.String
	LastStatementBalance   field.Float64
	LastStatementIssueDate field.String
	LinkId                 field.Uint64
	MinimumAmountDueDate   field.Float64
	MinimumPaymentAmount   field.Float64
	Name                   field.String
	NextPaymentDate        field.String
	NextPaymentDueDate     field.String
	Number                 field.String
	PlaidAccountId         field.String
	Status                 field.String
	Subtype                field.String
	Type                   field.String
	UserId                 field.String
	Aprs                   creditAccountORMHasManyAprs

	Pockets creditAccountORMHasManyPockets

	RecurringTransactions creditAccountORMHasManyRecurringTransactions

	Transactions creditAccountORMHasManyTransactions

	fieldMap map[string]field.Expr
}

func (c creditAccountORM) Table(newTableName string) *creditAccountORM {
	c.creditAccountORMDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c creditAccountORM) As(alias string) *creditAccountORM {
	c.creditAccountORMDo.DO = *(c.creditAccountORMDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *creditAccountORM) updateTableName(table string) *creditAccountORM {
	c.ALL = field.NewAsterisk(table)
	c.Balance = field.NewFloat32(table, "balance")
	c.BalanceLimit = field.NewUint64(table, "balance_limit")
	c.CurrentFunds = field.NewFloat64(table, "current_funds")
	c.Id = field.NewUint64(table, "id")
	c.IsOverdue = field.NewBool(table, "is_overdue")
	c.LastPaymentAmount = field.NewFloat64(table, "last_payment_amount")
	c.LastPaymentDate = field.NewString(table, "last_payment_date")
	c.LastStatementBalance = field.NewFloat64(table, "last_statement_balance")
	c.LastStatementIssueDate = field.NewString(table, "last_statement_issue_date")
	c.LinkId = field.NewUint64(table, "link_id")
	c.MinimumAmountDueDate = field.NewFloat64(table, "minimum_amount_due_date")
	c.MinimumPaymentAmount = field.NewFloat64(table, "minimum_payment_amount")
	c.Name = field.NewString(table, "name")
	c.NextPaymentDate = field.NewString(table, "next_payment_date")
	c.NextPaymentDueDate = field.NewString(table, "next_payment_due_date")
	c.Number = field.NewString(table, "number")
	c.PlaidAccountId = field.NewString(table, "plaid_account_id")
	c.Status = field.NewString(table, "status")
	c.Subtype = field.NewString(table, "subtype")
	c.Type = field.NewString(table, "type")
	c.UserId = field.NewString(table, "user_id")

	c.fillFieldMap()

	return c
}

func (c *creditAccountORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *creditAccountORM) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 25)
	c.fieldMap["balance"] = c.Balance
	c.fieldMap["balance_limit"] = c.BalanceLimit
	c.fieldMap["current_funds"] = c.CurrentFunds
	c.fieldMap["id"] = c.Id
	c.fieldMap["is_overdue"] = c.IsOverdue
	c.fieldMap["last_payment_amount"] = c.LastPaymentAmount
	c.fieldMap["last_payment_date"] = c.LastPaymentDate
	c.fieldMap["last_statement_balance"] = c.LastStatementBalance
	c.fieldMap["last_statement_issue_date"] = c.LastStatementIssueDate
	c.fieldMap["link_id"] = c.LinkId
	c.fieldMap["minimum_amount_due_date"] = c.MinimumAmountDueDate
	c.fieldMap["minimum_payment_amount"] = c.MinimumPaymentAmount
	c.fieldMap["name"] = c.Name
	c.fieldMap["next_payment_date"] = c.NextPaymentDate
	c.fieldMap["next_payment_due_date"] = c.NextPaymentDueDate
	c.fieldMap["number"] = c.Number
	c.fieldMap["plaid_account_id"] = c.PlaidAccountId
	c.fieldMap["status"] = c.Status
	c.fieldMap["subtype"] = c.Subtype
	c.fieldMap["type"] = c.Type
	c.fieldMap["user_id"] = c.UserId

}

func (c creditAccountORM) clone(db *gorm.DB) creditAccountORM {
	c.creditAccountORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c creditAccountORM) replaceDB(db *gorm.DB) creditAccountORM {
	c.creditAccountORMDo.ReplaceDB(db)
	return c
}

type creditAccountORMHasManyAprs struct {
	db *gorm.DB

	field.RelationField
}

func (a creditAccountORMHasManyAprs) Where(conds ...field.Expr) *creditAccountORMHasManyAprs {
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

func (a creditAccountORMHasManyAprs) WithContext(ctx context.Context) *creditAccountORMHasManyAprs {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a creditAccountORMHasManyAprs) Session(session *gorm.Session) *creditAccountORMHasManyAprs {
	a.db = a.db.Session(session)
	return &a
}

func (a creditAccountORMHasManyAprs) Model(m *financial_servicev1.CreditAccountORM) *creditAccountORMHasManyAprsTx {
	return &creditAccountORMHasManyAprsTx{a.db.Model(m).Association(a.Name())}
}

type creditAccountORMHasManyAprsTx struct{ tx *gorm.Association }

func (a creditAccountORMHasManyAprsTx) Find() (result []*financial_servicev1.AprORM, err error) {
	return result, a.tx.Find(&result)
}

func (a creditAccountORMHasManyAprsTx) Append(values ...*financial_servicev1.AprORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a creditAccountORMHasManyAprsTx) Replace(values ...*financial_servicev1.AprORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a creditAccountORMHasManyAprsTx) Delete(values ...*financial_servicev1.AprORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a creditAccountORMHasManyAprsTx) Clear() error {
	return a.tx.Clear()
}

func (a creditAccountORMHasManyAprsTx) Count() int64 {
	return a.tx.Count()
}

type creditAccountORMHasManyPockets struct {
	db *gorm.DB

	field.RelationField

	Goals struct {
		field.RelationField
		Forecasts struct {
			field.RelationField
		}
		Milestones struct {
			field.RelationField
			Budget struct {
				field.RelationField
				Category struct {
					field.RelationField
				}
			}
		}
		Notes struct {
			field.RelationField
		}
	}
}

func (a creditAccountORMHasManyPockets) Where(conds ...field.Expr) *creditAccountORMHasManyPockets {
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

func (a creditAccountORMHasManyPockets) WithContext(ctx context.Context) *creditAccountORMHasManyPockets {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a creditAccountORMHasManyPockets) Session(session *gorm.Session) *creditAccountORMHasManyPockets {
	a.db = a.db.Session(session)
	return &a
}

func (a creditAccountORMHasManyPockets) Model(m *financial_servicev1.CreditAccountORM) *creditAccountORMHasManyPocketsTx {
	return &creditAccountORMHasManyPocketsTx{a.db.Model(m).Association(a.Name())}
}

type creditAccountORMHasManyPocketsTx struct{ tx *gorm.Association }

func (a creditAccountORMHasManyPocketsTx) Find() (result []*financial_servicev1.PocketORM, err error) {
	return result, a.tx.Find(&result)
}

func (a creditAccountORMHasManyPocketsTx) Append(values ...*financial_servicev1.PocketORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a creditAccountORMHasManyPocketsTx) Replace(values ...*financial_servicev1.PocketORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a creditAccountORMHasManyPocketsTx) Delete(values ...*financial_servicev1.PocketORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a creditAccountORMHasManyPocketsTx) Clear() error {
	return a.tx.Clear()
}

func (a creditAccountORMHasManyPocketsTx) Count() int64 {
	return a.tx.Count()
}

type creditAccountORMHasManyRecurringTransactions struct {
	db *gorm.DB

	field.RelationField

	Notes struct {
		field.RelationField
	}
}

func (a creditAccountORMHasManyRecurringTransactions) Where(conds ...field.Expr) *creditAccountORMHasManyRecurringTransactions {
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

func (a creditAccountORMHasManyRecurringTransactions) WithContext(ctx context.Context) *creditAccountORMHasManyRecurringTransactions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a creditAccountORMHasManyRecurringTransactions) Session(session *gorm.Session) *creditAccountORMHasManyRecurringTransactions {
	a.db = a.db.Session(session)
	return &a
}

func (a creditAccountORMHasManyRecurringTransactions) Model(m *financial_servicev1.CreditAccountORM) *creditAccountORMHasManyRecurringTransactionsTx {
	return &creditAccountORMHasManyRecurringTransactionsTx{a.db.Model(m).Association(a.Name())}
}

type creditAccountORMHasManyRecurringTransactionsTx struct{ tx *gorm.Association }

func (a creditAccountORMHasManyRecurringTransactionsTx) Find() (result []*financial_servicev1.PlaidAccountRecurringTransactionORM, err error) {
	return result, a.tx.Find(&result)
}

func (a creditAccountORMHasManyRecurringTransactionsTx) Append(values ...*financial_servicev1.PlaidAccountRecurringTransactionORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a creditAccountORMHasManyRecurringTransactionsTx) Replace(values ...*financial_servicev1.PlaidAccountRecurringTransactionORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a creditAccountORMHasManyRecurringTransactionsTx) Delete(values ...*financial_servicev1.PlaidAccountRecurringTransactionORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a creditAccountORMHasManyRecurringTransactionsTx) Clear() error {
	return a.tx.Clear()
}

func (a creditAccountORMHasManyRecurringTransactionsTx) Count() int64 {
	return a.tx.Count()
}

type creditAccountORMHasManyTransactions struct {
	db *gorm.DB

	field.RelationField

	Notes struct {
		field.RelationField
	}
	Splits struct {
		field.RelationField
	}
}

func (a creditAccountORMHasManyTransactions) Where(conds ...field.Expr) *creditAccountORMHasManyTransactions {
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

func (a creditAccountORMHasManyTransactions) WithContext(ctx context.Context) *creditAccountORMHasManyTransactions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a creditAccountORMHasManyTransactions) Session(session *gorm.Session) *creditAccountORMHasManyTransactions {
	a.db = a.db.Session(session)
	return &a
}

func (a creditAccountORMHasManyTransactions) Model(m *financial_servicev1.CreditAccountORM) *creditAccountORMHasManyTransactionsTx {
	return &creditAccountORMHasManyTransactionsTx{a.db.Model(m).Association(a.Name())}
}

type creditAccountORMHasManyTransactionsTx struct{ tx *gorm.Association }

func (a creditAccountORMHasManyTransactionsTx) Find() (result []*financial_servicev1.PlaidAccountTransactionORM, err error) {
	return result, a.tx.Find(&result)
}

func (a creditAccountORMHasManyTransactionsTx) Append(values ...*financial_servicev1.PlaidAccountTransactionORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a creditAccountORMHasManyTransactionsTx) Replace(values ...*financial_servicev1.PlaidAccountTransactionORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a creditAccountORMHasManyTransactionsTx) Delete(values ...*financial_servicev1.PlaidAccountTransactionORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a creditAccountORMHasManyTransactionsTx) Clear() error {
	return a.tx.Clear()
}

func (a creditAccountORMHasManyTransactionsTx) Count() int64 {
	return a.tx.Count()
}

type creditAccountORMDo struct{ gen.DO }

type ICreditAccountORMDo interface {
	gen.SubQuery
	Debug() ICreditAccountORMDo
	WithContext(ctx context.Context) ICreditAccountORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICreditAccountORMDo
	WriteDB() ICreditAccountORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICreditAccountORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICreditAccountORMDo
	Not(conds ...gen.Condition) ICreditAccountORMDo
	Or(conds ...gen.Condition) ICreditAccountORMDo
	Select(conds ...field.Expr) ICreditAccountORMDo
	Where(conds ...gen.Condition) ICreditAccountORMDo
	Order(conds ...field.Expr) ICreditAccountORMDo
	Distinct(cols ...field.Expr) ICreditAccountORMDo
	Omit(cols ...field.Expr) ICreditAccountORMDo
	Join(table schema.Tabler, on ...field.Expr) ICreditAccountORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICreditAccountORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICreditAccountORMDo
	Group(cols ...field.Expr) ICreditAccountORMDo
	Having(conds ...gen.Condition) ICreditAccountORMDo
	Limit(limit int) ICreditAccountORMDo
	Offset(offset int) ICreditAccountORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICreditAccountORMDo
	Unscoped() ICreditAccountORMDo
	Create(values ...*financial_servicev1.CreditAccountORM) error
	CreateInBatches(values []*financial_servicev1.CreditAccountORM, batchSize int) error
	Save(values ...*financial_servicev1.CreditAccountORM) error
	First() (*financial_servicev1.CreditAccountORM, error)
	Take() (*financial_servicev1.CreditAccountORM, error)
	Last() (*financial_servicev1.CreditAccountORM, error)
	Find() ([]*financial_servicev1.CreditAccountORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_servicev1.CreditAccountORM, err error)
	FindInBatches(result *[]*financial_servicev1.CreditAccountORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_servicev1.CreditAccountORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICreditAccountORMDo
	Assign(attrs ...field.AssignExpr) ICreditAccountORMDo
	Joins(fields ...field.RelationField) ICreditAccountORMDo
	Preload(fields ...field.RelationField) ICreditAccountORMDo
	FirstOrInit() (*financial_servicev1.CreditAccountORM, error)
	FirstOrCreate() (*financial_servicev1.CreditAccountORM, error)
	FindByPage(offset int, limit int) (result []*financial_servicev1.CreditAccountORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICreditAccountORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result financial_servicev1.CreditAccountORM, err error)
	GetRecordByIDs(ids []int) (result []financial_servicev1.CreditAccountORM, err error)
	CreateRecord(item financial_servicev1.CreditAccountORM) (err error)
	UpdateRecordByID(id int, item financial_servicev1.CreditAccountORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []financial_servicev1.CreditAccountORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result financial_servicev1.CreditAccountORM, err error)
	GetByIDs(ids []uint64) (result []financial_servicev1.CreditAccountORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c creditAccountORMDo) GetRecordByID(id int) (result financial_servicev1.CreditAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_accounts ")
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
func (c creditAccountORMDo) GetRecordByIDs(ids []int) (result []financial_servicev1.CreditAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_accounts ")
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
func (c creditAccountORMDo) CreateRecord(item financial_servicev1.CreditAccountORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO credit_accounts (columns) VALUES (values) ")

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
func (c creditAccountORMDo) UpdateRecordByID(id int, item financial_servicev1.CreditAccountORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE credit_accounts SET columns=values ")
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
func (c creditAccountORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM credit_accounts ")
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
func (c creditAccountORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []financial_servicev1.CreditAccountORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_accounts ORDER BY " + c.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (c creditAccountORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM credit_accounts ")

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
func (c creditAccountORMDo) GetByID(id uint64) (result financial_servicev1.CreditAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_accounts ")
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
func (c creditAccountORMDo) GetByIDs(ids []uint64) (result []financial_servicev1.CreditAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM credit_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c creditAccountORMDo) Debug() ICreditAccountORMDo {
	return c.withDO(c.DO.Debug())
}

func (c creditAccountORMDo) WithContext(ctx context.Context) ICreditAccountORMDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c creditAccountORMDo) ReadDB() ICreditAccountORMDo {
	return c.Clauses(dbresolver.Read)
}

func (c creditAccountORMDo) WriteDB() ICreditAccountORMDo {
	return c.Clauses(dbresolver.Write)
}

func (c creditAccountORMDo) Session(config *gorm.Session) ICreditAccountORMDo {
	return c.withDO(c.DO.Session(config))
}

func (c creditAccountORMDo) Clauses(conds ...clause.Expression) ICreditAccountORMDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c creditAccountORMDo) Returning(value interface{}, columns ...string) ICreditAccountORMDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c creditAccountORMDo) Not(conds ...gen.Condition) ICreditAccountORMDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c creditAccountORMDo) Or(conds ...gen.Condition) ICreditAccountORMDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c creditAccountORMDo) Select(conds ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c creditAccountORMDo) Where(conds ...gen.Condition) ICreditAccountORMDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c creditAccountORMDo) Order(conds ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c creditAccountORMDo) Distinct(cols ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c creditAccountORMDo) Omit(cols ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c creditAccountORMDo) Join(table schema.Tabler, on ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c creditAccountORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c creditAccountORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c creditAccountORMDo) Group(cols ...field.Expr) ICreditAccountORMDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c creditAccountORMDo) Having(conds ...gen.Condition) ICreditAccountORMDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c creditAccountORMDo) Limit(limit int) ICreditAccountORMDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c creditAccountORMDo) Offset(offset int) ICreditAccountORMDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c creditAccountORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICreditAccountORMDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c creditAccountORMDo) Unscoped() ICreditAccountORMDo {
	return c.withDO(c.DO.Unscoped())
}

func (c creditAccountORMDo) Create(values ...*financial_servicev1.CreditAccountORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c creditAccountORMDo) CreateInBatches(values []*financial_servicev1.CreditAccountORM, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c creditAccountORMDo) Save(values ...*financial_servicev1.CreditAccountORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c creditAccountORMDo) First() (*financial_servicev1.CreditAccountORM, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.CreditAccountORM), nil
	}
}

func (c creditAccountORMDo) Take() (*financial_servicev1.CreditAccountORM, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.CreditAccountORM), nil
	}
}

func (c creditAccountORMDo) Last() (*financial_servicev1.CreditAccountORM, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.CreditAccountORM), nil
	}
}

func (c creditAccountORMDo) Find() ([]*financial_servicev1.CreditAccountORM, error) {
	result, err := c.DO.Find()
	return result.([]*financial_servicev1.CreditAccountORM), err
}

func (c creditAccountORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_servicev1.CreditAccountORM, err error) {
	buf := make([]*financial_servicev1.CreditAccountORM, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c creditAccountORMDo) FindInBatches(result *[]*financial_servicev1.CreditAccountORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c creditAccountORMDo) Attrs(attrs ...field.AssignExpr) ICreditAccountORMDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c creditAccountORMDo) Assign(attrs ...field.AssignExpr) ICreditAccountORMDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c creditAccountORMDo) Joins(fields ...field.RelationField) ICreditAccountORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c creditAccountORMDo) Preload(fields ...field.RelationField) ICreditAccountORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c creditAccountORMDo) FirstOrInit() (*financial_servicev1.CreditAccountORM, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.CreditAccountORM), nil
	}
}

func (c creditAccountORMDo) FirstOrCreate() (*financial_servicev1.CreditAccountORM, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.CreditAccountORM), nil
	}
}

func (c creditAccountORMDo) FindByPage(offset int, limit int) (result []*financial_servicev1.CreditAccountORM, count int64, err error) {
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

func (c creditAccountORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c creditAccountORMDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c creditAccountORMDo) Delete(models ...*financial_servicev1.CreditAccountORM) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *creditAccountORMDo) withDO(do gen.Dao) *creditAccountORMDo {
	c.DO = *do.(*gen.DO)
	return c
}
