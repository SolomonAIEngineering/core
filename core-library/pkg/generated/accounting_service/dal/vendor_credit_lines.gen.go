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

func newVendorCreditLineORM(db *gorm.DB, opts ...gen.DOOption) vendorCreditLineORM {
	_vendorCreditLineORM := vendorCreditLineORM{}

	_vendorCreditLineORM.vendorCreditLineORMDo.UseDB(db, opts...)
	_vendorCreditLineORM.vendorCreditLineORMDo.UseModel(&accounting_servicev1.VendorCreditLineORM{})

	tableName := _vendorCreditLineORM.vendorCreditLineORMDo.TableName()
	_vendorCreditLineORM.ALL = field.NewAsterisk(tableName)
	_vendorCreditLineORM.Account = field.NewString(tableName, "account")
	_vendorCreditLineORM.Company = field.NewString(tableName, "company")
	_vendorCreditLineORM.Description = field.NewString(tableName, "description")
	_vendorCreditLineORM.ExchangeRate = field.NewString(tableName, "exchange_rate")
	_vendorCreditLineORM.Id = field.NewUint64(tableName, "id")
	_vendorCreditLineORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_vendorCreditLineORM.NetAmount = field.NewFloat64(tableName, "net_amount")
	_vendorCreditLineORM.RemoteId = field.NewString(tableName, "remote_id")
	_vendorCreditLineORM.TrackingCategories = field.NewField(tableName, "tracking_categories")
	_vendorCreditLineORM.TrackingCategory = field.NewString(tableName, "tracking_category")
	_vendorCreditLineORM.VendorCreditId = field.NewUint64(tableName, "vendor_credit_id")

	_vendorCreditLineORM.fillFieldMap()

	return _vendorCreditLineORM
}

type vendorCreditLineORM struct {
	vendorCreditLineORMDo

	ALL                field.Asterisk
	Account            field.String
	Company            field.String
	Description        field.String
	ExchangeRate       field.String
	Id                 field.Uint64
	ModifiedAt         field.Time
	NetAmount          field.Float64
	RemoteId           field.String
	TrackingCategories field.Field
	TrackingCategory   field.String
	VendorCreditId     field.Uint64

	fieldMap map[string]field.Expr
}

func (v vendorCreditLineORM) Table(newTableName string) *vendorCreditLineORM {
	v.vendorCreditLineORMDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vendorCreditLineORM) As(alias string) *vendorCreditLineORM {
	v.vendorCreditLineORMDo.DO = *(v.vendorCreditLineORMDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vendorCreditLineORM) updateTableName(table string) *vendorCreditLineORM {
	v.ALL = field.NewAsterisk(table)
	v.Account = field.NewString(table, "account")
	v.Company = field.NewString(table, "company")
	v.Description = field.NewString(table, "description")
	v.ExchangeRate = field.NewString(table, "exchange_rate")
	v.Id = field.NewUint64(table, "id")
	v.ModifiedAt = field.NewTime(table, "modified_at")
	v.NetAmount = field.NewFloat64(table, "net_amount")
	v.RemoteId = field.NewString(table, "remote_id")
	v.TrackingCategories = field.NewField(table, "tracking_categories")
	v.TrackingCategory = field.NewString(table, "tracking_category")
	v.VendorCreditId = field.NewUint64(table, "vendor_credit_id")

	v.fillFieldMap()

	return v
}

func (v *vendorCreditLineORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vendorCreditLineORM) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 11)
	v.fieldMap["account"] = v.Account
	v.fieldMap["company"] = v.Company
	v.fieldMap["description"] = v.Description
	v.fieldMap["exchange_rate"] = v.ExchangeRate
	v.fieldMap["id"] = v.Id
	v.fieldMap["modified_at"] = v.ModifiedAt
	v.fieldMap["net_amount"] = v.NetAmount
	v.fieldMap["remote_id"] = v.RemoteId
	v.fieldMap["tracking_categories"] = v.TrackingCategories
	v.fieldMap["tracking_category"] = v.TrackingCategory
	v.fieldMap["vendor_credit_id"] = v.VendorCreditId
}

