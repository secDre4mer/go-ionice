package ionice

type Niceness int

const (
	VeryLow Niceness = iota
	Low
	Standard
	High
	VeryHigh
)

func SetIoPriority(niceness Niceness) error {
	return setIoPriority(niceness)
}

func GetIoPriority() (Niceness, error) {
	return getIoPriority()
}
