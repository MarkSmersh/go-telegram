package cli

import (
	"strconv"
)

type Option struct {
	Prefix string
	Value  string
}

// Converts option value into an int value. Returs error if value is inconvertable.
func (o Option) Atoi() (int, bool) {
	v, err := strconv.Atoi(o.Value)
	if err != nil {
		return v, false
	} else {
		return v, true
	}
}

// Checks is value exists, but empty - string equals "".
func (o Option) IsEmpty() bool {
	if o.Value == "" {
		return true
	}
	return false
}

// Checks is value in range of min and max. Returns bool
func (o Option) IsInRange(min int, max int) bool {
	_, ok := o.AtoiRange(min, max)

	if !ok {
		return false
	} else {
		return true
	}
}

// Checks is value in range of min and max.
func (o Option) AtoiRange(min int, max int) (int, bool) {
	v, ok := o.Atoi()
	wrong := false

	if !ok {
		wrong = true
	}

	if v < min {
		wrong = true
	}

	if v > max {
		wrong = true
	}

	if wrong {
		return 0, false
	}

	return v, true
}

// Checks are conditions true. If one is not returns position of that condition. If all are true returns -1
func (o Option) Is(conditions ...bool) int {
	for i, c := range conditions {
		if !c {
			return i
		}
	}

	return -1
}

func (o Option) Exists() bool {
	if len(o.Prefix) <= 0 {
		return false
	} else {
		return true
	}
}
