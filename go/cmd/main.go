package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alexvoksa/rosatom-hack/go/internal/processor"
)

func main() {
	hub, err := processor.NewHub(
		"free",
		"free",
		"postgres://hackathon:hackathon@localhost:5432/postgres",
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = hub.Process("fcs_regions/Adygeja_Resp/contracts", time.Time{}, time.Time{})
	fmt.Println(err)
}
