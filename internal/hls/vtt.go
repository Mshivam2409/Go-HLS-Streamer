package hls

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func SegmentVTT(inVTTPath string, rid string, segment_duration int) error {
	_, err := os.Stat(inVTTPath)
	if os.IsNotExist(err) {
		return err
	}
	audioStream := ffmpeg_go.Input(inVTTPath)
	args := ffmpeg_go.KwArgs{
		"f":                 "segment",
		"sc_threshold":      "0",
		"segment_timex":      fmt.Sprint(segment_duration),
		"segment_list":      fmt.Sprintf("%s/%s/sub.m3u8", viper.GetString("cache.static"), rid),
		"segment_list_size": 0,
		"segment_format":    "webvtt",
		"scodec":            "copy",
	}
	HlSStream := ffmpeg_go.Output([]*ffmpeg_go.Stream{audioStream}, fmt.Sprintf("%s/%s/sub%%d.vtt", viper.GetString("cache.static"), rid), args)
	err = HlSStream.Run()

	if err != nil {
		return err
	}

	if err = os.Remove(inVTTPath); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
