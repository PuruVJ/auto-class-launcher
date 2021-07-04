GOOS=windows GOARCH=amd64 go build -o bin/auto-class-launcher-windows.exe index.go
GOOS=darwin GOARCH=arm64 go build -o bin/auto-class-launcher-windows.exe index.go
GOOS=darwin GOARCH=amd64 go build -o bin/auto-class-launcher-windows.exe index.go
GOOS=linux GOARCH=amd64 go build -o bin/auto-class-launcher-windows.exe index.go
