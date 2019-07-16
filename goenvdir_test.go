package main

import (
	"testing"
)

const (
	evnDir  = "./internal/prog/env"
	progDir = "./internal/prog/bin/prog"
)

func TestReadVars(t *testing.T) {
	const expectedLen = 2
	const expectedA = "A_ENV=123"
	const expectedB = "B_VAR=another_val"

	vars, _ := readVars(evnDir)
	gotLen := len(vars)

	if expectedLen != gotLen {
		t.Error("Expected: len of vars slice:", expectedLen, "got:", gotLen)
	} else {
		if vars[0] != expectedA {
			t.Error("Expected: value of A", expectedA, "got:", vars[0])
		}

		if vars[1] != expectedB {
			t.Error("Expected: value of B", expectedB, "got:", vars[1])
		}
	}
}

func TestStart(t *testing.T) {
	vars, _ := readVars(evnDir)
	out := start(progDir, vars)

	expected := "Received:  A_ENV = 123\nReceived:  B_VAR = another_val\n"
	got := string(out)

	if expected != got {
		t.Error("Expected out:", expected, "got:", got)
	}
}
