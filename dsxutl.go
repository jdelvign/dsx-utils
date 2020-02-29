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
// Package main provides the entry point of the 'dsxutl' commands
package main

import (
	"fmt"
	"os"
)

// Some Constants
const (
	beginHeader string = "BEGIN HEADER"
	endHeader   string = "END HEADER"

	beginDSJOB string = "BEGIN DSJOB"
	endDSJOB   string = "END DSJOB"

	toolInstanceID string = "   ToolInstanceID"

	dsjobIDENTIFIER string = "   Identifier"
	dsjobCATEGORY   string = "      Category"
)

// Interface command
type command interface {
	process()
}

// Function main do something with `dsxutl` Command
// Subcommands :
// 		`dsxutl grep -substr <substring> -dsxfile <dsxfile>` : find the substring inside the Job Designs
//		`dsxutl header -dsxfile <dsxfile>` : Print the DSX header
func main() {

	//start := time.Now()

	m := make(map[string]command)
	m["grep"] = new(commandGrep)
	m["header"] = new(commandHeader)
	m["ljobs"] = new(commandLJobs)

	if len(os.Args) < 2 {
		fmt.Println("expected 'grep', 'header', 'ljobs' subcommands")
		os.Exit(1)
	}

	// Select the Subcommand
	cmd := m[os.Args[1]]

	if cmd != nil {
		cmd.process()
	} else {
		fmt.Println("expected 'grep', 'header', 'ljobs' subcommands")
		os.Exit(1)
	}

	//t := time.Now()
	//fmt.Printf("Elapsed Time: %v\n", t.Sub(start))
}

// Check error returned by I/O functions
func check(e error) {
	if e != nil {
		if os.IsNotExist(e) {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		} else {

		}
		panic(e)
	}
}

// open a file
func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	check(err)
	return f
}
