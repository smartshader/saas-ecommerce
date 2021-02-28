package engine

const (
	// ContextOriginalPath holds the original request url
	ContextOriginalPath key = iota

	// ContextRequestStart holds the request start time
	ContextRequestStart

	// ContextUserID holds the user ID
	ContextUserID
)

type key int
