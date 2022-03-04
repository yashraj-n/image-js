package lib

import (
	"fmt"
)

func Help(commandName string) {
	switch commandName {
	case "compile":
		fmt.Println("Usage: compile [Javascript file] [Image file]")
		fmt.Println("")
		fmt.Println("Encodes a Javascript file into a image file.")
		fmt.Println("Example: ijs compile test.js test.png --debug")
		fmt.Println("")
		break
	case "run":
		fmt.Println("Usage: run [Image File]")
		fmt.Println("")
		fmt.Println("Decodes a image file into a Javascript file and runs it. (Must have node.js installed)")
		fmt.Println("Example: ijs run test.png")
		fmt.Println("")
		break
	case "export":
		fmt.Println("Usage: export [Image File]")
		fmt.Println("")
		fmt.Println("Decodes a image file into a Javascript file")
		fmt.Println("Example: ijs export test.png")
		fmt.Println("")
		break
	default:
		fmt.Println("")
		fmt.Println("	IJS (Image Javascript) version 1.0.0")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("")
		fmt.Println(" ijs [command] [...arguments]")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("")
		fmt.Println("	compile		Encodes Javascript file to Image")
		fmt.Println("	run		Executes the Image file as an node.js file")
		fmt.Println("	export		Exports the Image file to a .js file")
		fmt.Println("	help		Displays this help message. Use 'help [command]' for more information about a command.")
		fmt.Println("")
	}
}
