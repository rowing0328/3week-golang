package main

import (
	"fmt"
	"io"
	"os"
)

// run은 인사말과 자기소개 문장을 w에 출력합니다.
func run(w io.Writer) {
	fmt.Fprintln(w, "Hello, Go!")
	fmt.Fprintln(w, "Go 프로그램은 package main의 main 함수에서 시작합니다.")
	fmt.Fprintln(w, "안녕하세요! 저는 Go를 배우는 멘티입니다.")
}

func main() {
	run(os.Stdout)
}
