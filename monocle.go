package monocle

import (
	"strings"
	"time"

	"github.com/dsincl12/wyrand"
)

// number of characters available
const base62 = 62

// contants for fast random character generation
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// lexicographical ordering (based on Unicode table)
const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Monocle is a struct for the configuration and the instance methods.
type Monocle struct {
	Config Config
	now    func() time.Time
	wyrand *wyrand.WyRand
}

// Config is a struct that is used to configure Monocle.
type Config struct {
	TimestampResolution      time.Duration
	NumberOfRandomCharacters int
}

// New returns a new instance of Monocle with the applied configuration.
func New(c Config) *Monocle {
	return &Monocle{
		Config: c,
		now: func() time.Time {
			return time.Now().UTC()
		},
		wyrand: wyrand.New(uint64(time.Now().UnixNano())),
	}
}

// Next returns a value in the format: [base62 encoded timestamp][random characters].
func (m *Monocle) Next() string {
	b := strings.Builder{}
	b.Grow(m.Config.NumberOfRandomCharacters * 2)

	b.Write(m.timestamp())
	b.Write(m.rand())

	return b.String()
}

// ParseTimestamp parses the value and returns the timestamp.
func (m *Monocle) ParseTimestamp(value string) time.Time {
	timestamp := value[:len(value)-m.Config.NumberOfRandomCharacters]

	return time.Unix(0, int64(decode(reverse(timestamp)))*int64(m.Config.TimestampResolution)).UTC()
}

// ParseRandomString parses the value and returns the random string.
func (m *Monocle) ParseRandomString(value string) string {
	return value[len(value)-m.Config.NumberOfRandomCharacters:]
}

func (m *Monocle) timestamp() []byte {
	t := uint64(m.now().UnixNano() / int64(m.Config.TimestampResolution))

	return encode(t)
}

func (m *Monocle) rand() []byte {
	b := make([]byte, m.Config.NumberOfRandomCharacters)

	for i, cache, remain := m.Config.NumberOfRandomCharacters-1, m.wyrand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = m.wyrand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(characters) {
			b[i] = characters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

func encode(val uint64) []byte {
	count := 0
	valCopy := val

	for valCopy > 0 {
		count++
		valCopy /= base62
	}

	b := make([]byte, count)

	for val > 0 {
		count--
		rem := val % base62
		b[count] = characters[rem]
		val /= base62
	}

	return b
}

func decode(str string) uint64 {
	var val uint64
	baseMul := uint64(1)

	for _, ch := range str {
		rem := getRemainderForChar(byte(ch))
		val += rem * baseMul
		baseMul *= base62
	}

	return val
}

func getRemainderForChar(ch byte) uint64 {
	if ch >= '0' && ch <= '9' {
		return (uint64(ch) % uint64('0'))
	}

	if ch >= 'A' && ch <= 'Z' {
		return (uint64(ch) % uint64('A')) + 10
	}

	if ch >= 'a' && ch <= 'z' {
		return (uint64(ch) % uint64('a')) + 36
	}

	return 0
}

func reverse(s string) string {
	b := []byte(s)

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
