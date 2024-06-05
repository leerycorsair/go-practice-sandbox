package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantDst string
		wantErr error
	}{
		{
			name:    "baseTest",
			args:    args{src: "a4bc2d5e"},
			wantDst: "aaaabccddddde",
			wantErr: nil,
		},
		{
			name:    "noUnpackingTest",
			args:    args{src: "abcd"},
			wantDst: "abcd",
			wantErr: nil,
		},
		{
			name:    "onlyDigitsStringTest",
			args:    args{src: "45"},
			wantDst: "",
			wantErr: errors.New("invalid str"),
		},
		{
			name:    "emptyStringTest",
			args:    args{src: ""},
			wantDst: "",
			wantErr: nil,
		},
		{
			name:    "withEscapeDigitsTest",
			args:    args{src: "qwe\\4\\5"},
			wantDst: "qwe45",
			wantErr: nil,
		},
		{
			name:    "withDigitsUnpackingTest",
			args:    args{src: "qwe\\45"},
			wantDst: "qwe44444",
			wantErr: nil,
		},
		{
			name:    "withEscapingEscapeTest",
			args:    args{src: "qwe\\\\5"},
			wantDst: "qwe\\\\\\\\\\",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := UnpackString(tt.args.src)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("UnpackString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDst != tt.wantDst {
				t.Errorf("UnpackString() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}
