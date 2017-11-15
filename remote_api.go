package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Client used to for remote-service-api
type Client struct {
	HTTPClient *http.Client
	ApiURL     string
}

func (client *Client) createAPIURL(u string) string {
	if client.ApiURL == "" {
		return "defaultapiurlgoeshere"
	}

	return client.ApiURL
}

func createClient(httpClient *http.Client, apiURL string) *Client {
	client := new(Client)
	client.HTTPClient = httpClient
	client.ApiURL = apiURL

	return client
}

func (client *Client) getURL(URL string) (string, error) {

	log.Printf("Making remote request: %s", URL)
	URL = client.createAPIURL(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", errors.New("Could not create request for " + URL + " - " + err.Error())
	}

	// Make a request to the sourceURL
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", errors.New("Could not get " + URL + " - " + err.Error())
	}
	defer res.Body.Close()

	// Read the whole body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Problem reading the body for " + URL + " - " + err.Error())
	}

	return string(body[:]), nil
}

type PriceLookup struct {
	ID    string
	Price int
}

func (client *Client) LookupPrice(prodID string) (int, error) {
	body, err := client.getURL("price/" + prodID)

	if err != nil {
		return -1, errors.New("Problem getting URL for image info ID " + prodID + " - " + err.Error())
	}

	dec := json.NewDecoder(strings.NewReader(body))
	var lookup PriceLookup

	if err := dec.Decode(&lookup); err != nil {
		return -1, errors.New("Problem decoding json for product " + prodID + " - " + err.Error())
	}

	return lookup.Price, nil
}
