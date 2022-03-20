package handlers

import (
	"io"
	"log"
	"net/http"
)

func CreateUrl(originalUrl string) string{

	baseUrl := "http://tinyurl.com/api-create.php?url="
	urlToShorten := originalUrl
	getReqUrl := baseUrl + urlToShorten

	req, err := http.NewRequest("GET", getReqUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "text/html")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	return (string(b))

}
