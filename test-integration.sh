#!/usr/bin/env bash
go build
go test ./... -tags=integration
rm scheme
