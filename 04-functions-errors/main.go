package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func greet(name string) string {
	return "안녕하세요, " + name + "님!"
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil
}

func add(a, b int) int {
	return a + b
}

// run은 함수 호출, 다중 반환값, 에러 처리 예제의 결과를 w에 출력합니다.
func run(w io.Writer) {
	fmt.Fprintln(w, greet("Alice"))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Fprintln(w, "error:", err)
	} else {
		fmt.Fprintf(w, "result: %.2f\n", result)
	}

	fmt.Fprintln(w, "sum:", add(2, 3))

	result, err = divide(5, 0)
	if err != nil {
		fmt.Fprintln(w, "error:", err)
		return
	}

	fmt.Fprintln(w, "result:", result)
}

func main() {
	run(os.Stdout)
}
