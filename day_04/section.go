package main

type Section struct {
	start int
	end   int
}

func (s Section) Contains(other Section) bool {
	return s.start <= other.start && s.end >= other.end
}

func (s Section) Overlaps(other Section) bool {
	overlapsEarlier := func(left Section, right Section) bool {
		return left.start <= right.start && left.end >= right.start
	}
	return overlapsEarlier(s, other) || overlapsEarlier(other, s)
}
