package main

import "testing"

func TestJSONDiff_FindDifferences_WithDifferences_WrongShape(t *testing.T) {
	j := JSONDiff{
		File1: File{
			Bytes: []byte(`{"key1": "value1"}`),
			Map:   map[string]interface{}{"key1": "value1"},
		},
		File2: File{
			Bytes: []byte(`{"key1": "value2"}`),
			Map:   map[string]interface{}{"key1": "value2"},
		},
	}

	_, err := j.FindDifferences()
	if err == nil {
		t.Errorf("JSONDiff.FindDifferences() error = %v", err)
	}
}

func TestJSONDiff_FindDifferences_NoBytes(t *testing.T) {
	j := JSONDiff{
		File1: File{
			Bytes: nil,
			Map:   map[string]interface{}{"key1": "value1"},
		},
		File2: File{
			Bytes: []byte(`{"key1": "value1"}`),
			Map:   map[string]interface{}{"key1": "value1"},
		},
	}

	if got, _ := j.FindDifferences(); got != "No bytes defined for File1 and/or File2." {
		t.Errorf("JSONDiff.FindDifferences() = %v, want %v", got, "No bytes defined for File1 and/or File2.")
	}
}

func TestJSONDiff_FindDifferences_NoMap(t *testing.T) {
	j := JSONDiff{
		File1: File{
			Bytes: []byte(`{"key1": "value1"}`),
			Map:   nil,
		},
		File2: File{
			Bytes: []byte(`{"key1": "value1"}`),
			Map:   map[string]interface{}{"key1": "value1"},
		},
	}

	if got, _ := j.FindDifferences(); got != "No map defined for File1 and/or File2." {
		t.Errorf("JSONDiff.FindDifferences() = %v, want %v", got, "No map defined for File1 and/or File2.")
	}
}
