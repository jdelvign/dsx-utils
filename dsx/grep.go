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
		dsxFileName string
		subString   string
		ignoreCase  bool
	)

	grepCmd := flag.NewFlagSet("grep", flag.ExitOnError)
	grepCmd.StringVar(&subString, "substr", "", "The substring to find in the DSX file")
	grepCmd.StringVar(&dsxFileName, "dsxfile", "", "The DSX file to search in")
	grepCmd.BoolVar(&ignoreCase, "ignoreCase", false, "Search the substring in case sensitive (false/default) or not (true)")

	grepCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl grep -substr <SUBSTRING> [-ignoreCase] -dsxfile <DSXFILE>\n")
		grepCmd.PrintDefaults()
	}

	grepCmd.Parse(os.Args[2:])

	if (grepCmd.NFlag() > 3) || (grepCmd.NFlag() == 0) {
		grepCmd.Usage()
		os.Exit(1)
	}

	fmt.Printf("Searching \"%s\" in %s, ignoreCase=%t\n", subString, dsxFileName, ignoreCase)

	f := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

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

			if ignoreCase {
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
}