func (v vendorCreditLineORM) clone(db *gorm.DB) vendorCreditLineORM {
	v.vendorCreditLineORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vendorCreditLineORM) replaceDB(db *gorm.DB) vendorCreditLineORM {
	v.vendorCreditLineORMDo.ReplaceDB(db)
	return v
}

type vendorCreditLineORMDo struct{ gen.DO }

type IVendorCreditLineORMDo interface {
	gen.SubQuery
	Debug() IVendorCreditLineORMDo
	WithContext(ctx context.Context) IVendorCreditLineORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVendorCreditLineORMDo
	WriteDB() IVendorCreditLineORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVendorCreditLineORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVendorCreditLineORMDo
	Not(conds ...gen.Condition) IVendorCreditLineORMDo
	Or(conds ...gen.Condition) IVendorCreditLineORMDo
	Select(conds ...field.Expr) IVendorCreditLineORMDo
	Where(conds ...gen.Condition) IVendorCreditLineORMDo
	Order(conds ...field.Expr) IVendorCreditLineORMDo
	Distinct(cols ...field.Expr) IVendorCreditLineORMDo
	Omit(cols ...field.Expr) IVendorCreditLineORMDo
	Join(table schema.Tabler, on ...field.Expr) IVendorCreditLineORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVendorCreditLineORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVendorCreditLineORMDo
	Group(cols ...field.Expr) IVendorCreditLineORMDo
	Having(conds ...gen.Condition) IVendorCreditLineORMDo
	Limit(limit int) IVendorCreditLineORMDo
	Offset(offset int) IVendorCreditLineORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVendorCreditLineORMDo
	Unscoped() IVendorCreditLineORMDo
	Create(values ...*accounting_servicev1.VendorCreditLineORM) error
	CreateInBatches(values []*accounting_servicev1.VendorCreditLineORM, batchSize int) error
	Save(values ...*accounting_servicev1.VendorCreditLineORM) error
	First() (*accounting_servicev1.VendorCreditLineORM, error)
	Take() (*accounting_servicev1.VendorCreditLineORM, error)
	Last() (*accounting_servicev1.VendorCreditLineORM, error)
	Find() ([]*accounting_servicev1.VendorCreditLineORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.VendorCreditLineORM, err error)
	FindInBatches(result *[]*accounting_servicev1.VendorCreditLineORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.VendorCreditLineORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVendorCreditLineORMDo
	Assign(attrs ...field.AssignExpr) IVendorCreditLineORMDo
	Joins(fields ...field.RelationField) IVendorCreditLineORMDo
	Preload(fields ...field.RelationField) IVendorCreditLineORMDo
	FirstOrInit() (*accounting_servicev1.VendorCreditLineORM, error)
	FirstOrCreate() (*accounting_servicev1.VendorCreditLineORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.VendorCreditLineORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVendorCreditLineORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.VendorCreditLineORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.VendorCreditLineORM, err error)
	CreateRecord(item accounting_servicev1.VendorCreditLineORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.VendorCreditLineORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.VendorCreditLineORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.VendorCreditLineORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.VendorCreditLineORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (v vendorCreditLineORMDo) GetRecordByID(id int) (result accounting_servicev1.VendorCreditLineORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM vendor_credit_lines ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (v vendorCreditLineORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.VendorCreditLineORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM vendor_credit_lines ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (v vendorCreditLineORMDo) CreateRecord(item accounting_servicev1.VendorCreditLineORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO vendor_credit_lines (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (v vendorCreditLineORMDo) UpdateRecordByID(id int, item accounting_servicev1.VendorCreditLineORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE vendor_credit_lines SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (v vendorCreditLineORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM vendor_credit_lines ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (v vendorCreditLineORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.VendorCreditLineORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM vendor_credit_lines ORDER BY " + v.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (v vendorCreditLineORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM vendor_credit_lines ")

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (v vendorCreditLineORMDo) GetByID(id uint64) (result accounting_servicev1.VendorCreditLineORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM vendor_credit_lines ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (v vendorCreditLineORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.VendorCreditLineORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM vendor_credit_lines ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (v vendorCreditLineORMDo) Debug() IVendorCreditLineORMDo {
	return v.withDO(v.DO.Debug())
}

func (v vendorCreditLineORMDo) WithContext(ctx context.Context) IVendorCreditLineORMDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vendorCreditLineORMDo) ReadDB() IVendorCreditLineORMDo {
	return v.Clauses(dbresolver.Read)
}

func (v vendorCreditLineORMDo) WriteDB() IVendorCreditLineORMDo {
	return v.Clauses(dbresolver.Write)
}

func (v vendorCreditLineORMDo) Session(config *gorm.Session) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Session(config))
}

func (v vendorCreditLineORMDo) Clauses(conds ...clause.Expression) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vendorCreditLineORMDo) Returning(value interface{}, columns ...string) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vendorCreditLineORMDo) Not(conds ...gen.Condition) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vendorCreditLineORMDo) Or(conds ...gen.Condition) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vendorCreditLineORMDo) Select(conds ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vendorCreditLineORMDo) Where(conds ...gen.Condition) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vendorCreditLineORMDo) Order(conds ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vendorCreditLineORMDo) Distinct(cols ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vendorCreditLineORMDo) Omit(cols ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vendorCreditLineORMDo) Join(table schema.Tabler, on ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vendorCreditLineORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vendorCreditLineORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vendorCreditLineORMDo) Group(cols ...field.Expr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vendorCreditLineORMDo) Having(conds ...gen.Condition) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vendorCreditLineORMDo) Limit(limit int) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vendorCreditLineORMDo) Offset(offset int) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vendorCreditLineORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vendorCreditLineORMDo) Unscoped() IVendorCreditLineORMDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vendorCreditLineORMDo) Create(values ...*accounting_servicev1.VendorCreditLineORM) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vendorCreditLineORMDo) CreateInBatches(values []*accounting_servicev1.VendorCreditLineORM, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vendorCreditLineORMDo) Save(values ...*accounting_servicev1.VendorCreditLineORM) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vendorCreditLineORMDo) First() (*accounting_servicev1.VendorCreditLineORM, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.VendorCreditLineORM), nil
	}
}

