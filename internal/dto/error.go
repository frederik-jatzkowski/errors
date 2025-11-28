// Package dto provides a more suitable intermediate representation for marshalling errors into different output formats.
package dto

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors/internal/settings"
)

type Error struct {
	StackTrace *StackTrace `json:"stack_trace,omitempty"`
	Type       string      `json:"type"`
	Components []any       `json:"components"`
	Wrapped    int         `json:"wrapped"`
}

var _ DTO = (*Error)(nil)

func NewError(err error, stack *StackTrace, s *settings.Settings) *Error {
	toDTOer, ok := err.(ToDTOer)
	if !ok {
		var wrapped string
		if s.ShouldForwardVerbs && s.ShowStackTrace {
			wrapped = fmt.Sprintf("%+v", err)
		} else {
			wrapped = err.Error()
		}

		return &Error{
			Type:       "external",
			Components: []any{wrapped},
			StackTrace: stack,
		}
	}

	return toDTOer.ToDTO(stack, s)
}

func (e *Error) Add(component any, s *settings.Settings) {
	err, ok := component.(error)
	if ok {
		e.Components = append(e.Components, NewError(err, nil, s))

		return
	}

	e.Components = append(e.Components, component)
}

func (e *Error) Write(w *Writer) error {
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

				err := dto.Write(w)
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
		err := e.StackTrace.Write(w)
		if err != nil {
			return err
		}
	}

	return nil
}
