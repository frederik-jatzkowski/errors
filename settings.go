package errors

import "github.com/frederik-jatzkowski/errors/internal/settings"

// EnableAdvancedFormattingOfExternalErrors will allow forwarding of the %+v verb to external errors.
// This setting allows for more stack trace information if you still rely on stack traces from other error handling libraries.
//
// Keep in mind that this adds valuable debugging information but might also lead to redundant stack traces.
// Additionally, the formatting of the wrapped errors might not work as nicely with the formatting of this package.
//
// This is a package wide option and should only be called once in the main function of your application.
func EnableAdvancedFormattingOfExternalErrors() {
	settings.Defaults.ShouldForwardVerbs = true
}
