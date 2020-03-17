package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type table struct {
	name    string
	args    []string
	fixture string
}

var binaryName = "dsxutl"

func fixturePath(t *testing.T, fixture string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("problems recovering caller information")
	}

	return filepath.Join(filepath.Dir(filename), fixture)
}

func loadFixture(t *testing.T, fixture string) string {
	content, err := ioutil.ReadFile(fixturePath(t, fixture))
	if err != nil {
		t.Fatal(err)
	}

	return string(content)
}

func runTest(t *testing.T, tests []table) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				if err.Error() != "exit status 1" {
					t.Fatal(err)
				}
			}

			actual := string(output)

			expected := loadFixture(t, tt.fixture)

			// Using go-cmp instead of reflect.DeepEqual
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Errorf("%s:\tmismatch (-expected +actual):\n%s", tt.name, diff)
			}

		})
	}
}
