package fswatch

// #include <libfswatch/c/libfswatch.h>
import "C"

type FilterType C.enum_fsw_filter_type

var (
	FilterInclude FilterType = C.filter_include
	FilterExclude FilterType = C.filter_exclude
)

type Filter struct {
	// POSIX regular expression used to match the paths (see https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap09.html).
	Text       string
	FilterType FilterType
	// Indicates if the regular expression is case sensitive or not.
	CaseSensitive bool
	// Indicates if the regular exoression is an extended regular expression or not (see https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap09.html#tag_09_04).
	Extended bool
}
