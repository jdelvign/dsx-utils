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
	"flag"
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
)

// Check error returned by I/O functions
func check(e error) {
	if e != nil {
			panic(e)
	}
}

const BEGIN_HEADER string = "BEGIN HEADER"
const END_HEADER string = "END HEADER" 

const BEGIN_DSJOB string = "BEGIN DSJOB"
const END_DSJOB string = "END DSJOB"

const DSJOB_IDENTIFIER string = "   Identifier"

// `dsxgrep` Command
//
func main() {

	headerPtr := flag.Bool("header", false, "Display DSX HEADER informations")
	ljobsPtr := flag.Bool("ljobs", false, "Display a list of all jobs in the specified DSX")
	//jobPtr := flag.String("job", "", "Job name for -lstages")
	//lstagesPtr := flag.Bool("lstages", false, "Display a list of all stages for the given job in the specified DSX")

	flag.Parse()

	start := time.Now()
	
	dsxFileName := flag.Args()[0]

	mainFlag := true

	// OPTION "-header" 
	if *headerPtr {

		mainFlag = false

		f, err := os.Open(dsxFileName)
		check(err)

		scanner := bufio.NewScanner(f)
	
		display := false

		for scanner.Scan() {
			line := scanner.Text() 
			if (line == BEGIN_HEADER) {
				display = true
			}

			if (display) {
				fmt.Println(line) // Println will add back the final '\n'
			}

			if (line == END_HEADER) {
				display = false
			}
		}		
	}

	// OPTION "-ljobs"
	if *ljobsPtr {

		mainFlag = false

		f, err := os.Open(dsxFileName)
		check(err)

		scanner := bufio.NewScanner(f)
	
		dsjob := false

		for scanner.Scan() {
			line := scanner.Text() 
			if (line == BEGIN_DSJOB) {
				dsjob = true
			}

			if (dsjob) {
				if strings.HasPrefix(line, DSJOB_IDENTIFIER) {
					fmt.Println(strings.Split(line, "\"") [1])
				}
			}

			if (line == END_DSJOB) {
				dsjob = false
			}
		}
	}

	// MAIN method 
	if mainFlag {

		subString := flag.Args()[1]
		fmt.Printf("Searching \"%s\" in %s\n", subString, dsxFileName)

		f, err := os.Open(dsxFileName)
		check(err)

		scanner := bufio.NewScanner(f)

		dsjob := false
		displayJobName := false
		dsJobName := "<not available>"
		lineCounter := 1

		for scanner.Scan() {
			line := scanner.Text() 
			if (line == BEGIN_DSJOB) {
				dsjob = true
			}

			if (dsjob) {
				if strings.HasPrefix(line, DSJOB_IDENTIFIER) {
					dsJobName = strings.Split(line, "\"") [1]
				}
				if (strings.Index(line, subString) != -1) {
					if !displayJobName {
						fmt.Printf("%s:\n", dsJobName)
						displayJobName = true
					}					
					fmt.Printf("%d:\t%s\n", lineCounter, line)
				}
			}

			if (line == END_DSJOB) {
				dsjob = false
				dsJobName = "<not available>"
				displayJobName = false
			}

			lineCounter++
		}
	}
	t := time.Now()
	fmt.Printf("Elapsed Time: %v\n", t.Sub(start))
}
