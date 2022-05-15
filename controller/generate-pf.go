package controller

import "fmt"

type GeneratePF struct {
	ApplicationId string
	ModuleName    string
	PackageName   string
}

func (data *GeneratePF) Generate() {
	fmt.Printf("PF generate: %v\n", data.ModuleName)
}
