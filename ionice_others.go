//go:build !windows && !darwin && !linux

package ionice

import (
	"errors"
)

func setIoPriority(niceness Niceness) error {
	return errors.New("IO priority change not supported on this OS")
}

func getIoPriority() (Niceness, error) {
	return 0, errors.New("IO priority query not supported on this OS")
}
