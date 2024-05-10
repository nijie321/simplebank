package api

import (
	"simplebank/db/util"

	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(filedLevel validator.FieldLevel) bool {
	if currency, ok := filedLevel.Field().Interface().(string); ok {
		// check if currency is supported
		return util.IsSupportedCurrency(currency)
	}

	return false
}
