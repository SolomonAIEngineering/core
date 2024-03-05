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

func newDependentsORM(db *gorm.DB, opts ...gen.DOOption) dependentsORM {
	_dependentsORM := dependentsORM{}

	_dependentsORM.dependentsORMDo.UseDB(db, opts...)
	_dependentsORM.dependentsORMDo.UseModel(&accounting_servicev1.DependentsORM{})

	tableName := _dependentsORM.dependentsORMDo.TableName()
	_dependentsORM.ALL = field.NewAsterisk(tableName)
	_dependentsORM.CreatedAt = field.NewTime(tableName, "created_at")
	_dependentsORM.DateOfBirth = field.NewTime(tableName, "date_of_birth")
	_dependentsORM.DependentRelationshipToEmployee = field.NewString(tableName, "dependent_relationship_to_employee")
	_dependentsORM.EmployeeId = field.NewUint64(tableName, "employee_id")
	_dependentsORM.FirstName = field.NewString(tableName, "first_name")
	_dependentsORM.Gender = field.NewString(tableName, "gender")
	_dependentsORM.Id = field.NewUint64(tableName, "id")
	_dependentsORM.IsStudent = field.NewBool(tableName, "is_student")
	_dependentsORM.LastName = field.NewString(tableName, "last_name")
	_dependentsORM.MergeAccountId = field.NewString(tableName, "merge_account_id")
	_dependentsORM.MiddleName = field.NewString(tableName, "middle_name")
	_dependentsORM.ModifiedAt = field.NewTime(tableName, "modified_at")
	_dependentsORM.PhoneNumber = field.NewString(tableName, "phone_number")
	_dependentsORM.RemoteId = field.NewString(tableName, "remote_id")
	_dependentsORM.RemoteWasDeleted = field.NewBool(tableName, "remote_was_deleted")
	_dependentsORM.Ssn = field.NewString(tableName, "ssn")
	_dependentsORM.HomeLocation = dependentsORMHasOneHomeLocation{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("HomeLocation", "accounting_servicev1.LocationAddressORM"),
	}

	_dependentsORM.fillFieldMap()

	return _dependentsORM
}

type dependentsORM struct {
	dependentsORMDo

	ALL                             field.Asterisk
	CreatedAt                       field.Time
	DateOfBirth                     field.Time
	DependentRelationshipToEmployee field.String
	EmployeeId                      field.Uint64
	FirstName                       field.String
	Gender                          field.String
	Id                              field.Uint64
	IsStudent                       field.Bool
	LastName                        field.String
	MergeAccountId                  field.String
	MiddleName                      field.String
	ModifiedAt                      field.Time
	PhoneNumber                     field.String
	RemoteId                        field.String
	RemoteWasDeleted                field.Bool
	Ssn                             field.String
	HomeLocation                    dependentsORMHasOneHomeLocation

	fieldMap map[string]field.Expr
}

func (d dependentsORM) Table(newTableName string) *dependentsORM {
	d.dependentsORMDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dependentsORM) As(alias string) *dependentsORM {
	d.dependentsORMDo.DO = *(d.dependentsORMDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dependentsORM) updateTableName(table string) *dependentsORM {
	d.ALL = field.NewAsterisk(table)
	d.CreatedAt = field.NewTime(table, "created_at")
	d.DateOfBirth = field.NewTime(table, "date_of_birth")
	d.DependentRelationshipToEmployee = field.NewString(table, "dependent_relationship_to_employee")
	d.EmployeeId = field.NewUint64(table, "employee_id")
	d.FirstName = field.NewString(table, "first_name")
	d.Gender = field.NewString(table, "gender")
	d.Id = field.NewUint64(table, "id")
	d.IsStudent = field.NewBool(table, "is_student")
	d.LastName = field.NewString(table, "last_name")
	d.MergeAccountId = field.NewString(table, "merge_account_id")
	d.MiddleName = field.NewString(table, "middle_name")
	d.ModifiedAt = field.NewTime(table, "modified_at")
	d.PhoneNumber = field.NewString(table, "phone_number")
	d.RemoteId = field.NewString(table, "remote_id")
	d.RemoteWasDeleted = field.NewBool(table, "remote_was_deleted")
	d.Ssn = field.NewString(table, "ssn")

	d.fillFieldMap()

	return d
}

func (d *dependentsORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dependentsORM) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 17)
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["date_of_birth"] = d.DateOfBirth
	d.fieldMap["dependent_relationship_to_employee"] = d.DependentRelationshipToEmployee
	d.fieldMap["employee_id"] = d.EmployeeId
	d.fieldMap["first_name"] = d.FirstName
	d.fieldMap["gender"] = d.Gender
	d.fieldMap["id"] = d.Id
	d.fieldMap["is_student"] = d.IsStudent
	d.fieldMap["last_name"] = d.LastName
	d.fieldMap["merge_account_id"] = d.MergeAccountId
	d.fieldMap["middle_name"] = d.MiddleName
	d.fieldMap["modified_at"] = d.ModifiedAt
	d.fieldMap["phone_number"] = d.PhoneNumber
	d.fieldMap["remote_id"] = d.RemoteId
	d.fieldMap["remote_was_deleted"] = d.RemoteWasDeleted
	d.fieldMap["ssn"] = d.Ssn

}

