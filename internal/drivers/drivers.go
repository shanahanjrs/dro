package drivers

import (
	"errors"
)

type Driver struct {
	Cmd          string
	InstallCmd   string
	UninstallCmd string
	SearchCmd    string
}

func LoadDriver(driverName string) (Driver, error) {
	switch driverName {
	case "dnf":
		return dnf(), nil
	case "apk":
		return apk(), nil
	case "brew":
		return brew(), nil
	case "pacman":
		return pacman(), nil
	default:
		return Driver{}, errors.New("driver not found")
	}
}

func dnf() Driver {
	return Driver{
		Cmd:          "dnf",
		InstallCmd:   "install",
		UninstallCmd: "remove",
		SearchCmd:    "search",
	}
}

func apk() Driver {
	return Driver{
		Cmd:          "apk",
		InstallCmd:   "add",
		UninstallCmd: "del",
		SearchCmd:    "search",
	}
}

func brew() Driver {
	return Driver{
		Cmd:          "brew",
		InstallCmd:   "install",
		UninstallCmd: "uninstall",
		SearchCmd:    "search",
	}
}

func pacman() Driver {
	return Driver{
		Cmd:          "pacman",
		InstallCmd:   "-S",
		UninstallCmd: "-R",
		SearchCmd:    "-Q",
	}
}
