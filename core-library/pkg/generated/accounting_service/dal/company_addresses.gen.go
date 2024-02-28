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

func newCompanyAddressORM(db *gorm.DB, opts ...gen.DOOption) companyAddressORM {
	_companyAddressORM := companyAddressORM{}

	_companyAddressORM.companyAddressORMDo.UseDB(db, opts...)
	_companyAddressORM.companyAddressORMDo.UseModel(&accounting_servicev1.CompanyAddressORM{})

	tableName := _companyAddressORM.companyAddressORMDo.TableName()
	_companyAddressORM.ALL = field.NewAsterisk(tableName)
	_companyAddressORM.City = field.NewString(tableName, "city")
	_companyAddressORM.CompanyInfoId = field.NewUint64(tableName, "company_info_id")
	_companyAddressORM.Country = field.NewString(tableName, "country")
	_companyAddressORM.CountrySubdivision = field.NewString(tableName, "country_subdivision")
	_companyAddressORM.Id = field.NewUint64(tableName, "id")
	_companyAddressORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_companyAddressORM.PurchaseOrderId = field.NewUint64(tableName, "purchase_order_id")
	_companyAddressORM.State = field.NewString(tableName, "state")
	_companyAddressORM.Street_1 = field.NewString(tableName, "street_1")
	_companyAddressORM.Street_2 = field.NewString(tableName, "street_2")
	_companyAddressORM.Type = field.NewString(tableName, "type")
	_companyAddressORM.ZipCode = field.NewString(tableName, "zip_code")

	_companyAddressORM.fillFieldMap()

	return _companyAddressORM
}

type companyAddressORM struct {
	companyAddressORMDo

	ALL                field.Asterisk
	City               field.String
	CompanyInfoId      field.Uint64
	Country            field.String
	CountrySubdivision field.String
	Id                 field.Uint64
	ModifiedAt         field.Time
	PurchaseOrderId    field.Uint64
	State              field.String
	Street_1           field.String
	Street_2           field.String
	Type               field.String
	ZipCode            field.String

	fieldMap map[string]field.Expr
}

func (c companyAddressORM) Table(newTableName string) *companyAddressORM {
	c.companyAddressORMDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c companyAddressORM) As(alias string) *companyAddressORM {
	c.companyAddressORMDo.DO = *(c.companyAddressORMDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *companyAddressORM) updateTableName(table string) *companyAddressORM {
	c.ALL = field.NewAsterisk(table)
	c.City = field.NewString(table, "city")
	c.CompanyInfoId = field.NewUint64(table, "company_info_id")
	c.Country = field.NewString(table, "country")
	c.CountrySubdivision = field.NewString(table, "country_subdivision")
	c.Id = field.NewUint64(table, "id")
	c.ModifiedAt = field.NewTime(table, "modified_at")
	c.PurchaseOrderId = field.NewUint64(table, "purchase_order_id")
	c.State = field.NewString(table, "state")
	c.Street_1 = field.NewString(table, "street_1")
	c.Street_2 = field.NewString(table, "street_2")
	c.Type = field.NewString(table, "type")
	c.ZipCode = field.NewString(table, "zip_code")

	c.fillFieldMap()

	return c
}

func (c *companyAddressORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *companyAddressORM) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["city"] = c.City
	c.fieldMap["company_info_id"] = c.CompanyInfoId
	c.fieldMap["country"] = c.Country
	c.fieldMap["country_subdivision"] = c.CountrySubdivision
	c.fieldMap["id"] = c.Id
	c.fieldMap["modified_at"] = c.ModifiedAt
	c.fieldMap["purchase_order_id"] = c.PurchaseOrderId
	c.fieldMap["state"] = c.State
	c.fieldMap["street_1"] = c.Street_1
	c.fieldMap["street_2"] = c.Street_2
	c.fieldMap["type"] = c.Type
	c.fieldMap["zip_code"] = c.ZipCode
}

