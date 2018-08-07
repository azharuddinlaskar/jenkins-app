package main

import (
	"testing"
	)

func TestApp(t *testing.T) {
	m := matrix()
	if m == nil {
		t.Errorf("Not a valid matrix")
	}
}
