package main

import (
	"testing"
)

func TestFunction(t *testing.T) {

	inputs := [...]string{
		"(test string)",
		"()",
		"(lala",
		"lala)",
		"(test string 2)",
		"(asdf)(test)",
		"((a(a(a)",
		"((aaa))",
	}
	expectedOutputs := [...]string{
		"test string",
		"",
		"",
		"",
		"test string 2",
		"asdf",
		"(a(a(a",
		"(aaa",
	}

	for i, element := range inputs {
		output := findFirstStringInBracket(element)
		if output != expectedOutputs[i] {
			t.Fatalf("test %d result expected %s but found %s", i+1, expectedOutputs[i], output)
		}
		t.Logf("test %d success", i+1)
	}
}
