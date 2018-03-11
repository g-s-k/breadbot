package main

import "fmt"
import "time"
import "net/url"
import "net/http"
import "encoding/json"
import "io/ioutil"

//import "image"

const keyFile = "keys.json"
const searchInterval = 5

func main() {
	// get api info
	apiKeys := getKeys()
	// pipelines
	c1 := make(chan string)
	// separate goroutines for different functions
	go initiate(c1)
	go imgGet(c1, apiKeys["google-search"])
	// don't close out program until interrupted
	for {
	}
}

func getKeys() map[string]map[string]string {
	// read file into bytearray
	raw, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}
	// decode as json
	var keys map[string]map[string]string
	json.Unmarshal(raw, &keys)
	return keys
}

func initiate(ch chan<- string) {
	for {
		// get user input
		var input string
		fmt.Scanln(&input)
		// send it into the channel
		ch <- url.QueryEscape(input)
		// wait for some interval
		time.Sleep(time.Second * searchInterval)
	}
}

type searchPage struct {
	Kind  string
	Items []imgResult
}

type imgResult struct {
	Link string
	Mime string
}

func imgGet(cin chan string, googKey map[string]string) {
	for {
		// get query
		msg := <-cin
		// make http request
		urlTemplate := "%s?key=%s&cx=%s&q=%s&searchType=image&fields=kind,items(link,mime)"
		reqUrl := fmt.Sprintf(urlTemplate, googKey["address"], googKey["key"], googKey["id"], msg)
		fmt.Println(reqUrl)
		// send it to google
		client := &http.Client{
			Timeout: time.Second * 15,
		}
		resp, err := client.Get(reqUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println("received")
		// process response
		defer resp.Body.Close()
		bodyJson, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("read")
		fmt.Println(string(bodyJson))
		var body searchPage
		json.Unmarshal(bodyJson, &body)
		fmt.Println("decoded")
		fmt.Println(body.Items[0].Link)
	}
}
