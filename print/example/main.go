package main

import (
	"fmt"

	"github.com/aca/x/print"
)

func main(){
    x := struct{
        X int
        Y string
    }{
        X: 3,
        Y: "wer",
    }

    print.JSON(x)
    fmt.Println("\n------\n")
    print.YAML(x)
}
