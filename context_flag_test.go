package ixgo

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

type invokingTestProcessor struct{}

func (invokingTestProcessor) LoadMain(_ *build.Context, pkg *build.Package) ([]byte, error) {
	return []byte(fmt.Sprintf(`package main
import tested %q
func main() {
	tested.TestSample()
	if !tested.TestSampleRan() { panic("TestSample was not executed") }
}
`, pkg.ImportPath)), nil
}

func TestRunTestPreservesTestingFlags(t *testing.T) {
	oldWorkingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := os.MkdirTemp(".", "runtest-flag-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	dir, err = filepath.Abs(dir)
	if err != nil {
		t.Fatal(err)
	}
	source := []byte(`package sample

var testSampleRan bool

func TestSample() {
	testSampleRan = true
}

func TestSampleRan() bool {
	return testSampleRan
}
`)
	if err := os.WriteFile(filepath.Join(dir, "sample_test.go"), source, 0o644); err != nil {
		t.Fatal(err)
	}

	oldTestProcessor := testProcessor
	testProcessor = invokingTestProcessor{}
	t.Cleanup(func() { testProcessor = oldTestProcessor })

	if err := NewContext(0).RunTest(dir, nil); err != nil {
		t.Fatal(err)
	}
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if workingDir != oldWorkingDir {
		t.Fatalf("working directory changed from %q to %q", oldWorkingDir, workingDir)
	}
	_ = testing.Verbose()
}
