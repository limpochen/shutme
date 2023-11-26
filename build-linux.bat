@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
echo on
go mod tidy
rem go generate
go build -ldflags "-s -w"
@echo off
SET CGO_ENABLED=
SET GOOS=
SET GOARCH=
