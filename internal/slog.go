package internal

import (
	"fmt"
	"log/slog"
)

func (e *ErrorfMany) LogValue() slog.Value {
	return logValue(e)
}

func (e *ErrorfSingle) LogValue() slog.Value {
	return logValue(e)
}

func (err *Join) LogValue() slog.Value {
	return logValue(err)
}

func (err *Simple) LogValue() slog.Value {
	return logValue(err)
}

func (e *WithStack) LogValue() slog.Value {
	return logValue(e)
}

func logValue(err error) slog.Value {
	return slog.GroupValue(
		slog.String("short", err.Error()),
		slog.String("long", fmt.Sprintf("%+v", err)),
	)
}
