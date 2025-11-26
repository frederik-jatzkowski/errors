package dto

import "github.com/frederik-jatzkowski/errors/internal/settings"

type ToDTOer interface {
	ToDTO(stack *StackTrace, settings *settings.Settings) *Error
}

type DTO interface {
	Write(w *Writer) error
}
