package internal

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors/internal/dto"
	"github.com/frederik-jatzkowski/errors/internal/settings"
)

func (e *ErrorfMany) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		_ = e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailFullStackTrace)).
			Write(dto.NewWriter(s, 0))
	case shouldPrintErrorMessage(s, verb):
		_ = e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailSimple)).
			Write(dto.NewWriter(s, 0))
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func (e *ErrorfSingle) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		_ = e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailFullStackTrace)).
			Write(dto.NewWriter(s, 0))
	case shouldPrintErrorMessage(s, verb):
		_ = e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailSimple)).
			Write(dto.NewWriter(s, 0))
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func (err *Join) Format(s fmt.State, verb rune) {
	if shouldPrintStack(s, verb) {
		_ = err.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailFullStackTrace)).
			Write(dto.NewWriter(s, -1))
	} else {
		_ = err.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailSimple)).Write(dto.NewWriter(s, -1))
	}
}

func (e *WithStack) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		_ = e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailFullStackTrace)).
			Write(dto.NewWriter(s, 0))
	case shouldPrintErrorMessage(s, verb):
		_ = e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailSimple)).
			Write(dto.NewWriter(s, 0))
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func shouldPrintStack(s fmt.State, verb rune) bool {
	return verb == 'v' && s.Flag('+')
}

func shouldPrintErrorMessage(s fmt.State, verb rune) bool {
	return verb == 'v' && !s.Flag('+') || verb == 's'
}
