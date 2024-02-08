package un_test

import "testing"

func Test_totalPrice(t *testing.T) {
	// test #1
	expected := uint(0)
	actual := totalPrice(0, 150, 12)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
	// test #2
	expected = uint(112)
	actual = totalPrice(0, 100, 12)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
	// test #3
	expected = uint(224)
	actual = totalPrice(2, 100, 12)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}
