package tic

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
