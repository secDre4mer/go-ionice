package ionice

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

// IO priority values from https://github.com/processhacker/phnt
const (
	ioPriorityVeryLow  = iota // Defragging, content indexing and other background I/Os.
	ioPriorityLow             // Prefetching for applications.
	ioPriorityNormal          // Normal I/Os.
	ioPriorityHigh            // Used by filesystems for checkpoint I/O.
	ioPriorityCritical        // Used by memory manager. Not available for applications.
)

var priorityMapping = map[Niceness]uint32{
	VeryLow:  ioPriorityVeryLow,
	Low:      ioPriorityLow,
	Standard: ioPriorityNormal,
	High:     ioPriorityHigh,
	VeryHigh: ioPriorityHigh, // ioPriorityCritical is not available from user space
}

func setIoPriority(niceness Niceness) error {
	ioPrio, isValid := priorityMapping[niceness]
	if !isValid {
		return errors.New("invalid niceness specified")
	}
	return windows.NtSetInformationProcess(windows.CurrentProcess(), windows.ProcessIoPriority, unsafe.Pointer(&ioPrio), uint32(unsafe.Sizeof(ioPrio)))
}

var inversePriorityMapping = map[uint32]Niceness{
	ioPriorityVeryLow:  VeryLow,
	ioPriorityLow:      Low,
	ioPriorityNormal:   Standard,
	ioPriorityHigh:     High,
	ioPriorityCritical: VeryHigh,
}

func getIoPriority() (Niceness, error) {
	var ioPrio uint32
	var returnLength uint32
	err := windows.NtQueryInformationProcess(windows.CurrentProcess(), windows.ProcessIoPriority, unsafe.Pointer(&ioPrio), uint32(unsafe.Sizeof(ioPrio)), &returnLength)
	if err != nil {
		return 0, err
	}
	return inversePriorityMapping[ioPrio], nil
}
