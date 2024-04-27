// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                    = new(Query)
	AccountStatementsORM                 *accountStatementsORM
	ActionableInsightORM                 *actionableInsightORM
	AddressORM                           *addressORM
	AprORM                               *aprORM
	BankAccountORM                       *bankAccountORM
	BudgetORM                            *budgetORM
	CategoryORM                          *categoryORM
	CreditAccountORM                     *creditAccountORM
	FinancialUserProfileORM              *financialUserProfileORM
	ForecastORM                          *forecastORM
	InvesmentHoldingORM                  *invesmentHoldingORM
	InvestmentAccountORM                 *investmentAccountORM
	InvestmentSecurityORM                *investmentSecurityORM
	LinkORM                              *linkORM
	MilestoneORM                         *milestoneORM
	MortgageAccountORM                   *mortgageAccountORM
	PersonalActionableInsightORM         *personalActionableInsightORM
	PlaidAccountInvestmentTransactionORM *plaidAccountInvestmentTransactionORM
	PlaidAccountRecurringTransactionORM  *plaidAccountRecurringTransactionORM
	PlaidAccountTransactionORM           *plaidAccountTransactionORM
	PlaidLinkORM                         *plaidLinkORM
	PlaidSyncORM                         *plaidSyncORM
	PocketORM                            *pocketORM
	SmartGoalORM                         *smartGoalORM
	SmartNoteORM                         *smartNoteORM
	StripeSubscriptionORM                *stripeSubscriptionORM
	StudentLoanAccountORM                *studentLoanAccountORM
	TokenORM                             *tokenORM
	TransactionSplitORM                  *transactionSplitORM
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AccountStatementsORM = &Q.AccountStatementsORM
	ActionableInsightORM = &Q.ActionableInsightORM
	AddressORM = &Q.AddressORM
	AprORM = &Q.AprORM
	BankAccountORM = &Q.BankAccountORM
	BudgetORM = &Q.BudgetORM
	CategoryORM = &Q.CategoryORM
	CreditAccountORM = &Q.CreditAccountORM
	FinancialUserProfileORM = &Q.FinancialUserProfileORM
	ForecastORM = &Q.ForecastORM
	InvesmentHoldingORM = &Q.InvesmentHoldingORM
	InvestmentAccountORM = &Q.InvestmentAccountORM
	InvestmentSecurityORM = &Q.InvestmentSecurityORM
	LinkORM = &Q.LinkORM
	MilestoneORM = &Q.MilestoneORM
	MortgageAccountORM = &Q.MortgageAccountORM
	PersonalActionableInsightORM = &Q.PersonalActionableInsightORM
	PlaidAccountInvestmentTransactionORM = &Q.PlaidAccountInvestmentTransactionORM
	PlaidAccountRecurringTransactionORM = &Q.PlaidAccountRecurringTransactionORM
	PlaidAccountTransactionORM = &Q.PlaidAccountTransactionORM
	PlaidLinkORM = &Q.PlaidLinkORM
	PlaidSyncORM = &Q.PlaidSyncORM
	PocketORM = &Q.PocketORM
	SmartGoalORM = &Q.SmartGoalORM
	SmartNoteORM = &Q.SmartNoteORM
	StripeSubscriptionORM = &Q.StripeSubscriptionORM
	StudentLoanAccountORM = &Q.StudentLoanAccountORM
	TokenORM = &Q.TokenORM
	TransactionSplitORM = &Q.TransactionSplitORM
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                                   db,
		AccountStatementsORM:                 newAccountStatementsORM(db, opts...),
		ActionableInsightORM:                 newActionableInsightORM(db, opts...),
		AddressORM:                           newAddressORM(db, opts...),
		AprORM:                               newAprORM(db, opts...),
		BankAccountORM:                       newBankAccountORM(db, opts...),
		BudgetORM:                            newBudgetORM(db, opts...),
		CategoryORM:                          newCategoryORM(db, opts...),
		CreditAccountORM:                     newCreditAccountORM(db, opts...),
		FinancialUserProfileORM:              newFinancialUserProfileORM(db, opts...),
		ForecastORM:                          newForecastORM(db, opts...),
		InvesmentHoldingORM:                  newInvesmentHoldingORM(db, opts...),
		InvestmentAccountORM:                 newInvestmentAccountORM(db, opts...),
		InvestmentSecurityORM:                newInvestmentSecurityORM(db, opts...),
		LinkORM:                              newLinkORM(db, opts...),
		MilestoneORM:                         newMilestoneORM(db, opts...),
		MortgageAccountORM:                   newMortgageAccountORM(db, opts...),
		PersonalActionableInsightORM:         newPersonalActionableInsightORM(db, opts...),
		PlaidAccountInvestmentTransactionORM: newPlaidAccountInvestmentTransactionORM(db, opts...),
		PlaidAccountRecurringTransactionORM:  newPlaidAccountRecurringTransactionORM(db, opts...),
		PlaidAccountTransactionORM:           newPlaidAccountTransactionORM(db, opts...),
		PlaidLinkORM:                         newPlaidLinkORM(db, opts...),
		PlaidSyncORM:                         newPlaidSyncORM(db, opts...),
		PocketORM:                            newPocketORM(db, opts...),
		SmartGoalORM:                         newSmartGoalORM(db, opts...),
		SmartNoteORM:                         newSmartNoteORM(db, opts...),
		StripeSubscriptionORM:                newStripeSubscriptionORM(db, opts...),
		StudentLoanAccountORM:                newStudentLoanAccountORM(db, opts...),
		TokenORM:                             newTokenORM(db, opts...),
		TransactionSplitORM:                  newTransactionSplitORM(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AccountStatementsORM                 accountStatementsORM
	ActionableInsightORM                 actionableInsightORM
	AddressORM                           addressORM
	AprORM                               aprORM
	BankAccountORM                       bankAccountORM
	BudgetORM                            budgetORM
	CategoryORM                          categoryORM
	CreditAccountORM                     creditAccountORM
	FinancialUserProfileORM              financialUserProfileORM
	ForecastORM                          forecastORM
	InvesmentHoldingORM                  invesmentHoldingORM
	InvestmentAccountORM                 investmentAccountORM
	InvestmentSecurityORM                investmentSecurityORM
	LinkORM                              linkORM
	MilestoneORM                         milestoneORM
	MortgageAccountORM                   mortgageAccountORM
	PersonalActionableInsightORM         personalActionableInsightORM
	PlaidAccountInvestmentTransactionORM plaidAccountInvestmentTransactionORM
	PlaidAccountRecurringTransactionORM  plaidAccountRecurringTransactionORM
	PlaidAccountTransactionORM           plaidAccountTransactionORM
	PlaidLinkORM                         plaidLinkORM
	PlaidSyncORM                         plaidSyncORM
	PocketORM                            pocketORM
	SmartGoalORM                         smartGoalORM
	SmartNoteORM                         smartNoteORM
	StripeSubscriptionORM                stripeSubscriptionORM
	StudentLoanAccountORM                studentLoanAccountORM
	TokenORM                             tokenORM
	TransactionSplitORM                  transactionSplitORM
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                                   db,
		AccountStatementsORM:                 q.AccountStatementsORM.clone(db),
		ActionableInsightORM:                 q.ActionableInsightORM.clone(db),
		AddressORM:                           q.AddressORM.clone(db),
		AprORM:                               q.AprORM.clone(db),
		BankAccountORM:                       q.BankAccountORM.clone(db),
		BudgetORM:                            q.BudgetORM.clone(db),
		CategoryORM:                          q.CategoryORM.clone(db),
		CreditAccountORM:                     q.CreditAccountORM.clone(db),
		FinancialUserProfileORM:              q.FinancialUserProfileORM.clone(db),
		ForecastORM:                          q.ForecastORM.clone(db),
		InvesmentHoldingORM:                  q.InvesmentHoldingORM.clone(db),
		InvestmentAccountORM:                 q.InvestmentAccountORM.clone(db),
		InvestmentSecurityORM:                q.InvestmentSecurityORM.clone(db),
		LinkORM:                              q.LinkORM.clone(db),
		MilestoneORM:                         q.MilestoneORM.clone(db),
		MortgageAccountORM:                   q.MortgageAccountORM.clone(db),
		PersonalActionableInsightORM:         q.PersonalActionableInsightORM.clone(db),
		PlaidAccountInvestmentTransactionORM: q.PlaidAccountInvestmentTransactionORM.clone(db),
		PlaidAccountRecurringTransactionORM:  q.PlaidAccountRecurringTransactionORM.clone(db),
		PlaidAccountTransactionORM:           q.PlaidAccountTransactionORM.clone(db),
		PlaidLinkORM:                         q.PlaidLinkORM.clone(db),
		PlaidSyncORM:                         q.PlaidSyncORM.clone(db),
		PocketORM:                            q.PocketORM.clone(db),
		SmartGoalORM:                         q.SmartGoalORM.clone(db),
		SmartNoteORM:                         q.SmartNoteORM.clone(db),
		StripeSubscriptionORM:                q.StripeSubscriptionORM.clone(db),
		StudentLoanAccountORM:                q.StudentLoanAccountORM.clone(db),
		TokenORM:                             q.TokenORM.clone(db),
		TransactionSplitORM:                  q.TransactionSplitORM.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                                   db,
		AccountStatementsORM:                 q.AccountStatementsORM.replaceDB(db),
		ActionableInsightORM:                 q.ActionableInsightORM.replaceDB(db),
		AddressORM:                           q.AddressORM.replaceDB(db),
		AprORM:                               q.AprORM.replaceDB(db),
		BankAccountORM:                       q.BankAccountORM.replaceDB(db),
		BudgetORM:                            q.BudgetORM.replaceDB(db),
		CategoryORM:                          q.CategoryORM.replaceDB(db),
		CreditAccountORM:                     q.CreditAccountORM.replaceDB(db),
		FinancialUserProfileORM:              q.FinancialUserProfileORM.replaceDB(db),
		ForecastORM:                          q.ForecastORM.replaceDB(db),
		InvesmentHoldingORM:                  q.InvesmentHoldingORM.replaceDB(db),
		InvestmentAccountORM:                 q.InvestmentAccountORM.replaceDB(db),
		InvestmentSecurityORM:                q.InvestmentSecurityORM.replaceDB(db),
		LinkORM:                              q.LinkORM.replaceDB(db),
		MilestoneORM:                         q.MilestoneORM.replaceDB(db),
		MortgageAccountORM:                   q.MortgageAccountORM.replaceDB(db),
		PersonalActionableInsightORM:         q.PersonalActionableInsightORM.replaceDB(db),
		PlaidAccountInvestmentTransactionORM: q.PlaidAccountInvestmentTransactionORM.replaceDB(db),
		PlaidAccountRecurringTransactionORM:  q.PlaidAccountRecurringTransactionORM.replaceDB(db),
		PlaidAccountTransactionORM:           q.PlaidAccountTransactionORM.replaceDB(db),
		PlaidLinkORM:                         q.PlaidLinkORM.replaceDB(db),
		PlaidSyncORM:                         q.PlaidSyncORM.replaceDB(db),
		PocketORM:                            q.PocketORM.replaceDB(db),
		SmartGoalORM:                         q.SmartGoalORM.replaceDB(db),
		SmartNoteORM:                         q.SmartNoteORM.replaceDB(db),
		StripeSubscriptionORM:                q.StripeSubscriptionORM.replaceDB(db),
		StudentLoanAccountORM:                q.StudentLoanAccountORM.replaceDB(db),
		TokenORM:                             q.TokenORM.replaceDB(db),
		TransactionSplitORM:                  q.TransactionSplitORM.replaceDB(db),
	}
}

