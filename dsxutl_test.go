package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/jdelvign/dsxutl/test"
)

var binaryName = "dsxutl"

func TestMain(m *testing.M) {
	fmt.Printf("Making \"%s\" ...\n", binaryName)

	make := exec.Command("make")
	err := make.Run()
	if err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestMainProgram(t *testing.T) {
	test.RunTest(t, binaryName, []test.Table{
		{Name: "no_command", Args: []string{}, Golden: "main-no-command.golden"},
		{Name: "unknown_command", Args: []string{"dummy"}, Golden: "main-unknown-command.golden"},
	})
}

func TestCmdLjobs(t *testing.T) {
	test.RunTest(t, binaryName, []test.Table{
		{Name: "no_arguments", Args: []string{"ljobs"}, Golden: "ljobs-no-args.golden"},
		{Name: "category", Args: []string{"ljobs", "test/test.dsx"}, Golden: "ljobs-with-category.golden"},
	})
}

func TestCmdHeader(t *testing.T) {
	test.RunTest(t, binaryName, []test.Table{
		{Name: "no_arguments", Args: []string{"header"}, Golden: "header-no-args.golden"},
		{Name: "header_output", Args: []string{"header", "test/test.dsx"}, Golden: "header-output.golden"},
	})
}

func TestCmdGrep(t *testing.T) {
	test.RunTest(t, binaryName, []test.Table{
		{Name: "no_arguments", Args: []string{"grep"}, Golden: "grep-no-args.golden"},
		{Name: "grep_case_sensitive", Args: []string{"grep", "-substr", "MCHO", "test/test.dsx"}, Golden: "grep-cs.golden"},
		{Name: "grep_case_insensitive", Args: []string{"grep", "-substr", "mcho", "-ci", "test/test.dsx"}, Golden: "grep-ci.golden"},
	})
}
