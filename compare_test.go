package main

import "testing"

func TestJSONDiff_FindDifferences(t *testing.T) {
	tests := []struct {
		name string
		j    JSONDiff
		want string
	}{
		{
			name: "no differences",
			j: JSONDiff{
				File1: File{
					Bytes: []byte(`{"key1": "value1"}`),
					Map:   map[string]interface{}{"key1": "value1"},
				},
				File2: File{
					Bytes: []byte(`{"key1": "value1"}`),
					Map:   map[string]interface{}{"key1": "value1"},
				},
			},
			want: "No differences found.",
		},
		{
			name: "differences found",
			j: JSONDiff{
				File1: File{
					Bytes: []byte(`{"key1": "value1"}`),
					Map:   map[string]interface{}{"key1": "value1"},
				},
				File2: File{
					Bytes: []byte(`{"key1": "value2"}`),
					Map:   map[string]interface{}{"key1": "value2"},
				},
			},
			want: "Differences found:\nmap[key1]: value1 != value2",
		},
		{
			name: "differences found with nested map",
			j: JSONDiff{
				File1: File{
					Bytes: []byte(`{"key1": {"key2": "value1"}}`),
					Map:   map[string]interface{}{"key1": map[string]interface{}{"key2": "value1"}},
				},
				File2: File{
					Bytes: []byte(`{"key1": {"key2": "value2"}}`),
					Map:   map[string]interface{}{"key1": map[string]interface{}{"key2": "value1"}},
				},
			},
			want: "No differences found.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.j.FindDifferences(); got != tt.want {
				t.Errorf("JSONDiff.FindDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONDiff_FindDifferencesWithSnykShape_Same(t *testing.T) {
	snykTests := []struct {
		name string
		j    JSONDiff
	}{
		{
			name: "snyk test case no differences",
			j: JSONDiff{
				File1: File{
					Bytes: []byte(`{
						"vulnerabilities": [],
						"ok": true,
						"dependencyCount": 0,
						"org": "example-org",
						"policy": "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate": true,
						"licensesPolicy": {
							"severities": {},
							"orgLicenseRules": {
								"AGPL-1.0": {
									"licenseType": "AGPL-1.0",
									"severity": "high",
									"instructions": ""
								}
							}
						},
						"packageManager": "yarn",
						"ignoreSettings": {
							"adminOnly": false,
							"reasonRequired": false,
							"disregardFilesystemIgnores": false
						},
						"summary": "No known vulnerabilities",
						"filesystemPolicy": false,
						"uniqueCount": 0,
						"projectName": "package.json",
						"foundProjectCount": 1,
						"displayTargetFile": "yarn.lock",
						"hasUnknownVersions": false,
						"path": "/path/to/project"
				    }`),
					Map: map[string]interface{}{
						"vulnerabilities": []interface{}{},
						"ok":              true,
						"dependencyCount": 0,
						"org":             "example-org",
						"policy":          "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate":       true,
						"licensesPolicy": map[string]interface{}{
							"severities": map[string]interface{}{},
							"orgLicenseRules": map[string]interface{}{
								"AGPL-1.0": map[string]interface{}{
									"licenseType":  "AGPL-1.0",
									"severity":     "high",
									"instructions": "",
								},
							},
						},
						"packageManager": "yarn",
						"ignoreSettings": map[string]interface{}{
							"adminOnly":                  false,
							"reasonRequired":             false,
							"disregardFilesystemIgnores": false,
						},
						"summary":            "No known vulnerabilities",
						"filesystemPolicy":   false,
						"uniqueCount":        0,
						"projectName":        "package.json",
						"foundProjectCount":  1,
						"displayTargetFile":  "yarn.lock",
						"hasUnknownVersions": false,
						"path":               "/path/to/project",
					},
				},
				File2: File{
					Bytes: []byte(`{
						"vulnerabilities": [],
						"ok": true,
						"dependencyCount": 0,
						"org": "example-org",
						"policy": "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate": true,
						"licensesPolicy": {
							"severities": {},
							"orgLicenseRules": {
								"AGPL-1.0": {
									"licenseType": "AGPL-1.0",
									"severity": "high",
									"instructions": ""
								}
							}
						},
						"packageManager": "yarn",
						"ignoreSettings": {
							"adminOnly": false,
							"reasonRequired": false,
							"disregardFilesystemIgnores": false
						},
						"summary": "No known vulnerabilities",
						"filesystemPolicy": false,
						"uniqueCount": 0,
						"projectName": "package.json",
						"foundProjectCount": 1,
						"displayTargetFile": "yarn.lock",
						"hasUnknownVersions": false,
						"path": "/path/to/project"
				    }`),
					Map: map[string]interface{}{
						"vulnerabilities": []interface{}{},
						"ok":              true,
						"dependencyCount": 0,
						"org":             "example-org",
						"policy":          "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate":       true,
						"licensesPolicy": map[string]interface{}{
							"severities": map[string]interface{}{},
							"orgLicenseRules": map[string]interface{}{
								"AGPL-1.0": map[string]interface{}{
									"licenseType":  "AGPL-1.0",
									"severity":     "high",
									"instructions": "",
								},
							},
						},
						"packageManager": "yarn",
						"ignoreSettings": map[string]interface{}{
							"adminOnly":                  false,
							"reasonRequired":             false,
							"disregardFilesystemIgnores": false,
						},
						"summary":            "No known vulnerabilities",
						"filesystemPolicy":   false,
						"uniqueCount":        0,
						"projectName":        "package.json",
						"foundProjectCount":  1,
						"displayTargetFile":  "yarn.lock",
						"hasUnknownVersions": false,
						"path":               "/path/to/project",
					},
				},
			},
		},
	}

	for _, tt := range snykTests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.j.FindDifferences(); got != "No differences found." {
				t.Errorf("JSONDiff.FindDifferences() = %v)", got)
			}
		})
	}
}