type queryCtx struct {
	AccountStatementsORM                 IAccountStatementsORMDo
	ActionableInsightORM                 IActionableInsightORMDo
	AddressORM                           IAddressORMDo
	AprORM                               IAprORMDo
	BankAccountORM                       IBankAccountORMDo
	BudgetORM                            IBudgetORMDo
	CategoryORM                          ICategoryORMDo
	CreditAccountORM                     ICreditAccountORMDo
	FinancialUserProfileORM              IFinancialUserProfileORMDo
	ForecastORM                          IForecastORMDo
	InvesmentHoldingORM                  IInvesmentHoldingORMDo
	InvestmentAccountORM                 IInvestmentAccountORMDo
	InvestmentSecurityORM                IInvestmentSecurityORMDo
	LinkORM                              ILinkORMDo
	MilestoneORM                         IMilestoneORMDo
	MortgageAccountORM                   IMortgageAccountORMDo
	PersonalActionableInsightORM         IPersonalActionableInsightORMDo
	PlaidAccountInvestmentTransactionORM IPlaidAccountInvestmentTransactionORMDo
	PlaidAccountRecurringTransactionORM  IPlaidAccountRecurringTransactionORMDo
	PlaidAccountTransactionORM           IPlaidAccountTransactionORMDo
	PlaidLinkORM                         IPlaidLinkORMDo
	PlaidSyncORM                         IPlaidSyncORMDo
	PocketORM                            IPocketORMDo
	SmartGoalORM                         ISmartGoalORMDo
	SmartNoteORM                         ISmartNoteORMDo
	StripeSubscriptionORM                IStripeSubscriptionORMDo
	StudentLoanAccountORM                IStudentLoanAccountORMDo
	TokenORM                             ITokenORMDo
	TransactionSplitORM                  ITransactionSplitORMDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccountStatementsORM:                 q.AccountStatementsORM.WithContext(ctx),
		ActionableInsightORM:                 q.ActionableInsightORM.WithContext(ctx),
		AddressORM:                           q.AddressORM.WithContext(ctx),
		AprORM:                               q.AprORM.WithContext(ctx),
		BankAccountORM:                       q.BankAccountORM.WithContext(ctx),
		BudgetORM:                            q.BudgetORM.WithContext(ctx),
		CategoryORM:                          q.CategoryORM.WithContext(ctx),
		CreditAccountORM:                     q.CreditAccountORM.WithContext(ctx),
		FinancialUserProfileORM:              q.FinancialUserProfileORM.WithContext(ctx),
		ForecastORM:                          q.ForecastORM.WithContext(ctx),
		InvesmentHoldingORM:                  q.InvesmentHoldingORM.WithContext(ctx),
		InvestmentAccountORM:                 q.InvestmentAccountORM.WithContext(ctx),
		InvestmentSecurityORM:                q.InvestmentSecurityORM.WithContext(ctx),
		LinkORM:                              q.LinkORM.WithContext(ctx),
		MilestoneORM:                         q.MilestoneORM.WithContext(ctx),
		MortgageAccountORM:                   q.MortgageAccountORM.WithContext(ctx),
		PersonalActionableInsightORM:         q.PersonalActionableInsightORM.WithContext(ctx),
		PlaidAccountInvestmentTransactionORM: q.PlaidAccountInvestmentTransactionORM.WithContext(ctx),
		PlaidAccountRecurringTransactionORM:  q.PlaidAccountRecurringTransactionORM.WithContext(ctx),
		PlaidAccountTransactionORM:           q.PlaidAccountTransactionORM.WithContext(ctx),
		PlaidLinkORM:                         q.PlaidLinkORM.WithContext(ctx),
		PlaidSyncORM:                         q.PlaidSyncORM.WithContext(ctx),
		PocketORM:                            q.PocketORM.WithContext(ctx),
		SmartGoalORM:                         q.SmartGoalORM.WithContext(ctx),
		SmartNoteORM:                         q.SmartNoteORM.WithContext(ctx),
		StripeSubscriptionORM:                q.StripeSubscriptionORM.WithContext(ctx),
		StudentLoanAccountORM:                q.StudentLoanAccountORM.WithContext(ctx),
		TokenORM:                             q.TokenORM.WithContext(ctx),
		TransactionSplitORM:                  q.TransactionSplitORM.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
