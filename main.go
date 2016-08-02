package main

import (
	"github.com/MarvinMenzerath/UpAndRunning2-Terminal-Client/lib"
	"os"
	"strconv"
)

const VERSION = "2.2.0"

func main() {
	lib.PrintHead()

	lib.ReadConfigurationFromFile("config/local.json")

	if len(os.Args) <= 1 {
		lib.PrintWrongArguments()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "help" || os.Args[1] == "-h" {
			lib.PrintHelp()
		} else if os.Args[1] == "version" || os.Args[1] == "-v" {
			lib.PrintVersion(VERSION)
		} else if os.Args[1] == "websites" || os.Args[1] == "-w" {
			lib.WebsitesCommand()
		} else {
			lib.PrintWrongArguments()
		}
	} else if len(os.Args) == 3 {
		if os.Args[1] == "status" || os.Args[1] == "-s" {
			lib.StatusCommand(os.Args[2])
		} else if os.Args[1] == "results" || os.Args[1] == "-r" {
			lib.ResultsCommand(os.Args[2], 50, 0)
		} else {
			lib.PrintWrongArguments()
		}
	} else if len(os.Args) == 4 {
		if os.Args[1] == "results" || os.Args[1] == "-r" {
			limit, err := strconv.Atoi(os.Args[3])
			if err != nil {
				lib.PrintWrongArguments()
			}
			lib.ResultsCommand(os.Args[2], limit, 0)
		} else {
			lib.PrintWrongArguments()
		}
	} else if len(os.Args) == 5 {
		if os.Args[1] == "results" || os.Args[1] == "-r" {
			limit, err := strconv.Atoi(os.Args[3])
			if err != nil {
				lib.PrintWrongArguments()
			}
			offset, err := strconv.Atoi(os.Args[4])
			if err != nil {
				lib.PrintWrongArguments()
			}
			lib.ResultsCommand(os.Args[2], limit, offset)
		} else {
			lib.PrintWrongArguments()
		}
	} else if len(os.Args) > 5 {
		lib.PrintWrongArguments()
	}
}
