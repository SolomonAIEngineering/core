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

func newFileMetadataORM(db *gorm.DB, opts ...gen.DOOption) fileMetadataORM {
	_fileMetadataORM := fileMetadataORM{}

	_fileMetadataORM.fileMetadataORMDo.UseDB(db, opts...)
	_fileMetadataORM.fileMetadataORMDo.UseModel(&workspace_servicev1.FileMetadataORM{})

	tableName := _fileMetadataORM.fileMetadataORMDo.TableName()
	_fileMetadataORM.ALL = field.NewAsterisk(tableName)
	_fileMetadataORM.CreatedAt = field.NewTime(tableName, "created_at")
	_fileMetadataORM.FileType = field.NewString(tableName, "file_type")
	_fileMetadataORM.FolderMetadataId = field.NewUint64(tableName, "folder_metadata_id")
	_fileMetadataORM.Id = field.NewUint64(tableName, "id")
	_fileMetadataORM.IsDeleted = field.NewBool(tableName, "is_deleted")
	_fileMetadataORM.Location = field.NewString(tableName, "location")
	_fileMetadataORM.MarkdownContent = field.NewString(tableName, "markdown_content")
	_fileMetadataORM.Name = field.NewString(tableName, "name")
	_fileMetadataORM.S3Acl = field.NewString(tableName, "s3_acl")
	_fileMetadataORM.S3BucketName = field.NewString(tableName, "s3_bucket_name")
	_fileMetadataORM.S3ContentDisposition = field.NewString(tableName, "s3_content_disposition")
	_fileMetadataORM.S3ContentEncoding = field.NewString(tableName, "s3_content_encoding")
	_fileMetadataORM.S3ContentLength = field.NewInt64(tableName, "s3_content_length")
	_fileMetadataORM.S3ContentType = field.NewString(tableName, "s3_content_type")
	_fileMetadataORM.S3Etag = field.NewString(tableName, "s3_etag")
	_fileMetadataORM.S3Key = field.NewString(tableName, "s3_key")
	_fileMetadataORM.S3LastModified = field.NewTime(tableName, "s3_last_modified")
	_fileMetadataORM.S3Region = field.NewString(tableName, "s3_region")
	_fileMetadataORM.S3ServerSideEncryption = field.NewString(tableName, "s3_server_side_encryption")
	_fileMetadataORM.S3StorageClass = field.NewString(tableName, "s3_storage_class")
	_fileMetadataORM.S3VersionId = field.NewString(tableName, "s3_version_id")
	_fileMetadataORM.Size = field.NewInt64(tableName, "size")
	_fileMetadataORM.Tags = field.NewField(tableName, "tags")
	_fileMetadataORM.UpdatedAt = field.NewTime(tableName, "updated_at")
	_fileMetadataORM.UploadId = field.NewString(tableName, "upload_id")
	_fileMetadataORM.Version = field.NewInt32(tableName, "version")
	_fileMetadataORM.VersionId = field.NewString(tableName, "version_id")

	_fileMetadataORM.fillFieldMap()

	return _fileMetadataORM
}

type fileMetadataORM struct {
	fileMetadataORMDo

	ALL                    field.Asterisk
	CreatedAt              field.Time
	FileType               field.String
	FolderMetadataId       field.Uint64
	Id                     field.Uint64
	IsDeleted              field.Bool
	Location               field.String
	MarkdownContent        field.String
	Name                   field.String
	S3Acl                  field.String
	S3BucketName           field.String
	S3ContentDisposition   field.String
	S3ContentEncoding      field.String
	S3ContentLength        field.Int64
	S3ContentType          field.String
	S3Etag                 field.String
	S3Key                  field.String
	S3LastModified         field.Time
	S3Region               field.String
	S3ServerSideEncryption field.String
	S3StorageClass         field.String
	S3VersionId            field.String
	Size                   field.Int64
	Tags                   field.Field
	UpdatedAt              field.Time
	UploadId               field.String
	Version                field.Int32
	VersionId              field.String

	fieldMap map[string]field.Expr
}

