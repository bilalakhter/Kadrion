package main

import (
	"flag"
	"fmt"
	"github.com/bilalakhter/kadrion/internal/customtypes"
	"github.com/bilalakhter/kadrion/internal/testprocess"
	"github.com/bilalakhter/kadrion/internal/toolhelp"
	"os"
)

var tconfigs customtypes.Tconfig

const Version string = "0.10.0"

func main() {

	helpFlag := flag.Bool("help", false, "Tool Usage Information")
	versionFlag := flag.Bool("version", false, "Provide current version rolling")
	args := os.Args[1:]

	flag.Parse()

	if *helpFlag == true {
		toolhelp.ToolInfo()
		os.Exit(0)
	}
	if *versionFlag == true {
		fmt.Printf(" kadrion %s\n", Version)
	}
	if len(args) == 2 && args[0] == "apply" && args[1] == "tconfig.yaml" {
		_, err := os.Stat("tconfig.yaml")
		if os.IsNotExist(err) {
			fmt.Println("tconfig.yaml file not found")
			os.Exit(1)
		} else {
			readfile, err := os.ReadFile("tconfig.yaml")
			if err != nil {
				fmt.Println("Error reading tconfig.yaml:", err)
				os.Exit(1)

			} else {
				testprocess.ProcessYaml(readfile)
			}
		}
	} else {
		toolhelp.ToolInfo()
	}

}
