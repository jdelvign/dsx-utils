// Copyright 2019 Jerome Delvigne
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// empty struct that hold the process() function
type commandRemove struct{}

func (t *commandRemove) process() {

	var (
		dsxFileName    string
		outputFileName string
		objectName     string
	)

	rmCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	rmCmd.StringVar(&outputFileName, "outputFileName", "", "The output DSX file")
	rmCmd.StringVar(&dsxFileName, "dsxfile", "", "The input DSX file")
	rmCmd.StringVar(&objectName, "objectName", "", "Name of the object to be removed from <dsxfile>")

	rmCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl rm -dsxfile <DSXFILE> -outputFileName <OUTPUT> -objectName <OBJECT>\n")
		rmCmd.PrintDefaults()
	}

	rmCmd.Parse(os.Args[2:])

	if rmCmd.NFlag() != 3 {
		rmCmd.Usage()
		os.Exit(1)
	}

	fileIn := openFile(dsxFileName)
	defer fileIn.Close()

	scanner := bufio.NewScanner(fileIn)

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
			}
		}

		if line == endDSJOB {
			// Print job info now !

			fmt.Printf("%s\n", dsJobName)

			dsjob = false
			dsJobName = "<not available>"
		}

	}

}
