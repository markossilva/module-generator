package controller

import "fmt"

type GenerateIds struct {
	ApplicationId string
	ModuleName    string
	PackageName   string
}

func (data *GenerateIds) Generate() {
	fmt.Printf("IDS generate: %v\n", data.ModuleName)
}
