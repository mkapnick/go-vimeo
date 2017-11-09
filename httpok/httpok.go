package httpok

import (
	"io/ioutil"
	"net/http"
)

// Get request wrapper
func Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

// IsValidSource checks to see if the server accepts byte ranges
func IsValidSource(source string) (bool, error) {
	req, err := http.NewRequest("GET", source, nil)

	if err == nil {
		return false, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()

	value := res.Header.Get("Accept-Ranges")

	if value == "" || value != "bytes" {
		return false, nil
	}

	return true, nil
}
