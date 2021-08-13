install:
	go mod download
	go build main.go

depend:
	ffmpeg -version >/dev/null 2>&1  || echo "ffmpeg is not installed, Install using sudo apt get install ffmpeg"

build:
	go build main.go