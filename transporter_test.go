package mycommander

import "testing"

func TestTransporter_valid(t *testing.T) {
	type fields struct {
		transports map[string]struct{}
	}
	type args struct {
		scheme string
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
				transports: map[string]struct{}{
					"ssh":     {},
					"git":     {},
					"git+ssh": {},
					"http":    {},
					"https":   {},
				},
			},
			args: args{
				scheme: "https",
			},
			want: true,
		},
		{
			name: "Fail",
			fields: fields{
				transports: map[string]struct{}{
					"ssh":     {},
					"git":     {},
					"git+ssh": {},
					"http":    {},
					"https":   {},
				},
			},
			args: args{
				scheme: "ftp",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transporter{
				transports: tt.fields.transports,
			}
			if got := tr.valid(tt.args.scheme); got != tt.want {
				t.Errorf("Transporter.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
