package internalerror

import (
	"errors"

	"github.com/jackc/pgerrcode"
)

// Определение переменной для ошибки конфликта
var ErrConflict = errors.New(pgerrcode.UniqueViolation)
