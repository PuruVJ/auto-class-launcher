env GOOS=windows GOARCH=amd64 go build -o bin/auto-class-launcher-windows.exe 
env GOOS=darwin GOARCH=amd64 go build -o bin/auto-class-launcher-macos-intel 
env GOOS=darwin GOARCH=arm64 go build -o bin/auto-class-launcher-macos-m1 
env GOOS=linux GOARCH=amd64 go build -o bin/auto-class-launcher-linux 
