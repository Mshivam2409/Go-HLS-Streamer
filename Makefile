install:
	go mod download
	go build main.go

depend:
	ffmpeg -version >/dev/null 2>&1  || sudo apt install ffmpeg

build:
	go build main.go