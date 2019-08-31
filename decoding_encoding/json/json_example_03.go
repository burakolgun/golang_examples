package json

import (
	"encoding/json"
	"fmt"
)

func ExampleThree()  {
	// Create a map of key/value pairs.

	c := make(map[string] interface{})
	c["name"] = "gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// Marshal the map into a JSON string.

	err, data := json.MarshalIndent(c, "s", "    ")

	if err != nil {
		fmt.Println(string(err))
	}

	fmt.Println(data)
}