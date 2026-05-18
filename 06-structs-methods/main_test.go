package main

import (
	"bytes"
	"testing"
)

func TestIntroduceIncludesPersonFields(t *testing.T) {
	person := Person{
		Name:  "Alice",
		Age:   20,
		City:  "Seoul",
		Email: "alice@example.com",
	}

	got := person.Introduce()
	want := "저는 Alice이고 20살입니다. 사는 곳은 Seoul입니다. 이메일은 alice@example.com입니다."

	if got != want {
		t.Fatalf("예상과 다른 자기소개입니다: got %q, want %q", got, want)
	}
}

func TestIsAdultReturnsTrueForAdults(t *testing.T) {
	person := Person{Age: 18}

	if !person.IsAdult() {
		t.Fatalf("18세 이상은 성인으로 판단해야 합니다")
	}
}

func TestIsAdultReturnsFalseForMinors(t *testing.T) {
	person := Person{Age: 17}

	if person.IsAdult() {
		t.Fatalf("18세 미만은 성인이 아니어야 합니다")
	}
}

func TestHaveBirthdayIncreasesAge(t *testing.T) {
	person := Person{Age: 20}

	person.HaveBirthday()

	want := 21
	if person.Age != want {
		t.Fatalf("생일 후 나이가 증가해야 합니다: got %d, want %d", person.Age, want)
	}
}

func TestRunPrintsStructsAndMethods(t *testing.T) {
	var output bytes.Buffer

	run(&output)

	expected := "저는 Alice이고 20살입니다. 사는 곳은 Seoul입니다. 이메일은 alice@example.com입니다.\n" +
		"Alice is an adult.\n" +
		"after birthday: 21\n"

	if output.String() != expected {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected, output.String())
	}
}
