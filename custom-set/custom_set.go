package stringset

import "fmt"

// Set is a collection of unique string values.
type Set struct {
	data map[string]interface{}
}

// New creates a new set.
func New() Set {
	return Set{data: make(map[string]interface{})}
}

// NewFromSlice creates a new set from a slice of strings.
func NewFromSlice(slice []string) Set {
	set := New()
	for _, x := range slice {
		set.Add(x)
	}

	return set
}

// String returns a string representation of the set.
func (s Set) String() string {
	stringifiedSet := "{"

	length := len(s.data) - 1
	index := 0
	for x := range s.data {
		x = fmt.Sprintf("\"%s\"", x)
		if length != 0 && index != length {
			x += ", "
		}
		stringifiedSet += x
		index++
	}
	stringifiedSet += "}"
	return stringifiedSet
}

// IsEmpty checks is the set is empty
func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

// Has checks for a key in the set
func (s Set) Has(key string) bool {
	return s.data[key] != nil
}

// Subset returns true if all of its elements are contained in the other set
func Subset(s1, s2 Set) bool {
	for key := range s1.data {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

// Disjoint returns true are disjoint if they share no elements
func Disjoint(s1, s2 Set) bool {
	for key := range s1.data {
		if s2.Has(key) {
			return false
		}
	}
	return true
}

// Equal returns true if all elements are in both sets
func Equal(s1, s2 Set) bool {
	if len(s1.data) != len(s2.data) {
		return false
	}

	for key := range s1.data {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

// Add element to Set
func (s Set) Add(key string) {
	s.data[key] = key
}

// Intersection returns a set of all shared elements
func Intersection(s1, s2 Set) Set {
	matched := New()
	for key := range s1.data {
		if s2.Has(key) {
			matched.Add(key)
		}
	}

	return matched
}

// Difference (or Complement) of a set is a set of all elements that are only in the first set
func Difference(s1, s2 Set) Set {
	matched := New()
	for key := range s1.data {
		if !s2.Has(key) {
			matched.Add(key)
		}
	}

	return matched
}

// Union returns a set of all elements in either set
func Union(s1, s2 Set) Set {
	union := New()

	for key := range s1.data {
		union.Add(key)
	}

	for key := range s2.data {
		union.Add(key)
	}
	return union
}