func (d dependentsORM) clone(db *gorm.DB) dependentsORM {
	d.dependentsORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dependentsORM) replaceDB(db *gorm.DB) dependentsORM {
	d.dependentsORMDo.ReplaceDB(db)
	return d
}

type dependentsORMHasOneHomeLocation struct {
	db *gorm.DB

	field.RelationField
}

func (a dependentsORMHasOneHomeLocation) Where(conds ...field.Expr) *dependentsORMHasOneHomeLocation {
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

func (a dependentsORMHasOneHomeLocation) WithContext(ctx context.Context) *dependentsORMHasOneHomeLocation {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a dependentsORMHasOneHomeLocation) Session(session *gorm.Session) *dependentsORMHasOneHomeLocation {
	a.db = a.db.Session(session)
	return &a
}

func (a dependentsORMHasOneHomeLocation) Model(m *accounting_servicev1.DependentsORM) *dependentsORMHasOneHomeLocationTx {
	return &dependentsORMHasOneHomeLocationTx{a.db.Model(m).Association(a.Name())}
}

type dependentsORMHasOneHomeLocationTx struct{ tx *gorm.Association }

func (a dependentsORMHasOneHomeLocationTx) Find() (result *accounting_servicev1.LocationAddressORM, err error) {
	return result, a.tx.Find(&result)
}

func (a dependentsORMHasOneHomeLocationTx) Append(values ...*accounting_servicev1.LocationAddressORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a dependentsORMHasOneHomeLocationTx) Replace(values ...*accounting_servicev1.LocationAddressORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a dependentsORMHasOneHomeLocationTx) Delete(values ...*accounting_servicev1.LocationAddressORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a dependentsORMHasOneHomeLocationTx) Clear() error {
	return a.tx.Clear()
}

func (a dependentsORMHasOneHomeLocationTx) Count() int64 {
	return a.tx.Count()
}

type dependentsORMDo struct{ gen.DO }

type IDependentsORMDo interface {
	gen.SubQuery
	Debug() IDependentsORMDo
	WithContext(ctx context.Context) IDependentsORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDependentsORMDo
	WriteDB() IDependentsORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDependentsORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDependentsORMDo
	Not(conds ...gen.Condition) IDependentsORMDo
	Or(conds ...gen.Condition) IDependentsORMDo
	Select(conds ...field.Expr) IDependentsORMDo
	Where(conds ...gen.Condition) IDependentsORMDo
	Order(conds ...field.Expr) IDependentsORMDo
	Distinct(cols ...field.Expr) IDependentsORMDo
	Omit(cols ...field.Expr) IDependentsORMDo
	Join(table schema.Tabler, on ...field.Expr) IDependentsORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDependentsORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDependentsORMDo
	Group(cols ...field.Expr) IDependentsORMDo
	Having(conds ...gen.Condition) IDependentsORMDo
	Limit(limit int) IDependentsORMDo
	Offset(offset int) IDependentsORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDependentsORMDo
	Unscoped() IDependentsORMDo
	Create(values ...*accounting_servicev1.DependentsORM) error
	CreateInBatches(values []*accounting_servicev1.DependentsORM, batchSize int) error
	Save(values ...*accounting_servicev1.DependentsORM) error
	First() (*accounting_servicev1.DependentsORM, error)
	Take() (*accounting_servicev1.DependentsORM, error)
	Last() (*accounting_servicev1.DependentsORM, error)
	Find() ([]*accounting_servicev1.DependentsORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.DependentsORM, err error)
	FindInBatches(result *[]*accounting_servicev1.DependentsORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*accounting_servicev1.DependentsORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDependentsORMDo
	Assign(attrs ...field.AssignExpr) IDependentsORMDo
	Joins(fields ...field.RelationField) IDependentsORMDo
	Preload(fields ...field.RelationField) IDependentsORMDo
	FirstOrInit() (*accounting_servicev1.DependentsORM, error)
	FirstOrCreate() (*accounting_servicev1.DependentsORM, error)
	FindByPage(offset int, limit int) (result []*accounting_servicev1.DependentsORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDependentsORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRecordByID(id int) (result accounting_servicev1.DependentsORM, err error)
	GetRecordByIDs(ids []int) (result []accounting_servicev1.DependentsORM, err error)
	CreateRecord(item accounting_servicev1.DependentsORM) (err error)
	UpdateRecordByID(id int, item accounting_servicev1.DependentsORM) (err error)
	DeleteRecordByID(id int) (err error)
	GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.DependentsORM, err error)
	CountAll() (result int, err error)
	GetByID(id uint64) (result accounting_servicev1.DependentsORM, err error)
	GetByIDs(ids []uint64) (result []accounting_servicev1.DependentsORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (d dependentsORMDo) GetRecordByID(id int) (result accounting_servicev1.DependentsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM dependents ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (d dependentsORMDo) GetRecordByIDs(ids []int) (result []accounting_servicev1.DependentsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM dependents ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table (columns) VALUES (values)
func (d dependentsORMDo) CreateRecord(item accounting_servicev1.DependentsORM) (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("INSERT INTO dependents (columns) VALUES (values) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET columns=values
// {{where}}
//
//	id=@id
//
// {{end}}
func (d dependentsORMDo) UpdateRecordByID(id int, item accounting_servicev1.DependentsORM) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE dependents SET columns=values ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (d dependentsORMDo) DeleteRecordByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("DELETE FROM dependents ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// ORDER BY @@orderColumn
func (d dependentsORMDo) GetAllRecords(orderColumn string, limit int, offset int) (result []accounting_servicev1.DependentsORM, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM dependents ORDER BY " + d.Quote(orderColumn) + " ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Additional Operations
// SELECT COUNT(*) FROM @@table
func (d dependentsORMDo) CountAll() (result int, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("Additional Operations SELECT COUNT(*) FROM dependents ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (d dependentsORMDo) GetByID(id uint64) (result accounting_servicev1.DependentsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM dependents ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (d dependentsORMDo) GetByIDs(ids []uint64) (result []accounting_servicev1.DependentsORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM dependents ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (d dependentsORMDo) Debug() IDependentsORMDo {
	return d.withDO(d.DO.Debug())
}

func (d dependentsORMDo) WithContext(ctx context.Context) IDependentsORMDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dependentsORMDo) ReadDB() IDependentsORMDo {
	return d.Clauses(dbresolver.Read)
}

func (d dependentsORMDo) WriteDB() IDependentsORMDo {
	return d.Clauses(dbresolver.Write)
}

func (d dependentsORMDo) Session(config *gorm.Session) IDependentsORMDo {
	return d.withDO(d.DO.Session(config))
}

func (d dependentsORMDo) Clauses(conds ...clause.Expression) IDependentsORMDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dependentsORMDo) Returning(value interface{}, columns ...string) IDependentsORMDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dependentsORMDo) Not(conds ...gen.Condition) IDependentsORMDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dependentsORMDo) Or(conds ...gen.Condition) IDependentsORMDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dependentsORMDo) Select(conds ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dependentsORMDo) Where(conds ...gen.Condition) IDependentsORMDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dependentsORMDo) Order(conds ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dependentsORMDo) Distinct(cols ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dependentsORMDo) Omit(cols ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dependentsORMDo) Join(table schema.Tabler, on ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dependentsORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dependentsORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dependentsORMDo) Group(cols ...field.Expr) IDependentsORMDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dependentsORMDo) Having(conds ...gen.Condition) IDependentsORMDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dependentsORMDo) Limit(limit int) IDependentsORMDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dependentsORMDo) Offset(offset int) IDependentsORMDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dependentsORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDependentsORMDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dependentsORMDo) Unscoped() IDependentsORMDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dependentsORMDo) Create(values ...*accounting_servicev1.DependentsORM) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dependentsORMDo) CreateInBatches(values []*accounting_servicev1.DependentsORM, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dependentsORMDo) Save(values ...*accounting_servicev1.DependentsORM) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dependentsORMDo) First() (*accounting_servicev1.DependentsORM, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DependentsORM), nil
	}
}