func TestJSONDiff_FindDifferencesWithSnykShape_Vulnerabilities(t *testing.T) {
	snykTests := []struct {
		name string
		j    JSONDiff
	}{
		{
			name: "snyk test case with vulnerabilities",
			j: JSONDiff{
				File1: File{
					Bytes: []byte(`{
						"vulnerabilities": [
						  {
							"id": "vuln-1",
							"title": "Vulnerability 1",
							"severity": "high"
						  }
						],
						"ok": true,
						"dependencyCount": 0,
						"org": "example-org",
						"policy": "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate": true,
						"licensesPolicy": {
							"severities": {},
							"orgLicenseRules": {
								"AGPL-1.0": {
									"licenseType": "AGPL-1.0",
									"severity": "high",
									"instructions": ""
								}
							}
						},
						"packageManager": "yarn",
						"ignoreSettings": {
							"adminOnly": false,
							"reasonRequired": false,
							"disregardFilesystemIgnores": false
						},
						"summary": "No known vulnerabilities",
						"filesystemPolicy": false,
						"uniqueCount": 0,
						"projectName": "package.json",
						"foundProjectCount": 1,
						"displayTargetFile": "yarn.lock",
						"hasUnknownVersions": false,
							"path": "/path/to/project"
				    }`),
					Map: map[string]interface{}{
						"vulnerabilities": []interface{}{
							map[string]interface{}{
								"id":       "vuln-1",
								"title":    "Vulnerability 1",
								"severity": "high",
							},
						},
						"ok":              true,
						"dependencyCount": 0,
						"org":             "example-org",
						"policy":          "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate":       true,
						"licensesPolicy": map[string]interface{}{
							"severities": map[string]interface{}{},
							"orgLicenseRules": map[string]interface{}{
								"AGPL-1.0": map[string]interface{}{
									"licenseType":  "AGPL-1.0",
									"severity":     "high",
									"instructions": "",
								},
							},
						},
						"packageManager": "yarn",
						"ignoreSettings": map[string]interface{}{
							"adminOnly":                  false,
							"reasonRequired":             false,
							"disregardFilesystemIgnores": false,
						},
						"summary":            "No known vulnerabilities",
						"filesystemPolicy":   false,
						"uniqueCount":        0,
						"projectName":        "package.json",
						"foundProjectCount":  1,
						"displayTargetFile":  "yarn.lock",
						"hasUnknownVersions": false,
						"path":               "/path/to/project",
					},
				},
				File2: File{
					Bytes: []byte(`{
						"vulnerabilities": [
						  {
							"id": "vuln-2",
							"title": "Vulnerability 2",
							"severity": "medium"
						  }
						],
						"ok": true,
						"dependencyCount": 0,
						"org": "example-org",
						"policy": "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate": true,
						"licensesPolicy": {
							"severities": {},
							"orgLicenseRules": {
								"AGPL-1.0": {
									"licenseType": "AGPL-1.0",
									"severity": "high",
									"instructions": ""
								}
							}
						},
						"packageManager": "yarn",
						"ignoreSettings": {
							"adminOnly": false,
							"reasonRequired": false,
							"disregardFilesystemIgnores": false
						},
						"summary": "No known vulnerabilities",
						"filesystemPolicy": false,
						"uniqueCount": 0,
						"projectName": "package.json",
						"foundProjectCount": 1,
						"displayTargetFile": "yarn.lock",
						"hasUnknownVersions": false,
						"path": "/path/to/project"
					}`),
					Map: map[string]interface{}{
						"vulnerabilities": []interface{}{
							map[string]interface{}{
								"id":       "vuln-2",
								"title":    "Vulnerability 2",
								"severity": "medium",
							},
						},
						"ok":              true,
						"dependencyCount": 0,
						"org":             "example-org",
						"policy":          "# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities.\nversion: v1.25.1\nignore: {}\npatch: {}\n",
						"isPrivate":       true,
						"licensesPolicy": map[string]interface{}{
							"severities": map[string]interface{}{},
							"orgLicenseRules": map[string]interface{}{
								"AGPL-1.0": map[string]interface{}{
									"licenseType":  "AGPL-1.0",
									"severity":     "high",
									"instructions": "",
								},
							},
						},
						"packageManager": "yarn",
						"ignoreSettings": map[string]interface{}{
							"adminOnly":                  false,
							"reasonRequired":             false,
							"disregardFilesystemIgnores": false,
						},
						"summary":            "No known vulnerabilities",
						"filesystemPolicy":   false,
						"uniqueCount":        0,
						"projectName":        "package.json",
						"foundProjectCount":  1,
						"displayTargetFile":  "yarn.lock",
						"hasUnknownVersions": false,
						"path":               "/path/to/project",
					},
				},
			},
		},
	}

	for _, tt := range snykTests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.j.FindDifferences(); got == "No differences found." {
				t.Errorf("JSONDiff.FindDifferences() = %v)", got)
			}
		})
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
