/*
Copyright 2022 Jerome Delvigne

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
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// CommandDrop some comment
type CommandDrop struct{}

// Process ...
func (t *CommandDrop) Process() {

	var (
		dsxFileName string
		joblist     string
	)

	dropCmd := flag.NewFlagSet("drop", flag.ExitOnError)
	dropCmd.StringVar(&joblist, "joblist", "", "A comma separated job list to drop from the DSX file")

	dropCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl drop -joblist <JOB1,JOB2,JOB3> <DSXFILE>\n")
		dropCmd.PrintDefaults()
	}

	dropCmd.Parse(os.Args[2:])

	if joblist == "" ||
		len(dropCmd.Args()) != 1 {
		dropCmd.Usage()
		os.Exit(1)
	}

	dsxFileName = dropCmd.Args()[0]

	fmt.Printf("Drop \"%s\" from %s\n", joblist, dsxFileName)

	f, r := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(r)
	buffer := make([]byte, bufferSize)
	scanner.Buffer(buffer, bufferSize)

	dsjob := false
	dsexecjob := false
	dsJobName := "<not available>"
	dsCategory := "<not available>"
	lineCounter := 1
	searchIndex := -1

	var matches []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == beginDSJOB {
			dsjob = true
		}
		if line == beginDSEXECJOB {
			dsexecjob = true
		}

		if dsjob || dsexecjob {
			if strings.HasPrefix(line, dsjobIDENTIFIER) {
				dsJobName = strings.Split(line, "\"")[1]
			}

			if searchIndex != -1 {
				matches = append(matches, fmt.Sprintf("%d:\t%s", lineCounter, line))
			}
		}

		if line == endDSJOB {

			if len(matches) != 0 {
				fmt.Printf("%s;%s:\n", dsJobName, dsCategory)
				for _, match := range matches {
					fmt.Printf("%s\n", match)
				}
			}

			dsjob = false
			dsJobName = "<not available>"
			dsCategory = "<not available>"
			matches = nil
		}

		lineCounter++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading dsx file: %e", err)
	}
}
