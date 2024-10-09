package main

import (
	"encoding/json"
	"errors"
	"os"
)

type FileReader interface {
	ReadFile(name string) ([]byte, error)
}

type OSFileReader struct{}

func (o OSFileReader) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

type File struct {
	Bytes  []byte
	Map    map[string]interface{}
	Reader FileReader
}

func (f *File) ReadFile(filePath string) error {
	file, err := f.Reader.ReadFile(filePath)
	if err != nil {
		return errors.New("error reading file: \n" + err.Error())
	}
	f.Bytes = file

	f.Map = make(map[string]interface{})
	err = json.Unmarshal(f.Bytes, &f.Map)
	if err != nil {
		return errors.New("error unmarshalling file: \n" + err.Error())
	}

	return nil
}
