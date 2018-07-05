package main

import (
	"github.com/ssgo/s"
)

type In struct {
	Name string
}
type Out struct {
	FirstName string
}

func main() {

	firstNames := map[string]string{
		"Tom": "Lee",
		"Sam": "Wang",
		"Default": "Wang",
	}

	s.Restful(1, "POST", "/getLastName", func(in In) Out {
		firstName := firstNames[in.Name]
		if firstName == "" {
			firstName = firstNames["Default"]
		}

		return Out{
			FirstName: firstName,
		}
	})
	s.Start()
}
