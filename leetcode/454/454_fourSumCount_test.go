package _54

import "testing"

func Test_maxMin(t *testing.T) {
	type args struct {
		a     int
		b     int
		isMax bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a:     0,
				b:     1,
				isMax: false,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				a:     0,
				b:     1,
				isMax: true,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMin(tt.args.a, tt.args.b, tt.args.isMax); got != tt.want {
				t.Errorf("maxMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
