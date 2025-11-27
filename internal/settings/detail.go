package settings

// Detail represents the level of detail that should go be present in the dto
type Detail byte

const (
	// DetailSimple represents a simple error message
	DetailSimple Detail = iota
	// DetailStackTrace represents a error message with full stack trace information
	DetailStackTrace
)
