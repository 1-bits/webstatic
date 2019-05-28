package webstatic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Web struct {
	title string `json:"title"`
}

func getSiteTitle() string {
	jsonFile, err := os.Open("web.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var unaweb Web

	json.Unmarshal(byteValue, &unaweb)
	return unaweb.title
}
