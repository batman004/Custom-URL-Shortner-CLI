package main

import (
	"fmt"
	"os"
	"flag"
	handler "github.com/batman004/Custom-URL-Shortner-CLI/handlers"
)

func main() {
	fmt.Println("Custom URL generator")

	//'url get' subcommand
  getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `URL get` command
  getAll := getCmd.Bool("all", false, "Get all custom bookmarked URLs")
  getCustom := getCmd.String("custom", "", "custom name given for a URL")

	//'URL add' subcommand
  addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for `videos add` command
  addUrl := addCmd.String("url", "", "Add the URL you want to shorten")
  addCustomTag := addCmd.String("custom-tag", "", "Create new custom tag for URL")
  addDescription := addCmd.String("description", "", "Add Description of bookmarked URL")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
  }

	//look at the 2nd argument's value
	switch os.Args[1] {
		case "get": // if its the 'get' command
		handler.HandleGet(getCmd, getAll, getCustom)
		case "add": // if its the 'add' command
		handler.HandleAdd(addCmd, addUrl,addCustomTag, addDescription)
		default: // if we don't understand the input
	}


}