package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/aca/x/log"
	"github.com/aca/x/noti"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
)

func main() {
	log.SetLevel(zerolog.InfoLevel)
	var msg string
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		msg = string(b)
	} else {
		msg = strings.Join(os.Args[1:], " ")
	}

    err := noti.Send(msg)
    if err != nil {
        log.Fatal(err)
    }
}
