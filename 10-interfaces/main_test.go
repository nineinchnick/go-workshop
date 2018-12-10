package main

import (
	"testing"
	"time"
)

type fakeClock struct {
	time time.Time
}

func (f *fakeClock) now() time.Time {
	return f.time
}

func TestIsFridaySuccess(t *testing.T) {
	// Arrange
	date := time.Date(2018, 12, 11, 0, 0, 0, 0, time.UTC)
	clock := fakeClock{date}

	// Act
	result := isFriday(&clock)

	// Assert
	if result {
		t.Errorf("%s is NOT Friday", date)
	}
}

func TestIsFridayFail(t *testing.T) {
	// Arrange
	date := time.Date(2018, 12, 14, 0, 0, 0, 0, time.UTC)
	clock := fakeClock{date}

	// Act
	result := isFriday(&clock)

	// Assert
	if !result {
		t.Errorf("%s is Friday", date)
	}
}
