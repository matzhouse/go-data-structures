package list

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_NewList(t *testing.T) {

	if l := NewList(); l == nil {
		t.Fail()
	}

}

func Test_NewListWithValue(t *testing.T) {

	l := NewListWithValue("test value")

	if l.Root == nil {
		t.Fail()
	}

	if v := l.Root.Value; v.(string) != "test value" {
		t.Fail()
	}

}

func Test_NewListFromSlice(t *testing.T) {

	s := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	l := NewListFromSlice(s)

	n := l.Root
	c := 0

	// Loop over both data structures and compare the result
	for n.Next != nil {
		if reflect.DeepEqual(n.Value, s[c]) == false {
			t.Log(s[c], "=", n.Value)

			t.Log("interface is not the same")
			t.Fail()
		}

		c++
		n = n.Next
	}

}

func Test_NewListFromGoList(t *testing.T) {

	ol := list.New()
	for i := 1; i <= 10; i++ {
		ol.PushBack(i)
	}

	l, err := NewListFromGoList(ol)

	if err != nil {
		t.Fail()
	}

	// Get the first element of the original list
	el := ol.Front()

	if ol.Front() == nil {
		t.Fail()
	}

	// Loop over both data structures and compare the result
	for n := l.Root; n.Next != nil; n = n.Next {

		if reflect.DeepEqual(n.Value, el.Value) == false {
			t.Fail()
		}

		// Move forward on both lists
		el = el.Next()
	}

}

func Test_InsertAtBeginning(t *testing.T) {

	l := NewListWithValue("world")

	l.InsertAtBeginning("hello")

	if v := l.Root.Value; v.(string) != "hello" {
		t.Fail()
	}

	if v := l.Root.Next.Value; v.(string) != "world" {
		t.Fail()
	}

}

func Test_InsertAtEnd(t *testing.T) {

	l := NewListFromSlice([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	l.InsertAtEnd("!")

	for n := l.Root; n.Next != nil; n = n.Next {

		// check we're at the end
		if n.Next == nil {
			if v := n.Value; v.(string) != "!" {
				t.Fail()
			}
		}
	}

}

func Test_InsertAtEndEmptyList(t *testing.T) {

	l := NewList()
	l.InsertAtEnd("test")

	if v := l.Root.Value; v.(string) != "test" {
		t.Fail()
	}

}

func Test_InsertAfter(t *testing.T) {

	l := NewListWithValue("world")
	l.InsertAtBeginning("hello")
	l.InsertAtBeginning("...")

	el := l.Root.Next // quickly go to the second element

	// Add an element after the second element
	l.InsertAfter(el, "!")

	elv := l.Root.Next.Next.Value

	if elv.(string) != "!" {
		t.Fail()
	}

}

func Test_InsertAfterEmptyList(t *testing.T) {

	l := NewList()

	err := l.InsertAfter(l.Root, "test")

	if err == nil {
		t.Fail()
	}

}

func Test_Contains(t *testing.T) {

	l := NewListWithValue("world")
	l.InsertAtBeginning("hello")

	if l.Contains("world") == nil {
		t.Fail()
	}

}

func Test_ContainsNegative(t *testing.T) {

	l := NewListWithValue("world")
	l.InsertAtBeginning("hello")

	if l.Contains("world123") != nil {
		t.Fail()
	}

}

func Test_ContainsEmptyList(t *testing.T) {

	l := NewList()

	if l.Contains("world123") != nil {
		t.Fail()
	}

}

func Test_Length(t *testing.T) {

	l := NewListWithValue("world")

	for i := 0; i < 10; i++ {
		l.InsertAtBeginning("hello")
	}

	if ll := l.Length(); ll != 11 {
		t.Fail()
	}

}
