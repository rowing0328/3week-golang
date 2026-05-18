package main

import (
	"fmt"
	"io"
	"os"
)

func averageScore(scores map[string]int) float64 {
	if len(scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range scores {
		total += score
	}

	return float64(total) / float64(len(scores))
}

// run은 배열, 슬라이스, 맵 예제의 결과를 w에 출력합니다.
func run(w io.Writer) {
	numbers := [3]int{10, 20, 30}
	fmt.Fprintln(w, "array:", numbers)
	fmt.Fprintln(w, "first number:", numbers[0])

	names := []string{"Alice", "Bob", "Carol"}
	names = append(names, "Dave")

	fmt.Fprintln(w, "slice:", names)
	fmt.Fprintln(w, "slice length:", len(names))

	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
	}
	scores["Carol"] = 88
	scores["Mentee"] = 95

	aliceScore, ok := scores["Alice"]
	if ok {
		fmt.Fprintln(w, "Alice score:", aliceScore)
	}

	for _, name := range []string{"Alice", "Bob", "Carol", "Mentee"} {
		fmt.Fprintf(w, "%s: %d점\n", name, scores[name])
	}

	fmt.Fprintf(w, "average score: %.2f\n", averageScore(scores))
}

func main() {
	run(os.Stdout)
}
