package main

import (
	"bytes"
	"testing"
)

func TestAverageScoreReturnsMean(t *testing.T) {
	scores := map[string]int{
		"Alice":  90,
		"Bob":    75,
		"Carol":  88,
		"Mentee": 95,
	}

	got := averageScore(scores)
	want := 87.0

	if got != want {
		t.Fatalf("예상과 다른 평균 점수입니다: got %.2f, want %.2f", got, want)
	}
}

func TestAverageScoreReturnsZeroForEmptyMap(t *testing.T) {
	got := averageScore(map[string]int{})
	want := 0.0

	if got != want {
		t.Fatalf("빈 맵의 평균은 0이어야 합니다: got %.2f, want %.2f", got, want)
	}
}

func TestRunPrintsCollections(t *testing.T) {
	var output bytes.Buffer

	run(&output)

	expected := "array: [10 20 30]\n" +
		"first number: 10\n" +
		"slice: [Alice Bob Carol Dave]\n" +
		"slice length: 4\n" +
		"Alice score: 90\n" +
		"Alice: 90점\n" +
		"Bob: 75점\n" +
		"Carol: 88점\n" +
		"Mentee: 95점\n" +
		"average score: 87.00\n"

	if output.String() != expected {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected, output.String())
	}
}
