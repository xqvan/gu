package gu

import (
	"reflect"
	"testing"
)

func TestCommonPostJson(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		body    map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommonPostJson(tt.args.url, tt.args.headers, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommonPostJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonPostJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
