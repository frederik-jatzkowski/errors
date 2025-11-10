package internal

import "encoding/json"

func (e *ErrorfMany) MarshalJSON() ([]byte, error) {
	// nolint: wrapcheck
	return json.Marshal(e.ToDTO(nil))
}

func (e *ErrorfSingle) MarshalJSON() ([]byte, error) {
	// nolint: wrapcheck
	return json.Marshal(e.ToDTO(nil))
}

func (e *Join) MarshalJSON() ([]byte, error) {
	// nolint: wrapcheck
	return json.Marshal(e.ToDTO(nil))
}

func (e *Simple) MarshalJSON() ([]byte, error) {
	// nolint: wrapcheck
	return json.Marshal(e.ToDTO(nil))
}

func (e *WithStack) MarshalJSON() ([]byte, error) {
	// nolint: wrapcheck
	return json.Marshal(e.ToDTO(nil))
}
