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

func newStudentLoanAccountORM(db *gorm.DB, opts ...gen.DOOption) studentLoanAccountORM {
	_studentLoanAccountORM := studentLoanAccountORM{}

	_studentLoanAccountORM.studentLoanAccountORMDo.UseDB(db, opts...)
	_studentLoanAccountORM.studentLoanAccountORMDo.UseModel(&financial_servicev1.StudentLoanAccountORM{})

	tableName := _studentLoanAccountORM.studentLoanAccountORMDo.TableName()
	_studentLoanAccountORM.ALL = field.NewAsterisk(tableName)
	_studentLoanAccountORM.DisbursementDates = field.NewField(tableName, "disbursement_dates")
	_studentLoanAccountORM.ExpectedPayoffDate = field.NewString(tableName, "expected_payoff_date")
	_studentLoanAccountORM.Guarantor = field.NewString(tableName, "guarantor")
	_studentLoanAccountORM.Id = field.NewUint64(tableName, "id")
	_studentLoanAccountORM.InterestRatePercentage = field.NewFloat64(tableName, "interest_rate_percentage")
	_studentLoanAccountORM.IsOverdue = field.NewBool(tableName, "is_overdue")
	_studentLoanAccountORM.LastPaymentAmount = field.NewFloat64(tableName, "last_payment_amount")
	_studentLoanAccountORM.LastPaymentDate = field.NewString(tableName, "last_payment_date")
	_studentLoanAccountORM.LastStatementIssueDate = field.NewString(tableName, "last_statement_issue_date")
	_studentLoanAccountORM.LinkId = field.NewUint64(tableName, "link_id")
	_studentLoanAccountORM.LoanEndDate = field.NewString(tableName, "loan_end_date")
	_studentLoanAccountORM.LoanName = field.NewString(tableName, "loan_name")
	_studentLoanAccountORM.LoanType = field.NewString(tableName, "loan_type")
	_studentLoanAccountORM.MinimumPaymentAmount = field.NewFloat64(tableName, "minimum_payment_amount")
	_studentLoanAccountORM.Name = field.NewString(tableName, "name")
	_studentLoanAccountORM.NextPaymentDueDate = field.NewString(tableName, "next_payment_due_date")
	_studentLoanAccountORM.OriginationDate = field.NewString(tableName, "origination_date")
	_studentLoanAccountORM.OriginationPrincipalAmount = field.NewFloat64(tableName, "origination_principal_amount")
	_studentLoanAccountORM.OutstandingInterestAmount = field.NewFloat64(tableName, "outstanding_interest_amount")
	_studentLoanAccountORM.PaymentReferenceNumber = field.NewString(tableName, "payment_reference_number")
	_studentLoanAccountORM.PlaidAccountId = field.NewString(tableName, "plaid_account_id")
	_studentLoanAccountORM.PslfStatusEstimatedEligibilityDate = field.NewString(tableName, "pslf_status_estimated_eligibility_date")
	_studentLoanAccountORM.PslfStatusPaymentsMade = field.NewInt32(tableName, "pslf_status_payments_made")
	_studentLoanAccountORM.PslfStatusPaymentsRemaining = field.NewInt32(tableName, "pslf_status_payments_remaining")
	_studentLoanAccountORM.RepaymentPlanDescription = field.NewString(tableName, "repayment_plan_description")
	_studentLoanAccountORM.RepaymentPlanType = field.NewString(tableName, "repayment_plan_type")
	_studentLoanAccountORM.SequenceNumber = field.NewString(tableName, "sequence_number")
	_studentLoanAccountORM.ServicerAddressCity = field.NewString(tableName, "servicer_address_city")
	_studentLoanAccountORM.ServicerAddressCountry = field.NewString(tableName, "servicer_address_country")
	_studentLoanAccountORM.ServicerAddressPostalCode = field.NewString(tableName, "servicer_address_postal_code")
	_studentLoanAccountORM.ServicerAddressRegion = field.NewString(tableName, "servicer_address_region")
	_studentLoanAccountORM.ServicerAddressState = field.NewString(tableName, "servicer_address_state")
	_studentLoanAccountORM.ServicerAddressStreet = field.NewString(tableName, "servicer_address_street")
	_studentLoanAccountORM.Status = field.NewString(tableName, "status")
	_studentLoanAccountORM.UserId = field.NewString(tableName, "user_id")
	_studentLoanAccountORM.YtdInterestPaid = field.NewFloat64(tableName, "ytd_interest_paid")
	_studentLoanAccountORM.YtdPrincipalPaid = field.NewFloat64(tableName, "ytd_principal_paid")
	_studentLoanAccountORM.Statements = studentLoanAccountORMHasManyStatements{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Statements", "financial_servicev1.AccountStatementsORM"),
	}

	_studentLoanAccountORM.fillFieldMap()

	return _studentLoanAccountORM
}

