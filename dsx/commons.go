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
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// Some Constants
const (
	beginHeader string = "BEGIN HEADER"
	endHeader   string = "END HEADER"

	characterSet string = "   CharacterSet"

	beginDSJOB string = "BEGIN DSJOB"
	endDSJOB   string = "END DSJOB"

	beginDSROUTINES string = "BEGIN DSROUTINES"
	endDSROUTINES   string = "END DSROUTINES"

	dsroutineIDENTIFIER string = "      Identifier"
	dsroutineCATEGORY   string = "      Category"

	beginDSRECORD string = "   BEGIN DSRECORD"
	endDSRECORD   string = "   END DSRECORD"

	toolInstanceID string = "   ToolInstanceID"

	dsjobIDENTIFIER string = "   Identifier"
	dsjobCATEGORY   string = "      Category"
)

// Command ...
type Command interface {
	Process()
}

var csMap = map[string]encoding.Encoding{
	"CP1252": charmap.Windows1252,
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
// TODO: check the DSX charset here and return the appropriate scanner
func openFile(fileName string) (*os.File, *transform.Reader) {

	cs, err := extractCharset(fileName)
	check(err)

	f, err := os.Open(fileName)
	check(err)

	r := transform.NewReader(f, csMap[cs].NewDecoder())

	return f, r
}

func extractCharset(fileName string) (charset string, err error) {

	f, err := os.Open(fileName)
	defer f.Close()

	check(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, characterSet) {
			return strings.Split(line, "\"")[1], nil
		}
	}

	return "", fmt.Errorf("CharacterSet not found in %s", f.Name())
}
