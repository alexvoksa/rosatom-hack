package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexvoksa/rosatom-hack/go/internal/closer"
	"github.com/alexvoksa/rosatom-hack/go/internal/processor"
)

func main() {
	closerLocal := closer.New(os.Kill, os.Interrupt)

	hub, err := processor.NewHub(
		"free",
		"free",
		"postgres://hackathon:hackathon@localhost:5432/hackathon",
	)
	closerLocal.Add(hub.Close)
	if err != nil {
		log.Fatalln(err)
	}

	err = hub.Process("fcs_regions/Adygeja_Resp/contracts", time.Time{}, time.Time{})
	fmt.Println(err)
}
