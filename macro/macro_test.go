package macro_test

import (
	"fmt"

	. "github.com/aca/x/macro"
)

func ExampleMacro() {
	fmt.Println(S("hello %v.", "john"))
}
