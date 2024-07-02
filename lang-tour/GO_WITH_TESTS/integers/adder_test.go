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

//func TestAnotherAdder(t *testing.T) {
//	got := Add(2, 501)
//	want := 503
//
//	if got != want {
//		t.Errorf("Test case 02 : \n got %d, but want %d \n", got, want)
//	} else {
//		fmt.Printf("Test case 02 pass : \n got %d, but want %d \n", got, want)
//	}
//}
