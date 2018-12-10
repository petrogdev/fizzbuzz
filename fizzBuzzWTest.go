package main

import (
	"fmt"
	"testing"
)

func FizzBuzz(amount int) string {
	results := ""
	for i := 1; i <= amount; i++ {
		result := ""
		if i%3 == 0 { result += "Fizz" }
		if i%5 == 0 { result += "Buzz" }
		if result != "" {
			results += result + "\n"
			continue
		}
		results += fmt.Sprintf("%d\n", i)
	}
	return results
}

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	want :=
`1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
`
	if got != want {
		t.Errorf("FizzBuzz(15) \n got: \n%v \n want: \n%v", got, want)
	}
}

func main() {
	tests := []testing.InternalTest{{"TestFizzBuzz", TestFizzBuzz}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)

}
