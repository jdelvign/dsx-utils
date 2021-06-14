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
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// CommandGrep some comment
type CommandGrep struct{}

// Process ...
func (t *CommandGrep) Process() {

	var (
		dsxFileName     string
		subString       string
		caseInsensitive bool
	)

	grepCmd := flag.NewFlagSet("grep", flag.ExitOnError)
	grepCmd.StringVar(&subString, "substr", "", "The substring to find in the DSX file")
	grepCmd.BoolVar(&caseInsensitive, "ci", false, "Search the substring in case insensitive")

	grepCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl grep -substr <SUBSTRING> [-ci] <DSXFILE>\n")
		grepCmd.PrintDefaults()
	}

	grepCmd.Parse(os.Args[2:])

	if subString == "" ||
		len(grepCmd.Args()) != 1 {
		grepCmd.Usage()
		os.Exit(1)
	}

	dsxFileName = grepCmd.Args()[0]

	fmt.Printf("Searching \"%s\" in %s, CaseInsensitive=%t\n", subString, dsxFileName, caseInsensitive)

	f, r := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(r)
	buffer := make([]byte, bufferSize)
	scanner.Buffer(buffer, bufferSize)

	dsjob := false
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

		if dsjob {
			if strings.HasPrefix(line, dsjobIDENTIFIER) {
				dsJobName = strings.Split(line, "\"")[1]
			}

			if strings.HasPrefix(line, dsjobCATEGORY) {
				dsCategory = strings.Split(line, "\"")[1]
			}

			if caseInsensitive {
				searchIndex = strings.Index(strings.ToLower(line), strings.ToLower(subString))
			} else {
				searchIndex = strings.Index(line, subString)
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
