package main

import (
	"testing"
)

func TestChop_normalCase(t *testing.T) {
	if chop(1, []int32{1, 3, 5}) != 0 {
		t.Fatal("faild!")
	}
	if chop(3, []int32{1, 3, 5}) != 1 {
		t.Fatal("faild!")
	}
	if chop(5, []int32{1, 3, 5}) != 2 {
		t.Fatal("faild!")
	}
}
func TestChop_abnormalCase(t *testing.T) {
	if chop(0, []int32{1, 3, 5}) != -1 {
		t.Fatal("faild!")
	}
	if chop(0, []int32{}) != -1 {
		t.Fatal("faild!")
	}
	if chop(1, []int32{1, 1, 3, 5}) != 0 {
		t.Fatal("faild!")
	}

}
