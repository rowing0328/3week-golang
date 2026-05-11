package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainPrintsHelloAndIntroduction(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create stdout pipe: %v", err)
	}

	originalStdout := os.Stdout
	os.Stdout = writer

	main()

	os.Stdout = originalStdout
	if err := writer.Close(); err != nil {
		t.Fatalf("close stdout writer: %v", err)
	}

	var output bytes.Buffer
	if _, err := io.Copy(&output, reader); err != nil {
		t.Fatalf("read stdout: %v", err)
	}

	expected := "Hello, Go!\n" +
		"Go 프로그램은 package main의 main 함수에서 시작합니다.\n" +
		"안녕하세요! 저는 Go를 배우는 멘티입니다.\n"

	if output.String() != expected {
		t.Fatalf("unexpected output:\nwant:\n%q\ngot:\n%q", expected, output.String())
	}
}
