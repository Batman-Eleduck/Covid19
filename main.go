package main

import (
	"log"
	"os"

	"batman.covid19/diagram"
)

func main() {
	os.RemoveAll("./go-diagrams/")
	d := diagram.GetDiagram()
	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
