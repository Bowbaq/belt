package belt_test

import (
	"errors"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/Bowbaq/belt"
)

func TestCheckNil(t *testing.T) {
	if os.Getenv("TEST_CHECK_NIL") == "1" {
		belt.Check(nil)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCheckNil")
	cmd.Env = append(os.Environ(), "TEST_CHECK_NIL=1")

	if err := cmd.Run(); err != nil {
		t.Fatalf("process ran with err %v, want exit status 0", err)
	}
}

func TestCheckErr(t *testing.T) {
	if os.Getenv("TEST_CHECK_ERR") == "1" {
		belt.Check(errors.New("This should crash"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCheckErr")
	cmd.Env = append(os.Environ(), "TEST_CHECK_ERR=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func ExampleCheck() {
	_, err := http.Get("http://www.example.com/")
	belt.Check(err)
}
