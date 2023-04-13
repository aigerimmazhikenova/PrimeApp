package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_intro_prompt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	outBytes, err := io.ReadAll(r)
	if err != nil {
		t.Errorf("error reading output: %v", err)
	}

	outStr := string(outBytes)
	expected := "Enter a whole number"

	if !strings.Contains(outStr, expected) {
		t.Errorf("expected output %q not found in actual output %q", expected, outStr)
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	stdin := strings.NewReader("1\nq\n")

	go readUserInput(stdin, doneChan)

	<-doneChan

	close(doneChan)
}
func Test_checkNumbers(t *testing.T) {
	checktest := []struct {
		num   string
		check string
		res   string
	}{
		{"empty", "", "Please enter a whole number!"},
		{"zero", "0", "0 is not prime, by definition!"},
		{"one", "1", "1 is not prime, by definition!"},
		{"two", "2", "2 is a prime number!"},
		{"negative", "-1", "Negative numbers are not prime, by definition!"},
		{"typed", "three", "Please enter a whole number!"},
		{"float", "1.1", "Please enter a whole number!"},
		{"exit", "q", ""},
		{"QUIT", "Q", ""},
	}
	for _, e := range checktest {
		input := strings.NewReader(e.check)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.res) {
			t.Errorf("%s: expected %s, but got %s", e.num, e.res, res)
		}
	}

}