func (v vendorCreditLineORMDo) Take() (*accounting_servicev1.VendorCreditLineORM, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.VendorCreditLineORM), nil
	}
}

func (v vendorCreditLineORMDo) Last() (*accounting_servicev1.VendorCreditLineORM, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.VendorCreditLineORM), nil
	}
}

func (v vendorCreditLineORMDo) Find() ([]*accounting_servicev1.VendorCreditLineORM, error) {
	result, err := v.DO.Find()
	return result.([]*accounting_servicev1.VendorCreditLineORM), err
}

func (v vendorCreditLineORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.VendorCreditLineORM, err error) {
	buf := make([]*accounting_servicev1.VendorCreditLineORM, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vendorCreditLineORMDo) FindInBatches(result *[]*accounting_servicev1.VendorCreditLineORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vendorCreditLineORMDo) Attrs(attrs ...field.AssignExpr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vendorCreditLineORMDo) Assign(attrs ...field.AssignExpr) IVendorCreditLineORMDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vendorCreditLineORMDo) Joins(fields ...field.RelationField) IVendorCreditLineORMDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vendorCreditLineORMDo) Preload(fields ...field.RelationField) IVendorCreditLineORMDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vendorCreditLineORMDo) FirstOrInit() (*accounting_servicev1.VendorCreditLineORM, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.VendorCreditLineORM), nil
	}
}

func (v vendorCreditLineORMDo) FirstOrCreate() (*accounting_servicev1.VendorCreditLineORM, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.VendorCreditLineORM), nil
	}
}

func (v vendorCreditLineORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.VendorCreditLineORM, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vendorCreditLineORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vendorCreditLineORMDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vendorCreditLineORMDo) Delete(models ...*accounting_servicev1.VendorCreditLineORM) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vendorCreditLineORMDo) withDO(do gen.Dao) *vendorCreditLineORMDo {
	v.DO = *do.(*gen.DO)
	return v
}