type studentLoanAccountORM struct {
	studentLoanAccountORMDo

	ALL                                field.Asterisk
	DisbursementDates                  field.Field
	ExpectedPayoffDate                 field.String
	Guarantor                          field.String
	Id                                 field.Uint64
	InterestRatePercentage             field.Float64
	IsOverdue                          field.Bool
	LastPaymentAmount                  field.Float64
	LastPaymentDate                    field.String
	LastStatementIssueDate             field.String
	LinkId                             field.Uint64
	LoanEndDate                        field.String
	LoanName                           field.String
	LoanType                           field.String
	MinimumPaymentAmount               field.Float64
	Name                               field.String
	NextPaymentDueDate                 field.String
	OriginationDate                    field.String
	OriginationPrincipalAmount         field.Float64
	OutstandingInterestAmount          field.Float64
	PaymentReferenceNumber             field.String
	PlaidAccountId                     field.String
	PslfStatusEstimatedEligibilityDate field.String
	PslfStatusPaymentsMade             field.Int32
	PslfStatusPaymentsRemaining        field.Int32
	RepaymentPlanDescription           field.String
	RepaymentPlanType                  field.String
	SequenceNumber                     field.String
	ServicerAddressCity                field.String
	ServicerAddressCountry             field.String
	ServicerAddressPostalCode          field.String
	ServicerAddressRegion              field.String
	ServicerAddressState               field.String
	ServicerAddressStreet              field.String
	Status                             field.String
	UserId                             field.String
	YtdInterestPaid                    field.Float64
	YtdPrincipalPaid                   field.Float64
	Statements                         studentLoanAccountORMHasManyStatements

	fieldMap map[string]field.Expr
}

