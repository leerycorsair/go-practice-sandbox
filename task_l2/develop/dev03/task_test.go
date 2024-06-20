package main

import (
	"flag"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMySort(t *testing.T) {
	tests := []struct {
		name    string
		lines   []string
		conf    Config
		want    []string
		wantErr bool
	}{
		{
			name:  "simple string sort",
			lines: []string{"banana", "apple", "cherry"},
			conf:  Config{Key: 1},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:  "numeric sort",
			lines: []string{"10", "1", "20", "2"},
			conf:  Config{Key: 1, Numeric: true},
			want:  []string{"1", "2", "10", "20"},
		},
		{
			name:  "reverse string sort",
			lines: []string{"banana", "apple", "cherry"},
			conf:  Config{Key: 1, Reverse: true},
			want:  []string{"cherry", "banana", "apple"},
		},
		{
			name:  "unique lines",
			lines: []string{"apple", "banana", "apple", "cherry", "banana"},
			conf:  Config{Key: 1, Unique: true},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:  "month sort",
			lines: []string{"March", "January", "February"},
			conf:  Config{Key: 1, Month: true},
			want:  []string{"January", "February", "March"},
		},
		{
			name:  "mixed month and non-month values",
			lines: []string{"March", "banana", "January", "apple", "February"},
			conf:  Config{Key: 1, Month: true},
			want:  []string{"January", "February", "March", "banana", "apple"},
		},
		{
			name:  "case insensitive month sort",
			lines: []string{"march", "January", "february"},
			conf:  Config{Key: 1, Month: true},
			want:  []string{"January", "february", "march"},
		},
		{
			name:  "month sort with invalid month names",
			lines: []string{"March", "apple", "June", "banana", "Feb"},
			conf:  Config{Key: 1, Month: true},
			want:  []string{"March", "June", "apple", "banana", "Feb"},
		},
		{
			name:  "ignore leading blanks",
			lines: []string{"  banana", " apple", "  cherry"},
			conf:  Config{Key: 1, Blanks: true},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:    "check sorted",
			lines:   []string{"apple", "banana", "cherry"},
			conf:    Config{Key: 1, Check: true},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "empty input",
			lines: []string{},
			conf:  Config{Key: 1},
			want:  []string{},
		},
		{
			name:  "all identical lines",
			lines: []string{"apple", "apple", "apple"},
			conf:  Config{Key: 1},
			want:  []string{"apple", "apple", "apple"},
		},
		{
			name:  "mixed case sensitivity",
			lines: []string{"Banana", "apple", "Cherry"},
			conf:  Config{Key: 1},
			want:  []string{"Banana", "Cherry", "apple"},
		},
		{
			name:  "trailing spaces without blanks option",
			lines: []string{"banana ", "apple", "cherry "},
			conf:  Config{Key: 1},
			want:  []string{"apple", "banana ", "cherry "},
		},
		{
			name:  "complex numeric values",
			lines: []string{"3.14", "2.718", "1.618"},
			conf:  Config{Key: 1, Numeric: true},
			want:  []string{"1.618", "2.718", "3.14"},
		},
		{
			name:  "invalid numeric values",
			lines: []string{"10", "apple", "20", "banana"},
			conf:  Config{Key: 1, Numeric: true},
			want:  []string{"10", "20", "apple", "banana"},
		},
		{
			name:  "special characters",
			lines: []string{"apple!", "apple#", "apple$"},
			conf:  Config{Key: 1},
			want:  []string{"apple!", "apple#", "apple$"},
		},
		{
			name:    "check already sorted unique",
			lines:   []string{"apple", "banana", "cherry"},
			conf:    Config{Key: 1, Check: true, Unique: true},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mySort(tt.lines, &tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("mySort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		flags    []string
		expected []string
	}{
		{
			name:     "app simple string sort",
			input:    []string{"banana", "apple", "cherry"},
			flags:    []string{"-k", "1"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "app numeric sort",
			input:    []string{"10", "1", "20", "2"},
			flags:    []string{"-k", "1", "-n"},
			expected: []string{"1", "2", "10", "20"},
		},
		{
			name:     "app reverse string sort",
			input:    []string{"banana", "apple", "cherry"},
			flags:    []string{"-k", "1", "-r"},
			expected: []string{"cherry", "banana", "apple"},
		},
		{
			name:     "app unique lines",
			input:    []string{"apple", "banana", "apple", "cherry", "banana"},
			flags:    []string{"-k", "1", "-u"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "app month sort",
			input:    []string{"March", "January", "February"},
			flags:    []string{"-k", "1", "-M"},
			expected: []string{"January", "February", "March"},
		},
		{
			name:     "app ignore leading blanks",
			input:    []string{"  banana", " apple", "  cherry"},
			flags:    []string{"-k", "1", "-b"},
			expected: []string{"apple", "banana", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile, err := os.CreateTemp("", "input.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(inputFile.Name())

			for _, line := range tt.input {
				inputFile.WriteString(line + "\n")
			}
			inputFile.Close()

			outputFile, err := os.CreateTemp("", "output.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(outputFile.Name())

			os.Args = append([]string{"cmd"}, append(tt.flags, inputFile.Name(), outputFile.Name())...)
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			main()

			outputData, err := os.ReadFile(outputFile.Name())
			if err != nil {
				t.Fatal(err)
			}

			got := strings.Split(strings.TrimSpace(string(outputData)), "\n")
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
