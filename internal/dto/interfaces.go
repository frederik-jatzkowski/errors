package dto

type ToDTOer interface {
	ToDTO(stack *StackTrace) *Error
}

type DTO interface {
	WriteShort(w *Writer) error
	WriteLong(w *Writer) error
}
