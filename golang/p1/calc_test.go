package p1

import (
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{4, 7, 28, 4917}
	result := 4956
	s := sum(nums)
	if s != result {
		t.Errorf("expected: %d; got: %d", result, s)
	}
}

func TestProduct(t *testing.T) {
	nums := []int{597, 392, 38, 50183}
	result := 446273002896
	p := product(nums)
	if p != result {
		t.Errorf("expected: %d; got: %d", result, p)
	}
}

func TestMean(t *testing.T) {
	nums := []int{59, 3492, 57, 38, 10, 348}
	result := 667
	m := mean(nums)
	if m != result {
		t.Errorf("expected: %d; got: %d", result, m)
	}
}

func TestSqrt(t *testing.T) {

}
