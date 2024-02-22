package Testing

import "testing"

func TestAddNumber(t *testing.T) {
	got := AddNumber(3, 4)
	want := 7

	if got != want {
		t.Errorf("got %d, want%d", got, want)
	}
}

// 'go test' to run
