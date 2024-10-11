package main

import (
	"bytes"
	"errors"
)

type JSONDiff struct {
	File1 File
	File2 File
}

func (j JSONDiff) FindDifferences() (string, error) {
	if j.File1.Bytes == nil || j.File2.Bytes == nil {
		return "No bytes defined for File1 and/or File2.", nil
	}

	if j.File1.Map == nil || j.File2.Map == nil {
		return "No map defined for File1 and/or File2.", nil
	}

	if bytes.Equal(j.File1.Bytes, j.File2.Bytes) {
		return "No differences found.", nil
	}

	output, err := check(j.File1.Map, j.File2.Map)
	if err != nil {
		return "", err
	}

	return output, errors.New("new vulnerability(s) found")
}
