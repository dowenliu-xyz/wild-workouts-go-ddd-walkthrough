package errors

import "github.com/pkg/errors"

type ErrorType struct {
	t string
}

var (
	errorTypeUnknown        = ErrorType{"unknown"}
	errorTypeAuthorization  = ErrorType{"authorization"}
	errorTypeIncorrectInput = ErrorType{"incorrect-input"}
)

func IsErrorTypeUnknown(errorType ErrorType) bool {
	return errorType == errorTypeUnknown
}

func IsErrorTypeAuthorization(errorType ErrorType) bool {
	return errorType == errorTypeAuthorization
}

func IsErrorTypeIncorrectInput(errorType ErrorType) bool {
	return errorType == errorTypeIncorrectInput
}

type SlugError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewSlugError(error string, slug string) error {
	return errors.WithStack(SlugError{
		error:     error,
		slug:      slug,
		errorType: errorTypeUnknown,
	})
}

func NewAuthorizationError(error string, slug string) error {
	return errors.WithStack(SlugError{
		error:     error,
		slug:      slug,
		errorType: errorTypeAuthorization,
	})
}

func NewIncorrectInputError(error string, slug string) error {
	return errors.WithStack(SlugError{
		error:     error,
		slug:      slug,
		errorType: errorTypeIncorrectInput,
	})
}
