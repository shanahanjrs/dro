package main

import (
	"fmt"
	"github.com/shanahanjrs/dro/internal/drivers"
	"github.com/shanahanjrs/dro/internal/utils"
	"os"
	"os/exec"
	"strings"
)

const helpString = `   _
 _| |___ ___
| . |  _| . |
|___|_| |___|

Author John Shanahan <shanahan.jrs@gmail.com>

Usage
    dro [action] [package(s)]

Actions
    install           | install chosen package(s)
    uninstall         | uninstall chosen package(s)
    search            | searches for package

    --list-supported  | list supported package managers
    -h, --help        | print this message

Example
    dro install vim git
`

// prints the help string
func help() {
	fmt.Println(helpString)
}

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
		fmt.Println("Not enough args...")
		os.Exit(1)
	}

	return args
}

func checkHelp() {
	for _, val := range os.Args[1:] {
		if (val == "-h") || (val == "--help") {
			help()
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

func main() {
	actionName := utils.GetAction()
	checkHelp()
	checkListSupported()

	driverName, err := utils.GetBasePackageManagerName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

	cmd := exec.Command("/bin/sh", "-c", cmdString)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
}
