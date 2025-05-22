package validation

import (
	"fmt"
	"strings"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/common/httpfilter"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s any) error {
	if err := validate.Struct(s); err != nil {
		if verrs, ok := err.(validator.ValidationErrors); ok {
			msgs := make([]string, len(verrs))
			for i, e := range verrs {
				msgs[i] = fmt.Sprintf("%s is invalid (tag=%s, param=%s)", e.Field(), e.ActualTag(), e.Param())
			}
			return httpfilter.NewValidationError(strings.Join(msgs, ", "))
		}
		return httpfilter.NewValidationError("Validation failed")
	}
	return nil
}
