package clock

import (
	"fmt"
)

// Clock holding time as hh:mm
type Clock struct {
	m int
	h int
}

// New clock
func New(h int, m int) Clock {

	for m < 0 {
		m += 60
		h--
	}

	h += m / 60
	for h < 0 {
		h += 24
	}

	return Clock{h: h % 24, m: m % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add minutes to clock
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

// Subtract minutes from clock
func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}
