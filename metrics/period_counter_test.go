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

func TestResetIfBeginningOfPeriod(t *testing.T) {
	Now = func() time.Time { return time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC) }

	counter := NewPeriodCounter()
	periodCounter := counter.(*PeriodCounter)
	periodCounter.PeriodType = minutelyPeriodType
	currentPeriod := periodCounter.CurrentPeriod()

	periodCounter.StoredPeriod = currentPeriod - 1
	periodCounter.ResetIfBeginningOfPeriod()

	assert.Equal(t, periodCounter.StandardCounter.count, int64(0))
	assert.Equal(t, periodCounter.StoredPeriod, currentPeriod)
	assert.True(t, periodCounter.CounterResetted)

	periodCounter.StandardCounter.count = 2
	periodCounter.ResetIfBeginningOfPeriod()
	assert.Equal(t, periodCounter.StandardCounter.count, int64(2))

	oldPeriod := currentPeriod - 2
	periodCounter.StoredPeriod = oldPeriod
	periodCounter.ResetIfBeginningOfPeriod()

	assert.Equal(t, periodCounter.StandardCounter.count, int64(0))
	assert.Equal(t, periodCounter.StoredPeriod, currentPeriod)
	assert.True(t, periodCounter.CounterResetted)

	periodCounter.StandardCounter.count = 1
	periodCounter.StoredPeriod = oldPeriod
	periodCounter.ResetIfBeginningOfPeriod()

	assert.Equal(t, periodCounter.StandardCounter.count, int64(1))
	assert.Equal(t, periodCounter.StoredPeriod, oldPeriod)
	assert.False(t, periodCounter.CounterResetted)

	periodCounter.ResetIfBeginningOfPeriod()
	assert.Equal(t, periodCounter.StandardCounter.count, int64(0))
	assert.Equal(t, periodCounter.StoredPeriod, currentPeriod)
	assert.True(t, periodCounter.CounterResetted)
}

func TestGetOrRegisterPeriodCounter(t *testing.T) {
	Now = func() time.Time { return time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC) }
	method := "testingMethod"
	counter := GetOrRegisterPeriodCounter(method, nil)
	periodCounter := counter.(*PeriodCounter)

	assert.Equal(t, periodCounter.StandardCounter.count, int64(0))
	assert.Equal(t, periodCounter.PeriodType, hourlyPeriodType)
	assert.Equal(t, periodCounter.StoredPeriod, 1)
	assert.False(t, periodCounter.CounterResetted)

	newCounter := GetOrRegisterPeriodCounter(method, nil)
	assert.Equal(t, counter, newCounter)
}
