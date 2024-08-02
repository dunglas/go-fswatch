package fswatch

// #include <libfswatch.h>
import "C"

type MonitorType C.enum_fsw_monitor_type

const (
	/* System default */
	SystemDefaultMonitor MonitorType = C.system_default_monitor_type
	/* macOS FSEvents */
	FseventsMonitor = C.fsevents_monitor_type
	/* BSD `kqueue` */
	KqueueMonitor = C.kqueue_monitor_type
	/* Linux `inotify` */
	InotifyMonitor = C.inotify_monitor_type
	/* Windows */
	WindowsMonitor = C.windows_monitor_type
	/* `stat()`-based poll  */
	PollMonitor = C.poll_monitor_type
	/* Solaris/Illumos */
	FenMonitor = C.fen_monitor_type
)
