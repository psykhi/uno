package levenshtein

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistanceK(t *testing.T) {
	d := LevenshteinDistanceK([]byte("kitten"), []byte("sitting"), nil, nil, 3)
	assert.Equal(t, 3, d)

	d = LevenshteinDistanceK([]byte("kitten"), []byte("sitting"), nil, nil, 2)
	assert.Equal(t, -1, d)

	d = LevenshteinDistanceK([]byte("elephant"), []byte("hippo"), nil, nil, 8)
	assert.Equal(t, 7, d)

	d = LevenshteinDistanceK([]byte("elephant"), []byte(""), nil, nil, 100)
	assert.Equal(t, 8, d)

	d = LevenshteinDistanceK([]string{"hello", "world"}, []string{"hello", "earth"}, nil, nil, 100)
	assert.Equal(t, 1, d)
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	b.Run("Long log line K bound", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistanceK(x, y, nil, nil, len(x)/4)
		}
	})

}
