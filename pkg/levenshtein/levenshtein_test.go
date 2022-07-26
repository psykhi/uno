package levenshtein

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type seq struct {
	chars []byte
}

func (s *seq) Val(i int) interface{} {
	return s.chars[i]
}

func (s *seq) Len() int {
	return len(s.chars)
}

func TestDistanceK(t *testing.T) {
	d := LevenshteinDistanceK([]byte("kitten"), []byte("sitting"), nil, nil, 3)
	assert.Equal(t, 3, d)

	d = LevenshteinDistanceK([]byte("kitten"), []byte("sitting"), nil, nil, 2)
	assert.Equal(t, -1, d)

	d = LevenshteinDistanceK([]byte("elephant"), []byte("hippo"), nil, nil, 8)
	assert.Equal(t, 7, d)

	d = LevenshteinDistanceK([]byte("elephant"), []byte(""), nil, nil, 100)
	assert.Equal(t, 8, d)
}

func TestDistance(t *testing.T) {
	d := LevenshteinDistance([]byte("kitten"), []byte("sitting"), nil, nil)
	assert.Equal(t, 3, d)
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	b.Run("6 char string", func(b *testing.B) {
		x := []byte("ABCDEF")
		y := []byte("ABCCDEF")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, nil, nil)
		}
	})
	b.Run("25-30 char string", func(b *testing.B) {
		x := []byte("This is a longer string")
		y := []byte("This is a much  longer string")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, nil, nil)
		}
	})
	b.Run("Long log line", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, nil, nil)
		}
	})
	b.Run("Long log line K bound", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistanceK(x, y, nil, nil, len(x)/4)
		}
	})
	b.Run("Long log line buffer reuse", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")
		n := len(y)
		v0 := make([]int, n+1, n+1)
		v1 := make([]int, n+1, n+1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, v0, v1)
		}
	})

}
