package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainPrintsHelloAndIntroduction(t *testing.T) {
	// main 함수가 표준 출력(stdout)에 쓴 내용을 테스트에서 읽기 위해 파이프를 만듭니다.
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("stdout 캡처용 파이프 생성 실패: %v", err)
	}

	// 기존 stdout을 보관한 뒤, 테스트 중에는 stdout이 파이프 writer를 바라보게 바꿉니다.
	originalStdout := os.Stdout
	os.Stdout = writer

	// 실제 프로그램 시작점인 main 함수를 호출해 출력 동작을 검증합니다.
	main()

	// stdout을 원래 상태로 되돌리고 writer를 닫아 reader가 출력 내용을 끝까지 읽을 수 있게 합니다.
	os.Stdout = originalStdout
	if err := writer.Close(); err != nil {
		t.Fatalf("stdout writer 닫기 실패: %v", err)
	}

	// 파이프 reader에 쌓인 출력 내용을 문자열 비교가 쉬운 버퍼로 복사합니다.
	var output bytes.Buffer
	if _, err := io.Copy(&output, reader); err != nil {
		t.Fatalf("stdout 읽기 실패: %v", err)
	}

	// fmt.Println은 각 줄 끝에 줄바꿈을 붙이므로 기대값에도 \n을 포함합니다.
	expected := "Hello, Go!\n" +
		"Go 프로그램은 package main의 main 함수에서 시작합니다.\n" +
		"안녕하세요! 저는 Go를 배우는 멘티입니다.\n"

	// 출력 순서와 내용이 하나라도 달라지면 테스트가 실패합니다.
	if output.String() != expected {
		t.Fatalf("예상과 다른 출력입니다:\n기대값:\n%q\n실제값:\n%q", expected, output.String())
	}
}
