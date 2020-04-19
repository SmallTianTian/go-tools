package slice

import (
	"reflect"
	"testing"
)

func TestUintPtrInSlice(t *testing.T) {
	type args struct {
		slice []uintptr
		key   uintptr
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []uintptr{1, 2, 3},
				key:   uintptr(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []uintptr{1, 2, 3},
				key:   uintptr(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintPtrInSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("UintPtrInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8InSlice(t *testing.T) {
	type args struct {
		slice []uint8
		key   uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []uint8{1, 2, 3},
				key:   uint8(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []uint8{1, 2, 3},
				key:   uint8(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Uint8InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16InSlice(t *testing.T) {
	type args struct {
		slice []uint16
		key   uint16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []uint16{1, 2, 3},
				key:   uint16(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []uint16{1, 2, 3},
				key:   uint16(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Uint16InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32InSlice(t *testing.T) {
	type args struct {
		slice []uint32
		key   uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []uint32{1, 2, 3},
				key:   uint32(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []uint32{1, 2, 3},
				key:   uint32(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Uint32InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintInSlice(t *testing.T) {
	type args struct {
		slice []uint
		key   uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []uint{1, 2, 3},
				key:   uint(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []uint{1, 2, 3},
				key:   uint(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintInSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("UintInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64InSlice(t *testing.T) {
	type args struct {
		slice []uint64
		key   uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []uint64{1, 2, 3},
				key:   uint64(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []uint64{1, 2, 3},
				key:   uint64(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Uint64InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8InSlice(t *testing.T) {
	type args struct {
		slice []int8
		key   int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []int8{1, 2, 3},
				key:   int8(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []int8{1, 2, 3},
				key:   int8(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Int8InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16InSlice(t *testing.T) {
	type args struct {
		slice []int16
		key   int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []int16{1, 2, 3},
				key:   int16(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []int16{1, 2, 3},
				key:   int16(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Int16InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32InSlice(t *testing.T) {
	type args struct {
		slice []int32
		key   int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []int32{1, 2, 3},
				key:   int32(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []int32{1, 2, 3},
				key:   int32(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Int32InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntInSlice(t *testing.T) {
	type args struct {
		slice []int
		key   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []int{1, 2, 3},
				key:   int(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []int{1, 2, 3},
				key:   int(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntInSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("IntInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64InSlice(t *testing.T) {
	type args struct {
		slice []int64
		key   int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []int64{1, 2, 3},
				key:   int64(1),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []int64{1, 2, 3},
				key:   int64(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Int64InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32InSlice(t *testing.T) {
	type args struct {
		slice []float32
		key   float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []float32{1.2, 3.4},
				key:   float32(1.2),
			},
			want: true,
		},
		{
			name: "Must be in with over limit.",
			args: args{
				slice: []float32{1.99999999, 3.4},
				key:   float32(2),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []float32{1.9999999, 3.4},
				key:   float32(2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Float32InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64InSlice(t *testing.T) {
	type args struct {
		slice []float64
		key   float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must be in.",
			args: args{
				slice: []float64{1.2, 3.4},
				key:   float64(1.2),
			},
			want: true,
		},
		{
			name: "Must be in with over limit.",
			args: args{
				slice: []float64{1.999_999_999_999_999_9, 3.4},
				key:   float64(2),
			},
			want: true,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []float64{1.999_999_999_999_999, 3.4},
				key:   float64(2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Float64InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSlice(t *testing.T) {
	type args struct {
		slice interface{}
		key   interface{}
	}
	one := &args{slice: []uint{1, 2, 3}, key: 1}
	two := &args{slice: []uint{1, 2, 3}, key: 1}
	thr := &args{slice: []uint{4, 5, 6}, key: 2}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Must in.",
			args: args{
				slice: []*args{one, thr},
				key:   one,
			},
			want: true,
		},
		{
			name: "Not same object.",
			args: args{
				slice: []*args{one, thr},
				key:   two,
			},
			want: false,
		},
		{
			name: "Must not in.",
			args: args{
				slice: []*args{one, two},
				key:   thr,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctUintptrSlice(t *testing.T) {
	type args struct {
		origin []uintptr
	}
	tests := []struct {
		name string
		args args
		want []uintptr
	}{
		{
			name: "Can remove.",
			args: args{[]uintptr{1, 2, 3, 4, 1, 2}},
			want: []uintptr{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]uintptr{1, 2, 3, 4}},
			want: []uintptr{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctUintptrSlice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctUintptrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctUint8Slice(t *testing.T) {
	type args struct {
		origin []uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "Can remove.",
			args: args{[]uint8{1, 2, 3, 4, 1, 2}},
			want: []uint8{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]uint8{1, 2, 3, 4}},
			want: []uint8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctUint8Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctUint16Slice(t *testing.T) {
	type args struct {
		origin []uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "Can remove.",
			args: args{[]uint16{1, 2, 3, 4, 1, 2}},
			want: []uint16{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]uint16{1, 2, 3, 4}},
			want: []uint16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctUint16Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctUint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctUint32Slice(t *testing.T) {
	type args struct {
		origin []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "Can remove.",
			args: args{[]uint32{1, 2, 3, 4, 1, 2}},
			want: []uint32{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]uint32{1, 2, 3, 4}},
			want: []uint32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctUint32Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctUintSlice(t *testing.T) {
	type args struct {
		origin []uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "Can remove.",
			args: args{[]uint{1, 2, 3, 4, 1, 2}},
			want: []uint{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]uint{1, 2, 3, 4}},
			want: []uint{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctUintSlice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctUintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctUint64Slice(t *testing.T) {
	type args struct {
		origin []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "Can remove.",
			args: args{[]uint64{1, 2, 3, 4, 1, 2}},
			want: []uint64{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]uint64{1, 2, 3, 4}},
			want: []uint64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctUint64Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt8Slice(t *testing.T) {
	type args struct {
		origin []int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "Can remove.",
			args: args{[]int8{1, 2, 3, 4, 1, 2}},
			want: []int8{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]int8{1, 2, 3, 4}},
			want: []int8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt8Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt16Slice(t *testing.T) {
	type args struct {
		origin []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "Can remove.",
			args: args{[]int16{1, 2, 3, 4, 1, 2}},
			want: []int16{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]int16{1, 2, 3, 4}},
			want: []int16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt16Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt32Slice(t *testing.T) {
	type args struct {
		origin []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Can remove.",
			args: args{[]int32{1, 2, 3, 4, 1, 2}},
			want: []int32{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]int32{1, 2, 3, 4}},
			want: []int32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt32Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctIntSlice(t *testing.T) {
	type args struct {
		origin []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Can remove.",
			args: args{[]int{1, 2, 3, 4, 1, 2}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctIntSlice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctInt64Slice(t *testing.T) {
	type args struct {
		origin []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Can remove.",
			args: args{[]int64{1, 2, 3, 4, 1, 2}},
			want: []int64{1, 2, 3, 4},
		},
		{
			name: "No remove.",
			args: args{[]int64{1, 2, 3, 4}},
			want: []int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctInt64Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctFloat32Slice(t *testing.T) {
	type args struct {
		origin []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "Can remove.",
			args: args{[]float32{1.9999999, 1.9999999, 2, 3}},
			want: []float32{1.9999999, 2, 3},
		},
		{
			name: "No remove.",
			args: args{[]float32{1.9999999, 2, 3}},
			want: []float32{1.9999999, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctFloat32Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctFloat32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctFloat64Slice(t *testing.T) {
	type args struct {
		origin []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Can remove.",
			args: args{[]float64{1.999_999_999_999_999, 1.999_999_999_999_999, 2, 3}},
			want: []float64{1.999_999_999_999_999, 2, 3},
		},
		{
			name: "No remove.",
			args: args{[]float64{1.999_999_999_999_999, 2, 3}},
			want: []float64{1.999_999_999_999_999, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctFloat64Slice(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctFloat64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctSlice(t *testing.T) {
	type args struct {
		origin interface{}
	}
	one := &args{[]uint{1, 2, 3}}
	two := &args{[]uint{1, 2, 3}}
	thr := &args{[]uint{4, 5, 6}}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Can remove.",
			args: args{origin: []*args{one, one, thr}},
			want: []*args{one, thr},
		},
		{
			name: "Not remove.",
			args: args{origin: []*args{one, thr}},
			want: []*args{one, thr},
		},
		{
			name: "Not remove when not same object.",
			args: args{origin: []*args{one, two}},
			want: []*args{one, two},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistinctSlice(tt.args.origin)
			expect := tt.want.([]*args)
			exist := got.([]*args)
			if len(expect) != len(exist) {
				t.Errorf("DistinctSlice() = %v, want %v", got, tt.want)
			}
			for i, item := range expect {
				if item != expect[i] {
					t.Errorf("DistinctSlice() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
