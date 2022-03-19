package handlers

import (

	"io/ioutil"
	"log"
	"net/http"

)

func createUrl(originalUrl string)[]byte{

	baseUrl := "http://tinyurl.com/api-create.php?url="
	urlToShorten := originalUrl
	getReqUrl := baseUrl + urlToShorten

	request, err := http.NewRequest(
		http.MethodGet, //method
		getReqUrl,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not process a custom URL %v", err)
	}


	request.Header.Add("Accept", "text/html")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes

}
