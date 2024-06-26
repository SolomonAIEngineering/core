// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	workspace_servicev1 "github.com/SolomonAIEngineering/core/core-library/pkg/generated/workspace_service/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newFolderMetadataORM(db *gorm.DB, opts ...gen.DOOption) folderMetadataORM {
	_folderMetadataORM := folderMetadataORM{}

	_folderMetadataORM.folderMetadataORMDo.UseDB(db, opts...)
	_folderMetadataORM.folderMetadataORMDo.UseModel(&workspace_servicev1.FolderMetadataORM{})

	tableName := _folderMetadataORM.folderMetadataORMDo.TableName()
	_folderMetadataORM.ALL = field.NewAsterisk(tableName)
	_folderMetadataORM.CreatedAt = field.NewTime(tableName, "created_at")
	_folderMetadataORM.FolderMetadataId = field.NewUint64(tableName, "folder_metadata_id")
	_folderMetadataORM.Id = field.NewUint64(tableName, "id")
	_folderMetadataORM.IsDeleted = field.NewBool(tableName, "is_deleted")
	_folderMetadataORM.Name = field.NewString(tableName, "name")
	_folderMetadataORM.S3Acl = field.NewString(tableName, "s3_acl")
	_folderMetadataORM.S3BucketName = field.NewString(tableName, "s3_bucket_name")
	_folderMetadataORM.S3FolderPath = field.NewString(tableName, "s3_folder_path")
	_folderMetadataORM.S3LastModified = field.NewTime(tableName, "s3_last_modified")
	_folderMetadataORM.S3Region = field.NewString(tableName, "s3_region")
	_folderMetadataORM.UpdatedAt = field.NewTime(tableName, "updated_at")
	_folderMetadataORM.VersionId = field.NewString(tableName, "version_id")
	_folderMetadataORM.WorkspaceId = field.NewUint64(tableName, "workspace_id")
	_folderMetadataORM.ChildFolder = folderMetadataORMHasManyChildFolder{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ChildFolder", "workspace_servicev1.FolderMetadataORM"),
		ChildFolder: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("ChildFolder.ChildFolder", "workspace_servicev1.FolderMetadataORM"),
		},
		Files: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("ChildFolder.Files", "workspace_servicev1.FileMetadataORM"),
		},
	}

	_folderMetadataORM.Files = folderMetadataORMHasManyFiles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Files", "workspace_servicev1.FileMetadataORM"),
	}

	_folderMetadataORM.fillFieldMap()

	return _folderMetadataORM
}

type folderMetadataORM struct {
	folderMetadataORMDo

	ALL              field.Asterisk
	CreatedAt        field.Time
	FolderMetadataId field.Uint64
	Id               field.Uint64
	IsDeleted        field.Bool
	Name             field.String
	S3Acl            field.String
	S3BucketName     field.String
	S3FolderPath     field.String
	S3LastModified   field.Time
	S3Region         field.String
	UpdatedAt        field.Time
	VersionId        field.String
	WorkspaceId      field.Uint64
	ChildFolder      folderMetadataORMHasManyChildFolder

	Files folderMetadataORMHasManyFiles

	fieldMap map[string]field.Expr
}

func (f folderMetadataORM) Table(newTableName string) *folderMetadataORM {
	f.folderMetadataORMDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f folderMetadataORM) As(alias string) *folderMetadataORM {
	f.folderMetadataORMDo.DO = *(f.folderMetadataORMDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *folderMetadataORM) updateTableName(table string) *folderMetadataORM {
	f.ALL = field.NewAsterisk(table)
	f.CreatedAt = field.NewTime(table, "created_at")
	f.FolderMetadataId = field.NewUint64(table, "folder_metadata_id")
	f.Id = field.NewUint64(table, "id")
	f.IsDeleted = field.NewBool(table, "is_deleted")
	f.Name = field.NewString(table, "name")
	f.S3Acl = field.NewString(table, "s3_acl")
	f.S3BucketName = field.NewString(table, "s3_bucket_name")
	f.S3FolderPath = field.NewString(table, "s3_folder_path")
	f.S3LastModified = field.NewTime(table, "s3_last_modified")
	f.S3Region = field.NewString(table, "s3_region")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.VersionId = field.NewString(table, "version_id")
	f.WorkspaceId = field.NewUint64(table, "workspace_id")

	f.fillFieldMap()

	return f
}

func (f *folderMetadataORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *folderMetadataORM) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 15)
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["folder_metadata_id"] = f.FolderMetadataId
	f.fieldMap["id"] = f.Id
	f.fieldMap["is_deleted"] = f.IsDeleted
	f.fieldMap["name"] = f.Name
	f.fieldMap["s3_acl"] = f.S3Acl
	f.fieldMap["s3_bucket_name"] = f.S3BucketName
	f.fieldMap["s3_folder_path"] = f.S3FolderPath
	f.fieldMap["s3_last_modified"] = f.S3LastModified
	f.fieldMap["s3_region"] = f.S3Region
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["version_id"] = f.VersionId
	f.fieldMap["workspace_id"] = f.WorkspaceId

}

