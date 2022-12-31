package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/danny270793/tfgo/tfgo"
)

const VERSION = "1.0.0"

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("Usage of tfgo:\n")
		fmt.Printf("  variables\n")
		fmt.Printf("    missing\n")
		fmt.Printf("      search for variables declared but not ussed\n")
		fmt.Printf("    ussage\n")
		fmt.Printf("      search where are used each variabel declared\n")
		os.Exit(1)
	}

	command := os.Args[1]
	subcommand := os.Args[2]
	switch command {
	case "variables":
		flagSet := flag.NewFlagSet("missing", flag.ExitOnError)
		modulesBasePath := flagSet.String("path", "", "path where to seach for modules")
		verbose := flagSet.String("verbose", "DEBUG", "level of the logs to show")
		flagSet.Parse(os.Args[3:])

		os.Setenv("VERBOSE_LEVEL", *verbose)

		if *modulesBasePath == "" {
			fmt.Println("missing -path flag")
			os.Exit(0)
		}

		modules := tfgo.GetVariablesDetailByModule(*modulesBasePath)
		fmt.Printf("\"%d\" modules found on path \"%s\"\n\n", len(modules), *modulesBasePath)
		switch subcommand {
		case "missing":
			for _, module := range modules {
				for key, ussages := range module.VariablesDeclared {
					if len(module.VariablesUssed[key]) == 0 {
						fmt.Printf("variable \"%s\" declared on \"%s:%d\" was not used\n", key, ussages[0].File, ussages[0].Line)
					}
				}
			}
		case "ussage":
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
		}
	case "version":
		fmt.Println(VERSION)
		os.Exit(0)
	default:
		fmt.Printf("invalid command: %s\n", command)
		os.Exit(1)
	}
}
