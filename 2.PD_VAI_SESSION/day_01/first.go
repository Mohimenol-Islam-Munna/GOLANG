package main

import "fmt"

	const banglaHelloPrefix = "আসসালামুয়ালাইকুম "
const bangla = "Bangla"
const hello = "Hello "


func main() {
	var msg = message("M", "")

	fmt.Println(msg)
}

func message(name string, lang string) string {


	if name == "" {
        name = "world"
    }
    if lang == bangla {
        return banglaHelloPrefix + name
    }
    return hello + name
}


