package main

import (
	"reflect"
	"testing"
)

func TestCutLines(t *testing.T) {
	type args struct {
		lines []string
		conf  *Config
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "baseTest",
			args: args{
				lines: []string{"ssd"},
				conf: &Config{
					Fields:    []int{0},
					Delimeter: "\t",
					Separated: false,
				},
			},
			want: []string{"ssd"},
		},
		{
			name: "emptyInput",
			args: args{
				lines: nil,
				conf: &Config{
					Fields:    []int{1},
					Delimeter: "\t",
					Separated: false,
				},
			},
			want: nil,
		},
		{
			name: "singleField",
			args: args{
				lines: []string{"a\tb\tc"},
				conf: &Config{
					Fields:    []int{1},
					Delimeter: "\t",
					Separated: false,
				},
			},
			want: []string{"b"},
		},
		{
			name: "multipleFields",
			args: args{
				lines: []string{"a\tb\tc", "d\te\tf"},
				conf: &Config{
					Fields:    []int{0, 2},
					Delimeter: "\t",
					Separated: false,
				},
			},
			want: []string{"a\tc", "d\tf"},
		},
		{
			name: "differentDelimiter",
			args: args{
				lines: []string{"a,b,c", "d,e,f"},
				conf: &Config{
					Fields:    []int{1},
					Delimeter: ",",
					Separated: false,
				},
			},
			want: []string{"b", "e"},
		},
		{
			name: "separatedFieldTrue",
			args: args{
				lines: []string{"a b c", "d e f", "ghi"},
				conf: &Config{
					Fields:    []int{0, 1},
					Delimeter: " ",
					Separated: true,
				},
			},
			want: []string{"a b", "d e"},
		},
		{
			name: "separatedFieldFalse",
			args: args{
				lines: []string{"a b c", "d e f", "ghi"},
				conf: &Config{
					Fields:    []int{0, 1},
					Delimeter: " ",
					Separated: false,
				},
			},
			want: []string{"a b", "d e", "ghi"},
		},
		{
			name: "noDelimiterInSeparatedMode",
			args: args{
				lines: []string{"a b c", "def"},
				conf: &Config{
					Fields:    []int{0, 1},
					Delimeter: " ",
					Separated: true,
				},
			},
			want: []string{"a b"},
		},
		{
			name: "indexOutOfRange",
			args: args{
				lines: []string{"a\tb"},
				conf: &Config{
					Fields:    []int{2},
					Delimeter: "\t",
					Separated: false,
				},
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutLines(tt.args.lines, tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cutLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
