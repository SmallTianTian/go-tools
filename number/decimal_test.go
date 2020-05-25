package number

import (
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		num   float64
		digit uint8
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "be int by floor.",
			args: args{
				num:   123.456,
				digit: 0,
			},
			want: 123,
		},
		{
			name: "be int by ceiling.",
			args: args{
				num:   123.5,
				digit: 0,
			},
			want: 124,
		},
		{
			name: "keep three decimal by floor.",
			args: args{
				num:   123.456432,
				digit: 3,
			},
			want: 123.456,
		},
		{
			name: "be int by ceiling by ceiling.",
			args: args{
				num:   123.456789,
				digit: 3,
			},
			want: 123.457,
		},
		{
			name: "with carry",
			args: args{
				num:   123.9996,
				digit: 3,
			},
			want: 124.0,
		},
		{
			name: "max 16 decimal",
			args: args{
				num:   float64(1.999_999_999_999_999_163),
				digit: 100,
			},
			want: float64(1.999_999_999_999_999_1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.num, tt.args.digit); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
