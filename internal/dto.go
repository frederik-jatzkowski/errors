package internal

import (
	"runtime"
	"strings"

	"github.com/frederik-jatzkowski/errors/internal/dto"
	"github.com/frederik-jatzkowski/errors/internal/settings"
)

func (e *ErrorfMany) ToDTO(stack *dto.StackTrace, s *settings.Settings) *dto.Error {
	result := &dto.Error{
		Type:       "errorf",
		Wrapped:    len(e.Components.Errs),
		StackTrace: stack,
	}

	for _, component := range e.Components.Components {
		result.Add(component, s)
	}

	return result
}

func (e *ErrorfSingle) ToDTO(stack *dto.StackTrace, s *settings.Settings) *dto.Error {
	result := &dto.Error{
		Type:       "errorf",
		Wrapped:    1,
		StackTrace: stack,
	}

	for _, component := range e.components.Components {
		result.Add(component, s)
	}

	return result
}

func (err *Join) ToDTO(stack *dto.StackTrace, s *settings.Settings) *dto.Error {
	result := &dto.Error{
		Type:       "join",
		Wrapped:    len(err.Wrapped),
		StackTrace: stack,
	}

	for i, wrapped := range err.Wrapped {
		result.Add(wrapped, s)
		if i < len(err.Wrapped)-1 {
			result.Add("\n", s)
		}
	}

	return result
}

func (err *Simple) ToDTO(stack *dto.StackTrace, s *settings.Settings) *dto.Error {
	result := &dto.Error{
		Type:       "new",
		StackTrace: stack,
	}

	result.Add(err.Msg, s)

	return result
}

func (e *WithStack) ToDTO(_ *dto.StackTrace, s *settings.Settings) *dto.Error {
	if s.ShowStackTrace {
		return dto.NewError(e.Inner, e.St.ToDTO(s), s)
	}

	return dto.NewError(e.Inner, nil, s)
}

func (st *StackTrace) ToDTO(s *settings.Settings) *dto.StackTrace {
	dtoStack := &dto.StackTrace{}

funcloop:
	for _, pc := range st.Stack0[:] {
		if pc == 0 {
			break
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		file, line := fn.FileLine(pc)
		name := fn.Name()

		for _, prefix := range s.IgnoredFunctionPrefixes {
			if strings.HasPrefix(name, prefix) {
				continue funcloop
			}
		}

		dtoStack.Functions = append(dtoStack.Functions, dto.Function{
			Name: name,
			File: file,
			Line: line,
		})
	}

	return dtoStack
}
