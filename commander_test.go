package mycommander

import (
	"reflect"
	"testing"
)

func TestCommander_valid(t *testing.T) {
	type fields struct {
		currentPath string
		timeout     int64
		debugMode   bool
		commands    map[string]*Command
	}
	type args struct {
		cmd string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Valid Command",
			fields: fields{
				currentPath: "./",
				timeout:     0,
				debugMode:   false,
				commands: map[string]*Command{
					"mkdir.p":   NewCommand("mkdir.p", "mkdir -p [PATH] [PATH_2]", []string{"PATH", "PATH_2"}...),
					"grep.echo": NewCommand("grep.echo", "echo \"this is grep text\" | grep '^[FILTERS]'", []string{"FILTERS"}...),
				},
			},
			args: args{
				cmd: "mkdir.p",
			},
			want: true,
		},
		{
			name: "Unknown Command",
			fields: fields{
				currentPath: "./",
				timeout:     0,
				debugMode:   false,
				commands: map[string]*Command{
					"mkdir.p":   NewCommand("mkdir.p", "mkdir -p [PATH] [PATH_2]", []string{"PATH", "PATH_2"}...),
					"grep.echo": NewCommand("grep.echo", "echo \"this is grep text\" | grep '^[FILTERS]'", []string{"FILTERS"}...),
				},
			},
			args: args{
				cmd: "grep.list",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Commander{
				currentPath: tt.fields.currentPath,
				timeout:     tt.fields.timeout,
				debugMode:   tt.fields.debugMode,
				commands:    tt.fields.commands,
			}
			if got := c.valid(tt.args.cmd); got != tt.want {
				t.Errorf("Commander.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommander_run(t *testing.T) {
	type fields struct {
		currentPath string
		timeout     int64
		debugMode   bool
		commands    map[string]*Command
	}
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success Command Run",
			fields: fields{
				currentPath: "./",
				timeout:     0,
				debugMode:   false,
				commands: map[string]*Command{
					"grep.echo": NewCommand("grep.echo", "echo \"this is grep text\" | grep '^[FILTERS]'", []string{"FILTERS"}...),
				},
			},
			args: args{
				name: "grep.echo",
				args: []string{"this"},
			},
			want:    []byte("this is grep text\n"),
			wantErr: false,
		},
		{
			name: "Fail Unknown Command Run",
			fields: fields{
				currentPath: "./",
				timeout:     0,
				debugMode:   false,
				commands: map[string]*Command{
					"grep.echo": NewCommand("grep.echo", "echo \"this is grep text\" | mgrep '^[FILTERS]'", []string{"FILTERS"}...),
				},
			},
			args: args{
				name: "grep.echo",
				args: []string{"is"},
			},
			want:    []byte("\n"),
			wantErr: true,
		},
		{
			name: "Fail Unknown Command Run #2",
			fields: fields{
				currentPath: "./",
				timeout:     0,
				debugMode:   false,
				commands: map[string]*Command{
					"grep.echo": NewCommand("grep.echo", "echo \"this is grep text\" | grep '^[FILTERS]'", []string{"FILTERS"}...),
				},
			},
			args: args{
				name: "grep.echo.multi",
				args: []string{"is"},
			},
			want:    []byte("\n"),
			wantErr: true,
		},
		{
			name: "Not Enough Argument Command Run",
			fields: fields{
				currentPath: "./",
				timeout:     0,
				debugMode:   false,
				commands: map[string]*Command{
					"grep.echo": NewCommand("grep.echo", "echo \"this is grep text\" | grep '^[FILTERS]'", []string{"FILTERS"}...),
				},
			},
			args: args{
				name: "grep.echo",
				args: nil,
			},
			want:    []byte("\n"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Commander{
				currentPath: tt.fields.currentPath,
				timeout:     tt.fields.timeout,
				debugMode:   tt.fields.debugMode,
				commands:    tt.fields.commands,
			}
			got, err := c.run(tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commander.run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Commander.DeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
