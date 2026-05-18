package main

import (
	"bytes"
	"testing"
)

func TestGreetReturnsKoreanGreeting(t *testing.T) {
	got := greet("Alice")
	want := "안녕하세요, Alice님!"

	if got != want {
		t.Fatalf("예상과 다른 인사말입니다: got %q, want %q", got, want)
	}
}

func TestDivideReturnsQuotient(t *testing.T) {
	got, err := divide(10, 2)
	if err != nil {
		t.Fatalf("예상하지 못한 에러입니다: %v", err)
	}

	want := 5.0
	if got != want {
		t.Fatalf("예상과 다른 나눗셈 결과입니다: got %.2f, want %.2f", got, want)
	}
}

func TestDivideReturnsErrorWhenDividingByZero(t *testing.T) {
	got, err := divide(5, 0)
	if err == nil {
		t.Fatalf("0으로 나눌 때 에러가 발생해야 합니다")
	}
	if got != 0 {
		t.Fatalf("에러가 발생하면 결과값은 0이어야 합니다: got %.2f", got)
	}

	want := "0으로 나눌 수 없습니다"
	if err.Error() != want {
		t.Fatalf("예상과 다른 에러 메시지입니다: got %q, want %q", err.Error(), want)
	}
}

func TestAddReturnsSum(t *testing.T) {
	got := add(2, 3)
	want := 5

	if got != want {
		t.Fatalf("예상과 다른 덧셈 결과입니다: got %d, want %d", got, want)
	}
}

func TestRunPrintsFunctionsAndErrors(t *testing.T) {
	var output bytes.Buffer

	run(&output)

	expected := "안녕하세요, Alice님!\n" +
		"result: 3.33\n" +
		"sum: 5\n" +
		"error: 0으로 나눌 수 없습니다\n"

	if output.String() != expected {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected, output.String())
	}
}
