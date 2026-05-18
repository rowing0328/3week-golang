package main

import (
	"fmt"
	"io"
	"os"
)

// run은 Go의 기본 타입 변수들을 선언하고 w에 출력합니다.
func run(w io.Writer) {
	var name string = "Alice"
	var age int = 25
	city := "Seoul"
	score := 98.5
	isStudent := true

	var height float64 = 170.5
	email := "alice@example.com"
	isAdmin := false

	fmt.Fprintln(w, "name:", name)
	fmt.Fprintln(w, "age:", age)
	fmt.Fprintln(w, "city:", city)
	fmt.Fprintln(w, "score:", score)
	fmt.Fprintln(w, "isStudent:", isStudent)
	fmt.Fprintln(w, "height:", height)
	fmt.Fprintln(w, "email:", email)
	fmt.Fprintln(w, "isAdmin:", isAdmin)

	age = age + 1
	fmt.Fprintf(w, "%s is %d years old next year.\n", name, age)

	fmt.Fprintf(w, "type of score: %T\n", score)
	fmt.Fprintf(w, "score with one decimal place: %.1f\n", score)
}

func main() {
	run(os.Stdout)
}
