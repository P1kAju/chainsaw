CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o build/chainsaw_linux_386
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/chainsaw_linux_amd64
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/chainsaw_windows_386.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/chainsaw_windows_amd64.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o build/chainsaw_darwin_386
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/chainsaw_darwin_amd64