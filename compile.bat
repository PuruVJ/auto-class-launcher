SET GOOS = windows
SET GOARCH = amd
go build -o bin/auto-class-launcher-windows.exe index.go

SET GOOS = darwin
SET GOARCH = arm
go build -o bin/auto-class-launcher-macos-m1 index.go

SET GOOS = darwin
SET GOARCH = amd
go build -o bin/auto-class-launcher-macos-intel index.go

SET GOOS = linux
SET GOARCH = amd
go build -o bin/auto-class-launcher-linux index.go