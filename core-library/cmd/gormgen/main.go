package main

import (
	accountingServiceProto "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/accounting_service/v1"
	financialServiceProto "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/financial_service/v1"
	socialServiceProto "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/social_service/v2"
	userServiceProto "github.com/PlaybookMediaEngineering/core/core-library/pkg/generated/user_service/v1"

	"gorm.io/gen"
)

// Querier is the interface of added queries that the various
// types should implement.
type Querier interface {
	// SELECT * FROM @@table
	//  {{where}}
	//      id=@id
	//  {{end}}
	GetRecordByID(id int) (gen.T, error) // returns struct and error

	// SELECT * FROM @@table
	//  {{where}}
	//      id IN (@ids)
	//  {{end}}
	GetRecordByIDs(ids []int) ([]gen.T, error) // returns slice of struct and error

	// INSERT INTO @@table (columns) VALUES (values)
	CreateRecord(item gen.T) error // inserts a new record

	// UPDATE @@table SET columns=values
	//  {{where}}
	//      id=@id
	//  {{end}}
	UpdateRecordByID(id int, item gen.T) error // updates a record by ID

	// DELETE FROM @@table
	//  {{where}}
	//      id=@id
	//  {{end}}
	DeleteRecordByID(id int) error // deletes a record by ID

	// SELECT * FROM @@table
	// ORDER BY @@orderColumn
	GetAllRecords(orderColumn string, limit, offset int) ([]gen.T, error)

	// Additional Operations
	// SELECT COUNT(*) FROM @@table
	CountAll() (int, error) // returns count of all records

	// SELECT * FROM @@table
	//  {{where}}
	//      id=@id
	//  {{end}}
	GetByID(id uint64) (gen.T, error) // returns struct and error
	// SELECT * FROM @@table
	//  {{where}}
	//      id IN (@ids)
	//  {{end}}
	GetByIDs(ids []uint64) ([]gen.T, error) // returns slice of struct and error
}

func main() {
	socialModels := getAllSocialModels()
	socialOutputPath := "../../pkg/generated/social_service/dal"
	autogenerateDal(&socialOutputPath, socialModels)

	userOutputPath := "../../pkg/generated/user_service/dal"
	userModels := getAllUserModels()
	autogenerateDal(&userOutputPath, userModels)

	financialOutputPath := "../../pkg/generated/financial_service/dal"
	models := getAllFinancialModels()
	autogenerateDal(&financialOutputPath, models)

	accountingOutputPath := "../../pkg/generated/accounting_service/dal"
	accountingModels := getAllAccountingModels()
	autogenerateDal(&accountingOutputPath, accountingModels)
}

func getAllSocialModels() []interface{} {
	models := socialServiceProto.GetDatabaseSchemas()
	return models
}

func getAllUserModels() []interface{} {
	models := userServiceProto.GetDatabaseSchemas()
	return models
}

// getAllFinancialModels combines database schemas from different sources.
func getAllFinancialModels() []interface{} {
	models := financialServiceProto.GetDatabaseSchemas()
	return models
}

func getAllAccountingModels() []interface{} {
	models := accountingServiceProto.GetDatabaseSchemas()
	return models
}

// autogenerateDal generates data access layer for provided models.
func autogenerateDal(outputPath *string, models []interface{}) {
	if outputPath == nil {
		return
	}

	g := createGenerator(*outputPath)
	g.ApplyBasic(models...)
	g.ApplyInterface(func(Querier) {}, models...)
	g.Execute()
}

// createGenerator initializes a new gen.Generator with the given configuration.
func createGenerator(outputPath string) *gen.Generator {
	// define the generator configuration
	return gen.NewGenerator(gen.Config{
		OutPath:           outputPath,
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
}
