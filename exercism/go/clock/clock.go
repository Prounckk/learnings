package clock

import "fmt"

type Clock struct{ h, m int }

func New(h, m int) Clock {
	c := Clock{h, m}
	return c.normilizer()
}

func (c Clock) Add(m int) Clock {
	c.m += m
	return c.normilizer()
}

func (c Clock) Subtract(m int) Clock {
	c.m -= m
	return c.normilizer()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) normilizer() Clock {
	if c.m >= 60 {
		c.h += c.m / 60 // hours
		c.m %= 60       // minutes
	}
	if c.m < 0 {
		c.h--
		c.m += 60
	}
	if c.h >= 24 { // hours
		c.h %= 24 // hours
	}
	if c.h < 0 { // hours	// negative hours
		c.h = 24 + c.h // hours
	}
	if c.h < 0 || c.m > 60 || c.m < 0 {
		return c.normilizer()
	}

	return c
}
