package list

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"
)

// List stores a reference to the Root node of a linked list
type List struct {
	Root *Element

	mu sync.Mutex // Mutex to protect the list on an insert
}

// Element stores a value in the list with a reference to the next item
type Element struct {
	Value interface{}
	Next  *Element
}

// NewList returns an empty list
func NewList() *List {
	return &List{}
}

// NewListWithValue takes an initial value and returns a list value with the value set
// as the Root node
func NewListWithValue(value interface{}) *List {

	return &List{
		Root: &Element{
			Value: value,
		},
	}

}

// NewListFromGoList takes an initial value and returns a list value with the value set
// as the Root node
func NewListFromGoList(values *list.List) (*List, error) {

	// Get the first value from the input list
	f := values.Front()

	// Create a new list
	l := NewListWithValue(f.Value)

	// We gotta start somewhere..
	n := l.Root

	// Iterate through list and print its contents.
	for el := f.Next(); el != nil; el = el.Next() {

		// Create a new Element with the appropriate value
		n.Next = &Element{
			Value: el.Value,
		}

		// Move along out list in time with the input list
		n = n.Next
	}

	return l, nil

}

// NewListFromSlice takes an initial value and returns a list value with the value set
// as the Root node
func NewListFromSlice(values []interface{}) *List {

	// Make a list with the last value in the slice
	l := NewListWithValue(values[0])

	n := l.Root
	// As we loop through the slice we can build out the list
	for i := 1; i < len(values)-1; i++ {
		n.Next = &Element{
			Value: values[i],
		}
		n = n.Next
	}

	return l

}

// InsertAtBeginning takes a value and adds it to the beginning of the list referenced
// in the pointer receiver
func (l *List) InsertAtBeginning(value interface{}) {

	// Create a new element with the current root as its next
	el := &Element{
		Value: value,
		Next:  l.Root,
	}

	// Set the new element as the root
	l.Root = el

}

// InsertAtEnd takes a value and adds it to the end of the list referenced
// in the pointer receiver. This is a simple loop to find the end
func (l *List) InsertAtEnd(value interface{}) {

	nel := &Element{
		Value: value,
	}

	if l.Root == nil {
		l.Root = nel
		return
	}

	// Traverse the list until we find the current end element, at which point we
	// can add a new one
	for n := l.Root; n.Next != nil; n = n.Next {

		if n.Next == nil {
			// set the new element
			n.Next = nel
		}

	}

}

// InsertAfter takes a value and adds it to the list referenced in the pointer receiver
func (l *List) InsertAfter(el *Element, value interface{}) error {

	if el == nil {
		return fmt.Errorf("element cannot be nil")
	}

	// Insert the new element into the list, and set it's next as the next element
	// from the element we're inserting after.
	nel := &Element{
		Value: value,
		Next:  el.Next,
	}

	// Join up the original element to the new one
	el.Next = nel

	return nil

}

// Length returns an integer value of the number of elements in the list
func (l *List) Length() (len int) {

	// len should be 0
	len = 1

	// We start at the Root of the list
	n := l.Root

	// Traverse the list until we find the current end element, at which point we
	// can add a new one
	for n != nil {
		if n.Next != nil {
			// Add a value of 1 to the count
			len++
		}
		n = n.Next
	}

	return len

}

// Contains takes a value and returns a boolean based on whether the value
// appears in the list or not
func (l *List) Contains(v interface{}) *Element {

	// check for empty list
	if l.Root != nil {

		el := &Element{
			Value: v,
		}

		// Traverse the list and check each element for comparison with the supplied
		// element input
		for n := l.Root; n != nil; n = n.Next {
			if reflect.DeepEqual(n.Value, el.Value) {
				return n
			}
		}

	}

	return nil

}
