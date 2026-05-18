package main

import (
	"fmt"
	"io"
	"os"
)

// run은 조건문과 반복문 예제의 결과를 w에 출력합니다.
func run(w io.Writer) {
	score := 85

	if score >= 90 {
		fmt.Fprintln(w, "grade: A")
	} else if score >= 80 {
		fmt.Fprintln(w, "grade: B")
	} else if score >= 70 {
		fmt.Fprintln(w, "grade: C")
	} else {
		fmt.Fprintln(w, "grade: F")
	}

	day := "sat"
	switch day {
	case "sat", "sun":
		fmt.Fprintln(w, "weekend")
	default:
		fmt.Fprintln(w, "weekday")
	}

	for i := 1; i <= 5; i++ {
		fmt.Fprintf(w, "%d번째 반복\n", i)
	}

	n := 0
	for n < 3 {
		fmt.Fprintln(w, "n =", n)
		n++
	}

	fruits := []string{"apple", "banana", "cherry"}
	for index, value := range fruits {
		fmt.Fprintf(w, "[%d] %s\n", index, value)
	}

	for number := 1; number <= 100; number++ {
		if number%2 == 0 {
			fmt.Fprintln(w, number)
		}
	}
}

func main() {
	run(os.Stdout)
}
