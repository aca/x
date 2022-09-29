package yamldb_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aca/go/yamldb"
)

func ExampleYamlDB() {
	f, _ := ioutil.TempFile("", "yamldb")
	defer os.RemoveAll(f.Name())

	type DB struct {
		X string `yaml:"x"`
		A string `yaml:"a"`
	}

	f.WriteString(`x: y
a: b`)
	f.Close()

	db, err := yamldb.Open[DB](f.Name())
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
	// x: z
	// a: b
}
