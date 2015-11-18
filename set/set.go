// Package set provides the functions for creating and useing a set data
// structure which has no duplicates
package set

import (
	"fmt"

	"github.com/matzhouse/go-data-structures/list"
)

// Set implements a structure where only a single copy of an element may
// exist in a set. Set has a list embedded in it and can therefore use the
// list methods, including contains.
type Set struct {
	*list.List
}

// NewSet returns a new initialised set value
func NewSet() *Set {

	return &Set{
		List: list.NewList(),
	}

}

// Add takes a value and adds it to the Set so long as it doesn't exist.
func (s *Set) Add(value interface{}) (err error) {

	if s.List.Contains(value) != nil {
		return fmt.Errorf("element already in the list")
	}

	s.List.InsertAtBeginning(value)

	return nil

}
