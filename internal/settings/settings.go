// Package settings contains package wide settings
package settings

// Defaults contains all the default Settings of this package. It will be used whenever.
var Defaults = Settings{
	Detail:                  DetailSimple,
	ShouldForwardVerbs:      false,
	IgnoredFunctionPrefixes: []string{"runtime", "internal/runtime", "testing"},
}

// Settings represents the available settings for error formatting
type Settings struct {
	Detail                  Detail
	ShouldForwardVerbs      bool
	IgnoredFunctionPrefixes []string
}

func (s *Settings) CloneWithDetail(detail Detail) *Settings {
	result := *s
	result.Detail = detail
	return &result
}
