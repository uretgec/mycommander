package mycommander

import "testing"

func TestHoster_valid(t *testing.T) {
	type fields struct {
		hosts map[string]struct{}
	}
	type args struct {
		host string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Success",
			fields: fields{
				hosts: map[string]struct{}{
					"github.com":    {},
					"gitlab.com":    {},
					"bitbucket.org": {},
				},
			},
			args: args{
				host: "github.com",
			},
			want: true,
		},
		{
			name: "Fail",
			fields: fields{
				hosts: map[string]struct{}{
					"github.com":    {},
					"gitlab.com":    {},
					"bitbucket.org": {},
				},
			},
			args: args{
				host: "api.github.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hoster{
				hosts: tt.fields.hosts,
			}
			if got := h.valid(tt.args.host); got != tt.want {
				t.Errorf("Hoster.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
