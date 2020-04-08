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

// Package dsx contains code for the dsxutl command
package dsx

import (
	"fmt"
	"os"
)

// Some Constants
const (
	beginHeader string = "BEGIN HEADER"
	endHeader   string = "END HEADER"

	dsCharset string = "   CharacterSet"

	beginDSJOB string = "BEGIN DSJOB"
	endDSJOB   string = "END DSJOB"

	toolInstanceID string = "   ToolInstanceID"

	dsjobIDENTIFIER string = "   Identifier"
	dsjobCATEGORY   string = "      Category"
)

// Command ...
type Command interface {
	Process()
}

// Check error returned by I/O functions
func check(e error) {
	if e != nil {
		if os.IsNotExist(e) {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}
		panic(e)
	}
}

// open a file
// TODO: check the encoding here and return it
func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	check(err)
	return f
}

func extractCharset() {
	
}
