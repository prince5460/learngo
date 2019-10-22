package testing

import "testing"

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("error,%v", res)
	}
	t.Logf("success")
}
