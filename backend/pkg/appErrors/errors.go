package apperrors

import "fmt"

// Error define uma estrutura para tratamento de erros com exposição controlada.
type Error struct {
	Code        string // Código interno do erro
	Description string // Descrição genérica do erro para o cliente
	Internal    string // Mensagem detalhada para logs internos
}

// Error implementa a interface padrão de erros do Go.
func (e *Error) Error() string {
	return e.Description
}

// Funções para criar erros específicos:

func ErrBadRequest(entity string) *Error {
	return &Error{
		Code:        "BadRequest",
		Description: "Invalid request data.",
		Internal:    fmt.Sprintf("missing data for %s handler request", entity),
	}
}

func ErrUnprocessableEntity(entity string) *Error {
	return &Error{
		Code:        "UnprocessableEntity",
		Description: "The request could not be processed.",
		Internal:    fmt.Sprintf("incomplete or malformed %s information", entity),
	}
}

func ErrNotFound(entity string) *Error {
	return &Error{
		Code:        "NotFound",
		Description: "The requested resource was not found.",
		Internal:    fmt.Sprintf("%s not found", entity),
	}
}

func ErrConflict(entity string) *Error {
	return &Error{
		Code:        "Conflict",
		Description: "A conflict occurred with the request.",
		Internal:    fmt.Sprintf("%s already exists", entity),
	}
}

func ErrInternalServerError(entity string) *Error {
	return &Error{
		Code:        "InternalServerError",
		Description: "An internal server error occurred.",
		Internal:    fmt.Sprintf("%s: internal server error", entity),
	}
}
