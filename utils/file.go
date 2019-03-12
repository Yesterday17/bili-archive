package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func WriteFile(fpath, fname string, data []byte) error {
	return ioutil.WriteFile(path.Join(fpath, fname), data, 0644)
}

func WriteJson(fpath, fname string, data interface{}) error {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return WriteFile(fpath, fname, json)
}

func WriteJsonS(fpath, fname string, data interface{}) error {
	if !FileExist(fpath, fname) {
		return WriteJson(fpath, fname, data)
	}
	return nil
}

func WriteLockFile(fpath, fname string) error {
	return WriteFile(fpath, fname, []byte{})
}

func MkDir(fpath string) error {
	return os.MkdirAll(fpath, os.ModePerm)
}

func MKDir(fpath string) {
	if err := MkDir(fpath); err != nil {
		log.Println(fmt.Sprintf("[MK] %s", err))
	}
}

func MKDirs(paths ...string) {
	for _, path := range paths {
		MKDir(path)
	}
}

func FileExist(fpath, fname string) bool {
	_, err := os.Stat(path.Join(fpath, fname))
	return !os.IsNotExist(err)
}

func FilePath(fpath, fname, ext string) string {
	return path.Join(fpath, fname+"."+ext)
}

func FileSize(filePath string) (int64, bool, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return file.Size(), true, nil
}
