package main

import (
	"sort"
	"strings"
	"testing"
)

func TestSortInts(t *testing.T) {
	got := []int{4, 5, 1, 2, 7, 276, 34}
	want := []int{1, 2, 4, 5, 7, 34, 276}
	sort.Ints(got)
	for k := range want {

		if want[k] != got[k] {
			t.Fatalf("хотели это %v, а получили это %v", want, got)
		}
	}

}

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "test 1",
			got:  []string{"b", "a", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "test 2",
			got:  []string{"2", "1", "3"},
			want: []string{"1", "2", "3"},
		},
		{
			name: "test 3",
			got:  []string{"в", "а", "д"},
			want: []string{"а", "в", "д"},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := append([]string(nil), v.got...)
			sort.Strings(got)
			for k := range v.want {
				if v.want[k] != got[k] {
					t.Fatalf("хотели это %v, а получили это %v", v.want, got)
				}
			}
		})
	}
}
func TestStringsToLower(t *testing.T) {
	got := "aBcAA"
	want := "abcaa"
	got = strings.ToLower(got)
	if got != want {
		t.Fatalf("хотели:%s, а получили:%s", got, want)
	}
}
func TestStringsTrimSpace(t *testing.T) {
	test := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "test 1",
			got:  " faskolfea ",
			want: "faskolfea",
		},
		{
			name: "test 2",
			got:  "faskolfea",
			want: "faskolfea",
		},
		{
			name: "test 3",
			got:  "",
			want: "",
		},
		{
			name: "test 4",
			got:  "     ",
			want: "",
		},
		{
			name: "test 5",
			got:  "faw\n\t",
			want: "faw",
		},
	}
	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := strings.TrimSpace(v.got)
			if got != v.want {
				t.Fatalf("бро мы хотели это:%s, а получили это:%s", v.want, got)
			}
		})
	}
}
