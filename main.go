
package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	fmt.Println("Custom URL generator")

	//'url get' subcommand
  getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `URL get` command
  getAll := getCmd.Bool("all", false, "Get all custom bookmarked URL")
  getCustom := getCmd.String("custom", "", "custom name given for a url")

	//'URL add' subcommand
  addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for `videos add` command
  addUrl := addCmd.String("url", "", "Create new shortened URL")
  addCustomUrl := addCmd.String("custom-url", "", "Create new custom URL")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
  }

	//look at the 2nd argument's value
	switch os.Args[1] {
		case "get": // if its the 'get' command
			HandleGet(getCmd, getAll, getCustom)
		case "add": // if its the 'add' command
			HandleAdd(addCmd,addUrl, addCustomUrl)
		default: // if we don't understand the input
	}


}