func (f fileMetadataORM) Table(newTableName string) *fileMetadataORM {
	f.fileMetadataORMDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fileMetadataORM) As(alias string) *fileMetadataORM {
	f.fileMetadataORMDo.DO = *(f.fileMetadataORMDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fileMetadataORM) updateTableName(table string) *fileMetadataORM {
	f.ALL = field.NewAsterisk(table)
	f.CreatedAt = field.NewTime(table, "created_at")
	f.FileType = field.NewString(table, "file_type")
	f.FolderMetadataId = field.NewUint64(table, "folder_metadata_id")
	f.Id = field.NewUint64(table, "id")
	f.IsDeleted = field.NewBool(table, "is_deleted")
	f.Location = field.NewString(table, "location")
	f.MarkdownContent = field.NewString(table, "markdown_content")
	f.Name = field.NewString(table, "name")
	f.S3Acl = field.NewString(table, "s3_acl")
	f.S3BucketName = field.NewString(table, "s3_bucket_name")
	f.S3ContentDisposition = field.NewString(table, "s3_content_disposition")
	f.S3ContentEncoding = field.NewString(table, "s3_content_encoding")
	f.S3ContentLength = field.NewInt64(table, "s3_content_length")
	f.S3ContentType = field.NewString(table, "s3_content_type")
	f.S3Etag = field.NewString(table, "s3_etag")
	f.S3Key = field.NewString(table, "s3_key")
	f.S3LastModified = field.NewTime(table, "s3_last_modified")
	f.S3Region = field.NewString(table, "s3_region")
	f.S3ServerSideEncryption = field.NewString(table, "s3_server_side_encryption")
	f.S3StorageClass = field.NewString(table, "s3_storage_class")
	f.S3VersionId = field.NewString(table, "s3_version_id")
	f.Size = field.NewInt64(table, "size")
	f.Tags = field.NewField(table, "tags")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.UploadId = field.NewString(table, "upload_id")
	f.Version = field.NewInt32(table, "version")
	f.VersionId = field.NewString(table, "version_id")

	f.fillFieldMap()

	return f
}

func (f *fileMetadataORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fileMetadataORM) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 27)
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["file_type"] = f.FileType
	f.fieldMap["folder_metadata_id"] = f.FolderMetadataId
	f.fieldMap["id"] = f.Id
	f.fieldMap["is_deleted"] = f.IsDeleted
	f.fieldMap["location"] = f.Location
	f.fieldMap["markdown_content"] = f.MarkdownContent
	f.fieldMap["name"] = f.Name
	f.fieldMap["s3_acl"] = f.S3Acl
	f.fieldMap["s3_bucket_name"] = f.S3BucketName
	f.fieldMap["s3_content_disposition"] = f.S3ContentDisposition
	f.fieldMap["s3_content_encoding"] = f.S3ContentEncoding
	f.fieldMap["s3_content_length"] = f.S3ContentLength
	f.fieldMap["s3_content_type"] = f.S3ContentType
	f.fieldMap["s3_etag"] = f.S3Etag
	f.fieldMap["s3_key"] = f.S3Key
	f.fieldMap["s3_last_modified"] = f.S3LastModified
	f.fieldMap["s3_region"] = f.S3Region
	f.fieldMap["s3_server_side_encryption"] = f.S3ServerSideEncryption
	f.fieldMap["s3_storage_class"] = f.S3StorageClass
	f.fieldMap["s3_version_id"] = f.S3VersionId
	f.fieldMap["size"] = f.Size
	f.fieldMap["tags"] = f.Tags
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["upload_id"] = f.UploadId
	f.fieldMap["version"] = f.Version
	f.fieldMap["version_id"] = f.VersionId
}

func (f fileMetadataORM) clone(db *gorm.DB) fileMetadataORM {
	f.fileMetadataORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fileMetadataORM) replaceDB(db *gorm.DB) fileMetadataORM {
	f.fileMetadataORMDo.ReplaceDB(db)
	return f
}

type fileMetadataORMDo struct{ gen.DO }

