package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteWAV(qid string) (string, error) {
	b, err := GetWAV(qid)

	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println(len(b))
	fname := uuid.NewV4().String()
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.wav", viper.GetString("cache.tempdir"), fname), b, 0777)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("%s/%s.wav", viper.GetString("cache.tempdir"), fname), err
}

func WriteVTT(qid string) (string, error) {
	b, err := GetVTT(qid)

	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println(len(b))
	fname := uuid.NewV4().String()
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.vtt", viper.GetString("cache.tempdir"), fname), b, 0777)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("%s/%s.vtt", viper.GetString("cache.tempdir"), fname), err
}
