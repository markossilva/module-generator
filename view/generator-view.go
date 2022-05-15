package view

import (
	"module-generator/controller"
	"module-generator/enums"
)

type IGeneratorView interface {
	InitView()
}

type generatorView struct {
	selectedProject string
	moduleName      string
	packageName     string
	applicationName string
}

func (data *generatorView) initProjectGenerator() {
	var projectType controller.GenerateProject

	switch data.selectedProject {
	case enums.IDS:
		projectType = &controller.GenerateIds{
			ModuleName:    data.moduleName,
			PackageName:   data.packageName,
			ApplicationId: data.applicationName,
		}
		break
	case enums.PF:
		projectType = &controller.GeneratePF{
			ModuleName:    data.moduleName,
			PackageName:   data.packageName,
			ApplicationId: data.applicationName,
		}
		break
	case enums.PJ:
		projectType = &controller.GeneratePJ{
			ModuleName:    data.moduleName,
			PackageName:   data.packageName,
			ApplicationId: data.applicationName,
		}
		break
	}

	controller.GenerateProject.Generate(projectType)
}
