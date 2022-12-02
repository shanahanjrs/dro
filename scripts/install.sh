#!/bin/sh

bail() {
  echo "Something went wrong: $1"
  exit 1
}

command_exists() {
	command -v "$@" > /dev/null 2>&1
}

install_dro() {
  if command_exists "wget"; then
    wget -qO- https://raw.githubusercontent.com/shanahanjrs/dro/master/bin/dro-"$(uname)"-"$(uname -m)" -o /tmp/dro || bail "Could not wget dro install script"
  elif command_exists "curl"; then
    curl -fsSL https://raw.githubusercontent.com/shanahanjrs/dro/master/bin/dro-"$(uname)"-"$(uname -m)" -o /tmp/dro || bail "Could not curl dro install script"
  else
    bail "This script requires wget or curl be installed"
  fi

  chmod +x /tmp/dro || bail

  sudo mv /tmp/dro /usr/local/bin
}

install_dro