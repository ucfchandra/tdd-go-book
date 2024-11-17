package calculator_test

import (
	"os"
	"log"
	"testing"
	"github.com/ucfchandra/tdd-go-book/chapter02/calculator"
)

func TestMain(m *testing.M) {
	// setup statements
	setup()
	// run the tests
	e := m.Run()
	// cleanup statements
	teardown()
	// report the exit code
	os.Exit(e)
}

func setup() {
	log.Println("Setting up")
}

func teardown() {
	log.Println("Tearing down")
}

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0

	// Act
	got := e.Add(2.5,3.5)

	// Assert
	if got != want {
		t.Errorf("Add(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}


