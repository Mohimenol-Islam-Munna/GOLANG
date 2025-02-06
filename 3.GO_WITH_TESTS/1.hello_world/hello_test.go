package main

import ("testing")

func TestHello(t *testing.T){
	output := hello("Munna")

	want := "Munna"

	if output != want {
        t.Errorf("Expected '%s', but got '%s'", want, output)
    }
}