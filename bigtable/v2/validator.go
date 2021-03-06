// Copyright 2019, Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	bigtable_v2_tests "github.com/googleapis/conformance-tests/bigtable/v2/generated/google/cloud/conformance/bigtable/v2"
)

func main() {
	args := os.Args
	dir := args[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "not-a-test.json") ||
			!strings.Contains(f.Name(), ".json") {
			continue
		}

		log.Printf("Validating: %v/%v", dir, f.Name())

		inBytes, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		var testfile bigtable_v2_tests.TestFile
		if err := jsonpb.Unmarshal(bytes.NewBuffer(inBytes), &testfile); err != nil {
			log.Fatal(err)
		}
	}
}
