package print

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func JSON(v any) error {
    b, err := json.MarshalIndent(v, "", "\t")
    if err != nil{ 
        return err
    }

    fmt.Print(string(b))
    return nil
}

func YAML(v any) error {
    b, err := yaml.Marshal(v)
    if err != nil{ 
        return err
    }

    fmt.Print(string(b))
    return nil
}
