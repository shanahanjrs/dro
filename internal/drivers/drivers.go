package drivers

import (
	"errors"
)

type Driver struct {
	Cmd              string
	InstallCmd       []string
	UninstallCmd     []string
	SearchCmd        []string
	ListInstalledCmd []string
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
	case "apt":
		return apt(), nil
	case "zypper":
		return zypper(), nil
	default:
		return Driver{}, errors.New("driver not found")
	}
}

func GetSupportedPackageManagers() []string {
	return []string{
		"dnf",
		"apk",
		"brew",
		"pacman",
		"apt",
		"zypper",
	}
}

func dnf() Driver {
	return Driver{
		Cmd:              "dnf",
		InstallCmd:       []string{"install"},
		UninstallCmd:     []string{"remove"},
		SearchCmd:        []string{"search"},
		ListInstalledCmd: []string{"list --installed"},
	}
}

func apk() Driver {
	return Driver{
		Cmd:              "apk",
		InstallCmd:       []string{"add"},
		UninstallCmd:     []string{"del"},
		SearchCmd:        []string{"search"},
		ListInstalledCmd: []string{"info"},
	}
}

func brew() Driver {
	return Driver{
		Cmd:              "brew",
		InstallCmd:       []string{"install"},
		UninstallCmd:     []string{"uninstall"},
		SearchCmd:        []string{"search"},
		ListInstalledCmd: []string{"list"},
	}
}

func pacman() Driver {
	return Driver{
		Cmd:              "pacman",
		InstallCmd:       []string{"-S"},
		UninstallCmd:     []string{"-R"},
		SearchCmd:        []string{"-Q"},
		ListInstalledCmd: []string{"-Qe"},
	}
}

func apt() Driver {
	return Driver{
		Cmd:              "apt",
		InstallCmd:       []string{"install"},
		UninstallCmd:     []string{"remove"},
		SearchCmd:        []string{"search"},
		ListInstalledCmd: []string{"list --installed"},
	}
}

func zypper() Driver {
	return Driver{
		Cmd:              "zypper",
		InstallCmd:       []string{"in"},
		UninstallCmd:     []string{"rm"},
		SearchCmd:        []string{"se"},
		ListInstalledCmd: []string{"se -i"},
	}
}
