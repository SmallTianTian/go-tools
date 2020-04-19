package slice

import "testing"

func TestRemoveSlice(t *testing.T) {
	type args struct {
		slice interface{}
		i     int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "remove first int",
			args: args{
				slice: &[]interface{}{1, 2, 3, 4},
				i:     0,
			},
			want: []interface{}{2, 3, 4},
		},
		{
			name: "remove last int",
			args: args{
				slice: &[]interface{}{1, 2, 3, 4},
				i:     3,
			},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "remove middle int",
			args: args{
				slice: &[]interface{}{1, 2, 3, 4},
				i:     2,
			},
			want: []interface{}{1, 2, 4},
		},
		{
			name: "remove max then len",
			args: args{
				slice: &[]interface{}{1, 2, 3, 4},
				i:     10,
			},
			want: []interface{}{1, 2, 3, 4},
		},
		{
			name: "remove minus",
			args: args{
				slice: &[]interface{}{1, 2, 3, 4},
				i:     -1,
			},
			want: []interface{}{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveSlice(tt.args.slice, tt.args.i)
			old := tt.args.slice.(*[]interface{})
			e := tt.want.([]interface{})
			for i, item := range *old {
				if item != e[i] {
					t.Errorf("[%s] not equal. Except: %v, Found: %v.", tt.name, tt.want, tt.args.slice)
					break
				}
			}
		})
	}
}
