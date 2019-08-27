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
type commandLJobs struct{}

func (t *commandLJobs) process() {

	var (
		withCategory bool
		dsxFileName  string
	)

	ljobsCmd := flag.NewFlagSet("ljobs", flag.ExitOnError)
	ljobsCmd.BoolVar(&withCategory, "withCategory", false, "Display the Category where job resides")
	ljobsCmd.StringVar(&dsxFileName, "dsxfile", "", "The DSX file to search in")

	ljobsCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl ljobs [-withCategory] -dsxfile DSXFILE\n")
		ljobsCmd.PrintDefaults()
	}

	ljobsCmd.Parse(os.Args[2:])

	if (ljobsCmd.NFlag() > 2) || (ljobsCmd.NFlag() == 0) {
		ljobsCmd.Usage()
		os.Exit(1)
	}

	f := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dsjob := false
	dsJobName := "<not available>"

	for scanner.Scan() {
		line := scanner.Text()
		if line == beginDSJOB {
			dsjob = true
		}

		if dsjob {
			if strings.HasPrefix(line, dsjobIDENTIFIER) {
				dsJobName = strings.Split(line, "\"")[1]
				fmt.Printf("%s\n", dsJobName)
			}
		}

		if line == endDSJOB {
			dsjob = false
			dsJobName = "<not available>"
		}

	}

}
