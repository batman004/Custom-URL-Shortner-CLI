package handlers

import (
	"flag"
	"fmt"
	"os"
	gocontroller "github.com/batman004/Custom-URL-Shortner-CLI/controllers"	
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
		Urls := gocontroller.GetUrls()
		
		fmt.Printf("Original \t Shortened \t CustomTag \t Description \n")
		for _, url := range Urls {
			fmt.Printf("%v \t %v \t %v \t %v \n",url.OriginalUrl, url.ShortenedUrl, url.CustomTag, url.Description)
		}

		return
	}

	if *customTag != "" {
		urls := gocontroller.GetUrls()
		customTag := *customTag
		for _, url := range urls {
			if customTag == url.CustomTag {
				fmt.Printf("Original \t Shortened \t CustomTag \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \n",url.OriginalUrl, url.ShortenedUrl, url.CustomTag, url.Description)
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

func HandleAdd(addCmd *flag.FlagSet,customTag *string, originalUrl *string, description *string){

	ValidateUrl(addCmd, customTag, originalUrl,description )


	shortenedUrl := CreateUrl(*originalUrl)


	url := gocontroller.Url {
		CustomTag: *customTag,
		Description: *description,
		OriginalUrl: *originalUrl, 
		ShortenedUrl: shortenedUrl,
	}

	urls := gocontroller.GetUrls()
	urls = append(urls,url)

	gocontroller.SaveUrl(urls)

}