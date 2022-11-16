package main

import (
	"fmt"
	"github.com/shanahanjrs/dro/internal/drivers"
	"github.com/shanahanjrs/dro/internal/utils"
	"os"
	"os/exec"
	"strings"
)

func getActionCmdFromDriver(actionName string, driver *drivers.Driver) []string {
	switch actionName {
	case "install":
		return driver.InstallCmd
	case "uninstall":
		return driver.UninstallCmd
	case "search":
		return driver.SearchCmd
	case "list":
		return driver.ListInstalledCmd
	default:
		return []string{}
	}
}

func getPackages() []string {
	args := os.Args[2:]

	if len(args) < 1 {
		utils.Help()
		os.Exit(1)
	}

	return args
}

func checkHelp() {
	helpStrings := []string{"-h", "--help", "help"}

	for _, val := range os.Args[1:] {
		if utils.In(val, helpStrings) {
			utils.Help()
			os.Exit(0)
		}
	}
}

func checkListSupported() {
	if utils.In("--list-supported", os.Args) {
		for _, pkgMngr := range drivers.GetSupportedPackageManagers() {
			fmt.Println(pkgMngr)
		}
		os.Exit(0)
	}
}

func checkVersion() {
	versionStrings := []string{"version", "--version"}

	for _, val := range os.Args[1:] {
		if utils.In(val, versionStrings) {
			utils.GetVersion()
			os.Exit(0)
		}
	}
}

func main() {
	// what does the user want to do; (un)install, search, help, version, etc
	actionName := utils.GetAction()
	checkHelp()
	checkListSupported()
	checkVersion()

	// which package manager should we use
	driverName, err := utils.GetBasePackageManagerName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// load up the correct Driver
	driver, err := drivers.LoadDriver(driverName)
	if err != nil {
		fmt.Println("Could not find supported driver")
		os.Exit(1)
	}

	action := getActionCmdFromDriver(actionName, &driver)

	// start to build up arg string
	var cmdArgs []string
	cmdArgs = append(cmdArgs, driver.Cmd)
	cmdArgs = append(cmdArgs, action...)

	// if the action the user specified requires them to specify package names
	// we need to collect them and add them to the command
	if utils.DoesActionRequirePackageList(actionName) {
		pkgs := getPackages()
		for _, val := range pkgs {
			cmdArgs = append(cmdArgs, val)
		}
	}

	cmdString := strings.Join(cmdArgs, " ")

	// prep the cmd object to be used and connect the child proc to the users term
	cmd := exec.Command("/bin/sh", "-c", cmdString)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run
	_ = cmd.Run()
}
