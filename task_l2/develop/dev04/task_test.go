package main

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "baseTest",
			args: args{strs: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "лол", "kek"}},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name: "singleAnagramSubsetTest",
			args: args{strs: []string{"пятак", "пятка", "тяпка"}},
			want: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name: "notSingleAnagramSubsetTest",
			args: args{strs: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name: "emptyInputTest",
			args: args{strs: []string{}},
			want: map[string][]string{},
		},
		{
			name: "noAnagramsTest",
			args: args{strs: []string{"пятак", "пяток", "тряпочка", "листок", "стол"}},
			want: map[string][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagram(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nresult = %v\nwant   = %v", got, tt.want)
			}
		})
	}
}
