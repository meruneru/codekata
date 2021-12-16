package main

import (
	"testing"
)

func TestChop_Case1(t *testing.T) {
	if chop(3, []int{}) != -1 {
		t.Fatal("faild!")
	}
	if chop(3, []int{1}) != -1 {
		t.Fatal("faild!")
	}
	if chop(1, []int{1}) != 0 {
		t.Fatal("faild!")
	}
}
func TestChop_Case2(t *testing.T) {
	if chop(1, []int{1, 3, 5}) != 0 {
		t.Fatal("faild!")
	}
	if chop(3, []int{1, 3, 5}) != 1 {
		t.Fatal("faild!")
	}
	if chop(5, []int{1, 3, 5}) != 2 {
		t.Fatal("faild!")
	}
	if chop(0, []int{1, 3, 5}) != -1 {
		t.Fatal("faild!")
	}
	if chop(2, []int{1, 3, 5}) != -1 {
		t.Fatal("faild!")
	}
	if chop(4, []int{1, 3, 5}) != -1 {
		t.Fatal("faild!")
	}
	if chop(6, []int{1, 3, 5}) != -1 {
		t.Fatal("faild!")
	}
}
func TestChop_abnormalCase(t *testing.T) {
	if chop(1, []int{1, 3, 5, 7}) != 0 {
		t.Fatal("faild!")
	}
	if chop(3, []int{1, 3, 5, 7}) != 1 {
		t.Fatal("faild!")
	}
	if chop(5, []int{1, 3, 5, 7}) != 2 {
		t.Fatal("faild!")
	}
	if chop(7, []int{1, 3, 5, 7}) != 3 {
		t.Fatal("faild!")
	}
	if chop(0, []int{1, 3, 5, 7}) != -1 {
		t.Fatal("faild!")
	}
	if chop(2, []int{1, 3, 5, 7}) != -1 {
		t.Fatal("faild!")
	}
	if chop(4, []int{1, 3, 5, 7}) != -1 {
		t.Fatal("faild!")
	}
	if chop(6, []int{1, 3, 5, 7}) != -1 {
		t.Fatal("faild!")
	}
	if chop(8, []int{1, 3, 5, 7}) != -1 {
		t.Fatal("faild!")
	}
}
