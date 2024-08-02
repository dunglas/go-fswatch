#include "fswatch.h"
#include "_cgo_export.h"
#include <stdint.h>

void process_events(fsw_cevent const *const events,
                    const unsigned int event_num, void *data) {
  go_callback((fsw_cevent *)events, event_num, (uintptr_t)data);
}