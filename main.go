package main

import (
	"log"

	"github.com/firo-18/pnc/info"
)

func main() {
	// meta := info.NewMetaData()
	// meta.WriteYAML()

	// checkDollSkills("Abigail")
}

func checkDollSkills(name string) {
	d := info.NewDoll()
	d.ReadYAML(name)
	log.Println(d.Skills.Passive.Desc)
	log.Println(d.Skills.Auto.Desc)
	log.Println(d.Skills.Ultimate.Desc)
}
