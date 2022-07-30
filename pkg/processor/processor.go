package processor

import "github.com/blevesearch/segment"

type Processor struct {
	le *levenshtein
}

type Line struct {
	Input  []byte
	Tokens []string
	IsNew  bool
}

func NewProcessor(maxDiffRatio float64) *Processor {
	le := newLevenshtein(maxDiffRatio)
	return &Processor{le: le}
}

func (p *Processor) Process(in Line) Line {
	s := segment.NewSegmenterDirect(in.Input)
	in.Tokens = make([]string, 0)
	for s.Segment() {
		t := s.Text()
		// Number tokens will always be considered equal, a simple way to ignore timestamps, ids etc.
		if s.Type() == segment.Number {
			t = "*"
		}
		in.Tokens = append(in.Tokens, t)
	}
	in = p.le.process(in)
	return in
}
