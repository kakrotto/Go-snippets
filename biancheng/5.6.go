package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "")

func main() {
	flag.Parse()

	var skill = map[string]func(){
		"python": func() {
			println("python")
		},
		"java": func() {
			println("java")
		},
		"go": func() {
			println("go")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		println("no skill")
		fmt.Println()
	}
}
