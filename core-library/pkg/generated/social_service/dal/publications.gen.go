// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	social_servicev2 "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/social_service/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newPublicationORM(db *gorm.DB, opts ...gen.DOOption) publicationORM {
	_publicationORM := publicationORM{}

	_publicationORM.publicationORMDo.UseDB(db, opts...)
	_publicationORM.publicationORMDo.UseModel(&social_servicev2.PublicationORM{})

	tableName := _publicationORM.publicationORMDo.TableName()
	_publicationORM.ALL = field.NewAsterisk(tableName)
	_publicationORM.AdminBackendPlatformUserId = field.NewString(tableName, "admin_backend_platform_user_id")
	_publicationORM.BookmarkId = field.NewUint64(tableName, "bookmark_id")
	_publicationORM.CreatedAt = field.NewString(tableName, "created_at")
	_publicationORM.Description = field.NewString(tableName, "description")
	_publicationORM.Id = field.NewUint64(tableName, "id")
	_publicationORM.PostIds = field.NewField(tableName, "post_ids")
	_publicationORM.PublicationName = field.NewString(tableName, "publication_name")
	_publicationORM.Subjects = field.NewField(tableName, "subjects")
	_publicationORM.Tags = field.NewField(tableName, "tags")
	_publicationORM.Type = field.NewString(tableName, "type")
	_publicationORM.Admin = publicationORMHasOneAdmin{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Admin", "social_servicev2.UserProfileORM"),
		Bookmarks: struct {
			field.RelationField
			Publications struct {
				field.RelationField
				Admin struct {
					field.RelationField
				}
				Editors struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Admin.Bookmarks", "social_servicev2.BookmarkORM"),
			Publications: struct {
				field.RelationField
				Admin struct {
					field.RelationField
				}
				Editors struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Admin.Bookmarks.Publications", "social_servicev2.PublicationORM"),
				Admin: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Admin.Bookmarks.Publications.Admin", "social_servicev2.UserProfileORM"),
				},
				Editors: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Admin.Bookmarks.Publications.Editors", "social_servicev2.UserProfileORM"),
				},
			},
		},
		Tags: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Admin.Tags", "social_servicev2.UserTagsORM"),
		},
	}

	_publicationORM.Editors = publicationORMHasManyEditors{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Editors", "social_servicev2.UserProfileORM"),
	}

	_publicationORM.fillFieldMap()

	return _publicationORM
}

type publicationORM struct {
	publicationORMDo

	ALL                        field.Asterisk
	AdminBackendPlatformUserId field.String
	BookmarkId                 field.Uint64
	CreatedAt                  field.String
	Description                field.String
	Id                         field.Uint64
	PostIds                    field.Field
	PublicationName            field.String
	Subjects                   field.Field
	Tags                       field.Field
	Type                       field.String
	Admin                      publicationORMHasOneAdmin

	Editors publicationORMHasManyEditors

	fieldMap map[string]field.Expr
}

func (p publicationORM) Table(newTableName string) *publicationORM {
	p.publicationORMDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publicationORM) As(alias string) *publicationORM {
	p.publicationORMDo.DO = *(p.publicationORMDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publicationORM) updateTableName(table string) *publicationORM {
	p.ALL = field.NewAsterisk(table)
	p.AdminBackendPlatformUserId = field.NewString(table, "admin_backend_platform_user_id")
	p.BookmarkId = field.NewUint64(table, "bookmark_id")
	p.CreatedAt = field.NewString(table, "created_at")
	p.Description = field.NewString(table, "description")
	p.Id = field.NewUint64(table, "id")
	p.PostIds = field.NewField(table, "post_ids")
	p.PublicationName = field.NewString(table, "publication_name")
	p.Subjects = field.NewField(table, "subjects")
	p.Tags = field.NewField(table, "tags")
	p.Type = field.NewString(table, "type")

	p.fillFieldMap()

	return p
}

func (p *publicationORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publicationORM) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["admin_backend_platform_user_id"] = p.AdminBackendPlatformUserId
	p.fieldMap["bookmark_id"] = p.BookmarkId
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["description"] = p.Description
	p.fieldMap["id"] = p.Id
	p.fieldMap["post_ids"] = p.PostIds
	p.fieldMap["publication_name"] = p.PublicationName
	p.fieldMap["subjects"] = p.Subjects
	p.fieldMap["tags"] = p.Tags
	p.fieldMap["type"] = p.Type

}

func (p publicationORM) clone(db *gorm.DB) publicationORM {
	p.publicationORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publicationORM) replaceDB(db *gorm.DB) publicationORM {
	p.publicationORMDo.ReplaceDB(db)
	return p
}

type publicationORMHasOneAdmin struct {
	db *gorm.DB

	field.RelationField

	Bookmarks struct {
		field.RelationField
		Publications struct {
			field.RelationField
			Admin struct {
				field.RelationField
			}
			Editors struct {
				field.RelationField
			}
		}
	}
	Tags struct {
		field.RelationField
	}
}

