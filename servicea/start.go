package main

import (
	"github.com/ssgo/s"
	"github.com/ssgo/s/discover"
)

type In struct {
	Name string
}
type Out struct {
	FirstName string
	LastName  string
}

func main() {
	s.Restful(0, "POST", "/parseName", func(in In, caller *discover.Caller) Out {
		result := struct {
			FirstName string
		}{}
		caller.Post("svc-b", "/getLastName", in).To(&result)
		return Out{
			LastName:  in.Name,
			FirstName: result.FirstName,
		}
	})
	s.Start()
}
