package controller

import (
	"log"
	"os"
	"text/template"
)

type GenerateIds struct {
	ApplicationId string
	ModuleName    string
	PackageName   string
}

func (data *GenerateIds) Generate() {
	data.generateParsedFiles(
		"template/template-01.txt",
		"module/",
		"template-02.txt",
	)
}

func (data *GenerateIds) generateParsedFiles(src, targetFolder, fileName string) {
	tParsed, err := template.ParseFiles(src)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(targetFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create(targetFolder + fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := tParsed.Execute(newFile, data); err != nil {
		log.Fatal(err)
	}

	if err := newFile.Close(); err != nil {
		log.Fatal(err)
	}
}
