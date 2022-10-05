![](assets/logo.png "dro")

---

## what

**dro** :: The tiny universal remote for package managers

- Very small 1.5MB binary
- Provides a practical, universal package manager interface for any system
- Attempts to adhere to the [suckless philosophy](https://suckless.org/philosophy)
- To be used on any/all Linux machines, macOS, inside containers, and in scripts


## why

- Between work and personal machines I have to use over half a dozen package managers. This solves a problem I have.
- It's fun

## install

`curl -sL "https://raw.githubusercontent.com/shanahanjrs/dro/HEAD/scripts/install.sh" | bash`

or

`wget -q "https://raw.githubusercontent.com/shanahanjrs/dro/HEAD/scripts/install.sh" -O - | bash`


## usage

#### install a package
`dro install <pkg>`

#### uninstall a package
`dro uninstall <pkg>`

#### search for a package
`dro search <pkg>`

#### list the package managers dro currently supports
`dro --list-supported`


## documentation

> if the `DRO_PKG_MNGR` env var is not set it will scan the system to find a supported package manager

Supported package managers:

| pkg mngr | install | uninstall | search |
|----------|---------|-----------|--------|
| dnf      | X       | X         | X      |
| apk      | X       | X         | X      |
| brew     | X       | X         | X      |
| pacman   | X       | X         | X      |
| apt      | X       | X         | X      |
| zypper   | X       | X         | X      |


## coming soon

- more package managers, if one you want is missing please send a PR or open an issue
- more commands; update/upgrade, list installed packages, etc
