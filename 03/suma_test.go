package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data []int
		answer int
	}

	tests := []test {
		test{
			data: []int{21, 21},
			answer: 42,
		},
		test{
			data: []int{3, 4, 5},
			answer: 12,
		},
		test{
			data: []int{1, 1},
			answer: 2,
		},
		test{
			data: []int{-1, 0, 1},
			answer: 0,
		},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		
		if x != v.answer {
			t.Error("Expected: ", v.answer,  " Got: ", x)
		}
	}
}