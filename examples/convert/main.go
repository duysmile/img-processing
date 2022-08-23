package main

import (
	"github.com/duysmile/img-processing/convert"
	"io/ioutil"
	"log"
)

func main() {
	inputFile := "img/camel.heic"
	output, err := convert.ConvertFromFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	outputFile := "img/camel.jpg"
	err = ioutil.WriteFile(outputFile, output, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("convert image successfully")
}
