package mappers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"strings"
)

var validate = validator.New()

// Função genérica para converter DTO para Model
func toModel[T any, M any](dto T, model M) (M, error) {
	// Validação do DTO
	if err := validate.Struct(dto); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field()+" is invalid")
		}
		return model, errors.New("validation failed: " + strings.Join(validationErrors, ", "))
	}

	// Conversão automática com copier
	err := copier.Copy(&model, &dto)
	if err != nil {
		return model, errors.New("failed to copy fields")
	}

	return model, nil
}