func (s studentLoanAccountORM) Table(newTableName string) *studentLoanAccountORM {
	s.studentLoanAccountORMDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s studentLoanAccountORM) As(alias string) *studentLoanAccountORM {
	s.studentLoanAccountORMDo.DO = *(s.studentLoanAccountORMDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *studentLoanAccountORM) updateTableName(table string) *studentLoanAccountORM {
	s.ALL = field.NewAsterisk(table)
	s.DisbursementDates = field.NewField(table, "disbursement_dates")
	s.ExpectedPayoffDate = field.NewString(table, "expected_payoff_date")
	s.Guarantor = field.NewString(table, "guarantor")
	s.Id = field.NewUint64(table, "id")
	s.InterestRatePercentage = field.NewFloat64(table, "interest_rate_percentage")
	s.IsOverdue = field.NewBool(table, "is_overdue")
	s.LastPaymentAmount = field.NewFloat64(table, "last_payment_amount")
	s.LastPaymentDate = field.NewString(table, "last_payment_date")
	s.LastStatementIssueDate = field.NewString(table, "last_statement_issue_date")
	s.LinkId = field.NewUint64(table, "link_id")
	s.LoanEndDate = field.NewString(table, "loan_end_date")
	s.LoanName = field.NewString(table, "loan_name")
	s.LoanType = field.NewString(table, "loan_type")
	s.MinimumPaymentAmount = field.NewFloat64(table, "minimum_payment_amount")
	s.Name = field.NewString(table, "name")
	s.NextPaymentDueDate = field.NewString(table, "next_payment_due_date")
	s.OriginationDate = field.NewString(table, "origination_date")
	s.OriginationPrincipalAmount = field.NewFloat64(table, "origination_principal_amount")
	s.OutstandingInterestAmount = field.NewFloat64(table, "outstanding_interest_amount")
	s.PaymentReferenceNumber = field.NewString(table, "payment_reference_number")
	s.PlaidAccountId = field.NewString(table, "plaid_account_id")
	s.PslfStatusEstimatedEligibilityDate = field.NewString(table, "pslf_status_estimated_eligibility_date")
	s.PslfStatusPaymentsMade = field.NewInt32(table, "pslf_status_payments_made")
	s.PslfStatusPaymentsRemaining = field.NewInt32(table, "pslf_status_payments_remaining")
	s.RepaymentPlanDescription = field.NewString(table, "repayment_plan_description")
	s.RepaymentPlanType = field.NewString(table, "repayment_plan_type")
	s.SequenceNumber = field.NewString(table, "sequence_number")
	s.ServicerAddressCity = field.NewString(table, "servicer_address_city")
	s.ServicerAddressCountry = field.NewString(table, "servicer_address_country")
	s.ServicerAddressPostalCode = field.NewString(table, "servicer_address_postal_code")
	s.ServicerAddressRegion = field.NewString(table, "servicer_address_region")
	s.ServicerAddressState = field.NewString(table, "servicer_address_state")
	s.ServicerAddressStreet = field.NewString(table, "servicer_address_street")
	s.Status = field.NewString(table, "status")
	s.UserId = field.NewString(table, "user_id")
	s.YtdInterestPaid = field.NewFloat64(table, "ytd_interest_paid")
	s.YtdPrincipalPaid = field.NewFloat64(table, "ytd_principal_paid")

	s.fillFieldMap()

	return s
}

func (s *studentLoanAccountORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *studentLoanAccountORM) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 38)
	s.fieldMap["disbursement_dates"] = s.DisbursementDates
	s.fieldMap["expected_payoff_date"] = s.ExpectedPayoffDate
	s.fieldMap["guarantor"] = s.Guarantor
	s.fieldMap["id"] = s.Id
	s.fieldMap["interest_rate_percentage"] = s.InterestRatePercentage
	s.fieldMap["is_overdue"] = s.IsOverdue
	s.fieldMap["last_payment_amount"] = s.LastPaymentAmount
	s.fieldMap["last_payment_date"] = s.LastPaymentDate
	s.fieldMap["last_statement_issue_date"] = s.LastStatementIssueDate
	s.fieldMap["link_id"] = s.LinkId
	s.fieldMap["loan_end_date"] = s.LoanEndDate
	s.fieldMap["loan_name"] = s.LoanName
	s.fieldMap["loan_type"] = s.LoanType
	s.fieldMap["minimum_payment_amount"] = s.MinimumPaymentAmount
	s.fieldMap["name"] = s.Name
	s.fieldMap["next_payment_due_date"] = s.NextPaymentDueDate
	s.fieldMap["origination_date"] = s.OriginationDate
	s.fieldMap["origination_principal_amount"] = s.OriginationPrincipalAmount
	s.fieldMap["outstanding_interest_amount"] = s.OutstandingInterestAmount
	s.fieldMap["payment_reference_number"] = s.PaymentReferenceNumber
	s.fieldMap["plaid_account_id"] = s.PlaidAccountId
	s.fieldMap["pslf_status_estimated_eligibility_date"] = s.PslfStatusEstimatedEligibilityDate
	s.fieldMap["pslf_status_payments_made"] = s.PslfStatusPaymentsMade
	s.fieldMap["pslf_status_payments_remaining"] = s.PslfStatusPaymentsRemaining
	s.fieldMap["repayment_plan_description"] = s.RepaymentPlanDescription
	s.fieldMap["repayment_plan_type"] = s.RepaymentPlanType
	s.fieldMap["sequence_number"] = s.SequenceNumber
	s.fieldMap["servicer_address_city"] = s.ServicerAddressCity
	s.fieldMap["servicer_address_country"] = s.ServicerAddressCountry
	s.fieldMap["servicer_address_postal_code"] = s.ServicerAddressPostalCode
	s.fieldMap["servicer_address_region"] = s.ServicerAddressRegion
	s.fieldMap["servicer_address_state"] = s.ServicerAddressState
	s.fieldMap["servicer_address_street"] = s.ServicerAddressStreet
	s.fieldMap["status"] = s.Status
	s.fieldMap["user_id"] = s.UserId
	s.fieldMap["ytd_interest_paid"] = s.YtdInterestPaid
	s.fieldMap["ytd_principal_paid"] = s.YtdPrincipalPaid

}

