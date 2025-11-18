package internal

import (
	"runtime"

	"github.com/frederik-jatzkowski/errors/internal/dto"
)

func (e *ErrorfMany) ToDTO(stack *dto.StackTrace) *dto.Error {
	result := &dto.Error{
		Type:       "errorf",
		Wrapped:    len(e.Components.Errs),
		StackTrace: stack,
	}

	for _, component := range e.Components.Components {
		result.Add(component)
	}

	return result
}

func (e *ErrorfSingle) ToDTO(stack *dto.StackTrace) *dto.Error {
	result := &dto.Error{
		Type:       "errorf",
		Wrapped:    1,
		StackTrace: stack,
	}

	for _, component := range e.components.Components {
		result.Add(component)
	}

	return result
}

func (err *Join) ToDTO(stack *dto.StackTrace) *dto.Error {
	result := &dto.Error{
		Type:       "join",
		Wrapped:    len(err.Wrapped),
		StackTrace: stack,
	}

	for i, wrapped := range err.Wrapped {
		result.Add(wrapped)
		if i < len(err.Wrapped)-1 {
			result.Add("\n")
		}
	}

	return result
}

func (err *Simple) ToDTO(stack *dto.StackTrace) *dto.Error {
	result := &dto.Error{
		Type:       "new",
		StackTrace: stack,
	}

	result.Add(err.Msg)

	return result
}

func (err *WithStack) ToDTO(_ *dto.StackTrace) *dto.Error {
	return dto.NewError(err.Inner, err.St.ToDTO())
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
