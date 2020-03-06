package object

import "testing"

func TestStringIsBlank(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{content: ""},
			want: true,
		},
		{
			name: "space",
			args: args{content: "  "},
			want: true,
		},
		{
			name: "windows newline",
			args: args{content: "\r\n"},
			want: true,
		},
		{
			name: "linux newline",
			args: args{content: "\n"},
			want: true,
		},
		{
			name: "mac newline",
			args: args{content: "\r"},
			want: true,
		},
		{
			name: "space + windows newline",
			args: args{content: "  \r\n  "},
			want: true,
		},
		{
			name: "tab",
			args: args{content: "\t"},
			want: true,
		},
		{
			name: "tab + windows newline",
			args: args{content: "\t\r\n"},
			want: true,
		},
		{
			name: "number",
			args: args{content: "123"},
			want: false,
		},
		{
			name: "space + number",
			args: args{content: " 123 "},
			want: false,
		},
		{
			name: "newline + number",
			args: args{content: "\r\n123"},
			want: false,
		},
		{
			name: "tab + number",
			args: args{content: "\t123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringIsBlank(tt.args.content); got != tt.want {
				t.Errorf("StringIsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustUint64(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want uint64
		pic  bool
	}{
		{
			name: "number",
			args: args{"123"},
			want: 123,
		},
		{
			name: "alphabet",
			args: args{"abc"},
			pic:  true,
		},
		{
			name: "decimals",
			args: args{"1.1"},
			pic:  true,
		},
		{
			name: "minus",
			args: args{"-1"},
			pic:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				e := recover()
				if tt.pic != (e != nil) {
					t.Errorf("MustUint64() panic %v, want %v", e != nil, tt.pic)
				}
			}()
			got := MustUint64(tt.args.value)
			if got != tt.want {
				t.Errorf("MustUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustInt64(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int64
		pic  bool
	}{
		{
			name: "number",
			args: args{"123"},
			want: 123,
		},
		{
			name: "minus",
			args: args{"-1"},
			want: -1,
		},
		{
			name: "alphabet",
			args: args{"abc"},
			pic:  true,
		},
		{
			name: "decimals",
			args: args{"1.1"},
			pic:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				e := recover()
				if tt.pic != (e != nil) {
					t.Errorf("MustUint64() panic %v, want %v", e != nil, tt.pic)
				}
			}()
			got := MustInt64(tt.args.value)
			if got != tt.want {
				t.Errorf("MustUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
