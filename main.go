package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/firo-18/pnc/info"
)

func main() {
	QualityControl()
	meta := info.NewMetaData()
	meta.WriteYAML()
}

func QualityControl() {
	path := "data/dolls/wip/"
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("read-dir:", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			doll := info.DollProfile{}
			doll.ReadYAML(file.Name(), path)
			if doll.Verify() {
				err := os.Rename(path+file.Name(), "data/dolls/"+file.Name())
				if err != nil {
					log.Fatalln("rename:", err)
				}
				err = os.Rename("asset/dolls/icons/wip/"+doll.Name+".png", "asset/dolls/icons/"+doll.Name+".png")
				if err != nil {
					log.Fatalln("rename:", err)
				}
				log.Printf("Successfully move file and assets of Doll `%v' to production.", doll.Name)
			}
		}
	}
}

func RestructureDolls() {
	path := info.Path.DollData
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("read-dir:", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			doll := info.DollProfile{}
			doll.ReadYAML(file.Name(), path)

			// Make changes
			doll.Bio.Release = time.Date(2022, time.Now().Month(), 21, 11, 0, 0, 0, time.UTC)

			doll.WriteYAML("data/")
		}
	}
}
