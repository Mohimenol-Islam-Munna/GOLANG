package iteration 

import "testing"

func TestRepeat(t *testing.T){
	want := "aaaaa"
	got := Repeat("a")

	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i<b.N; i++{
		Repeat("a")
	}
}