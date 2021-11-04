package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInterval(t *testing.T) {
	assert := assert.New(t)

	var interval *Interval
	assert.True(interval.IsEmpty(), "empty Interval")

	interval = NewInterval(0, 1)
	assert.False(interval.IsEmpty(), "non-empty Interval")
}

func TestEmptyIntervalSet(t *testing.T) {
	assert := assert.New(t)

	var intervalset *IntervalSet
	assert.True(intervalset.IsEmpty(), "empty IntervalSet")

	intervalset = NewIntervalSet()
	assert.True(intervalset.IsEmpty(), "empty IntervalSet")

	var interval *Interval
	intervalset.AddInterval(interval)
	assert.True(intervalset.IsEmpty(), "empty IntervalSet")

	intervalset.AddInterval(NewInterval(0, 1))
	assert.False(intervalset.IsEmpty(), "non-empty IntervalSet")
}
