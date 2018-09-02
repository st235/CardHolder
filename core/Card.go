package core

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var Cards CardArray

type Card struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Emoji string `json:"emoji"`
	Description string `json:"description"`
	Background string `json:"background"`
	Tags []string `json:"tags"`
}

type CardArray struct {
	CardsArray []Card `json:"cards"`
}

func Parse(filename string) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &Cards)

	defer jsonFile.Close()
}
