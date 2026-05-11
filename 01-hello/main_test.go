package main

import (
	"bytes"
	"testing"
)

func TestRunPrintsHelloAndIntroduction(t *testing.T) {
	// bytes.Buffer는 메모리에 출력 내용을 저장하는 writer입니다.
	var output bytes.Buffer

	// run 함수에 버퍼를 넘기면 실제 stdout을 바꾸지 않고 출력 결과를 검증할 수 있습니다.
	run(&output)

	// fmt.Println은 각 줄 끝에 줄바꿈을 붙이므로 기대값에도 \n을 포함합니다.
	expected := "Hello, Go!\n" +
		"Go 프로그램은 package main의 main 함수에서 시작합니다.\n" +
		"안녕하세요! 저는 Go를 배우는 멘티입니다.\n"

	// 출력 순서와 내용이 하나라도 달라지면 테스트가 실패합니다.
	if output.String() != expected {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected, output.String())
	}
}
