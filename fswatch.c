#include "fswatch.h"
#include "_cgo_export.h"

static inline void process_events(fsw_cevent const *const events,
                                  const unsigned int event_num, void *data) {
  go_callback((fsw_cevent *)events, event_num, (uintptr_t)data);
}

void set_callback(const FSW_HANDLE handle, uintptr_t data) {
  fsw_set_callback(handle, process_events, (void *)data);
}
