package gu

import (
	"errors"
	"testing"
)

func TestFatalError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				err: errors.New("aaa"),
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FatalError(tt.args.err)
		})
	}
}
