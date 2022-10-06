package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(date interface{}) error
}

type Type string

const (
	FyleType Type = "file"
	MonoType Type = "mongo"
)

func New(store Type, fileName string) Store {
	switch store {
	case FyleType:
		return &fileStore{fileName}
	}
	return nil
}

type fileStore struct {
	path string
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.path, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
