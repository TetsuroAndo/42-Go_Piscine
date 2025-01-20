#!/bin/bash

for i in {0..6}
do
    dir="ex0$i"
    if [ -d "$dir" ]; then
        cd "$dir"
        if [ -f "go.mod" ]; then
            rm go.mod
        fi
        go mod init "ex0$i"
        cd ..
    fi
done