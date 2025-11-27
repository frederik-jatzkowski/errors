package errors

import "github.com/frederik-jatzkowski/errors/internal/settings"

// FormatSetting is an option that can modify this package's formatting behavior.
type FormatSetting func(*settings.Settings)

// GlobalFormatSettings applies the given settings to the package's default settings.
//
// This has global effects and should only be called once in the main function of your application.
func GlobalFormatSettings(s ...FormatSetting) {
	for _, setting := range s {
		setting(&settings.Defaults)
	}
}

// WithAdvancedFormattingOfExternalErrors is a [FormatSetting] that will allow forwarding of the %+v verb to external errors.
// This setting allows for more stack trace information if you still rely on stack traces from other error handling libraries.
//
// Keep in mind that this adds valuable debugging information but might also lead to redundant stack traces.
// Additionally, the formatting of the wrapped errors might not work as nicely with the formatting of this package.
func WithAdvancedFormattingOfExternalErrors() FormatSetting {
	return func(settings *settings.Settings) {
		settings.ShouldForwardVerbs = true
	}
}

// WithIgnoredFunctionPrefixes is a [FormatSetting] that will prevent functions with names starting with the given prefix from showing up in stack traces.
// This is useful to keep go internals out of your application logs.
//
// The default for this package is already set to []string{"runtime", "internal/runtime", "testing"}.
func WithIgnoredFunctionPrefixes(prefixes ...string) FormatSetting {
	return func(settings *settings.Settings) {
		settings.IgnoredFunctionPrefixes = prefixes
	}
}
