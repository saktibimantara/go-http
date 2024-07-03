package go_http

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoHTTP_Get(t *testing.T) {
	type args struct {
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				url: "https://api.zippopotam.us/us/33162",
			},
			want: &Response{
				Code: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := New(&Config{})
			got, err := g.Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Code != tt.want.Code {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}

			fmt.Println(string(got.Data))
		})
	}
}

func TestNewBearerToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBearerToken(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBearerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