func (d dependentsORMDo) Take() (*accounting_servicev1.DependentsORM, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DependentsORM), nil
	}
}

func (d dependentsORMDo) Last() (*accounting_servicev1.DependentsORM, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DependentsORM), nil
	}
}

func (d dependentsORMDo) Find() ([]*accounting_servicev1.DependentsORM, error) {
	result, err := d.DO.Find()
	return result.([]*accounting_servicev1.DependentsORM), err
}

func (d dependentsORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*accounting_servicev1.DependentsORM, err error) {
	buf := make([]*accounting_servicev1.DependentsORM, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dependentsORMDo) FindInBatches(result *[]*accounting_servicev1.DependentsORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dependentsORMDo) Attrs(attrs ...field.AssignExpr) IDependentsORMDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dependentsORMDo) Assign(attrs ...field.AssignExpr) IDependentsORMDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dependentsORMDo) Joins(fields ...field.RelationField) IDependentsORMDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dependentsORMDo) Preload(fields ...field.RelationField) IDependentsORMDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dependentsORMDo) FirstOrInit() (*accounting_servicev1.DependentsORM, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DependentsORM), nil
	}
}

func (d dependentsORMDo) FirstOrCreate() (*accounting_servicev1.DependentsORM, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*accounting_servicev1.DependentsORM), nil
	}
}

func (d dependentsORMDo) FindByPage(offset int, limit int) (result []*accounting_servicev1.DependentsORM, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dependentsORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dependentsORMDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dependentsORMDo) Delete(models ...*accounting_servicev1.DependentsORM) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dependentsORMDo) withDO(do gen.Dao) *dependentsORMDo {
	d.DO = *do.(*gen.DO)
	return d
}