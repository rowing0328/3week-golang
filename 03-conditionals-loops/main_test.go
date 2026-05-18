package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunPrintsConditionalsLoopsAndEvenNumbers(t *testing.T) {
	var output bytes.Buffer

	run(&output)

	var expected strings.Builder
	expected.WriteString("grade: B\n")
	expected.WriteString("weekend\n")
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(&expected, "%d번째 반복\n", i)
	}
	for n := 0; n < 3; n++ {
		fmt.Fprintf(&expected, "n = %d\n", n)
	}
	for index, value := range []string{"apple", "banana", "cherry"} {
		fmt.Fprintf(&expected, "[%d] %s\n", index, value)
	}
	for number := 1; number <= 100; number++ {
		if number%2 == 0 {
			fmt.Fprintf(&expected, "%d\n", number)
		}
	}

	if output.String() != expected.String() {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected.String(), output.String())
	}
}
