package processor

type Processor struct {
	le *levenshtein
}

type Line struct {
	Input []byte
	IsNew bool
}

func NewProcessor(maxDiffRatio float64) *Processor {
	le := newLevenshtein(maxDiffRatio)
	return &Processor{le: le}
}

func (p *Processor) Process(in Line) Line {
	in = p.le.process(in)
	return in
}
