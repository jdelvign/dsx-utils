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
)

type commandHeader struct{}

func (t *commandHeader) process() {

	var dsxFileName string

	headerCmd := flag.NewFlagSet("header", flag.ExitOnError)
	headerCmd.StringVar(&dsxFileName, "dsxfile", "", "The DSX file to inspect")

	headerCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dsxutl header -dsxfile DSXFILE\n")
		headerCmd.PrintDefaults()
	}

	headerCmd.Parse(os.Args[2:])

	if headerCmd.NFlag() != 1 {
		headerCmd.Usage()
		os.Exit(1)
	}

	f := openFile(dsxFileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	display := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == beginHeader {
			display = true
		}

		if display {
			fmt.Println(line) // Println will add back the final '\n'
		}

		if line == endHeader {
			display = false
		}
	}
}
