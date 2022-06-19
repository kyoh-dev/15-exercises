package p1

import (
	"reflect"
	"testing"
)

func TestStrToIntSlice(t *testing.T) {
	slc := []string{"1", "4", "28", "67", "419"}
	result := []int{1, 4, 28, 67, 419}
	nslc := StrToIntSlice(&slc)
	if !reflect.DeepEqual(nslc, result) {
		t.Errorf("expected: %v; got: %v", result, nslc)
	}
}
