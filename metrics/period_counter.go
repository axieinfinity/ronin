package metrics

import (
	"sync/atomic"
	"time"
)

const (
	minutelyPeriodType = iota
	hourlyPeriodType   = iota
	dailyPeriodType    = iota
	monthlyPeriodType  = iota
)

type PeriodCounter struct {
	StandardCounter
	PeriodType      int
	StoredPeriod    int
	CounterResetted bool
}

func (c *PeriodCounter) Inc(i int64) {
	c.resetIfBeginningOfPeriod()
	atomic.AddInt64(&c.count, i)
}

func (c *PeriodCounter) Dec(i int64) {
	c.resetIfBeginningOfPeriod()
	atomic.AddInt64(&c.count, -i)
}

func (c *PeriodCounter) resetIfBeginningOfPeriod() {
	if c.currentPeriod() != c.StoredPeriod && !c.CounterResetted {
		c.CounterResetted = true
		c.StoredPeriod = c.currentPeriod()
		c.Clear()
	} else {
		c.CounterResetted = false
	}
}

func (c *PeriodCounter) currentPeriod() int {
	switch c.PeriodType {
	case minutelyPeriodType:
		return time.Now().Minute()
	case hourlyPeriodType:
		return time.Now().Hour()
	case dailyPeriodType:
		return time.Now().Day()
	case monthlyPeriodType:
		return int(time.Now().Month())
	default:
		panic("invalid period type")
	}
}
