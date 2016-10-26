#!/bin/bash

go tool vet -all -shadowstrict . 2>&1 | grep -v ^vendor | reviewdog -f=govet -ci="circle-ci"
go list ./... | grep -v vendor/ | golint | reviewdog -f=golint -ci="circle-ci"