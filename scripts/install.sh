#!/bin/bash

curl -sL https://raw.githubusercontent.com/shanahanjrs/dro/master/bin/dro-"$(uname)"-"$(uname -m)" -o /tmp/dro || exit

chmod +x /tmp/dro || exit

sudo mv /tmp/dro /usr/local/bin
