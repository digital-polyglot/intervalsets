package domain

type Interval struct {
	Start uint64
	Stop  uint64
}

type IntervalSet struct {
	Intervals []Interval
}

func NewInterval(start, stop uint64) *Interval {
	return &Interval{
		Start: start,
		Stop:  stop,
	}
}

func NewIntervalSet() *IntervalSet {
	return &IntervalSet{}
}

func (iv *Interval) IsEmpty() bool {
	return iv == nil
}

func (ivs *IntervalSet) IsEmpty() bool {
	if ivs == nil || len(ivs.Intervals) == 0 {
		return true
	}
	return false
}

func (ivs *IntervalSet) AddInterval(r *Interval) *IntervalSet {
	if r == nil {
		return ivs
	}
	ivs.Intervals = append(ivs.Intervals, *r)
	return ivs
}
