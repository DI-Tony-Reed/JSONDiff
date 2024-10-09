package main

import (
	"testing"
	"testing/fstest"
)

func TestFile_ReadFile(t *testing.T) {
	m := fstest.MapFS{
		"file.json": {
			Data: []byte(`{"key": "value"}`),
		},
		"invalid.json": {
			Data: []byte(`{"key": "value"`),
		},
	}

	file := File{
		Reader: m,
	}
	err := file.ReadFile("file.json")
	if err != nil {
		t.Errorf("File.ReadFile() error = %v", err)
	}

	err = file.ReadFile("invalid.json")
	if err == nil {
		t.Errorf("File.ReadFile() error = %v, want error", err)
	}

	err = file.ReadFile("notfound.json")
	if err == nil {
		t.Errorf("File.ReadFile() error = %v, want error", err)
	}
}
