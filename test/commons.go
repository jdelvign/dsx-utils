/*
Copyright 2019 Jerome Delvigne

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package test contains code for the dsxutl integration test
package test

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

// Table represents the TableDrivenTest struct for dsxutl
type Table struct {
	Name   string
	Args   []string
	Golden string
}

func goldenPath(t *testing.T, golden string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("problems recovering caller information")
	}

	return filepath.Join(filepath.Dir(filename), golden)
}

func loadGolden(t *testing.T, golden string) string {
	content, err := ioutil.ReadFile(goldenPath(t, golden))
	if err != nil {
		t.Fatal(err)
	}

	return string(content)
}

// RunTest iterate over each Table elements and run the sub test
func RunTest(t *testing.T, binaryName string, tests []Table) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.Args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				if err.Error() != "exit status 1" {
					t.Fatal(err)
				}
			}

			actual := string(output)

			expected := loadGolden(t, tt.Golden)

			// Using go-cmp instead of reflect.DeepEqual
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Errorf("%s:\tmismatch (-expected +actual):\n%s", tt.Name, diff)
			}

		})
	}
}
