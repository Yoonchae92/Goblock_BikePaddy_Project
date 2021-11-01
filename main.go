package main

import (
	"log"
	"net/http"
	"web22-1/app"

	"github.com/urfave/negroni"
)

func main() {

	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Print("Started App")
	err := http.ListenAndServe(":3001", n)
	if err != nil {
		panic(err)
	}
}
