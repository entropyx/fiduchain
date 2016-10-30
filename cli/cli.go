package cli

import (
	"bytes"
	"encoding/json"
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
	if err != nil {
		return err
	}
	return nil
}

func Post(path string, reqSchema, schema interface{}) error {
	URL := getURL()
	u, err := url.Parse(URL + path)
	if err != nil {
		return err
	}
	reqBody, err := json.Marshal(reqSchema)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, schema)
	if err != nil {
		return err
	}
	return nil
}

func getURL() string {
	return os.Getenv("BCDB")
}
