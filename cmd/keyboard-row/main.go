package main

import (
	"fmt"
	"reflect"
	"strings"
)

func findWords(words []string) (res []string) {
	m := map[int32]map[int32]bool{}
	m1 := map[int32]bool{}
	for _, i := range "qwertyuiop" {
		m1[i] = true
		m[i] = m1
	}
	m2 := map[int32]bool{}
	for _, i := range "asdfghjkl" {
		m2[i] = true
		m[i] = m2
	}
	m3 := map[int32]bool{}
	for _, i := range "zxcvbnm" {
		m3[i] = true
		m[i] = m3
	}

next:
	for _, word := range words {
		lw := strings.ToLower(word)
		mp := m[int32(lw[0])]
		for _, i := range lw {
			if !mp[i] {
				continue next
			}
		}
		res = append(res, word)
	}
	return
}

func findWords2(words []string) (res []string) {
	rowIdx := []int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}
next:
	for _, word := range words {
		lw := strings.ToLower(word)
		idx := rowIdx[lw[0]-'a']
		for _, i := range lw {
			if rowIdx[i-'a'] != idx {
				continue next
			}
		}
		res = append(res, word)
	}
	return
}

func main() {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"1", args{[]string{"Hello", "Alaska", "Dad", "Peace"}}, []string{"Alaska", "Dad"}},
		{"2", args{[]string{"omk"}}, []string{}},
		{"3", args{[]string{"adsdf", "sfd"}}, []string{"adsdf", "sfd"}},
	}
	for _, tt := range tests {
		if gotRes := findWords(tt.args.words); !reflect.DeepEqual(gotRes, tt.wantRes) {
			fmt.Errorf("findWords() = %v, want %v", gotRes, tt.wantRes)
		}
	}
}
