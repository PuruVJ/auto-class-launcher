env GOOS=windows GOARCH=amd64 go build -o bin/auto-class-launcher-windows.exe index.go
env GOOS=darwin GOARCH=amd64 go build -o bin/auto-class-launcher-macos-intel index.go
env GOOS=darwin GOARCH=arm64 go build -o bin/auto-class-launcher-macos-m1 index.go
env GOOS=linux GOARCH=amd64 go build -o bin/auto-class-launcher-linux index.go
