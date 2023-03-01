/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/3/1 下午8:57.
 * Author:  guojia(https://github.com/guojia99)
 */

package _7

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "2",
			args: args{
				x: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
