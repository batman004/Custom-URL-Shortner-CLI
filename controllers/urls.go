package controllers

import(
	"io/ioutil"
	"encoding/json"

)



func GetUrls()(urls []Url){
	
	fileBytes, err := ioutil.ReadFile("data/saved_urls.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &urls)

	if err != nil {
		panic(err)
	}

	return urls
}

func SaveUrl(urls []Url)(){

	videoBytes, err  := json.Marshal(urls)
	if err != nil {
		  panic(err)
	  }
  
	  err = ioutil.WriteFile("data/saved_urls.json", videoBytes, 0644)
	if err != nil {
		  panic(err)
	  }
  
  }
