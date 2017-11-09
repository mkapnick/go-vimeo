package source

import (
	"io/ioutil"
	"net/http"
)

type Source struct {
	Url      string
	Response *http.Response
	IsValid  bool
}

// New a new httpok source
func New(source string) (*Source, error) {
	return Initialize(source)
}

// Get request wrapper
func Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

// Initialize checks to see if the server accepts byte ranges
func Initialize(source string) (*Source, error) {
	req, err := http.NewRequest("GET", source, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	value := res.Header.Get("Accept-Ranges")

	if value == "" || value != "bytes" {
		return &Source{
			Url:      source,
			Response: res,
			IsValid:  false,
		}, nil
	}

	return &Source{
		Url:      source,
		Response: res,
		IsValid:  true,
	}, nil
}
