package main

import (
	"log"

	"github.com/aca/x/downloader"
)

func main(){
    err := downloader.Download("https://www.gstatic.com/webp/gallery3/1.png", "")

    if err != nil {
      log.Fatal(err)
    }
}
