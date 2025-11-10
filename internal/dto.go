package internal

import (
	"runtime"

	"github.com/frederik-jatzkowski/errors/internal/dto"
)

func (e *ErrorfMany) ToDTO(stack *dto.StackTrace) *dto.Error {
	err := &dto.Error{
		Type:       "errorf",
		Wrapped:    len(e.Components.Errs),
		StackTrace: stack,
	}

	for _, component := range e.Components.Components {
		err.Add(component)
	}

	return err
}

func (e *ErrorfSingle) ToDTO(stack *dto.StackTrace) *dto.Error {
	err := &dto.Error{
		Type:       "errorf",
		Wrapped:    1,
		StackTrace: stack,
	}

	for _, component := range e.components.Components {
		err.Add(component)
	}

	return err
}

func (e *Join) ToDTO(stack *dto.StackTrace) *dto.Error {
	err := &dto.Error{
		Type:       "join",
		Wrapped:    len(e.Wrapped),
		StackTrace: stack,
	}

	for i, wrapped := range e.Wrapped {
		err.Add(wrapped)
		if i < len(e.Wrapped)-1 {
			err.Add("\n")
		}
	}

	return err
}

func (e *Simple) ToDTO(stack *dto.StackTrace) *dto.Error {
	err := &dto.Error{
		Type:       "new",
		StackTrace: stack,
	}

	err.Add(e.Msg)

	return err
}

func (e *WithStack) ToDTO(_ *dto.StackTrace) *dto.Error {
	return dto.NewError(e.Inner, e.St.ToDTO())
}

func (st *StackTrace) ToDTO() *dto.StackTrace {
	dtoStack := &dto.StackTrace{}
	for _, pc := range st.Stack0[:] {
		if pc == 0 {
			break
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		file, line := fn.FileLine(pc)

		dtoStack.Functions = append(dtoStack.Functions, dto.Function{
			Name: fn.Name(),
			File: file,
			Line: line,
		})
	}

	return dtoStack
}
