package view

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"module-generator/enums"
	"strconv"
)

type PromptuiView struct {
	selectedProject string
	applicationName string
	moduleName      string
	packageName     string
}

func (pv *PromptuiView) InitView() {
	generator := generatorView{}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: pv.validate,
	}

	result, err := prompt.Run()

	if pv.errorHandler(err) {
		return
	}

	generator.moduleName = result

	promptSelect := promptui.Select{
		Label: "Selecione o tipo de modulo para ser gerado:",
		Items: []string{enums.IDS, enums.PF, enums.PJ},
	}

	_, selected, selectErr := promptSelect.Run()

	if pv.errorHandler(selectErr) {
		return
	}

	generator.selectedProject = selected

	generator.initProjectGenerator()
}

func (pv *PromptuiView) validate(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}

func (pv *PromptuiView) errorHandler(err error) bool {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return true
	}
	return false
}
