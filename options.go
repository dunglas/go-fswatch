package fswatch

// #include <stdlib.h>
// #include <libfswatch/c/libfswatch.h>
import "C"

// Option instances allow to configure monitors.
type Option func(h *opt) error

// opt contains the available options.
type opt struct {
	monitorType    C.enum_fsw_monitor_type
	properties     map[*C.char]*C.char
	allowOverflow  bool
	latency        C.double
	recursive      bool
	directoryOnly  bool
	followSymlinks bool
	eventTypes     []EventType
	filters        []C.fsw_cmonitor_filter
}

func WithMonitorType(monitorType MonitorType) Option {
	return func(o *opt) error {
		o.monitorType = C.enum_fsw_monitor_type(monitorType)

		return nil
	}
}

func WithProperties(properties map[string]string) Option {
	return func(o *opt) error {
		o.properties = make(map[*C.char]*C.char, len(properties))
		for k, v := range properties {
			o.properties[C.CString(k)] = C.CString(v)
		}

		return nil
	}
}

func WithAllowOverflow(allowOverflow bool) Option {
	return func(o *opt) error {
		o.allowOverflow = allowOverflow

		return nil
	}
}

func WithLatency(latency float64) Option {
	return func(o *opt) error {
		o.latency = C.double(latency)

		return nil
	}
}

func WithRecursive(recursive bool) Option {
	return func(o *opt) error {
		o.recursive = recursive

		return nil
	}
}

func WithDirectoryOnly(directoryOnly bool) Option {
	return func(o *opt) error {
		o.directoryOnly = directoryOnly

		return nil
	}
}

func WithFollowSymlinks(followSymlinks bool) Option {
	return func(o *opt) error {
		o.followSymlinks = followSymlinks

		return nil
	}
}

func WithEventTypeFilters(eventTypes []EventType) Option {
	return func(o *opt) error {
		o.eventTypes = eventTypes

		return nil
	}
}

func WithFilters(filters []Filter) Option {
	return func(o *opt) error {
		o.filters = make([]C.fsw_cmonitor_filter, len(filters))

		for i, f := range filters {
			o.filters[i] = C.fsw_cmonitor_filter{text: C.CString(f.Text), _type: C.enum_fsw_filter_type(f.FilterType), case_sensitive: C.bool(f.CaseSensitive), extended: C.bool(f.Extended)}
		}

		return nil
	}
}
