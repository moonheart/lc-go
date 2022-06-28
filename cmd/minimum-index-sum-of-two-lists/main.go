package main

import (
	"fmt"
	"math"
	"reflect"
)

func findRestaurant(list1 []string, list2 []string) (res []string) {
	m1 := map[string]int{}
	for i, s := range list1 {
		m1[s] = i
	}
	min := math.MaxInt

	for i, s := range list2 {
		idx1, ok := m1[s]
		if !ok {
			continue
		}
		if idx1+i < min {
			min = idx1 + i
			res = []string{s}
		} else if idx1+i == min {
			res = append(res, s)
		}
	}
	return
}

func main() {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"1", args{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}}, []string{"Shogun"}},
		{"2", args{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}}, []string{"Shogun"}},
	}
	for _, tt := range tests {
		if gotRes := findRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(gotRes, tt.wantRes) {
			fmt.Errorf("findRestaurant() = %v, want %v", gotRes, tt.wantRes)
		}
	}
}
