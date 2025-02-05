package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 5)
	want := 7

	if got != want {
		t.Errorf("Test case 01 fail : \n got %d, but want %d \n", got, want)
	} else {
		fmt.Printf("Test case 01 pass : \n got %d, but want %d \n", got, want)
	}

}
