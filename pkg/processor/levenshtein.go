package processor

import (
	levenshtein2 "github.com/psykhi/uno/pkg/levenshtein"
	"math"
)

type levenshtein struct {
	seen         [][]string
	maxDiffRatio float64
}

func newLevenshtein(maxDiffRatio float64) *levenshtein {
	seen := make([][]string, 0)
	return &levenshtein{seen: seen, maxDiffRatio: maxDiffRatio}
}

func (le *levenshtein) process(in Line) Line {
	in.IsNew = true
	for _, l := range le.seen {
		maxDiff := int(math.Ceil(float64(len(in.Tokens)) * le.maxDiffRatio))
		d := levenshtein2.LevenshteinDistanceK(in.Tokens, l, nil, nil, maxDiff)
		if d <= maxDiff && d >= 0 {
			in.IsNew = false
			return in
		}
	}
	le.seen = append(le.seen, in.Tokens)
	return in
}
