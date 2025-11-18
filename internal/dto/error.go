// Package dto provides a more suitable intermediate representation for marshalling errors into different output formats.
package dto

type Error struct {
	StackTrace *StackTrace `json:"-"`
	Type       string      `json:"type"`
	Components []any       `json:"components"`
	Wrapped    int         `json:"wrapped"`
}

var _ DTO = (*Error)(nil)

func NewError(err error, stack *StackTrace) *Error {
	toDTOer, ok := err.(ToDTOer)
	if !ok {
		return &Error{
			Type:       "external",
			Components: []any{err.Error()},
			StackTrace: stack,
		}
	}

	return toDTOer.ToDTO(stack)
}

func (e *Error) Add(component any) {
	err, ok := component.(error)
	if ok {
		e.Components = append(e.Components, NewError(err, nil))

		return
	}

	e.Components = append(e.Components, component)
}

func (e *Error) WriteShort(w *Writer) error {
	for _, component := range e.Components {
		dto, ok := component.(DTO)
		if ok {
			err := func() error {
				if e.Wrapped > 1 {
					if !w.IsAfterNewline() {
						w.AddNewline()
					}

					w.Descend()
					defer w.Ascend()

					_, err := w.Write([]byte("=> "))
					if err != nil {
						return err
					}
				}

				err := dto.WriteShort(w)
				if err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}
		}

		str, ok := component.(string)
		if ok {
			_, err := w.Write([]byte(str))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Error) WriteLong(w *Writer) error {
	for _, component := range e.Components {
		dto, ok := component.(DTO)
		if ok {
			err := func() error {
				if e.Wrapped > 1 {
					if !w.IsAfterNewline() {
						w.AddNewline()
					}

					w.Descend()
					defer w.Ascend()

					_, err := w.Write([]byte("=> "))
					if err != nil {
						return err
					}
				}

				err := dto.WriteLong(w)
				if err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}
		}

		str, ok := component.(string)
		if ok {
			_, err := w.Write([]byte(str))
			if err != nil {
				return err
			}
		}
	}

	if e.StackTrace != nil {
		err := e.StackTrace.WriteLong(w)
		if err != nil {
			return err
		}
	}

	return nil
}
