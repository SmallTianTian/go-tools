package slice

import (
	"reflect"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	type args struct {
		slice []string
		key   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must in",
			args: args{
				slice: []string{"a", "b", "c"},
				key:   "a",
			},
			want: true,
		},
		{
			name: "Not in",
			args: args{
				slice: []string{"a", "b", "c"},
				key:   "",
			},
			want: false,
		},
		{
			name: "Blank not equal",
			args: args{
				slice: []string{"\n", "\t", "\n\t"},
				key:   "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctStringSlice(t *testing.T) {
	type args struct {
		origin []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Can remove",
			args: args{
				origin: []string{"a", "b", "c", "a"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "No remove",
			args: args{
				origin: []string{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Blank not remove",
			args: args{
				origin: []string{"\n", "\t", "\n\t"},
			},
			want: []string{"\n", "\t", "\n\t"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctStringSlice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
