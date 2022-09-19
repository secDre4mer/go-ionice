package ionice

import (
	"errors"
	"fmt"
)

// Unfortunately, setiopolicy_np is not available in the syscall pkg, so we need to use CGO instead.

// #include <sys/resource.h>
import "C"

var priorityMapping = map[Niceness]C.int{
	VeryLow:  C.IOPOL_THROTTLE,
	Low:      C.IOPOL_THROTTLE,
	Standard: C.IOPOL_STANDARD,
	High:     C.IOPOL_IMPORTANT,
	VeryHigh: C.IOPOL_IMPORTANT,
}

func setIoPriority(niceness Niceness) error {
	priority, isValid := priorityMapping[niceness]
	if !isValid {
		return errors.New("invalid niceness specified")
	}

	r1, err := C.setiopolicy_np(C.IOPOL_TYPE_DISK, C.IOPOL_SCOPE_PROCESS, priority)
	if r1 != 0 {
		return err
	}
	return nil
}

var inversePriorityMapping = map[C.int]Niceness{
	C.IOPOL_DEFAULT:   Standard,
	C.IOPOL_PASSIVE:   VeryLow,
	C.IOPOL_THROTTLE:  Low,
	C.IOPOL_STANDARD:  Standard,
	C.IOPOL_IMPORTANT: High,
}

func getIoPriority() (Niceness, error) {
	r1, err := C.getiopolicy_np(C.IOPOL_TYPE_DISK, C.IOPOL_SCOPE_PROCESS)
	if r1 == -1 {
		return 0, err
	}
	niceness, isValid := inversePriorityMapping[r1]
	if !isValid {
		return 0, fmt.Errorf("unrecognized IO policy: %d", int(r1))
	}
	return niceness, nil
}
