package jsondb_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aca/go/jsondb"
)

func ExampleJsonDB() {
	f, _ := ioutil.TempFile("", "jsondb")
	defer os.RemoveAll(f.Name())

	type DB struct {
		X string `json:"x"`
		A string `json:"a"`
	}

	f.WriteString(`{ "x": "y", "a": "b" }`)
	f.Close()

	db, err := jsondb.Open[DB](f.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Data.X)

	db.Data.X = "z"
	err = db.Save()
	if err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadFile(f.Name())
	fmt.Printf("%v", string(b))

	// Output: y
	// {"x":"z","a":"b"}
}
