/*
Copyright 2023 The pdfcpu Authors.

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

package test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/flamacue/pdfcpu/pkg/cli"
)

func TestPageMode(t *testing.T) {
	msg := "testPageMode"

	pageMode := "UseOutlines"

	inFile := filepath.Join(inDir, "test.pdf")
	outFile := filepath.Join(outDir, "test.pdf")

	cmd := cli.ListPageModeCommand(inFile, conf)
	ss, err := cli.Process(cmd)
	if err != nil {
		t.Fatalf("%s %s: list pageMode: %v\n", msg, inFile, err)
	}
	if len(ss) > 0 && !strings.HasPrefix(ss[0], "No page mode") {
		t.Fatalf("%s %s: list pageMode, unexpected: %s\n", msg, inFile, ss[0])
	}

	cmd = cli.SetPageModeCommand(inFile, outFile, pageMode, nil)
	if _, err = cli.Process(cmd); err != nil {
		t.Fatalf("%s %s: set pageMode: %v\n", msg, outFile, err)
	}

	cmd = cli.ListPageModeCommand(outFile, conf)
	ss, err = cli.Process(cmd)
	if err != nil {
		t.Fatalf("%s %s: list pageMode: %v\n", msg, outFile, err)
	}
	if len(ss) == 0 {
		t.Fatalf("%s %s: list pageMode, missing pageMode\n", msg, outFile)
	}
	if ss[0] != pageMode {
		t.Fatalf("%s %s: list pageMode, want:%s, got:%s\n", msg, outFile, pageMode, ss[0])
	}

	cmd = cli.ResetPageModeCommand(outFile, "", nil)
	if _, err = cli.Process(cmd); err != nil {
		t.Fatalf("%s %s: reset pageMode: %v\n", msg, outFile, err)
	}

	cmd = cli.ListPageModeCommand(outFile, conf)
	ss, err = cli.Process(cmd)
	if err != nil {
		t.Fatalf("%s %s: list pageMode: %v\n", msg, outFile, err)
	}
	if len(ss) > 0 && !strings.HasPrefix(ss[0], "No page mode") {
		t.Fatalf("%s %s: list pageMode, unexpected: %s\n", msg, outFile, ss[0])
	}
}