func (s studentLoanAccountORM) clone(db *gorm.DB) studentLoanAccountORM {
	s.studentLoanAccountORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s studentLoanAccountORM) replaceDB(db *gorm.DB) studentLoanAccountORM {
	s.studentLoanAccountORMDo.ReplaceDB(db)
	return s
}

type studentLoanAccountORMHasManyStatements struct {
	db *gorm.DB

	field.RelationField
}

func (a studentLoanAccountORMHasManyStatements) Where(conds ...field.Expr) *studentLoanAccountORMHasManyStatements {
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

func (a studentLoanAccountORMHasManyStatements) WithContext(ctx context.Context) *studentLoanAccountORMHasManyStatements {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a studentLoanAccountORMHasManyStatements) Session(session *gorm.Session) *studentLoanAccountORMHasManyStatements {
	a.db = a.db.Session(session)
	return &a
}

func (a studentLoanAccountORMHasManyStatements) Model(m *financial_servicev1.StudentLoanAccountORM) *studentLoanAccountORMHasManyStatementsTx {
	return &studentLoanAccountORMHasManyStatementsTx{a.db.Model(m).Association(a.Name())}
}

type studentLoanAccountORMHasManyStatementsTx struct{ tx *gorm.Association }

func (a studentLoanAccountORMHasManyStatementsTx) Find() (result []*financial_servicev1.AccountStatementsORM, err error) {
	return result, a.tx.Find(&result)
}

func (a studentLoanAccountORMHasManyStatementsTx) Append(values ...*financial_servicev1.AccountStatementsORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a studentLoanAccountORMHasManyStatementsTx) Replace(values ...*financial_servicev1.AccountStatementsORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a studentLoanAccountORMHasManyStatementsTx) Delete(values ...*financial_servicev1.AccountStatementsORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a studentLoanAccountORMHasManyStatementsTx) Clear() error {
	return a.tx.Clear()
}

func (a studentLoanAccountORMHasManyStatementsTx) Count() int64 {
	return a.tx.Count()
}

type studentLoanAccountORMDo struct{ gen.DO }

type IStudentLoanAccountORMDo interface {
	gen.SubQuery
	Debug() IStudentLoanAccountORMDo
	WithContext(ctx context.Context) IStudentLoanAccountORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStudentLoanAccountORMDo
	WriteDB() IStudentLoanAccountORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStudentLoanAccountORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStudentLoanAccountORMDo
	Not(conds ...gen.Condition) IStudentLoanAccountORMDo
	Or(conds ...gen.Condition) IStudentLoanAccountORMDo
	Select(conds ...field.Expr) IStudentLoanAccountORMDo
	Where(conds ...gen.Condition) IStudentLoanAccountORMDo
	Order(conds ...field.Expr) IStudentLoanAccountORMDo
	Distinct(cols ...field.Expr) IStudentLoanAccountORMDo
	Omit(cols ...field.Expr) IStudentLoanAccountORMDo
	Join(table schema.Tabler, on ...field.Expr) IStudentLoanAccountORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStudentLoanAccountORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStudentLoanAccountORMDo
	Group(cols ...field.Expr) IStudentLoanAccountORMDo
	Having(conds ...gen.Condition) IStudentLoanAccountORMDo
	Limit(limit int) IStudentLoanAccountORMDo
	Offset(offset int) IStudentLoanAccountORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentLoanAccountORMDo
	Unscoped() IStudentLoanAccountORMDo
	Create(values ...*financial_servicev1.StudentLoanAccountORM) error
	CreateInBatches(values []*financial_servicev1.StudentLoanAccountORM, batchSize int) error
	Save(values ...*financial_servicev1.StudentLoanAccountORM) error
	First() (*financial_servicev1.StudentLoanAccountORM, error)
	Take() (*financial_servicev1.StudentLoanAccountORM, error)
	Last() (*financial_servicev1.StudentLoanAccountORM, error)
	Find() ([]*financial_servicev1.StudentLoanAccountORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_servicev1.StudentLoanAccountORM, err error)
	FindInBatches(result *[]*financial_servicev1.StudentLoanAccountORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_servicev1.StudentLoanAccountORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStudentLoanAccountORMDo
	Assign(attrs ...field.AssignExpr) IStudentLoanAccountORMDo
	Joins(fields ...field.RelationField) IStudentLoanAccountORMDo
	Preload(fields ...field.RelationField) IStudentLoanAccountORMDo
	FirstOrInit() (*financial_servicev1.StudentLoanAccountORM, error)
	FirstOrCreate() (*financial_servicev1.StudentLoanAccountORM, error)
	FindByPage(offset int, limit int) (result []*financial_servicev1.StudentLoanAccountORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStudentLoanAccountORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result financial_servicev1.StudentLoanAccountORM, err error)
	GetRecordByIDs(ids []int) (result []financial_servicev1.StudentLoanAccountORM, err error)
	CreateRecord(item financial_servicev1.StudentLoanAccountORM) (err error)
	UpdateRecordByID(id int, item financial_servicev1.StudentLoanAccountORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []financial_servicev1.StudentLoanAccountORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result financial_servicev1.StudentLoanAccountORM, err error)
	GetByIDs(ids []uint64) (result []financial_servicev1.StudentLoanAccountORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (s studentLoanAccountORMDo) GetRecordByID(id int) (result financial_servicev1.StudentLoanAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM student_loan_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (s studentLoanAccountORMDo) GetRecordByIDs(ids []int) (result []financial_servicev1.StudentLoanAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM student_loan_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (s studentLoanAccountORMDo) CreateRecord(item financial_servicev1.StudentLoanAccountORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO student_loan_accounts (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (s studentLoanAccountORMDo) UpdateRecordByID(id int, item financial_servicev1.StudentLoanAccountORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE student_loan_accounts SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (s studentLoanAccountORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM student_loan_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (s studentLoanAccountORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []financial_servicev1.StudentLoanAccountORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM student_loan_accounts ORDER BY " + s.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (s studentLoanAccountORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM student_loan_accounts ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (s studentLoanAccountORMDo) GetByID(id uint64) (result financial_servicev1.StudentLoanAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM student_loan_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (s studentLoanAccountORMDo) GetByIDs(ids []uint64) (result []financial_servicev1.StudentLoanAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM student_loan_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s studentLoanAccountORMDo) Debug() IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Debug())
}

func (s studentLoanAccountORMDo) WithContext(ctx context.Context) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studentLoanAccountORMDo) ReadDB() IStudentLoanAccountORMDo {
	return s.Clauses(dbresolver.Read)
}

func (s studentLoanAccountORMDo) WriteDB() IStudentLoanAccountORMDo {
	return s.Clauses(dbresolver.Write)
}

func (s studentLoanAccountORMDo) Session(config *gorm.Session) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Session(config))
}

func (s studentLoanAccountORMDo) Clauses(conds ...clause.Expression) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studentLoanAccountORMDo) Returning(value interface{}, columns ...string) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studentLoanAccountORMDo) Not(conds ...gen.Condition) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s studentLoanAccountORMDo) Or(conds ...gen.Condition) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s studentLoanAccountORMDo) Select(conds ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s studentLoanAccountORMDo) Where(conds ...gen.Condition) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s studentLoanAccountORMDo) Order(conds ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s studentLoanAccountORMDo) Distinct(cols ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studentLoanAccountORMDo) Omit(cols ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studentLoanAccountORMDo) Join(table schema.Tabler, on ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studentLoanAccountORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studentLoanAccountORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studentLoanAccountORMDo) Group(cols ...field.Expr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s studentLoanAccountORMDo) Having(conds ...gen.Condition) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s studentLoanAccountORMDo) Limit(limit int) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s studentLoanAccountORMDo) Offset(offset int) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s studentLoanAccountORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studentLoanAccountORMDo) Unscoped() IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Unscoped())
}

