package controllers

import(
	"io/ioutil"
	"encoding/json"
)

func getUrls()(urls []Url){
	
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &urls)

	if err != nil {
		panic(err)
	}

	return urls
}

func saveUrl(urls []Url)(){

	videoBytes, err  := json.Marshal(urls)
	if err != nil {
		  panic(err)
	  }
  
	  err = ioutil.WriteFile("./urls.json", videoBytes, 0644)
	if err != nil {
		  panic(err)
	  }
  
  }
