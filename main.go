package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
)

type Climber struct {
	Name                string `csv:"NAME"`
	AfterSemis          int    `csv:"After Semis"`
	AfterQualifications int    `csv:"After Qualifications"`
	Progressed          bool   `csv:"Progressed?"`
}

func main() {
	log.Print("Hello")
	f, err := os.OpenFile("fixtures/2021-lead-women-briancon.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("Error opening file")
		return
	}
	defer f.Close()

	climbers := []*Climber{}

	if err := gocsv.UnmarshalFile(f, &climbers); err != nil { // Load clients from file
		log.Error().Err(err).Msg("Error unmarshalling file")
	}

	for _, c := range climbers {
		fmt.Printf("Climber: %s", c.Name)
	}
}