func (c companyAddressORM) clone(db *gorm.DB) companyAddressORM {
	c.companyAddressORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c companyAddressORM) replaceDB(db *gorm.DB) companyAddressORM {
	c.companyAddressORMDo.ReplaceDB(db)
	return c
}

type companyAddressORMDo struct{ gen.DO }

type ICompanyAddressORMDo interface {
	gen.SubQuery
	Debug() ICompanyAddressORMDo
	WithContext(ctx context.Context) ICompanyAddressORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICompanyAddressORMDo
	WriteDB() ICompanyAddressORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICompanyAddressORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICompanyAddressORMDo
	Not(conds ...gen.Condition) ICompanyAddressORMDo
	Or(conds ...gen.Condition) ICompanyAddressORMDo
	Select(conds ...field.Expr) ICompanyAddressORMDo
	Where(conds ...gen.Condition) ICompanyAddressORMDo
	Order(conds ...field.Expr) ICompanyAddressORMDo
	Distinct(cols ...field.Expr) ICompanyAddressORMDo
	Omit(cols ...field.Expr) ICompanyAddressORMDo
	Join(table schema.Tabler, on ...field.Expr) ICompanyAddressORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICompanyAddressORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICompanyAddressORMDo
	Group(cols ...field.Expr) ICompanyAddressORMDo
	Having(conds ...gen.Condition) ICompanyAddressORMDo
	Limit(limit int) ICompanyAddressORMDo
	Offset(offset int) ICompanyAddressORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICompanyAddressORMDo
	Unscoped() ICompanyAddressORMDo
	Create(values ...*accounting_servicev1.CompanyAddressORM) error
	CreateInBatches(values []*accounting_servicev1.CompanyAddressORM, batchSize int) error
	Save(values ...*accounting_servicev1.CompanyAddressORM) error
	First() (*accounting_servicev1.CompanyAddressORM, error)
	Take() (*accounting_servicev1.CompanyAddressORM, error)
	Last() (*accounting_servicev1.CompanyAddressORM, error)
	Find() ([]*accounting_servicev1.CompanyAddressORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.CompanyAddressORM, err error)
	FindInBatches(result *[]*accounting_servicev1.CompanyAddressORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.CompanyAddressORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICompanyAddressORMDo
	Assign(attrs ...field.AssignExpr) ICompanyAddressORMDo
	Joins(fields ...field.RelationField) ICompanyAddressORMDo
	Preload(fields ...field.RelationField) ICompanyAddressORMDo
	FirstOrInit() (*accounting_servicev1.CompanyAddressORM, error)
	FirstOrCreate() (*accounting_servicev1.CompanyAddressORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.CompanyAddressORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICompanyAddressORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.CompanyAddressORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.CompanyAddressORM, err error)
	CreateRecord(item accounting_servicev1.CompanyAddressORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.CompanyAddressORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.CompanyAddressORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.CompanyAddressORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.CompanyAddressORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c companyAddressORMDo) GetRecordByID(id int) (result accounting_servicev1.CompanyAddressORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM company_addresses ")
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
func (c companyAddressORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.CompanyAddressORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM company_addresses ")
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
func (c companyAddressORMDo) CreateRecord(item accounting_servicev1.CompanyAddressORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO company_addresses (columns) VALUES (values) ")

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
func (c companyAddressORMDo) UpdateRecordByID(id int, item accounting_servicev1.CompanyAddressORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE company_addresses SET columns=values ")
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
func (c companyAddressORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM company_addresses ")
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
func (c companyAddressORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.CompanyAddressORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM company_addresses ORDER BY " + c.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (c companyAddressORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM company_addresses ")

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
func (c companyAddressORMDo) GetByID(id uint64) (result accounting_servicev1.CompanyAddressORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM company_addresses ")
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
func (c companyAddressORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.CompanyAddressORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM company_addresses ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c companyAddressORMDo) Debug() ICompanyAddressORMDo {
	return c.withDO(c.DO.Debug())
}

func (c companyAddressORMDo) WithContext(ctx context.Context) ICompanyAddressORMDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c companyAddressORMDo) ReadDB() ICompanyAddressORMDo {
	return c.Clauses(dbresolver.Read)
}

func (c companyAddressORMDo) WriteDB() ICompanyAddressORMDo {
	return c.Clauses(dbresolver.Write)
}

func (c companyAddressORMDo) Session(config *gorm.Session) ICompanyAddressORMDo {
	return c.withDO(c.DO.Session(config))
}

func (c companyAddressORMDo) Clauses(conds ...clause.Expression) ICompanyAddressORMDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c companyAddressORMDo) Returning(value interface{}, columns ...string) ICompanyAddressORMDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c companyAddressORMDo) Not(conds ...gen.Condition) ICompanyAddressORMDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c companyAddressORMDo) Or(conds ...gen.Condition) ICompanyAddressORMDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c companyAddressORMDo) Select(conds ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c companyAddressORMDo) Where(conds ...gen.Condition) ICompanyAddressORMDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c companyAddressORMDo) Order(conds ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c companyAddressORMDo) Distinct(cols ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c companyAddressORMDo) Omit(cols ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c companyAddressORMDo) Join(table schema.Tabler, on ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c companyAddressORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c companyAddressORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c companyAddressORMDo) Group(cols ...field.Expr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c companyAddressORMDo) Having(conds ...gen.Condition) ICompanyAddressORMDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c companyAddressORMDo) Limit(limit int) ICompanyAddressORMDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c companyAddressORMDo) Offset(offset int) ICompanyAddressORMDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c companyAddressORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICompanyAddressORMDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c companyAddressORMDo) Unscoped() ICompanyAddressORMDo {
	return c.withDO(c.DO.Unscoped())
}

