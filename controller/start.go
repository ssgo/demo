package main

import (
	"github.com/ssgo/discover"
	"github.com/ssgo/s"
)

func main() {
	s.Register(0, "/hello", func(caller *discover.Caller) string {
		result := struct {
			FirstName string
			LastName  string
		}{}
		caller.Post("svc-a", "/parseName", map[string]string{"Name": "Sam"}).To(&result)
		return result.LastName + " " + result.FirstName
	})
	s.Start()
}
