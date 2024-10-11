package main

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		file1 map[string]interface{}
		file2 map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "No new issues",
			args: args{
				file1: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": []interface{}{
								map[string]interface{}{
									"fingerprints": map[string]interface{}{
										"identity": "12345",
									},
									"level": "note",
									"message": map[string]interface{}{
										"text": "Issue 1",
									},
									"locations": []interface{}{
										map[string]interface{}{
											"physicalLocation": map[string]interface{}{
												"artifactLocation": map[string]interface{}{
													"uri": "file1.go",
												},
												"region": map[string]interface{}{
													"startLine": 10.0,
												},
											},
										},
									},
								},
							},
						},
					},
				},
				file2: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": []interface{}{
								map[string]interface{}{
									"fingerprints": map[string]interface{}{
										"identity": "12345",
									},
									"level": "note",
									"message": map[string]interface{}{
										"text": "Issue 1",
									},
									"locations": []interface{}{
										map[string]interface{}{
											"physicalLocation": map[string]interface{}{
												"artifactLocation": map[string]interface{}{
													"uri": "file1.go",
												},
												"region": map[string]interface{}{
													"startLine": 10.0,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "New issue found",
			args: args{
				file1: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": []interface{}{},
						},
					},
				},
				file2: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": []interface{}{
								map[string]interface{}{
									"fingerprints": map[string]interface{}{
										"identity": "67890",
									},
									"level": "error",
									"message": map[string]interface{}{
										"text": "New Issue",
									},
									"locations": []interface{}{
										map[string]interface{}{
											"physicalLocation": map[string]interface{}{
												"artifactLocation": map[string]interface{}{
													"uri": "file2.go",
												},
												"region": map[string]interface{}{
													"startLine": 20.0,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want:    "✗ Severity: [High]\nPath: file2.go\nStart Line: 20\nMessage: New Issue\n\n",
			wantErr: false,
		},
		{
			name: "Error extracting baseline results",
			args: args{
				file1: map[string]interface{}{
					"invalid_key": "invalid_value",
				},
				file2: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": []interface{}{},
						},
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Error extracting PR results",
			args: args{
				file1: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": []interface{}{},
						},
					},
				},
				file2: map[string]interface{}{
					"invalid_key": "invalid_value",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := check(tt.args.file1, tt.args.file2)
			if (err != nil) != tt.wantErr {
				t.Errorf("check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractResults(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name     string
		args     args
		want     []interface{}
		wantBool bool
	}{
		{
			name: "Invalid data",
			args: args{
				data: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{
							"results": "invalid",
						},
					},
				},
			},
			want:     nil,
			wantBool: false,
		},
		{
			name: "No results",
			args: args{
				data: map[string]interface{}{
					"runs": []interface{}{
						map[string]interface{}{},
					},
				},
			},
			want:     nil,
			wantBool: false,
		},
		{
			name: "Runs exist, but empty",
			args: args{
				data: map[string]interface{}{
					"runs": []interface{}{},
				},
			},
			want:     nil,
			wantBool: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotBool := extractResults(tt.args.data)
			if gotBool != tt.wantBool {
				t.Errorf("extractResults() gotBool = %v, want %v", gotBool, tt.wantBool)
			}
			if gotBool {
				if len(got) != len(tt.want) {
					t.Errorf("extractResults() got = %v, want %v", got, tt.want)
				} else {
					for i := range got {
						if got[i] != tt.want[i] {
							t.Errorf("extractResults() got = %v, want %v", got, tt.want)
						}
					}
				}
			}
		})
	}
}