func (c companyAddressORMDo) Create(values ...*accounting_servicev1.CompanyAddressORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c companyAddressORMDo) CreateInBatches(values []*accounting_servicev1.CompanyAddressORM, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c companyAddressORMDo) Save(values ...*accounting_servicev1.CompanyAddressORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c companyAddressORMDo) First() (*accounting_servicev1.CompanyAddressORM, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CompanyAddressORM), nil
	}
}

func (c companyAddressORMDo) Take() (*accounting_servicev1.CompanyAddressORM, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CompanyAddressORM), nil
	}
}

func (c companyAddressORMDo) Last() (*accounting_servicev1.CompanyAddressORM, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CompanyAddressORM), nil
	}
}

func (c companyAddressORMDo) Find() ([]*accounting_servicev1.CompanyAddressORM, error) {
	result, err := c.DO.Find()
	return result.([]*accounting_servicev1.CompanyAddressORM), err
}

func (c companyAddressORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.CompanyAddressORM, err error) {
	buf := make([]*accounting_servicev1.CompanyAddressORM, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c companyAddressORMDo) FindInBatches(result *[]*accounting_servicev1.CompanyAddressORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c companyAddressORMDo) Attrs(attrs ...field.AssignExpr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c companyAddressORMDo) Assign(attrs ...field.AssignExpr) ICompanyAddressORMDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c companyAddressORMDo) Joins(fields ...field.RelationField) ICompanyAddressORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c companyAddressORMDo) Preload(fields ...field.RelationField) ICompanyAddressORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c companyAddressORMDo) FirstOrInit() (*accounting_servicev1.CompanyAddressORM, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CompanyAddressORM), nil
	}
}

func (c companyAddressORMDo) FirstOrCreate() (*accounting_servicev1.CompanyAddressORM, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.CompanyAddressORM), nil
	}
}

func (c companyAddressORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.CompanyAddressORM, count int64, err error) {
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

func (c companyAddressORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c companyAddressORMDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c companyAddressORMDo) Delete(models ...*accounting_servicev1.CompanyAddressORM) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *companyAddressORMDo) withDO(do gen.Dao) *companyAddressORMDo {
	c.DO = *do.(*gen.DO)
	return c
}
