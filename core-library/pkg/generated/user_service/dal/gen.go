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
	Q                        = new(Query)
	AddressORM               *addressORM
	BusinessAccountORM       *businessAccountORM
	DigitalWorkerSettingsORM *digitalWorkerSettingsORM
	FinancialPreferencesORM  *financialPreferencesORM
	NotificationSettingsORM  *notificationSettingsORM
	RoleAuditEventsORM       *roleAuditEventsORM
	RoleORM                  *roleORM
	SettingsORM              *settingsORM
	TagsORM                  *tagsORM
	UserAccountORM           *userAccountORM
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AddressORM = &Q.AddressORM
	BusinessAccountORM = &Q.BusinessAccountORM
	DigitalWorkerSettingsORM = &Q.DigitalWorkerSettingsORM
	FinancialPreferencesORM = &Q.FinancialPreferencesORM
	NotificationSettingsORM = &Q.NotificationSettingsORM
	RoleAuditEventsORM = &Q.RoleAuditEventsORM
	RoleORM = &Q.RoleORM
	SettingsORM = &Q.SettingsORM
	TagsORM = &Q.TagsORM
	UserAccountORM = &Q.UserAccountORM
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                       db,
		AddressORM:               newAddressORM(db, opts...),
		BusinessAccountORM:       newBusinessAccountORM(db, opts...),
		DigitalWorkerSettingsORM: newDigitalWorkerSettingsORM(db, opts...),
		FinancialPreferencesORM:  newFinancialPreferencesORM(db, opts...),
		NotificationSettingsORM:  newNotificationSettingsORM(db, opts...),
		RoleAuditEventsORM:       newRoleAuditEventsORM(db, opts...),
		RoleORM:                  newRoleORM(db, opts...),
		SettingsORM:              newSettingsORM(db, opts...),
		TagsORM:                  newTagsORM(db, opts...),
		UserAccountORM:           newUserAccountORM(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AddressORM               addressORM
	BusinessAccountORM       businessAccountORM
	DigitalWorkerSettingsORM digitalWorkerSettingsORM
	FinancialPreferencesORM  financialPreferencesORM
	NotificationSettingsORM  notificationSettingsORM
	RoleAuditEventsORM       roleAuditEventsORM
	RoleORM                  roleORM
	SettingsORM              settingsORM
	TagsORM                  tagsORM
	UserAccountORM           userAccountORM
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		AddressORM:               q.AddressORM.clone(db),
		BusinessAccountORM:       q.BusinessAccountORM.clone(db),
		DigitalWorkerSettingsORM: q.DigitalWorkerSettingsORM.clone(db),
		FinancialPreferencesORM:  q.FinancialPreferencesORM.clone(db),
		NotificationSettingsORM:  q.NotificationSettingsORM.clone(db),
		RoleAuditEventsORM:       q.RoleAuditEventsORM.clone(db),
		RoleORM:                  q.RoleORM.clone(db),
		SettingsORM:              q.SettingsORM.clone(db),
		TagsORM:                  q.TagsORM.clone(db),
		UserAccountORM:           q.UserAccountORM.clone(db),
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
		db:                       db,
		AddressORM:               q.AddressORM.replaceDB(db),
		BusinessAccountORM:       q.BusinessAccountORM.replaceDB(db),
		DigitalWorkerSettingsORM: q.DigitalWorkerSettingsORM.replaceDB(db),
		FinancialPreferencesORM:  q.FinancialPreferencesORM.replaceDB(db),
		NotificationSettingsORM:  q.NotificationSettingsORM.replaceDB(db),
		RoleAuditEventsORM:       q.RoleAuditEventsORM.replaceDB(db),
		RoleORM:                  q.RoleORM.replaceDB(db),
		SettingsORM:              q.SettingsORM.replaceDB(db),
		TagsORM:                  q.TagsORM.replaceDB(db),
		UserAccountORM:           q.UserAccountORM.replaceDB(db),
	}
}

type queryCtx struct {
	AddressORM               IAddressORMDo
	BusinessAccountORM       IBusinessAccountORMDo
	DigitalWorkerSettingsORM IDigitalWorkerSettingsORMDo
	FinancialPreferencesORM  IFinancialPreferencesORMDo
	NotificationSettingsORM  INotificationSettingsORMDo
	RoleAuditEventsORM       IRoleAuditEventsORMDo
	RoleORM                  IRoleORMDo
	SettingsORM              ISettingsORMDo
	TagsORM                  ITagsORMDo
	UserAccountORM           IUserAccountORMDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AddressORM:               q.AddressORM.WithContext(ctx),
		BusinessAccountORM:       q.BusinessAccountORM.WithContext(ctx),
		DigitalWorkerSettingsORM: q.DigitalWorkerSettingsORM.WithContext(ctx),
		FinancialPreferencesORM:  q.FinancialPreferencesORM.WithContext(ctx),
		NotificationSettingsORM:  q.NotificationSettingsORM.WithContext(ctx),
		RoleAuditEventsORM:       q.RoleAuditEventsORM.WithContext(ctx),
		RoleORM:                  q.RoleORM.WithContext(ctx),
		SettingsORM:              q.SettingsORM.WithContext(ctx),
		TagsORM:                  q.TagsORM.WithContext(ctx),
		UserAccountORM:           q.UserAccountORM.WithContext(ctx),
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