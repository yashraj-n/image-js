package main

import (
	"github.com/yashraj-n/ijs/lib"
	"github.com/yashraj-n/ijs/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		utils.Error("Invalid command usage, Use ijs --help for more info")
		return
	}
	debug := false
	if len(args) > 2 && args[2] == "--debug" {
		debug = true
		utils.Debug("--debug detected. Debug mode enabled.")
	}

	if len(args) > 3 && args[3] == "--debug" {
		debug = true
		utils.Debug("--debug detected. Debug mode enabled.")
	}

	switch args[0] {
	case "--help":
		if len(args) > 1 {
			lib.Help(args[1])
		} else {
			lib.Help("")
		}

		break
	case "compile":
		if len(args) == 2 {
			utils.Error("Invalid compile command usage, Use ijs --help for more info")
			return
		}
		lib.Compile(args[1], args[2], debug)
		break
	case "run":
		if !utils.NodeExists() {
			utils.Error("Please install Node.js")
		}
		lib.Execute(args[1], debug)
		break
	case "export":
		lib.Decode(args[1], debug)
		break
	default:
		utils.Error("Invalid command, Use ijs --help for more info")
	}

}
