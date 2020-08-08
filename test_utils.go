package main

import "testing"

func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("\nactual:\n%s\nexpected:\n%s", actual, expected)
	}
}
