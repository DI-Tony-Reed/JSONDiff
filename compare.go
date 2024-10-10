package main

import (
	"bytes"
	"fmt"
	"github.com/go-test/deep"
)

type JSONDiff struct {
	File1    File
	File2    File
	ByteSkip bool
}

func (j JSONDiff) FindDifferences() string {
	if j.File1.Bytes == nil || j.File2.Bytes == nil {
		return "No bytes defined for File1 and/or File2."
	}

	if j.File1.Map == nil || j.File2.Map == nil {
		return "No map defined for File1 and/or File2."
	}

	if bytes.Equal(j.File1.Bytes, j.File2.Bytes) {
		return "No differences found."
	}

	if j.ByteSkip && len(j.File2.Bytes) < len(j.File1.Bytes) {
		return "Second file smaller than first and byteskip enabled"
	}

	if diff := deep.Equal(j.File1.Map, j.File2.Map); diff != nil {
		differences := "Differences found:"
		for _, d := range diff {
			differences += fmt.Sprintf("\n%v", d)
		}

		return differences
	}

	return "No differences found."
}
