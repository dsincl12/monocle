package monocle

import (
	"strings"
	"testing"
	"time"
)

func TestNextTimestampResolutionNanosecond(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Nanosecond,
		NumberOfRandomCharacters: 0,
	})

	length := 11
	value := m.Next()

	if len(value) != length {
		t.Errorf("expected length %d, got %d", length, len(value))
	}
}

func TestNextTimestampResolutionMillisecond(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Millisecond,
		NumberOfRandomCharacters: 0,
	})

	length := 7
	value := m.Next()

	if len(value) != length {
		t.Errorf("expected length %d, got %d", length, len(value))
	}
}

func TestNextTimestampResolutionSecond(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Second,
		NumberOfRandomCharacters: 0,
	})

	length := 6
	value := m.Next()

	if len(value) != length {
		t.Errorf("expected length %d, got %d", length, len(value))
	}
}

func TestNextTimestampResolutionMinute(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Minute,
		NumberOfRandomCharacters: 0,
	})

	length := 5
	value := m.Next()

	if len(value) != length {
		t.Errorf("expected length %d, got %d", length, len(value))
	}
}

func TestNextValuesShouldBeSortable(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Millisecond,
		NumberOfRandomCharacters: 0,
	})

	value1 := m.Next()
	value2 := m.Next()

	if !strings.HasPrefix(value1, value2[:3]) {
		t.Errorf("expected %s to begin with %s", value1, value2[:3])
	}
}

func TestParseTimestampNanosecond(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Nanosecond,
		NumberOfRandomCharacters: 0,
	})

	m.now = func() time.Time {
		return time.Date(2019, 6, 3, 20, 34, 58, 651387237, time.UTC)
	}

	value := m.Next()
	timestamp := m.ParseTimestamp(value)

	if timestamp != m.now() {
		t.Errorf("expected %v, got %v", m.now(), timestamp)
	}
}

func TestParseTimestampMillisecond(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Millisecond,
		NumberOfRandomCharacters: 0,
	})

	m.now = func() time.Time {
		return time.Date(2019, 6, 3, 20, 34, 58, int(651*time.Millisecond), time.UTC)
	}

	value := m.Next()
	timestamp := m.ParseTimestamp(value)

	if timestamp != m.now() {
		t.Errorf("expected %v, got %v", m.now(), timestamp)
	}
}

func TestParseTimestampSecond(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Second,
		NumberOfRandomCharacters: 0,
	})

	m.now = func() time.Time {
		return time.Date(2019, 6, 3, 20, 34, 58, 0, time.UTC)
	}

	value := m.Next()
	timestamp := m.ParseTimestamp(value)

	if timestamp != m.now() {
		t.Errorf("expected %v, got %v", m.now(), timestamp)
	}
}

func TestParseTimestampMinute(t *testing.T) {
	m := New(Config{
		TimestampResolution:      time.Minute,
		NumberOfRandomCharacters: 0,
	})

	m.now = func() time.Time {
		return time.Date(2019, 6, 3, 20, 34, 0, 0, time.UTC)
	}

	value := m.Next()
	timestamp := m.ParseTimestamp(value)

	if timestamp != m.now() {
		t.Errorf("expected %v, got %v", m.now(), timestamp)
	}
}

func TestParseRandStringLength4(t *testing.T) {
	length := 4

	m := New(Config{
		TimestampResolution:      time.Millisecond,
		NumberOfRandomCharacters: length,
	})

	value := m.Next()
	randString := m.ParseRandomString(value)

	if len(randString) != length {
		t.Errorf("expected length %d, got %d", length, len(randString))
	}
}

func TestParseRandStringLength12(t *testing.T) {
	length := 12

	m := New(Config{
		TimestampResolution:      time.Millisecond,
		NumberOfRandomCharacters: length,
	})

	value := m.Next()
	randString := m.ParseRandomString(value)

	if len(randString) != length {
		t.Errorf("expected length %d, got %d", length, len(randString))
	}
}

func BenchmarkNext(b *testing.B) {
	b.ReportAllocs()

	m := New(Config{
		TimestampResolution:      time.Millisecond,
		NumberOfRandomCharacters: 8,
	})

	for i := 0; i < b.N; i++ {
		m.Next()
	}
}
