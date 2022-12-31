package main

import (
	"flag"
	"fmt"
	"os"

	"github.io/danny270793/tfgo/tfgo"
)

const VERSION = "1.0.0"

func main() {
	modulesBasePath := flag.String("path", "", "path where to seach for modules")
	showVersion := flag.Bool("version", false, "show version")
	verbose := flag.String("verbose", "DEBUG", "level of the logs to show")
	flag.Parse()

	os.Setenv("VERBOSE_LEVEL", *verbose)

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if *modulesBasePath == "" {
		fmt.Println("missing -path flag")
		os.Exit(0)
	}

	if len(os.Args) <= 2 {
		fmt.Printf("Usage of tfgo:\n")
		fmt.Printf("  variable\n")
		fmt.Printf("        missing\n")
		fmt.Printf("        ussage\n")
		os.Exit(1)
	}

	command := os.Args[1]
	action := os.Args[2]
	if command == "variables" {
		modules := tfgo.GetVariablesDetailByModule(*modulesBasePath)
		fmt.Printf("\"%d\" modules found on path \"%s\"\n\n", len(modules), *modulesBasePath)

		if action == "missing" {
			for _, module := range modules {
				for key, ussages := range module.VariablesDeclared {
					if len(module.VariablesUssed[key]) == 0 {
						fmt.Printf("variable \"%s\" declared on \"%s:%d\" was not used\n", key, ussages[0].File, ussages[0].Line)
					}
				}
			}
		} else if action == "ussage" {
			for _, module := range modules {
				fmt.Printf("module \"%s\" from path \"%s\"\n", module.Name, module.FullPath)
				for key, ussages := range module.VariablesDeclared {
					fmt.Printf("\tvariable \"%s\" used \"%d\" times", key, len(module.VariablesUssed[key]))
					for _, ussage := range ussages {
						fmt.Printf(" was declared on \"%s:%d\"\n", ussage.File, ussage.Line)
					}
					for _, ussage := range module.VariablesUssed[key] {
						fmt.Printf("\t\tused on \"%s:%d\"\n", ussage.File, ussage.Line)
					}
				}
			}
		} else {
			fmt.Printf("invalid action %s\n", action)
			os.Exit(0)
		}
	} else {
		fmt.Printf("invalid command %s\n", command)
		os.Exit(0)
	}
}
