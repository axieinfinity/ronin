package metrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCurrentPeriod(t *testing.T) {
	Now = func() time.Time { return time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC) }

	counter := NewPeriodCounter()
	periodCounter := counter.(*PeriodCounter)
	periodCounter.PeriodType = minutelyPeriodType
	assert.Equal(t, periodCounter.CurrentPeriod(), 10)

	periodCounter.PeriodType = hourlyPeriodType
	assert.Equal(t, periodCounter.CurrentPeriod(), 1)

	periodCounter.PeriodType = dailyPeriodType
	assert.Equal(t, periodCounter.CurrentPeriod(), 21)

	periodCounter.PeriodType = monthlyPeriodType
	assert.Equal(t, periodCounter.CurrentPeriod(), 2)
}
