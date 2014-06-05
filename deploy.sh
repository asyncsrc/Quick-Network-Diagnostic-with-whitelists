#!/bin/bash

GOOS=windows GOARCH=386 go build . && mv whitelist_scanner2.exe "Windows-32 PM Whitelist Verification.exe" && zip "Windows-32 PM Whitelist Verification.zip" "Windows-32 PM Whitelist Verification.exe" && rm "Windows-32 PM Whitelist Verification.exe"
GOOS=windows GOARCH=amd64 go build . && mv whitelist_scanner2.exe "Windows-64 PM Whitelist Verification.exe" && zip "Windows-64 PM Whitelist Verification.zip" "Windows-64 PM Whitelist Verification.exe" && rm "Windows-64 PM Whitelist Verification.exe"
GOOS=linux GOARCH=amd64 go build . && mv whitelist_scanner2 "Linux-64 PM Whitelist Verification" && zip "Linux-64 PM Whitelist Verification.zip" "Linux-64 PM Whitelist Verification" && rm "Linux-64 PM Whitelist Verification"
GOOS=linux GOARCH=386 go build . && mv whitelist_scanner2 "Linux-32 PM Whitelist Verification" && zip "Linux-32 PM Whitelist Verification.zip" "Linux-32 PM Whitelist Verification" && rm "Linux-32 PM Whitelist Verification"
GOOS=darwin GOARCH=amd64 go build . && mv whitelist_scanner2 "MacOS PM Whitelist Verification" && zip "MacOS PM Whitelist Verification.zip" "MacOS PM Whitelist Verification" && rm "MacOS PM Whitelist Verification"
