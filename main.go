package main

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func checkTrellisRequirements() {
	ansiblePlaybook, err := exec.LookPath("ansible-playbook")
	checkMsg(err, "You need ansible-playbook")
	ansibleVersionCommand := exec.Command(ansiblePlaybook, "--version")
	ansibleVersionOutput, err := ansibleVersionCommand.Output()
	versionText := string(ansibleVersionOutput)
	versionArray := strings.Split(versionText, "\n")
	userVersion, err := version.NewVersion(strings.Split(versionArray[0], " ")[1])
	check(err)

	acceptableVersions := ">= 1.9.2, < 2.0.0"
	constraint, err := version.NewConstraint(acceptableVersions)
	check(err)

	if !constraint.Check(userVersion) {
		fmt.Println("Invalid version of ansible. We require " + acceptableVersions)
		os.Exit(1)
	}
}

func cloneRepo(repo string, location string) {
	git, err := exec.LookPath("git")
	check(err)

	gitCommand := exec.Command(
		git,
		"clone",
		"https://github.com/roots/"+repo,
		location,
	)
	gitCommand.Stderr = os.Stdout
	gitCommand.Run()
}

func scaffoldProject(args []string) {
	target, err := filepath.Abs(args[0])
	check(err)
	checkTrellisRequirements()
	cloneRepo("trellis", target)
	fmt.Println(target)
}

func main() {
	fmt.Println("hello")
	args := os.Args[1:]

	command := args[0]

	if command == "new" {
		scaffoldProject(args[1:])
	}
}
