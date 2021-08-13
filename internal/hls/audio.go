package hls

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func SegmentAudio(inWavPath string, rid string, segment_duration int) error {
	_, err := os.Stat(inWavPath)
	if os.IsNotExist(err) {
		return err
	}
	audioStream := ffmpeg_go.Input(inWavPath)
	args := ffmpeg_go.KwArgs{
		"muxdelay":       "0",
		"c:a":            "aac",
		"b:a":            "128k",
		"f":              "segment",
		"sc_threshold":   "0",
		"segment_time":   fmt.Sprint(segment_duration),
		"segment_list":   fmt.Sprintf("%s/%s/playlist.m3u8", viper.GetString("cache.static"), rid),
		"segment_format": "mpegts",
	}
	HlSStream := ffmpeg_go.Output([]*ffmpeg_go.Stream{audioStream}, fmt.Sprintf("%s/%s/file%%d.m4a", viper.GetString("cache.static"), rid), args)
	err = HlSStream.Run()
	if err != nil {
		return err
	}

	if err = os.Remove(inWavPath); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
