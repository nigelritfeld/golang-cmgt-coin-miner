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

func TestHash(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test example hash",
			args: args{"000078454c038871fa4d67b0022a30baaf25eaa231f8991b108e2624f052f3f8CMGT Mining CorporationBob PIKAB11548689513858154874778871610312"},
			want: "00005d430ce77ad654b5309a770350bfb4cf49171c682330a2eccc98fd8853cf",
		},
		{
			name: "Test second hash",
			args: args{"0000a41723a9797cd18a8c8f462fbd04ae1bcc010c814704161e570c5a6791acCMGT Mining CorporationNigel 1004416116768966634841676908077111116100"},
			want: "00000e04de1e13322b500c5ac442a2f6f99454aeac164bc52af0be17774e3c96",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mod10.HashPayload(tt.args.phrase)
			if got != tt.want {
				t.Errorf("TryNonce() got = %v, want %v", got, tt.want)
			}
		})
	}
}
