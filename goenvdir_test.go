package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	AVar = "./env/A"
	BVar = "./env/B"
)

func TestStart(t *testing.T)  {

}

func TestReadVars(t *testing.T)  {
	_ = os.Mkdir("env", 700)
	_, _ = os.Create(AVar)
	_, _ = os.Create(BVar)

	_ = ioutil.WriteFile(AVar, []byte("123"), 700)
	_ = ioutil.WriteFile(BVar, []byte("b"), 700)

	vars, _ := readVars("./env")

	const expectedLen = 2
	const expectedA = "A=123"
	const expectedB = "B=b"

	if expectedLen != len(vars) {
		t.Error("Expected: len of vars slice:", expectedLen, "got:", len(vars))
	} else {
		if vars[0] != expectedA {
			t.Error("Expected: value of A", expectedA, "got:", vars[0])
		}

		if vars[1] != expectedB {
			t.Error("Expected: value of B", expectedB, "got:", vars[1])
		}
	}

	_ = os.RemoveAll("env")
}