func (f folderMetadataORM) clone(db *gorm.DB) folderMetadataORM {
	f.folderMetadataORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f folderMetadataORM) replaceDB(db *gorm.DB) folderMetadataORM {
	f.folderMetadataORMDo.ReplaceDB(db)
	return f
}

type folderMetadataORMHasManyChildFolder struct {
	db *gorm.DB

	field.RelationField

	ChildFolder struct {
		field.RelationField
	}
	Files struct {
		field.RelationField
	}
}

func (a folderMetadataORMHasManyChildFolder) Where(conds ...field.Expr) *folderMetadataORMHasManyChildFolder {
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

func (a folderMetadataORMHasManyChildFolder) WithContext(ctx context.Context) *folderMetadataORMHasManyChildFolder {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a folderMetadataORMHasManyChildFolder) Session(session *gorm.Session) *folderMetadataORMHasManyChildFolder {
	a.db = a.db.Session(session)
	return &a
}

func (a folderMetadataORMHasManyChildFolder) Model(m *workspace_servicev1.FolderMetadataORM) *folderMetadataORMHasManyChildFolderTx {
	return &folderMetadataORMHasManyChildFolderTx{a.db.Model(m).Association(a.Name())}
}

type folderMetadataORMHasManyChildFolderTx struct{ tx *gorm.Association }

func (a folderMetadataORMHasManyChildFolderTx) Find() (result []*workspace_servicev1.FolderMetadataORM, err error) {
	return result, a.tx.Find(&result)
}

func (a folderMetadataORMHasManyChildFolderTx) Append(values ...*workspace_servicev1.FolderMetadataORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a folderMetadataORMHasManyChildFolderTx) Replace(values ...*workspace_servicev1.FolderMetadataORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a folderMetadataORMHasManyChildFolderTx) Delete(values ...*workspace_servicev1.FolderMetadataORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a folderMetadataORMHasManyChildFolderTx) Clear() error {
	return a.tx.Clear()
}

func (a folderMetadataORMHasManyChildFolderTx) Count() int64 {
	return a.tx.Count()
}

type folderMetadataORMHasManyFiles struct {
	db *gorm.DB

	field.RelationField
}

func (a folderMetadataORMHasManyFiles) Where(conds ...field.Expr) *folderMetadataORMHasManyFiles {
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

func (a folderMetadataORMHasManyFiles) WithContext(ctx context.Context) *folderMetadataORMHasManyFiles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a folderMetadataORMHasManyFiles) Session(session *gorm.Session) *folderMetadataORMHasManyFiles {
	a.db = a.db.Session(session)
	return &a
}

func (a folderMetadataORMHasManyFiles) Model(m *workspace_servicev1.FolderMetadataORM) *folderMetadataORMHasManyFilesTx {
	return &folderMetadataORMHasManyFilesTx{a.db.Model(m).Association(a.Name())}
}

type folderMetadataORMHasManyFilesTx struct{ tx *gorm.Association }

func (a folderMetadataORMHasManyFilesTx) Find() (result []*workspace_servicev1.FileMetadataORM, err error) {
	return result, a.tx.Find(&result)
}

func (a folderMetadataORMHasManyFilesTx) Append(values ...*workspace_servicev1.FileMetadataORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a folderMetadataORMHasManyFilesTx) Replace(values ...*workspace_servicev1.FileMetadataORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a folderMetadataORMHasManyFilesTx) Delete(values ...*workspace_servicev1.FileMetadataORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a folderMetadataORMHasManyFilesTx) Clear() error {
	return a.tx.Clear()
}

func (a folderMetadataORMHasManyFilesTx) Count() int64 {
	return a.tx.Count()
}

type folderMetadataORMDo struct{ gen.DO }

type IFolderMetadataORMDo interface {
	gen.SubQuery
	Debug() IFolderMetadataORMDo
	WithContext(ctx context.Context) IFolderMetadataORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFolderMetadataORMDo
	WriteDB() IFolderMetadataORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFolderMetadataORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFolderMetadataORMDo
	Not(conds ...gen.Condition) IFolderMetadataORMDo
	Or(conds ...gen.Condition) IFolderMetadataORMDo
	Select(conds ...field.Expr) IFolderMetadataORMDo
	Where(conds ...gen.Condition) IFolderMetadataORMDo
	Order(conds ...field.Expr) IFolderMetadataORMDo
	Distinct(cols ...field.Expr) IFolderMetadataORMDo
	Omit(cols ...field.Expr) IFolderMetadataORMDo
	Join(table schema.Tabler, on ...field.Expr) IFolderMetadataORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFolderMetadataORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFolderMetadataORMDo
	Group(cols ...field.Expr) IFolderMetadataORMDo
	Having(conds ...gen.Condition) IFolderMetadataORMDo
	Limit(limit int) IFolderMetadataORMDo
	Offset(offset int) IFolderMetadataORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFolderMetadataORMDo
	Unscoped() IFolderMetadataORMDo
	Create(values ...*workspace_servicev1.FolderMetadataORM) error
	CreateInBatches(values []*workspace_servicev1.FolderMetadataORM, batchSize int) error
	Save(values ...*workspace_servicev1.FolderMetadataORM) error
	First() (*workspace_servicev1.FolderMetadataORM, error)
	Take() (*workspace_servicev1.FolderMetadataORM, error)
	Last() (*workspace_servicev1.FolderMetadataORM, error)
	Find() ([]*workspace_servicev1.FolderMetadataORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*workspace_servicev1.FolderMetadataORM, err error)
	FindInBatches(result *[]*workspace_servicev1.FolderMetadataORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*workspace_servicev1.FolderMetadataORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFolderMetadataORMDo
	Assign(attrs ...field.AssignExpr) IFolderMetadataORMDo
	Joins(fields ...field.RelationField) IFolderMetadataORMDo
	Preload(fields ...field.RelationField) IFolderMetadataORMDo
	FirstOrInit() (*workspace_servicev1.FolderMetadataORM, error)
	FirstOrCreate() (*workspace_servicev1.FolderMetadataORM, error)
	FindByPage(offset int, limit int) (result []*workspace_servicev1.FolderMetadataORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFolderMetadataORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result workspace_servicev1.FolderMetadataORM, err error)
	GetRecordByIDs(ids []int) (result []workspace_servicev1.FolderMetadataORM, err error)
	CreateRecord(item workspace_servicev1.FolderMetadataORM) (err error)
	UpdateRecordByID(id int, item workspace_servicev1.FolderMetadataORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []workspace_servicev1.FolderMetadataORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result workspace_servicev1.FolderMetadataORM, err error)
	GetByIDs(ids []uint64) (result []workspace_servicev1.FolderMetadataORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (f folderMetadataORMDo) GetRecordByID(id int) (result workspace_servicev1.FolderMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM folder_metadata ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (f folderMetadataORMDo) GetRecordByIDs(ids []int) (result []workspace_servicev1.FolderMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM folder_metadata ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (f folderMetadataORMDo) CreateRecord(item workspace_servicev1.FolderMetadataORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO folder_metadata (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (f folderMetadataORMDo) UpdateRecordByID(id int, item workspace_servicev1.FolderMetadataORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE folder_metadata SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (f folderMetadataORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM folder_metadata ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (f folderMetadataORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []workspace_servicev1.FolderMetadataORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM folder_metadata ORDER BY " + f.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (f folderMetadataORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM folder_metadata ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (f folderMetadataORMDo) GetByID(id uint64) (result workspace_servicev1.FolderMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM folder_metadata ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (f folderMetadataORMDo) GetByIDs(ids []uint64) (result []workspace_servicev1.FolderMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM folder_metadata ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (f folderMetadataORMDo) Debug() IFolderMetadataORMDo {
	return f.withDO(f.DO.Debug())
}

func (f folderMetadataORMDo) WithContext(ctx context.Context) IFolderMetadataORMDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f folderMetadataORMDo) ReadDB() IFolderMetadataORMDo {
	return f.Clauses(dbresolver.Read)
}

func (f folderMetadataORMDo) WriteDB() IFolderMetadataORMDo {
	return f.Clauses(dbresolver.Write)
}

func (f folderMetadataORMDo) Session(config *gorm.Session) IFolderMetadataORMDo {
	return f.withDO(f.DO.Session(config))
}

func (f folderMetadataORMDo) Clauses(conds ...clause.Expression) IFolderMetadataORMDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f folderMetadataORMDo) Returning(value interface{}, columns ...string) IFolderMetadataORMDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f folderMetadataORMDo) Not(conds ...gen.Condition) IFolderMetadataORMDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f folderMetadataORMDo) Or(conds ...gen.Condition) IFolderMetadataORMDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f folderMetadataORMDo) Select(conds ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f folderMetadataORMDo) Where(conds ...gen.Condition) IFolderMetadataORMDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f folderMetadataORMDo) Order(conds ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f folderMetadataORMDo) Distinct(cols ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f folderMetadataORMDo) Omit(cols ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f folderMetadataORMDo) Join(table schema.Tabler, on ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f folderMetadataORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f folderMetadataORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f folderMetadataORMDo) Group(cols ...field.Expr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f folderMetadataORMDo) Having(conds ...gen.Condition) IFolderMetadataORMDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f folderMetadataORMDo) Limit(limit int) IFolderMetadataORMDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f folderMetadataORMDo) Offset(offset int) IFolderMetadataORMDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f folderMetadataORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFolderMetadataORMDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f folderMetadataORMDo) Unscoped() IFolderMetadataORMDo {
	return f.withDO(f.DO.Unscoped())
}

func (f folderMetadataORMDo) Create(values ...*workspace_servicev1.FolderMetadataORM) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f folderMetadataORMDo) CreateInBatches(values []*workspace_servicev1.FolderMetadataORM, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f folderMetadataORMDo) Save(values ...*workspace_servicev1.FolderMetadataORM) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f folderMetadataORMDo) First() (*workspace_servicev1.FolderMetadataORM, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FolderMetadataORM), nil
	}
}

