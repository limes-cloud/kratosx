package main

import "testing"

func Test_case2Camel(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "snake1",
			args: args{name: "SYSTEM_ERROR"},
			want: "SystemError",
		},
		{
			name: "snake2",
			args: args{name: "System_Error"},
			want: "SystemError",
		},
		{
			name: "snake3",
			args: args{name: "system_error"},
			want: "SystemError",
		},
		{
			name: "snake4",
			args: args{name: "System_error"},
			want: "SystemError",
		},
		{
			name: "upper1",
			args: args{name: "UNKNOWN"},
			want: "Unknown",
		},
		{
			name: "camel1",
			args: args{name: "SystemError"},
			want: "SystemError",
		},
		{
			name: "camel2",
			args: args{name: "systemError"},
			want: "SystemError",
		},
		{
			name: "lower1",
			args: args{name: "system"},
			want: "System",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := case2Camel(tt.args.name); got != tt.want {
				t.Errorf("case2Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
