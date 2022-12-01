package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func GetInput(year int, day int) string {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		log.Fatal("AOC_SESSION environment variable not set")
	}

	inputFile := fmt.Sprintf("%d.txt", day)

	var body []byte
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {

		client := &http.Client{}

		req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(inputFile, body, 0600)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		body, err = os.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	return string(body)
}

type TestCase struct {
	Input    string
	Expected string
}

func GetTest(year int, day int, part int, testcases []TestCase, solution func(string) string) func(t *testing.T) {
	timeSinceStart := time.Since(time.Date(year, 12, day, 0, 0, 0, 0, time.UTC)).Hours()/24 + 1

	if timeSinceStart < 1 {
		return func(t *testing.T) {
			t.Skipf("Day %d not yet available", day)
		}
	}

	return func(t *testing.T) {
		failed := false
		for _, testcase := range testcases {
			t.Run(testcase.Input, func(t *testing.T) {
				actual := solution(testcase.Input)
				if actual != testcase.Expected {
					failed = true
					t.Errorf("Expected %s, got %s", testcase.Expected, actual)
				}
			})
		}

		if !failed {
			input := GetInput(year, day)
			s := solution(input)
			fmt.Printf("Solution: %s\n", s)
		}
	}
}