func (f folderMetadataORMDo) Take() (*workspace_servicev1.FolderMetadataORM, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FolderMetadataORM), nil
	}
}

func (f folderMetadataORMDo) Last() (*workspace_servicev1.FolderMetadataORM, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FolderMetadataORM), nil
	}
}

func (f folderMetadataORMDo) Find() ([]*workspace_servicev1.FolderMetadataORM, error) {
	result, err := f.DO.Find()
	return result.([]*workspace_servicev1.FolderMetadataORM), err
}

func (f folderMetadataORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*workspace_servicev1.FolderMetadataORM, err error) {
	buf := make([]*workspace_servicev1.FolderMetadataORM, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f folderMetadataORMDo) FindInBatches(result *[]*workspace_servicev1.FolderMetadataORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f folderMetadataORMDo) Attrs(attrs ...field.AssignExpr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f folderMetadataORMDo) Assign(attrs ...field.AssignExpr) IFolderMetadataORMDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f folderMetadataORMDo) Joins(fields ...field.RelationField) IFolderMetadataORMDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f folderMetadataORMDo) Preload(fields ...field.RelationField) IFolderMetadataORMDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f folderMetadataORMDo) FirstOrInit() (*workspace_servicev1.FolderMetadataORM, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FolderMetadataORM), nil
	}
}

func (f folderMetadataORMDo) FirstOrCreate() (*workspace_servicev1.FolderMetadataORM, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FolderMetadataORM), nil
	}
}

func (f folderMetadataORMDo) FindByPage(offset int, limit int) (result []*workspace_servicev1.FolderMetadataORM, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f folderMetadataORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f folderMetadataORMDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f folderMetadataORMDo) Delete(models ...*workspace_servicev1.FolderMetadataORM) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *folderMetadataORMDo) withDO(do gen.Dao) *folderMetadataORMDo {
	f.DO = *do.(*gen.DO)
	return f
}
