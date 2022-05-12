package db

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"gopkg.in/stash.v1"
)

func GetWAV(qid string) ([]byte, error) {
	rc, err := GoStreamer.StashClient.Get(fmt.Sprintf("q:wav:%s", qid))
	if err == nil {
		log.Println("Cache hit!")
		b, err := io.ReadAll(rc)
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		return b, nil
	} else if err == stash.ErrNotFound {
		r, err := http.Get(fmt.Sprintf("%s/%s?batch", viper.GetString("url.audio"), qid))
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		err = GoStreamer.StashClient.Put(fmt.Sprintf("q:wav:%s", qid), b)
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		return b, nil
	} else {
		return nil, err
	}
}

func GetVTT(qid string) ([]byte, error) {
	rc, err := GoStreamer.StashClient.Get(fmt.Sprintf("q:vtt:%s", qid))
	if err == nil {
		log.Println("Cache hit!")
		b, err := io.ReadAll(rc)
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		return b, nil
	} else if err == stash.ErrNotFound {
		r, err := http.Get(fmt.Sprintf("%s/%s?batch", viper.GetString("url.vtt"), qid))
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		err = GoStreamer.StashClient.Put(fmt.Sprintf("q:vtt:%s", qid), b)
		if err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		return b, nil
	} else {
		return nil, err
	}
}
