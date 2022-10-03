package main

import (
	"fmt"
	"github.com/shanahanjrs/dro/internal/drivers"
	"github.com/shanahanjrs/dro/internal/utils"
	"os"
	"os/exec"
)

const helpString = `   _
 _| |___ ___
| . |  _| . |
|___|_| |___|

Author John Shanahan <shanahan.jrs@gmail.com>

Usage
    dro [action] [package(s)]

Actions
    install    | install chosen package(s)
    uninstall  | uninstall chosen package(s)
    search     | searches for package

Example
    dro install vim git
`

var validActions = []string{
	"install",
	"uninstall",
	"search",
}

// prints the help string
func help() {
	fmt.Println(helpString)
}

func getAction(driver *drivers.Driver) string {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Not enough args...")
		os.Exit(1)
	}

	// install/uninstall/search/etc...
	action := args[0]
	if !utils.In(action, validActions) {
		fmt.Println("Please provide a valid action...")
		os.Exit(1)
	}

	switch action {
	case "install":
		return driver.InstallCmd
	case "uninstall":
		return driver.UninstallCmd
	case "search":
		return driver.SearchCmd
	default:
		return ""
	}
}

func getPackages() []string {
	args := os.Args[2:]

	if len(args) < 1 {
		fmt.Println("Not enough args...")
		help()
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

func main() {
	checkHelp()

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

	action := getAction(&driver)

	pkgs := getPackages()

	// build up arg string
	var cmdArgs []string
	cmdArgs = append(cmdArgs, action)

	for _, val := range pkgs {
		cmdArgs = append(cmdArgs, val)
	}

	cmd := exec.Command(driver.Cmd, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
}
