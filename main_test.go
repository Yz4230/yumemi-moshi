package main

import (
	"os"
	"testing"
)

func TestValidateHeader(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"#", false},
		{"create_timestamp,", false},
		{"create_timestamp,palyer_id", false},
		{"create_timestamp,player_id,score", true},
		{"create_timestamp,palyer_id,score,score", false},
		{"create_timestamp,player_id,score,score2", false},
	}
	for _, test := range tests {
		got := validateHeader(test.input)
		if got != test.want {
			t.Errorf("validateHeader(%q) = %t; shoud be %t", test.input, got, test.want)
		}
	}
}

func TestValidatePlayerID(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"&&", false},
		{"-1", false},
		{"0", true},
		{"1", true},
		{"123abcABC", true},
	}
	for _, test := range tests {
		got := validatePlayerID(test.input)
		if got != test.want {
			t.Errorf("validatePlayerID(%q) = %t; shoud be %t", test.input, got, test.want)
		}
	}
}

func TestParseCSV(t *testing.T) {
	tests := []struct {
		filename           string
		finishSuccessFully bool
	}{
		{"./sample1.csv", true},
	}
	for _, test := range tests {
		f, err := os.Open(test.filename)
		if err != nil {
			t.Errorf("Failed to open file %s", test.filename)
		}
		defer f.Close()
		_, err = parseCSV(f)
		if err != nil {
			t.Errorf("Failed to parse file %s: %s", test.filename, err)
		}
	}
}
