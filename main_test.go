package main

import (
	"golang-cmgt-coin-miner/mod10"
	"reflect"
	"testing"
)

// TestSumSequence Tests mod10 sum sequence
func TestSumSequence(t *testing.T) {
	type args struct {
		chunks [][]int
	}
	chunks := [][]int{
		{1, 1, 6, 1, 0, 1, 1, 2, 0, 1},
		{1, 6, 0, 1, 2, 3, 4, 5, 6, 7},
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test sum sequence mod10",
			args: args{chunks},
			want: []int{2, 7, 6, 2, 2, 4, 5, 7, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mod10.SumSequence(tt.args.chunks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestHashSHA256 Test hashing algorithm
func TestHashSHA256(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test sha256 hashing",
			args: args{"2762245768"},
			want: "d0b3cb0cc9100ef243a1023b2a129d15c28489e387d3f8b687a7299afb4b5079",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mod10.HashSHA256(tt.args.input); got != tt.want {
				t.Errorf("HashSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}
