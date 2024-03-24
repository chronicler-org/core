package appUtil

import (
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	gormErrors = map[error]struct {
		Message string
		Status  int
	}{
		logger.ErrRecordNotFound:              {Message: "Record not found", Status: http.StatusNotFound},
		gorm.ErrInvalidTransaction:            {Message: "Invalid transaction", Status: http.StatusInternalServerError},
		gorm.ErrNotImplemented:                {Message: "Not implemented", Status: http.StatusNotImplemented},
		gorm.ErrMissingWhereClause:            {Message: "Missing WHERE clause", Status: http.StatusBadRequest},
		gorm.ErrUnsupportedRelation:           {Message: "Unsupported relations", Status: http.StatusBadRequest},
		gorm.ErrPrimaryKeyRequired:            {Message: "Primary keys required", Status: http.StatusBadRequest},
		gorm.ErrModelValueRequired:            {Message: "Model value required", Status: http.StatusBadRequest},
		gorm.ErrModelAccessibleFieldsRequired: {Message: "Model accessible fields required", Status: http.StatusBadRequest},
		gorm.ErrSubQueryRequired:              {Message: "Sub query required", Status: http.StatusBadRequest},
		gorm.ErrInvalidData:                   {Message: "Invalid data", Status: http.StatusBadRequest},
		gorm.ErrUnsupportedDriver:             {Message: "Unsupported driver", Status: http.StatusBadRequest},
		gorm.ErrRegistered:                    {Message: "Registered", Status: http.StatusBadRequest},
		gorm.ErrInvalidField:                  {Message: "Invalid field", Status: http.StatusBadRequest},
		gorm.ErrEmptySlice:                    {Message: "Empty slice found", Status: http.StatusBadRequest},
		gorm.ErrDryRunModeUnsupported:         {Message: "Dry run mode unsupported", Status: http.StatusBadRequest},
		gorm.ErrInvalidDB:                     {Message: "Invalid DB", Status: http.StatusBadRequest},
		gorm.ErrInvalidValue:                  {Message: "Invalid value, should be pointer to struct or slice", Status: http.StatusBadRequest},
		gorm.ErrInvalidValueOfLength:          {Message: "Invalid association values, length doesn't match", Status: http.StatusBadRequest},
		gorm.ErrPreloadNotAllowed:             {Message: "Preload not allowed when count is used", Status: http.StatusBadRequest},
		gorm.ErrDuplicatedKey:                 {Message: "Duplicated key not allowed", Status: http.StatusBadRequest},
		gorm.ErrForeignKeyViolated:            {Message: "Violates foreign key constraint", Status: http.StatusBadRequest},
	}
)
