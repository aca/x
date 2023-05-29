package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aca/x/filedb/csvdb"
)

func main() {
	f, _ := ioutil.TempFile("", "csvdb")
	defer os.RemoveAll(f.Name())

	type DB struct {
		X string `csv:"x"`
		A string `csv:"a"`
	}

    // type DBS []DB
	f.WriteString(`a,x
aaa,xxx
aaa2,xxx2`)
	f.Close()

	db, err := csvdb.Open[DB](f.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((*db.Data)[0].X)
	fmt.Println((*db.Data)[0].A)

	// ((*db.Data)[0]).X = "z"
	// err = db.Save()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// b, _ := ioutil.ReadFile(f.Name())
	// fmt.Printf("%v", string(b))

	// Output: y
	// {"x":"z","a":"b"}
}
