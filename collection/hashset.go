package collection

type HashSet struct {
	data map[interface{}]struct{}
}

// NewHashSet create new set instance
func NewHashSet() HashSet {
	return HashSet{
		data: make(map[interface{}]struct{}),
	}
}

// HashSetFrom new set from multiple elements
func HashSetFrom(es ...interface{}) HashSet {
	hs := HashSet{data: make(map[interface{}]struct{}, len(es))}
	for _, e := range es {
		hs.data[e] = struct{}{}
	}

	return hs
}

// Add new element to the set.
func (h *HashSet) Add(e interface{}) {
	h.data[e] = struct{}{}
}

// Remove element from the set.
func (h *HashSet) Remove(e interface{}) {
	delete(h.data, e)
}

// Pop pick some random element of the set,
// remove it from the set and return it.
func (h *HashSet) Pop() (interface{}, bool) {
	var e interface{} = nil
	for v := range h.data {
		e = v
		break
	}

	if e != nil {
		delete(h.data, e)
		return e, true
	}

	return e, false
}

// Find try to find element in the set.
func (h HashSet) Find(e interface{}) bool {
	_, found := h.data[e]
	return found
}

// Len return total number of elements in the set.
func (h HashSet) Len() int {
	return len(h.data)
}

// IsEmpty check empty set
func (h HashSet) IsEmpty() bool {
	return h.Len() == 0
}

// List return slice of set elements.
func (h HashSet) List() []interface{} {
	flatten := make([]interface{}, 0, len(h.data))
	for v := range h.data {
		flatten = append(flatten, v)
	}

	return flatten
}

// Difference between two sets: (only_in_s1, only_in_s2).
func Difference(s1, s2 HashSet) (HashSet, HashSet) {
	var (
		onlyS1 = NewHashSet()
		onlyS2 = NewHashSet()
	)

	if s1.IsEmpty() && s2.IsEmpty() {
		return onlyS1, onlyS2
	} else if s1.IsEmpty() {
		return onlyS1, s2
	} else if s2.IsEmpty() {
		return s1, onlyS2
	}

	for v1 := range s1.data {
		if !s2.Find(v1) {
			onlyS1.Add(v1)
		}
	}

	for v2 := range s2.data {
		if !s1.Find(v2) {
			onlyS2.Add(v2)
		}
	}

	return onlyS1, onlyS2
}

// Intersection return new set contains
// equals elements from two sets
func Intersection(s1, s2 HashSet) HashSet {
	both := NewHashSet()
	for v1 := range s1.data {
		if s2.Find(v1) {
			both.Add(v1)
		}
	}

	return both
}

// Union combine arbitrary number of sets into single one.
// Warning: in order to make sense, values
// in the union must be of the same type!
func Union(sets ...HashSet) HashSet {
	union := NewHashSet()
	for _, s := range sets {
		for v := range s.data {
			union.Add(v)
		}
	}

	return union
}

// Equal compare two set on equals
func Equal(s1, s2 HashSet) bool {
	l, r := Difference(s1, s2)
	return l.IsEmpty() && r.IsEmpty()
}
