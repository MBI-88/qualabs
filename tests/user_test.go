package tests

import (
	"qualabs/src"
	"testing"
)


func TestSolutionA(t *testing.T) {
	user := src.NewUser()

	if !user.LoadData("./mock") {
		t.Fatalf("Error loading data")
	}

	result, err := user.SolutionA()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatalf("Solution A error")
	}

}

func TestSolutionB(t *testing.T) {
	user := src.NewUser()

	if !user.LoadData("./mock") {
		t.Fatalf("Error loading data")
	}

	if len(user.SolutionB()) == 0 {
		t.Fatalf("Solution B error")
	}

}