func (a publicationORMHasOneAdmin) Where(conds ...field.Expr) *publicationORMHasOneAdmin {
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

func (a publicationORMHasOneAdmin) WithContext(ctx context.Context) *publicationORMHasOneAdmin {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a publicationORMHasOneAdmin) Session(session *gorm.Session) *publicationORMHasOneAdmin {
	a.db = a.db.Session(session)
	return &a
}

func (a publicationORMHasOneAdmin) Model(m *social_servicev2.PublicationORM) *publicationORMHasOneAdminTx {
	return &publicationORMHasOneAdminTx{a.db.Model(m).Association(a.Name())}
}

type publicationORMHasOneAdminTx struct{ tx *gorm.Association }

func (a publicationORMHasOneAdminTx) Find() (result *social_servicev2.UserProfileORM, err error) {
	return result, a.tx.Find(&result)
}

func (a publicationORMHasOneAdminTx) Append(values ...*social_servicev2.UserProfileORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a publicationORMHasOneAdminTx) Replace(values ...*social_servicev2.UserProfileORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a publicationORMHasOneAdminTx) Delete(values ...*social_servicev2.UserProfileORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a publicationORMHasOneAdminTx) Clear() error {
	return a.tx.Clear()
}

func (a publicationORMHasOneAdminTx) Count() int64 {
	return a.tx.Count()
}

type publicationORMHasManyEditors struct {
	db *gorm.DB

	field.RelationField
}

func (a publicationORMHasManyEditors) Where(conds ...field.Expr) *publicationORMHasManyEditors {
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

func (a publicationORMHasManyEditors) WithContext(ctx context.Context) *publicationORMHasManyEditors {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a publicationORMHasManyEditors) Session(session *gorm.Session) *publicationORMHasManyEditors {
	a.db = a.db.Session(session)
	return &a
}

func (a publicationORMHasManyEditors) Model(m *social_servicev2.PublicationORM) *publicationORMHasManyEditorsTx {
	return &publicationORMHasManyEditorsTx{a.db.Model(m).Association(a.Name())}
}

type publicationORMHasManyEditorsTx struct{ tx *gorm.Association }

func (a publicationORMHasManyEditorsTx) Find() (result []*social_servicev2.UserProfileORM, err error) {
	return result, a.tx.Find(&result)
}

func (a publicationORMHasManyEditorsTx) Append(values ...*social_servicev2.UserProfileORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a publicationORMHasManyEditorsTx) Replace(values ...*social_servicev2.UserProfileORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a publicationORMHasManyEditorsTx) Delete(values ...*social_servicev2.UserProfileORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a publicationORMHasManyEditorsTx) Clear() error {
	return a.tx.Clear()
}

func (a publicationORMHasManyEditorsTx) Count() int64 {
	return a.tx.Count()
}

type publicationORMDo struct{ gen.DO }

type IPublicationORMDo interface {
	gen.SubQuery
	Debug() IPublicationORMDo
	WithContext(ctx context.Context) IPublicationORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublicationORMDo
	WriteDB() IPublicationORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublicationORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublicationORMDo
	Not(conds ...gen.Condition) IPublicationORMDo
	Or(conds ...gen.Condition) IPublicationORMDo
	Select(conds ...field.Expr) IPublicationORMDo
	Where(conds ...gen.Condition) IPublicationORMDo
	Order(conds ...field.Expr) IPublicationORMDo
	Distinct(cols ...field.Expr) IPublicationORMDo
	Omit(cols ...field.Expr) IPublicationORMDo
	Join(table schema.Tabler, on ...field.Expr) IPublicationORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublicationORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublicationORMDo
	Group(cols ...field.Expr) IPublicationORMDo
	Having(conds ...gen.Condition) IPublicationORMDo
	Limit(limit int) IPublicationORMDo
	Offset(offset int) IPublicationORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublicationORMDo
	Unscoped() IPublicationORMDo
	Create(values ...*social_servicev2.PublicationORM) error
	CreateInBatches(values []*social_servicev2.PublicationORM, batchSize int) error
	Save(values ...*social_servicev2.PublicationORM) error
	First() (*social_servicev2.PublicationORM, error)
	Take() (*social_servicev2.PublicationORM, error)
	Last() (*social_servicev2.PublicationORM, error)
	Find() ([]*social_servicev2.PublicationORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*social_servicev2.PublicationORM, err error)
	FindInBatches(result *[]*social_servicev2.PublicationORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*social_servicev2.PublicationORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublicationORMDo
	Assign(attrs ...field.AssignExpr) IPublicationORMDo
	Joins(fields ...field.RelationField) IPublicationORMDo
	Preload(fields ...field.RelationField) IPublicationORMDo
	FirstOrInit() (*social_servicev2.PublicationORM, error)
	FirstOrCreate() (*social_servicev2.PublicationORM, error)
	FindByPage(offset int, limit int) (result []*social_servicev2.PublicationORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublicationORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result social_servicev2.PublicationORM, err error)
	GetRecordByIDs(ids []int) (result []social_servicev2.PublicationORM, err error)
	CreateRecord(item social_servicev2.PublicationORM) (err error)
	UpdateRecordByID(id int, item social_servicev2.PublicationORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []social_servicev2.PublicationORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result social_servicev2.PublicationORM, err error)
	GetByIDs(ids []uint64) (result []social_servicev2.PublicationORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (p publicationORMDo) GetRecordByID(id int) (result social_servicev2.PublicationORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM publications ")
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
func (p publicationORMDo) GetRecordByIDs(ids []int) (result []social_servicev2.PublicationORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM publications ")
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
func (p publicationORMDo) CreateRecord(item social_servicev2.PublicationORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO publications (columns) VALUES (values) ")

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
func (p publicationORMDo) UpdateRecordByID(id int, item social_servicev2.PublicationORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE publications SET columns=values ")
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
func (p publicationORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM publications ")
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
func (p publicationORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []social_servicev2.PublicationORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM publications ORDER BY " + p.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (p publicationORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM publications ")

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
func (p publicationORMDo) GetByID(id uint64) (result social_servicev2.PublicationORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM publications ")
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
func (p publicationORMDo) GetByIDs(ids []uint64) (result []social_servicev2.PublicationORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM publications ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p publicationORMDo) Debug() IPublicationORMDo {
	return p.withDO(p.DO.Debug())
}

func (p publicationORMDo) WithContext(ctx context.Context) IPublicationORMDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publicationORMDo) ReadDB() IPublicationORMDo {
	return p.Clauses(dbresolver.Read)
}

func (p publicationORMDo) WriteDB() IPublicationORMDo {
	return p.Clauses(dbresolver.Write)
}

func (p publicationORMDo) Session(config *gorm.Session) IPublicationORMDo {
	return p.withDO(p.DO.Session(config))
}

func (p publicationORMDo) Clauses(conds ...clause.Expression) IPublicationORMDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publicationORMDo) Returning(value interface{}, columns ...string) IPublicationORMDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publicationORMDo) Not(conds ...gen.Condition) IPublicationORMDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publicationORMDo) Or(conds ...gen.Condition) IPublicationORMDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publicationORMDo) Select(conds ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publicationORMDo) Where(conds ...gen.Condition) IPublicationORMDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publicationORMDo) Order(conds ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publicationORMDo) Distinct(cols ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publicationORMDo) Omit(cols ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publicationORMDo) Join(table schema.Tabler, on ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publicationORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publicationORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publicationORMDo) Group(cols ...field.Expr) IPublicationORMDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publicationORMDo) Having(conds ...gen.Condition) IPublicationORMDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publicationORMDo) Limit(limit int) IPublicationORMDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publicationORMDo) Offset(offset int) IPublicationORMDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publicationORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublicationORMDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publicationORMDo) Unscoped() IPublicationORMDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publicationORMDo) Create(values ...*social_servicev2.PublicationORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publicationORMDo) CreateInBatches(values []*social_servicev2.PublicationORM, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publicationORMDo) Save(values ...*social_servicev2.PublicationORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publicationORMDo) First() (*social_servicev2.PublicationORM, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*social_servicev2.PublicationORM), nil
	}
}

