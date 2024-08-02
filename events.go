package fswatch

// #include <libfswatch.h>
import "C"
import "time"

type EventType C.enum_fsw_event_flag

var (
	/* No event has occurred. */
	NoOp EventType = C.NoOp
	/* Platform-specific placeholder for event type that cannot currently be mapped. */
	PlatformSpecific EventType = C.PlatformSpecific
	/* An object was created. */
	Created EventType = C.Created
	/* An object was updated. */
	Updated EventType = C.Updated
	/* An object was removed. */
	Removed EventType = C.Removed
	/* An object was renamed. */
	Renamed EventType = C.Renamed
	/* The owner of an object was modified. */
	OwnerModified EventType = C.OwnerModified
	/* The attributes of an object were modified. */
	AttributeModified EventType = C.AttributeModified
	/* An object was moved from this location. */
	MovedFrom EventType = C.MovedFrom
	/* An object was moved to this location. */
	MovedTo EventType = C.MovedTo
	/* The object is a file. */
	IsFile EventType = C.IsFile
	/* The object is a directory. */
	IsDir EventType = C.IsDir
	/* The object is a symbolic link. */
	IsSymLink EventType = C.IsSymLink
	/* The link count of an object has changed. */
	Link EventType = C.Link
	/* The event queue has overflowed. */
	Overflow EventType = C.Overflow
)

type Event struct {
	Path  string
	Time  time.Time
	Types []EventType
}