type IFileMetadataORMDo interface {
	gen.SubQuery
	Debug() IFileMetadataORMDo
	WithContext(ctx context.Context) IFileMetadataORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFileMetadataORMDo
	WriteDB() IFileMetadataORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFileMetadataORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFileMetadataORMDo
	Not(conds ...gen.Condition) IFileMetadataORMDo
	Or(conds ...gen.Condition) IFileMetadataORMDo
	Select(conds ...field.Expr) IFileMetadataORMDo
	Where(conds ...gen.Condition) IFileMetadataORMDo
	Order(conds ...field.Expr) IFileMetadataORMDo
	Distinct(cols ...field.Expr) IFileMetadataORMDo
	Omit(cols ...field.Expr) IFileMetadataORMDo
	Join(table schema.Tabler, on ...field.Expr) IFileMetadataORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFileMetadataORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFileMetadataORMDo
	Group(cols ...field.Expr) IFileMetadataORMDo
	Having(conds ...gen.Condition) IFileMetadataORMDo
	Limit(limit int) IFileMetadataORMDo
	Offset(offset int) IFileMetadataORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFileMetadataORMDo
	Unscoped() IFileMetadataORMDo
	Create(values ...*workspace_servicev1.FileMetadataORM) error
	CreateInBatches(values []*workspace_servicev1.FileMetadataORM, batchSize int) error
	Save(values ...*workspace_servicev1.FileMetadataORM) error
	First() (*workspace_servicev1.FileMetadataORM, error)
	Take() (*workspace_servicev1.FileMetadataORM, error)
	Last() (*workspace_servicev1.FileMetadataORM, error)
	Find() ([]*workspace_servicev1.FileMetadataORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*workspace_servicev1.FileMetadataORM, err error)
	FindInBatches(result *[]*workspace_servicev1.FileMetadataORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*workspace_servicev1.FileMetadataORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFileMetadataORMDo
	Assign(attrs ...field.AssignExpr) IFileMetadataORMDo
	Joins(fields ...field.RelationField) IFileMetadataORMDo
	Preload(fields ...field.RelationField) IFileMetadataORMDo
	FirstOrInit() (*workspace_servicev1.FileMetadataORM, error)
	FirstOrCreate() (*workspace_servicev1.FileMetadataORM, error)
	FindByPage(offset int, limit int) (result []*workspace_servicev1.FileMetadataORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFileMetadataORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result workspace_servicev1.FileMetadataORM, err error)
	GetRecordByIDs(ids []int) (result []workspace_servicev1.FileMetadataORM, err error)
	CreateRecord(item workspace_servicev1.FileMetadataORM) (err error)
	UpdateRecordByID(id int, item workspace_servicev1.FileMetadataORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []workspace_servicev1.FileMetadataORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result workspace_servicev1.FileMetadataORM, err error)
	GetByIDs(ids []uint64) (result []workspace_servicev1.FileMetadataORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (f fileMetadataORMDo) GetRecordByID(id int) (result workspace_servicev1.FileMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM file_metadata ")
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
func (f fileMetadataORMDo) GetRecordByIDs(ids []int) (result []workspace_servicev1.FileMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM file_metadata ")
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
func (f fileMetadataORMDo) CreateRecord(item workspace_servicev1.FileMetadataORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO file_metadata (columns) VALUES (values) ")

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
func (f fileMetadataORMDo) UpdateRecordByID(id int, item workspace_servicev1.FileMetadataORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE file_metadata SET columns=values ")
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
func (f fileMetadataORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM file_metadata ")
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
func (f fileMetadataORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []workspace_servicev1.FileMetadataORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM file_metadata ORDER BY " + f.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (f fileMetadataORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM file_metadata ")

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
func (f fileMetadataORMDo) GetByID(id uint64) (result workspace_servicev1.FileMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM file_metadata ")
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
func (f fileMetadataORMDo) GetByIDs(ids []uint64) (result []workspace_servicev1.FileMetadataORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM file_metadata ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (f fileMetadataORMDo) Debug() IFileMetadataORMDo {
	return f.withDO(f.DO.Debug())
}

func (f fileMetadataORMDo) WithContext(ctx context.Context) IFileMetadataORMDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileMetadataORMDo) ReadDB() IFileMetadataORMDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileMetadataORMDo) WriteDB() IFileMetadataORMDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileMetadataORMDo) Session(config *gorm.Session) IFileMetadataORMDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileMetadataORMDo) Clauses(conds ...clause.Expression) IFileMetadataORMDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileMetadataORMDo) Returning(value interface{}, columns ...string) IFileMetadataORMDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileMetadataORMDo) Not(conds ...gen.Condition) IFileMetadataORMDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileMetadataORMDo) Or(conds ...gen.Condition) IFileMetadataORMDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileMetadataORMDo) Select(conds ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileMetadataORMDo) Where(conds ...gen.Condition) IFileMetadataORMDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileMetadataORMDo) Order(conds ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileMetadataORMDo) Distinct(cols ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileMetadataORMDo) Omit(cols ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileMetadataORMDo) Join(table schema.Tabler, on ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileMetadataORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileMetadataORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileMetadataORMDo) Group(cols ...field.Expr) IFileMetadataORMDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileMetadataORMDo) Having(conds ...gen.Condition) IFileMetadataORMDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileMetadataORMDo) Limit(limit int) IFileMetadataORMDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileMetadataORMDo) Offset(offset int) IFileMetadataORMDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileMetadataORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFileMetadataORMDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileMetadataORMDo) Unscoped() IFileMetadataORMDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileMetadataORMDo) Create(values ...*workspace_servicev1.FileMetadataORM) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileMetadataORMDo) CreateInBatches(values []*workspace_servicev1.FileMetadataORM, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileMetadataORMDo) Save(values ...*workspace_servicev1.FileMetadataORM) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileMetadataORMDo) First() (*workspace_servicev1.FileMetadataORM, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FileMetadataORM), nil
	}
}

