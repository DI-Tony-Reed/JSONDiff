package main

import (
	"errors"
)

type Runner struct {
	Arguments []string
}

func (r Runner) Run(reader FileReader) (string, error) {
	if len(r.Arguments) < 2 {
		return "", errors.New("please provide two JSON files to compare")
	}

	fileArgs := r.Arguments[0:2]
	var files []File
	for _, arg := range fileArgs {
		file := File{
			Reader: reader,
		}
		err := file.ReadFile(arg)
		if err != nil {
			return "", err
		}
		files = append(files, file)
	}

	comparator := JSONDiff{
		File1: files[0],
		File2: files[1],
	}

	return comparator.FindDifferences()
}
