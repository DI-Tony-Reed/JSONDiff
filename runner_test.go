package main

import (
	"errors"
	"fmt"
	"testing"
)

type MockFileReader struct {
	Content []byte
	Err     error
}

func (m MockFileReader) ReadFile(path string) ([]byte, error) {
	return m.Content, m.Err
}

func TestRunner_Run_NoArgs(t *testing.T) {
	runner := Runner{
		Arguments: []string{"runner"},
	}

	_, err := runner.Run(MockFileReader{})

	expected := fmt.Errorf("Please provide two JSON files to compare.\n")
	if errors.Is(err, expected) {
		t.Errorf("Expected %q, but got %q", expected, expected)
	}
}

func TestRunner_Run_ValidFiles(t *testing.T) {
	mockFileReader := MockFileReader{
		Content: []byte(`{"key1": "value1"}`),
		Err:     nil,
	}

	runner := Runner{
		Arguments: []string{"file1.json", "file2.json"},
	}

	output, err := runner.Run(mockFileReader)

	expected := "No differences found."
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestRunner_Run_InvalidFiles(t *testing.T) {
	mockFileReader := MockFileReader{
		Content: nil,
		Err:     nil,
	}

	runner := Runner{
		Arguments: []string{"file1.json", "file2.json"},
	}

	_, err := runner.Run(mockFileReader)

	if err == nil {
		t.Errorf("Expected error, but got did not get one")
	}
}

func TestRunner_Run_ByteSkip(t *testing.T) {
	mockFileReader := MockFileReader{
		Content: []byte(`{"key1": "value1"}`),
		Err:     nil,
	}

	runner := Runner{
		Arguments: []string{"file1.json", "file2.json", "--byteskip"},
	}

	_, err := runner.Run(mockFileReader)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
}
