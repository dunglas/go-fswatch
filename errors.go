package fswatch

import "errors"
import "C"

var (
	UnknownError               = errors.New("an unknown error has occurred")
	SessionUnknownError        = errors.New("the session specified by the handle is unknown")
	MonitorAlreadyExistsError  = errors.New("the session already contains a monitor")
	MemoryError                = errors.New("an error occurred while invoking a memory management routine")
	UnknownMonitorTypeError    = errors.New("the specified monitor type does not exist")
	CallbackNotSetError        = errors.New("the callback has not been set")
	PathsNotSetError           = errors.New("the paths to watch have not been set")
	MissingContextError        = errors.New("the callback context has not been set")
	InvalidPathError           = errors.New("the path is invalid")
	InvalidCallbackError       = errors.New("the callback is invalid")
	InvalidLatencyError        = errors.New("the latency is invalid")
	InvalidRegexError          = errors.New("the regular expression is invalid")
	MonitorAlreadyRunningError = errors.New("a monitor is already running in the specified session")
	UnknownValueError          = errors.New("the value is unknown")
	InvalidPropertyError       = errors.New("the property is invalid")
)

var errs map[C.int]error = map[C.int]error{
	0:       nil,
	1 << 0:  UnknownError,
	1 << 1:  SessionUnknownError,
	1 << 2:  MonitorAlreadyExistsError,
	1 << 3:  MemoryError,
	1 << 4:  UnknownMonitorTypeError,
	1 << 5:  CallbackNotSetError,
	1 << 6:  PathsNotSetError,
	1 << 7:  MissingContextError,
	1 << 8:  InvalidPathError,
	1 << 9:  InvalidCallbackError,
	1 << 10: InvalidLatencyError,
	1 << 11: InvalidRegexError,
	1 << 12: MonitorAlreadyRunningError,
	1 << 13: UnknownValueError,
	1 << 14: InvalidPropertyError,
}
