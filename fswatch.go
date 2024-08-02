// Package fswatch provides bindings for libfswatch.
package fswatch

// #cgo LDFLAGS: -L/usr/local/lib -lfswatch
// #cgo CFLAGS: -I/usr/local/include/libfswatch/c
// #include <stdint.h>
// #include <stdlib.h>
// #include <libfswatch.h>
// #include "fswatch.h"
import "C"
import (
	"runtime/cgo"
	"time"
	"unsafe"
)

var initialized = false

type Callback func([]Event)

type Session struct {
	handle    C.FSW_HANDLE
	paths     []*C.char
	callback  Callback
	cgoHandle cgo.Handle
	*opt
}

// Start starts the monitor. Start must be called in a dedicated goroutine.
// Depending on the type of monitor this call might return when a monitor is stopped or not.
func (s *Session) Start() error {
	return errs[C.fsw_start_monitor(s.handle)]
}

// Stop stops a running monitor.
func (s *Session) Stop() error {
	return errs[C.fsw_stop_monitor(s.handle)]
}

// Destroy frees allocated system resources.
// It must be called when the session is not used anymore.
func (s *Session) Destroy() error {
	err := errs[C.fsw_destroy_session(s.handle)]

	for _, p := range s.paths {
		C.free(unsafe.Pointer(p))
	}

	for n, v := range s.properties {
		C.free(unsafe.Pointer(n))
		C.free(unsafe.Pointer(v))
	}

	for _, f := range s.filters {
		C.free(unsafe.Pointer(f.text))
	}

	s.cgoHandle.Delete()

	return err
}

// NewSession creates a new monitor session.
//
// At least one path must be passed.
// The callback is invoked when the monitor receives change events satisfying all the session criteria.
func NewSession(paths []string, c Callback, options ...Option) (*Session, error) {
	if !initialized {
		if err := errs[C.fsw_init_library()]; err != nil {
			return nil, err
		}

		initialized = true
	}

	opt := &opt{}
	for _, o := range options {
		if err := o(opt); err != nil {
			return nil, err
		}
	}

	s := &Session{handle: C.fsw_init_session(opt.monitorType), opt: opt}
	s.paths = make([]*C.char, len(paths))
	for i, p := range paths {
		s.paths[i] = C.CString(p)
		if err := errs[C.fsw_add_path(s.handle, s.paths[i])]; err != nil {
			return nil, err
		}
	}

	for n, v := range opt.properties {
		if err := errs[C.fsw_add_property(s.handle, n, v)]; err != nil {
			return nil, err
		}
	}

	if opt.allowOverflow {
		if err := errs[C.fsw_set_allow_overflow(s.handle, true)]; err != nil {
			return nil, err
		}
	}

	if opt.latency != 0 {
		if err := errs[C.fsw_set_latency(s.handle, opt.latency)]; err != nil {
			return nil, err
		}
	}

	if opt.recursive {
		if err := errs[C.fsw_set_recursive(s.handle, true)]; err != nil {
			return nil, err
		}
	}

	if opt.directoryOnly {
		if err := errs[C.fsw_set_directory_only(s.handle, true)]; err != nil {
			return nil, err
		}
	}

	if opt.followSymlinks {
		if err := errs[C.fsw_set_follow_symlinks(s.handle, true)]; err != nil {
			return nil, err
		}
	}

	for _, et := range opt.eventTypes {
		if err := errs[C.fsw_add_event_type_filter(s.handle, C.fsw_event_type_filter{flag: C.enum_fsw_event_flag(et)})]; err != nil {
			return nil, err
		}
	}

	for _, f := range opt.filters {
		if err := errs[C.fsw_add_filter(s.handle, f)]; err != nil {
			return nil, err
		}
	}

	h := cgo.NewHandle(s)
	s.callback = c
	s.cgoHandle = h
	C.fsw_set_callback(s.handle, C.FSW_CEVENT_CALLBACK(C.process_events), unsafe.Pointer(h)) //nolint:unsafeptr

	return s, nil
}

//export go_callback
func go_callback(events *C.fsw_cevent, eventNum C.uint, handle C.uintptr_t) {
	e := make([]Event, eventNum)

	var i C.uint
	for ; i < eventNum; i++ {
		types := make([]EventType, events.flags_num)
		var j C.uint

		flags := events.flags
		for ; j < events.flags_num; j++ {
			types[j] = EventType(*flags)

			if j < events.flags_num-1 {
				flags = (*C.enum_fsw_event_flag)(unsafe.Add(unsafe.Pointer(flags), C.sizeof_enum_fsw_event_flag))
			}
		}

		e[i] = Event{
			C.GoString(events.path),
			time.Unix(int64(events.evt_time), 0),
			types,
		}

		if i < eventNum-1 {
			events = (*C.fsw_cevent)(unsafe.Add(unsafe.Pointer(events), C.sizeof_fsw_cevent))
		}
	}

	h := cgo.Handle(handle)
	s := h.Value().(*Session)
	s.callback(e)
}
