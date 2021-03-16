run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/sf-golang-7-2-darwin .
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o bin/sf-golang-7-2-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/sf-golang-7-2-win64.exe .
