package set

import (
	"testing"
)

func Test_NewSet(t *testing.T) {

	s := NewSet()

	if s.List == nil {
		t.Fail()
	}

}

func Test_Add(t *testing.T) {

	s := NewSet()

	err := s.Add("test")

	if err != nil {
		t.Fail()
	}


}

func Test_Add_Error(t *testing.T) {

	s := NewSet()

	err := s.Add("test")

	if err != nil {
		t.Fail()
	}

	err = s.Add("test")

	if err == nil {
		t.Fail()
	}

}

func Test_Contains(t *testing.T) {

	s := NewSet()
	s.Add("hello")

	if s.Contains("hello") == nil {
		t.Fail()
	}

}
