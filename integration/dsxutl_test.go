package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Printf("Making \"%s\" ...\n", binaryName)
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
	make := exec.Command("make")
	err = make.Run()
	if err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestMainProgram(t *testing.T) {
	runTest(t, []table{
		{"no_command", []string{}, "main-no-command.golden"},
		{"unknown_command", []string{"dummy"}, "main-unknown-command.golden"},
	})
}

func TestCmdLjobs(t *testing.T) {
	runTest(t, []table{
		{"no_arguments", []string{"ljobs"}, "ljobs-no-args.golden"},
		{"no_dsxfile_but_category", []string{"ljobs", "-withCategory"}, "ljobs-no-args.golden"},
		{"without_category", []string{"ljobs", "-dsxfile", "integration/test.dsx"}, "ljobs-no-category.golden"},
		{"with_category", []string{"ljobs", "-withCategory", "-dsxfile", "integration/test.dsx"}, "ljobs-with-category.golden"},
	})
}

func TestCmdHeader(t *testing.T) {
	runTest(t, []table{
		{"no_arguments", []string{"header"}, "header-no-args.golden"},
		{"header_output", []string{"header", "-dsxfile", "integration/test.dsx"}, "header-output.golden"},
	})
}

func TestCmdGrep(t *testing.T) {
	runTest(t, []table{
		{"no_arguments", []string{"grep"}, "grep-no-args.golden"},
		//{"header_output", []string{"header", "-dsxfile", "integration/test.dsx"}, "header-output.golden"},
	})
}
