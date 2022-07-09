build:
	GOOS=windows go build -ldflags "-s -w" -o defaultinator-cli.exe
	GOOS=linux go build -ldflags "-s -w" -o defaultinator-cli
	GOOS=darwin go build -ldflags "-s -w" -o defaultinator-cli-mac
	upx defaultinator-cli*
