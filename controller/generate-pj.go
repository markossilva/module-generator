package controller

import "fmt"

type GeneratePJ struct {
	ApplicationId string
	ModuleName    string
	PackageName   string
}

func (data *GeneratePJ) Generate() {
	fmt.Printf("PJ generate: %v\n", data.ModuleName)
}
