package internal

import "encoding/json"

func (e *ErrorfMany) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToDTO(nil))
}

func (e *ErrorfSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToDTO(nil))
}

func (err *Join) MarshalJSON() ([]byte, error) {
	return json.Marshal(err.ToDTO(nil))
}

func (err *Simple) MarshalJSON() ([]byte, error) {
	return json.Marshal(err.ToDTO(nil))
}

func (err *WithStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(err.ToDTO(nil))
}
