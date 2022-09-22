package mycommander

import (
	"net/http"
	"testing"
)

func TestHttpClient_IsOK(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Valid Url",
			fields: fields{
				Client: &http.Client{},
			},
			args: args{
				url: "https://github.com",
			},
			want: true,
		},
		{
			name: "Not Valid Url",
			fields: fields{
				Client: &http.Client{},
			},
			args: args{
				url: "http://notvalidurl.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := &HttpClient{
				Client: tt.fields.Client,
			}
			if got := hc.IsOK(tt.args.url); got != tt.want {
				t.Errorf("HttpClient.IsOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpClient_IsNotFound(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Not Found Url",
			fields: fields{
				Client: &http.Client{},
			},
			args: args{
				url: "https://github.com/cli/cli/test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := &HttpClient{
				Client: tt.fields.Client,
			}
			if got := hc.IsNotFound(tt.args.url); got != tt.want {
				t.Errorf("HttpClient.IsNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