func (f fileMetadataORMDo) Take() (*workspace_servicev1.FileMetadataORM, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FileMetadataORM), nil
	}
}

func (f fileMetadataORMDo) Last() (*workspace_servicev1.FileMetadataORM, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FileMetadataORM), nil
	}
}

func (f fileMetadataORMDo) Find() ([]*workspace_servicev1.FileMetadataORM, error) {
	result, err := f.DO.Find()
	return result.([]*workspace_servicev1.FileMetadataORM), err
}

func (f fileMetadataORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*workspace_servicev1.FileMetadataORM, err error) {
	buf := make([]*workspace_servicev1.FileMetadataORM, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileMetadataORMDo) FindInBatches(result *[]*workspace_servicev1.FileMetadataORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileMetadataORMDo) Attrs(attrs ...field.AssignExpr) IFileMetadataORMDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileMetadataORMDo) Assign(attrs ...field.AssignExpr) IFileMetadataORMDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileMetadataORMDo) Joins(fields ...field.RelationField) IFileMetadataORMDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileMetadataORMDo) Preload(fields ...field.RelationField) IFileMetadataORMDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileMetadataORMDo) FirstOrInit() (*workspace_servicev1.FileMetadataORM, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FileMetadataORM), nil
	}
}

func (f fileMetadataORMDo) FirstOrCreate() (*workspace_servicev1.FileMetadataORM, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*workspace_servicev1.FileMetadataORM), nil
	}
}

func (f fileMetadataORMDo) FindByPage(offset int, limit int) (result []*workspace_servicev1.FileMetadataORM, count int64, err error) {
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

func (f fileMetadataORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileMetadataORMDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileMetadataORMDo) Delete(models ...*workspace_servicev1.FileMetadataORM) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileMetadataORMDo) withDO(do gen.Dao) *fileMetadataORMDo {
	f.DO = *do.(*gen.DO)
	return f
}
