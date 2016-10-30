package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func Get(path string, query *url.Values, schema interface{}) error {
	URL := getURL()
	u, err := url.Parse(URL + path)
	if err != nil {
		return err
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, schema)
	fmt.Println(schema)
	if err != nil {
		return err
	}
	return nil
}

func getURL() string {
	return os.Getenv("BCDB")
}
