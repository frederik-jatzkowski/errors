package internal

import (
	"encoding/json"

	"github.com/frederik-jatzkowski/errors/internal/settings"
)

func (e *ErrorfMany) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailStackTrace)),
	)
}

func (e *ErrorfSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailStackTrace)),
	)
}

func (err *Join) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		err.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailStackTrace)),
	)
}

func (err *Simple) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		err.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailStackTrace)),
	)
}

func (e *WithStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		e.ToDTO(nil, settings.Defaults.CloneWithDetail(settings.DetailStackTrace)),
	)
}
