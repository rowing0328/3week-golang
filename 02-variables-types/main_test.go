package main

import (
	"bytes"
	"testing"
)

func TestRunPrintsVariablesAndTypes(t *testing.T) {
	var output bytes.Buffer

	run(&output)

	expected := "name: Alice\n" +
		"age: 25\n" +
		"city: Seoul\n" +
		"score: 98.5\n" +
		"isStudent: true\n" +
		"height: 170.5\n" +
		"email: alice@example.com\n" +
		"isAdmin: false\n" +
		"Alice is 26 years old next year.\n" +
		"type of score: float64\n" +
		"score with one decimal place: 98.5\n"

	if output.String() != expected {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected, output.String())
	}
}
