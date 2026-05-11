package main

import (
	"fmt"
	"io"
	"os"
)

func run(w io.Writer) {
	fmt.Fprintln(w, "Hello, Go!")
	fmt.Fprintln(w, "Go 프로그램은 package main의 main 함수에서 시작합니다.")
	fmt.Fprintln(w, "안녕하세요! 저는 Go를 배우는 멘티입니다.")
}

func main() {
	run(os.Stdout)
}
