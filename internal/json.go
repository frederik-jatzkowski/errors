package internal

import (
	"encoding/json"

	"github.com/frederik-jatzkowski/errors/internal/settings"
)

func (e *ErrorfMany) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		e.ToDTO(nil, settings.Defaults.CloneWithStackTrace()),
	)
}

func (e *ErrorfSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		e.ToDTO(nil, settings.Defaults.CloneWithStackTrace()),
	)
}

func (err *Join) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		err.ToDTO(nil, settings.Defaults.CloneWithStackTrace()),
	)
}

func (err *Simple) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		err.ToDTO(nil, settings.Defaults.CloneWithStackTrace()),
	)
}

func (e *WithStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		e.ToDTO(nil, settings.Defaults.CloneWithStackTrace()),
	)
}
