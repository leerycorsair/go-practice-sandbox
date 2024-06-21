package main

import (
	"reflect"
	"testing"
)

func TestFindMatches(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		pattern  string
		conf     Config
		expected interface{}
	}{
		{
			name:    "simple match",
			lines:   []string{"apple", "banana", "cherry", "apple pie"},
			pattern: "apple",
			conf:    Config{},
			expected: map[int]string{
				0: "apple",
				3: "apple pie",
			},
		},
		{
			name:    "ignore case",
			lines:   []string{"Apple", "banana", "CHERRY", "apple pie"},
			pattern: "apple",
			conf:    Config{IgnoreCase: true},
			expected: map[int]string{
				0: "Apple",
				3: "apple pie",
			},
		},
		{
			name:    "fixed string",
			lines:   []string{"apple", "banana", "apple", "apple pie"},
			pattern: "apple",
			conf:    Config{Fixed: true},
			expected: map[int]string{
				0: "apple",
				2: "apple",
			},
		},
		{
			name:    "invert match",
			lines:   []string{"apple", "banana", "cherry", "apple pie"},
			pattern: "apple",
			conf:    Config{Invert: true},
			expected: map[int]string{
				1: "banana",
				2: "cherry",
			},
		},
		{
			name:    "context before and after",
			lines:   []string{"one", "two", "three", "four", "five"},
			pattern: "three",
			conf:    Config{Before: 1, After: 1},
			expected: map[int]string{
				1: "two",
				2: "three",
				3: "four",
			},
		},
		{
			name:    "line numbers",
			lines:   []string{"apple", "banana", "cherry", "apple pie"},
			pattern: "apple",
			conf:    Config{LineNum: true},
			expected: map[int]string{
				0: "apple",
				3: "apple pie",
			},
		},
		{
			name:     "count matches",
			lines:    []string{"apple", "banana", "cherry", "apple pie"},
			pattern:  "apple",
			conf:     Config{Count: true},
			expected: 2,
		},
		{
			name:    "regexp match",
			lines:   []string{"apple", "banana", "cherry", "apple pie"},
			pattern: "a[a-z]+e",
			conf:    Config{Regexp: true},
			expected: map[int]string{
				0: "apple",
				3: "apple pie",
			},
		},
		{
			name:    "ignore case with regexp",
			lines:   []string{"Apple", "banana", "CHERRY", "apple pie"},
			pattern: "a[A-Z]+e",
			conf:    Config{IgnoreCase: true, Regexp: true},
			expected: map[int]string{
				0: "Apple",
				3: "apple pie",
			},
		},
		{
			name:    "lines after match",
			lines:   []string{"one", "two", "three", "four", "five"},
			pattern: "three",
			conf:    Config{After: 2},
			expected: map[int]string{
				2: "three",
				3: "four",
				4: "five",
			},
		},
		{
			name:    "lines before match",
			lines:   []string{"one", "two", "three", "four", "five"},
			pattern: "three",
			conf:    Config{Before: 2},
			expected: map[int]string{
				0: "one",
				1: "two",
				2: "three",
			},
		},
		{
			name:    "lines around match",
			lines:   []string{"one", "two", "three", "four", "five"},
			pattern: "three",
			conf:    Config{Context: 1},
			expected: map[int]string{
				1: "two",
				2: "three",
				3: "four",
			},
		},
		{
			name:    "combine invert and ignore case",
			lines:   []string{"Apple", "banana", "CHERRY", "apple pie"},
			pattern: "apple",
			conf:    Config{IgnoreCase: true, Invert: true},
			expected: map[int]string{
				1: "banana",
				2: "CHERRY",
			},
		},
		{
			name:     "empty input",
			lines:    []string{},
			pattern:  "apple",
			conf:     Config{},
			expected: map[int]string{},
		},
		{
			name:     "no matches",
			lines:    []string{"banana", "cherry", "date"},
			pattern:  "apple",
			conf:     Config{},
			expected: map[int]string{},
		},
		{
			name:    "special characters in regexp",
			lines:   []string{"apple", "banana.?", "cherry?", "apple pie"},
			pattern: `\.\?`,
			conf:    Config{Regexp: true},
			expected: map[int]string{
				1: "banana.?",
			},
		},
		{
			name:    "lines before and after match",
			lines:   []string{"one", "two", "three", "four", "five"},
			pattern: "three",
			conf:    Config{Before: 1, After: 1},
			expected: map[int]string{
				1: "two",
				2: "three",
				3: "four",
			},
		},
		{
			name:    "context with regexp",
			lines:   []string{"one", "two", "three", "four", "five"},
			pattern: "t.o",
			conf:    Config{Context: 1, Regexp: true},
			expected: map[int]string{
				0: "one",
				1: "two",
				2: "three",
			},
		},
		{
			name:    "multiple matches in a single line",
			lines:   []string{"apple banana cherry", "date fig grape"},
			pattern: "a",
			conf:    Config{Fixed: false},
			expected: map[int]string{
				0: "apple banana cherry",
				1: "date fig grape",
			},
		},
		{
			name:    "complex regexp pattern",
			lines:   []string{"apple1", "banana2", "cherry3", "apple pie"},
			pattern: `\w+\d`,
			conf:    Config{Regexp: true},
			expected: map[int]string{
				0: "apple1",
				1: "banana2",
				2: "cherry3",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := findMatches(tt.lines, tt.pattern, &tt.conf)
			if tt.conf.Count {
				count := len(got)
				if count != tt.expected {
					t.Errorf("findMatches() got %v, expected %v", count, tt.expected)
				}
			} else {
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("findMatches() got %v, expected %v", got, tt.expected)
				}
			}
		})
	}
}
