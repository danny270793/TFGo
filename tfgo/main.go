package tfgo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/danny270793/tfgo/logger"
)

type Ussage struct {
	Line int
	File string
}

type Module struct {
	Name              string
	FullPath          string
	VariablesDeclared map[string][]Ussage
	VariablesUssed    map[string][]Ussage
}

func GetVariablesDetailByModule(modulesBasePath string) []Module {
	variablesDeclarationRegexp, err := regexp.Compile("variable \"[a-zA-Z0-9_]*\"")
	if err != nil {
		panic(err)
	}
	variableUsssageRegexp, err := regexp.Compile("var\\.[a-zA-Z0-9_]*")
	if err != nil {
		panic(err)
	}

	modules, err := ioutil.ReadDir(modulesBasePath)
	if err != nil {
		panic(err)
	}

	modulesFound := make([]Module, 0)
	for _, module := range modules {
		moduleFullPath := path.Join(modulesBasePath, module.Name())
		moduleFound := Module{VariablesDeclared: map[string][]Ussage{}, VariablesUssed: map[string][]Ussage{}, Name: module.Name(), FullPath: moduleFullPath}

		files, err := ioutil.ReadDir(moduleFullPath)
		if err != nil {
			panic(err)
		}
		logger.Trace(fmt.Sprintf("opening module \"%s\" from path \"%s\"\n", module.Name(), moduleFullPath))
		for _, file := range files {
			fileFullPath := path.Join(moduleFullPath, file.Name())
			logger.Trace(fmt.Sprintf("\topening file \"%s\" from path \"%s\"\n", file.Name(), fileFullPath))
			lines, err := os.ReadFile(fileFullPath)
			if err != nil {
				panic(err)
			}

			for index, line := range strings.Split(string(lines), "\n") {
				for _, match := range variableUsssageRegexp.FindAllString(line, -1) {
					tokens := strings.Split(match, "var.")
					variable := tokens[1]
					ussage := Ussage{Line: index + 1, File: fileFullPath}
					moduleFound.VariablesUssed[variable] = append(moduleFound.VariablesUssed[variable], ussage)
				}

				for _, match := range variablesDeclarationRegexp.FindAllString(line, -1) {
					tokens := strings.Split(match, " ")
					variable := tokens[1][1 : len(tokens[1])-1]
					ussage := Ussage{Line: index + 1, File: fileFullPath}
					moduleFound.VariablesDeclared[variable] = append(moduleFound.VariablesDeclared[variable], ussage)
				}
			}
		}
		modulesFound = append(modulesFound, moduleFound)
	}

	return modulesFound
}
