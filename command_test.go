package mycommander

import "testing"

func TestCommand_generate(t *testing.T) {
	type fields struct {
		name      string
		template  string
		delimeter Delimeter
		args      []string
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Simple Command Template",
			fields: fields{
				name:     "mkdir.p",
				template: "mkdir -p [PATH] [PATH_2]",
				delimeter: Delimeter{
					start: "[",
					end:   "]",
				},
				args: []string{
					"PATH",
					"PATH_2",
				},
			},
			args: args{
				args: []string{
					"/var/log/nginx",
					"/var/log/statsd",
				},
			},
			want: "mkdir -p /var/log/nginx /var/log/statsd",
		},
		{
			name: "Different Delimeter Command Template",
			fields: fields{
				name:     "mkdir.p",
				template: "mkdir -p {PATH} {PATH_2}",
				delimeter: Delimeter{
					start: "{",
					end:   "}",
				},
				args: []string{
					"PATH",
					"PATH_2",
				},
			},
			args: args{
				args: []string{
					"/var/log/nginx",
					"/var/log/statsd",
				},
			},
			want: "mkdir -p /var/log/nginx /var/log/statsd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				name:      tt.fields.name,
				template:  tt.fields.template,
				delimeter: tt.fields.delimeter,
				args:      tt.fields.args,
			}
			if got := c.generate(tt.args.args...); got != tt.want {
				t.Errorf("Command.generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
