package main

import (
	"log"
	"os"
	"strings"

	"github.com/firo-18/pnc/info"
)

func main() {
	QualityControl()
	meta := info.NewMetaData()
	meta.WriteYAML()
}

func QualityControl() {
	path := "data/dolls/upcoming/"
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
				err = os.Rename("asset/dolls/icons/upcoming/"+doll.Name+".png", "asset/dolls/icons/"+doll.Name+".png")
				if err != nil {
					log.Fatalln("rename:", err)
				}
				log.Printf("Successfully move file and assets of Doll `%v' to production.", doll.Name)
			}
		}
	}
}
