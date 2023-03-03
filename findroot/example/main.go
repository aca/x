package main

import (
	"log"

	"github.com/aca/x/findroot"
)

func main() {
	x, err := findroot.FindRoot(".git")
	if err != nil {
		panic(err)
	}
	log.Println(x)
}
