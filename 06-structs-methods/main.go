package main

import (
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name  string
	Age   int
	City  string
	Email string
}

func (p Person) Introduce() string {
	return fmt.Sprintf("저는 %s이고 %d살입니다. 사는 곳은 %s입니다. 이메일은 %s입니다.", p.Name, p.Age, p.City, p.Email)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) HaveBirthday() {
	p.Age++
}

// run은 구조체와 메서드 예제의 결과를 w에 출력합니다.
func run(w io.Writer) {
	person := Person{
		Name:  "Alice",
		Age:   20,
		City:  "Seoul",
		Email: "alice@example.com",
	}

	fmt.Fprintln(w, person.Introduce())

	if person.IsAdult() {
		fmt.Fprintln(w, person.Name, "is an adult.")
	}

	person.HaveBirthday()
	fmt.Fprintln(w, "after birthday:", person.Age)
}

func main() {
	run(os.Stdout)
}