func (s studentLoanAccountORMDo) Create(values ...*financial_servicev1.StudentLoanAccountORM) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studentLoanAccountORMDo) CreateInBatches(values []*financial_servicev1.StudentLoanAccountORM, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studentLoanAccountORMDo) Save(values ...*financial_servicev1.StudentLoanAccountORM) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studentLoanAccountORMDo) First() (*financial_servicev1.StudentLoanAccountORM, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.StudentLoanAccountORM), nil
	}
}

func (s studentLoanAccountORMDo) Take() (*financial_servicev1.StudentLoanAccountORM, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.StudentLoanAccountORM), nil
	}
}

func (s studentLoanAccountORMDo) Last() (*financial_servicev1.StudentLoanAccountORM, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.StudentLoanAccountORM), nil
	}
}

func (s studentLoanAccountORMDo) Find() ([]*financial_servicev1.StudentLoanAccountORM, error) {
	result, err := s.DO.Find()
	return result.([]*financial_servicev1.StudentLoanAccountORM), err
}

func (s studentLoanAccountORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_servicev1.StudentLoanAccountORM, err error) {
	buf := make([]*financial_servicev1.StudentLoanAccountORM, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studentLoanAccountORMDo) FindInBatches(result *[]*financial_servicev1.StudentLoanAccountORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studentLoanAccountORMDo) Attrs(attrs ...field.AssignExpr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studentLoanAccountORMDo) Assign(attrs ...field.AssignExpr) IStudentLoanAccountORMDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studentLoanAccountORMDo) Joins(fields ...field.RelationField) IStudentLoanAccountORMDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studentLoanAccountORMDo) Preload(fields ...field.RelationField) IStudentLoanAccountORMDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studentLoanAccountORMDo) FirstOrInit() (*financial_servicev1.StudentLoanAccountORM, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.StudentLoanAccountORM), nil
	}
}

func (s studentLoanAccountORMDo) FirstOrCreate() (*financial_servicev1.StudentLoanAccountORM, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_servicev1.StudentLoanAccountORM), nil
	}
}

func (s studentLoanAccountORMDo) FindByPage(offset int, limit int) (result []*financial_servicev1.StudentLoanAccountORM, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s studentLoanAccountORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studentLoanAccountORMDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studentLoanAccountORMDo) Delete(models ...*financial_servicev1.StudentLoanAccountORM) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studentLoanAccountORMDo) withDO(do gen.Dao) *studentLoanAccountORMDo {
	s.DO = *do.(*gen.DO)
	return s
}
