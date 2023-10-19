package main

// see generate.go in the directory above. this needs to be in a separate
// directory so that it's `go run`-able by `go generate`.

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	mainTemplate = `package main

// Code generated - DO NOT EDIT.

import "github.com/{{username}}/oracles-randomizer-ng/randomizer"

func main() {
	randomizer.Main()
}
`
	versionTemplate = `package randomizer

// Code generated - DO NOT EDIT.

const version = {{version}}
`
)

var (
	usernamePattern = strings.ReplaceAll(
		filepath.FromSlash(`github.com/(.+)/oracles-randomizer-ng`), `\`, `\\`)
	usernameRegexp = regexp.MustCompile(usernamePattern)
	versionRegexp  = regexp.MustCompile(`/(.+)-\d+-g(.+)`)
)

func main() {
	generateMain()
	generateVersion()
}

func generateMain() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	matches := usernameRegexp.FindStringSubmatch(wd)
	if matches == nil {
		panic("error getting import path from working directory")
	}

	s := strings.ReplaceAll(mainTemplate, "{{username}}", matches[1])
	if err := ioutil.WriteFile("main.go", []byte(s), 0644); err != nil {
		panic(err)
	}
}

func generateVersion() {
	// try matching an exact tag first
	describeCmd := exec.Command("git", "describe")
	output, err := describeCmd.Output()
	version := ""

	// not an exact tag; use long format
	if err != nil || strings.Contains(string(output), "-g") {
		branch, branchErr := exec.Command("git", "branch", "--show-current").Output()
		hash, hashErr := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		if branchErr != nil || hashErr != nil {
			panic("error getting version string from git")
		}
		branchStr := strings.TrimSpace(string(branch))
		hashStr := strings.TrimSpace(string(hash))
		version = fmt.Sprintf(`"%s-%s"`, branchStr, hashStr)
	} else {
		version = fmt.Sprintf(`"%s"`, strings.TrimSpace(string(output)))
	}

	s := strings.ReplaceAll(versionTemplate, "{{version}}", version)
	err = ioutil.WriteFile("randomizer/version.go", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
