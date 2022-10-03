package utils

import (
	"testing"
)

func TestIn(t *testing.T) {
	pass1 := In("john", []string{"matt", "mark", "john"})
	if !pass1 {
		t.Fail()
	}

	fail1 := In("John", []string{"matt", "mark", "john"})
	if fail1 {
		t.Fail()
	}
}
