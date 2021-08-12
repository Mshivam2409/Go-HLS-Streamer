package main

import (
	"log"

	"github.com/Mshivam2409/hls-streamer/internal/cmd"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	cmd.Execute()
}
