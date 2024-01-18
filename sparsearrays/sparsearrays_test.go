package main

import (
	"fmt"
	"testing"
)

func TestMatchingStrings(t *testing.T) {

	var tests = []struct {
		strList   []string
		queryList []string
		expected  []int32
	}{
		{[]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}, []int32{2, 1, 0}},
		{[]string{"def", "de", "fgh"}, []string{"de", "lmn", "fgh"}, []int32{1, 0, 1}},
		{[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
			[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}, []int32{1, 3, 4, 3, 2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("String list: %v, Query list: %v, Expected: %v", test.strList, test.queryList, test.expected), func(t *testing.T) {
			res := matchingStrings(test.strList, test.queryList)

			if len(res) != len(test.expected) {
				t.Errorf("Expected a length of %v, got %v", len(test.expected), len(res))
			} else {
				for k := range test.expected {
					if res[k] != test.expected[k] {
						t.Errorf("Expected %v, got %v", test.expected, res)
					}
				}
			}
		})
	}

}
