package Testing

import "testing"

//file ending with _test.go tells go test command that this file contains test funcitons
func TestAddNumber(t *testing.T) { //test function starts with Test
	got := AddNumber(3, 4) // takes *testing.T as the only parameter
	want := 7

	if got != want {
		t.Errorf("got %d, want%d", got, want) // t.Errof indicates that test failed
	}
}

// 'go test' to run
