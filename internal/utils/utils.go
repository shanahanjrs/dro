package utils

import (
	"errors"
	"fmt"
	"github.com/shanahanjrs/dro/internal/drivers"
	"os"
	"os/exec"
)

// In - Checks if <needle> exists in <haystack>
// ("dog", ["cat" "dog" "bird"]) --> true
func In[T comparable](needle T, haystack []T) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

// GetEnv - gets the value of the env var or returns default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Which - a dirty version of the unix which command
func Which(command string) bool {
	_, err := exec.Command("which", command).Output()
	if err != nil {
		return false
	}
	return true

}

var GetValidActions = []string{
	"install",
	"uninstall",
	"search",
	"list",
	"--list-supported",
	"--help", "-h",
}

// GetAction
// handles counting input args, making sure it's actually valid and then returning the string
func GetAction() string {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Not enough args...")
		os.Exit(1)
	}

	action := os.Args[1]
	if !In(action, GetValidActions) {
		fmt.Println("Invalid action " + action + "...")
	}

	return action
}

// DoesActionRequirePackageList
// takes an action and returns t/f whether we need to check
// for packages specified on the cli by the user
// required: install, uninstall, search
// !required: list, etc...
func DoesActionRequirePackageList(action string) bool {
	requiresPackageList := []string{
		"install",
		"uninstall",
		"search",
	}

	if In(action, requiresPackageList) {
		return true
	}
	return false
}

// GetBasePackageManagerName
// will return the name of a package manager to the caller
// decides which name to return by first checking if one is specified via
// the DRO_PKG_MNGR env var, otherwise it'll loop through a list of supported
// package managers until it finds the first supported one that is installed
func GetBasePackageManagerName() (string, error) {
	if droPkgMngr := GetEnv("DRO_PKG_MNGR", ""); len(droPkgMngr) > 0 {
		return droPkgMngr, nil
	}

	for _, val := range drivers.GetSupportedPackageManagers() {
		if Which(val) == true {
			return val, nil
		}
	}

	return "", errors.New("no installed package manager found")
}
