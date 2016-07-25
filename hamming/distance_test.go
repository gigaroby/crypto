package hamming

import "testing"

func TestEditDistance(t *testing.T) {
	d1 := []byte("this is a test")
	d2 := []byte("wokka wokka!!!")
	expected := 37
	got, err := EditDistance(d1, d2)
	if err != nil {
		t.Error("error computing edit distance: ", err)
	}

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}

}