func (p publicationORMDo) Take() (*social_servicev2.PublicationORM, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*social_servicev2.PublicationORM), nil
	}
}

func (p publicationORMDo) Last() (*social_servicev2.PublicationORM, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*social_servicev2.PublicationORM), nil
	}
}

func (p publicationORMDo) Find() ([]*social_servicev2.PublicationORM, error) {
	result, err := p.DO.Find()
	return result.([]*social_servicev2.PublicationORM), err
}

func (p publicationORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*social_servicev2.PublicationORM, err error) {
	buf := make([]*social_servicev2.PublicationORM, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publicationORMDo) FindInBatches(result *[]*social_servicev2.PublicationORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publicationORMDo) Attrs(attrs ...field.AssignExpr) IPublicationORMDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publicationORMDo) Assign(attrs ...field.AssignExpr) IPublicationORMDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publicationORMDo) Joins(fields ...field.RelationField) IPublicationORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publicationORMDo) Preload(fields ...field.RelationField) IPublicationORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publicationORMDo) FirstOrInit() (*social_servicev2.PublicationORM, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*social_servicev2.PublicationORM), nil
	}
}

func (p publicationORMDo) FirstOrCreate() (*social_servicev2.PublicationORM, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*social_servicev2.PublicationORM), nil
	}
}

func (p publicationORMDo) FindByPage(offset int, limit int) (result []*social_servicev2.PublicationORM, count int64, err error) {
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

func (p publicationORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publicationORMDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publicationORMDo) Delete(models ...*social_servicev2.PublicationORM) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publicationORMDo) withDO(do gen.Dao) *publicationORMDo {
	p.DO = *do.(*gen.DO)
	return p
}
