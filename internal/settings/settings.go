// Package settings contains package wide settings
package settings

// Defaults contains all the default Settings of this package. It will be used whenever.
var Defaults = Settings{
	ShowStackTrace:          false,
	ShouldForwardVerbs:      false,
	IgnoredFunctionPrefixes: []string{"runtime", "internal/runtime", "testing"},
	StrippedFileNamePrefix:  "",
	StrippedFuncNamePrefix:  "",
}

// Settings represents the available settings for error formatting
type Settings struct {
	StrippedFileNamePrefix  string
	StrippedFuncNamePrefix  string
	IgnoredFunctionPrefixes []string
	ShowStackTrace          bool
	ShouldForwardVerbs      bool
}

func (s *Settings) CloneWithStackTrace() *Settings {
	result := *s
	result.ShowStackTrace = true
	return &result
}
