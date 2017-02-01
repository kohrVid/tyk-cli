package utils

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCheckDomain(t *testing.T) {
	expectedResult := "http://www.example.com"
	result := CheckDomain("http://www.example.com")
	if result != expectedResult {
		t.Fatalf("Expected %s, got %s", expectedResult, result)
	}
}

func TestCheckDomainErrorsWithBadInput(t *testing.T) {
	if os.Getenv("ERRS") == "1" {
		CheckDomain("www.example.com")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckDomainErrorsWithBadInput")
	cmd.Env = append(os.Environ(), "ERRS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("test ran with err %v, want exit status 1", err)
}

func TestIsProtocolPresentOutputsTrueIfPresent(t *testing.T) {
	expectedResult := true
	result := isProtocolPresent("http://www.example.com")
	if result != expectedResult {
		t.Fatalf("Expected %s, got %b", expectedResult, result)
	}
}

func TestIsProtocolPresentOutputsFalseIfMissing(t *testing.T) {
	expectedResult := false
	result := isProtocolPresent("/www.example.com")
	if result != expectedResult {
		t.Fatalf("Expected %s, got %b", expectedResult, result)
	}
}
