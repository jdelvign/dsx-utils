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
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// some comment
type commandGrep struct{}

func (t *commandGrep) process() {

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
	displayJobName := false
	dsJobName := "<not available>"
	lineCounter := 1
	searchIndex := -1

	for scanner.Scan() {
		line := scanner.Text()
		if line == beginDSJOB {
			dsjob = true
		}

		if dsjob {
			if strings.HasPrefix(line, dsjobIDENTIFIER) {
				dsJobName = strings.Split(line, "\"")[1]
			}

			if ignoreCase {
				searchIndex = strings.Index(strings.ToLower(line), strings.ToLower(subString))
			} else {
				searchIndex = strings.Index(line, subString)
			}

			if searchIndex != -1 {
				if !displayJobName {
					fmt.Printf("%s:\n", dsJobName)
					displayJobName = true
				}
				fmt.Printf("%d:\t%s\n", lineCounter, line)
			}
		}

		if line == endDSJOB {
			dsjob = false
			dsJobName = "<not available>"
			displayJobName = false
		}

		lineCounter++
	}
}
