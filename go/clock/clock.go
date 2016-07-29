package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	c := Clock{hour, minute}
	return c.fix()
}

func (c Clock) fix() Clock {
	if c.minute >= 60 || c.minute < 0 {
		c.hour += c.minute / 60
		c.minute %= 60
		if c.minute < 0 {
			c.minute += 60
			c.hour -= 1
		}
	}
	c.hour %= 24
	if c.hour < 0 {
		c.hour += 24
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.fix()
}
