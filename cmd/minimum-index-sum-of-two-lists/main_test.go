package main

import (
	"reflect"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findRestaurant() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
