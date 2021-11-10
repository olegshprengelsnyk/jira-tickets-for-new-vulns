package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndToEndFunc(t *testing.T) {

	server := HTTPResponseEndToEnd()
	defer server.Close()

	os.Args = append(os.Args, "-orgID=123")
	os.Args = append(os.Args, "-token=123")
	os.Args = append(os.Args, "-jiraProjectID=123")
	os.Args = append(os.Args, "-api="+server.URL)

	// Get the console output
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// Test finished, read the output and compare with expectation
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	// Delete the file created for the test
	removeLogFile()

	compare := strings.Contains(string(out), "Number of tickets created: 3")

	assert.Equal(t, compare, true)

}

// comment for now, error with the arguments that are redefined
// Probably need a clear somewhere in the previous TestEndToEndFunc

// func TestEndToEndDryRunFunc(t *testing.T) {

// 	server := HTTPResponseEndToEnd()
// 	defer server.Close()

// 	fmt.Println(os.Args)

// 	os.Args = append(os.Args, "-orgID=123")
// 	os.Args = append(os.Args, "-token=123")
// 	os.Args = append(os.Args, "-jiraProjectID=123")
// 	os.Args = append(os.Args, "-api="+server.URL)
// 	os.Args = append(os.Args, "-dryRun=true")

// 	// Get the console output
// 	rescueStdout := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	main()

// 	os.Args = []string{}

// 	// Test finished, read the output and compare with expectation
// 	w.Close()
// 	out, _ := ioutil.ReadAll(r)
// 	os.Stdout = rescueStdout

// 	compare := strings.Contains(string(out), "Number of tickets created: 0")
// 	dryRunResult := strings.Contains(string(out), "Dry run result can be found in .log file")

// 	// Delete the file created for the test
// 	removeLogFile()

// 	assert.Equal(t, compare, true)
// 	assert.Equal(t, dryRunResult, true)

// }
