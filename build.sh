#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build cmd/vortex/main.go && cp main.exe /home/user/shared/
