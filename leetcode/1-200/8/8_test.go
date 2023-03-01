/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/3/1 下午10:29.
 * Author:  guojia(https://github.com/guojia99)
 */

package main

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "42",
			args: args{
				"42",
			},
			want: 42,
		},
		{
			name: "+-12",
			args: args{
				"+-12",
			},
			want: 0,
		},
		{
			name: "-42",
			args: args{" -42"},
			want: -42,
		},
		{
			name: "--42",
			args: args{" -42"},
			want: -42,
		},
		{
			name: "9223372036854775808",
			args: args{
				"9223372036854775808",
			},
			want: 2147483647,
		},
		{
			name: "+5+",
			args: args{
				"+5+",
			},
			want: 5,
		},
		{
			name: "  +  413",
			args: args{
				"  +  413",
			},
			want: 0,
		},
		{
			name: "   +0 123",
			args: args{
				"   +0 123",
			},
			want: 0,
		},
		{
			name: "    -88827   5655  U",
			args: args{
				ss: "    -88827   5655  U",
			},
			want: -88827,
		},
		{
			name: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459",
			args: args{
				ss: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459",
			},
			want: 1<<31 - 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.ss); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
