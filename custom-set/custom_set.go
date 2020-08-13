package stringset

import "strings"

// Set is a collection of unique string values.
type Set map[string]struct{}

// New creates a new set.
func New() Set {
	return make(map[string]struct{})
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

	items := []string{}
	for x := range s {
		items = append(items, "\""+x+"\"")
	}
	stringifiedSet += strings.Join(items, ", ")
	stringifiedSet += "}"

	return stringifiedSet
}

// IsEmpty checks is the set is empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has checks for a key in the set
func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

// Subset returns true if all of its elements are contained in the other set
func Subset(s1, s2 Set) bool {
	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

// Disjoint returns true are disjoint if they share no elements
func Disjoint(s1, s2 Set) bool {
	for key := range s1 {
		if s2.Has(key) {
			return false
		}
	}
	return true
}

// Equal returns true if all elements are in both sets
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

// Add element to Set
func (s Set) Add(key string) {
	s[key] = struct{}{}
}

// Intersection returns a set of all shared elements
func Intersection(s1, s2 Set) Set {
	matched := New()
	for key := range s1 {
		if s2.Has(key) {
			matched.Add(key)
		}
	}

	return matched
}

// Difference (or Complement) of a set is a set of all elements that are only in the first set
func Difference(s1, s2 Set) Set {
	matched := New()
	for key := range s1 {
		if !s2.Has(key) {
			matched.Add(key)
		}
	}

	return matched
}

// Union returns a set of all elements in either set
func Union(s1, s2 Set) Set {
	union := New()

	for key := range s1 {
		union.Add(key)
	}

	for key := range s2 {
		union.Add(key)
	}
	return union
}
