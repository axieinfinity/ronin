package metrics

import (
	"time"
)

const (
	minutelyPeriodType = iota
	hourlyPeriodType   = iota
	dailyPeriodType    = iota
	monthlyPeriodType  = iota
)

var Now = time.Now

type PeriodCounter struct {
	*StandardCounter
	PeriodType      int
	StoredPeriod    int
	CounterResetted bool
}

func (c *PeriodCounter) SetPeriodType(t int) {
	c.PeriodType = t
}

func (c *PeriodCounter) Inc(i int64) {
	c.ResetIfBeginningOfPeriod()
	c.StandardCounter.Inc(i)
}

func (c *PeriodCounter) Dec(i int64) {
	c.ResetIfBeginningOfPeriod()
	c.StandardCounter.Dec(i)
}

// Clear sets the counter to zero.
func (c *PeriodCounter) Clear() {
	c.StandardCounter.Clear()
}

// Count returns the current count.
func (c *PeriodCounter) Count() int64 {
	return c.StandardCounter.Count()
}

// Snapshot returns a read-only copy of the counter.
func (c *PeriodCounter) Snapshot() Counter {
	return c.StandardCounter.Snapshot()
}

// if current time is a new period compared to stored period, and counter is not resetted yet
// then reset the counter
func (c *PeriodCounter) ResetIfBeginningOfPeriod() {
	if c.CurrentPeriod() != c.StoredPeriod && !c.CounterResetted {
		c.CounterResetted = true
		c.StoredPeriod = c.CurrentPeriod()
		c.Clear()
	} else {
		c.CounterResetted = false
	}
}

// return the period value of current time.
// For example, if period type is minutelyPeriodType and current time is 12:03,
// it should return 3. If period type is monthlyPeriodType and current datetime is 3/12
// it should return 12
func (c *PeriodCounter) CurrentPeriod() int {
	switch c.PeriodType {
	case minutelyPeriodType:
		return Now().Minute()
	case hourlyPeriodType:
		return Now().Hour()
	case dailyPeriodType:
		return Now().Day()
	case monthlyPeriodType:
		return int(Now().Month())
	default:
		panic("invalid period type")
	}
}

// NewPeriodCounter constructs a new PeriodCounter.
func NewPeriodCounter() Counter {
	if !Enabled {
		return NilCounter{}
	}
	return &PeriodCounter{&StandardCounter{0}, hourlyPeriodType, 0, false}
}

// GetOrRegisterPeriodCounter returns an existing Counter or constructs and registers
// a new PeriodCounter.
func GetOrRegisterPeriodCounter(name string, r Registry) Counter {
	if nil == r {
		r = DefaultRegistry
	}
	return r.GetOrRegister(name, NewPeriodCounter).(Counter)
}
