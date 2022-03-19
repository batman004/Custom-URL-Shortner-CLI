package handlers

import (
	"flag"
	"fmt"
	"os"
) 

func HandleGet(getCmd *flag.FlagSet, all *bool, customTag *string){

	getCmd.Parse(os.Args[2:])

	if *all == false && *customTag == "" {
    fmt.Print("custom-tag is required or specify --all for all urls")
    getCmd.PrintDefaults()
    os.Exit(1)
  }

	if *all {
		//return all videos
		urls := getUrls()
		
		fmt.Printf("Original \t Shortened \t CustomTag \t Description \n")
		for _, url := range urls {
			fmt.Printf("%v \t %v \t %v \t %v \n",url.original, url.shorten, url.customTag, url.description)
		}

		return
	}

	if *customTag != "" {
		urls := getUrls()
		customTag := *customTag
		for _, url := range urls {
			if customTag == url.Id {
				fmt.Printf("Original \t Shortened \t CustomTag \t Description \n")
				fmt.Printf(" %v \t %v \t %v \t %v \n",url.original, url.shorten, url.customTag, url.description)
			}
		}
	} 



}

func ValidateUrl(addCmd *flag.FlagSet,customTag *string, originalUrl *string, description *string ){

	addCmd.Parse(os.Args[2:])

	if *customTag == "" || *originalUrl == "" || *description == "" {
		fmt.Print("all fields are required for bookmarking a url")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet,id *string, customTag *string, originalUrl *string, description *string, shortenedUrl *string ){

	ValidateUrl(addCmd, customTag, originalUrl,description )



	url := Url {
		CustomTag: *customTag,
		Description: *description,
		OriginalUrl: *originalUrl, 
		ShortenedUrl: *shortenedUrl,
	}

	urls := getUrls()
	urls = append(urls,url)

	saveUrls(urls)

}